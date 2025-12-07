package generator

import (
	"git.sr.ht/~madcapjake/rhi/internal/code"
	"git.sr.ht/~madcapjake/rhi/internal/object"
	P "git.sr.ht/~madcapjake/rhi/internal/parser"
	"github.com/antlr4-go/antlr/v4"
)

// Visit a parse tree produced by RhumbParser#conjunctive.
func (v *RhumbVisitor) VisitConjunctive(
	ctx *P.ConjunctiveContext,
) (any, error) {
	viLogger.Println("Conjunctive!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil, nil
}

// Visit a parse tree produced by RhumbParser#applicative.
func (v *RhumbVisitor) VisitApplicative(
	ctx *P.ApplicativeContext,
) (any, error) {
	viLogger.Println("Applicative!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil, nil
}

// Visit a parse tree produced by RhumbParser#comparative.
func (v *RhumbVisitor) VisitComparative(
	ctx *P.ComparativeContext,
) (any, error) {
	viLogger.Println("Comparative!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil, nil
}

// Visit a parse tree produced by RhumbParser#multiplicative.
func (v *RhumbVisitor) VisitMultiplicative(
	ctx *P.MultiplicativeContext,
) (any, error) {
	viLogger.Println("Multiplicative!")
	// v.visitBinaryLeaf(ctx.Expression(0))
	// v.visitBinaryLeaf(ctx.Expression(1))
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil, nil
}

// Visit a parse tree produced by RhumbParser#additive.
func (v *RhumbVisitor) VisitAdditive(
	ctx *P.AdditiveContext,
) (any, error) {
	viLogger.Println("Additive!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil, nil
}

// Visit a parse tree produced by RhumbParser#disjunctive.
func (v *RhumbVisitor) VisitDisjunctive(
	ctx *P.DisjunctiveContext,
) (any, error) {
	viLogger.Println("Disjunctive!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil, nil
}

// Visit a parse tree produced by RhumbParser#power.
func (v *RhumbVisitor) VisitPower(
	ctx *P.PowerContext,
) (any, error) {
	viLogger.Println("Power!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil, nil
}

// Visit a parse tree produced by RhumbParser#pipe.
func (v *RhumbVisitor) VisitPipe(ctx *P.PipeContext) (any, error) {
	viLogger.Println("Pipe!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#default.
func (v *RhumbVisitor) VisitDefault(ctx *P.DefaultContext) (any, error) {
	viLogger.Println("Default!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#addition.
func (v *RhumbVisitor) VisitAddition(ctx *P.AdditionContext) (any, error) {
	viLogger.Println("Addition!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#deviation.
func (v *RhumbVisitor) VisitDeviation(ctx *P.DeviationContext) (any, error) {
	viLogger.Println("Deviation!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#subtraction.
func (v *RhumbVisitor) VisitSubtraction(ctx *P.SubtractionContext) (any, error) {
	viLogger.Println("Subtraction!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#concatenate.
func (v *RhumbVisitor) VisitConcatenate(ctx *P.ConcatenateContext) (any, error) {
	viLogger.Println("Concatenate!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#multiplication.
func (v *RhumbVisitor) VisitMultiplication(ctx *P.MultiplicationContext) (any, error) {
	viLogger.Println("Multiplication!")
	viLogger.Println(ctx.GetText())

	op := v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, "_**_"),
	)
	s := ctx.GetStart()
	v.VM.Write(code.NewLocal(s.GetLine(), s.GetColumn(), op))
	return nil, nil
}

// Visit a parse tree produced by RhumbParser#rationalDivision.
func (v *RhumbVisitor) VisitRationalDivision(ctx *P.RationalDivisionContext) (any, error) {
	viLogger.Println("RationalDivision!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#wholeDivision.
func (v *RhumbVisitor) VisitWholeDivision(ctx *P.WholeDivisionContext) (any, error) {
	viLogger.Println("WholeDivision!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#remainderDivision.
func (v *RhumbVisitor) VisitRemainderDivision(ctx *P.RemainderDivisionContext) (any, error) {
	viLogger.Println("RemainderDivision!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#exponent.
func (v *RhumbVisitor) VisitExponent(ctx *P.ExponentContext) (any, error) {
	viLogger.Println("Exponent!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#rootExtraction.
func (v *RhumbVisitor) VisitRootExtraction(ctx *P.RootExtractionContext) (any, error) {
	viLogger.Println("RootExtraction!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#range.
func (v *RhumbVisitor) VisitRange(ctx *P.RangeContext) (any, error) {
	viLogger.Println("Range!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#scientific.
func (v *RhumbVisitor) VisitScientific(ctx *P.ScientificContext) (any, error) {
	viLogger.Println("Scientific!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}
