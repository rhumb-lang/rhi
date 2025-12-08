package testing

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/compiler"
	"git.sr.ht/~madcapjake/rhi/internal/config"
	"git.sr.ht/~madcapjake/rhi/internal/grammar"
	"git.sr.ht/~madcapjake/rhi/internal/parser_util" // Import the new package
	"git.sr.ht/~madcapjake/rhi/internal/vm"
	"git.sr.ht/~madcapjake/rhi/internal/visitor"
	"git.sr.ht/~madcapjake/rhi/internal/map"

	"github.com/antlr4-go/antlr/v4"
)

// RunRhumbTest executes a .rhs test file, interpreting code line by line
// and processing assertions defined with the %= meta-operator.
func RunRhumbTest(t *testing.T, path string) {
	t.Helper()

	code, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read test file: %v", err)
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("test panicked: %v", r)
		}
	}()

	// Initialize compiler and VM for the test file
	cfg := config.DefaultConfig() // No tracing for tests by default
	c := compiler.NewCompiler()
	v := vm.NewVM()
	v.Config = cfg
	
	errorListener := parser_util.NewRhumbErrorListener() // Use from new package

	lines := strings.Split(string(code), "\n")
	for lineNum, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "%") {
			continue // Skip empty lines and comments
		}

		if strings.Contains(trimmedLine, "%=") {
			// Assertion line
			parts := strings.SplitN(trimmedLine, "%=", 2)
			if len(parts) != 2 {
				t.Fatalf("Line %d: Malformed assertion: %s", lineNum+1, line)
			}
			lhsCode := strings.TrimSpace(parts[0])
			rhsCode := strings.TrimSpace(parts[1])

			// Evaluate LHS
			lhsResult, compileErr, runtimeErr := evaluateRhumbCode(t, c, v, lhsCode, errorListener)
			if compileErr != nil {
				t.Errorf("Line %d: LHS compilation error: %v in '%s'", lineNum+1, compileErr, lhsCode)
				continue
			}
			if runtimeErr != nil {
				t.Errorf("Line %d: LHS runtime error: %v in '%s'", lineNum+1, runtimeErr, lhsCode)
				continue
			}

			// Evaluate RHS
			rhsResult, compileErr, runtimeErr := evaluateRhumbCode(t, c, v, rhsCode, errorListener)
			if compileErr != nil {
				t.Errorf("Line %d: RHS compilation error: %v in '%s'", lineNum+1, compileErr, rhsCode)
				continue
			}
			if runtimeErr != nil {
				t.Errorf("Line %d: RHS runtime error: %v in '%s'", lineNum+1, runtimeErr, rhsCode)
				continue
			}

			// Compare results
			if !valuesEqual(lhsResult, rhsResult) {
				t.Errorf("Line %d: Assertion failed: '%s' (evaluated to %s) != '%s' (evaluated to %s)",
					lineNum+1, lhsCode, lhsResult.String(), rhsCode, rhsResult.String())
			}
		} else {
			// Regular code execution
			_, compileErr, runtimeErr := evaluateRhumbCode(t, c, v, trimmedLine, errorListener)
			if compileErr != nil {
				t.Errorf("Line %d: Compilation error: %v in '%s'", lineNum+1, compileErr, trimmedLine)
			}
			if runtimeErr != nil {
				t.Errorf("Line %d: Runtime error: %v in '%s'", lineNum+1, runtimeErr, trimmedLine)
			}
		}
	}
}

// evaluateRhumbCode compiles and runs a single line of Rhumb code.
// It returns the last value on the stack, or empty, and any errors.
func evaluateRhumbCode(t *testing.T, c *compiler.Compiler, v *vm.VM, code string, errorListener *parser_util.RhumbErrorListener) (mapval.Value, error, error) { // Updated parameter type
	t.Helper()
	errorListener.Errors = []string{} // Clear errors for each evaluation

	is := antlr.NewInputStream(code)
	lexer := grammar.NewRhumbLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := grammar.NewRhumbParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)
	
	tree := p.Document()
	
	if len(errorListener.Errors) > 0 {
		return mapval.NewEmpty(), fmt.Errorf("syntax errors: %v", errorListener.Errors), nil
	}

	builder := visitor.NewASTBuilder()
	astNode := builder.Visit(tree)
	doc, ok := astNode.(*ast.Document)
	if !ok {
		return mapval.NewEmpty(), fmt.Errorf("failed to build AST"), nil
	}

	startOffset, compileErr := c.CompileIncremental(doc)
	if compileErr != nil {
		return mapval.NewEmpty(), compileErr, nil
	}
    
	var res vm.Result
	var runtimeErr error

	if startOffset == 0 {
		res, runtimeErr = v.Interpret(c.Chunk())
	} else {
		res, runtimeErr = v.Continue(startOffset)
	}

	if runtimeErr != nil {
		return mapval.NewEmpty(), nil, runtimeErr
	}
	if res != vm.Ok {
		return mapval.NewEmpty(), nil, fmt.Errorf("VM exited with status: %v", res)
	}
    
    var lastValue mapval.Value
    if v.SP > 0 {
        lastValue = v.Stack[v.SP-1]
    } else {
        lastValue = mapval.NewEmpty()
    }
    
	return lastValue, nil, nil
}

// valuesEqual performs a deep equality check between two Rhumb values.
func valuesEqual(a, b mapval.Value) bool {
	if a.Type != b.Type {
		return false
	}
	switch a.Type {
	case mapval.ValInteger:
		return a.Integer == b.Integer
	case mapval.ValFloat:
		return a.Float == b.Float
	case mapval.ValText:
		return a.Str == b.Str
	case mapval.ValBoolean:
		return a.Integer == b.Integer
	case mapval.ValEmpty:
		return true // Both are empty
	case mapval.ValObject:
		// Deep comparison for maps would be complex. For now, rely on String() representation
		// or if we have mapval.Map.Equals()
		return a.String() == b.String()
	case mapval.ValRange:
		r1, ok1 := a.Obj.(*mapval.Range)
		r2, ok2 := b.Obj.(*mapval.Range)
		if !ok1 || !ok2 { return false }
		return r1.Start == r2.Start && r1.End == r2.End
	case mapval.ValVersion:
		return a.Integer == b.Integer
	case mapval.ValKey:
		return a.Integer == b.Integer
	default:
		return false
	}
}
