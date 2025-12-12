package compiler

import (
	"fmt"
	"os"
	"strings"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/grammar"
	mapval "git.sr.ht/~madcapjake/rhi/internal/map"
	"git.sr.ht/~madcapjake/rhi/internal/visitor"
	"github.com/antlr4-go/antlr/v4"
)

func (c *Compiler) CompileShelf(sourceFiles []string, entryPoint string) (*mapval.Chunk, error) {
	var docs []*ast.Document

	// 1. Parse ALL files
	allFiles := make([]string, len(sourceFiles))
	copy(allFiles, sourceFiles)
	if entryPoint != "" {
		allFiles = append(allFiles, entryPoint)
	}

	for _, path := range allFiles {
		content, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		is := antlr.NewInputStream(string(content))
		lexer := grammar.NewRhumbLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		parser := grammar.NewRhumbParser(stream)

		tree := parser.Document()
		builder := visitor.NewASTBuilder(stream, false)
		res := builder.Visit(tree)

		doc, ok := res.(*ast.Document)
		if !ok {
			return nil, fmt.Errorf("failed to parse document: %s", path)
		}
		docs = append(docs, doc)
	}

	// 2. Global Hoisting (Pass 1)
	//    Scan every doc. Add ALL top-level assignments to c.Scope.
	hoister := NewHoister()
	for _, doc := range docs {
		locals := hoister.Hoist(doc)
		for _, name := range locals {
			if !c.isDeclared(name) {
				c.Scope.addLocal(name)
				c.emitConstant(mapval.NewEmpty()) // Reserve slot
			}
		}
	}

	// 3. Compile Expressions (Pass 2)
	for _, doc := range docs {
		for _, expr := range doc.Expressions {
			if err := c.compileExpression(expr); err != nil {
				return nil, err
			}
			c.emit(mapval.OP_POP) // Discard top-level statement results
		}
	}

	// 4. Auto-Export (The Harvest)
	//    Generate instructions to create a Map containing only Public variables.
	var publicLocals []string
	for _, local := range c.Scope.Locals {
		if !strings.HasPrefix(local.Name, "_") {
			publicLocals = append(publicLocals, local.Name)
		}
	}

	c.emit(mapval.OP_MAKE_MAP)

	for _, name := range publicLocals {
		// 1. Dup Map (Receiver)
		c.emit(mapval.OP_DUP)

		// 2. Load Value
		idx := c.Scope.resolveLocal(name)
		c.emit(mapval.OP_LOAD_LOC)
		c.Chunk().WriteByte(byte(idx), 0)

		// 3. Set Field
		keyIdx := c.makeConstant(mapval.NewText(name))
		c.emit(mapval.OP_SET_FIELD)
		c.Chunk().WriteByte(byte(keyIdx), 0)
		c.Chunk().WriteByte(0, 0) // Immutable

		c.emit(mapval.OP_POP) // Pop result of Set
	}

	c.emit(mapval.OP_RETURN) // Return the export map
	return c.Chunk(), nil
}
