package vm

import "git.sr.ht/~madcapjake/rhi/internal/map"

// Chunk represents a block of bytecode and its associated constants.
type Chunk struct {
	Code      []byte
	Constants []mapval.Value
	Lines     []int // Line number for each byte (for debugging)
}

// NewChunk creates a new empty chunk.
func NewChunk() *Chunk {
	return &Chunk{
		Code:      make([]byte, 0),
		Constants: make([]mapval.Value, 0),
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
func (c *Chunk) AddConstant(val mapval.Value) int {
	c.Constants = append(c.Constants, val)
	return len(c.Constants) - 1
}
