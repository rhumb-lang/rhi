package config

// Config holds the configuration for the Rhumb runtime/compiler.
type Config struct {
	TraceParser   bool
	TraceBytecode bool
	TraceStack    bool
}

// DefaultConfig returns the default configuration.
func DefaultConfig() *Config {
	return &Config{
		TraceParser:   false,
		TraceBytecode: false,
		TraceStack:    false,
	}
}
