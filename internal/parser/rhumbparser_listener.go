// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // RhumbParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// RhumbParserListener is a complete listener for a parse tree produced by RhumbParser.
type RhumbParserListener interface {
	antlr.ParseTreeListener

	// EnterSequence is called when entering the sequence production.
	EnterSequence(c *SequenceContext)

	// EnterTerminator is called when entering the terminator production.
	EnterTerminator(c *TerminatorContext)

	// EnterFloatLiteral is called when entering the FloatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterIntegerLiteral is called when entering the IntegerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterStringLiteral is called when entering the StringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterReferenceLiteral is called when entering the ReferenceLiteral production.
	EnterReferenceLiteral(c *ReferenceLiteralContext)

	// EnterLabelLiteral is called when entering the LabelLiteral production.
	EnterLabelLiteral(c *LabelLiteralContext)

	// EnterConjunctive is called when entering the conjunctive production.
	EnterConjunctive(c *ConjunctiveContext)

	// EnterAccess is called when entering the access production.
	EnterAccess(c *AccessContext)

	// EnterApplicative is called when entering the applicative production.
	EnterApplicative(c *ApplicativeContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterPrefix is called when entering the prefix production.
	EnterPrefix(c *PrefixContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterComparative is called when entering the comparative production.
	EnterComparative(c *ComparativeContext)

	// EnterSimple is called when entering the simple production.
	EnterSimple(c *SimpleContext)

	// EnterMultiplicative is called when entering the multiplicative production.
	EnterMultiplicative(c *MultiplicativeContext)

	// EnterAdditive is called when entering the additive production.
	EnterAdditive(c *AdditiveContext)

	// EnterInvocation is called when entering the invocation production.
	EnterInvocation(c *InvocationContext)

	// EnterRoutine is called when entering the routine production.
	EnterRoutine(c *RoutineContext)

	// EnterDisjunctive is called when entering the disjunctive production.
	EnterDisjunctive(c *DisjunctiveContext)

	// EnterIdentity is called when entering the identity production.
	EnterIdentity(c *IdentityContext)

	// EnterEffect is called when entering the effect production.
	EnterEffect(c *EffectContext)

	// EnterMember is called when entering the member production.
	EnterMember(c *MemberContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterPower is called when entering the power production.
	EnterPower(c *PowerContext)

	// EnterMap is called when entering the map production.
	EnterMap(c *MapContext)

	// EnterFreeze is called when entering the freeze production.
	EnterFreeze(c *FreezeContext)

	// EnterInner is called when entering the inner production.
	EnterInner(c *InnerContext)

	// EnterLength is called when entering the length production.
	EnterLength(c *LengthContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterJunction is called when entering the junction production.
	EnterJunction(c *JunctionContext)

	// EnterGreaterThan is called when entering the greaterThan production.
	EnterGreaterThan(c *GreaterThanContext)

	// EnterGreaterThanOrEqualTo is called when entering the greaterThanOrEqualTo production.
	EnterGreaterThanOrEqualTo(c *GreaterThanOrEqualToContext)

	// EnterLessThan is called when entering the lessThan production.
	EnterLessThan(c *LessThanContext)

	// EnterLessThanOrEqualTo is called when entering the lessThanOrEqualTo production.
	EnterLessThanOrEqualTo(c *LessThanOrEqualToContext)

	// EnterIsEqual is called when entering the isEqual production.
	EnterIsEqual(c *IsEqualContext)

	// EnterIsLike is called when entering the isLike production.
	EnterIsLike(c *IsLikeContext)

	// EnterIsIn is called when entering the isIn production.
	EnterIsIn(c *IsInContext)

	// EnterIsOverlayed is called when entering the isOverlayed production.
	EnterIsOverlayed(c *IsOverlayedContext)

	// EnterIsTopmost is called when entering the isTopmost production.
	EnterIsTopmost(c *IsTopmostContext)

	// EnterNotEqual is called when entering the notEqual production.
	EnterNotEqual(c *NotEqualContext)

	// EnterNotLike is called when entering the notLike production.
	EnterNotLike(c *NotLikeContext)

	// EnterNotIn is called when entering the notIn production.
	EnterNotIn(c *NotInContext)

	// EnterNotOverlayed is called when entering the notOverlayed production.
	EnterNotOverlayed(c *NotOverlayedContext)

	// EnterNotTopmost is called when entering the notTopmost production.
	EnterNotTopmost(c *NotTopmostContext)

	// EnterConjunctiveOp is called when entering the conjunctiveOp production.
	EnterConjunctiveOp(c *ConjunctiveOpContext)

	// EnterDisjunctiveOp is called when entering the disjunctiveOp production.
	EnterDisjunctiveOp(c *DisjunctiveOpContext)

	// EnterOtherwise is called when entering the otherwise production.
	EnterOtherwise(c *OtherwiseContext)

	// EnterDefault is called when entering the default production.
	EnterDefault(c *DefaultContext)

	// EnterForeach is called when entering the foreach production.
	EnterForeach(c *ForeachContext)

	// EnterWhile is called when entering the while production.
	EnterWhile(c *WhileContext)

	// EnterThen is called when entering the then production.
	EnterThen(c *ThenContext)

	// EnterElse is called when entering the else production.
	EnterElse(c *ElseContext)

	// EnterAddition is called when entering the addition production.
	EnterAddition(c *AdditionContext)

	// EnterDeviation is called when entering the deviation production.
	EnterDeviation(c *DeviationContext)

	// EnterSubtraction is called when entering the subtraction production.
	EnterSubtraction(c *SubtractionContext)

	// EnterMultiplication is called when entering the multiplication production.
	EnterMultiplication(c *MultiplicationContext)

	// EnterDivision is called when entering the division production.
	EnterDivision(c *DivisionContext)

	// EnterIntegerDivision is called when entering the integerDivision production.
	EnterIntegerDivision(c *IntegerDivisionContext)

	// EnterModulo is called when entering the modulo production.
	EnterModulo(c *ModuloContext)

	// EnterBind is called when entering the bind production.
	EnterBind(c *BindContext)

	// EnterExponent is called when entering the exponent production.
	EnterExponent(c *ExponentContext)

	// EnterRootExtraction is called when entering the rootExtraction production.
	EnterRootExtraction(c *RootExtractionContext)

	// EnterScientific is called when entering the scientific production.
	EnterScientific(c *ScientificContext)

	// EnterImmutablePair is called when entering the immutablePair production.
	EnterImmutablePair(c *ImmutablePairContext)

	// EnterImmutableLabel is called when entering the immutableLabel production.
	EnterImmutableLabel(c *ImmutableLabelContext)

	// EnterMutablePair is called when entering the mutablePair production.
	EnterMutablePair(c *MutablePairContext)

	// EnterMutableLabel is called when entering the mutableLabel production.
	EnterMutableLabel(c *MutableLabelContext)

	// EnterNumericalNegate is called when entering the numericalNegate production.
	EnterNumericalNegate(c *NumericalNegateContext)

	// EnterOuterScope is called when entering the outerScope production.
	EnterOuterScope(c *OuterScopeContext)

	// EnterLogicalNegate is called when entering the logicalNegate production.
	EnterLogicalNegate(c *LogicalNegateContext)

	// EnterAssert is called when entering the assert production.
	EnterAssert(c *AssertContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterSlurpSpread is called when entering the slurpSpread production.
	EnterSlurpSpread(c *SlurpSpreadContext)

	// EnterBaseClone is called when entering the baseClone production.
	EnterBaseClone(c *BaseCloneContext)

	// EnterNumericalPosit is called when entering the numericalPosit production.
	EnterNumericalPosit(c *NumericalPositContext)

	// EnterLogicalPosit is called when entering the logicalPosit production.
	EnterLogicalPosit(c *LogicalPositContext)

	// EnterOverlay is called when entering the overlay production.
	EnterOverlay(c *OverlayContext)

	// EnterExistentialPosit is called when entering the existentialPosit production.
	EnterExistentialPosit(c *ExistentialPositContext)

	// EnterImmutableDestruct is called when entering the immutableDestruct production.
	EnterImmutableDestruct(c *ImmutableDestructContext)

	// EnterMutableDestruct is called when entering the mutableDestruct production.
	EnterMutableDestruct(c *MutableDestructContext)

	// EnterNestedLabel is called when entering the nestedLabel production.
	EnterNestedLabel(c *NestedLabelContext)

	// EnterNestedOverlay is called when entering the nestedOverlay production.
	EnterNestedOverlay(c *NestedOverlayContext)

	// EnterRange is called when entering the range production.
	EnterRange(c *RangeContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterLabelInterp is called when entering the LabelInterp production.
	EnterLabelInterp(c *LabelInterpContext)

	// EnterRoutineInterp is called when entering the routineInterp production.
	EnterRoutineInterp(c *RoutineInterpContext)

	// EnterSelectorInterp is called when entering the selectorInterp production.
	EnterSelectorInterp(c *SelectorInterpContext)

	// EnterNamedRef is called when entering the namedRef production.
	EnterNamedRef(c *NamedRefContext)

	// EnterFunctionRef is called when entering the functionRef production.
	EnterFunctionRef(c *FunctionRefContext)

	// EnterComputedRef is called when entering the computedRef production.
	EnterComputedRef(c *ComputedRefContext)

	// EnterJunctionRef is called when entering the junctionRef production.
	EnterJunctionRef(c *JunctionRefContext)

	// ExitSequence is called when exiting the sequence production.
	ExitSequence(c *SequenceContext)

	// ExitTerminator is called when exiting the terminator production.
	ExitTerminator(c *TerminatorContext)

	// ExitFloatLiteral is called when exiting the FloatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitIntegerLiteral is called when exiting the IntegerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitStringLiteral is called when exiting the StringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitReferenceLiteral is called when exiting the ReferenceLiteral production.
	ExitReferenceLiteral(c *ReferenceLiteralContext)

	// ExitLabelLiteral is called when exiting the LabelLiteral production.
	ExitLabelLiteral(c *LabelLiteralContext)

	// ExitConjunctive is called when exiting the conjunctive production.
	ExitConjunctive(c *ConjunctiveContext)

	// ExitAccess is called when exiting the access production.
	ExitAccess(c *AccessContext)

	// ExitApplicative is called when exiting the applicative production.
	ExitApplicative(c *ApplicativeContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitPrefix is called when exiting the prefix production.
	ExitPrefix(c *PrefixContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitComparative is called when exiting the comparative production.
	ExitComparative(c *ComparativeContext)

	// ExitSimple is called when exiting the simple production.
	ExitSimple(c *SimpleContext)

	// ExitMultiplicative is called when exiting the multiplicative production.
	ExitMultiplicative(c *MultiplicativeContext)

	// ExitAdditive is called when exiting the additive production.
	ExitAdditive(c *AdditiveContext)

	// ExitInvocation is called when exiting the invocation production.
	ExitInvocation(c *InvocationContext)

	// ExitRoutine is called when exiting the routine production.
	ExitRoutine(c *RoutineContext)

	// ExitDisjunctive is called when exiting the disjunctive production.
	ExitDisjunctive(c *DisjunctiveContext)

	// ExitIdentity is called when exiting the identity production.
	ExitIdentity(c *IdentityContext)

	// ExitEffect is called when exiting the effect production.
	ExitEffect(c *EffectContext)

	// ExitMember is called when exiting the member production.
	ExitMember(c *MemberContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitPower is called when exiting the power production.
	ExitPower(c *PowerContext)

	// ExitMap is called when exiting the map production.
	ExitMap(c *MapContext)

	// ExitFreeze is called when exiting the freeze production.
	ExitFreeze(c *FreezeContext)

	// ExitInner is called when exiting the inner production.
	ExitInner(c *InnerContext)

	// ExitLength is called when exiting the length production.
	ExitLength(c *LengthContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitJunction is called when exiting the junction production.
	ExitJunction(c *JunctionContext)

	// ExitGreaterThan is called when exiting the greaterThan production.
	ExitGreaterThan(c *GreaterThanContext)

	// ExitGreaterThanOrEqualTo is called when exiting the greaterThanOrEqualTo production.
	ExitGreaterThanOrEqualTo(c *GreaterThanOrEqualToContext)

	// ExitLessThan is called when exiting the lessThan production.
	ExitLessThan(c *LessThanContext)

	// ExitLessThanOrEqualTo is called when exiting the lessThanOrEqualTo production.
	ExitLessThanOrEqualTo(c *LessThanOrEqualToContext)

	// ExitIsEqual is called when exiting the isEqual production.
	ExitIsEqual(c *IsEqualContext)

	// ExitIsLike is called when exiting the isLike production.
	ExitIsLike(c *IsLikeContext)

	// ExitIsIn is called when exiting the isIn production.
	ExitIsIn(c *IsInContext)

	// ExitIsOverlayed is called when exiting the isOverlayed production.
	ExitIsOverlayed(c *IsOverlayedContext)

	// ExitIsTopmost is called when exiting the isTopmost production.
	ExitIsTopmost(c *IsTopmostContext)

	// ExitNotEqual is called when exiting the notEqual production.
	ExitNotEqual(c *NotEqualContext)

	// ExitNotLike is called when exiting the notLike production.
	ExitNotLike(c *NotLikeContext)

	// ExitNotIn is called when exiting the notIn production.
	ExitNotIn(c *NotInContext)

	// ExitNotOverlayed is called when exiting the notOverlayed production.
	ExitNotOverlayed(c *NotOverlayedContext)

	// ExitNotTopmost is called when exiting the notTopmost production.
	ExitNotTopmost(c *NotTopmostContext)

	// ExitConjunctiveOp is called when exiting the conjunctiveOp production.
	ExitConjunctiveOp(c *ConjunctiveOpContext)

	// ExitDisjunctiveOp is called when exiting the disjunctiveOp production.
	ExitDisjunctiveOp(c *DisjunctiveOpContext)

	// ExitOtherwise is called when exiting the otherwise production.
	ExitOtherwise(c *OtherwiseContext)

	// ExitDefault is called when exiting the default production.
	ExitDefault(c *DefaultContext)

	// ExitForeach is called when exiting the foreach production.
	ExitForeach(c *ForeachContext)

	// ExitWhile is called when exiting the while production.
	ExitWhile(c *WhileContext)

	// ExitThen is called when exiting the then production.
	ExitThen(c *ThenContext)

	// ExitElse is called when exiting the else production.
	ExitElse(c *ElseContext)

	// ExitAddition is called when exiting the addition production.
	ExitAddition(c *AdditionContext)

	// ExitDeviation is called when exiting the deviation production.
	ExitDeviation(c *DeviationContext)

	// ExitSubtraction is called when exiting the subtraction production.
	ExitSubtraction(c *SubtractionContext)

	// ExitMultiplication is called when exiting the multiplication production.
	ExitMultiplication(c *MultiplicationContext)

	// ExitDivision is called when exiting the division production.
	ExitDivision(c *DivisionContext)

	// ExitIntegerDivision is called when exiting the integerDivision production.
	ExitIntegerDivision(c *IntegerDivisionContext)

	// ExitModulo is called when exiting the modulo production.
	ExitModulo(c *ModuloContext)

	// ExitBind is called when exiting the bind production.
	ExitBind(c *BindContext)

	// ExitExponent is called when exiting the exponent production.
	ExitExponent(c *ExponentContext)

	// ExitRootExtraction is called when exiting the rootExtraction production.
	ExitRootExtraction(c *RootExtractionContext)

	// ExitScientific is called when exiting the scientific production.
	ExitScientific(c *ScientificContext)

	// ExitImmutablePair is called when exiting the immutablePair production.
	ExitImmutablePair(c *ImmutablePairContext)

	// ExitImmutableLabel is called when exiting the immutableLabel production.
	ExitImmutableLabel(c *ImmutableLabelContext)

	// ExitMutablePair is called when exiting the mutablePair production.
	ExitMutablePair(c *MutablePairContext)

	// ExitMutableLabel is called when exiting the mutableLabel production.
	ExitMutableLabel(c *MutableLabelContext)

	// ExitNumericalNegate is called when exiting the numericalNegate production.
	ExitNumericalNegate(c *NumericalNegateContext)

	// ExitOuterScope is called when exiting the outerScope production.
	ExitOuterScope(c *OuterScopeContext)

	// ExitLogicalNegate is called when exiting the logicalNegate production.
	ExitLogicalNegate(c *LogicalNegateContext)

	// ExitAssert is called when exiting the assert production.
	ExitAssert(c *AssertContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitSlurpSpread is called when exiting the slurpSpread production.
	ExitSlurpSpread(c *SlurpSpreadContext)

	// ExitBaseClone is called when exiting the baseClone production.
	ExitBaseClone(c *BaseCloneContext)

	// ExitNumericalPosit is called when exiting the numericalPosit production.
	ExitNumericalPosit(c *NumericalPositContext)

	// ExitLogicalPosit is called when exiting the logicalPosit production.
	ExitLogicalPosit(c *LogicalPositContext)

	// ExitOverlay is called when exiting the overlay production.
	ExitOverlay(c *OverlayContext)

	// ExitExistentialPosit is called when exiting the existentialPosit production.
	ExitExistentialPosit(c *ExistentialPositContext)

	// ExitImmutableDestruct is called when exiting the immutableDestruct production.
	ExitImmutableDestruct(c *ImmutableDestructContext)

	// ExitMutableDestruct is called when exiting the mutableDestruct production.
	ExitMutableDestruct(c *MutableDestructContext)

	// ExitNestedLabel is called when exiting the nestedLabel production.
	ExitNestedLabel(c *NestedLabelContext)

	// ExitNestedOverlay is called when exiting the nestedOverlay production.
	ExitNestedOverlay(c *NestedOverlayContext)

	// ExitRange is called when exiting the range production.
	ExitRange(c *RangeContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitLabelInterp is called when exiting the LabelInterp production.
	ExitLabelInterp(c *LabelInterpContext)

	// ExitRoutineInterp is called when exiting the routineInterp production.
	ExitRoutineInterp(c *RoutineInterpContext)

	// ExitSelectorInterp is called when exiting the selectorInterp production.
	ExitSelectorInterp(c *SelectorInterpContext)

	// ExitNamedRef is called when exiting the namedRef production.
	ExitNamedRef(c *NamedRefContext)

	// ExitFunctionRef is called when exiting the functionRef production.
	ExitFunctionRef(c *FunctionRefContext)

	// ExitComputedRef is called when exiting the computedRef production.
	ExitComputedRef(c *ComputedRefContext)

	// ExitJunctionRef is called when exiting the junctionRef production.
	ExitJunctionRef(c *JunctionRefContext)
}
