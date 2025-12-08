package visitor_test

import (
	"testing"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/grammar"
	"git.sr.ht/~madcapjake/rhi/internal/visitor"
	"github.com/antlr4-go/antlr/v4"
)

func parseExpr(input string) ast.Expression {
	is := antlr.NewInputStream(input)
	lexer := grammar.NewRhumbLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := grammar.NewRhumbParser(stream)
	
	// We parse 'expression' rule directly for testing
	tree := p.Expression()
	builder := visitor.NewASTBuilder()
	return builder.Visit(tree).(ast.Expression)
}

func TestParser_ChainFolding(t *testing.T) {
	// 1. Simple Chain Flattening: a\b\c
	// Should be Chain(Base: a, Steps: [\b, \c])
	expr1 := parseExpr("a\\b\\c")
	
	chain1, ok := expr1.(*ast.ChainExpression)
	if !ok {
		t.Fatalf("Expected ChainExpression, got %T", expr1)
	}
	
	if len(chain1.Steps) != 2 {
		t.Errorf("Expected 2 steps, got %d", len(chain1.Steps))
	}
	if chain1.Steps[0].Ident != "b" || chain1.Steps[1].Ident != "c" {
		t.Errorf("Unexpected steps: %v", chain1.Steps)
	}
	
	// 2. Interrupted Chain: a\b[0]\c
	// Should be Chain(Base: Binary(Chain(a\b), Index, 0), Steps: [\c])
	expr2 := parseExpr("a\\b[0]\\c")
	
	chain2, ok := expr2.(*ast.ChainExpression)
	if !ok {
		t.Fatalf("Expected ChainExpression for outer, got %T", expr2)
	}
	
	// Check Steps of outer chain
	if len(chain2.Steps) != 1 {
		t.Errorf("Expected 1 step for outer chain, got %d", len(chain2.Steps))
	}
	if chain2.Steps[0].Ident != "c" {
		t.Errorf("Expected outer step 'c', got %s", chain2.Steps[0].Ident)
	}
	
	// Check Base of outer chain (Should be Binary Index)
	binExpr, ok := chain2.Base.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("Expected BinaryExpression base, got %T", chain2.Base)
	}
	if binExpr.Op != ast.OpIndex {
		t.Errorf("Expected OpIndex, got %v", binExpr.Op)
	}
	
	// Check Left of Binary (Should be Chain a\b)
	innerChain, ok := binExpr.Left.(*ast.ChainExpression)
	if !ok {
		t.Fatalf("Expected inner ChainExpression, got %T", binExpr.Left)
	}
	if len(innerChain.Steps) != 1 || innerChain.Steps[0].Ident != "b" {
		t.Errorf("Inner chain malformed")
	}
}
