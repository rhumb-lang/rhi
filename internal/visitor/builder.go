package visitor

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/grammar"
	"github.com/antlr4-go/antlr/v4"
	"github.com/cockroachdb/apd/v3"
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
	for _, e := range ctx.AllExpression() {
		res := b.Visit(e)
		if expr, ok := res.(ast.Expression); ok {
			if b.IsTestMode {
				if prc, ok := e.(antlr.ParserRuleContext); ok {
					op, content, found := b.getMetaOp(prc)
					if found {
						switch op {
						case "%=":
							// 1. Handle "Test Name" separator (%)
							parts := strings.SplitN(content, "%", 2)

							// 2. Treat the value strictly as a trimmed string (No Parsing!)
							rawString := strings.TrimSpace(parts[0])
							testName := ""
							if len(parts) > 1 {
								testName = strings.TrimSpace(parts[1])
							}

							expected := &ast.TextLiteral{
								IsRaw: true,
								Segments: []ast.TextSegment{
									&ast.StringSegment{Value: rawString},
								},
							}

							expr = &ast.AssertionWrapper{
								Actual:   expr,
								Expected: expected,
								Name:     testName,
							}
						case "%?":
							expr = &ast.InspectionWrapper{
								Expr: expr,
							}
						}
					}
				}
			}
			exprs = append(exprs, expr)
		}
	}
	return exprs
}

func (b *ASTBuilder) getMetaOp(ctx antlr.ParserRuleContext) (string, string, bool) {
	stop := ctx.GetStop()
	if stop == nil {
		return "", "", false
	}
	stopIndex := stop.GetTokenIndex()
	hiddenTokens := b.TokenStream.GetHiddenTokensToRight(stopIndex, antlr.TokenHiddenChannel)

	for _, t := range hiddenTokens {
		text := t.GetText()
		if after, ok := strings.CutPrefix(text, "%="); ok {
			return "%=", strings.TrimSpace(after), true
		}
		if after, ok := strings.CutPrefix(text, "%?"); ok {
			return "%?", strings.TrimSpace(after), true
		}
	}
	return "", "", false
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

	// Check for Decimal (Leading Zero distinction)
	// e.g. 01.5 (Decimal) vs 1.5 (Float) vs 0.5 (Float)
	// We check the part before the dot.
	parts := strings.Split(text, ".")
	integerPart := parts[0]
	isDecimal := len(integerPart) > 1 && strings.HasPrefix(integerPart, "0")

	if isDecimal {
		// Remove ".-" suffix if present for clean parsing
		cleanText := strings.TrimSuffix(text, ".-")

		d, _, err := apd.NewFromString(cleanText)
		if err != nil {
			panic(fmt.Errorf("error creating apd.Decimal from string: %s", err))
		}

		return &ast.DecimalLiteral{Value: d}
	}

	// Handle 10.- format (10.0) for standard floats
	if strings.HasSuffix(text, ".-") || strings.HasSuffix(text, ",-") {
		text = text[:len(text)-1] + "0"
	}

	val, _ := strconv.ParseFloat(text, 64)
	return &ast.RationalLiteral{Value: val}
}

func (b *ASTBuilder) VisitWholeNumber(ctx *grammar.WholeNumberContext) interface{} {
	text := ctx.GetText()

	// Check for ".-" suffix (Wildcard or Decimal Precision)
	if strings.HasSuffix(text, ".-") || strings.HasSuffix(text, ",-") {
		baseText := text[:len(text)-2]

		// Case 1: Decimal (Leading Zero) -> 01.-
		if len(baseText) > 1 && strings.HasPrefix(baseText, "0") {

			d, _, err := apd.NewFromString(baseText)
			if err != nil {
				panic(fmt.Errorf("error creating apd.Decimal from string: %s", err))
			}
			return &ast.DecimalLiteral{Value: d}
		}

		// Case 2: Version Wildcard -> 1.-
		// We treat this as a VersionLiteral with IsWildcard=true
		// 1.- means 1.*.*. Set Minor and Patch to MAX to indicate wildcard at Minor level (and below).
		val, _ := strconv.ParseUint(baseText, 10, 16)
		return &ast.VersionLiteral{
			Major:      uint16(val),
			Minor:      0xFFFF,     // MaxUint16
			Patch:      0xFFFFFFFF, // MaxUint32
			IsWildcard: true,
		}
	}

	val, _ := strconv.ParseInt(text, 10, 64)
	return &ast.IntegerLiteral{Value: val}
}

func (b *ASTBuilder) VisitVersionNumber(ctx *grammar.VersionNumberContext) interface{} {
	text := ctx.GetText()

	suffix := ""
	base := text

	// Detect suffix (starts with - or +)
	// Need to be careful not to confuse with ".-" wildcard
	// Regex would be easiest, but manual scan is fine.
	// Version structure is roughly: N.N.N-Suffix
	// But ".-" is valid at the end.
	if idx := strings.IndexAny(text, "-+"); idx != -1 {
		// Check if it is the ".-" wildcard at the end
		if idx == len(text)-1 && text[idx] == '-' && text[idx-1] == '.' {
			// It is ".-" wildcard, not a suffix
		} else if idx > 0 && text[idx] == '-' && text[idx-1] == '.' {
			// It matches ".-" pattern
		} else {
			// It is a real suffix
			suffix = text[idx:]
			base = text[:idx]
		}
	}

	isWildcard := strings.HasSuffix(base, ".-")
	if isWildcard {
		base = strings.TrimSuffix(base, ".-")
	}

	parts := strings.Split(strings.ReplaceAll(base, ",", "."), ".")
	v := &ast.VersionLiteral{IsWildcard: isWildcard, Suffix: suffix}

	if len(parts) > 0 {
		i, _ := strconv.ParseUint(parts[0], 10, 16)
		v.Major = uint16(i)
	}
	if len(parts) > 1 {
		i, _ := strconv.ParseUint(parts[1], 10, 16)
		v.Minor = uint16(i)
	} else if isWildcard {
		// If 1.- (though usually handled by WholeNumber), or 1. (invalid?)
		// If 1.-, Minor is Wild.
		v.Minor = 0xFFFF
	}
	
	if len(parts) > 2 {
		i, _ := strconv.ParseUint(parts[2], 10, 32)
		v.Patch = uint32(i)
	} else if isWildcard {
		// 1.2.- -> Patch is Wild.
		v.Patch = 0xFFFFFFFF
	}

	return v
}

func (b *ASTBuilder) VisitDateNumber(ctx *grammar.DateNumberContext) interface{} {
	text := ctx.GetText()
	// Formats: 2025/01/01, 2025/01/01@12:00:00, 2025/01/01@12:00:00.000

	if strings.Contains(text, "/00") { // Crude check for 00 month/day
		// Manual Parse
		parts := strings.Split(text, "@")
		datePart := parts[0]
		timePart := ""
		if len(parts) > 1 {
			timePart = parts[1]
		}

		dateParts := strings.Split(datePart, "/")
		if len(dateParts) == 3 {
			y, _ := strconv.ParseInt(dateParts[0], 10, 64)
			m, _ := strconv.ParseInt(dateParts[1], 10, 64)
			d, _ := strconv.ParseInt(dateParts[2], 10, 64)

			ms := int64(0)
			ms += y * 365 * 24 * 3600 * 1000
			ms += m * 30 * 24 * 3600 * 1000
			ms += d * 24 * 3600 * 1000

			if timePart != "" {
				t, err := time.Parse("15:04:05", timePart)
				if err == nil {
					// Add hours/min/sec from t (relative to 0000-01-01 00:00:00)
					ms += int64(t.Hour()) * 3600 * 1000
					ms += int64(t.Minute()) * 60 * 1000
					ms += int64(t.Second()) * 1000
					ms += int64(t.Nanosecond()) / 1000000
				}
			}

			return &ast.DateTimeLiteral{Value: ms}
		}
	}

	layouts := []string{
		"2006/01/02@15:04:05.000",
		"2006/01/02@15:04:05",
		"2006/01/02",
		"15:04:05",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, text)
		if err == nil {
			return &ast.DateTimeLiteral{Value: t.UnixMilli()}
		}
	}

	panic(fmt.Errorf("not able to convert value '%s' to DateTime", text))
}

func (b *ASTBuilder) VisitDurationNumber(ctx *grammar.DurationNumberContext) interface{} {
	text := ctx.GetText()
	// "Time Only" literal (e.g. 00:05:00)
	// Defaults to Unix Epoch Date: 1970/01/01
	layouts := []string{
		"15:04:05.000",
		"15:04:05",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, text)
		if err == nil {
			// Force date to 1970-01-01
			t = time.Date(1970, 1, 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
			return &ast.DateTimeLiteral{Value: t.UnixMilli()}
		}
	}

	return &ast.DateTimeLiteral{Value: 0}
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
