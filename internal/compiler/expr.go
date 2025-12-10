package compiler

import (
	"fmt"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	mapval "git.sr.ht/~madcapjake/rhi/internal/map"
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
	case *ast.DateTimeLiteral:
		c.emitConstant(mapval.Value{
			Type:    mapval.ValDateTime,
			Integer: e.Value,
		})
	case *ast.VersionLiteral:
		val := mapval.NewVersion(e.Major, e.Minor, e.Patch, e.IsWildcard)
		val.Str = e.Suffix
		c.emitConstant(val)
	case *ast.DecimalLiteral:
		c.emitConstant(mapval.Value{
			Type: mapval.ValDecimal,
			Obj:  &mapval.Decimal{D: e.Value, Original: e.Original},
		})
	case *ast.LabelLiteral:
		// 1. Variable Access (Local)
		idx := c.Scope.resolveLocal(e.Value)
		if idx != -1 {
			c.emit(mapval.OP_LOAD_LOC)
			c.Chunk().WriteByte(byte(idx), 0)
			return nil
		}

		// 2. Variable Access (Upvalue)
		up := c.resolveUpvalue(e.Value)
		if up != -1 {
			c.emit(mapval.OP_LOAD_UPVALUE)
			c.Chunk().WriteByte(byte(up), 0)
			return nil
		}

		// 3. Special Boolean Literals
		// TODO: Add international support
		if e.Value == "yes" {
			c.emitConstant(mapval.NewBoolean(true))
			return nil
		}
		if e.Value == "no" {
			c.emitConstant(mapval.NewBoolean(false))
			return nil
		}

		// 4. Fallback
		return fmt.Errorf("undefined variable: %s", e.Value)
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
	case *ast.UnaryExpression:
		return c.compileUnary(e)
	case *ast.RoutineExpression:
		return c.compileRoutine(e)
	case *ast.ChildRealmLiteral:
		c.emit(mapval.OP_NEW_REALM)
		c.Chunk().WriteByte(0, 0) // Flag 0 for Child
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
		// Push Test Name
		c.emitConstant(mapval.NewText(e.Name))
		// Assert (Pops Name, Expected, Actual)
		c.emit(mapval.OP_ASSERT_EQ)
	case *ast.EffectExpression:
		// Compile Target (Closure)
		if err := c.compileExpression(e.Target); err != nil {
			return err
		}
		// Compile Selector
		if err := c.compileExpression(e.Selector); err != nil {
			return err
		}
		// Emit MONITOR
		c.emit(mapval.OP_MONITOR)
	case *ast.InspectionWrapper:
		// Compile Expression
		if err := c.compileExpression(e.Expr); err != nil {
			return err
		}
		// Dup so value remains for program
		c.emit(mapval.OP_DUP)
		// Inspect (Pops Value)
		c.emit(mapval.OP_INSPECT)
	default:
		return fmt.Errorf("unsupported expression type: %T", expr)
	}
	return nil
}
