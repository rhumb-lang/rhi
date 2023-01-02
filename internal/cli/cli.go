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

	"git.sr.ht/~madcapjake/grhumb/internal/generator"
	"git.sr.ht/~madcapjake/grhumb/internal/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type visitorContextKey int
type disassembleContextKey int

const VisitorCK visitorContextKey = iota
const DisassembleCK disassembleContextKey = iota

func disassembleFlag(ctx *context.Context, args *[]string) ([]string, error) {
	flags := flag.NewFlagSet("rhumb", flag.ContinueOnError)
	var disassemble bool
	flags.BoolVar(&disassemble, "disassemble", false, "Disassemble the chunk of code")
	if err := flags.Parse(*args); err != nil {
		return []string{}, err
	}
	flags.Args()
	if disassemble {
		*ctx = context.WithValue(*ctx, DisassembleCK, true)
	} else {
		*ctx = context.WithValue(*ctx, DisassembleCK, false)
	}
	return flags.Args(), nil
}

func parse(ctx context.Context, chars antlr.CharStream) {
	lexer := parser.NewRhumbLexer(chars)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRhumbParser(stream)
	p.AddErrorListener(new(generator.RhumbErrorListener))
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	// TODO: BailErrorStrategy is broken right now
	// p.SetErrorHandler(antlr.NewBailErrorStrategy())
	// p.BuildParseTrees = true

	visitor := ctx.Value(VisitorCK).(*generator.RhumbVisitor)
	visitor.Visit(p.Sequence())
	vm := visitor.GetVM()
	if ctx.Value(DisassembleCK).(bool) {
		vm.Disassemble()
	}
	vm.Execute()
}

func ParseFile(ctx context.Context, args []string) error {
	var err error
	if args, err = disassembleFlag(&ctx, &args); err != nil {
		return err
	}
	input, err := antlr.NewFileStream(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating filestream: %s\n", err)
		os.Exit(1)
	}
	parse(ctx, input)
	return nil
}

func ParseLines(ctx context.Context, args []string) error {
	var err error
	if args, err = disassembleFlag(&ctx, &args); err != nil {
		return err
	}
	input := antlr.NewInputStream(strings.Join(args, "\n"))
	parse(ctx, input)
	return nil
}

func ReadEvalPrintLoop(ctx context.Context, args []string) error {
	var err error
	if args, err = disassembleFlag(&ctx, &args); err != nil {
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
		if err = ParseLines(ctx, []string{text}); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		}
	}
}
