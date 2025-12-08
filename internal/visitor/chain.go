package visitor

import (
	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/grammar"
	"github.com/antlr4-go/antlr/v4"
)

// --- Chain / Member ---

func (b *ASTBuilder) VisitMember(ctx *grammar.MemberContext) interface{} {
	// Delegated to ChainExpression
	return b.Visit(ctx.ChainExpression())
}

func (b *ASTBuilder) VisitChainExpression(ctx *grammar.ChainExpressionContext) interface{} {
	// "Fold Left" Strategy: Iterate children to build expression tree
	
	children := ctx.GetChildren()
	if len(children) == 0 {
		return nil
	}

	// Child 0 is FieldLiteral (Base)
	currentExpr := toExpr(b.Visit(children[0].(antlr.ParseTree)))

	// Helper to get field name
	getFieldText := func(n antlr.ParseTree) string {
		// Reuse visitFieldLiteral if accessible or just replicate logic
		// b.visitFieldLiteral returns ast.Node
		if litCtx, ok := n.(grammar.IFieldLiteralContext); ok {
			node := b.visitFieldLiteral(litCtx)
			if lit, ok := node.(*ast.LabelLiteral); ok {
				return lit.Value
			}
			if key, ok := node.(*ast.KeyLiteral); ok {
				return key.Value
			}
			if txt, ok := node.(*ast.TextLiteral); ok {
				return txt.Segments[0].String()
			}
		}
		return "?"
	}

	// Iterate Tail
	for i := 1; i < len(children); i++ {
		child := children[i]

		// CASE A: Chain Operator (\, @, #, ^, $)
		if opCtx, ok := child.(grammar.IChainOpContext); ok {
			// Next is Prefix (optional) then FieldLiteral
			i++
			if i >= len(children) { break }

			if _, isPrefix := children[i].(grammar.IPrefixOpContext); isPrefix {
				// TODO: Handle Prefix Logic (e.g. \!field)
				i++
			}

			if i >= len(children) { break }
			fieldLit := children[i] // FieldLiteral

			opType := ast.ChainMember
			switch opCtx.GetText() {
			case "\\": opType = ast.ChainMember
			case "@": opType = ast.ChainSubfield
			case "#": opType = ast.ChainSignal
			case "^": opType = ast.ChainReply
			case "$": opType = ast.ChainProclamation
			}
			
			ident := getFieldText(fieldLit.(antlr.ParseTree))

			// Flatten into ChainExpression if possible, else nest
			if c, ok := currentExpr.(*ast.ChainExpression); ok {
				c.Steps = append(c.Steps, ast.ChainStep{Op: opType, Ident: ident})
			} else {
				currentExpr = &ast.ChainExpression{
					Base: currentExpr,
					Steps: []ast.ChainStep{{Op: opType, Ident: ident}},
				}
			}
			continue
		}
		
		// CASE B: Open Bracket [ (Index or AccessOp)
		if token, ok := child.(*antlr.TerminalNodeImpl); ok && token.GetSymbol().GetTokenType() == grammar.RhumbParserOpenBracket {
			i++ // Consume [
			if i >= len(children) { break }
			
			// Skip terminators
			for {
				if _, ok := children[i].(grammar.ITerminatorContext); ok {
					i++
					if i >= len(children) { break }
				} else {
					break
				}
			}
			if i >= len(children) { break }

			if exprCtx, ok := children[i].(grammar.IExpressionContext); ok {
				// Index: [ expr ]
				idxExpr := toExpr(b.Visit(exprCtx.(antlr.ParseTree)))
				currentExpr = &ast.BinaryExpression{Left: currentExpr, Op: ast.OpIndex, Right: idxExpr}
				i++ // Consume expression
			} else if accessOpCtx, ok := children[i].(grammar.IAccessOpContext); ok {
				// Postfix: [ # ]
				op := b.Visit(accessOpCtx).(ast.OpType)
				// Represent Postfix as Binary(Expr, Op, Nil) per convention established in builder
				currentExpr = &ast.BinaryExpression{Left: currentExpr, Op: op, Right: nil}
				i++ // Consume Op
			} else if term, ok := children[i].(*antlr.TerminalNodeImpl); ok && term.GetSymbol().GetTokenType() == grammar.RhumbParserCloseBracket {
				// Empty []
				currentExpr = &ast.BinaryExpression{Left: currentExpr, Op: ast.OpIndex, Right: &ast.EmptyLiteral{}}
			}
			continue
		}

		// CASE C: Open Paren ( (Invocation)
		if token, ok := child.(*antlr.TerminalNodeImpl); ok && token.GetSymbol().GetTokenType() == grammar.RhumbParserOpenParen {
			i++ // Consume (
			var args []ast.Expression
			if i < len(children) {
				if exprsCtx, ok := children[i].(*grammar.ExpressionsContext); ok {
					args = b.Visit(exprsCtx).([]ast.Expression)
					i++ // Consume expressions
				}
			}
			currentExpr = &ast.CallExpression{Callee: currentExpr, Args: args}
			continue
		}
	}
	
	return currentExpr
}

// --- Other Expressions ---

func (b *ASTBuilder) VisitInvocation(ctx *grammar.InvocationContext) interface{} {
	// expression OpenParen ...
	target := toExpr(b.Visit(ctx.Expression()))
	
	var args []ast.Expression
	if ctx.Expressions() != nil {
		res := b.Visit(ctx.Expressions())
		if as, ok := res.([]ast.Expression); ok {
			args = as
		}
	}
	
	return &ast.CallExpression{
		Callee: target,
		Args:   args,
	}
}

func (b *ASTBuilder) VisitAccess(ctx *grammar.AccessContext) interface{} {
	// expression OpenBracket ...
	target := toExpr(b.Visit(ctx.Expression()))
	// Check AccessOp
	if ctx.AccessOp() != nil {
		// Helper to get opcode
		op := b.Visit(ctx.AccessOp()).(ast.OpType)
		return &ast.BinaryExpression{Left: target, Op: op, Right: nil}
	}
	
	// Indexing
	if ctx.Expressions() != nil {
		res := b.Visit(ctx.Expressions())
		if exprs, ok := res.([]ast.Expression); ok && len(exprs) > 0 {
			return &ast.BinaryExpression{Left: target, Op: ast.OpIndex, Right: exprs[0]}
		}
		return &ast.BinaryExpression{Left: target, Op: ast.OpIndex, Right: &ast.EmptyLiteral{}}
	}
	
	return target
}

func (b *ASTBuilder) VisitAssignLabel(ctx *grammar.AssignLabelContext) interface{} {
	// LHS is chainExpression or fieldLiteral
	var lhs ast.Expression
	if ctx.ChainExpression() != nil {
		lhs = toExpr(b.Visit(ctx.ChainExpression()))
	} else {
		node := b.visitFieldLiteral(ctx.FieldLiteral())
		if expr, ok := node.(ast.Expression); ok {
			lhs = expr
		}
	}
	
	rhs := toExpr(b.Visit(ctx.Expression()))
	
	var op ast.OpType
	switch ctx.AssignmentOp().(type) {
	case *grammar.MutableLabelContext: op = ast.OpAssignMut
	default: op = ast.OpAssignImm
	}
	
	return &ast.BinaryExpression{Left: lhs, Op: op, Right: rhs}
}

// --- AccessOp Visitors (OpType) ---

func (b *ASTBuilder) VisitAppend(ctx *grammar.AppendContext) interface{} { return ast.OpAppend }
func (b *ASTBuilder) VisitUnshift(ctx *grammar.UnshiftContext) interface{} { return ast.OpUnshift }
func (b *ASTBuilder) VisitLength(ctx *grammar.LengthContext) interface{} { return ast.OpLength }
func (b *ASTBuilder) VisitEmpty(ctx *grammar.EmptyContext) interface{} { return ast.OpIsEmpty }
func (b *ASTBuilder) VisitAllSubfields(ctx *grammar.AllSubfieldsContext) interface{} { return ast.OpAllSub }
func (b *ASTBuilder) VisitAllFields(ctx *grammar.AllFieldsContext) interface{} { return ast.OpAllFld }
func (b *ASTBuilder) VisitElements(ctx *grammar.ElementsContext) interface{} { return ast.OpAllPos }
func (b *ASTBuilder) VisitFreeze(ctx *grammar.FreezeContext) interface{} { return ast.OpFreeze }
func (b *ASTBuilder) VisitCopy(ctx *grammar.CopyContext) interface{} { return ast.OpCopy }
func (b *ASTBuilder) VisitToDate(ctx *grammar.ToDateContext) interface{} { return ast.OpToDate }
func (b *ASTBuilder) VisitParameters(ctx *grammar.ParametersContext) interface{} { return ast.OpGetParams }
func (b *ASTBuilder) VisitConstructor(ctx *grammar.ConstructorContext) interface{} { return ast.OpGetCtor }
func (b *ASTBuilder) VisitBase(ctx *grammar.BaseContext) interface{} { return ast.OpGetBase }
func (b *ASTBuilder) VisitToNumber(ctx *grammar.ToNumberContext) interface{} { return ast.OpToNum }
func (b *ASTBuilder) VisitNegateNumber(ctx *grammar.NegateNumberContext) interface{} { return ast.OpNegNum }
func (b *ASTBuilder) VisitToTruth(ctx *grammar.ToTruthContext) interface{} { return ast.OpToBool }
func (b *ASTBuilder) VisitNegateTruth(ctx *grammar.NegateTruthContext) interface{} { return ast.OpNegBool }
func (b *ASTBuilder) VisitVariadic(ctx *grammar.VariadicContext) interface{} { return ast.OpSpread }
func (b *ASTBuilder) VisitToKey(ctx *grammar.ToKeyContext) interface{} { return ast.OpToKey }
