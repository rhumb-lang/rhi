// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // RhumbParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseRhumbParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRhumbParserVisitor) VisitSequence(ctx *SequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitTerminator(ctx *TerminatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitReferenceLiteral(ctx *ReferenceLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitLabelLiteral(ctx *LabelLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitConjunctive(ctx *ConjunctiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAccess(ctx *AccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitApplicative(ctx *ApplicativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitConditional(ctx *ConditionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitPrefix(ctx *PrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitComparative(ctx *ComparativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitSimple(ctx *SimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitMultiplicative(ctx *MultiplicativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAdditive(ctx *AdditiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitInvocation(ctx *InvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitRoutine(ctx *RoutineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitDisjunctive(ctx *DisjunctiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitIdentity(ctx *IdentityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitEffect(ctx *EffectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitMember(ctx *MemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitPower(ctx *PowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitMap(ctx *MapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitFreeze(ctx *FreezeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitInner(ctx *InnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitLength(ctx *LengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitJunction(ctx *JunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitGreaterThan(ctx *GreaterThanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitGreaterThanOrEqualTo(ctx *GreaterThanOrEqualToContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitLessThan(ctx *LessThanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitLessThanOrEqualTo(ctx *LessThanOrEqualToContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitIsEqual(ctx *IsEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitIsLike(ctx *IsLikeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitIsIn(ctx *IsInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitIsOverlayed(ctx *IsOverlayedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitIsTopmost(ctx *IsTopmostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNotEqual(ctx *NotEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNotLike(ctx *NotLikeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNotIn(ctx *NotInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNotOverlayed(ctx *NotOverlayedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNotTopmost(ctx *NotTopmostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitConjunctiveOp(ctx *ConjunctiveOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitDisjunctiveOp(ctx *DisjunctiveOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitOtherwise(ctx *OtherwiseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitDefault(ctx *DefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitForeach(ctx *ForeachContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitWhile(ctx *WhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitThen(ctx *ThenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitElse(ctx *ElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAddition(ctx *AdditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitDeviation(ctx *DeviationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitSubtraction(ctx *SubtractionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitMultiplication(ctx *MultiplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitDivision(ctx *DivisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitIntegerDivision(ctx *IntegerDivisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitModulo(ctx *ModuloContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitBind(ctx *BindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitExponent(ctx *ExponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitRootExtraction(ctx *RootExtractionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitScientific(ctx *ScientificContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitImmutablePair(ctx *ImmutablePairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitImmutableLabel(ctx *ImmutableLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitMutablePair(ctx *MutablePairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitMutableLabel(ctx *MutableLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNumericalNegate(ctx *NumericalNegateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitOuterScope(ctx *OuterScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitLogicalNegate(ctx *LogicalNegateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAssert(ctx *AssertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitSlurpSpread(ctx *SlurpSpreadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitBaseClone(ctx *BaseCloneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNumericalPosit(ctx *NumericalPositContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitLogicalPosit(ctx *LogicalPositContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitOverlay(ctx *OverlayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitExistentialPosit(ctx *ExistentialPositContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitImmutableDestruct(ctx *ImmutableDestructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitMutableDestruct(ctx *MutableDestructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNestedLabel(ctx *NestedLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNestedOverlay(ctx *NestedOverlayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitRange(ctx *RangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitLabelInterp(ctx *LabelInterpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitRoutineInterp(ctx *RoutineInterpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitSelectorInterp(ctx *SelectorInterpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNamedRef(ctx *NamedRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitFunctionRef(ctx *FunctionRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitComputedRef(ctx *ComputedRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitJunctionRef(ctx *JunctionRefContext) interface{} {
	return v.VisitChildren(ctx)
}
