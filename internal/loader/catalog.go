package loader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// ResourceMeta holds the parsed symbolic options for a file
type ResourceMeta struct {
	Mime     string
	Charset  string
	Checksum string   // "sha256:..."
	Options  []string // Raw options like "base64", "text/plain"
}

// DependencySpec handles both Code (String) and Resource (Array) deps
type DependencySpec struct {
	Version       string
	IsResource    bool
	InlineCatalog map[string]ResourceMeta // For inline definitions: [{ "file": "meta" }]
}

type Catalog struct {
	Author      string   `yaml:"👤"`
	License     string   `yaml:"🪪"`
	Repository  string   `yaml:"📦"`
	Keywords    []string `yaml:"🏷️"`
	Description string   `yaml:"📝"`
	SourceRoot  string   `yaml:"📂"`

	// Engines defines the required runtime versions
	// Key: "rhi" (or potentially others like "go" if strictly needed)
	// Value: Version with or without a wildcard (e.g. "1.-", "0.5.-", or "0.5.1")
	Engines map[string]string `yaml:"⚙️"`

	// Versions maps "0.1.0" -> VersionConfig
	// Note: Since YAML keys are dynamic strings, we parse into a map[string]interface{}
	// and manually decode the version keys vs metadata keys.
	Versions map[string]VersionConfig `yaml:"-"`
}

type VersionConfig struct {
	Dependencies map[string]DependencySpec
	Alias        string
	// Resources is populated ONLY if the catalog file itself is bracketed [name] @tests/fixtures/project_alpha/project_alpha@.rhy
	Resources map[string]ResourceMeta
}

func LoadCatalog(path string) (*Catalog, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Decode into raw map
	var raw map[string]interface{}
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	m := &Catalog{
		Versions: make(map[string]VersionConfig),
	}

	knownKeys := map[string]bool{
		"👤": true, "🪪": true, "📦": true, "🏷️": true, "📝": true, "📂": true, "⚙️": true,
	}

	// 1. Check for Flat Metadata
	hasFlatMetadata := false
	for k := range raw {
		if knownKeys[k] {
			hasFlatMetadata = true
			break
		}
	}

	if hasFlatMetadata {
		if err := yaml.Unmarshal(data, m); err != nil {
			return nil, err
		}
	}

	// Detect Resource Catalog
	filename := filepath.Base(path)
	isResourceCatalog := strings.Contains(filename, "[") && strings.Contains(filename, "]")

	// 2. Scan for Nested Metadata (if not flat) & Versions
	for key, val := range raw {
		if knownKeys[key] {
			continue // Already handled (if flat)
		}

		// Check if it's a Metadata Block (Project Name)
		isVersion := key == "-" || strings.Contains(key, ".") || (len(key) > 0 && key[0] >= '0' && key[0] <= '9')

		if !isVersion {
			if vMap, ok := val.(map[string]interface{}); ok {
				// Check for metadata keys inside
				isMetaBlock := false
				for subK := range vMap {
					if knownKeys[subK] {
						isMetaBlock = true
						break
					}
				}

				if isMetaBlock && !hasFlatMetadata {
					// Extract metadata from this nested map
					// Re-marshal and unmarshal to Catalog
					subData, _ := yaml.Marshal(val)
					yaml.Unmarshal(subData, m)
					continue
				}
			}
		}

		// It's a version (or treated as one)!
		vc := VersionConfig{
			Dependencies: make(map[string]DependencySpec),
			Resources:    make(map[string]ResourceMeta),
		}

		if vMap, ok := val.(map[string]interface{}); ok {
			for k, v := range vMap {
				if k == "<-" {
					if s, ok := v.(string); ok {
						vc.Alias = s
					}
					continue
				}

				valStr := fmt.Sprintf("%v", v)

				// MODE A: Resource Catalog
				if isResourceCatalog {
					vc.Resources[parseFilename(k)] = parseResourceMeta(k, valStr)
					continue
				}

				// MODE B: Standard Catalog
				spec := DependencySpec{}

				switch typedVal := v.(type) {
				case string:
					// Code Dependency: "math": "1.0.0"
					spec.Version = typedVal
					spec.IsResource = false

				case []interface{}:
					// Resource Dependency: "assets": ["1.0.0"]
					spec.IsResource = true
					spec.Version = "-" // Default

					if len(typedVal) > 0 {
						// Check first item to determine type
						first := typedVal[0]
						if verStr, ok := first.(string); ok {
							spec.Version = verStr
						} else {
							// Inline Catalog
							spec.InlineCatalog = make(map[string]ResourceMeta)
							for _, item := range typedVal {
								if mapDef, ok := item.(map[string]interface{}); ok {
									for fKey, fVal := range mapDef {
										fMeta := parseResourceMeta(fKey, fmt.Sprintf("%v", fVal))
										spec.InlineCatalog[parseFilename(fKey)] = fMeta
									}
								}
							}
						}
					}
				}
				vc.Dependencies[k] = spec
			}
		}

		m.Versions[key] = vc
	}

	// 3. Validation & Flattening
	flattened := make(map[string]bool)
	checking := make(map[string]bool)

	var flatten func(ver string) error
	flatten = func(ver string) error {
		if flattened[ver] {
			return nil
		}
		if checking[ver] {
			return fmt.Errorf("cycle detected in version aliases: %s", ver)
		}

		vc, exists := m.Versions[ver]
		if !exists {
			return nil
		}

		checking[ver] = true
		defer func() { delete(checking, ver) }()

		if vc.Alias != "" {
			baseVer := vc.Alias
			if _, ok := m.Versions[baseVer]; !ok {
				return fmt.Errorf("version %s points to missing base version %s", ver, baseVer)
			}

			if err := flatten(baseVer); err != nil {
				return err
			}

			baseVC := m.Versions[baseVer]

			// Validate Shadowing
			for key := range vc.Dependencies {
				if _, exists := baseVC.Dependencies[key]; exists {
					return fmt.Errorf("pointer Invalid: Shadowing Detected in %s overriding %s from %s", ver, key, baseVer)
				}
			}
			for key := range vc.Resources {
				if _, exists := baseVC.Resources[key]; exists {
					return fmt.Errorf("pointer Invalid: Resource Shadowing Detected in %s overriding %s", ver, key)
				}
			}

			// Merge (Copy base dependencies to current)
			for k, v := range baseVC.Dependencies {
				vc.Dependencies[k] = v
			}
			for k, v := range baseVC.Resources {
				vc.Resources[k] = v
			}
		}

		flattened[ver] = true
		return nil
	}

	for v := range m.Versions {
		if err := flatten(v); err != nil {
			return nil, err
		}
	}

	if _, ok := m.Versions["-"]; !ok {
		return nil, fmt.Errorf("invalid catalog: missing tip version ('-')")
	}

	return m, nil
}

// Helper: "config.json;utf-8" -> "config.json"
func parseFilename(k string) string {
	if idx := strings.Index(k, ";"); idx != -1 {
		return k[:idx]
	}
	return k
}

// Helper: Parses key options and value checksums
func parseResourceMeta(k, v string) ResourceMeta {
	meta := ResourceMeta{}

	// 1. Parse Key Options (e.g. "file.png;base64")
	if idx := strings.Index(k, ";"); idx != -1 {
		meta.Options = append(meta.Options, strings.Split(k[idx+1:], ";")...)
	}

	// 2. Parse Value Tokens (e.g. "text/plain sha256:...")
	// We handle the value being potentially just "___" or a complex string
	if v != "" && v != "___" {
		tokens := strings.Fields(v) // Splits by whitespace
		for _, token := range tokens {
			if strings.HasPrefix(token, "sha256:") {
				meta.Checksum = token
			} else {
				// Treat non-checksum tokens as Options/MIME
				meta.Options = append(meta.Options, token)
			}
		}
	}

	// 3. Normalize Fields (MIME vs Charset)
	// Consolidate options from both Key and Value sources
	for _, opt := range meta.Options {
		if strings.Contains(opt, "/") {
			meta.Mime = opt
		} else if opt == "utf-8" || opt == "iso-8859-1" {
			meta.Charset = opt
		}
	}
	
	return meta
}
