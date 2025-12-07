package vm

import (
	"fmt"

	mapval "git.sr.ht/~madcapjake/rhi/internal/map"
)

func (vm *VM) isFalsy(val mapval.Value) bool {
	return val.Type == mapval.ValEmpty || (val.Type == mapval.ValBoolean && val.Integer == 0)
}

func (vm *VM) isTruthy(val mapval.Value) bool {
	return !vm.isFalsy(val)
}

func asFloat(v mapval.Value) float64 {
	if v.Type == mapval.ValInteger {
		return float64(v.Integer)
	}
	if v.Type == mapval.ValFloat {
		return v.Float
	}
	// TODO: Handle other types that can coerce to float
	return 0.0 // Default or error
}

// isEqual compares two values for equality.
// Handles different types based on Rhumb semantics.
func isEqual(a, b mapval.Value) bool {
	if a.Type != b.Type {
		// Attempt coercion for numbers
		if (a.Type == mapval.ValInteger || a.Type == mapval.ValFloat) && (b.Type == mapval.ValInteger || b.Type == mapval.ValFloat) {
			return asFloat(a) == asFloat(b)
		}
		return false
	}

	switch a.Type {
	case mapval.ValInteger:
		return a.Integer == b.Integer
	case mapval.ValFloat:
		return a.Float == b.Float
	case mapval.ValBoolean:
		return a.Integer == b.Integer
	case mapval.ValEmpty:
		return true
	case mapval.ValText:
		return a.Str == b.Str
	// TODO: Keys, Dates, Versions, Objects
	default:
		return false
	}
}

// numericCompare compares two numeric values. Returns -1 if a<b, 0 if a==b, 1 if a>b.
func numericCompare(a, b mapval.Value) (int, error) {
	if (a.Type != mapval.ValInteger && a.Type != mapval.ValFloat) || (b.Type != mapval.ValInteger && b.Type != mapval.ValFloat) {
		return 0, fmt.Errorf("operands must be numbers for comparison")
	}

	fa := asFloat(a)
	fb := asFloat(b)

	if fa < fb {
		return -1, nil
	}
	if fa > fb {
		return 1, nil
	}
	return 0, nil
}
