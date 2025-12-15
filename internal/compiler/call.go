package compiler

import (
	"github.com/rhumb-lang/rhi/internal/ast"
	mapval "github.com/rhumb-lang/rhi/internal/map"
)

func (c *Compiler) compileCall(call *ast.CallExpression) error {
	// SPECIAL CASE: Space Operations (Signal/Reply/Proclamation with Args)
	// pattern: obj#signal(args) or obj^reply(args)
	if chain, ok := call.Callee.(*ast.ChainExpression); ok && len(chain.Steps) > 0 {
		lastStep := chain.Steps[len(chain.Steps)-1]
		if lastStep.Op == ast.ChainSignal || lastStep.Op == ast.ChainReply || lastStep.Op == ast.ChainProclamation {
			// 1. Compile Base
			if err := c.compileExpression(chain.Base); err != nil {
				return err
			}

			// 2. Compile Intermediate Steps (Property Access)
			for i := 0; i < len(chain.Steps)-1; i++ {
				step := chain.Steps[i]
				idx := c.makeConstant(mapval.NewText(step.Ident))
				switch step.Op {
				case ast.ChainMember:
					c.emit(mapval.OP_SEND)
				case ast.ChainSubfield:
					// @ident is treated as SEND "@ident"
					idx = c.makeConstant(mapval.NewText("@" + step.Ident))
					c.emit(mapval.OP_SEND)
				default:
					// Signal/Reply/Proclamation in middle of chain?
					// e.g. obj#sig\field
					// This effectively means "Post Signal, then access field on result (Empty)".
					// This matches chain.go logic.
					switch step.Op {
					case ast.ChainSignal:
						c.emit(mapval.OP_POST)
						c.Chunk().WriteByte(byte(idx), 0) // Signal Name
						c.Chunk().WriteByte(0, 0)         // Arg Count (0)
					case ast.ChainReply:
						c.emit(mapval.OP_INJECT)
						c.Chunk().WriteByte(byte(idx), 0)
						c.Chunk().WriteByte(0, 0)
					case ast.ChainProclamation:
						c.emit(mapval.OP_WRITE)
						c.Chunk().WriteByte(byte(idx), 0)
						c.Chunk().WriteByte(0, 0)
					}
					continue // Skip the WriteByte below as we handled it
				}
				c.Chunk().WriteByte(byte(idx), 0)
			}

			// 3. Compile Args
			for _, arg := range call.Args {
				if err := c.compileExpression(arg); err != nil {
					return err
				}
			}

			// 4. Emit Space Op with Arg Count
			nameIdx := c.makeConstant(mapval.NewText(lastStep.Ident))
			switch lastStep.Op {
			case ast.ChainSignal:
				c.emit(mapval.OP_POST)
			case ast.ChainReply:
				c.emit(mapval.OP_INJECT)
			case ast.ChainProclamation:
				c.emit(mapval.OP_WRITE)
			}
			c.Chunk().WriteByte(byte(nameIdx), 0)
			c.Chunk().WriteByte(byte(len(call.Args)), 0) // The Argument Count!

			return nil
		}
	}

	// STANDARD CALL

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
