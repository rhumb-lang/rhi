package compiler_test

import (
	"testing"

	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/compiler"
	mapval "github.com/rhumb-lang/rhi/internal/map"
	"github.com/rhumb-lang/rhi/internal/vm"
)

func TestCompiler_Selector(t *testing.T) {
	// res1 := 10 {
	//   10 .. 100
	//   x  .. x + 1
	// }

	patterns := []ast.Pattern{
		&ast.PatternDefinition{
			Target:    &ast.IntegerLiteral{Value: 10},
			Action:    &ast.IntegerLiteral{Value: 100},
			IsConsume: true,
		},
		&ast.PatternDefinition{
			Target: &ast.LabelLiteral{Value: "x"},
			Action: &ast.BinaryExpression{
				Left:  &ast.LabelLiteral{Value: "x"},
				Op:    ast.OpAdd,
				Right: &ast.IntegerLiteral{Value: 1},
			},
			IsConsume: true,
		},
	}

	selExpr := &ast.SelectorExpression{Patterns: patterns}

	// 10 { ... } -> Call(sel, [10])
	call1 := &ast.CallExpression{
		Callee: selExpr,
		Args:   []ast.Expression{&ast.IntegerLiteral{Value: 10}},
	}
	assignRes1 := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "res1"},
		Op:    ast.OpAssignMut,
		Right: call1,
	}

	// 20 { ... } -> Call(sel, [20])
	call2 := &ast.CallExpression{
		Callee: selExpr,
		Args:   []ast.Expression{&ast.IntegerLiteral{Value: 20}},
	}
	assignRes2 := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "res2"},
		Op:    ast.OpAssignMut,
		Right: call2,
	}

	doc := &ast.Document{
		Expressions: []ast.Expression{
			assignRes1,
			assignRes2,
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

	if machine.SP != 3 {
		t.Errorf("Expected 3 values on stack, got %d", machine.SP)
	} else {
		r1 := machine.Stack[0]
		if r1.Type != mapval.ValInteger || r1.Integer != 100 {
			t.Errorf("Expected res1=100, got %v", r1)
		}
		r2 := machine.Stack[1]
		if r2.Type != mapval.ValInteger || r2.Integer != 21 {
			t.Errorf("Expected res2=21, got %v", r2)
		}
	}
}

func TestCompiler_Selector_Advanced(t *testing.T) {
	// res := 50 {
	//   _ >> 100 .. "High"
	//   _ >> 10  .. "Med"
	//   _        .. "Low" (Default)
	// }

	patterns := []ast.Pattern{
		&ast.PatternDefinition{
			Target: &ast.BinaryExpression{
				Left:  &ast.LabelLiteral{Value: "_"},
				Op:    ast.OpGt,
				Right: &ast.IntegerLiteral{Value: 100},
			},
			Action:    &ast.TextLiteral{Segments: []ast.TextSegment{&ast.StringSegment{Value: "High"}}},
			IsConsume: true,
		},
		&ast.PatternDefinition{
			Target: &ast.BinaryExpression{
				Left:  &ast.LabelLiteral{Value: "_"},
				Op:    ast.OpGt,
				Right: &ast.IntegerLiteral{Value: 10},
			},
			Action:    &ast.TextLiteral{Segments: []ast.TextSegment{&ast.StringSegment{Value: "Med"}}},
			IsConsume: true,
		},
		&ast.PatternDefault{
			Value: &ast.TextLiteral{Segments: []ast.TextSegment{&ast.StringSegment{Value: "Low"}}},
		},
	}

	selExpr := &ast.SelectorExpression{Patterns: patterns}

	call := &ast.CallExpression{
		Callee: selExpr,
		Args:   []ast.Expression{&ast.IntegerLiteral{Value: 50}},
	}

	assign := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "res"},
		Op:    ast.OpAssignMut,
		Right: call,
	}

	doc := &ast.Document{
		Expressions: []ast.Expression{assign},
	}

	c := compiler.NewCompiler()
	chunk, err := c.Compile(doc)
	if err != nil {
		t.Fatalf("Compilation failed: %v", err)
	}

	machine := vm.NewVM()
	machine.Interpret(chunk)

	// Stack: [res, res_value]
	// res (local 0) -> Stack[0]
	// res_value (expression result) -> Stack[1]

	if machine.SP != 2 {
		t.Errorf("Expected 2 values on stack, got %d", machine.SP)
	} else {
		val := machine.Stack[1]
		if val.Type != mapval.ValText || val.Str != "Med" {
			t.Errorf("Expected 'Med', got %v", val)
		}
	}
}

func TestCompiler_Selector_Binding(t *testing.T) {

	// Pinning not yet implemented (requires Upvalues).

	// Currently `target` in selector binds to NEW local.

	// target := 42

	// res := 42 {

	//   target .. "Match"

	//   _      .. "No Match"

	// }

	assignTarget := &ast.BinaryExpression{

		Left: &ast.LabelLiteral{Value: "target"},

		Op: ast.OpAssignMut,

		Right: &ast.IntegerLiteral{Value: 42},
	}

	patterns := []ast.Pattern{

		&ast.PatternDefinition{

			Target: &ast.LabelLiteral{Value: "target"}, // Binds new 'target'

			Action: &ast.TextLiteral{Segments: []ast.TextSegment{&ast.StringSegment{Value: "Match"}}},

			IsConsume: true,
		},

		&ast.PatternDefinition{

			Target: &ast.LabelLiteral{Value: "_"},

			Action: &ast.TextLiteral{Segments: []ast.TextSegment{&ast.StringSegment{Value: "No Match"}}},

			IsConsume: true,
		},
	}

	selExpr := &ast.SelectorExpression{Patterns: patterns}

	// 42 { ... }

	call1 := &ast.CallExpression{

		Callee: selExpr,

		Args: []ast.Expression{&ast.IntegerLiteral{Value: 42}},
	}

	assignRes := &ast.BinaryExpression{

		Left: &ast.LabelLiteral{Value: "res"},

		Op: ast.OpAssignMut,

		Right: call1,
	}

	// 99 { ... }

	call2 := &ast.CallExpression{

		Callee: selExpr,

		Args: []ast.Expression{&ast.IntegerLiteral{Value: 99}},
	}

	assignRes2 := &ast.BinaryExpression{

		Left: &ast.LabelLiteral{Value: "res2"},

		Op: ast.OpAssignMut,

		Right: call2,
	}

	doc := &ast.Document{

		Expressions: []ast.Expression{

			assignTarget,

			assignRes,

			assignRes2,
		},
	}

	c := compiler.NewCompiler()

	chunk, err := c.Compile(doc)

	if err != nil {

		t.Fatalf("Compilation failed: %v", err)

	}

	machine := vm.NewVM()

	machine.Interpret(chunk)

	if machine.SP != 4 {

		t.Errorf("Expected 4 values on stack, got %d", machine.SP)

	} else {

		res := machine.Stack[1]

		if res.Type != mapval.ValText || res.Str != "Match" {

			t.Errorf("Expected 'Match', got %v", res)

		}

		res2 := machine.Stack[2]

		// 99 matches 'target' (binding new var target=99)

		if res2.Type != mapval.ValText || res2.Str != "Match" {

			t.Errorf("Expected 'Match' (Binding), got %v", res2)

		}

	}

}
