package loader

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Manifest struct {
	Author      string   `yaml:"👤"`
	License     string   `yaml:"🪪"`
	Repository  string   `yaml:"📦"`
	Keywords    []string `yaml:"🏷️"`
	Description string   `yaml:"📝"`
	SourceRoot  string   `yaml:"📂"`

	// Versions maps "0.1.0" -> VersionConfig
	// Note: Since YAML keys are dynamic strings, we parse into a map[string]interface{}
	// and manually decode the version keys vs metadata keys.
	Versions map[string]VersionConfig `yaml:"-"`
}

type VersionConfig struct {
	Dependencies map[string]string // "math": "1.0.0"
	Alias        string            // For "<-: 0.1.0"
	IsResource   bool              // For "assets: true"
}

func LoadCatalog(path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Decode into raw map
	var raw map[string]interface{}
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	m := &Manifest{
		Versions: make(map[string]VersionConfig),
	}

	knownKeys := map[string]bool{
		"👤": true, "🪪": true, "📦": true, "🏷️": true, "📝": true, "📂": true,
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

	// 2. Scan for Nested Metadata (if not flat) & Versions
	for key, val := range raw {
		if knownKeys[key] {
			continue // Already handled (if flat)
		}

		// Check if it's a Metadata Block (Project Name)
		// Heuristic: Not a version number, and value is a map containing metadata keys
		// Or assume any non-version key is metadata?
		// Version keys: "-", "0.1.0", "1.0"
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
					// Re-marshal and unmarshal to Manifest
					subData, _ := yaml.Marshal(val)
					yaml.Unmarshal(subData, m)
					continue
				}
			}
		}

		// It's a version (or treated as one)!
		vc := VersionConfig{
			Dependencies: make(map[string]string),
		}

		if vMap, ok := val.(map[string]interface{}); ok {
			for k, v := range vMap {
				if k == "<-" {
					if s, ok := v.(string); ok {
						vc.Alias = s
					}
					continue
				}
				
				switch val := v.(type) {
				case string:
					vc.Dependencies[k] = val
				case bool:
					vc.Dependencies[k] = fmt.Sprintf("%v", val)
					if val {
						vc.IsResource = true
					}
				}
			}
		}

		m.Versions[key] = vc
	}
	return m, nil
}
