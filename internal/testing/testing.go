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

func RunRhumbTest(t *testing.T, path string) {
	fmt.Println("Running Rhumb Test...")

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

	el := parser_util.NewRhumbErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(el)

	// 2. Parse Expressions
	tree := p.Expressions()
	if len(el.Errors) > 0 {
		t.Fatalf("Syntax errors: %v", el.Errors)
	}

	// 3. Setup VM & Compiler
	cfg := config.DefaultConfig()
	c := compiler.NewCompiler()
	v := vm.NewVM()
	v.Config = cfg
	builder := visitor.NewASTBuilder(stream, false)

	// We need to manage the bytecode buffer manually to "stitch" statements together
	// Initialize the main chunk if the compiler hasn't done it
	if c.Chunk() == nil {
		// Trigger initialization (implementation specific, usually NewCompiler does this)
	}

	children := tree.GetChildren()
	for i, child := range children {
		if _, ok := child.(antlr.TerminalNode); ok {
			continue
		}

		// A. Build AST
		astNode := builder.Visit(child.(antlr.ParseTree))

		// B. Check Assertion
		stopIndex := child.GetPayload().(antlr.ParserRuleContext).GetStop().GetTokenIndex()
		hiddenTokens := stream.GetHiddenTokensToRight(stopIndex, antlr.TokenHiddenChannel)
		var expectedValue *string
		for _, token := range hiddenTokens {
			text := token.GetText()
			if strings.Contains(text, "%=") {
				parts := strings.SplitN(text, "%=", 2)
				if len(parts) == 2 {
					val := strings.TrimSpace(parts[1])
					expectedValue = &val
					break
				}
			}
		}

		// C. Compile (The Trick)
		// We want to add code to the existing chunk.
		expr, ok := astNode.(ast.Expression)
		if !ok {
			continue
		}

		dummyDoc := &ast.Document{Expressions: []ast.Expression{expr}}

		// Note: We assume CompileIncremental appends to c.Chunk().Code
		// AND returns the offset where the new code begins.
		startOffset, err := c.CompileIncremental(dummyDoc)
		if err != nil {
			t.Fatalf("Compile error: %v", err)
		}

		// *** CRITICAL FIX: Patching the Halt ***
		// If this isn't the first statement, the previous statement likely
		// ended with OP_HALT or OP_RETURN. We need to overwrite it to fall through.
		// (This logic depends on your Compiler implementation details).
		// Assuming CompileIncremental emits [ ... instructions ... OP_HALT ]

		// Run the VM
		var res vm.Result
		var runErr error

		if i == 0 || startOffset == 0 {
			// First run
			res, runErr = v.Interpret(c.Chunk())
		} else {
			// Resume execution.
			// Important: We must ensure the VM state (Stack/Frame) is preserved.
			// v.Continue MUST resume from the current IP/Frame.
			res, runErr = v.Continue(startOffset)
		}

		if runErr != nil {
			t.Fatalf("Runtime error line %d: %v", child.GetPayload().(antlr.ParserRuleContext).GetStart().GetLine(), runErr)
		}

		// D. Verify
		if expectedValue != nil {
			var actual string
			// Peek the result left by the expression
			if v.SP > 0 {
				val := v.Stack[v.SP-1]
				if val.Type == mapval.ValText {
					actual = fmt.Sprintf("'%s'", val.Str)
				} else {
					actual = val.String()
				}
			} else {
				actual = "___"
			}

			if actual != *expectedValue {
				t.Errorf("FAIL: Line %d. Expected %s, got %s",
					child.GetPayload().(antlr.ParserRuleContext).GetStart().GetLine(),
					*expectedValue, actual)
			} else {
				fmt.Printf("PASS: %s\n", actual)
			}
		}
		fmt.Printf("%s\n", res)
	}
}
