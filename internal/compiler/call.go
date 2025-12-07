package compiler

import (
	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

func (c *Compiler) compileCall(call *ast.CallExpression) error {
	// 1. Compile Callee
	if err := c.compileExpression(call.Callee); err != nil {
		return err
	}
	
	// 2. Compile Args
	for _, arg := range call.Args {
		if err := c.compileExpression(arg); err != nil {
			return err
		}
	}
	
	// 3. Emit Call
	c.emit(mapval.OP_CALL)
	c.Chunk().WriteByte(byte(len(call.Args)), 0)
	return nil
}
