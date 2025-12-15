package compiler_test

import (
	"testing"

	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/compiler"
	mapval "github.com/rhumb-lang/rhi/internal/map"
	"github.com/rhumb-lang/rhi/internal/vm"
)

func TestCompiler_Map(t *testing.T) {
	// m := [ x: 1; y: 2 ]
	// val := m\x

	mapExpr := &ast.MapExpression{
		Fields: []ast.Field{
			&ast.FieldDefinition{
				Key:       &ast.LabelLiteral{Value: "x"},
				Value:     &ast.IntegerLiteral{Value: 1},
				IsMutable: true, // :
			},
			&ast.FieldDefinition{
				Key:       &ast.LabelLiteral{Value: "y"},
				Value:     &ast.IntegerLiteral{Value: 2},
				IsMutable: false, // .
			},
		},
	}

	assignM := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "m"},
		Op:    ast.OpAssignMut,
		Right: mapExpr,
	}

	// m\x
	accessX := &ast.ChainExpression{
		Base: &ast.LabelLiteral{Value: "m"},
		Steps: []ast.ChainStep{
			{Op: ast.ChainMember, Ident: "x"},
		},
	}

	doc := &ast.Document{
		Expressions: []ast.Expression{
			assignM,
			accessX,
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

	// Stack should have `m` (Map) then result of `m\x` (1).
	// Wait, assignM pushes m? No, assignment leaves m on stack.
	// accessX pushes 1.
	// So stack: [m, 1]. SP=2.
	if machine.SP != 2 {
		t.Errorf("Expected 2 values on stack, got %d", machine.SP)
	} else {
		val := machine.Stack[1]
		if val.Type != mapval.ValInteger || val.Integer != 1 {
			t.Errorf("Expected 1, got %v", val)
		}
	}
}
