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

	// Arg 0 is Subject. Name it "*" so patterns can reference it.
	// Since it's an argument, it's already on the stack (Slot 0).
	// We do NOT emitConstant for it.
	child.Scope.addLocal("*")

	// Hoist locals
	hoister := NewHoister()
	for _, pat := range s.Patterns {
		if p, ok := pat.(*ast.PatternDefinition); ok {
			// Check if this is a value pattern (anything other than a signal match)
			isSignalPattern := false
			if unary, ok := p.Target.(*ast.UnaryExpression); ok && unary.Op == ast.OpSignal {
				isSignalPattern = true
			}

			if !isSignalPattern {
				child.Function.HasValuePatterns = true
			}

			// Binding in Target
			if label, ok := p.Target.(*ast.LabelLiteral); ok {
				// Only hoist if NOT declared in parent (Binding)
				// If declared, it is a Pinned Match (Comparison)
				if !c.isDeclared(label.Value) && label.Value != "*" {
					hoister.add(label.Value)
				}
			}
			// Recursive Binding search in Expressions (e.g. x -/ 15 == 0)
			binds := collectBindingLabels(p.Target)
			for _, name := range binds {
				if !c.isDeclared(name) && name != "*" {
					hoister.add(name)
				}
			}

			// Binding in Signal Pattern #topic(arg)
			if unary, ok := p.Target.(*ast.UnaryExpression); ok {
				if unary.Op == ast.OpSignal {
					if call, ok := unary.Expr.(*ast.CallExpression); ok {
						for _, arg := range call.Args {
							if label, ok := arg.(*ast.LabelLiteral); ok {
								if !c.isDeclared(label.Value) && label.Value != "*" {
									hoister.add(label.Value)
								}
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
		if name == "*" {
			continue
		} // Already added

		if child.Scope.resolveLocal(name) != -1 {
			continue
		}

		child.Scope.addLocal(name)
		child.emitConstant(mapval.NewEmpty()) // Reserve slot
	}

	for _, pat := range s.Patterns {
		// Load Subject for matching
		child.emit(mapval.OP_LOAD_LOC)
		child.Chunk().WriteByte(0, 0) // Load Subject

		switch p := pat.(type) {
		case *ast.PatternDefinition:
			if err := child.compilePattern(p.Target, false); err != nil {
				return err
			}
			// Jump if False (No Match) -> Next Pattern
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

func (c *Compiler) compilePattern(target ast.Expression, allowEmpty bool) error {
	// Stack: [Subject]

	switch t := target.(type) {
	case *ast.BinaryExpression:
		// Condition Match (e.g. x -/ 15 == 0)
		// We must bind Subject to any new variables found in expression.
		binds := collectBindingLabels(t)
		for _, name := range binds {
			idx := c.Scope.resolveLocal(name)
			if idx != -1 {
				c.emit(mapval.OP_STORE_LOC) // Peeks subject, stores
				c.Chunk().WriteByte(byte(idx), 0)
			}
		}

		// The expression uses variables (like x or *) to check condition.
		// It does NOT consume the Subject on stack implicitly.
		// So we must POP the Subject provided by the loop.
		c.emit(mapval.OP_POP)
		return c.compileExpression(t)

	case *ast.LabelLiteral:
		// Special Wildcard "*"
		if t.Value == "*" {
			c.emit(mapval.OP_POP)
			c.emitConstant(mapval.NewBoolean(true))
			return nil
		}

		// 1. Check for Boolean Literals (yes/no)
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

		// 2. Bind variable (Local)
		idx := c.Scope.resolveLocal(t.Value)
		if idx != -1 {
			if !allowEmpty {
				// Check if Subject is Empty (Bind fails on Empty)
				c.emit(mapval.OP_DUP)           // [Subject, Subject]
				c.emit(mapval.OP_IS_EMPTY)      // [Subject, IsEmpty]
				failJump := c.emitJump(mapval.OP_JUMP_IF_TRUE) // Jumps if Empty

				// Not Empty: Bind
				c.emit(mapval.OP_STORE_LOC)     // [Subject] (stored)
				c.Chunk().WriteByte(byte(idx), 0)
				c.emit(mapval.OP_POP)           // []
				c.emitConstant(mapval.NewBoolean(true)) // [True]
				endJump := c.emitJump(mapval.OP_JUMP)

				// Fail Label
				c.patchJump(failJump)
				c.emit(mapval.OP_POP)           // [Subject] -> []
				c.emitConstant(mapval.NewBoolean(false)) // [False]

				c.patchJump(endJump)
				return nil
			} else {
				// Allow Empty: Bind directly
				c.emit(mapval.OP_STORE_LOC)
				c.Chunk().WriteByte(byte(idx), 0)
				c.emit(mapval.OP_POP)
				c.emitConstant(mapval.NewBoolean(true))
				return nil
			}
		}

		// 3. Pinned Match (Upvalue)
		up := c.resolveUpvalue(t.Value)
		if up != -1 {
			// Load Upvalue
			c.emit(mapval.OP_LOAD_UPVALUE)
			c.Chunk().WriteByte(byte(up), 0) // [Subject, Upvalue]

			// Check if Upvalue is Empty (Pinned fails on Empty)
			c.emit(mapval.OP_DUP)            // [Subject, Upvalue, Upvalue]
			c.emit(mapval.OP_IS_EMPTY)       // [Subject, Upvalue, IsEmpty]
			failJump := c.emitJump(mapval.OP_JUMP_IF_TRUE)

			// Not Empty: Check Equality
			c.emit(mapval.OP_EQ)             // [Bool]
			endJump := c.emitJump(mapval.OP_JUMP)

			// Fail Label
			c.patchJump(failJump)
			c.emit(mapval.OP_POP)            // [Subject, Upvalue] -> [Subject]
			c.emit(mapval.OP_POP)            // [Subject] -> []
			c.emitConstant(mapval.NewBoolean(false))

			c.patchJump(endJump)
			return nil
		}

		// Should not happen if hoisted correctly
		return fmt.Errorf("undeclared binding variable: %s", t.Value)

	case *ast.IntegerLiteral, *ast.RationalLiteral, *ast.TextLiteral, *ast.EmptyLiteral:
		// Literal Match
		c.compileExpression(t) // Push Literal
		c.emit(mapval.OP_EQ)   // Subject == Literal
		return nil

	case *ast.MapExpression:
		// Structural Map Match
		// 1. Check if Subject is a Map
		c.emit(mapval.OP_LOAD_LOC)
		c.Chunk().WriteByte(0, 0) // Load Subject
		c.emit(mapval.OP_IS_MAP)
		
		var failJumps []int
		failJumps = append(failJumps, c.emitJump(mapval.OP_JUMP_IF_FALSE))

		for i, field := range t.Fields {
			// Subject is at the bottom of the current pattern stack [Subject]
			// We need to duplicate it to extract field
			c.emit(mapval.OP_LOAD_LOC)
			c.Chunk().WriteByte(0, 0) // Load Subject

			var key string
			var valPattern ast.Expression

			switch f := field.(type) {
			case *ast.FieldElement:
				key = fmt.Sprintf("%d", i+1)
				valPattern = f.Value
			case *ast.FieldDefinition:
				if label, ok := f.Key.(*ast.LabelLiteral); ok {
					key = label.Value
				}
				valPattern = f.Value
			case *ast.FieldPun:
				if label, ok := f.Key.(*ast.LabelLiteral); ok {
					key = label.Value
				}
				valPattern = f.Key.(ast.Expression)
			}

			if key == "" || valPattern == nil {
				return fmt.Errorf("unsupported map pattern field")
			}

			// Extract field from subject
			keyIdx := c.makeConstant(mapval.NewText(key))
			c.emit(mapval.OP_SEND)
			c.Chunk().WriteByte(byte(keyIdx), 0)

			// Stack: [Subject, FieldValue]
			// Match field value
			if err := c.compilePattern(valPattern, false); err != nil { // Map fields don't allow empty by default
				return err
			}
			// Stack: [Subject, FieldBool]

			// If FieldBool is false, whole match fails
			failJump := c.emitJump(mapval.OP_JUMP_IF_FALSE)
			failJumps = append(failJumps, failJump)
		}

		// All fields matched!
		c.emit(mapval.OP_POP) // Pop Subject
		c.emitConstant(mapval.NewBoolean(true))
		endJump := c.emitJump(mapval.OP_JUMP)

		// Fail Target
		for _, fj := range failJumps {
			c.patchJump(fj)
		}
		c.emit(mapval.OP_POP) // Pop Subject
		c.emitConstant(mapval.NewBoolean(false))

		c.patchJump(endJump)
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

// collectBindingLabels recursively finds labels in an expression that should be bound
func collectBindingLabels(expr ast.Expression) []string {
	var labels []string
	switch e := expr.(type) {
	case *ast.LabelLiteral:
		if e.Value != "*" && e.Value != "yes" && e.Value != "no" {
			labels = append(labels, e.Value)
		}
	case *ast.BinaryExpression:
		labels = append(labels, collectBindingLabels(e.Left)...)
		labels = append(labels, collectBindingLabels(e.Right)...)
	case *ast.UnaryExpression:
		labels = append(labels, collectBindingLabels(e.Expr)...)
	case *ast.MapExpression:
		for _, f := range e.Fields {
			switch field := f.(type) {
			case *ast.FieldElement:
				labels = append(labels, collectBindingLabels(field.Value)...)
			case *ast.FieldDefinition:
				labels = append(labels, collectBindingLabels(field.Value)...)
			case *ast.FieldPun:
				// FieldPun.Key is ast.Node, but in patterns it should be a label/expression
				if expr, ok := field.Key.(ast.Expression); ok {
					labels = append(labels, collectBindingLabels(expr)...)
				}
			}
		}
	case *ast.CallExpression:
		for _, arg := range e.Args {
			labels = append(labels, collectBindingLabels(arg)...)
		}
	}
	// Unique
	seen := make(map[string]bool)
	var uniq []string
	for _, l := range labels {
		if !seen[l] {
			uniq = append(uniq, l)
			seen[l] = true
		}
	}
	return uniq
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

		// Match Arg - ALLOW EMPTY BINDINGS IN SIGNALS
		if err := c.compilePattern(arg, true); err != nil {
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
