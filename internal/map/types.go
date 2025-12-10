package mapval

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/cockroachdb/apd/v3"
)

// ---------------------------------------------------------
// 1. The Value Struct (Stack Allocated / Passed by Value)
// ---------------------------------------------------------

type ValueType uint8

const (
	ValInteger ValueType = iota
	ValFloat
	ValDecimal  // apd.Decimal (Heap Object)
	ValText     // Standard string
	ValObject   // Heap Object (Map, Func, etc.)
	ValEmpty    // The ___ literal
	ValBoolean  // Packed in Integer
	ValDateTime // Packed in Integer (Unix Milliseconds)
	ValDuration // Packed in Integer (Milliseconds)
	ValRange    // Range Object
	ValVersion  // Packed SemVer
	ValKey      // Interned Global ID
)

type Value struct {
	Type ValueType

	// Slot A: Packed Integers (Int64, Date, Version, Key, Bool)
	Integer int64

	// Slot B: Dedicated Float (Performance optimization)
	Float float64

	// Slot C: String Header (Native Go string)
	Str string

	// Slot D: Heap Pointer (Interface avoids "Fat Struct")
	Obj Object
}

func (v Value) String() string {
	switch v.Type {
	case ValInteger:
		return fmt.Sprintf("%d", v.Integer)
	case ValFloat:
		// Suppress trailing .0 for integers represented as floats (e.g., 5.0 -> 5)
		s := fmt.Sprintf("%f", v.Float)
		s = strings.TrimSuffix(s, ".000000") // Default Go float format
		// If after trimming it becomes just "5", it's fine.
		return s
	case ValDecimal:
		if d, ok := v.Obj.(*Decimal); ok {
			return d.String()
		}
		return "0.0"
	case ValText:
		return fmt.Sprintf("'%s'", v.Str)
	case ValBoolean:
		if v.Integer == 1 {
			return "yes"
		}
		return "no"
	case ValEmpty:
		return "___"
	case ValDateTime:
		t := time.UnixMilli(v.Integer).UTC()
		
		// If value is exactly 0 (Epoch at Midnight), prints as "1970/01/01"
		if v.Integer == 0 {
			return "1970/01/01"
		}

		s := ""
		// Midnight Suppression: If the time is 00:00:00.000, the time part is hidden.
		if t.Hour() == 0 && t.Minute() == 0 && t.Second() == 0 && t.Nanosecond() == 0 {
			s = t.Format("2006/01/02")
		} else {
			s = t.Format("2006/01/02@15:04:05")
			if t.Nanosecond() > 0 {
				s += fmt.Sprintf(".%03d", t.Nanosecond()/1000000) // Add milliseconds
			}
		}

		// Epoch Suppression: If the date is 1970/01/01, the date part is hidden.
		if t.Year() == 1970 && t.Month() == time.January && t.Day() == 1 {
			if s == "1970/01/01" { 
				// This case is already covered by v.Integer == 0, where it prints "1970/01/01"
				// but if it's 1970/01/01@time, we want just the time
				return strings.TrimPrefix(s, "1970/01/01@")
			}
		}
		return s
	case ValDuration:
		sign := "+"
		ms := v.Integer
		if ms < 0 {
			sign = "-"
			ms = -ms
		}

		const DayMS = 24 * 3600 * 1000
		const YearMS = 365 * DayMS
		const MonthMS = 30 * DayMS // Simplistic month

		// Waterfall Normalization & Zero-Suppression
		years := ms / YearMS
		rem := ms % YearMS
		months := rem / MonthMS
		rem %= MonthMS
		days := rem / DayMS
		rem %= DayMS

		var dateStr string
		if years > 0 || months > 0 || days > 0 {
			dateStr = fmt.Sprintf("%04d/%02d/%02d", years, months, days)
		}

		// Time part
		hours := rem / (3600 * 1000)
		rem %= (3600 * 1000)
		minutes := rem / (60 * 1000)
		rem %= (60 * 1000)
		seconds := rem / 1000
		millis := rem % 1000

		var timeStr string
		if hours > 0 || minutes > 0 || seconds > 0 || millis > 0 {
			timeStr = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
			if millis > 0 {
				timeStr += fmt.Sprintf(".%03d", millis)
			}
		}

		// Combine with Zero-Suppression
		s := ""
		if dateStr != "" && timeStr != "" {
			s = dateStr + "@" + timeStr 
		} else if dateStr != "" {
			s = dateStr
		} else if timeStr != "" {
			s = timeStr
		} else { // Exactly 0ms
			return "+00:00:00" // Ensure leading + for zero duration
		}
		return sign + s
	case ValRange:
		return "<Range>"
	case ValObject:
		if v.Obj == nil {
			return "nil"
		}
		if m, ok := v.Obj.(*Map); ok {
			return m.String()
		}
		if t, ok := v.Obj.(*Tuple); ok {
			kind := ""
			switch t.Kind {
			case TupleSignal:
				kind = "#"
			case TupleReply:
				kind = "^"
			case TupleProclamation:
				kind = "$"
			}
			return fmt.Sprintf("<%s%s>", kind, t.Topic)
		}
		if f, ok := v.Obj.(*Function); ok {
			return fmt.Sprintf("<%s>", f.Name)
		}
		if c, ok := v.Obj.(*Closure); ok {
			return fmt.Sprintf("<%s>", c.Fn.Name)
		}
		return fmt.Sprintf("<Object %T>", v.Obj)
	case ValVersion:
		maj, min, pat, wild := v.VersionUnpack()
		s := ""
		if wild {
			// Zero-suppression for wildcard: don't print 0.0 if wildcard exists
			if pat == 0 {
				if min == 0 {
					s = fmt.Sprintf("%d.-", maj)
				} else {
					s = fmt.Sprintf("%d.%d.-", maj, min)
				}
			} else { // This case should not be hit if parsing correctly sets patch to 0 for x.y.-
				s = fmt.Sprintf("%d.%d.%d.-", maj, min, pat)
			}
		} else {
			s = fmt.Sprintf("%d.%d.%d", maj, min, pat)
		}

		// Suffix
		if v.Str != "" {
			s += v.Str
		}
		return s
	case ValKey:
		return fmt.Sprintf(":key(%d)", v.Integer)
	default:
		return fmt.Sprintf("?%d", v.Type)
	}
}

// ---------------------------------------------------------
// 2. The Heap Objects (Interfaces)
// ---------------------------------------------------------

type ObjectType uint8

const (
	ObjTypeMap ObjectType = iota
	ObjTypeFunction
	ObjTypeClosure
	ObjTypeNative
	ObjTypeRange
	ObjTypeTuple
	ObjTypeDecimal
)

// Decimal wrapper
type Decimal struct {
	D        *apd.Decimal
	Original string // Preserves strict formatting (e.g. 01.5)
}

func (d *Decimal) Type() ObjectType { return ObjTypeDecimal }
func (d *Decimal) String() string {
	if d.Original != "" {
		return d.Original
	}
	if d.D == nil {
		return "0.0"
	}
	s := d.D.String()
	// apd sometimes outputs .0 for integers. Remove it if no fractional part.
	s = strings.TrimSuffix(s, ".0")
	return s
}

// ... rest of file (Tuple, Range, etc.)

// TupleKind
type TupleKind uint8

const (
	TupleSignal       TupleKind = iota // #
	TupleReply                         // ^
	TupleProclamation                  // $
)

// Tuple represents a message/event (Signal, Reply, Proclamation).
type Tuple struct {
	Kind    TupleKind
	Topic   string
	Payload []Value
	Source  interface{} // Opaque Source Frame (for Signals/Replies)
}

func (t *Tuple) Type() ObjectType { return ObjTypeTuple }

// Range represents a lazy sequence.
type Range struct {
	Start int64
	End   int64
}

func (r *Range) Type() ObjectType { return ObjTypeRange }

// ---------------------------------------------------------

// Object is the common interface for heap entities
type Object interface {
	Type() ObjectType
}

// ---------------------------------------------------------
// 3. Concrete Implementations
// ---------------------------------------------------------

// Map represents the Self-style Object
type Map struct {
	Legend    *Legend
	Fields    []Value // Linear storage matching the Legend offsets
	Listeners []Value // List of Selectors (Listeners)
}

func (m *Map) Type() ObjectType { return ObjTypeMap }

func (m *Map) String() string {
	var positionals []int
	posMap := make(map[int]Value)
	var named []FieldDesc

	if m.Legend == nil {
		return "[]"
	}

	for i, fieldDesc := range m.Legend.Fields {
		// Check if positional
		if idx, err := strconv.Atoi(fieldDesc.Name); err == nil && idx > 0 {
			positionals = append(positionals, idx)
			posMap[idx] = m.Fields[i]
		} else {
			named = append(named, fieldDesc)
		}
	}

	sort.Ints(positionals)

	var parts []string

	// 1. Positional Values
	for _, idx := range positionals {
		parts = append(parts, posMap[idx].String())
	}

	// 2. Named Fields (Names Only)
	for _, field := range named {
		// Filter Keys (starts with backtick)
		if strings.HasPrefix(field.Name, "`") {
			continue
		}

		prefix := "."
		if field.Kind == FieldMutable {
			prefix = ":"
		}

		parts = append(parts, prefix+field.Name)
	}

	return "[" + strings.Join(parts, "; ") + "]"
}

// FieldKind distinguishes between immutable (.) and mutable (:) fields.
type FieldKind uint8

const (
	FieldImmutable FieldKind = iota // Defined via '.'
	FieldMutable                    // Defined via ':'
)

// FieldDesc describes a single field in a Legend.
type FieldDesc struct {
	Name string    // The label/key of the field
	Kind FieldKind // Mutable or Immutable
}

// LegendType distinguishes between Map (Linear), Dictionary (Hash), etc.
type LegendType uint8

const (
	LegendMap        LegendType = iota // Linear scan
	LegendDictionary                   // Hash map
)

// Legend represents the Schema of a Map (Hidden Class).
type Legend struct {
	Kind   LegendType
	Fields []FieldDesc    // Linear scan (Map Mode)
	Lookup map[string]int // Hash map (Dictionary Mode)
}

// Function represents compiled bytecode (Prototype Activation Record)
type Function struct {
	Name         string
	Arity        int
	UpvalueCount int
	Chunk        *Chunk
}

func (f *Function) Type() ObjectType { return ObjTypeFunction }

// Closure represents a running instance of a function
type Closure struct {
	Fn       *Function
	Upvalues []*Upvalue
}

func (c *Closure) Type() ObjectType { return ObjTypeClosure }

// NativeFunction represents a native Go function callable from Rhumb
type NativeFunction struct {
	Fn func(args []Value) Value
}

func (n *NativeFunction) Type() ObjectType { return ObjTypeNative }

// ---------------------------------------------------------
// 4. "Code is Data" Definitions
// ---------------------------------------------------------

// Chunk belongs here because it is a data structure referenced by Objects.
// The VM package will import this, not the other way around.
type Chunk struct {
	Code      []byte
	Constants []Value
	Lines     []int
}

// NewChunk creates a new empty chunk.
func NewChunk() *Chunk {
	return &Chunk{
		Code:      make([]byte, 0),
		Constants: make([]Value, 0),
		Lines:     make([]int, 0),
	}
}

// WriteByte appends a byte to the code.
func (c *Chunk) WriteByte(b byte, line int) {
	c.Code = append(c.Code, b)
	c.Lines = append(c.Lines, line)
}

// WriteOp appends an OpCode to the code.
func (c *Chunk) WriteOp(op OpCode, line int) {
	c.Code = append(c.Code, byte(op))
	c.Lines = append(c.Lines, line)
}

// AddConstant adds a value to the constants pool and returns its index.
func (c *Chunk) AddConstant(val Value) int {
	c.Constants = append(c.Constants, val)
	return len(c.Constants) - 1
}

type Upvalue struct {
	Location *Value
	Closed   Value
	Next     *Upvalue
}

// ---------------------------------------------------------
// 5. Value Helper Constructors
// ---------------------------------------------------------

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

func NewFunction(f *Function) Value {
	return Value{Type: ValObject, Obj: f}
}

func NewSignal(topic string, source interface{}, args []Value) Value {
	return Value{Type: ValObject, Obj: &Tuple{
		Kind:    TupleSignal,
		Topic:   topic,
		Source:  source,
		Payload: args,
	}}
}

func NewVersion(major, minor uint16, patch uint32, wildcard bool) Value {
	// Use 63 bits. MSB (bit 63) is Wildcard.
	// Major: 15 bits (0-32767).
	// Minor: 16 bits.
	// Patch: 32 bits.
	
	encoded := int64(major&0x7FFF) << 48
	encoded |= int64(minor) << 32
	encoded |= int64(patch)
	
	if wildcard {
		// Set MSB (Bit 63) - which makes it negative in int64, but that's fine for storage
		encoded |= -9223372036854775808 
	}
	
	return Value{Type: ValVersion, Integer: encoded}
}

func NewKey(id int64) Value {

	return Value{Type: ValKey, Integer: id}

}

// VersionUnpack returns the major, minor, patch, and wildcard status.
func (v Value) VersionUnpack() (uint16, uint16, uint32, bool) {

	if v.Type != ValVersion {
		return 0, 0, 0, false
	}

	raw := uint64(v.Integer)
	
	// MSB is Wildcard
	wildcard := (raw & (1 << 63)) != 0
	
	// Major is bits 48-62 (15 bits)
	major := uint16((raw >> 48) & 0x7FFF)

	minor := uint16((raw >> 32) & 0xFFFF)

	patch := uint32(raw & 0xFFFFFFFF)

	return major, minor, patch, wildcard
}

// KeyID returns the unique ID of a Key value.

func (v Value) KeyID() int64 {

	if v.Type != ValKey {

		return -1

	}

	return v.Integer

}
