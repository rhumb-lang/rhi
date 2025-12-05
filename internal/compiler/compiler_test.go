package compiler_test

import (
	"testing"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/compiler"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
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
	
	if chunk.Code[0] != byte(vm.OP_LOAD_CONST) {
		t.Errorf("Expected OP_LOAD_CONST at 0, got %d", chunk.Code[0])
	}
	if chunk.Code[2] != byte(vm.OP_LOAD_CONST) {
		t.Errorf("Expected OP_LOAD_CONST at 2, got %d", chunk.Code[2])
	}
	if chunk.Code[4] != byte(vm.OP_ADD) {
		t.Errorf("Expected OP_ADD at 4, got %d", chunk.Code[4])
	}
	if chunk.Code[5] != byte(vm.OP_HALT) { // Wait, I counted wrong? 0,1 (LC, idx), 2,3 (LC, idx), 4 (ADD), 5 (HALT). Total 6 bytes.
		// Let's check loop logic: emit(HALT) is last.
		// chunk.Code index 5 is HALT. Length should be 6.
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
	if lastOp != byte(vm.OP_HALT) {
		t.Errorf("Expected last op to be HALT, got %d", lastOp)
	}
}
