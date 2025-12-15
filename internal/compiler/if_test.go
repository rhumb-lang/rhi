package compiler_test

import (
	"testing"

	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/compiler"
	mapval "github.com/rhumb-lang/rhi/internal/map"
	"github.com/rhumb-lang/rhi/internal/vm"
)

func TestCompiler_IfStatement(t *testing.T) {
	// if (1 == 1) => 10 else 20
	// This would be `(1 == 1) => 10` for an if-true, and then `(1 == 1) ~> 20` for an if-false block.
	// For simplicity, let's test `1 == 1 => 10`.
	doc := &ast.Document{
		Expressions: []ast.Expression{
			&ast.BinaryExpression{
				Left: &ast.BinaryExpression{
					Left:  &ast.IntegerLiteral{Value: 1},
					Op:    ast.OpEq,
					Right: &ast.IntegerLiteral{Value: 1},
				},
				Op:    ast.OpIfTrue,
				Right: &ast.IntegerLiteral{Value: 10},
			},
		},
	}

	c := compiler.NewCompiler()
	chunk, err := c.Compile(doc)
	if err != nil {
		t.Fatalf("Compilation failed: %v", err)
	}

	// Expected Bytecode for `1 == 1 => 10`:
	// The `=>` operator returns the CONDITION, not the body.
	// So stack should have `1` (True).

	// ... compilation ...

	vmMachine := vm.NewVM()
	res, err := vmMachine.Interpret(chunk)
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	if res != vm.Ok {
		t.Errorf("Expected Ok result, got %v", res)
	}

	if vmMachine.SP != 1 {
		t.Errorf("Expected 1 value on stack, got %d", vmMachine.SP)
	} else {
		val := vmMachine.Stack[0]
		if val.Type != mapval.ValBoolean || val.Integer != 1 {
			t.Errorf("Expected Boolean(True), got %v", val)
		}
	}
}

func TestCompiler_IfElseStatement(t *testing.T) {
	// 1 == 2 => 10
	// Condition is False (0). Body skipped. Result is Condition (0).
	doc := &ast.Document{
		Expressions: []ast.Expression{
			&ast.BinaryExpression{
				Left: &ast.BinaryExpression{
					Left:  &ast.IntegerLiteral{Value: 1},
					Op:    ast.OpEq,
					Right: &ast.IntegerLiteral{Value: 2},
				},
				Op:    ast.OpIfTrue,
				Right: &ast.IntegerLiteral{Value: 10},
			},
		},
	}

	c := compiler.NewCompiler()
	chunk, err := c.Compile(doc)
	if err != nil {
		t.Fatalf("Compilation failed: %v", err)
	}

	vmMachine := vm.NewVM()
	res, err := vmMachine.Interpret(chunk)
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	if res != vm.Ok {
		t.Errorf("Expected Ok result, got %v", res)
	}

	if vmMachine.SP != 1 {
		t.Errorf("Expected 1 value on stack, got %d", vmMachine.SP)
	} else {
		val := vmMachine.Stack[0]
		if val.Type != mapval.ValBoolean || val.Integer != 0 {
			t.Errorf("Expected Boolean(False), got %v", val)
		}
	}
}
