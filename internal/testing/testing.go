package testing

import (
	"fmt"
	"io/ioutil"
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

// RunRhumbTest executes a .rhs test file, interpreting the entire code
// and processing assertions defined with the %= meta-operator via the VM's test mode.
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
	
	errorListener := parser_util.NewRhumbErrorListener()

	// Execute the entire file contents
	_, compileErr, runtimeErr := evaluateRhumbCode(t, c, v, string(code), errorListener)
	
	if compileErr != nil {
		t.Errorf("Compilation error: %v", compileErr)
	}
	if runtimeErr != nil {
		t.Errorf("Runtime error: %v", runtimeErr)
	}
}

// evaluateRhumbCode compiles and runs Rhumb code.
// It returns the last value on the stack, or empty, and any errors.
func evaluateRhumbCode(t *testing.T, c *compiler.Compiler, v *vm.VM, code string, errorListener *parser_util.RhumbErrorListener) (mapval.Value, error, error) {
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

	builder := visitor.NewASTBuilder(stream, true)
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
	if res != vm.Ok && res != vm.Halt {
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
