package compiler

import (
	"fmt"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	mapval "git.sr.ht/~madcapjake/rhi/internal/map"
)

func (c *Compiler) compileBinary(bin *ast.BinaryExpression) error {
	// Function Definition
	if bin.Op == ast.OpMakeFn {
		// LHS is Params (Map), RHS is Body (Routine)

		// Create child compiler
		child := NewCompiler()
		child.Enclosing = c
		child.Function.Name = "<anonymous>"

		// Parse params from LHS (bin.Left)
		if mapExpr, ok := bin.Left.(*ast.MapExpression); ok {
			child.Function.Arity = len(mapExpr.Fields)
			for _, field := range mapExpr.Fields {
				var name string
				switch f := field.(type) {
				case *ast.FieldDefinition:
					if label, ok := f.Key.(*ast.LabelLiteral); ok {
						name = label.Value
					}
				case *ast.FieldPun:
					if label, ok := f.Key.(*ast.LabelLiteral); ok {
						name = label.Value
					}
				case *ast.FieldElement:
					if label, ok := f.Value.(*ast.LabelLiteral); ok {
						name = label.Value
					} else if unary, ok := f.Value.(*ast.UnaryExpression); ok {
						// Handle #pong(msg) pattern in params
						if unary.Op == ast.OpSignal {
							if call, ok := unary.Expr.(*ast.CallExpression); ok {
								if len(call.Args) > 0 {
									if argLabel, ok := call.Args[0].(*ast.LabelLiteral); ok {
										name = argLabel.Value
									}
								}
							}
						}
					}
				}

				if name != "" {
					child.Scope.addLocal(name)
				} else {
					return fmt.Errorf("unsupported parameter type: %T", field)
				}
			}
		}

		// Hoist locals for function body
		hoister := NewHoister()
		locals := hoister.Hoist(bin.Right)
		for _, name := range locals {
			// Check if already declared in enclosing scopes (for Upvalue support)
			// But wait, parameters (added above) shadow outer variables.
			// And we just added parameters to child.Scope.
			// child.isDeclared(name) will return true if it's a parameter.
			// If it's a parameter, it's a local.
			// If it's NOT a parameter, check enclosing.
			// `child.Scope.resolveLocal(name)` handles parameters.

			if child.Scope.resolveLocal(name) != -1 {
				// Already a parameter, ignore (it's local)
				continue
			}

			// Check Enclosing
			if child.Enclosing != nil && child.Enclosing.isDeclared(name) {
				// Exists in outer scope -> Don't make it local (Upvalue)
				continue
			}

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

		// Write Upvalue Descriptors
		for _, up := range child.Upvalues {
			isLocal := byte(0)
			if up.IsLocal {
				isLocal = 1
			}
			c.Chunk().WriteByte(isLocal, 0)
			c.Chunk().WriteByte(byte(up.Index), 0)
		}

		return nil
	}

	// Special handling for Assignment
	if bin.Op == ast.OpAssignImm || bin.Op == ast.OpAssignMut {
		// Check LHS: Label (Local)
		if label, ok := bin.Left.(*ast.LabelLiteral); ok {
			// Variable Assignment: x .= 1

			// 1. Compile RHS
			if err := c.compileExpression(bin.Right); err != nil {
				return err
			}

			// 2. Resolve Variable
			idx := c.Scope.resolveLocal(label.Value)
			if idx != -1 {
				// Local
				c.emit(mapval.OP_STORE_LOC)
				c.Chunk().WriteByte(byte(idx), 0)
				return nil
			}

			up := c.resolveUpvalue(label.Value)
			if up != -1 {
				// Upvalue
				c.emit(mapval.OP_STORE_UPVALUE)
				c.Chunk().WriteByte(byte(up), 0)
				return nil
			}

			return fmt.Errorf("undeclared variable (hoisting failed): %s", label.Value)
		}

		// Check LHS: Chain (Map Field) e.g. obj\x .= 1
		if chain, ok := bin.Left.(*ast.ChainExpression); ok {
			if len(chain.Steps) == 0 {
				return fmt.Errorf("invalid assignment target")
			}
			lastStep := chain.Steps[len(chain.Steps)-1]
			if lastStep.Op != ast.ChainMember && lastStep.Op != ast.ChainSubfield {
				return fmt.Errorf("can only assign to fields")
			}

			// Compile Receiver (Base + Steps[:-1])
			if err := c.compileExpression(chain.Base); err != nil {
				return err
			}
			for i := 0; i < len(chain.Steps)-1; i++ {
				step := chain.Steps[i]
				// Logic from compileChain
				switch step.Op {
				case ast.ChainMember:
					idx := c.makeConstant(mapval.NewText(step.Ident))
					c.emit(mapval.OP_SEND)
					c.Chunk().WriteByte(byte(idx), 0)
				case ast.ChainSubfield:
					idx := c.makeConstant(mapval.NewText("@" + step.Ident))
					c.emit(mapval.OP_SEND)
					c.Chunk().WriteByte(byte(idx), 0)
				default:
					return fmt.Errorf("unsupported chain step in assignment target")
				}
			}

			// Compile Value
			if err := c.compileExpression(bin.Right); err != nil {
				return err
			}

			// Emit Set
			flags := byte(0)
			if bin.Op == ast.OpAssignMut {
				flags = 1
			}

			keyName := lastStep.Ident
			if lastStep.Op == ast.ChainSubfield {
				keyName = "@" + keyName
			}

			keyIdx := c.makeConstant(mapval.NewText(keyName))
			c.emit(mapval.OP_SET_FIELD)
			c.Chunk().WriteByte(byte(keyIdx), 0)
			c.Chunk().WriteByte(flags, 0)
			return nil
		}

		return fmt.Errorf("invalid assignment target type: %T", bin.Left)
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
		jumpOp := mapval.OP_JUMP_IF_FALSE
		if bin.Op == ast.OpIfFalse {
			jumpOp = mapval.OP_JUMP_IF_TRUE
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
	case ast.OpAdd:
		c.emit(mapval.OP_ADD)
	case ast.OpSub:
		c.emit(mapval.OP_SUB)
	case ast.OpMult:
		c.emit(mapval.OP_MULT)
	case ast.OpDivFloat:
		c.emit(mapval.OP_DIV_FLOAT)
	case ast.OpDivInt:
		c.emit(mapval.OP_DIV_INT)
	case ast.OpMod:
		c.emit(mapval.OP_MOD)
	case ast.OpPow:
		c.emit(mapval.OP_POW)
	case ast.OpRoot:
		c.emit(mapval.OP_ROOT)
	case ast.OpSciNot:
		c.emit(mapval.OP_SCI_NOT)
	case ast.OpDev:
		c.emit(mapval.OP_DEV)
	case ast.OpAnd:
		c.emit(mapval.OP_AND)
	case ast.OpOr:
		c.emit(mapval.OP_OR)
	case ast.OpAssignImm:
		c.emit(mapval.OP_ASSIGN_IMM) // For map fields
	case ast.OpAssignMut:
		c.emit(mapval.OP_ASSIGN_MUT) // For map fields
	case ast.OpEq:
		c.emit(mapval.OP_EQ)
	case ast.OpNeq:
		c.emit(mapval.OP_NEQ)
	case ast.OpGt:
		c.emit(mapval.OP_GT)
	case ast.OpLt:
		c.emit(mapval.OP_LT)
	case ast.OpGte:
		c.emit(mapval.OP_GTE)
	case ast.OpLte:
		c.emit(mapval.OP_LTE)
	case ast.OpRange:
		c.emit(mapval.OP_RANGE)
	case ast.OpHasSub:
		c.emit(mapval.OP_HAS_SUBFIELD)
	case ast.OpNotHasSub:
		c.emit(mapval.OP_NOT_HAS_SUB)
	case ast.OpHasFld:
		c.emit(mapval.OP_HAS_FIELD)
	case ast.OpNotHasFld:
		c.emit(mapval.OP_NOT_HAS_FLD)
	case ast.OpTempSub:
		c.emit(mapval.OP_TEMP_SUBFIELD)
	case ast.OpConcat:
		c.emit(mapval.OP_CONCAT)
	case ast.OpNested:
		c.emit(mapval.OP_ACCESS_NESTED)
	case ast.OpCoalesce:
		c.emit(mapval.OP_COALESCE)
	case ast.OpPipe:
		c.emit(mapval.OP_PIPE)
	case ast.OpForeach:
		c.emit(mapval.OP_FOREACH)
	case ast.OpWhile:
		loopStart := len(c.Chunk().Code)

		// Compile Condition
		if err := c.compileExpression(bin.Left); err != nil {
			return err
		}

		// Jump to End if False
		exitJump := c.emitJump(mapval.OP_JUMP_IF_FALSE)

		// Compile Body
		if err := c.compileExpression(bin.Right); err != nil {
			return err
		}

		// Discard body result
		c.emit(mapval.OP_POP)

		// Jump back to Start
		c.emitJumpBack(mapval.OP_JUMP, loopStart)

		// Patch Exit
		c.patchJump(exitJump)

		// Loop expression result
		c.emitConstant(mapval.NewEmpty())

	default:
		return fmt.Errorf("unsupported binary op: %v", bin.Op)
	}
	return nil
}
