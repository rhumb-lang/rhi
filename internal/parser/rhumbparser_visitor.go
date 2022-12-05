// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // RhumbParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by RhumbParser.
type RhumbParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RhumbParser#sequence.
	VisitSequence(ctx *SequenceContext) interface{}

	// Visit a parse tree produced by RhumbParser#terminator.
	VisitTerminator(ctx *TerminatorContext) interface{}

	// Visit a parse tree produced by RhumbParser#FloatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by RhumbParser#IntegerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by RhumbParser#StringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by RhumbParser#ReferenceLiteral.
	VisitReferenceLiteral(ctx *ReferenceLiteralContext) interface{}

	// Visit a parse tree produced by RhumbParser#LabelLiteral.
	VisitLabelLiteral(ctx *LabelLiteralContext) interface{}

	// Visit a parse tree produced by RhumbParser#conjunctive.
	VisitConjunctive(ctx *ConjunctiveContext) interface{}

	// Visit a parse tree produced by RhumbParser#access.
	VisitAccess(ctx *AccessContext) interface{}

	// Visit a parse tree produced by RhumbParser#applicative.
	VisitApplicative(ctx *ApplicativeContext) interface{}

	// Visit a parse tree produced by RhumbParser#conditional.
	VisitConditional(ctx *ConditionalContext) interface{}

	// Visit a parse tree produced by RhumbParser#prefix.
	VisitPrefix(ctx *PrefixContext) interface{}

	// Visit a parse tree produced by RhumbParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by RhumbParser#comparative.
	VisitComparative(ctx *ComparativeContext) interface{}

	// Visit a parse tree produced by RhumbParser#simple.
	VisitSimple(ctx *SimpleContext) interface{}

	// Visit a parse tree produced by RhumbParser#multiplicative.
	VisitMultiplicative(ctx *MultiplicativeContext) interface{}

	// Visit a parse tree produced by RhumbParser#additive.
	VisitAdditive(ctx *AdditiveContext) interface{}

	// Visit a parse tree produced by RhumbParser#invocation.
	VisitInvocation(ctx *InvocationContext) interface{}

	// Visit a parse tree produced by RhumbParser#routine.
	VisitRoutine(ctx *RoutineContext) interface{}

	// Visit a parse tree produced by RhumbParser#disjunctive.
	VisitDisjunctive(ctx *DisjunctiveContext) interface{}

	// Visit a parse tree produced by RhumbParser#identity.
	VisitIdentity(ctx *IdentityContext) interface{}

	// Visit a parse tree produced by RhumbParser#effect.
	VisitEffect(ctx *EffectContext) interface{}

	// Visit a parse tree produced by RhumbParser#member.
	VisitMember(ctx *MemberContext) interface{}

	// Visit a parse tree produced by RhumbParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by RhumbParser#power.
	VisitPower(ctx *PowerContext) interface{}

	// Visit a parse tree produced by RhumbParser#map.
	VisitMap(ctx *MapContext) interface{}

	// Visit a parse tree produced by RhumbParser#freeze.
	VisitFreeze(ctx *FreezeContext) interface{}

	// Visit a parse tree produced by RhumbParser#inner.
	VisitInner(ctx *InnerContext) interface{}

	// Visit a parse tree produced by RhumbParser#length.
	VisitLength(ctx *LengthContext) interface{}

	// Visit a parse tree produced by RhumbParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by RhumbParser#junction.
	VisitJunction(ctx *JunctionContext) interface{}

	// Visit a parse tree produced by RhumbParser#greaterThan.
	VisitGreaterThan(ctx *GreaterThanContext) interface{}

	// Visit a parse tree produced by RhumbParser#greaterThanOrEqualTo.
	VisitGreaterThanOrEqualTo(ctx *GreaterThanOrEqualToContext) interface{}

	// Visit a parse tree produced by RhumbParser#lessThan.
	VisitLessThan(ctx *LessThanContext) interface{}

	// Visit a parse tree produced by RhumbParser#lessThanOrEqualTo.
	VisitLessThanOrEqualTo(ctx *LessThanOrEqualToContext) interface{}

	// Visit a parse tree produced by RhumbParser#isEqual.
	VisitIsEqual(ctx *IsEqualContext) interface{}

	// Visit a parse tree produced by RhumbParser#isLike.
	VisitIsLike(ctx *IsLikeContext) interface{}

	// Visit a parse tree produced by RhumbParser#isIn.
	VisitIsIn(ctx *IsInContext) interface{}

	// Visit a parse tree produced by RhumbParser#isOverlayed.
	VisitIsOverlayed(ctx *IsOverlayedContext) interface{}

	// Visit a parse tree produced by RhumbParser#isTopmost.
	VisitIsTopmost(ctx *IsTopmostContext) interface{}

	// Visit a parse tree produced by RhumbParser#notEqual.
	VisitNotEqual(ctx *NotEqualContext) interface{}

	// Visit a parse tree produced by RhumbParser#notLike.
	VisitNotLike(ctx *NotLikeContext) interface{}

	// Visit a parse tree produced by RhumbParser#notIn.
	VisitNotIn(ctx *NotInContext) interface{}

	// Visit a parse tree produced by RhumbParser#notOverlayed.
	VisitNotOverlayed(ctx *NotOverlayedContext) interface{}

	// Visit a parse tree produced by RhumbParser#notTopmost.
	VisitNotTopmost(ctx *NotTopmostContext) interface{}

	// Visit a parse tree produced by RhumbParser#conjunctiveOp.
	VisitConjunctiveOp(ctx *ConjunctiveOpContext) interface{}

	// Visit a parse tree produced by RhumbParser#disjunctiveOp.
	VisitDisjunctiveOp(ctx *DisjunctiveOpContext) interface{}

	// Visit a parse tree produced by RhumbParser#otherwise.
	VisitOtherwise(ctx *OtherwiseContext) interface{}

	// Visit a parse tree produced by RhumbParser#default.
	VisitDefault(ctx *DefaultContext) interface{}

	// Visit a parse tree produced by RhumbParser#foreach.
	VisitForeach(ctx *ForeachContext) interface{}

	// Visit a parse tree produced by RhumbParser#while.
	VisitWhile(ctx *WhileContext) interface{}

	// Visit a parse tree produced by RhumbParser#then.
	VisitThen(ctx *ThenContext) interface{}

	// Visit a parse tree produced by RhumbParser#else.
	VisitElse(ctx *ElseContext) interface{}

	// Visit a parse tree produced by RhumbParser#addition.
	VisitAddition(ctx *AdditionContext) interface{}

	// Visit a parse tree produced by RhumbParser#deviation.
	VisitDeviation(ctx *DeviationContext) interface{}

	// Visit a parse tree produced by RhumbParser#subtraction.
	VisitSubtraction(ctx *SubtractionContext) interface{}

	// Visit a parse tree produced by RhumbParser#multiplication.
	VisitMultiplication(ctx *MultiplicationContext) interface{}

	// Visit a parse tree produced by RhumbParser#division.
	VisitDivision(ctx *DivisionContext) interface{}

	// Visit a parse tree produced by RhumbParser#integerDivision.
	VisitIntegerDivision(ctx *IntegerDivisionContext) interface{}

	// Visit a parse tree produced by RhumbParser#modulo.
	VisitModulo(ctx *ModuloContext) interface{}

	// Visit a parse tree produced by RhumbParser#bind.
	VisitBind(ctx *BindContext) interface{}

	// Visit a parse tree produced by RhumbParser#exponent.
	VisitExponent(ctx *ExponentContext) interface{}

	// Visit a parse tree produced by RhumbParser#rootExtraction.
	VisitRootExtraction(ctx *RootExtractionContext) interface{}

	// Visit a parse tree produced by RhumbParser#scientific.
	VisitScientific(ctx *ScientificContext) interface{}

	// Visit a parse tree produced by RhumbParser#immutablePair.
	VisitImmutablePair(ctx *ImmutablePairContext) interface{}

	// Visit a parse tree produced by RhumbParser#immutableLabel.
	VisitImmutableLabel(ctx *ImmutableLabelContext) interface{}

	// Visit a parse tree produced by RhumbParser#mutablePair.
	VisitMutablePair(ctx *MutablePairContext) interface{}

	// Visit a parse tree produced by RhumbParser#mutableLabel.
	VisitMutableLabel(ctx *MutableLabelContext) interface{}

	// Visit a parse tree produced by RhumbParser#numericalNegate.
	VisitNumericalNegate(ctx *NumericalNegateContext) interface{}

	// Visit a parse tree produced by RhumbParser#bindBase.
	VisitBindBase(ctx *BindBaseContext) interface{}

	// Visit a parse tree produced by RhumbParser#logicalNegate.
	VisitLogicalNegate(ctx *LogicalNegateContext) interface{}

	// Visit a parse tree produced by RhumbParser#assert.
	VisitAssert(ctx *AssertContext) interface{}

	// Visit a parse tree produced by RhumbParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by RhumbParser#slurpSpread.
	VisitSlurpSpread(ctx *SlurpSpreadContext) interface{}

	// Visit a parse tree produced by RhumbParser#baseClone.
	VisitBaseClone(ctx *BaseCloneContext) interface{}

	// Visit a parse tree produced by RhumbParser#numericalPosit.
	VisitNumericalPosit(ctx *NumericalPositContext) interface{}

	// Visit a parse tree produced by RhumbParser#logicalPosit.
	VisitLogicalPosit(ctx *LogicalPositContext) interface{}

	// Visit a parse tree produced by RhumbParser#overlay.
	VisitOverlay(ctx *OverlayContext) interface{}

	// Visit a parse tree produced by RhumbParser#existentialPosit.
	VisitExistentialPosit(ctx *ExistentialPositContext) interface{}

	// Visit a parse tree produced by RhumbParser#immutableDestruct.
	VisitImmutableDestruct(ctx *ImmutableDestructContext) interface{}

	// Visit a parse tree produced by RhumbParser#mutableDestruct.
	VisitMutableDestruct(ctx *MutableDestructContext) interface{}

	// Visit a parse tree produced by RhumbParser#nestedLabel.
	VisitNestedLabel(ctx *NestedLabelContext) interface{}

	// Visit a parse tree produced by RhumbParser#nestedOverlay.
	VisitNestedOverlay(ctx *NestedOverlayContext) interface{}

	// Visit a parse tree produced by RhumbParser#range.
	VisitRange(ctx *RangeContext) interface{}

	// Visit a parse tree produced by RhumbParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by RhumbParser#LabelInterp.
	VisitLabelInterp(ctx *LabelInterpContext) interface{}

	// Visit a parse tree produced by RhumbParser#routineInterp.
	VisitRoutineInterp(ctx *RoutineInterpContext) interface{}

	// Visit a parse tree produced by RhumbParser#selectorInterp.
	VisitSelectorInterp(ctx *SelectorInterpContext) interface{}

	// Visit a parse tree produced by RhumbParser#namedRef.
	VisitNamedRef(ctx *NamedRefContext) interface{}

	// Visit a parse tree produced by RhumbParser#functionRef.
	VisitFunctionRef(ctx *FunctionRefContext) interface{}

	// Visit a parse tree produced by RhumbParser#computedRef.
	VisitComputedRef(ctx *ComputedRefContext) interface{}

	// Visit a parse tree produced by RhumbParser#junctionRef.
	VisitJunctionRef(ctx *JunctionRefContext) interface{}
}
