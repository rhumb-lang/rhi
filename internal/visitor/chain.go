package visitor

import (
	"git.sr.ht/~madcapjake/rhi/internal/ast"
	"git.sr.ht/~madcapjake/rhi/internal/grammar"
)

// --- Chain / Member ---

func (b *ASTBuilder) VisitMember(ctx *grammar.MemberContext) interface{} {
	// Delegated to ChainExpression
	return b.Visit(ctx.ChainExpression())
}

func (b *ASTBuilder) VisitChainExpression(ctx *grammar.ChainExpressionContext) interface{} {
	// Grammar:
	// chainExpression : fieldLiteral ( chainOp prefixOp? fieldLiteral | OpenBracket ... | OpenParen ... )+
	
	// We start with the base fieldLiteral
	baseNode := b.visitFieldLiteral(ctx.FieldLiteral(0))
	currentExpr, _ := baseNode.(ast.Expression)

	// We iterate manually because the structure is flat in the rule
	children := ctx.GetChildren()
	
	// Skip the first child (fieldLiteral)
	i := 1
	fieldLitIndex := 1 // Track which fieldLiteral we are on for context lookup
	
	for i < len(children) {
		child := children[i]
		
		if chainOp, ok := child.(grammar.IChainOpContext); ok {
			// It's a chain op. Type switch to find which one.
			i++ // Move past Op
			
			stepType := ast.ChainMember // Default to \ (NestedField)

			switch chainOp.(type) {
			case *grammar.NestedFieldContext:
				stepType = ast.ChainMember
			case *grammar.NestedSubfieldContext:
				stepType = ast.ChainSubfield
			case *grammar.SignalFieldContext:
				stepType = ast.ChainSignal
			case *grammar.ReplyFieldContext:
				stepType = ast.ChainReply
			case *grammar.ProclamationFieldContext:
				stepType = ast.ChainProclamation
			// TODO: Add DeeplyNested handlers if AST supports them
			}
			
			// Check for PrefixOp (optional)
			if i < len(children) {
				if _, isPrefix := children[i].(*grammar.PrefixOpContext); isPrefix {
					i++ // Skip prefix for now (TODO: Implement prefix logic in chain)
				}
			}
			
			// Next should be fieldLiteral
			if i < len(children) {
				if _, isLit := children[i].(*grammar.FieldLiteralContext); isLit {
					// We use GetFieldLiteral(index) because ctx provides typed accessors
					litCtx := ctx.FieldLiteral(fieldLitIndex)
					fieldLitIndex++
					
					ident := litCtx.GetText() // Simple text extraction for now
					
					// Append to Chain
					if chain, ok := currentExpr.(*ast.ChainExpression); ok {
						chain.Steps = append(chain.Steps, ast.ChainStep{Op: stepType, Ident: ident})
					} else {
						currentExpr = &ast.ChainExpression{
							Base:  currentExpr,
							Steps: []ast.ChainStep{{Op: stepType, Ident: ident}},
						}
					}
					i++
				} else {
					// Unexpected token after op
					i++
				}
			}
			
		} else if _, ok := child.(*grammar.ExpressionsContext); ok {
			// Handled via Invocation or similar, but here we skip to avoid crash
			i++
		} else {
			// OpenBracket/Paren tokens or Terminators
			i++
		}
	}
	
	return currentExpr
}

func (b *ASTBuilder) VisitAccess(ctx *grammar.AccessContext) interface{} {
	// expression OpenBracket ...
	target := toExpr(b.Visit(ctx.Expression()))
	// Check AccessOp
	if ctx.AccessOp() != nil {
		opCtx := ctx.AccessOp()
		var op ast.OpType
		
		switch opCtx.(type) {
		case *grammar.LengthContext: op = ast.OpLength
		case *grammar.EmptyContext: op = ast.OpIsEmpty
		case *grammar.AllSubfieldsContext: op = ast.OpAllSub
		case *grammar.AllFieldsContext: op = ast.OpAllFld
		case *grammar.ElementsContext: op = ast.OpAllPos
		case *grammar.FreezeContext: op = ast.OpFreeze
		case *grammar.CopyContext: op = ast.OpCopy
		case *grammar.ToDateContext: op = ast.OpToDate
		case *grammar.ParametersContext: op = ast.OpGetParams
		case *grammar.ConstructorContext: op = ast.OpGetCtor
		case *grammar.BaseContext: op = ast.OpGetBase
		case *grammar.ToNumberContext: op = ast.OpToNum
		case *grammar.NegateNumberContext: op = ast.OpNegNum
		case *grammar.ToTruthContext: op = ast.OpToBool
		case *grammar.NegateTruthContext: op = ast.OpNegBool
		case *grammar.VariadicContext: op = ast.OpSpread
		case *grammar.ToKeyContext: op = ast.OpToKey
		case *grammar.AppendContext: op = ast.OpAppend
		case *grammar.UnshiftContext: op = ast.OpUnshift
		default:
			// Fallback or error
			op = ast.OpEq
		}
		
		return &ast.UnaryExpression{Op: op, Expr: target}
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
