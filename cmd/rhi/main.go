package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/compiler"
	"git.sr.ht/~madcapjake/rhi/internal/grammar"
	"git.sr.ht/~madcapjake/rhi/internal/visitor"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

type Session struct {
	Compiler *compiler.Compiler
	VM       *vm.VM
	IsRepl   bool
}

func NewSession() *Session {
	return &Session{
		Compiler: compiler.NewCompiler(),
		VM:       vm.NewVM(),
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	
session := NewSession()
	
	if len(args) == 0 {
		session.IsRepl = true
		session.runREPL()
	} else {
		session.runFile(args[0])
	}
}

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
	fmt.Println("Rhumb REPL (v0.1)")
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

func (s *Session) execute(input string) {
	// 1. Parse
	is := antlr.NewInputStream(input)
	lexer := grammar.NewRhumbLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := grammar.NewRhumbParser(stream)
	
p.RemoveErrorListeners()
	errorListener := NewRhumbErrorListener()
	p.AddErrorListener(errorListener)
	
tree := p.Document()
	
	if len(errorListener.Errors) > 0 {
		fmt.Println("Syntax Errors:")
		for _, msg := range errorListener.Errors {
			fmt.Println("  " + msg)
		}
		return
	}

	// 2. Build AST
	builder := visitor.NewASTBuilder()
	astNode := builder.Visit(tree)
	if astNode == nil {
		return
	}
	
doc, ok := astNode.(*ast.Document)
	if !ok {
		return
	}

	// 3. Compile Incremental
	startOffset, err := s.Compiler.CompileIncremental(doc)
	if err != nil {
		fmt.Printf("Compilation failed: %v\n", err)
		return
	}

	// 4. Execute
	var res vm.Result
	var runtimeErr error
	
	if startOffset == 0 {
		res, runtimeErr = s.VM.Interpret(s.Compiler.Chunk())
	} else {
		res, runtimeErr = s.VM.Continue(startOffset)
	}
	
	if runtimeErr != nil {
		fmt.Printf("Runtime error: %v\n", runtimeErr)
		return
	}
	
	if res != vm.Ok {
		fmt.Printf("VM exited with status: %v\n", res)
		return
	}
	
	// 5. Print Result and Reset Stack (keep locals)
	if s.IsRepl {
		// If stack has values beyond locals, print top
		numLocals := len(s.Compiler.Scope.Locals)
		if s.VM.SP > numLocals {
			val := s.VM.Stack[s.VM.SP-1]
			fmt.Printf("=> %s\n", formatValue(val))
		}
		
		// Reset SP to just locals
		// NOTE: If locals were added during this execution, they are on stack.
		// Any operands above them are discarded.
		if s.VM.SP > numLocals {
			s.VM.SP = numLocals
		}
	}
}

func formatValue(val mapval.Value) string {
	switch val.Type {
	case mapval.ValInteger:
		return fmt.Sprintf("%d", val.Integer)
	case mapval.ValFloat:
		return fmt.Sprintf("%f", val.Float)
	case mapval.ValText:
		return fmt.Sprintf("'%s'", val.Str)
	case mapval.ValBoolean:
		if val.Integer == 1 { return "yes" }
		return "no"
	case mapval.ValEmpty:
		return "___"
	case mapval.ValObject:
		if m, ok := val.Obj.(*mapval.Map); ok {
			var sb strings.Builder
			sb.WriteString("[")
			for i, field := range m.Fields {
				if i > 0 { sb.WriteString("; ") }
				sb.WriteString(formatValue(field))
			}
			sb.WriteString("]")
			return sb.String()
		}
		return fmt.Sprintf("<Object %T>", val.Obj)
	case mapval.ValRange:
		if r, ok := val.Obj.(*mapval.Range); ok {
			return fmt.Sprintf("%d|%d", r.Start, r.End)
		}
		return "<Range>"
	default:
		return fmt.Sprintf("?%d", val.Type)
	}
}
