package loader

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
