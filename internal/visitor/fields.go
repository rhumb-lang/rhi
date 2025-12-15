package visitor

import (
	"fmt"

	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/grammar"
)

// Helper to extract literal node from fieldLiteral
// fieldLiteral : Zero | NumberPart | Key | text | reference | Bang | Label | TripleUnderscore | HookLabel
func (b *ASTBuilder) visitFieldLiteral(ctx grammar.IFieldLiteralContext) ast.Node {
	// Since this isn't a labeled alternative rule, we check children
	c := ctx.(*grammar.FieldLiteralContext)
	if c.Zero() != nil {
		return &ast.IntegerLiteral{Value: 0}
	}
	// TODO: Implement full literal extraction properly
	// For now, return a LabelLiteral with text to unblock
	return &ast.LabelLiteral{Value: c.GetText()}
}

func (b *ASTBuilder) VisitFieldLiteral(ctx *grammar.FieldLiteralContext) interface{} {
	return b.visitFieldLiteral(ctx)
}

// --- Fields ---

func (b *ASTBuilder) VisitPrefixAssignMutField(ctx *grammar.PrefixAssignMutFieldContext) interface{} {
	// Dot fieldLiteral -> .x (Immutable Shorthand)
	return &ast.FieldPun{
		Key:       b.visitFieldLiteral(ctx.FieldLiteral()),
		IsMutable: false, // . is Immutable
		IsSub:     false,
	}
}

func (b *ASTBuilder) VisitPrefixAssignMutSubField(ctx *grammar.PrefixAssignMutSubFieldContext) interface{} {
	// Dot At fieldLiteral -> .@x (Immutable Subfield Shorthand)
	return &ast.FieldPun{
		Key:       b.visitFieldLiteral(ctx.FieldLiteral()),
		IsMutable: false, // . is Immutable
		IsSub:     true,  // @ is Subfield
	}
}

func (b *ASTBuilder) VisitPrefixAssignImmSubField(ctx *grammar.PrefixAssignImmSubFieldContext) interface{} {
	// Colon fieldLiteral -> :x (Mutable Shorthand)
	// Note: Grammar label says ImmSubField but rule is `Colon fieldLiteral` or `Colon At fieldLiteral`.
	// We check for At manually or assume distinct methods if generated.
	// Since they share the label `prefixAssignImmSubField` in the grammar file provided earlier,
	// we need to handle both cases here.

	isSub := ctx.At() != nil
	return &ast.FieldPun{
		Key:       b.visitFieldLiteral(ctx.FieldLiteral()),
		IsMutable: true, // : is Mutable
		IsSub:     isSub,
	}
}

// Actually, let's implement a generic handler for the puns if possible, or specific ones.
// Since I can't see generated code, I'll implement specific ones based on the labels I saw.

func (b *ASTBuilder) VisitPrefixSlurpSpread(ctx *grammar.PrefixSlurpSpreadContext) interface{} {
	return &ast.FieldSpread{
		Key: b.visitFieldLiteral(ctx.FieldLiteral()),
	}
}

func (b *ASTBuilder) VisitAssignMutField(ctx *grammar.AssignMutFieldContext) interface{} {
	expr := ctx.Expression()
	if expr == nil {
		fmt.Println("Error: AssignMutField has nil expression context")
		return nil
	}

	valNode := b.Visit(expr)
	if valNode == nil {
		fmt.Println("Error: Visiting expression returned nil")
	}

	val := toExpr(valNode)

	return &ast.FieldDefinition{
		Key:       b.visitFieldLiteral(ctx.FieldLiteral()),
		Value:     val,
		IsMutable: true,
		IsSub:     false,
	}
}

func (b *ASTBuilder) VisitAssignMutSubField(ctx *grammar.AssignMutSubFieldContext) interface{} {
	return &ast.FieldDefinition{
		Key:       b.visitFieldLiteral(ctx.FieldLiteral()),
		Value:     toExpr(b.Visit(ctx.Expression())),
		IsMutable: true,
		IsSub:     true,
	}
}

func (b *ASTBuilder) VisitAssignImmField(ctx *grammar.AssignImmFieldContext) interface{} {
	return &ast.FieldDefinition{
		Key:       b.visitFieldLiteral(ctx.FieldLiteral()),
		Value:     toExpr(b.Visit(ctx.Expression())),
		IsMutable: false,
		IsSub:     false,
	}
}

func (b *ASTBuilder) VisitAssignImmSubField(ctx *grammar.AssignImmSubFieldContext) interface{} {
	return &ast.FieldDefinition{
		Key:       b.visitFieldLiteral(ctx.FieldLiteral()),
		Value:     toExpr(b.Visit(ctx.Expression())),
		IsMutable: false,
		IsSub:     true,
	}
}

func (b *ASTBuilder) VisitSimpleField(ctx *grammar.SimpleFieldContext) interface{} {
	return &ast.FieldElement{
		Value: toExpr(b.Visit(ctx.Expression())),
	}
}

// --- Patterns ---

func (b *ASTBuilder) VisitAssignBreakingPattern(ctx *grammar.AssignBreakingPatternContext) interface{} {
	return &ast.PatternDefinition{
		Target:    toExpr(b.Visit(ctx.Expression(0))),
		Action:    toExpr(b.Visit(ctx.Expression(1))),
		IsConsume: true, // ..
	}
}

func (b *ASTBuilder) VisitAssignFallthroughPattern(ctx *grammar.AssignFallthroughPatternContext) interface{} {
	return &ast.PatternDefinition{
		Target:    toExpr(b.Visit(ctx.Expression(0))),
		Action:    toExpr(b.Visit(ctx.Expression(1))),
		IsConsume: false, // ::
	}
}

func (b *ASTBuilder) VisitAssignDefaultPattern(ctx *grammar.AssignDefaultPatternContext) interface{} {
	return &ast.PatternDefault{
		Value: toExpr(b.Visit(ctx.Expression())),
	}
}
