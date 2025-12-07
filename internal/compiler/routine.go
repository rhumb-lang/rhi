package compiler

import (
	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

func (c *Compiler) compileRoutine(r *ast.RoutineExpression) error {
	if len(r.Expressions) == 0 {
		c.emitConstant(mapval.NewEmpty())
		return nil
	}
	for i, expr := range r.Expressions {
		if err := c.compileExpression(expr); err != nil {
			return err
		}
		// Pop result of previous expression if not last
		if i < len(r.Expressions)-1 {
			c.emit(mapval.OP_POP)
		}
	}
	return nil
}
