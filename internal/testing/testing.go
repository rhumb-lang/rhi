package testing

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/compiler"
	"git.sr.ht/~madcapjake/rhi/internal/config"
	"git.sr.ht/~madcapjake/rhi/internal/grammar"
	mapval "git.sr.ht/~madcapjake/rhi/internal/map"
	"git.sr.ht/~madcapjake/rhi/internal/parser_util"
	"git.sr.ht/~madcapjake/rhi/internal/visitor"
	"git.sr.ht/~madcapjake/rhi/internal/vm"

	"github.com/antlr4-go/antlr/v4"
)

// RunRhumbTest executes a .rhs test file statement-by-statement
// and verifies assertions defined with the %= meta-operator.
func RunRhumbTest(t *testing.T, path string) {
	t.Helper()

	codeBytes, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read test file: %v", err)
	}
	code := string(codeBytes)

	// 1. Setup Parser
	is := antlr.NewInputStream(code)
	lexer := grammar.NewRhumbLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := grammar.NewRhumbParser(stream)

	// Error Listener
	el := parser_util.NewRhumbErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(el)

	// 2. Parse Document
	tree := p.Expressions() // Use Expressions rule to get list
	if len(el.Errors) > 0 {
		t.Fatalf("Syntax errors: %v", el.Errors)
	}

	// 3. Setup VM & Compiler (Persistent State)
	cfg := config.DefaultConfig()
	c := compiler.NewCompiler()
	v := vm.NewVM()
	v.Config = cfg
	builder := visitor.NewASTBuilder(stream, false) // False = don't print AST

	// 4. Iterate Statements
	children := tree.GetChildren()
	for _, child := range children {
		// Skip terminals (EOF, separators)
		if _, ok := child.(antlr.TerminalNode); ok {
			continue
		}

		// A. Build AST for this single statement
		astNode := builder.Visit(child.(antlr.ParseTree))

		// B. Check for Assertion Comment (%=)
		// We look at the tokens to the RIGHT of this node on the HIDDEN channel.
		stopIndex := child.GetPayload().(antlr.ParserRuleContext).GetStop().GetTokenIndex()
		hiddenTokens := stream.GetHiddenTokensToRight(stopIndex, antlr.TokenHiddenChannel)

		var expectedValue *string
		for _, token := range hiddenTokens {
			text := token.GetText()
			if strings.Contains(text, "%=") {
				// Parse " % = value " -> "value"
				parts := strings.SplitN(text, "%=", 2)
				if len(parts) == 2 {
					val := strings.TrimSpace(parts[1])
					expectedValue = &val
					break // Only handle one assertion per line
				}
			}
		}

		// C. Compile & Run
		// Wrap the single node in a temporary Document or compile directly
		// Note: Compiler needs to maintain scope, so we use CompileIncremental logic
		// effectively by compiling just this expression.

		// Hack: Compile expects a Document usually, but we can compile expressions.
		// Let's assume CompileIncremental handles a list of nodes.
		// We'll create a dummy doc with 1 expression.
		expr, ok := astNode.(ast.Expression)
		if !ok {
			continue // Skip non-expressions
		}

		dummyDoc := &ast.Document{Expressions: []ast.Expression{expr}}

		startOffset, err := c.CompileIncremental(dummyDoc)
		if err != nil {
			t.Fatalf("Compile error: %v", err)
		}

		var res vm.Result
		var runErr error

		if startOffset == 0 && len(c.Chunk().Code) > 0 {
			res, runErr = v.Interpret(c.Chunk())
		} else {
			res, runErr = v.Continue(startOffset)
		}

		if runErr != nil {
			t.Fatalf("Runtime error on line %d: %v", child.GetPayload().(antlr.ParserRuleContext).GetStart().GetLine(), runErr)
		}
		if res != vm.Ok && res != vm.Halt {
			t.Fatalf("VM crashed on line %d", child.GetPayload().(antlr.ParserRuleContext).GetStart().GetLine())
		}

		// D. Verify Assertion
		if expectedValue != nil {
			// Get actual value from stack
			var actual string
			if v.SP > 0 {
				val := v.Stack[v.SP-1]
				actual = val.String() // Assuming Value has a String() method

				// Handle String quotes for comparison
				// If expected is 'foo', actual might be "foo".
				// Let's rely on string representation.
				if val.Type == mapval.ValText {
					actual = fmt.Sprintf("'%s'", val.Str)
				}
			} else {
				actual = "___"
			}

			if actual != *expectedValue {
				t.Errorf("Assertion Failed on line %d.\nExpected: %s\nActual:   %s",
					child.GetPayload().(antlr.ParserRuleContext).GetStart().GetLine(),
					*expectedValue, actual)
			}
		}
	}
}
