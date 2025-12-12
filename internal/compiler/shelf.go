package compiler

import (
	"strings"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	mapval "git.sr.ht/~madcapjake/rhi/internal/map"
)

func (c *Compiler) CompileShelf(sourceFiles []string, entryPoint string) (*mapval.Chunk, error) {
	var docs []*ast.Document

	// 1. Parse ALL files
	for _, src := range sourceFiles {
		// TODO: Parser logic
		docs = append(docs, ast)
	}
	if entryPoint != "" {
		// TODO: Parse entry point
		docs = append(docs, ast) // Add last
	}

	// 2. Global Hoisting (Pass 1)
	//    Scan every doc. Add ALL top-level assignments to c.Scope.
	for _, doc := range docs {
		locals := NewHoister().Hoist(doc)
		for _, name := range locals {
			if !strings.HasPrefix(name, "_") { // Only hoisting public/private logic?
				// Actually hoist EVERYTHING so internal files can see private vars.
				c.Scope.AddLocal(name)
				c.emitConstant(mapval.NewEmpty())
			}
		}
	}

	// 3. Compile Expressions (Pass 2)
	for _, doc := range docs {
		for _, expr := range doc.Expressions {
			c.compileExpression(expr)
			c.emit(mapval.OP_POP) // Discard statement results
		}
	}

	// 4. Auto-Export (The Harvest)
	//    Generate instructions to create a Map containing only Public variables.
	var exports []string
	for _, local := range c.Scope.Locals {
		if !strings.HasPrefix(local.Name, "_") {
			exports = append(exports, local.Name)
		}
	}

	// Emit [ key: val; key: val ]
	// Use `emitMapLiteral` helper or manual OP_MAKE_MAP sequence

	c.emit(mapval.OP_RETURN) // Return the export map
	return c.Chunk(), nil
}
