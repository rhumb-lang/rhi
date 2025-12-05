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
}

func NewCompiler() *Compiler {
	return &Compiler{
		Chunk: vm.NewChunk(),
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
	case ast.OpAssignImm: c.emit(vm.OP_ASSIGN_IMM)
	case ast.OpAssignMut: c.emit(vm.OP_ASSIGN_MUT)
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
