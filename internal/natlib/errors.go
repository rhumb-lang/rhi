package natlib

import (
	"fmt"
	mapval "github.com/rhumb-lang/rhi/internal/map"
)

// Standard Error Codes
const (
	ErrBadRequest   = 400 // Invalid args, type mismatch, range error
	ErrUnauthorized = 401 // Permission denied
	ErrNotFound     = 404 // File/Key not found
	ErrTimeout      = 408 // Deadline exceeded
	ErrConflict     = 409 // Resource exists
	ErrInternal     = 500 // Generic Go error
	ErrNotImpl      = 501 // Not supported
	ErrPanic        = 999 // Explicit panic
)

// Helper for generic "Bad Argument" errors
func ErrInvalidArg(fnName string, reason string) mapval.Value {
	msg := fmt.Sprintf("%s: %s", fnName, reason)
	return mapval.NewErrorSignal(ErrBadRequest, msg, mapval.NewEmpty())
}

// Helper for wrapping arbitrary Go errors
func ErrNative(err error) mapval.Value {
	return mapval.NewErrorSignal(ErrInternal, err.Error(), mapval.NewEmpty())
}
