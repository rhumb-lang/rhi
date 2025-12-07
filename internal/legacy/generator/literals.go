package generator

import (
	"strconv"

	"git.sr.ht/~madcapjake/rhi/internal/code"
	"git.sr.ht/~madcapjake/rhi/internal/object"
	P "git.sr.ht/~madcapjake/rhi/internal/parser"
)

// Visit a parse tree produced by RhumbParser#rationalNumber.
func (v *RhumbVisitor) VisitRationalNumber(
	ctx *P.RationalNumberContext,
) (any, error) {
	viLogger.Println("RationalNumber!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#dateNumber.
func (v *RhumbVisitor) VisitDateNumber(
	ctx *P.DateNumberContext,
) (any, error) {
	viLogger.Println("DateNumber!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#zeroNumber.
func (v *RhumbVisitor) VisitZeroNumber(
	ctx *P.ZeroNumberContext,
) (any, error) {
	viLogger.Println("ZeroNumber!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#wholeNumber.
func (v *RhumbVisitor) VisitWholeNumber(
	ctx *P.WholeNumberContext,
) (any, error) {
	viLogger.Println("WholeNumber!")
	viLogger.Println(ctx.GetText())
	num, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		panic("Integer conv error")
	}
	obj := object.NewNumber(v.VM.Memory, num)
	objID := v.VM.RegisterObject(obj)
	start := ctx.GetStart()
	v.VM.Write(code.NewValue(
		start.GetLine(),
		start.GetColumn(),
		objID,
	))
	return nil, nil
}

// Visit a parse tree produced by RhumbParser#keySymbol.
func (v *RhumbVisitor) VisitKeySymbol(
	ctx *P.KeySymbolContext,
) (any, error) {
	viLogger.Println("KeySymbol!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#textSymbol.
func (v *RhumbVisitor) VisitTextSymbol(
	ctx *P.TextSymbolContext,
) (any, error) {
	viLogger.Println("TextSymbol!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#referenceLiteral.
func (v *RhumbVisitor) VisitReferenceLiteral(
	ctx *P.ReferenceLiteralContext,
) (any, error) {
	viLogger.Println("ReferenceLiteral!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#labelSymbol.
func (v *RhumbVisitor) VisitLabelSymbol(
	ctx *P.LabelSymbolContext,
) (any, error) {
	viLogger.Println("LabelSymbol!")
	labelID := v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, ctx.GetText()),
	)
	start := ctx.GetStart()
	if _, ok := ctx.GetParent().GetParent().(*P.SimpleFieldContext); ok {
		v.VM.Write(code.NewValue(
			start.GetLine(),
			start.GetColumn(),
			labelID,
		))
	} else {
		v.VM.Write(code.NewLocal(
			start.GetLine(),
			start.GetColumn(),
			labelID,
		))
	}

	return nil, nil
}

// Visit a parse tree produced by RhumbParser#fieldLiteral.
func (v *RhumbVisitor) VisitFieldLiteral(
	ctx *P.FieldLiteralContext,
) (any, error) {

	viLogger.Println("FieldLiteral!")
	viLogger.Println(ctx.GetText())
	fieldId := v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, ctx.GetText()),
	)
	start := ctx.GetStart()
	v.VM.Write(code.NewValue(
		start.GetLine(),
		start.GetColumn(),
		fieldId,
	))
	return nil, nil
}

// Visit a parse tree produced by RhumbParser#map.
func (v *RhumbVisitor) VisitMap(
	ctx *P.MapContext,
) (any, error) {
	v.visitMapChildren(ctx.GetChildren(), "List", ctx)
	return nil, nil
}

// Visit a parse tree produced by RhumbParser#text.
func (v *RhumbVisitor) VisitText(
	ctx *P.TextContext,
) (any, error) {
	viLogger.Println("Text!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#date.
func (v *RhumbVisitor) VisitDate(
	ctx *P.DateContext,
) (any, error) {
	viLogger.Println("Date!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#datePart.
func (v *RhumbVisitor) VisitDatePart(
	ctx *P.DatePartContext,
) (any, error) {
	viLogger.Println("DatePart!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}
