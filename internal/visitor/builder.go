package visitor

import (
	"fmt"
	"strconv"
	"strings"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/grammar"
	"github.com/antlr4-go/antlr/v4"
)

// ASTBuilder converts the ANTLR parse tree into our internal AST.
type ASTBuilder struct {
	*grammar.BaseRhumbParserVisitor
	TokenStream *antlr.CommonTokenStream
	IsTestMode  bool
}

func (b *ASTBuilder) VisitDocument(ctx *grammar.DocumentContext) interface{} {
	// The root of the AST is a Document containing a list of expressions
	exprs := b.Visit(ctx.Expressions())

	// Safety check in case VisitExpressions returns nil
	if exprs == nil {
		return &ast.Document{Expressions: []ast.Expression{}}
	}

	return &ast.Document{
		Expressions: exprs.([]ast.Expression),
	}
}

func (b *ASTBuilder) VisitExpressions(ctx *grammar.ExpressionsContext) interface{} {
	var exprs []ast.Expression

	// Iterate over all children to preserve order and handle interleaved terminators
	// However, using ctx.AllExpression() is cleaner if we just want the expressions.
	// We need the Context for each expression to check for assertions.

	for _, exprCtx := range ctx.AllExpression() {
		res := b.Visit(exprCtx)
		if expr, ok := res.(ast.Expression); ok {

			// 1. Check for Assertion (Test Mode)
			// If -test flag is active, we look for the %= comment
			if b.IsTestMode {
				// We pass the rule context to find hidden tokens to its right
				if assertCode, found := b.getAssertion(exprCtx); found {
					// Parse the expected value (e.g. "10")
					expected := b.parseFragment(assertCode)

					// Wrap the expression in an AssertionWrapper
					// The Compiler will generate OP_ASSERT_EQ for this
					expr = &ast.AssertionWrapper{
						Actual:   expr,
						Expected: expected,
					}
				}
			}

			exprs = append(exprs, expr)
		}
	}

	return exprs
}

func (b *ASTBuilder) getAssertion(ctx antlr.ParserRuleContext) (string, bool) {
	stop := ctx.GetStop()
	if stop == nil {
		return "", false
	}
	stopIndex := stop.GetTokenIndex()
	hiddenTokens := b.TokenStream.GetHiddenTokensToRight(stopIndex, antlr.TokenHiddenChannel)

	// DEBUG PRINT
	fmt.Printf("DEBUG: getAssertion stopIndex=%d hiddenCount=%d Expr=%q\n", stopIndex, len(hiddenTokens), ctx.GetText())
	for i, t := range hiddenTokens {
		fmt.Printf("  Token[%d]: %q\n", i, t.GetText())
	}

	for _, t := range hiddenTokens {
		text := t.GetText()
		if after, ok := strings.CutPrefix(text, "%="); ok {
			return strings.TrimSpace(after), true
		}
	}
	return "", false
}

func NewASTBuilder(stream *antlr.CommonTokenStream, isTestMode bool) *ASTBuilder {
	return &ASTBuilder{
		BaseRhumbParserVisitor: &grammar.BaseRhumbParserVisitor{},
		TokenStream:            stream,
		IsTestMode:             isTestMode,
	}
}

func (b *ASTBuilder) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(b)
}

func (b *ASTBuilder) parseFragment(code string) ast.Expression {
	is := antlr.NewInputStream(code)
	lexer := grammar.NewRhumbLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := grammar.NewRhumbParser(stream)

	tree := parser.Expression()

	childBuilder := NewASTBuilder(stream, false)
	res := childBuilder.Visit(tree)

	if expr, ok := res.(ast.Expression); ok {
		return expr
	}
	// Fallback to empty if parse failed or returned something else
	return &ast.EmptyLiteral{}
}

// --- Specific Labeled Alternative Visitors ---

func (b *ASTBuilder) VisitSimpleExpression(ctx *grammar.SimpleExpressionContext) interface{} {
	return b.Visit(ctx.Literal())
}

func (b *ASTBuilder) VisitMap(ctx *grammar.MapContext) interface{} {
	mapExpr := &ast.MapExpression{Fields: []ast.Field{}}
	if ctx.Fields() != nil {
		res := b.Visit(ctx.Fields())
		if fields, ok := res.([]ast.Field); ok {
			mapExpr.Fields = fields
		}
	}
	return mapExpr
}

func (b *ASTBuilder) VisitRoutine(ctx *grammar.RoutineContext) interface{} {
	routine := &ast.RoutineExpression{Expressions: []ast.Expression{}}
	if ctx.Expressions() != nil {
		res := b.Visit(ctx.Expressions())
		if exprs, ok := res.([]ast.Expression); ok {
			routine.Expressions = exprs
		}
	}
	return routine
}

func (b *ASTBuilder) VisitSelector(ctx *grammar.SelectorContext) interface{} {
	sel := &ast.SelectorExpression{Patterns: []ast.Pattern{}}
	if ctx.Patterns() != nil {
		res := b.Visit(ctx.Patterns())
		if pats, ok := res.([]ast.Pattern); ok {
			sel.Patterns = pats
		}
	}
	return sel
}

func (b *ASTBuilder) VisitPatterns(ctx *grammar.PatternsContext) interface{} {
	var pats []ast.Pattern
	for _, p := range ctx.AllPattern() {
		res := b.Visit(p)
		if pat, ok := res.(ast.Pattern); ok {
			pats = append(pats, pat)
		}
	}
	return pats
}

func (b *ASTBuilder) VisitChildRealm(ctx *grammar.ChildRealmContext) interface{} {
	return &ast.ChildRealmLiteral{}
}

func (b *ASTBuilder) VisitDetachedRealm(ctx *grammar.DetachedRealmContext) interface{} {
	return &ast.DetachedRealmLiteral{}
}

// --- References ---

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

func (b *ASTBuilder) VisitRationalNumber(ctx *grammar.RationalNumberContext) interface{} {
	text := ctx.GetText()
	// Replace comma with dot for Go parsing if user used comma
	text = strings.Replace(text, ",", ".", 1)
	// Handle 10.- format (10.0)
	if strings.HasSuffix(text, ".-") || strings.HasSuffix(text, ",-") {
		text = text[:len(text)-1] + "0"
	}

	val, _ := strconv.ParseFloat(text, 64)
	return &ast.RationalLiteral{Value: val}
}

func (b *ASTBuilder) VisitWholeNumber(ctx *grammar.WholeNumberContext) interface{} {
	val, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return &ast.IntegerLiteral{Value: val}
}

func (b *ASTBuilder) VisitZeroNumber(ctx *grammar.ZeroNumberContext) interface{} {
	return &ast.IntegerLiteral{Value: 0}
}

func (b *ASTBuilder) VisitLabelSymbol(ctx *grammar.LabelSymbolContext) interface{} {
	return &ast.LabelLiteral{Value: ctx.GetText()}
}

func (b *ASTBuilder) VisitKeySymbol(ctx *grammar.KeySymbolContext) interface{} {
	// Remove backtick
	text := ctx.GetText()
	return &ast.KeyLiteral{Value: text[1:]}
}

func (b *ASTBuilder) VisitEmptyValue(ctx *grammar.EmptyValueContext) interface{} {
	return &ast.EmptyLiteral{}
}

func (b *ASTBuilder) VisitTextSymbol(ctx *grammar.TextSymbolContext) interface{} {
	// For now, just return a simple string literal wrapper.
	// We need to parse the interpolation parts properly later.
	// For MVP, let's just strip quotes.
	raw := ctx.GetText()
	isRaw := strings.HasPrefix(raw, "'")
	content := raw
	if isRaw {
		content = strings.Trim(raw, "'")
	} else {
		content = strings.Trim(raw, "\"")
	}

	return &ast.TextLiteral{
		IsRaw: isRaw,
		Segments: []ast.TextSegment{
			&ast.StringSegment{Value: content},
		},
	}
}

// --- Fields & Patterns ---

func (b *ASTBuilder) VisitFields(ctx *grammar.FieldsContext) interface{} {
	var fields []ast.Field
	for _, f := range ctx.AllField() {
		res := b.Visit(f)
		if field, ok := res.(ast.Field); ok {
			fields = append(fields, field)
		}
	}
	return fields
}

// --- Structural Visitors ---
// Moved to chain.go to avoid circular dependencies or massive files.
