package natlib

import (
	"math"
	"math/rand"

	mapval "github.com/rhumb-lang/rhi/internal/map"
)

func init() {
	Register("nat_math", map[string]mapval.Value{
		"pi":  mapval.NewFloat(math.Pi),
		"sin": mapval.NewNativeFunc("sin", natSin),
		"int": mapval.NewNativeFunc("int", natInt),
	})
}

func natSin(args []mapval.Value) mapval.Value {
	if len(args) != 1 {
		return ErrInvalidArg("sin", "expects 1 argument")
	}

	// Type Coercion logic (Float/Int -> Float)
	val := args[0]
	f := 0.0
	switch val.Type {
	case mapval.ValFloat:
		f = val.Float
	case mapval.ValInteger:
		f = float64(val.Integer)
	default:
		return ErrInvalidArg("sin", "expects number")
	}

	return mapval.NewFloat(math.Sin(f))
}

func natInt(args []mapval.Value) mapval.Value {
	if len(args) != 2 {
		return ErrInvalidArg("int", "expects (min, max)")
	}

	min, ok1 := args[0].Integer, args[0].Type == mapval.ValInteger
	max, ok2 := args[1].Integer, args[1].Type == mapval.ValInteger

	if !ok1 || !ok2 {
		return ErrInvalidArg("int", "expects integers")
	}

	return mapval.NewInt(rand.Int63n(max-min) + min)
}
