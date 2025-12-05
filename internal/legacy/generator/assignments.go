package generator

import (
	P "git.sr.ht/~madcapjake/rhi/internal/parser"
	"github.com/antlr4-go/antlr/v4"
)

// Visit a parse tree produced by RhumbParser#assignLabel.
func (v *RhumbVisitor) VisitAssignLabel(
	ctx *P.AssignLabelContext,
) (any, error) {
	viLogger.Println("AssignLabel!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil, nil
}

// Visit a parse tree produced by RhumbParser#assignBreakingPattern.
func (v *RhumbVisitor) VisitAssignBreakingPattern(
	ctx *P.AssignBreakingPatternContext,
) (any, error) {
	viLogger.Println("AssignBreakingPattern!")
	viLogger.Println(ctx.GetText())
	panic("VisitAssignBreakingPattern not implemented")
}

// Visit a parse tree produced by RhumbParser#assignFallthroughPattern.
func (v *RhumbVisitor) VisitAssignFallthroughPattern(
	ctx *P.AssignFallthroughPatternContext,
) (any, error) {
	viLogger.Println("AssignFallthroughPattern!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#assignDefaultPattern.
func (v *RhumbVisitor) VisitAssignDefaultPattern(
	ctx *P.AssignDefaultPatternContext,
) (any, error) {
	viLogger.Println("AssignDefaultPattern!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#assignMutField.
func (v *RhumbVisitor) VisitAssignMutField(
	ctx *P.AssignMutFieldContext,
) (any, error) {
	viLogger.Println("AssignMutField!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#assignMutSubField.
func (v *RhumbVisitor) VisitAssignMutSubField(
	ctx *P.AssignMutSubFieldContext,
) (any, error) {
	viLogger.Println("AssignMutSubField!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#assignImmField.
func (v *RhumbVisitor) VisitAssignImmField(
	ctx *P.AssignImmFieldContext,
) (any, error) {
	viLogger.Println("AssignImmField!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#assignImmSubField.
func (v *RhumbVisitor) VisitAssignImmSubField(
	ctx *P.AssignImmSubFieldContext,
) (any, error) {
	viLogger.Println("AssignImmSubField!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}
