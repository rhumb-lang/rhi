package visitor_test

import (
	"testing"

	"git.sr.ht/~madcapjake/rhi/internal/grammar"
	"git.sr.ht/~madcapjake/rhi/internal/visitor"
	"github.com/antlr4-go/antlr/v4"
)

func TestParser_BasicExpression(t *testing.T) {
	input := "x .= 1"
	// Document.String() outputs "Expr.String()\n" for each expr.
	// Expr.String() for Binary is "(Left Op Right)".
	// So we expect "(x .= 1)\n"
	expected := "(x .= 1)\n"

	// Setup ANTLR
	is := antlr.NewInputStream(input)
	lexer := grammar.NewRhumbLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := grammar.NewRhumbParser(stream)

	tree := p.Document()
	builder := visitor.NewASTBuilder()
	astDoc := builder.Visit(tree)

	if astDoc == nil {
		t.Fatal("AST Builder returned nil")
	}

	str := astDoc.(interface{ String() string }).String()
	if str != expected {
		t.Errorf("Expected AST:\n%q\nGot:\n%q", expected, str)
	}
}

func TestParser_ComplexMap(t *testing.T) {
	input := `Point .= <(
		x := ___
  		y := ___
  		!\\set-x .= [x1] !> (x := x1)
	)>`

	is := antlr.NewInputStream(input)
	lexer := grammar.NewRhumbLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := grammar.NewRhumbParser(stream)

	tree := p.Document()
	builder := visitor.NewASTBuilder()
	astDoc := builder.Visit(tree)

	if astDoc == nil {
		t.Fatal("AST Builder returned nil for complex map")
	}

	t.Logf("Parsed Complex Map AST:\n%s", astDoc.(interface{ String() string }).String())
}
