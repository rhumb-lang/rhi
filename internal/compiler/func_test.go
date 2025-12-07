package compiler_test

import (
	"testing"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/compiler"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

func TestCompiler_Function(t *testing.T) {
	// f := [] -> (10)
	// result := f()
	
	// 1. Function Definition: [] -> (10)
	funcDef := &ast.BinaryExpression{
		Left: &ast.MapExpression{}, // []
		Op:   ast.OpMakeFn,
		Right: &ast.RoutineExpression{
			Expressions: []ast.Expression{
				&ast.IntegerLiteral{Value: 10},
			},
		},
	}
	
	// 2. f := ...
	assignF := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "f"},
		Op:    ast.OpAssignMut,
		Right: funcDef,
	}
	
	// 3. f()
	callF := &ast.CallExpression{
		Callee: &ast.LabelLiteral{Value: "f"},
		Args:   []ast.Expression{},
	}
	
	// 4. result := f()
	assignResult := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "result"},
		Op:    ast.OpAssignMut,
		Right: callF,
	}
	
	doc := &ast.Document{
		Expressions: []ast.Expression{
			assignF,
			assignResult,
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
	
	// Stack should contain [f_closure, result_value, result_expression_value]
	// locals: f (0), result (1).
	// result_value is 10.
	
	if vmMachine.SP != 3 {
		t.Errorf("Expected 3 values on stack, got %d", vmMachine.SP)
	} else {
		// Check variable 'result' at index 1
		val := vmMachine.Stack[1] 
		if val.Type != mapval.ValInteger || val.Integer != 10 {
			t.Errorf("Expected 10, got %v", val)
		}
	}
}
