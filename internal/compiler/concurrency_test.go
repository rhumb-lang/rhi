package compiler_test

import (
	"testing"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/compiler"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
)

func TestCompiler_Concurrency(t *testing.T) {
	// obj := []
	// obj#click
	// obj^ack
	// obj$ready
	
	objDecl := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "obj"},
		Op:    ast.OpAssignMut,
		Right: &ast.MapExpression{},
	}
	
	sig := &ast.ChainExpression{
		Base: &ast.LabelLiteral{Value: "obj"},
		Steps: []ast.ChainStep{{Op: ast.ChainSignal, Ident: "click"}},
	}
	
	reply := &ast.ChainExpression{
		Base: &ast.LabelLiteral{Value: "obj"},
		Steps: []ast.ChainStep{{Op: ast.ChainReply, Ident: "ack"}},
	}
	
	proc := &ast.ChainExpression{
		Base: &ast.LabelLiteral{Value: "obj"},
		Steps: []ast.ChainStep{{Op: ast.ChainProclamation, Ident: "ready"}},
	}
	
	doc := &ast.Document{
		Expressions: []ast.Expression{
			objDecl,
			sig,
			reply,
			proc,
		},
	}

	c := compiler.NewCompiler()
	chunk, err := c.Compile(doc)
	if err != nil {
		t.Fatalf("Compilation failed: %v", err)
	}
	
	machine := vm.NewVM()
	// Should print debug messages from stubbed ops
	res, err := machine.Interpret(chunk)
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	if res != vm.Ok {
		t.Errorf("Expected Ok result, got %v", res)
	}
	
	// Stack state:
	// objDecl -> leaves obj on stack. POPPED by Compile loop (non-last).
	// sig -> leaves ?? Stubs in space_ops.go currently pop receiver and return nil (don't push).
	// wait, if they don't push, stack shrinks?
	// compileChain logic:
	// 1. Compile Base (pushes obj).
	// 2. Emit OP_POST.
	// OP_POST stub pops.
	// Result?
	// If OP_POST doesn't push result, expression returns nothing?
	// Stack imbalance?
	// The Compiler `compileExpression` assumes expression leaves a value?
	// If `compileChain` finishes, does it ensure value on stack?
	// `OP_SEND` pushes result.
	// `OP_POST` "Emit Event". Void?
	// In Rhumb, `obj#click` might return the event object or void.
	// If void, `Compile` loop might `POP` empty stack?
	// `POP` checks underflow.
	// If `sig` leaves nothing, `Compile`'s `POP` (for non-last) will panic.
	
	// Check `internal/vm/space_ops.go` implementation.
	// `func (vm *VM) opPost() error { ... vm.pop() ... return nil }`
	// It consumes receiver and pushes NOTHING.
	
	// This is a problem for `Compiler.Compile` which does `c.emit(mapval.OP_POP)` if i < len-1.
	// If `sig` pushes nothing, `OP_POP` will underflow stack (if stack was empty) or pop previous local?
	// Wait, `obj` was popped by `OP_POST`. So stack is effectively empty (relative to this expression).
	// Then `OP_POP` runs. Stack underflow.
	
	// I should update stubs to push `Empty` or `True` to signal success, maintaining expression semantics.
	// Rhumb expressions always return a value.
}
