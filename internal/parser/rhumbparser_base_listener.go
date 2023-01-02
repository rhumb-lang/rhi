// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // RhumbParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseRhumbParserListener is a complete listener for a parse tree produced by RhumbParser.
type BaseRhumbParserListener struct{}

var _ RhumbParserListener = &BaseRhumbParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRhumbParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRhumbParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRhumbParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRhumbParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSequence is called when production sequence is entered.
func (s *BaseRhumbParserListener) EnterSequence(ctx *SequenceContext) {}

// ExitSequence is called when production sequence is exited.
func (s *BaseRhumbParserListener) ExitSequence(ctx *SequenceContext) {}

// EnterTerminator is called when production terminator is entered.
func (s *BaseRhumbParserListener) EnterTerminator(ctx *TerminatorContext) {}

// ExitTerminator is called when production terminator is exited.
func (s *BaseRhumbParserListener) ExitTerminator(ctx *TerminatorContext) {}

// EnterFloatLiteral is called when production FloatLiteral is entered.
func (s *BaseRhumbParserListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production FloatLiteral is exited.
func (s *BaseRhumbParserListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterIntegerLiteral is called when production IntegerLiteral is entered.
func (s *BaseRhumbParserListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production IntegerLiteral is exited.
func (s *BaseRhumbParserListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterStringLiteral is called when production StringLiteral is entered.
func (s *BaseRhumbParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production StringLiteral is exited.
func (s *BaseRhumbParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterReferenceLiteral is called when production ReferenceLiteral is entered.
func (s *BaseRhumbParserListener) EnterReferenceLiteral(ctx *ReferenceLiteralContext) {}

// ExitReferenceLiteral is called when production ReferenceLiteral is exited.
func (s *BaseRhumbParserListener) ExitReferenceLiteral(ctx *ReferenceLiteralContext) {}

// EnterLabelLiteral is called when production LabelLiteral is entered.
func (s *BaseRhumbParserListener) EnterLabelLiteral(ctx *LabelLiteralContext) {}

// ExitLabelLiteral is called when production LabelLiteral is exited.
func (s *BaseRhumbParserListener) ExitLabelLiteral(ctx *LabelLiteralContext) {}

// EnterConjunctive is called when production conjunctive is entered.
func (s *BaseRhumbParserListener) EnterConjunctive(ctx *ConjunctiveContext) {}

// ExitConjunctive is called when production conjunctive is exited.
func (s *BaseRhumbParserListener) ExitConjunctive(ctx *ConjunctiveContext) {}

// EnterAccess is called when production access is entered.
func (s *BaseRhumbParserListener) EnterAccess(ctx *AccessContext) {}

// ExitAccess is called when production access is exited.
func (s *BaseRhumbParserListener) ExitAccess(ctx *AccessContext) {}

// EnterApplicative is called when production applicative is entered.
func (s *BaseRhumbParserListener) EnterApplicative(ctx *ApplicativeContext) {}

// ExitApplicative is called when production applicative is exited.
func (s *BaseRhumbParserListener) ExitApplicative(ctx *ApplicativeContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BaseRhumbParserListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BaseRhumbParserListener) ExitConditional(ctx *ConditionalContext) {}

// EnterPrefix is called when production prefix is entered.
func (s *BaseRhumbParserListener) EnterPrefix(ctx *PrefixContext) {}

// ExitPrefix is called when production prefix is exited.
func (s *BaseRhumbParserListener) ExitPrefix(ctx *PrefixContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseRhumbParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseRhumbParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterComparative is called when production comparative is entered.
func (s *BaseRhumbParserListener) EnterComparative(ctx *ComparativeContext) {}

// ExitComparative is called when production comparative is exited.
func (s *BaseRhumbParserListener) ExitComparative(ctx *ComparativeContext) {}

// EnterSimple is called when production simple is entered.
func (s *BaseRhumbParserListener) EnterSimple(ctx *SimpleContext) {}

// ExitSimple is called when production simple is exited.
func (s *BaseRhumbParserListener) ExitSimple(ctx *SimpleContext) {}

// EnterMultiplicative is called when production multiplicative is entered.
func (s *BaseRhumbParserListener) EnterMultiplicative(ctx *MultiplicativeContext) {}

// ExitMultiplicative is called when production multiplicative is exited.
func (s *BaseRhumbParserListener) ExitMultiplicative(ctx *MultiplicativeContext) {}

// EnterAdditive is called when production additive is entered.
func (s *BaseRhumbParserListener) EnterAdditive(ctx *AdditiveContext) {}

// ExitAdditive is called when production additive is exited.
func (s *BaseRhumbParserListener) ExitAdditive(ctx *AdditiveContext) {}

// EnterInvocation is called when production invocation is entered.
func (s *BaseRhumbParserListener) EnterInvocation(ctx *InvocationContext) {}

// ExitInvocation is called when production invocation is exited.
func (s *BaseRhumbParserListener) ExitInvocation(ctx *InvocationContext) {}

// EnterRoutine is called when production routine is entered.
func (s *BaseRhumbParserListener) EnterRoutine(ctx *RoutineContext) {}

// ExitRoutine is called when production routine is exited.
func (s *BaseRhumbParserListener) ExitRoutine(ctx *RoutineContext) {}

// EnterDisjunctive is called when production disjunctive is entered.
func (s *BaseRhumbParserListener) EnterDisjunctive(ctx *DisjunctiveContext) {}

// ExitDisjunctive is called when production disjunctive is exited.
func (s *BaseRhumbParserListener) ExitDisjunctive(ctx *DisjunctiveContext) {}

// EnterIdentity is called when production identity is entered.
func (s *BaseRhumbParserListener) EnterIdentity(ctx *IdentityContext) {}

// ExitIdentity is called when production identity is exited.
func (s *BaseRhumbParserListener) ExitIdentity(ctx *IdentityContext) {}

// EnterEffect is called when production effect is entered.
func (s *BaseRhumbParserListener) EnterEffect(ctx *EffectContext) {}

// ExitEffect is called when production effect is exited.
func (s *BaseRhumbParserListener) ExitEffect(ctx *EffectContext) {}

// EnterMember is called when production member is entered.
func (s *BaseRhumbParserListener) EnterMember(ctx *MemberContext) {}

// ExitMember is called when production member is exited.
func (s *BaseRhumbParserListener) ExitMember(ctx *MemberContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseRhumbParserListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseRhumbParserListener) ExitSelector(ctx *SelectorContext) {}

// EnterPower is called when production power is entered.
func (s *BaseRhumbParserListener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production power is exited.
func (s *BaseRhumbParserListener) ExitPower(ctx *PowerContext) {}

// EnterMap is called when production map is entered.
func (s *BaseRhumbParserListener) EnterMap(ctx *MapContext) {}

// ExitMap is called when production map is exited.
func (s *BaseRhumbParserListener) ExitMap(ctx *MapContext) {}

// EnterFreeze is called when production freeze is entered.
func (s *BaseRhumbParserListener) EnterFreeze(ctx *FreezeContext) {}

// ExitFreeze is called when production freeze is exited.
func (s *BaseRhumbParserListener) ExitFreeze(ctx *FreezeContext) {}

// EnterInner is called when production inner is entered.
func (s *BaseRhumbParserListener) EnterInner(ctx *InnerContext) {}

// ExitInner is called when production inner is exited.
func (s *BaseRhumbParserListener) ExitInner(ctx *InnerContext) {}

// EnterLength is called when production length is entered.
func (s *BaseRhumbParserListener) EnterLength(ctx *LengthContext) {}

// ExitLength is called when production length is exited.
func (s *BaseRhumbParserListener) ExitLength(ctx *LengthContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseRhumbParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseRhumbParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterJunction is called when production junction is entered.
func (s *BaseRhumbParserListener) EnterJunction(ctx *JunctionContext) {}

// ExitJunction is called when production junction is exited.
func (s *BaseRhumbParserListener) ExitJunction(ctx *JunctionContext) {}

// EnterGreaterThan is called when production greaterThan is entered.
func (s *BaseRhumbParserListener) EnterGreaterThan(ctx *GreaterThanContext) {}

// ExitGreaterThan is called when production greaterThan is exited.
func (s *BaseRhumbParserListener) ExitGreaterThan(ctx *GreaterThanContext) {}

// EnterGreaterThanOrEqualTo is called when production greaterThanOrEqualTo is entered.
func (s *BaseRhumbParserListener) EnterGreaterThanOrEqualTo(ctx *GreaterThanOrEqualToContext) {}

// ExitGreaterThanOrEqualTo is called when production greaterThanOrEqualTo is exited.
func (s *BaseRhumbParserListener) ExitGreaterThanOrEqualTo(ctx *GreaterThanOrEqualToContext) {}

// EnterLessThan is called when production lessThan is entered.
func (s *BaseRhumbParserListener) EnterLessThan(ctx *LessThanContext) {}

// ExitLessThan is called when production lessThan is exited.
func (s *BaseRhumbParserListener) ExitLessThan(ctx *LessThanContext) {}

// EnterLessThanOrEqualTo is called when production lessThanOrEqualTo is entered.
func (s *BaseRhumbParserListener) EnterLessThanOrEqualTo(ctx *LessThanOrEqualToContext) {}

// ExitLessThanOrEqualTo is called when production lessThanOrEqualTo is exited.
func (s *BaseRhumbParserListener) ExitLessThanOrEqualTo(ctx *LessThanOrEqualToContext) {}

// EnterIsEqual is called when production isEqual is entered.
func (s *BaseRhumbParserListener) EnterIsEqual(ctx *IsEqualContext) {}

// ExitIsEqual is called when production isEqual is exited.
func (s *BaseRhumbParserListener) ExitIsEqual(ctx *IsEqualContext) {}

// EnterIsLike is called when production isLike is entered.
func (s *BaseRhumbParserListener) EnterIsLike(ctx *IsLikeContext) {}

// ExitIsLike is called when production isLike is exited.
func (s *BaseRhumbParserListener) ExitIsLike(ctx *IsLikeContext) {}

// EnterIsIn is called when production isIn is entered.
func (s *BaseRhumbParserListener) EnterIsIn(ctx *IsInContext) {}

// ExitIsIn is called when production isIn is exited.
func (s *BaseRhumbParserListener) ExitIsIn(ctx *IsInContext) {}

// EnterIsOverlayed is called when production isOverlayed is entered.
func (s *BaseRhumbParserListener) EnterIsOverlayed(ctx *IsOverlayedContext) {}

// ExitIsOverlayed is called when production isOverlayed is exited.
func (s *BaseRhumbParserListener) ExitIsOverlayed(ctx *IsOverlayedContext) {}

// EnterIsTopmost is called when production isTopmost is entered.
func (s *BaseRhumbParserListener) EnterIsTopmost(ctx *IsTopmostContext) {}

// ExitIsTopmost is called when production isTopmost is exited.
func (s *BaseRhumbParserListener) ExitIsTopmost(ctx *IsTopmostContext) {}

// EnterNotEqual is called when production notEqual is entered.
func (s *BaseRhumbParserListener) EnterNotEqual(ctx *NotEqualContext) {}

// ExitNotEqual is called when production notEqual is exited.
func (s *BaseRhumbParserListener) ExitNotEqual(ctx *NotEqualContext) {}

// EnterNotLike is called when production notLike is entered.
func (s *BaseRhumbParserListener) EnterNotLike(ctx *NotLikeContext) {}

// ExitNotLike is called when production notLike is exited.
func (s *BaseRhumbParserListener) ExitNotLike(ctx *NotLikeContext) {}

// EnterNotIn is called when production notIn is entered.
func (s *BaseRhumbParserListener) EnterNotIn(ctx *NotInContext) {}

// ExitNotIn is called when production notIn is exited.
func (s *BaseRhumbParserListener) ExitNotIn(ctx *NotInContext) {}

// EnterNotOverlayed is called when production notOverlayed is entered.
func (s *BaseRhumbParserListener) EnterNotOverlayed(ctx *NotOverlayedContext) {}

// ExitNotOverlayed is called when production notOverlayed is exited.
func (s *BaseRhumbParserListener) ExitNotOverlayed(ctx *NotOverlayedContext) {}

// EnterNotTopmost is called when production notTopmost is entered.
func (s *BaseRhumbParserListener) EnterNotTopmost(ctx *NotTopmostContext) {}

// ExitNotTopmost is called when production notTopmost is exited.
func (s *BaseRhumbParserListener) ExitNotTopmost(ctx *NotTopmostContext) {}

// EnterConjunctiveOp is called when production conjunctiveOp is entered.
func (s *BaseRhumbParserListener) EnterConjunctiveOp(ctx *ConjunctiveOpContext) {}

// ExitConjunctiveOp is called when production conjunctiveOp is exited.
func (s *BaseRhumbParserListener) ExitConjunctiveOp(ctx *ConjunctiveOpContext) {}

// EnterDisjunctiveOp is called when production disjunctiveOp is entered.
func (s *BaseRhumbParserListener) EnterDisjunctiveOp(ctx *DisjunctiveOpContext) {}

// ExitDisjunctiveOp is called when production disjunctiveOp is exited.
func (s *BaseRhumbParserListener) ExitDisjunctiveOp(ctx *DisjunctiveOpContext) {}

// EnterOtherwise is called when production otherwise is entered.
func (s *BaseRhumbParserListener) EnterOtherwise(ctx *OtherwiseContext) {}

// ExitOtherwise is called when production otherwise is exited.
func (s *BaseRhumbParserListener) ExitOtherwise(ctx *OtherwiseContext) {}

// EnterDefault is called when production default is entered.
func (s *BaseRhumbParserListener) EnterDefault(ctx *DefaultContext) {}

// ExitDefault is called when production default is exited.
func (s *BaseRhumbParserListener) ExitDefault(ctx *DefaultContext) {}

// EnterForeach is called when production foreach is entered.
func (s *BaseRhumbParserListener) EnterForeach(ctx *ForeachContext) {}

// ExitForeach is called when production foreach is exited.
func (s *BaseRhumbParserListener) ExitForeach(ctx *ForeachContext) {}

// EnterWhile is called when production while is entered.
func (s *BaseRhumbParserListener) EnterWhile(ctx *WhileContext) {}

// ExitWhile is called when production while is exited.
func (s *BaseRhumbParserListener) ExitWhile(ctx *WhileContext) {}

// EnterThen is called when production then is entered.
func (s *BaseRhumbParserListener) EnterThen(ctx *ThenContext) {}

// ExitThen is called when production then is exited.
func (s *BaseRhumbParserListener) ExitThen(ctx *ThenContext) {}

// EnterElse is called when production else is entered.
func (s *BaseRhumbParserListener) EnterElse(ctx *ElseContext) {}

// ExitElse is called when production else is exited.
func (s *BaseRhumbParserListener) ExitElse(ctx *ElseContext) {}

// EnterAddition is called when production addition is entered.
func (s *BaseRhumbParserListener) EnterAddition(ctx *AdditionContext) {}

// ExitAddition is called when production addition is exited.
func (s *BaseRhumbParserListener) ExitAddition(ctx *AdditionContext) {}

// EnterDeviation is called when production deviation is entered.
func (s *BaseRhumbParserListener) EnterDeviation(ctx *DeviationContext) {}

// ExitDeviation is called when production deviation is exited.
func (s *BaseRhumbParserListener) ExitDeviation(ctx *DeviationContext) {}

// EnterSubtraction is called when production subtraction is entered.
func (s *BaseRhumbParserListener) EnterSubtraction(ctx *SubtractionContext) {}

// ExitSubtraction is called when production subtraction is exited.
func (s *BaseRhumbParserListener) ExitSubtraction(ctx *SubtractionContext) {}

// EnterMultiplication is called when production multiplication is entered.
func (s *BaseRhumbParserListener) EnterMultiplication(ctx *MultiplicationContext) {}

// ExitMultiplication is called when production multiplication is exited.
func (s *BaseRhumbParserListener) ExitMultiplication(ctx *MultiplicationContext) {}

// EnterDivision is called when production division is entered.
func (s *BaseRhumbParserListener) EnterDivision(ctx *DivisionContext) {}

// ExitDivision is called when production division is exited.
func (s *BaseRhumbParserListener) ExitDivision(ctx *DivisionContext) {}

// EnterIntegerDivision is called when production integerDivision is entered.
func (s *BaseRhumbParserListener) EnterIntegerDivision(ctx *IntegerDivisionContext) {}

// ExitIntegerDivision is called when production integerDivision is exited.
func (s *BaseRhumbParserListener) ExitIntegerDivision(ctx *IntegerDivisionContext) {}

// EnterModulo is called when production modulo is entered.
func (s *BaseRhumbParserListener) EnterModulo(ctx *ModuloContext) {}

// ExitModulo is called when production modulo is exited.
func (s *BaseRhumbParserListener) ExitModulo(ctx *ModuloContext) {}

// EnterBind is called when production bind is entered.
func (s *BaseRhumbParserListener) EnterBind(ctx *BindContext) {}

// ExitBind is called when production bind is exited.
func (s *BaseRhumbParserListener) ExitBind(ctx *BindContext) {}

// EnterExponent is called when production exponent is entered.
func (s *BaseRhumbParserListener) EnterExponent(ctx *ExponentContext) {}

// ExitExponent is called when production exponent is exited.
func (s *BaseRhumbParserListener) ExitExponent(ctx *ExponentContext) {}

// EnterRootExtraction is called when production rootExtraction is entered.
func (s *BaseRhumbParserListener) EnterRootExtraction(ctx *RootExtractionContext) {}

// ExitRootExtraction is called when production rootExtraction is exited.
func (s *BaseRhumbParserListener) ExitRootExtraction(ctx *RootExtractionContext) {}

// EnterScientific is called when production scientific is entered.
func (s *BaseRhumbParserListener) EnterScientific(ctx *ScientificContext) {}

// ExitScientific is called when production scientific is exited.
func (s *BaseRhumbParserListener) ExitScientific(ctx *ScientificContext) {}

// EnterImmutablePair is called when production immutablePair is entered.
func (s *BaseRhumbParserListener) EnterImmutablePair(ctx *ImmutablePairContext) {}

// ExitImmutablePair is called when production immutablePair is exited.
func (s *BaseRhumbParserListener) ExitImmutablePair(ctx *ImmutablePairContext) {}

// EnterImmutableLabel is called when production immutableLabel is entered.
func (s *BaseRhumbParserListener) EnterImmutableLabel(ctx *ImmutableLabelContext) {}

// ExitImmutableLabel is called when production immutableLabel is exited.
func (s *BaseRhumbParserListener) ExitImmutableLabel(ctx *ImmutableLabelContext) {}

// EnterMutablePair is called when production mutablePair is entered.
func (s *BaseRhumbParserListener) EnterMutablePair(ctx *MutablePairContext) {}

// ExitMutablePair is called when production mutablePair is exited.
func (s *BaseRhumbParserListener) ExitMutablePair(ctx *MutablePairContext) {}

// EnterMutableLabel is called when production mutableLabel is entered.
func (s *BaseRhumbParserListener) EnterMutableLabel(ctx *MutableLabelContext) {}

// ExitMutableLabel is called when production mutableLabel is exited.
func (s *BaseRhumbParserListener) ExitMutableLabel(ctx *MutableLabelContext) {}

// EnterNumericalNegate is called when production numericalNegate is entered.
func (s *BaseRhumbParserListener) EnterNumericalNegate(ctx *NumericalNegateContext) {}

// ExitNumericalNegate is called when production numericalNegate is exited.
func (s *BaseRhumbParserListener) ExitNumericalNegate(ctx *NumericalNegateContext) {}

// EnterOuterScope is called when production outerScope is entered.
func (s *BaseRhumbParserListener) EnterOuterScope(ctx *OuterScopeContext) {}

// ExitOuterScope is called when production outerScope is exited.
func (s *BaseRhumbParserListener) ExitOuterScope(ctx *OuterScopeContext) {}

// EnterLogicalNegate is called when production logicalNegate is entered.
func (s *BaseRhumbParserListener) EnterLogicalNegate(ctx *LogicalNegateContext) {}

// ExitLogicalNegate is called when production logicalNegate is exited.
func (s *BaseRhumbParserListener) ExitLogicalNegate(ctx *LogicalNegateContext) {}

// EnterAssert is called when production assert is entered.
func (s *BaseRhumbParserListener) EnterAssert(ctx *AssertContext) {}

// ExitAssert is called when production assert is exited.
func (s *BaseRhumbParserListener) ExitAssert(ctx *AssertContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseRhumbParserListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseRhumbParserListener) ExitArgument(ctx *ArgumentContext) {}

// EnterSlurpSpread is called when production slurpSpread is entered.
func (s *BaseRhumbParserListener) EnterSlurpSpread(ctx *SlurpSpreadContext) {}

// ExitSlurpSpread is called when production slurpSpread is exited.
func (s *BaseRhumbParserListener) ExitSlurpSpread(ctx *SlurpSpreadContext) {}

// EnterBaseClone is called when production baseClone is entered.
func (s *BaseRhumbParserListener) EnterBaseClone(ctx *BaseCloneContext) {}

// ExitBaseClone is called when production baseClone is exited.
func (s *BaseRhumbParserListener) ExitBaseClone(ctx *BaseCloneContext) {}

// EnterNumericalPosit is called when production numericalPosit is entered.
func (s *BaseRhumbParserListener) EnterNumericalPosit(ctx *NumericalPositContext) {}

// ExitNumericalPosit is called when production numericalPosit is exited.
func (s *BaseRhumbParserListener) ExitNumericalPosit(ctx *NumericalPositContext) {}

// EnterLogicalPosit is called when production logicalPosit is entered.
func (s *BaseRhumbParserListener) EnterLogicalPosit(ctx *LogicalPositContext) {}

// ExitLogicalPosit is called when production logicalPosit is exited.
func (s *BaseRhumbParserListener) ExitLogicalPosit(ctx *LogicalPositContext) {}

// EnterOverlay is called when production overlay is entered.
func (s *BaseRhumbParserListener) EnterOverlay(ctx *OverlayContext) {}

// ExitOverlay is called when production overlay is exited.
func (s *BaseRhumbParserListener) ExitOverlay(ctx *OverlayContext) {}

// EnterExistentialPosit is called when production existentialPosit is entered.
func (s *BaseRhumbParserListener) EnterExistentialPosit(ctx *ExistentialPositContext) {}

// ExitExistentialPosit is called when production existentialPosit is exited.
func (s *BaseRhumbParserListener) ExitExistentialPosit(ctx *ExistentialPositContext) {}

// EnterImmutableDestruct is called when production immutableDestruct is entered.
func (s *BaseRhumbParserListener) EnterImmutableDestruct(ctx *ImmutableDestructContext) {}

// ExitImmutableDestruct is called when production immutableDestruct is exited.
func (s *BaseRhumbParserListener) ExitImmutableDestruct(ctx *ImmutableDestructContext) {}

// EnterMutableDestruct is called when production mutableDestruct is entered.
func (s *BaseRhumbParserListener) EnterMutableDestruct(ctx *MutableDestructContext) {}

// ExitMutableDestruct is called when production mutableDestruct is exited.
func (s *BaseRhumbParserListener) ExitMutableDestruct(ctx *MutableDestructContext) {}

// EnterNestedLabel is called when production nestedLabel is entered.
func (s *BaseRhumbParserListener) EnterNestedLabel(ctx *NestedLabelContext) {}

// ExitNestedLabel is called when production nestedLabel is exited.
func (s *BaseRhumbParserListener) ExitNestedLabel(ctx *NestedLabelContext) {}

// EnterNestedOverlay is called when production nestedOverlay is entered.
func (s *BaseRhumbParserListener) EnterNestedOverlay(ctx *NestedOverlayContext) {}

// ExitNestedOverlay is called when production nestedOverlay is exited.
func (s *BaseRhumbParserListener) ExitNestedOverlay(ctx *NestedOverlayContext) {}

// EnterRange is called when production range is entered.
func (s *BaseRhumbParserListener) EnterRange(ctx *RangeContext) {}

// ExitRange is called when production range is exited.
func (s *BaseRhumbParserListener) ExitRange(ctx *RangeContext) {}

// EnterString is called when production string is entered.
func (s *BaseRhumbParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseRhumbParserListener) ExitString(ctx *StringContext) {}

// EnterLabelInterp is called when production LabelInterp is entered.
func (s *BaseRhumbParserListener) EnterLabelInterp(ctx *LabelInterpContext) {}

// ExitLabelInterp is called when production LabelInterp is exited.
func (s *BaseRhumbParserListener) ExitLabelInterp(ctx *LabelInterpContext) {}

// EnterRoutineInterp is called when production routineInterp is entered.
func (s *BaseRhumbParserListener) EnterRoutineInterp(ctx *RoutineInterpContext) {}

// ExitRoutineInterp is called when production routineInterp is exited.
func (s *BaseRhumbParserListener) ExitRoutineInterp(ctx *RoutineInterpContext) {}

// EnterSelectorInterp is called when production selectorInterp is entered.
func (s *BaseRhumbParserListener) EnterSelectorInterp(ctx *SelectorInterpContext) {}

// ExitSelectorInterp is called when production selectorInterp is exited.
func (s *BaseRhumbParserListener) ExitSelectorInterp(ctx *SelectorInterpContext) {}

// EnterNamedRef is called when production namedRef is entered.
func (s *BaseRhumbParserListener) EnterNamedRef(ctx *NamedRefContext) {}

// ExitNamedRef is called when production namedRef is exited.
func (s *BaseRhumbParserListener) ExitNamedRef(ctx *NamedRefContext) {}

// EnterFunctionRef is called when production functionRef is entered.
func (s *BaseRhumbParserListener) EnterFunctionRef(ctx *FunctionRefContext) {}

// ExitFunctionRef is called when production functionRef is exited.
func (s *BaseRhumbParserListener) ExitFunctionRef(ctx *FunctionRefContext) {}

// EnterComputedRef is called when production computedRef is entered.
func (s *BaseRhumbParserListener) EnterComputedRef(ctx *ComputedRefContext) {}

// ExitComputedRef is called when production computedRef is exited.
func (s *BaseRhumbParserListener) ExitComputedRef(ctx *ComputedRefContext) {}

// EnterJunctionRef is called when production junctionRef is entered.
func (s *BaseRhumbParserListener) EnterJunctionRef(ctx *JunctionRefContext) {}

// ExitJunctionRef is called when production junctionRef is exited.
func (s *BaseRhumbParserListener) ExitJunctionRef(ctx *JunctionRefContext) {}
