package compiler

import (
	"fmt"
	"strconv"
	
	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/map"
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
			
			// Load Local Variable
			localIdx := c.Scope.resolveLocal(keyName)
			if localIdx == -1 {
				return fmt.Errorf("undefined variable in pun: %s", keyName)
			}
			c.emit(mapval.OP_LOAD_LOC)
			c.Chunk().WriteByte(byte(localIdx), 0)
			
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
	}
	
	return nil
}