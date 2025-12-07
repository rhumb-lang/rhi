package compiler

import (
	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

// Compiler transforms AST into Bytecode.
type Compiler struct {
	Enclosing *Compiler
	Function  *mapval.Function
	Scope     *CompilerScope
}

func NewCompiler() *Compiler {
	fn := &mapval.Function{
		Name:  "<script>",
		Chunk: mapval.NewChunk(),
	}
	return &Compiler{
		Function: fn,
		Scope:    NewCompilerScope(),
	}
}

func (c *Compiler) Chunk() *mapval.Chunk {
	return c.Function.Chunk
}

// Compile processes a document and returns the resulting chunk.
func (c *Compiler) Compile(doc *ast.Document) (*mapval.Chunk, error) {
	// Hoist locals
	hoister := NewHoister()
	locals := hoister.Hoist(doc)
	for _, name := range locals {
		c.Scope.addLocal(name)
		c.emitConstant(mapval.NewEmpty()) // Reserve slot
	}

	if len(doc.Expressions) == 0 {
		c.emitConstant(mapval.NewEmpty()) // Return empty if empty doc
	} else {
		for i, expr := range doc.Expressions {
			if err := c.compileExpression(expr); err != nil {
				return nil, err
			}
			// Pop result if not last
			if i < len(doc.Expressions)-1 {
				c.emit(mapval.OP_POP)
			}
		}
	}
	c.emit(mapval.OP_HALT)
	return c.Chunk(), nil
}

// CompileIncremental appends compiled code for the doc to the existing chunk.
// It overwrites the previous OP_HALT if present.
// Returns the start offset of the new code.
func (c *Compiler) CompileIncremental(doc *ast.Document) (int, error) {
	chunk := c.Chunk()
	if len(chunk.Code) > 0 && mapval.OpCode(chunk.Code[len(chunk.Code)-1]) == mapval.OP_HALT {
		chunk.Code = chunk.Code[:len(chunk.Code)-1]
	}
	startOffset := len(chunk.Code)
	
	// Hoist locals (append to existing scope)
	hoister := NewHoister()
	locals := hoister.Hoist(doc)
	for _, name := range locals {
		// addLocal checks duplicates?
		// My Scope.addLocal appends blindly?
		// Let's check scope.go.
		// "Search from top...".
		// addLocal just appends.
		// If we add same local again, it shadows (new slot).
		// In REPL, reusing "x" should probably refer to existing "x"?
		// If I define "x := 1". Then "x := 2".
		// Should it reuse slot?
		// CompilerScope.resolveLocal searches.
		// Hoister finds declarations.
		// If "x" declared again, `resolveLocal` finds OLD `x`?
		// No, `addLocal` adds NEW `x`. `resolveLocal` finds NEW `x` (top down).
		// But we want to reuse if global?
		// For REPL, shadowing is fine (consumes stack space but correct).
		// To reuse, we need `addLocal` to check existence in current scope depth?
		// My `Scope.addLocal` implementation (step 52) appends.
		// I should check if it exists in current scope depth (0 for script) and reuse?
		// Let's stick to append for simplicity (Shadowing).
		
		c.Scope.addLocal(name)
		c.emitConstant(mapval.NewEmpty()) // Reserve slot
	}

	for _, expr := range doc.Expressions {
		if err := c.compileExpression(expr); err != nil {
			return 0, err
		}
		// Pop results? For REPL, we usually WANT to print result.
		// So we keep the last one?
		// If we keep it, stack grows.
		// REPL prints top of stack.
		// So we should keep last.
		// And POP others.
		// Same logic as Compile?
		// But previous `Compile` run might have left a value on stack (last result of prev line).
		// We should POP that if we are running new line?
		// VM doesn't auto-pop.
		// If user types `1`. Stack `[1]`.
		// User types `2`. Stack `[1, 2]`.
		// This is fine.
	}
	c.emit(mapval.OP_HALT)
	return startOffset, nil
}

func (c *Compiler) emit(op mapval.OpCode) {
	c.Chunk().WriteOp(op, 0)
}

func (c *Compiler) emitBytes(b1, b2 byte) {
	c.Chunk().WriteByte(b1, 0)
	c.Chunk().WriteByte(b2, 0)
}

// emitJump emits a jump instruction with a placeholder operand and returns the offset to patch.
func (c *Compiler) emitJump(op mapval.OpCode) int {
	c.emit(op)
	c.emitBytes(0xFF, 0xFF) // Placeholder 2-byte operand
	return len(c.Chunk().Code) - 2 // Return start of operand
}

// patchJump patches the operand of a jump instruction at the given offset.
func (c *Compiler) patchJump(offset int) {
	jump := len(c.Chunk().Code) - offset - 2 // Calculate actual jump distance
	
	if jump > 0xFFFF { // Max 2 bytes (65535)
		panic("Jump offset too large")
	}
	
	c.Chunk().Code[offset] = byte((jump >> 8) & 0xFF)
		c.Chunk().Code[offset+1] = byte(jump & 0xFF)
	}
	
	// emitJumpBack emits a backward jump to the target offset.
	func (c *Compiler) emitJumpBack(op mapval.OpCode, target int) {
		// Offset = target - (current_len + 3)
		// Because IP will be at current_len + 3 after reading this instruction
		offset := target - (len(c.Chunk().Code) + 3)
		
		if offset < -32768 {
			panic("Loop body too large")
		}
		
		c.emit(op)
		c.Chunk().WriteByte(byte((offset >> 8) & 0xFF), 0)
		c.Chunk().WriteByte(byte(offset & 0xFF), 0)
	}
	
	func (c *Compiler) makeConstant(val mapval.Value) int {
	
	return c.Chunk().AddConstant(val)
}

func (c *Compiler) emitConstant(val mapval.Value) {
	idx := c.makeConstant(val)
	if idx > 255 {
		panic("Too many constants")
	}
	c.emit(mapval.OP_LOAD_CONST)
	c.Chunk().WriteByte(byte(idx), 0)
}
