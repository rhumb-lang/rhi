package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/compiler"
	"git.sr.ht/~madcapjake/rhi/internal/config"
	"git.sr.ht/~madcapjake/rhi/internal/grammar"
	mapval "git.sr.ht/~madcapjake/rhi/internal/map"
	"git.sr.ht/~madcapjake/rhi/internal/parser_util" // Import the new package
	"git.sr.ht/~madcapjake/rhi/internal/visitor"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
	"github.com/antlr4-go/antlr/v4"
)

var (
	traceParser   = flag.Bool("trace-parser", false, "Enable parser tracing")
	traceBytecode = flag.Bool("trace-bytecode", false, "Enable bytecode tracing")
	traceStack    = flag.Bool("trace-stack", false, "Enable stack tracing")
	traceSpace    = flag.Bool("trace-space", false, "Enable space/concurrency tracing")
	lastValue     = flag.Bool("last-value", false, "Print the last value of the execution")
	testMode      = flag.Bool("test", false, "Run in test mode (check assertions)")
)

type Session struct {
	Compiler *compiler.Compiler
	VM       *vm.VM
	IsRepl   bool
}

func NewSession(cfg *config.Config) *Session {
	v := vm.NewVM()
	v.Config = cfg
	return &Session{
		Compiler: compiler.NewCompiler(),
		VM:       v,
	}
}

func main() {
	flag.Parse()

	// Environment variable overrides
	if !*traceParser && os.Getenv("RHI_TRACE_PARSER") == "1" {
		*traceParser = true
	}
	if !*traceBytecode && os.Getenv("RHI_TRACE_BYTECODE") == "1" {
		*traceBytecode = true
	}
	if !*traceStack && os.Getenv("RHI_TRACE_STACK") == "1" {
		*traceStack = true
	}
	if !*traceSpace && os.Getenv("RHI_TRACE_SPACE") == "1" {
		*traceSpace = true
	}

	cfg := &config.Config{
		TraceParser:   *traceParser,
		TraceBytecode: *traceBytecode,
		TraceStack:    *traceStack,
		TraceSpace:    *traceSpace,
	}

	args := flag.Args()

	session := NewSession(cfg)

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
	fmt.Println("Rhumb REPL (v0.3.0-dev)")
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
	errorListener := parser_util.NewRhumbErrorListener() // Use from new package
	p.AddErrorListener(errorListener)

	tree := p.Document()

	if len(errorListener.Errors) > 0 {
		fmt.Println("Syntax Errors:")
		for _, msg := range errorListener.Errors {
			fmt.Println("  " + msg)
		}
		return
	}

	if s.VM.Config.TraceParser {
		fmt.Println("=== Parse Tree ===")
		fmt.Println(tree.ToStringTree(p.GetRuleNames(), p))
	}

	// 2. Build AST
	builder := visitor.NewASTBuilder(stream, *testMode)
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

	if s.VM.Config.TraceBytecode {
		fmt.Println("=== Bytecode Chunk ===")
		chunk := s.Compiler.Chunk()
		fmt.Printf("Constants: %v\n", chunk.Constants)

		fmt.Println("Disassembly:")
		for i := 0; i < len(chunk.Code); {
			fmt.Printf("%04d ", i)

			op := mapval.OpCode(chunk.Code[i])
			i++

			// Simple check for operands (assume most are 0 or constant index)
			// This is a basic disassembler, might not cover all variable-length ops perfectly
			// without full instruction set knowledge here, but helpful for debugging.
			// Reusing logic similar to VM's read would be best but for now just print OpName.

			fmt.Printf("%s", op)

			// Heuristic: If it's a known opcode taking an operand, print it.
			// Bank 0
			switch op {
			case mapval.OP_JUMP, mapval.OP_CALL, mapval.OP_MAKE_FN:
				if i+1 < len(chunk.Code) {
					// 2 byte operand usually for jump? Or 1 for const?
					// OP_MAKE_FN takes 1 byte const index.
					// OP_JUMP takes 2 bytes.
					if op == mapval.OP_JUMP || op == mapval.OP_JUMP_IF_FALSE || op == mapval.OP_JUMP_IF_TRUE || op == mapval.OP_WHILE {
						// 2 byte
						val := uint16(chunk.Code[i])<<8 | uint16(chunk.Code[i+1])
						fmt.Printf(" %d", val)
						i += 2
					} else {
						// 1 byte
						fmt.Printf(" %d", chunk.Code[i])
						i++
					}
				}
			case mapval.OP_LOAD_CONST, mapval.OP_LOAD_LOC, mapval.OP_STORE_LOC, mapval.OP_SEND, mapval.OP_SET_FIELD:
				// 1 byte operand
				if i < len(chunk.Code) {
					fmt.Printf(" %d", chunk.Code[i])
					// For Constants, maybe print the value?
					if op == mapval.OP_LOAD_CONST || op == mapval.OP_SEND || op == mapval.OP_SET_FIELD {
						idx := chunk.Code[i]
						if int(idx) < len(chunk.Constants) {
							fmt.Printf(" (%s)", chunk.Constants[idx])
						}
					}
					i++
				}
				if op == mapval.OP_SET_FIELD {
					// Takes another byte for flags?
					if i < len(chunk.Code) {
						fmt.Printf(" flags:%d", chunk.Code[i])
						i++
					}
				}
			}

			fmt.Println()
		}
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

	if res != vm.Ok && res != vm.Halt {
		fmt.Printf("VM exited with status: %s\n", res)
		return
	}

	// 5. Print Result and Reset Stack (keep locals)
	if s.IsRepl || *lastValue {
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
	// The Value.String() method already provides the canonical string representation for all types.
	return val.String()
}
