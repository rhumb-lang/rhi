package compiler_test

import (
	"testing"

	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/compiler"
	mapval "github.com/rhumb-lang/rhi/internal/map"
	"github.com/rhumb-lang/rhi/internal/vm"
)

func TestCompiler_FizzBuzz(t *testing.T) {
	// fizzbuzz := [n] -> (
	//   list := []
	//   i := 1
	//   i <= n |> (
	//     val := i {
	//       _ -/ 15 == 0 .. "FizzBuzz"
	//       _ -/ 5 == 0 .. "Buzz"
	//       _ -/ 3 == 0 .. "Fizz"
	//       _ .. i
	//     }
	//     list := list ++ [val]
	//     i := i + 1
	//   )
	//   list
	// )
	// res := fizzbuzz(15)

	// Params [n]
	params := &ast.MapExpression{
		Fields: []ast.Field{
			&ast.FieldDefinition{
				Key:   &ast.LabelLiteral{Value: "n"},
				Value: &ast.EmptyLiteral{}, // value ignored for param definition
			},
		},
	}

	// Body Expressions
	// list := []
	initList := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "list"},
		Op:    ast.OpAssignMut,
		Right: &ast.MapExpression{},
	}

	// i := 1
	initI := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "i"},
		Op:    ast.OpAssignMut,
		Right: &ast.IntegerLiteral{Value: 1},
	}

	// Loop Condition: i <= n
	loopCond := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "i"},
		Op:    ast.OpLte,
		Right: &ast.LabelLiteral{Value: "n"},
	}

	// Selector for val
	// val := i { ... } -> Call(Selector, [i])
	selectorPatterns := []ast.Pattern{
		// _ -/ 15 == 0 .. "FizzBuzz"
		&ast.PatternDefinition{
			Target: &ast.BinaryExpression{
				Left: &ast.BinaryExpression{
					Left:  &ast.LabelLiteral{Value: "_"},
					Op:    ast.OpMod,
					Right: &ast.IntegerLiteral{Value: 15},
				},
				Op:    ast.OpEq,
				Right: &ast.IntegerLiteral{Value: 0},
			},
			Action:    &ast.TextLiteral{Segments: []ast.TextSegment{&ast.StringSegment{Value: "FizzBuzz"}}},
			IsConsume: true,
		},
		// _ -/ 5 == 0 .. "Buzz"
		&ast.PatternDefinition{
			Target: &ast.BinaryExpression{
				Left: &ast.BinaryExpression{
					Left:  &ast.LabelLiteral{Value: "_"},
					Op:    ast.OpMod,
					Right: &ast.IntegerLiteral{Value: 5},
				},
				Op:    ast.OpEq,
				Right: &ast.IntegerLiteral{Value: 0},
			},
			Action:    &ast.TextLiteral{Segments: []ast.TextSegment{&ast.StringSegment{Value: "Buzz"}}},
			IsConsume: true,
		},
		// _ -/ 3 == 0 .. "Fizz"
		&ast.PatternDefinition{
			Target: &ast.BinaryExpression{
				Left: &ast.BinaryExpression{
					Left:  &ast.LabelLiteral{Value: "_"},
					Op:    ast.OpMod,
					Right: &ast.IntegerLiteral{Value: 3},
				},
				Op:    ast.OpEq,
				Right: &ast.IntegerLiteral{Value: 0},
			},
			Action:    &ast.TextLiteral{Segments: []ast.TextSegment{&ast.StringSegment{Value: "Fizz"}}},
			IsConsume: true,
		},
		// Default: i
		// Use pattern definition with _ to match anything, or PatternDefault
		&ast.PatternDefault{
			Value: &ast.LabelLiteral{Value: "i"}, // Wait, i is in outer scope. Selector scope has _, but needs upvalue i?
			// My compiler doesn't support upvalues!
			// BUT, `_` is bound to Subject (which IS `i`).
			// So I can just return `_`.
		},
	}

	// Fix PatternDefault value to use `_` instead of `i` to avoid upvalue requirement.
	selectorPatterns[3] = &ast.PatternDefault{
		Value: &ast.LabelLiteral{Value: "_"},
	}

	selExpr := &ast.SelectorExpression{Patterns: selectorPatterns}
	callSel := &ast.CallExpression{
		Callee: selExpr,
		Args:   []ast.Expression{&ast.LabelLiteral{Value: "i"}},
	}

	calcVal := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "val"},
		Op:    ast.OpAssignMut,
		Right: callSel,
	}

	// list := list ++ [val]
	// [val] is MapExpression with 1 field (positional)
	wrapVal := &ast.MapExpression{
		Fields: []ast.Field{
			&ast.FieldElement{Value: &ast.LabelLiteral{Value: "val"}},
		},
	}
	appendList := &ast.BinaryExpression{
		Left: &ast.LabelLiteral{Value: "list"},
		Op:   ast.OpAssignMut,
		Right: &ast.BinaryExpression{
			Left:  &ast.LabelLiteral{Value: "list"},
			Op:    ast.OpAdd,
			Right: wrapVal,
		},
	}

	// i := i + 1
	incI := &ast.BinaryExpression{
		Left: &ast.LabelLiteral{Value: "i"},
		Op:   ast.OpAssignMut,
		Right: &ast.BinaryExpression{
			Left:  &ast.LabelLiteral{Value: "i"},
			Op:    ast.OpAdd,
			Right: &ast.IntegerLiteral{Value: 1},
		},
	}

	loopBody := &ast.RoutineExpression{
		Expressions: []ast.Expression{
			calcVal,
			appendList,
			incI,
		},
	}

	loop := &ast.BinaryExpression{
		Left:  loopCond,
		Op:    ast.OpWhile,
		Right: loopBody,
	}

	funcBody := &ast.RoutineExpression{
		Expressions: []ast.Expression{
			initList,
			initI,
			loop,
			&ast.LabelLiteral{Value: "list"},
		},
	}

	defFunc := &ast.BinaryExpression{
		Left:  params,
		Op:    ast.OpMakeFn,
		Right: funcBody,
	}

	assignFizz := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "fizzbuzz"},
		Op:    ast.OpAssignMut,
		Right: defFunc,
	}

	// res := fizzbuzz(15)
	callFizz := &ast.CallExpression{
		Callee: &ast.LabelLiteral{Value: "fizzbuzz"},
		Args:   []ast.Expression{&ast.IntegerLiteral{Value: 15}},
	}
	assignRes := &ast.BinaryExpression{
		Left:  &ast.LabelLiteral{Value: "res"},
		Op:    ast.OpAssignMut,
		Right: callFizz,
	}

	doc := &ast.Document{
		Expressions: []ast.Expression{
			assignFizz,
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

	// Expect: Stack has res (List)
	// Stack size depends on hoister. Locals: fizzbuzz (0), res (1).
	// Stack[1] is res.
	if machine.SP < 2 {
		t.Errorf("Stack too small: %d", machine.SP)
		return
	}

	resVal := machine.Stack[1]
	if resVal.Type != mapval.ValObject {
		t.Fatalf("Expected Map result, got %v", resVal)
	}
	resList := resVal.Obj.(*mapval.Map)

	if len(resList.Fields) != 15 {
		t.Errorf("Expected list length 15, got %d", len(resList.Fields))
	}

	// Verify some values
	// 1 (0): 1
	// 3 (2): Fizz
	// 5 (4): Buzz
	// 15 (14): FizzBuzz

	check := func(idx int, expected interface{}) {
		val := resList.Fields[idx]
		if str, ok := expected.(string); ok {
			if val.Type != mapval.ValText || val.Str != str {
				t.Errorf("Index %d: Expected %s, got %v", idx, str, val)
			}
		} else if i, ok := expected.(int); ok {
			if val.Type != mapval.ValInteger || val.Integer != int64(i) {
				t.Errorf("Index %d: Expected %d, got %v", idx, i, val)
			}
		}
	}

	check(0, 1)
	check(2, "Fizz")
	check(4, "Buzz")
	check(14, "FizzBuzz")
}
