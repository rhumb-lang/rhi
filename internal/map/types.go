package mapval

// ---------------------------------------------------------
// 1. The Value Struct (Stack Allocated / Passed by Value)
// ---------------------------------------------------------

type ValueType uint8

const (
	ValInteger ValueType = iota
	ValFloat             // The "Separate Slot" from architecture discussion
	ValText
	ValObject            // Heap Object (Map, Func, etc.)
	ValEmpty             // The ___ literal
	ValBoolean           // Packed in Integer
	ValDate              // Packed in Integer
	ValRange             // Range Object
	ValVersion           // Packed SemVer
	ValKey               // Interned Global ID
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
)

// ...

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
	Legend *Legend
	Fields []Value // Linear storage matching the Legend offsets
}

func (m *Map) Type() ObjectType { return ObjTypeMap }

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
	Kind        LegendType
	Fields      []FieldDesc       // Linear scan (Map Mode)
	Lookup      map[string]int    // Hash map (Dictionary Mode)
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

func NewVersion(major, minor uint16, patch uint32) Value {
	encoded := int64(major) << 48
	encoded |= int64(minor) << 32
	encoded |= int64(patch)
	return Value{Type: ValVersion, Integer: encoded}
}

func NewKey(id int64) Value {
	return Value{Type: ValKey, Integer: id}
}