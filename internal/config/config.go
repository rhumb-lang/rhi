package config

// Config holds the configuration for the Rhumb runtime/compiler.
type Config struct {
	TraceParser          bool
	TraceBytecode        bool
	TraceStack           bool
	TraceSpace           bool
	AllowUnsafeWildcards bool // If false, imports like {-|lib|-} are rejected
}

// DefaultConfig returns the default configuration.
func DefaultConfig() *Config {
	return &Config{
		TraceParser:          false,
		TraceBytecode:        false,
		TraceStack:           false,
		TraceSpace:           false,
		AllowUnsafeWildcards: false,
	}
}
