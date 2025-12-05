package compiler

import (
	"fmt"
	
	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

// Compiler transforms AST into Bytecode.
type Compiler struct {
	Chunk *vm.Chunk
	Scope *CompilerScope
}

func NewCompiler() *Compiler {
	return &Compiler{
		Chunk: vm.NewChunk(),
		Scope: NewCompilerScope(),
	}
}

// Compile processes a document and returns the resulting chunk.
func (c *Compiler) Compile(doc *ast.Document) (*vm.Chunk, error) {
	for _, expr := range doc.Expressions {
		if err := c.compileExpression(expr); err != nil {
			return nil, err
		}
	}
	c.emit(vm.OP_HALT)
	return c.Chunk, nil
}

func (c *Compiler) emit(op vm.OpCode) {
	// TODO: Track line numbers from AST
	c.Chunk.WriteOp(op, 0)
}

func (c *Compiler) emitBytes(b1, b2 byte) {
	c.Chunk.WriteByte(b1, 0)
	c.Chunk.WriteByte(b2, 0)
}

// emitJump emits a jump instruction with a placeholder operand and returns the offset to patch.
func (c *Compiler) emitJump(op vm.OpCode) int {
	c.emit(op)
	c.emitBytes(0xFF, 0xFF) // Placeholder 2-byte operand
	return len(c.Chunk.Code) - 2 // Return start of operand
}

// patchJump patches the operand of a jump instruction at the given offset.
func (c *Compiler) patchJump(offset int) {
	jump := len(c.Chunk.Code) - offset - 2 // Calculate actual jump distance
	
	if jump > 0xFFFF { // Max 2 bytes (65535)
		panic("Jump offset too large")
	}
	
	c.Chunk.Code[offset] = byte((jump >> 8) & 0xFF)
	c.Chunk.Code[offset+1] = byte(jump & 0xFF)
}


func (c *Compiler) makeConstant(val mapval.Value) int {
	return c.Chunk.AddConstant(val)
}

func (c *Compiler) emitConstant(val mapval.Value) {
	idx := c.makeConstant(val)
	if idx > 255 {
		// TODO: Handle OpConstantLong if > 255
		panic("Too many constants")
	}
	c.emit(vm.OP_LOAD_CONST)
	c.Chunk.WriteByte(byte(idx), 0)
}

func (c *Compiler) compileExpression(expr ast.Expression) error {
	switch e := expr.(type) {
	case *ast.BinaryExpression:
		return c.compileBinary(e)
	case *ast.IntegerLiteral:
		c.emitConstant(mapval.NewInt(e.Value))
	case *ast.RationalLiteral:
		c.emitConstant(mapval.NewFloat(e.Value))
	case *ast.LabelLiteral:
		// Variable Access
		idx := c.Scope.resolveLocal(e.Value)
		if idx != -1 {
			c.emit(vm.OP_LOAD_LOC)
			c.Chunk.WriteByte(byte(idx), 0)
		} else {
			return fmt.Errorf("undefined variable: %s", e.Value)
		}
	case *ast.TextLiteral:
		// Simple text support for now (no interp)
		if len(e.Segments) == 1 {
			if s, ok := e.Segments[0].(*ast.StringSegment); ok {
				c.emitConstant(mapval.NewText(s.Value))
			} else {
				return fmt.Errorf("interpolation not yet supported in compiler")
			}
		}
	case *ast.EmptyLiteral:
		c.emitConstant(mapval.NewEmpty())
	case *ast.MapExpression:
		return c.compileMap(e)
	default:
		return fmt.Errorf("unsupported expression type: %T", expr)
	}
	return nil
}

func (c *Compiler) compileBinary(bin *ast.BinaryExpression) error {
	// Special handling for Assignment
	if bin.Op == ast.OpAssignImm || bin.Op == ast.OpAssignMut {
		// Check LHS
		if label, ok := bin.Left.(*ast.LabelLiteral); ok {
			// Variable Assignment: x .= 1
			
			// 1. Compile RHS
			if err := c.compileExpression(bin.Right); err != nil {
				return err
			}
			
			// 2. Declare/Resolve Variable
			// Check if it exists in current scope (or any scope for update)
			idx := c.Scope.resolveLocal(label.Value)
			
			if idx != -1 {
				// Update existing
				// TODO: Check if immutable in scope metadata
				c.emit(vm.OP_STORE_LOC)
				c.Chunk.WriteByte(byte(idx), 0)
			} else {
				// Define new
				idx = c.Scope.addLocal(label.Value)
				c.emit(vm.OP_STORE_LOC)
				c.Chunk.WriteByte(byte(idx), 0)
			}
			return nil
		}
		// Fallthrough for Map assignment (e.g. obj.x .= 1)
	}

	// Control Flow
	if bin.Op == ast.OpIfTrue || bin.Op == ast.OpIfFalse {
		// Compile condition (LHS)
		if err := c.compileExpression(bin.Left); err != nil {
			return err
		}
		
		// Preserve condition for chaining
		c.emit(vm.OP_DUP)
		
		// Determine Jump Op
		// => (IfTrue): Jump if False (Skip body) -> OP_IF_TRUE (as implemented in VM: jumps if falsy)
		// ~> (IfFalse): Jump if True (Skip body) -> OP_IF_FALSE (as implemented in VM: jumps if truthy)
		jumpOp := vm.OP_IF_TRUE
		if bin.Op == ast.OpIfFalse {
			jumpOp = vm.OP_IF_FALSE
		}
		
		skipBody := c.emitJump(jumpOp)
		
		// Compile Body (RHS)
		if err := c.compileExpression(bin.Right); err != nil {
			return err
		}
		
		// Discard body result, return condition (which is under body result)
		c.emit(vm.OP_POP)
		
		// Target for skip
		c.patchJump(skipBody)
		return nil
	}

	// Compile Left
	if err := c.compileExpression(bin.Left); err != nil {
		return err
	}
	// Compile Right
	if err := c.compileExpression(bin.Right); err != nil {
		return err
	}
	
	// Emit Op
	switch bin.Op {
	case ast.OpAdd: c.emit(vm.OP_ADD)
	case ast.OpSub: c.emit(vm.OP_SUB)
	case ast.OpMult: c.emit(vm.OP_MULT)
	case ast.OpDivFloat: c.emit(vm.OP_DIV_FLOAT)
	case ast.OpAssignImm: c.emit(vm.OP_ASSIGN_IMM) // For map fields
	case ast.OpAssignMut: c.emit(vm.OP_ASSIGN_MUT) // For map fields
	case ast.OpEq: c.emit(vm.OP_EQ)
	case ast.OpNeq: c.emit(vm.OP_NEQ)
	case ast.OpGt: c.emit(vm.OP_GT)
	case ast.OpLt: c.emit(vm.OP_LT)
	case ast.OpGte: c.emit(vm.OP_GTE)
	case ast.OpLte: c.emit(vm.OP_LTE)
	default:
		return fmt.Errorf("unsupported binary op: %v", bin.Op)
	}
	return nil
}

func (c *Compiler) compileMap(m *ast.MapExpression) error {
	// For now, just create empty map
	c.emit(vm.OP_MAKE_MAP)
	// TODO: Compile fields
	return nil
}
