package cli

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"git.sr.ht/~madcapjake/grhumb/internal/codegen"
	"git.sr.ht/~madcapjake/grhumb/internal/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type visitorContextKey int

const VisitorCK visitorContextKey = iota

func parse(visitor *codegen.RhumbVisitor, chars antlr.CharStream) {
	lexer := parser.NewRhumbLexer(chars)

	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewRhumbParser(stream)
	p.AddErrorListener(new(codegen.RhumbErrorListener))
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	// TODO: BailErrorStrategy is broken right now
	// p.SetErrorHandler(antlr.NewBailErrorStrategy())
	// p.BuildParseTrees = true

	tree := p.Sequence()

	visitor.Visit(tree)

}

func ParseFile(ctx context.Context, args []string) error {
	fmt.Println("Parsing file...")
	input, err := antlr.NewFileStream(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating filestream: %s\n", err)
		os.Exit(1)
	}
	parse(ctx.Value(VisitorCK).(*codegen.RhumbVisitor), input)
	return nil
}

func ParseLine(ctx context.Context, args []string) error {
	fmt.Println("Parsing line...")
	input := antlr.NewInputStream(args[0])
	parse(ctx.Value(VisitorCK).(*codegen.RhumbVisitor), input)
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
		if err = ParseLine(ctx, []string{text}); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		}
	}
}
