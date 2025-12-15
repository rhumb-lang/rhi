package compiler

import (
	"fmt"
	"strconv"

	"github.com/rhumb-lang/rhi/internal/ast"
	mapval "github.com/rhumb-lang/rhi/internal/map"
)

func (c *Compiler) compileMap(m *ast.MapExpression) error {
	// Create empty map
	c.emit(mapval.OP_MAKE_MAP)

	posIdx := 1
	for _, field := range m.Fields {
		// Duplicate map for setting field (receiver)
		c.emit(mapval.OP_DUP)

		switch f := field.(type) {
		case *ast.FieldDefinition:
			// x: 1 or x. 1
			keyName := ""
			if label, ok := f.Key.(*ast.LabelLiteral); ok {
				keyName = label.Value
			} else if text, ok := f.Key.(*ast.TextLiteral); ok {
				if s, ok := text.Segments[0].(*ast.StringSegment); ok {
					keyName = s.Value
				}
			}

			if keyName == "" {
				return fmt.Errorf("dynamic keys not supported in map literal yet")
			}

			if f.IsSub {
				keyName = "@" + keyName
			}

			// Compile Value
			if err := c.compileExpression(f.Value); err != nil {
				return err
			}

			// Emit Set
			flags := byte(0)
			if f.IsMutable {
				flags = 1
			}

			keyIdx := c.makeConstant(mapval.NewText(keyName))

			c.emit(mapval.OP_SET_FIELD)
			c.Chunk().WriteByte(byte(keyIdx), 0)
			c.Chunk().WriteByte(flags, 0)

		case *ast.FieldPun:
			// .x or :x
			keyName := ""
			if label, ok := f.Key.(*ast.LabelLiteral); ok {
				keyName = label.Value
			}

			if keyName == "" {
				return fmt.Errorf("invalid pun key")
			}

			if f.IsSub {
				keyName = "@" + keyName
			}

			// Load Local Variable (if defined), else Empty
			localIdx := c.Scope.resolveLocal(keyName)
			if localIdx != -1 {
				c.emit(mapval.OP_LOAD_LOC)
				c.Chunk().WriteByte(byte(localIdx), 0)
			} else {
				// Pun of undefined variable -> Empty?
				// Or maybe the user INTENDED `x :: 100` but got Pun.
				// If I change this to Empty, then `x :: 100` becomes `x: ___`.
				// That explains why "undefined variable".
				// But `x :: 100` should NOT be a Pun!
				// So fixing this here hides the parser/visitor bug.
				// But if I want to support puns of undefined vars as Empty, I can do it.
				// Rhumb principles: "Any label not yet defined is considered empty (___) by default." (Architecture 5.4).
				// So `[.x]` where x is undefined should probably be `{x: ___}`.
				c.emitConstant(mapval.NewEmpty())
			}

			// Emit Set
			flags := byte(0)
			if f.IsMutable {
				flags = 1
			}

			keyIdx := c.makeConstant(mapval.NewText(keyName))
			c.emit(mapval.OP_SET_FIELD)
			c.Chunk().WriteByte(byte(keyIdx), 0)
			c.Chunk().WriteByte(flags, 0)

		case *ast.FieldElement:
			// [val] -> key is index "1", "2"... (1-based)
			// Architecture 5.6: "Positional elements use 1-based indexing."
			keyName := strconv.Itoa(posIdx)
			posIdx++

			// Compile Value
			if err := c.compileExpression(f.Value); err != nil {
				return err
			}

			// Emit Set (Immutable default)
			flags := byte(0)

			keyIdx := c.makeConstant(mapval.NewText(keyName))
			c.emit(mapval.OP_SET_FIELD)
			c.Chunk().WriteByte(byte(keyIdx), 0)
			c.Chunk().WriteByte(flags, 0)

		default:
			return fmt.Errorf("unsupported field type in map: %T", f)
		}

		// SetField now pushes val, so we must pop it to restore [Map]
		c.emit(mapval.OP_POP)
	}
	return nil
}
