package visitor

import (
	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/grammar"
)

// --- Additional Binary Operations ---

func (b *ASTBuilder) VisitApplicative(ctx *grammar.ApplicativeContext) interface{} {
	left := toExpr(b.Visit(ctx.Expression(0)))
	right := toExpr(b.Visit(ctx.Expression(1)))
	opCtx := ctx.ApplicativeOp()

	var op ast.OpType
	switch opCtx.(type) {
	case *grammar.FunctionContext: op = ast.OpMakeFn
	case *grammar.MethodContext: op = ast.OpBindFn
	case *grammar.LetFunctionContext: op = ast.OpLetFn
	default: op = ast.OpMakeFn
	}

	return &ast.BinaryExpression{Left: left, Op: op, Right: right}
}

func (b *ASTBuilder) VisitConditional(ctx *grammar.ConditionalContext) interface{} {
	left := toExpr(b.Visit(ctx.Expression(0)))
	right := toExpr(b.Visit(ctx.Expression(1)))
	opCtx := ctx.ConditionalOp()

	var op ast.OpType
	switch opCtx.(type) {
	case *grammar.PipeContext: op = ast.OpPipe
	case *grammar.DefaultContext: op = ast.OpCoalesce
	case *grammar.ForeachContext: op = ast.OpForeach
	case *grammar.WhileContext: op = ast.OpWhile
	case *grammar.ThenContext: op = ast.OpIfTrue
	case *grammar.ElseContext: op = ast.OpIfFalse
	case *grammar.DestructureContext: op = ast.OpDestruct
	default: op = ast.OpIfTrue
	}

	return &ast.BinaryExpression{Left: left, Op: op, Right: right}
}

func (b *ASTBuilder) VisitPower(ctx *grammar.PowerContext) interface{} {
	left := toExpr(b.Visit(ctx.Expression(0)))
	right := toExpr(b.Visit(ctx.Expression(1)))
	opCtx := ctx.ExponentiationOp()

	var op ast.OpType
	switch opCtx.(type) {
	case *grammar.ExponentContext: op = ast.OpPow
	case *grammar.RootExtractionContext: op = ast.OpRoot
	case *grammar.RangeContext: op = ast.OpRange
	case *grammar.ScientificContext: op = ast.OpSciNot
	default: op = ast.OpPow
	}

	return &ast.BinaryExpression{Left: left, Op: op, Right: right}
}

func (b *ASTBuilder) VisitConjunctive(ctx *grammar.ConjunctiveContext) interface{} {
	left := toExpr(b.Visit(ctx.Expression(0)))
	right := toExpr(b.Visit(ctx.Expression(1)))
	// op is always And
	return &ast.BinaryExpression{Left: left, Op: ast.OpAnd, Right: right}
}

func (b *ASTBuilder) VisitDisjunctive(ctx *grammar.DisjunctiveContext) interface{} {
	left := toExpr(b.Visit(ctx.Expression(0)))
	right := toExpr(b.Visit(ctx.Expression(1)))
	// op is always Or
	return &ast.BinaryExpression{Left: left, Op: ast.OpOr, Right: right}
}

// --- Prefix / Effect ---

func (b *ASTBuilder) VisitPrefix(ctx *grammar.PrefixContext) interface{} {
	expr := toExpr(b.Visit(ctx.Expression()))
	opCtx := ctx.PrefixOp()
	
	var op ast.OpType
	switch opCtx.(type) {
	case *grammar.EmptyPrefixContext: op = ast.OpIsEmpty
	case *grammar.FreezePrefixContext: op = ast.OpFreeze
	case *grammar.CopyPrefixContext: op = ast.OpCopy
	case *grammar.ToNumberPrefixContext: op = ast.OpToNum
	case *grammar.NegateNumberPrefixContext: op = ast.OpNegNum
	case *grammar.ToTruthPrefixContext: op = ast.OpToBool
	case *grammar.NegateTruthPrefixContext: op = ast.OpNegBool
	case *grammar.VariadicPrefixContext: op = ast.OpSpread
	case *grammar.ArgumentPrefixContext: op = ast.OpGetParams
	case *grammar.SignalOutwardPrefixContext: op = ast.OpSignal
	case *grammar.SignalInwardPrefixContext: op = ast.OpReply
	default: op = ast.OpEq
	}
	
	return &ast.UnaryExpression{Op: op, Expr: expr}
}

func (b *ASTBuilder) VisitEffect(ctx *grammar.EffectContext) interface{} {
	// expression OpenCurly ... patterns ... CloseCurly
	target := toExpr(b.Visit(ctx.Expression()))
	
	sel := &ast.SelectorExpression{Patterns: []ast.Pattern{}}
	if ctx.Patterns() != nil {
		res := b.Visit(ctx.Patterns())
		if pats, ok := res.([]ast.Pattern); ok {
			sel.Patterns = pats
		}
	}
	
	return &ast.EffectExpression{
		Target:   target,
		Selector: sel,
	}
}
