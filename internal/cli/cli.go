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
	"git.sr.ht/~madcapjake/grhumb/internal/vm"
	"github.com/antlr4-go/antlr/v4"
)

type visitorContextKey int
type disassembleContextKey int
type lastValueContextKey int

const VisitorCK visitorContextKey = iota
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
	lexer := parser.NewRhumbLexer(chars)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRhumbParser(stream)
	p.AddErrorListener(new(generator.RhumbErrorListener))
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	// TODO: BailErrorStrategy is broken right now
	// p.SetErrorHandler(antlr.NewBailErrorStrategy())
	// p.BuildParseTrees = true

	visitor := ctx.Value(VisitorCK).(*generator.RhumbVisitor)
	visitor.Visit(p.Expressions())
	if ctx.Value(DisassembleCK).(bool) {
		vm.DisassembleMain(visitor.VM)
	}
	vm.RunMain(visitor.VM)
}

func InterpretFile(ctx context.Context, args []string) error {
	var err error
	if args, err = setupFlags(&ctx, &args); err != nil {
		return err
	}
	input, err := antlr.NewFileStream(args[1])
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
