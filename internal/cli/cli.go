package cli

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"git.sr.ht/~madcapjake/grhumb/internal/generator"
	"git.sr.ht/~madcapjake/grhumb/internal/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type visitorContextKey int

const VisitorCK visitorContextKey = iota

func parse(visitor *generator.RhumbVisitor, chars antlr.CharStream) {
	lexer := parser.NewRhumbLexer(chars)

	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewRhumbParser(stream)
	p.AddErrorListener(new(generator.RhumbErrorListener))
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	// TODO: BailErrorStrategy is broken right now
	// p.SetErrorHandler(antlr.NewBailErrorStrategy())
	// p.BuildParseTrees = true

	tree := p.Sequence()
	visitor.Visit(tree)
	vm := visitor.GetVM()
	vm.Disassemble()
	vm.Execute()
}

func ParseFile(ctx context.Context, args []string) error {
	fmt.Println("Parsing file...")
	input, err := antlr.NewFileStream(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating filestream: %s\n", err)
		os.Exit(1)
	}

	parse(ctx.Value(VisitorCK).(*generator.RhumbVisitor), input)
	return nil
}

func ParseLines(ctx context.Context, args []string) error {
	fmt.Println("Parsing line...")
	input := antlr.NewInputStream(strings.Join(args, "\n"))
	parse(ctx.Value(VisitorCK).(*generator.RhumbVisitor), input)
	return nil
}

func ReadEvalPrintLoop(ctx context.Context, args []string) error {
	fmt.Println("REPL...")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("==>> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("")
				os.Exit(0)
			}
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
			os.Exit(1)
		}
		if err = ParseLines(ctx, []string{text}); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		}
	}
}
