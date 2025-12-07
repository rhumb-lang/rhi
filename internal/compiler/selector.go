package compiler

import (
	"fmt"
	
	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/map"
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
	
	// Hoist locals from patterns (Action bodies)
	hoister := NewHoister()
	locals := hoister.Hoist(s)
	for _, name := range locals {
		if name == "_" { continue } // Already added
		child.Scope.addLocal(name)
		child.emitConstant(mapval.NewEmpty()) // Reserve slot for locals
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
			nextJump := child.emitJump(mapval.OP_IF_FALSE)
			
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
		// Bind variable
		idx := c.Scope.resolveLocal(t.Value)
		if idx != -1 {
			// Pinning: Subject == existing local?
			// Logic: Load Local. EQ.
			// But if we are binding to "_" (Slot 0), and Subject IS "_", then "_ == _" is True.
			// However, "_" is usually the Subject argument.
			// If target is `x`, and `x` is not defined, it's binding.
			// If `x` is defined (e.g. outer scope? or previous match?), pinning.
			
			// Since we hoist, `x` is defined in current scope (initialized to Empty).
			// But is it "bound"?
			// Hoister adds all labels. So `idx` will always be found for `x` in pattern.
			// We need to know if `x` was meant to be a NEW binding or EXISTING.
			// Pattern matching semantics: LHS of `..` usually binds if it's a bare label.
			// Pinning happens if it matches an *existing* value.
			// Since `x` is local to this selector (hoisted), it is "new".
			// So we should bind.
			
			// Store Subject to Local
			c.emit(mapval.OP_STORE_LOC)
			c.Chunk().WriteByte(byte(idx), 0)
			
			// Push True
			c.emitConstant(mapval.NewBoolean(true)) // STORE peeks, so Subject still there? 
			// My STORE_LOC peeks. So Stack: [Subject].
			// We need [True].
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
		
	default:
		return fmt.Errorf("unsupported pattern match target: %T", target)
	}
}
