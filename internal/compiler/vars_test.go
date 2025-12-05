package compiler_test

import (
	"testing"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/compiler"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
)

func TestCompiler_Variables(t *testing.T) {
	// x := 1
	// y := x + 2
	doc := &ast.Document{
		Expressions: []ast.Expression{
			&ast.BinaryExpression{
				Left:  &ast.LabelLiteral{Value: "x"},
				Op:    ast.OpAssignMut,
				Right: &ast.IntegerLiteral{Value: 1},
			},
			&ast.BinaryExpression{
				Left:  &ast.LabelLiteral{Value: "y"},
				Op:    ast.OpAssignMut,
				Right: &ast.BinaryExpression{
					Left:  &ast.LabelLiteral{Value: "x"},
					Op:    ast.OpAdd,
					Right: &ast.IntegerLiteral{Value: 2},
				},
			},
		},
	}

	c := compiler.NewCompiler()
	chunk, err := c.Compile(doc)
	if err != nil {
		t.Fatalf("Compilation failed: %v", err)
	}
	
	machine := vm.NewVM()
	res, err := machine.Interpret(chunk)
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	if res != vm.Ok {
		t.Errorf("VM result not OK")
	}
	
	// Verify stack state
	// Stack should have: [x=1, y=3]
	// Wait, assignment expression leaves value on stack.
	// 1. x := 1 -> Stack: [1] (at index 0)
	// 2. y := x + 2 -> Pushes x(1), 2, Add(3). Stack: [1, 3]. Store y at index 1. Stack: [1, 3].
	// So final stack has 2 items.
	// x is at index 0. y is at index 1.
	
	if machine.SP != 2 {
		t.Errorf("Expected stack size 2, got %d", machine.SP)
	}
	
	xVal := machine.Stack[0]
	if xVal.Integer != 1 {
		t.Errorf("Expected x=1, got %v", xVal)
	}
	
	yVal := machine.Stack[1]
	if yVal.Integer != 3 {
		t.Errorf("Expected y=3, got %v", yVal)
	}
}
