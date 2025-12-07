package compiler

import (
	"fmt"
	
	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

func (c *Compiler) compileBinary(bin *ast.BinaryExpression) error {
	// Function Definition
	if bin.Op == ast.OpMakeFn {
		// LHS is Params (Map), RHS is Body (Routine)
		
		// Create child compiler
		child := NewCompiler()
		child.Enclosing = c
		child.Function.Name = "<anonymous>"
		
		// TODO: Parse params from LHS (bin.Left)
		if mapExpr, ok := bin.Left.(*ast.MapExpression); ok {
			child.Function.Arity = len(mapExpr.Fields)
			// We need to add these fields as locals in child scope
			// iterating fields... (skipping for MVP)
		}
		
		// Hoist locals for function body
		hoister := NewHoister()
		locals := hoister.Hoist(bin.Right)
		for _, name := range locals {
			child.Scope.addLocal(name)
			child.emitConstant(mapval.NewEmpty())
		}
		
		// Compile Body (RHS)
		if err := child.compileExpression(bin.Right); err != nil {
			return err
		}
		
		child.emit(mapval.OP_RETURN)
		
		// Create Function Constant
		fnVal := mapval.NewFunction(child.Function)
		
		// Add to current chunk constants (without emitting LOAD_CONST)
		idx := c.makeConstant(fnVal)
		c.emit(mapval.OP_MAKE_FN)
		c.Chunk().WriteByte(byte(idx), 0)
		
		return nil
	}

	// Special handling for Assignment
	if bin.Op == ast.OpAssignImm || bin.Op == ast.OpAssignMut {
		// Check LHS
		if label, ok := bin.Left.(*ast.LabelLiteral); ok {
			// Variable Assignment: x .= 1
			
			// 1. Compile RHS
			if err := c.compileExpression(bin.Right); err != nil {
				return err
			}
			
			// 2. Resolve Variable (Must exist due to hoisting)
			idx := c.Scope.resolveLocal(label.Value)
			if idx == -1 {
				return fmt.Errorf("undeclared variable (hoisting failed): %s", label.Value)
			}
			
			// Update existing
			c.emit(mapval.OP_STORE_LOC)
			c.Chunk().WriteByte(byte(idx), 0)
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
		c.emit(mapval.OP_DUP)
		
		// Determine Jump Op
		jumpOp := mapval.OP_IF_TRUE
		if bin.Op == ast.OpIfFalse {
			jumpOp = mapval.OP_IF_FALSE
		}
		
		skipBody := c.emitJump(jumpOp)
		
		// Compile Body (RHS)
		if err := c.compileExpression(bin.Right); err != nil {
			return err
		}
		
		// Discard body result, return condition (which is under body result)
		c.emit(mapval.OP_POP)
		
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
	case ast.OpAdd: c.emit(mapval.OP_ADD)
	case ast.OpSub: c.emit(mapval.OP_SUB)
	case ast.OpMult: c.emit(mapval.OP_MULT)
	case ast.OpDivFloat: c.emit(mapval.OP_DIV_FLOAT)
	case ast.OpAssignImm: c.emit(mapval.OP_ASSIGN_IMM) // For map fields
	case ast.OpAssignMut: c.emit(mapval.OP_ASSIGN_MUT) // For map fields
	case ast.OpEq: c.emit(mapval.OP_EQ)
	case ast.OpNeq: c.emit(mapval.OP_NEQ)
	case ast.OpGt: c.emit(mapval.OP_GT)
	case ast.OpLt: c.emit(mapval.OP_LT)
	case ast.OpGte: c.emit(mapval.OP_GTE)
	case ast.OpLte: c.emit(mapval.OP_LTE)
	default:
		return fmt.Errorf("unsupported binary op: %v", bin.Op)
	}
	return nil
}
