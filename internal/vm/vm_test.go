package vm_test

import (
	"testing"

	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/compiler"
	mapval "github.com/rhumb-lang/rhi/internal/map"
	"github.com/rhumb-lang/rhi/internal/vm"
)

func TestVM_Math(t *testing.T) {
	// 1 + 2 * 3 = 7
	doc := &ast.Document{
		Expressions: []ast.Expression{
			&ast.BinaryExpression{
				Left: &ast.IntegerLiteral{Value: 1},
				Op:   ast.OpAdd,
				Right: &ast.BinaryExpression{
					Left:  &ast.IntegerLiteral{Value: 2},
					Op:    ast.OpMult,
					Right: &ast.IntegerLiteral{Value: 3},
				},
			},
		},
	}

	// Compile
	c := compiler.NewCompiler()
	chunk, err := c.Compile(doc)
	if err != nil {
		t.Fatalf("Compilation failed: %v", err)
	}

	// Execute
	machine := vm.NewVM()
	res, err := machine.Interpret(chunk)
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	if res != vm.Ok {
		t.Errorf("Expected Ok result, got %v", res)
	}

	// Check result (should be on stack? No, expressions leave result on stack?
	// Compiler emits code. Does it POP results?
	// Currently, `compileExpression` leaves value on stack.
	// `compileExpression` called in loop.
	// So stack should contain the result.
	// Wait, if multiple expressions, stack accumulates?
	// Yes, for now.

	// Inspect VM stack (exported? No. Add helper or export SP/Stack).
	// I made VM.Stack public in vm.go.

	if machine.SP != 1 {
		t.Errorf("Expected 1 value on stack, got %d", machine.SP)
	} else {
		val := machine.Stack[0]
		if val.Type != mapval.ValInteger || val.Integer != 7 {
			t.Errorf("Expected 7, got %v", val)
		}
	}
}

func TestVM_Floats(t *testing.T) {
	// 1.5 + 2.5 = 4.0
	doc := &ast.Document{
		Expressions: []ast.Expression{
			&ast.BinaryExpression{
				Left:  &ast.RationalLiteral{Value: 1.5},
				Op:    ast.OpAdd,
				Right: &ast.RationalLiteral{Value: 2.5},
			},
		},
	}

	c := compiler.NewCompiler()
	chunk, _ := c.Compile(doc)
	machine := vm.NewVM()
	machine.Interpret(chunk)

	val := machine.Stack[0]
	if val.Type != mapval.ValFloat || val.Float != 4.0 {
		t.Errorf("Expected 4.0, got %v", val)
	}
}
