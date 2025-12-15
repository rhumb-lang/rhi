package compiler_test

import (
	"testing"

	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/compiler"
	mapval "github.com/rhumb-lang/rhi/internal/map"
	"github.com/rhumb-lang/rhi/internal/vm"
)

func TestCompiler_Adder(t *testing.T) {
	// add := [a; b] -> a ++ b
	// res := add(10; 20)

	// Params [a; b]
	params := &ast.MapExpression{
		Fields: []ast.Field{
			&ast.FieldElement{Value: &ast.LabelLiteral{Value: "a"}},
			&ast.FieldElement{Value: &ast.LabelLiteral{Value: "b"}},
		},
	}

	// Body: a ++ b
	body := &ast.RoutineExpression{
		Expressions: []ast.Expression{
			&ast.BinaryExpression{
				Left:  &ast.LabelLiteral{Value: "a"},
				Op:    ast.OpAdd,
				Right: &ast.LabelLiteral{Value: "b"},
			},
		},
	}

	defAdd := &ast.BinaryExpression{
		Left:  params,
		Op:    ast.OpMakeFn,
		Right: body,
	}

	assignAdd := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "add"},
		Op:    ast.OpAssignMut,
		Right: defAdd,
	}

	// call
	callAdd := &ast.CallExpression{
		Callee: &ast.LabelLiteral{Value: "add"},
		Args: []ast.Expression{
			&ast.IntegerLiteral{Value: 10},
			&ast.IntegerLiteral{Value: 20},
		},
	}

	assignRes := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "res"},
		Op:    ast.OpAssignMut,
		Right: callAdd,
	}

	doc := &ast.Document{
		Expressions: []ast.Expression{
			assignAdd,
			assignRes,
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
		t.Errorf("Expected Ok result, got %v", res)
	}

	// Stack: [add, res, res_val]
	if machine.SP != 3 {
		t.Errorf("Expected 3 values on stack, got %d", machine.SP)
	} else {
		val := machine.Stack[2] // res value (pushed by assignment result)
		if val.Type != mapval.ValInteger || val.Integer != 30 {
			t.Errorf("Expected 30, got %v", val)
		}
	}
}
