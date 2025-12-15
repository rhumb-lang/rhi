package compiler_test

import (
	"testing"

	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/compiler"
	mapval "github.com/rhumb-lang/rhi/internal/map"
)

func TestCompiler_SimpleMath(t *testing.T) {
	// AST for: 1 + 2
	doc := &ast.Document{
		Expressions: []ast.Expression{
			&ast.BinaryExpression{
				Left:  &ast.IntegerLiteral{Value: 1},
				Op:    ast.OpAdd,
				Right: &ast.IntegerLiteral{Value: 2},
			},
		},
	}

	c := compiler.NewCompiler()
	chunk, err := c.Compile(doc)
	if err != nil {
		t.Fatalf("Compilation failed: %v", err)
	}

	// Expected Bytecode:
	// LOAD_CONST 0 (1)
	// LOAD_CONST 1 (2)
	// ADD
	// HALT

	if len(chunk.Code) != 6 { // 2+2+1+1
		t.Errorf("Expected 6 bytes of code, got %d", len(chunk.Code))
	}

	if chunk.Code[0] != byte(mapval.OP_LOAD_CONST) {
		t.Errorf("Expected OP_LOAD_CONST at 0, got %d", chunk.Code[0])
	}
	if chunk.Code[2] != byte(mapval.OP_LOAD_CONST) {
		t.Errorf("Expected OP_LOAD_CONST at 2, got %d", chunk.Code[2])
	}
	if chunk.Code[4] != byte(mapval.OP_ADD) {
		t.Errorf("Expected OP_ADD at 4, got %d", chunk.Code[4])
	}
	if chunk.Code[5] != byte(mapval.OP_HALT) {
		t.Errorf("Expected OP_HALT at 5, got %d", chunk.Code[5])
	}
}

func TestCompiler_NestedMath(t *testing.T) {
	// AST for: (1 + 2) * 3
	doc := &ast.Document{
		Expressions: []ast.Expression{
			&ast.BinaryExpression{
				Left: &ast.BinaryExpression{
					Left:  &ast.IntegerLiteral{Value: 1},
					Op:    ast.OpAdd,
					Right: &ast.IntegerLiteral{Value: 2},
				},
				Op:    ast.OpMult,
				Right: &ast.IntegerLiteral{Value: 3},
			},
		},
	}

	c := compiler.NewCompiler()
	chunk, err := c.Compile(doc)
	if err != nil {
		t.Fatalf("Compilation failed: %v", err)
	}

	// 1. LOAD_CONST 1 (2 bytes)
	// 2. LOAD_CONST 2 (2 bytes)
	// 3. ADD (1 byte)
	// 4. LOAD_CONST 3 (2 bytes)
	// 5. MULT (1 byte)
	// 6. HALT (1 byte)
	// Total: 9 bytes

	if len(chunk.Code) != 9 {
		t.Errorf("Expected 9 bytes, got %d", len(chunk.Code))
	}

	lastOp := chunk.Code[len(chunk.Code)-1]
	if lastOp != byte(mapval.OP_HALT) {
		t.Errorf("Expected last op to be HALT, got %d", lastOp)
	}
}
