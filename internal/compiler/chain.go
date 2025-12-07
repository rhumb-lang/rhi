package compiler

import (
	"fmt"
	
	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

func (c *Compiler) compileChain(chain *ast.ChainExpression) error {
	// Compile Base
	if err := c.compileExpression(chain.Base); err != nil {
		return err
	}
	
	for _, step := range chain.Steps {
		switch step.Op {
		case ast.ChainMember:
			// \identifier -> OP_SEND identifier
			idx := c.makeConstant(mapval.NewText(step.Ident))
			c.emit(mapval.OP_SEND)
			c.Chunk().WriteByte(byte(idx), 0)
			
		case ast.ChainSubfield:
			// @identifier -> OP_SEND @identifier
			idx := c.makeConstant(mapval.NewText("@" + step.Ident))
			c.emit(mapval.OP_SEND)
			c.Chunk().WriteByte(byte(idx), 0)
			
		case ast.ChainSignal:
			// #identifier -> OP_POST identifier
			idx := c.makeConstant(mapval.NewText(step.Ident))
			c.emit(mapval.OP_POST)
			c.Chunk().WriteByte(byte(idx), 0)
			
		case ast.ChainReply:
			// ^identifier -> OP_INJECT identifier
			idx := c.makeConstant(mapval.NewText(step.Ident))
			c.emit(mapval.OP_INJECT)
			c.Chunk().WriteByte(byte(idx), 0)
			
		case ast.ChainProclamation:
			// $identifier -> OP_WRITE identifier
			idx := c.makeConstant(mapval.NewText(step.Ident))
			c.emit(mapval.OP_WRITE)
			c.Chunk().WriteByte(byte(idx), 0)
			
		default:
			return fmt.Errorf("unsupported chain op: %v", step.Op)
		}
	}
	return nil
}
