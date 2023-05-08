package main

import (
	"context"
	"os"
	"os/signal"

	"git.sr.ht/~madcapjake/grhumb/internal/cli"
	"git.sr.ht/~madcapjake/grhumb/internal/generator"
	"git.sr.ht/~madcapjake/grhumb/internal/vm"
	"github.com/cristalhq/acmd"
)

var echo bool

func main() {
	cmds := []acmd.Command{
		{
			Name:        "file",
			Description: "evaluates file and any connected files",
			ExecFunc:    cli.InterpretFile,
		},
		{
			Name:        "line",
			Description: "evaluates lines and any connected files",
			ExecFunc:    cli.InterpretLines,
		},
		{
			Name:        "repl",
			Description: "read eval print loop",
			ExecFunc:    cli.ReadEvalPrintLoop,
		},
	}

	ctx := context.Background()
	ctx = context.WithValue(
		ctx,
		cli.VisitorCK,
		generator.NewRhumbVisitor(vm.NewVirtualMachine()),
	)
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	runner := acmd.RunnerOf(cmds, acmd.Config{
		AppName:        "rhumb",
		AppDescription: "Rhumb bytecode compiler",
		Version:        "0.1.0",
		Context:        ctx,
	})
	if err := runner.Run(); err != nil {
		runner.Exit(err)
	}
}
