package visitor

import (
	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/grammar"
)

// Helper to safely cast interface{} to ast.Expression
func toExpr(v interface{}) ast.Expression {
	if e, ok := v.(ast.Expression); ok {
		return e
	}
	return nil
}

// --- Binary Operations ---

func (b *ASTBuilder) VisitAdditive(ctx *grammar.AdditiveContext) interface{} {
	left := toExpr(b.Visit(ctx.Expression(0)))
	right := toExpr(b.Visit(ctx.Expression(1)))
	opCtx := ctx.AdditiveOp()

	var op ast.OpType
	switch opCtx.(type) {
	case *grammar.AdditionContext:
		op = ast.OpAdd
	case *grammar.SubtractionContext:
		op = ast.OpSub
	case *grammar.ConcatenateContext:
		op = ast.OpConcat
	case *grammar.DeviationContext:
		op = ast.OpDev
	default:
		op = ast.OpAdd
	}

	return &ast.BinaryExpression{Left: left, Op: op, Right: right}
}

func (b *ASTBuilder) VisitMultiplicative(ctx *grammar.MultiplicativeContext) interface{} {
	left := toExpr(b.Visit(ctx.Expression(0)))
	right := toExpr(b.Visit(ctx.Expression(1)))
	opCtx := ctx.MultiplicativeOp()

	var op ast.OpType
	switch opCtx.(type) {
	case *grammar.MultiplicationContext:
		op = ast.OpMult
	case *grammar.RationalDivisionContext:
		op = ast.OpDivFloat
	case *grammar.WholeDivisionContext:
		op = ast.OpDivInt
	case *grammar.RemainderDivisionContext:
		op = ast.OpMod
	case *grammar.BindContext:
		op = ast.OpRebind
	default:
		op = ast.OpMult
	}

	return &ast.BinaryExpression{Left: left, Op: op, Right: right}
}

func (b *ASTBuilder) VisitComparative(ctx *grammar.ComparativeContext) interface{} {
	left := toExpr(b.Visit(ctx.Expression(0)))
	right := toExpr(b.Visit(ctx.Expression(1)))
	opCtx := ctx.ComparativeOp()

	var op ast.OpType
	switch opCtx.(type) {
	case *grammar.GreaterThanContext:
		op = ast.OpGt
	case *grammar.LessThanContext:
		op = ast.OpLt
	case *grammar.GreaterThanOrEqualToContext:
		op = ast.OpGte
	case *grammar.LessThanOrEqualToContext:
		op = ast.OpLte
	default:
		op = ast.OpGt
	}

	return &ast.BinaryExpression{Left: left, Op: op, Right: right}
}

func (b *ASTBuilder) VisitIdentity(ctx *grammar.IdentityContext) interface{} {
	left := toExpr(b.Visit(ctx.Expression(0)))
	right := toExpr(b.Visit(ctx.Expression(1)))
	opCtx := ctx.IdentityOp()

	var op ast.OpType
	switch opCtx.(type) {
	case *grammar.IsEqualContext:
		op = ast.OpEq
	case *grammar.NotEqualContext:
		op = ast.OpNeq
	case *grammar.IsUnderContext:
		op = ast.OpHasSub
	case *grammar.NotUnderContext:
		op = ast.OpNotHasSub
	case *grammar.IsInnerContext:
		op = ast.OpHasFld
	case *grammar.NotInnerContext:
		op = ast.OpNotHasFld
	default:
		op = ast.OpEq
	}

	return &ast.BinaryExpression{Left: left, Op: op, Right: right}
}

func (b *ASTBuilder) VisitAssignmentOp(ctx *grammar.AssignmentOpContext) interface{} {
	// This isn't called directly usually, handled in VisitAssignLabel
	return nil
}

// Note: ANTLR visitors return interface{}, so we build the AST node and return it.
// The parent node casts it.
