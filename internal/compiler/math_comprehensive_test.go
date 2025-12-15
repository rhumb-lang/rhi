package compiler_test

import (
	"math"
	"testing"

	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/compiler"
	mapval "github.com/rhumb-lang/rhi/internal/map"
	"github.com/rhumb-lang/rhi/internal/vm"
)

func TestCompiler_MathOps(t *testing.T) {
	// Helper to create Boolean Literal Expr (1==1 or 1==0)
	boolExpr := func(val bool) ast.Expression {
		right := int64(0)
		if val {
			right = 1
		}
		return &ast.BinaryExpression{
			Left:  &ast.IntegerLiteral{Value: 1},
			Op:    ast.OpEq,
			Right: &ast.IntegerLiteral{Value: right},
		}
	}

	tests := []struct {
		name     string
		op       ast.OpType
		left     ast.Expression
		right    ast.Expression
		expected func(val mapval.Value) bool
	}{
		{
			name:  "IntDiv",
			op:    ast.OpDivInt,
			left:  &ast.IntegerLiteral{Value: 10},
			right: &ast.IntegerLiteral{Value: 3},
			expected: func(val mapval.Value) bool {
				return val.Type == mapval.ValInteger && val.Integer == 3
			},
		},
		{
			name:  "Mod",
			op:    ast.OpMod,
			left:  &ast.IntegerLiteral{Value: 10},
			right: &ast.IntegerLiteral{Value: 3},
			expected: func(val mapval.Value) bool {
				return val.Type == mapval.ValInteger && val.Integer == 1
			},
		},
		{
			name:  "Pow",
			op:    ast.OpPow,
			left:  &ast.IntegerLiteral{Value: 2},
			right: &ast.IntegerLiteral{Value: 3},
			expected: func(val mapval.Value) bool {
				return val.Type == mapval.ValFloat && val.Float == 8.0
			},
		},
		{
			name:  "Root",
			op:    ast.OpRoot,
			left:  &ast.IntegerLiteral{Value: 8},
			right: &ast.IntegerLiteral{Value: 3},
			expected: func(val mapval.Value) bool {
				// 8^(1/3) approx 2
				return val.Type == mapval.ValFloat && math.Abs(val.Float-2.0) < 0.000001
			},
		},
		{
			name:  "SciNot",
			op:    ast.OpSciNot,
			left:  &ast.RationalLiteral{Value: 1.5},
			right: &ast.IntegerLiteral{Value: 2}, // 1.5 * 10^2 = 150
			expected: func(val mapval.Value) bool {
				return val.Type == mapval.ValFloat && val.Float == 150.0
			},
		},
		{
			name:  "And_True",
			op:    ast.OpAnd,
			left:  boolExpr(true),
			right: boolExpr(true),
			expected: func(val mapval.Value) bool {
				return val.Type == mapval.ValBoolean && val.Integer == 1
			},
		},
		{
			name:  "And_False",
			op:    ast.OpAnd,
			left:  boolExpr(true),
			right: boolExpr(false),
			expected: func(val mapval.Value) bool {
				return val.Type == mapval.ValBoolean && val.Integer == 0
			},
		},
		{
			name:  "Or_True",
			op:    ast.OpOr,
			left:  boolExpr(false),
			right: boolExpr(true),
			expected: func(val mapval.Value) bool {
				return val.Type == mapval.ValBoolean && val.Integer == 1
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := &ast.Document{
				Expressions: []ast.Expression{
					&ast.BinaryExpression{
						Left:  tt.left,
						Op:    tt.op,
						Right: tt.right,
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
				t.Errorf("Expected Ok result, got %v", res)
			}

			if machine.SP < 1 {
				t.Errorf("Stack empty")
				return
			}

			val := machine.Stack[machine.SP-1]
			if !tt.expected(val) {
				t.Errorf("Unexpected result: %v", val)
			}
		})
	}
}
