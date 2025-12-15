package compiler_test

import (
	"testing"

	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/compiler"
	mapval "github.com/rhumb-lang/rhi/internal/map"
	"github.com/rhumb-lang/rhi/internal/vm"
)

func TestCompiler_MapAdvanced(t *testing.T) {
	// m := [ a :: 10 ]
	// m\b := 20
	// res := m\a ++ m\b

	mapExpr := &ast.MapExpression{
		Fields: []ast.Field{
			&ast.FieldDefinition{
				Key:       &ast.LabelLiteral{Value: "a"},
				Value:     &ast.IntegerLiteral{Value: 10},
				IsMutable: true,
			},
		},
	}

	assignM := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "m"},
		Op:    ast.OpAssignMut,
		Right: mapExpr,
	}

	assignB := &ast.BinaryExpression{
		Left: &ast.ChainExpression{
			Base:  &ast.LabelLiteral{Value: "m"},
			Steps: []ast.ChainStep{{Op: ast.ChainMember, Ident: "b"}},
		},
		Op:    ast.OpAssignMut, // :=
		Right: &ast.IntegerLiteral{Value: 20},
	}

	calcRes := &ast.BinaryExpression{
		Left: &ast.LabelLiteral{Value: "res"},
		Op:   ast.OpAssignMut,
		Right: &ast.BinaryExpression{
			Left: &ast.ChainExpression{
				Base:  &ast.LabelLiteral{Value: "m"},
				Steps: []ast.ChainStep{{Op: ast.ChainMember, Ident: "a"}},
			},
			Op: ast.OpAdd,
			Right: &ast.ChainExpression{
				Base:  &ast.LabelLiteral{Value: "m"},
				Steps: []ast.ChainStep{{Op: ast.ChainMember, Ident: "b"}},
			},
		},
	}

	doc := &ast.Document{
		Expressions: []ast.Expression{
			assignM,
			assignB,
			calcRes,
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

	// Stack: [m, res, res_val]. SP=3?
	// m := ... (Stack[0] = m).
	// m\b := ... (Consumes value).
	// res := ... (Stack[1] = 30). Result pushed (Stack[2] = 30)? No.
	// Assignment `STORE_LOC` peeks.
	// `Compile` pops non-last.
	// Last expression `res := ...`. Result kept.
	// So Stack: [m, res]. SP=2.
	// m at 0. res (30) at 1.

	if machine.SP != 2 {
		t.Errorf("Expected 2 values on stack, got %d", machine.SP)
	} else {
		// res is at index 1
		val := machine.Stack[1]
		if val.Type != mapval.ValInteger || val.Integer != 30 {
			t.Errorf("Expected res=30, got %v", val)
		}
	}
}

func TestCompiler_MapDelegation(t *testing.T) {
	// parent := [ proto :: 99 ]
	// child := [ @p :: parent ]
	// res := child\proto

	parentExpr := &ast.MapExpression{
		Fields: []ast.Field{
			&ast.FieldDefinition{
				Key:       &ast.LabelLiteral{Value: "proto"},
				Value:     &ast.IntegerLiteral{Value: 99},
				IsMutable: true,
			},
		},
	}

	assignParent := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "parent"},
		Op:    ast.OpAssignMut,
		Right: parentExpr,
	}

	childExpr := &ast.MapExpression{
		Fields: []ast.Field{
			&ast.FieldDefinition{
				Key:       &ast.LabelLiteral{Value: "p"},
				Value:     &ast.LabelLiteral{Value: "parent"},
				IsMutable: true,
				IsSub:     true, // @p
			},
		},
	}

	assignChild := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "child"},
		Op:    ast.OpAssignMut,
		Right: childExpr,
	}

	assignRes := &ast.BinaryExpression{
		Left: &ast.LabelLiteral{Value: "res"},
		Op:   ast.OpAssignMut,
		Right: &ast.ChainExpression{
			Base:  &ast.LabelLiteral{Value: "child"},
			Steps: []ast.ChainStep{{Op: ast.ChainMember, Ident: "proto"}},
		},
	}

	doc := &ast.Document{
		Expressions: []ast.Expression{
			assignParent,
			assignChild,
			assignRes,
		},
	}

	c := compiler.NewCompiler()
	chunk, err := c.Compile(doc)
	if err != nil {
		t.Fatalf("Compilation failed: %v", err)
	}

	machine := vm.NewVM()
	machine.Interpret(chunk)

	// Stack: [parent, child, res_val, res_val] (SP=4)
	// Locals: parent(0), child(1), res(2).
	// Stack[2] updated to 99.
	// Stack[3] is expression result (99).

	if machine.SP != 4 {
		t.Errorf("Expected 4 values on stack, got %d", machine.SP)
	} else {
		val := machine.Stack[3]
		if val.Type != mapval.ValInteger || val.Integer != 99 {
			t.Errorf("Expected res=99, got %v", val)
		}
	}
}
