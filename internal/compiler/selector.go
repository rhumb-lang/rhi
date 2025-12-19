package compiler

import (
	"fmt"

	"github.com/rhumb-lang/rhi/internal/ast"
	mapval "github.com/rhumb-lang/rhi/internal/map"
)

func (c *Compiler) compileSelector(s *ast.SelectorExpression) error {
	// Selector is a Function that takes 1 arg (Subject)
	child := NewCompiler()
	child.Enclosing = c
	child.Function.Name = "<selector>"
	child.Function.Arity = 1

	// Arg 0 is Subject. Name it "_" so patterns can reference it.
	// Since it's an argument, it's already on the stack (Slot 0).
	// We do NOT emitConstant for it.
	child.Scope.addLocal("_")

	// Hoist locals
	hoister := NewHoister()
	for _, pat := range s.Patterns {
		if p, ok := pat.(*ast.PatternDefinition); ok {
			// Binding in Target
			if label, ok := p.Target.(*ast.LabelLiteral); ok {
				hoister.add(label.Value)
			}
			// Binding in Signal Pattern #topic(arg)
			if unary, ok := p.Target.(*ast.UnaryExpression); ok {
				if unary.Op == ast.OpSignal {
					if call, ok := unary.Expr.(*ast.CallExpression); ok {
						for _, arg := range call.Args {
							if label, ok := arg.(*ast.LabelLiteral); ok {
								hoister.add(label.Value)
							}
						}
					}
				}
			}
			// Locals in Action
			_ = hoister.Hoist(p.Action)
		} else if def, ok := pat.(*ast.PatternDefault); ok {
			hoister.Hoist(def.Value)
		}
	}

	locals := hoister.Locals
	for _, name := range locals {
		if name == "_" {
			continue
		} // Already added

		if child.Scope.resolveLocal(name) != -1 {
			continue
		}

		// Allow shadowing of enclosing variables in selectors
		// If we don't shadow, we can't bind new values to names that exist outside.
		// Pinning (matching against outer var) requires distinct handling not yet implemented.
		// if child.Enclosing != nil && child.Enclosing.isDeclared(name) {
		// 	continue
		// }

		child.Scope.addLocal(name)
		child.emitConstant(mapval.NewEmpty()) // Reserve slot
	}

	for _, pat := range s.Patterns {
		// Load Subject for matching
		// Note: We load from local "_" (Slot 0) to put on stack for matchers that expect it (like Literals)
		// Matchers that use "_" explicit variable will resolve LOAD_LOC 0 themselves.
		child.emit(mapval.OP_LOAD_LOC)
		child.Chunk().WriteByte(0, 0) // Load Subject

		switch p := pat.(type) {
		case *ast.PatternDefinition:
			if err := child.compilePattern(p.Target); err != nil {
				return err
			}

			// Jump if False (No Match) -> Next Pattern
			// OP_IF_TRUE jumps if Falsy (Branch If False) -> Renamed to OP_JUMP_IF_FALSE
			nextJump := child.emitJump(mapval.OP_JUMP_IF_FALSE)

			// Match! Execute Action.
			if err := child.compileExpression(p.Action); err != nil {
				return err
			}

			if p.IsConsume {
				child.emit(mapval.OP_RETURN)
			} else {
				child.emit(mapval.OP_POP)
			}

			child.patchJump(nextJump)

		case *ast.PatternDefault:
			child.emit(mapval.OP_POP) // Pop Subject
			if err := child.compileExpression(p.Value); err != nil {
				return err
			}
			child.emit(mapval.OP_RETURN)

		default:
			return fmt.Errorf("unsupported pattern type")
		}
	}

	// If fallthrough all patterns, return Empty
	child.emitConstant(mapval.NewEmpty())
	child.emit(mapval.OP_RETURN)

	// Create Function Constant
	fnVal := mapval.NewFunction(child.Function)
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

func (c *Compiler) compilePattern(target ast.Expression) error {
	// Stack: [Subject]

	switch t := target.(type) {
	case *ast.BinaryExpression:
		// Condition Match (e.g. _ > 0)
		// The expression uses variables (like _) to check condition.
		// It does NOT consume the Subject on stack implicitly.
		// So we must POP the Subject provided by the loop.
		c.emit(mapval.OP_POP)
		return c.compileExpression(t)

	case *ast.LabelLiteral:
		// Special Wildcard "_"
		if t.Value == "_" {
			c.emit(mapval.OP_POP)
			c.emitConstant(mapval.NewBoolean(true))
			return nil
		}

		// 1. Check for Boolean Literals (yes/no)
		// TODO: Add international support
		if t.Value == "yes" {
			c.emitConstant(mapval.NewBoolean(true))
			c.emit(mapval.OP_EQ)
			return nil
		}
		if t.Value == "no" {
			c.emitConstant(mapval.NewBoolean(false))
			c.emit(mapval.OP_EQ)
			return nil
		}

		// 2. Bind variable
		idx := c.Scope.resolveLocal(t.Value)
		if idx != -1 {
			// Store Subject to Local
			c.emit(mapval.OP_STORE_LOC)
			c.Chunk().WriteByte(byte(idx), 0)

			// Push True (Match Successful)
			// STORE_LOC leaves the value on stack (peek), but we need 'True' for the match result.
			// Stack: [Subject] -> OP_POP -> [] -> Push True
			c.emit(mapval.OP_POP)
			c.emitConstant(mapval.NewBoolean(true))
			return nil
		}
		// Should not happen if hoisted
		return fmt.Errorf("undeclared binding variable: %s", t.Value)

	case *ast.IntegerLiteral, *ast.RationalLiteral, *ast.TextLiteral, *ast.EmptyLiteral:
		// Literal Match
		c.compileExpression(t) // Push Literal
		c.emit(mapval.OP_EQ)   // Subject == Literal
		return nil

	case *ast.UnaryExpression:
		if t.Op == ast.OpSignal {
			if label, ok := t.Expr.(*ast.LabelLiteral); ok {
				// Match Tuple: Kind Signal, Topic label.Value

				topicIdx := c.makeConstant(mapval.NewText(label.Value))

				c.emit(mapval.OP_MATCH_TUPLE)
				c.Chunk().WriteByte(byte(mapval.TupleSignal), 0)
				// Topic Index (2 bytes)
				c.Chunk().WriteByte(byte(topicIdx>>8), 0)   // High byte
				c.Chunk().WriteByte(byte(topicIdx&0xFF), 0) // Low byte

				return nil
			}

			if call, ok := t.Expr.(*ast.CallExpression); ok {
				if label, ok := call.Callee.(*ast.LabelLiteral); ok {
					topicIdx := c.makeConstant(mapval.NewText(label.Value))
					return c.compileSignalPattern(call, topicIdx)
				}
			}
		}
		return fmt.Errorf("unsupported unary pattern: %v", t.Op)

	default:
		return fmt.Errorf("unsupported pattern match target: %T", target)
	}
}

func (c *Compiler) compileSignalPattern(call *ast.CallExpression, topicIdx int) error {
	// 1. Check Tuple
	c.emit(mapval.OP_MATCH_TUPLE)
	c.Chunk().WriteByte(byte(mapval.TupleSignal), 0)
	c.Chunk().WriteByte(byte(topicIdx>>8), 0)
	c.Chunk().WriteByte(byte(topicIdx&0xFF), 0)

	// Stack: [Subject, Bool]
	failJump := c.emitJump(mapval.OP_JUMP_IF_FALSE) // Consumes Bool. Stack: [Subject]

	// If True:
	c.emit(mapval.OP_POP) // Pop Subject. Stack Clean []

	var argFails []int

	// Check Args
	for i, arg := range call.Args {
		// Load Subject from Local 0
		c.emit(mapval.OP_LOAD_LOC)
		c.Chunk().WriteByte(0, 0)

		// Get Payload[i]
		idxStr := fmt.Sprintf("%d", i+1)
		idxConst := c.makeConstant(mapval.NewText(idxStr))
		c.emit(mapval.OP_SEND)
		c.Chunk().WriteByte(byte(idxConst), 0)

		// Match Arg
		if err := c.compilePattern(arg); err != nil {
			return err
		}
		// Stack: [ArgBool]

		// If ArgBool is False, we fail.
		argFail := c.emitJump(mapval.OP_JUMP_IF_FALSE)
		argFails = append(argFails, argFail)
	}

	// All success
	c.emitConstant(mapval.NewBoolean(true))
	endJump := c.emitJump(mapval.OP_JUMP)

	// Fail Target (from Tuple Match)
	c.patchJump(failJump)
	c.emit(mapval.OP_POP) // Pop Subject. Stack Clean []

	// Fail Target (from Args)
	for _, fail := range argFails {
		c.patchJump(fail)
	}

	c.emitConstant(mapval.NewBoolean(false))

	// End
	c.patchJump(endJump)
	return nil
}
