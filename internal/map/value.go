package mapval

import "time"

// ValueType is the tag for the discriminated union.
type ValueType uint8

const (
	ValEmpty   ValueType = iota // ___
	ValInteger                  // int64
	ValFloat                    // float64
	ValText                     // string
	ValBoolean                  // packed in Integer (1 or 0)
	ValDate                     // packed in Integer (Unix Nano)
	ValVersion                  // packed in Integer (SemVer)
	ValKey                      // packed in Integer (Interned ID)
	ValMap                      // *Object
)

// Object represents the Map/Object structure.
// (This is a placeholder for now until we implement the full Legend/Map logic)
type Object struct {
	// TODO: Implement Legend and Fields storage
}

// Value is the stack-allocated primitive struct.
type Value struct {
	Type ValueType // 1 byte (Tag)

	// SLOT A: The General Purpose Integer Slot
	// We use this 64-bit space for Ints, but ALSO for
	// packed Booleans, Dates, Versions, and Keys.
	Integer int64

	// SLOT B: The Dedicated Float Slot
	// We use this 64-bit space ONLY when Type == ValFloat.
	// This is the "separate slot".
	Float float64

	// SLOT C: The String Header
	Str string

	// SLOT D: The Object Pointer
	Obj *Object
}

// Helper Constructors

func NewEmpty() Value {
	return Value{Type: ValEmpty}
}

func NewInt(i int64) Value {
	return Value{Type: ValInteger, Integer: i}
}

func NewFloat(f float64) Value {
	return Value{Type: ValFloat, Float: f}
}

func NewText(s string) Value {
	return Value{Type: ValText, Str: s}
}

func NewBoolean(b bool) Value {
	val := int64(0)
	if b {
		val = 1
	}
	return Value{Type: ValBoolean, Integer: val}
}

func NewDate(t time.Time) Value {
	return Value{Type: ValDate, Integer: t.UnixNano()}
}

// TODO: Implement Version packing logic (SemVer to int64)
// TODO: Implement Key interning logic (Global Symbol Table)
