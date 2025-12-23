package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/grammar"
	mapval "github.com/rhumb-lang/rhi/internal/map"
	"github.com/rhumb-lang/rhi/internal/parser_util"
	"github.com/rhumb-lang/rhi/internal/visitor"
	"github.com/rhumb-lang/rhi/internal/vm"
)

func (s *Session) execute(input string) bool {
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
		return false
	}

	if s.VM.Config.TraceParser {
		fmt.Println("=== Parse Tree ===")
		fmt.Println(tree.ToStringTree(p.GetRuleNames(), p))
	}

	// 2. Build AST
	builder := visitor.NewASTBuilder(stream, *testMode)
	astNode := builder.Visit(tree)
	if astNode == nil {
		return false
	}

	doc, ok := astNode.(*ast.Document)
	if !ok {
		return false
	}

	// 3. Compile Incremental
	startOffset, err := s.Compiler.CompileIncremental(doc)
	if err != nil {
		fmt.Printf("Compilation failed: %v\n", err)
		return false
	}

	if s.VM.Config.TraceBytecode {
		fmt.Println("=== Bytecode Trace ===")
		printChunk(s.Compiler.Chunk(), "Main")
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
		return false
	}

	if res != vm.Ok && res != vm.Halt {
		fmt.Printf("VM exited with status: %s\n", res)
		return false
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
	return true
}

func printChunk(chunk *mapval.Chunk, name string) {
	fmt.Printf("--- Chunk: %s ---\n", name)
	fmt.Printf("Constants: %v\n", chunk.Constants)
	fmt.Println("Disassembly:")
	for i := 0; i < len(chunk.Code); {
		fmt.Printf("%04d ", i)

		op := mapval.OpCode(chunk.Code[i])
		i++

		fmt.Printf("%s", op)

		switch op {
		case mapval.OP_JUMP, mapval.OP_CALL, mapval.OP_MAKE_FN, mapval.OP_JUMP_IF_FALSE, mapval.OP_JUMP_IF_TRUE, mapval.OP_WHILE:
			if op == mapval.OP_MAKE_FN {
				if i < len(chunk.Code) {
					idx := chunk.Code[i]
					fmt.Printf(" %d", idx)
					i++
					// Recurse!
					if int(idx) < len(chunk.Constants) {
						cVal := chunk.Constants[idx]
						if cVal.Type == mapval.ValObject {
							if fn, ok := cVal.Obj.(*mapval.Function); ok {
								fmt.Printf(" (<Function %s>)", fn.Name)
								// Recursively print function chunk
								fmt.Println()
								printChunk(fn.Chunk, fn.Name)
								fmt.Printf("--- End %s ---\n", fn.Name)
								
								// Skip Upvalue Descriptors
								i += fn.UpvalueCount * 2
							}
						}
					}
				}
			} else if op == mapval.OP_CALL {
				if i < len(chunk.Code) {
					fmt.Printf(" %d", chunk.Code[i])
					i++
				}
			} else {
				// Jump ops (2 bytes)
				if i+1 < len(chunk.Code) {
					val := uint16(chunk.Code[i])<<8 | uint16(chunk.Code[i+1])
					fmt.Printf(" %d", val)
					i += 2
				}
			}
		case mapval.OP_LOAD_CONST, mapval.OP_LOAD_LOC, mapval.OP_STORE_LOC, mapval.OP_SEND, mapval.OP_SET_FIELD, mapval.OP_LOAD_UPVALUE:
			if i < len(chunk.Code) {
				fmt.Printf(" %d", chunk.Code[i])
				if op == mapval.OP_LOAD_CONST {
					idx := chunk.Code[i]
					if int(idx) < len(chunk.Constants) {
						fmt.Printf(" (%s)", chunk.Constants[idx])
					}
				}
				i++
			}
			if op == mapval.OP_SET_FIELD {
				if i < len(chunk.Code) {
					fmt.Printf(" flags:%d", chunk.Code[i])
					i++
				}
			}
		}
		fmt.Println()
	}
}

func formatValue(val mapval.Value) string {
	return val.Canonical()
}