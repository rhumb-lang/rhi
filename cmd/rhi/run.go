package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rhumb-lang/rhi/internal/config"
)

func (s *Session) runFile(filename string) {
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	s.execute(string(input))
}

func (s *Session) runREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Rhumb REPL (v%s)\n", config.RuntimeVersion.Canonical())
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		s.execute(line)
	}
}
