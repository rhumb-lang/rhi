package generator

import (
	P "git.sr.ht/~madcapjake/rhi/internal/parser"
)

func (v *RhumbVisitor) VisitExpressions(
	ctx *P.ExpressionsContext,
) (any, error) {
	viLogger.Println("expressions!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}
