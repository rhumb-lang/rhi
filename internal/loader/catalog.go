package loader

import (
	"fmt"
	"os"

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

	// 1. Decode Metadata
	var m Manifest
	if err := yaml.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	m.Versions = make(map[string]VersionConfig)

	// 2. Decode Dynamic Versions
	var raw map[string]interface{}
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	knownKeys := map[string]bool{
		"👤": true, "🪪": true, "📦": true, "🏷️": true, "📝": true, "📂": true,
	}

	for key, val := range raw {
		if knownKeys[key] {
			continue
		}

		// It's a version!
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
					// Store boolean as string for now, or flag IsResource if appropriate
					// Assuming simple dependency map for now.
					vc.Dependencies[k] = fmt.Sprintf("%v", val)
					if val {
						vc.IsResource = true
					}
				}
			}
		}

		m.Versions[key] = vc
	}
	return &m, nil
}
