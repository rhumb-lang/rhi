package cli

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/compiler"
	"git.sr.ht/~madcapjake/rhi/internal/grammar"
	"git.sr.ht/~madcapjake/rhi/internal/visitor"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
	"github.com/antlr4-go/antlr/v4"
)

type disassembleContextKey int
type lastValueContextKey int

const DisassembleCK disassembleContextKey = iota
const LastValueCK lastValueContextKey = iota

func setupFlags(ctx *context.Context, args *[]string) ([]string, error) {
	flags := flag.NewFlagSet("rhumb", flag.ContinueOnError)
	var disassemble bool
	flags.BoolVar(&disassemble, "disassemble", false, "Disassemble the chunk of code")
	var lastValue bool
	flags.BoolVar(&lastValue, "last-value", false, "Print the last value of the main routine")
	if err := flags.Parse(*args); err != nil {
		return []string{}, err
	}
	flags.Args()
	if disassemble {
		*ctx = context.WithValue(*ctx, DisassembleCK, true)
	} else {
		*ctx = context.WithValue(*ctx, DisassembleCK, false)
	}
	if lastValue {
		*ctx = context.WithValue(*ctx, LastValueCK, true)
	} else {
		*ctx = context.WithValue(*ctx, LastValueCK, false)
	}
	return flags.Args(), nil
}

func interpret(ctx context.Context, chars antlr.CharStream) {
	// 1. Parse
	lexer := grammar.NewRhumbLexer(chars)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := grammar.NewRhumbParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	
	tree := p.Document()
	
	// 2. Build AST
	builder := visitor.NewASTBuilder()
	astNode := builder.Visit(tree)
	doc, ok := astNode.(*ast.Document)
	if !ok {
		fmt.Println("Error: Failed to build AST")
		return
	}

	// 3. Compile
	comp := compiler.NewCompiler()
	chunk, err := comp.Compile(doc)
	if err != nil {
		fmt.Printf("Compilation Error: %v\n", err)
		return
	}

	// Disassemble if requested
	if ctx.Value(DisassembleCK).(bool) {
		fmt.Println("=== Bytecode Chunk ===")
		fmt.Printf("Constants: %v\n", chunk.Constants)
		fmt.Printf("Code: %v\n", chunk.Code)
		// TODO: Better disassembly
	}

	// 4. Interpret

	machine := vm.NewVM()
	res, err := machine.Interpret(chunk)
	if err != nil {
		fmt.Printf("Runtime Error: %v\n", err)
	} else if res != vm.Ok {
		fmt.Printf("VM Result: %v\n", res)
	} else {
		// Check if we should print the last value
		if ctx.Value(LastValueCK).(bool) && machine.SP > 0 {
			// Peek top
			val := machine.Stack[machine.SP-1]
			// Simple print
			fmt.Printf("%v\n", val)
		}
	}
}

func InterpretFile(ctx context.Context, args []string) error {
	var err error
	if args, err = setupFlags(&ctx, &args); err != nil {
		return err
	}
	if len(args) < 1 {
		return fmt.Errorf("no file specified")
	}
	input, err := antlr.NewFileStream(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating filestream: %s\n", err)
		os.Exit(1)
	}
	interpret(ctx, input)
	return nil
}

func InterpretLines(ctx context.Context, args []string) error {
	var err error
	if args, err = setupFlags(&ctx, &args); err != nil {
		return err
	}
	interpretMany(ctx, args)
	return nil
}

func interpretMany(ctx context.Context, args []string) {
	input := antlr.NewInputStream(strings.Join(args, "\n"))
	interpret(ctx, input)
}

func ReadEvalPrintLoop(ctx context.Context, args []string) error {
	var err error
	if args, err = setupFlags(&ctx, &args); err != nil {
		return err
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("=>> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("")
				os.Exit(0)
			}
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
			os.Exit(1)
		}
		ctx = context.WithValue(ctx, LastValueCK, true)
		interpretMany(ctx, []string{text})
	}
}