package compiler

import (
	"fmt"
	
	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

func (c *Compiler) compileExpression(expr ast.Expression) error {
	switch e := expr.(type) {
	case *ast.BinaryExpression:
		return c.compileBinary(e)
	case *ast.CallExpression:
		return c.compileCall(e)
	case *ast.IntegerLiteral:
		c.emitConstant(mapval.NewInt(e.Value))
	case *ast.RationalLiteral:
		c.emitConstant(mapval.NewFloat(e.Value))
	case *ast.LabelLiteral:
		// Variable Access
		idx := c.Scope.resolveLocal(e.Value)
		if idx != -1 {
			c.emit(mapval.OP_LOAD_LOC)
			c.Chunk().WriteByte(byte(idx), 0)
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
	case *ast.RoutineExpression:
		return c.compileRoutine(e)
	case *ast.ChainExpression:
		return c.compileChain(e)
	case *ast.SelectorExpression:
		return c.compileSelector(e)
	case *ast.AssertionWrapper:
		// Compile Actual
		if err := c.compileExpression(e.Actual); err != nil {
			return err
		}
		// Dup Actual so the program continues with the value
		c.emit(mapval.OP_DUP)
		// Compile Expected
		if err := c.compileExpression(e.Expected); err != nil {
			return err
		}
		// Assert (Pops Expected, Actual)
		c.emit(mapval.OP_ASSERT_EQ)
	default:
		return fmt.Errorf("unsupported expression type: %T", expr)
	}
	return nil
}
