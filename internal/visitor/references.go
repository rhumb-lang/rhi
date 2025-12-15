package visitor

import (
	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/grammar"
)

func (b *ASTBuilder) VisitReferenceLiteral(ctx *grammar.ReferenceLiteralContext) interface{} {
	return b.Visit(ctx.Reference())
}

func (b *ASTBuilder) VisitNamedRef(ctx *grammar.NamedRefContext) interface{} {
	return &ast.NamedReference{Label: ctx.Label().GetText()}
}

func (b *ASTBuilder) VisitFunctionRef(ctx *grammar.FunctionRefContext) interface{} {
	// <( ... )>
	// Inside is `expressions`.
	// We wrap it in RoutineExpression.
	// But `RoutineExpression` expects `RoutineContext`.
	// Here we have `ExpressionsContext`.
	// We can reuse logic or build manually.
	routine := &ast.RoutineExpression{Expressions: []ast.Expression{}}
	if ctx.Expressions() != nil {
		res := b.Visit(ctx.Expressions())
		if exprs, ok := res.([]ast.Expression); ok {
			routine.Expressions = exprs
		}
	}
	return &ast.RoutineReference{Routine: routine}
}

func (b *ASTBuilder) VisitVassalRef(ctx *grammar.VassalRefContext) interface{} {
	// <{ patterns? }>
	sel := &ast.SelectorExpression{Patterns: []ast.Pattern{}}
	if ctx.Patterns() != nil {
		res := b.Visit(ctx.Patterns())
		if pats, ok := res.([]ast.Pattern); ok {
			sel.Patterns = pats
		}
	}
	return &ast.VassalReference{Selector: sel}
}

func (b *ASTBuilder) VisitComputedRef(ctx *grammar.ComputedRefContext) interface{} {
	// < "text" > or < [ expr ] >
	if ctx.Text() != nil {
		// < "text" >
		// Visit text symbol
		res := b.Visit(ctx.Text())
		if txt, ok := res.(*ast.TextLiteral); ok {
			return &ast.ComputedTextReference{Text: txt}
		}
	} else if ctx.OpenBracket() != nil {
		// < [ expr ] >
		res := b.Visit(ctx.Expressions())
		// Wait, ComputedExpressionReference expects Single Expression?
		// Grammar: `expressions`. So it's a list?
		// AST `ComputedExpressionReference` has `Expression Expression`.
		// If multiple, implicitly last or tuple?
		// Let's assume Tuple/Sequence if multiple, or just take first.
		// For now, take first or wrap in Map if logic dictates.
		// Let's assume simple expression.
		if exprs, ok := res.([]ast.Expression); ok && len(exprs) > 0 {
			return &ast.ComputedExpressionReference{Expression: exprs[0]}
		}
	}
	return nil // Fallback
}
