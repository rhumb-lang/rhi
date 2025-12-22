package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	colorable "github.com/mattn/go-colorable"

	readline "github.com/nyaosorg/go-readline-ny"
	"github.com/nyaosorg/go-readline-ny/simplehistory"

	multiline "github.com/hymkor/go-multiline-ny"

	"github.com/rhumb-lang/rhi/internal/config"
)

func (s *Session) runFile(filename string) bool {
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return false
	}
	return s.execute(string(input))
}

func (s *Session) runREPL() {
	ctx := context.Background()
	fmt.Printf("Rhumb REPL (v%s)\n", config.RuntimeVersion.Canonical())
	fmt.Println("C-m or Enter      : Insert a linefeed")
	fmt.Println("C-p or UP         : Move to the previous line.")
	fmt.Println("C-n or DOWN       : Move to the next line")
	fmt.Println("C-j or Esc+Enter  : Submit")
	fmt.Println("C-c               : Abort.")
	fmt.Println("C-D with no chars : Quit.")
	fmt.Println("C-UP   or M-P     : Move to the previous history entry")
	fmt.Println("C-DOWN or M-N     : Move to the next history entry")
	ed, history := newREPL()
	for {
		lines, err := ed.Read(ctx)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}
		joinedLines := strings.Join(lines, "\n")
		fmt.Println("-----")
		_ = s.execute(joinedLines)
		fmt.Println("-----")
		history.Add(joinedLines)
	}
}

func newREPL() (ed multiline.Editor, history *simplehistory.Container) {
	ed.SetPrompt(func(w io.Writer, lnum int) (int, error) {
		return fmt.Fprintf(w, "[%d] ", lnum+1)
	})
	ed.SetPredictColor(readline.PredictColorBlueItalic)

	ed.Highlight = []readline.Highlight{
		// Words -> dark green
		{Pattern: regexp.MustCompile(`(?i)(yes|no)`), Sequence: "\x1B[33;49;22m"},
		// Double quotation -> light magenta
		{Pattern: regexp.MustCompile(`(?m)"([^"\n]*\\")*[^"\n]*$|"([^"\n]*\\")*[^"\n]*"`), Sequence: "\x1B[32;49;1m"},
		// Single quotation -> light red
		{Pattern: regexp.MustCompile(`(?m)'([^'\n]*\\')*[^'\n]*$|'([^'\n]*\\')*[^'\n]*'`), Sequence: "\x1B[31;49;1m"},
		// Number literal -> light blue
		{Pattern: regexp.MustCompile(`[0-9]+`), Sequence: "\x1B[34;49;1m"},
		// Inline Comment -> dark gray
		{Pattern: regexp.MustCompile(`(?s)%\(.*?%\)`), Sequence: "\x1B[30;49;1m"},
		// EOL Comment -> dark gray
		{Pattern: regexp.MustCompile(`(?s)%[^\(\)=\?].*`), Sequence: "\x1B[30;49;1m"},
	}
	ed.ResetColor = "\x1B[0m"
	ed.DefaultColor = "\x1B[37;49;1m"

	// To enable escape sequence on Windows.
	// (On other operating systems, it can be ommited)
	ed.SetWriter(colorable.NewColorableStdout())

	// enable history (optional)
	history = simplehistory.New()
	ed.SetHistory(history)
	ed.SetHistoryCycling(true)

	return
}
