package compiler_test

import (
	"testing"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/compiler"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

func TestCompiler_MapMixed(t *testing.T) {
	// foo := [ 10; a .. 10; 30 ]
	// AST construction
	mapExpr := &ast.MapExpression{
		Fields: []ast.Field{
			// 10 (Positional 1)
			&ast.FieldElement{
				Value: &ast.IntegerLiteral{Value: 10},
			},
			// a .. 10 (Named)
			&ast.FieldDefinition{
				Key:       &ast.LabelLiteral{Value: "a"},
				Value:     &ast.IntegerLiteral{Value: 10},
				IsMutable: false, // ..
			},
			// 30 (Positional 2)
			&ast.FieldElement{
				Value: &ast.IntegerLiteral{Value: 30},
			},
		},
	}
	
	assignFoo := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "foo"},
		Op:    ast.OpAssignMut, // :=
		Right: mapExpr,
	}
	
doc := &ast.Document{
		Expressions: []ast.Expression{
			assignFoo,
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
	
	// Stack: [foo, foo_val] (SP=2)
	if machine.SP != 2 {
		t.Fatalf("Expected stack size 2, got %d", machine.SP)
	}
	
	val := machine.Stack[1]
	if val.Type != mapval.ValObject {
		t.Fatalf("Expected Map, got %v", val)
	}
	m := val.Obj.(*mapval.Map)
	
	// Verify "1" -> 10
	v1, found1 := m.Get("1")
	if !found1 || v1.Integer != 10 {
		t.Errorf("Expected foo[1] == 10, got %v", v1)
	}
	
	// Verify "2" -> 30
	v2, found2 := m.Get("2")
	if !found2 || v2.Integer != 30 {
		t.Errorf("Expected foo[2] == 30, got %v", v2)
	}
	
	// Verify "a" -> 10
	va, foundA := m.Get("a")
	if !foundA || va.Integer != 10 {
		t.Errorf("Expected foo[a] == 10, got %v", va)
	}
}