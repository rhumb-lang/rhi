package generator

import (
	P "git.sr.ht/~madcapjake/rhi/internal/parser"
	"github.com/antlr4-go/antlr/v4"
)

// Visit a parse tree produced by RhumbParser#conditional.
func (v *RhumbVisitor) VisitConditional(
	ctx *P.ConditionalContext,
) (any, error) {
	viLogger.Println("Conditional!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil, nil
}

// Visit a parse tree produced by RhumbParser#while.
func (v *RhumbVisitor) VisitWhile(ctx *P.WhileContext) (any, error) {
	viLogger.Println("While!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#then.
func (v *RhumbVisitor) VisitThen(ctx *P.ThenContext) (any, error) {
	viLogger.Println("Then!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#else.
func (v *RhumbVisitor) VisitElse(ctx *P.ElseContext) (any, error) {
	viLogger.Println("Else!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}
