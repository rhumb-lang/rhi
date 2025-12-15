package natlib

import (
	mapval "github.com/rhumb-lang/rhi/internal/map"
)

// Registry stores the Go-side implementations.
// Key: Library Name (e.g. "nat_math")
// Value: Map of exports
var Registry = make(map[string]map[string]mapval.Value)

func Register(name string, exports map[string]mapval.Value) {
	Registry[name] = exports
}
