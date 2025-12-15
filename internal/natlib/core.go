package natlib

import (
	mapval "github.com/rhumb-lang/rhi/internal/map"
)

func init() {
	// This would typically be loaded by the standard library
	Register("nat_core", map[string]mapval.Value{
		"***": mapval.NewNativeFunc("***", natPanic), // Rhumb's panic function
	})
}

// natPanic simulates a Rhumb panic by returning an Error Signal.
// The test runner would ideally detect this signal and fail the test.
func natPanic(args []mapval.Value) mapval.Value {
	msg := ""
	if len(args) > 0 && args[0].Type == mapval.ValText {
		msg = args[0].Str
	} else if len(args) > 0 {
		msg = args[0].Canonical()
	}

	// Code 999: General panic/unhandled error
	return mapval.NewErrorSignal(ErrPanic, "Rhumb Panic: "+msg, mapval.NewEmpty())
}
