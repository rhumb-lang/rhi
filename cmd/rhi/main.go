package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/compiler"
	"git.sr.ht/~madcapjake/rhi/internal/grammar"
	"git.sr.ht/~madcapjake/rhi/internal/visitor"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: rhi <file.rhs>")
		os.Exit(1)
	}

	filename := args[0]
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// 1. Parse
	is := antlr.NewInputStream(string(input))
	lexer := grammar.NewRhumbLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := grammar.NewRhumbParser(stream)
	
	// Custom Error Listener
	p.RemoveErrorListeners()
	errorListener := NewRhumbErrorListener()
	p.AddErrorListener(errorListener)
	
	tree := p.Document()
	
	if len(errorListener.Errors) > 0 {
		fmt.Println("Syntax Errors:")
		for _, msg := range errorListener.Errors {
			fmt.Println("  " + msg)
		}
		os.Exit(1)
	}

	// 2. Build AST
	builder := visitor.NewASTBuilder()
	astNode := builder.Visit(tree)
	if astNode == nil {
		fmt.Printf("Failed to build AST.\n")
		os.Exit(1)
	}
	
	doc, ok := astNode.(*ast.Document)
	if !ok {
		fmt.Printf("AST root is not a Document.\n")
		os.Exit(1)
	}
	
	// Debug AST
	fmt.Printf("Parsed %d expressions:\n", len(doc.Expressions))
	for i, expr := range doc.Expressions {
		fmt.Printf("%d: %s\n", i, expr.String())
	}

	// 3. Compile
	c := compiler.NewCompiler()
	chunk, err := c.Compile(doc)
	if err != nil {
		fmt.Printf("Compilation failed: %v\n", err)
		os.Exit(1)
	}

	// 4. Execute
	machine := vm.NewVM()
	res, err := machine.Interpret(chunk)
	if err != nil {
		fmt.Printf("Runtime error: %v\n", err)
		os.Exit(1)
	}
	
		if res != vm.Ok {
	
			fmt.Printf("VM exited with status: %v\n", res)
	
			os.Exit(1)
	
		}
	
		
	
		fmt.Printf("Stack Size: %d\n", machine.SP)
		if machine.SP > 0 {
	
			val := machine.Stack[machine.SP-1]
	
			fmt.Printf("=> %v\n", val) // Simple print
	
		}
	
		
	
		// Success
	
	}
	
	