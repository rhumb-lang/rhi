// Code generated from /home/jake/Code/rhumb-grammar/grammar/RhumbParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

// EnterExpressions is called when production expressions is entered.
func (s *BaseRhumbParserListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseRhumbParserListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterFields is called when production fields is entered.
func (s *BaseRhumbParserListener) EnterFields(ctx *FieldsContext) {}

// ExitFields is called when production fields is exited.
func (s *BaseRhumbParserListener) ExitFields(ctx *FieldsContext) {}

// EnterPatterns is called when production patterns is entered.
func (s *BaseRhumbParserListener) EnterPatterns(ctx *PatternsContext) {}

// ExitPatterns is called when production patterns is exited.
func (s *BaseRhumbParserListener) ExitPatterns(ctx *PatternsContext) {}

// EnterTerminator is called when production terminator is entered.
func (s *BaseRhumbParserListener) EnterTerminator(ctx *TerminatorContext) {}

// ExitTerminator is called when production terminator is exited.
func (s *BaseRhumbParserListener) ExitTerminator(ctx *TerminatorContext) {}

// EnterRationalNumber is called when production rationalNumber is entered.
func (s *BaseRhumbParserListener) EnterRationalNumber(ctx *RationalNumberContext) {}

// ExitRationalNumber is called when production rationalNumber is exited.
func (s *BaseRhumbParserListener) ExitRationalNumber(ctx *RationalNumberContext) {}

// EnterDateNumber is called when production dateNumber is entered.
func (s *BaseRhumbParserListener) EnterDateNumber(ctx *DateNumberContext) {}

// ExitDateNumber is called when production dateNumber is exited.
func (s *BaseRhumbParserListener) ExitDateNumber(ctx *DateNumberContext) {}

// EnterZeroNumber is called when production zeroNumber is entered.
func (s *BaseRhumbParserListener) EnterZeroNumber(ctx *ZeroNumberContext) {}

// ExitZeroNumber is called when production zeroNumber is exited.
func (s *BaseRhumbParserListener) ExitZeroNumber(ctx *ZeroNumberContext) {}

// EnterWholeNumber is called when production wholeNumber is entered.
func (s *BaseRhumbParserListener) EnterWholeNumber(ctx *WholeNumberContext) {}

// ExitWholeNumber is called when production wholeNumber is exited.
func (s *BaseRhumbParserListener) ExitWholeNumber(ctx *WholeNumberContext) {}

// EnterKeySymbol is called when production keySymbol is entered.
func (s *BaseRhumbParserListener) EnterKeySymbol(ctx *KeySymbolContext) {}

// ExitKeySymbol is called when production keySymbol is exited.
func (s *BaseRhumbParserListener) ExitKeySymbol(ctx *KeySymbolContext) {}

// EnterTextSymbol is called when production textSymbol is entered.
func (s *BaseRhumbParserListener) EnterTextSymbol(ctx *TextSymbolContext) {}

// ExitTextSymbol is called when production textSymbol is exited.
func (s *BaseRhumbParserListener) ExitTextSymbol(ctx *TextSymbolContext) {}

// EnterReferenceLiteral is called when production referenceLiteral is entered.
func (s *BaseRhumbParserListener) EnterReferenceLiteral(ctx *ReferenceLiteralContext) {}

// ExitReferenceLiteral is called when production referenceLiteral is exited.
func (s *BaseRhumbParserListener) ExitReferenceLiteral(ctx *ReferenceLiteralContext) {}

// EnterLabelSymbol is called when production labelSymbol is entered.
func (s *BaseRhumbParserListener) EnterLabelSymbol(ctx *LabelSymbolContext) {}

// ExitLabelSymbol is called when production labelSymbol is exited.
func (s *BaseRhumbParserListener) ExitLabelSymbol(ctx *LabelSymbolContext) {}

// EnterFieldLiteral is called when production fieldLiteral is entered.
func (s *BaseRhumbParserListener) EnterFieldLiteral(ctx *FieldLiteralContext) {}

// ExitFieldLiteral is called when production fieldLiteral is exited.
func (s *BaseRhumbParserListener) ExitFieldLiteral(ctx *FieldLiteralContext) {}

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

// EnterComparative is called when production comparative is entered.
func (s *BaseRhumbParserListener) EnterComparative(ctx *ComparativeContext) {}

// ExitComparative is called when production comparative is exited.
func (s *BaseRhumbParserListener) ExitComparative(ctx *ComparativeContext) {}

// EnterSimpleExpression is called when production simpleExpression is entered.
func (s *BaseRhumbParserListener) EnterSimpleExpression(ctx *SimpleExpressionContext) {}

// ExitSimpleExpression is called when production simpleExpression is exited.
func (s *BaseRhumbParserListener) ExitSimpleExpression(ctx *SimpleExpressionContext) {}

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

// EnterLibrary is called when production library is entered.
func (s *BaseRhumbParserListener) EnterLibrary(ctx *LibraryContext) {}

// ExitLibrary is called when production library is exited.
func (s *BaseRhumbParserListener) ExitLibrary(ctx *LibraryContext) {}

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

// EnterAssignLabel is called when production assignLabel is entered.
func (s *BaseRhumbParserListener) EnterAssignLabel(ctx *AssignLabelContext) {}

// ExitAssignLabel is called when production assignLabel is exited.
func (s *BaseRhumbParserListener) ExitAssignLabel(ctx *AssignLabelContext) {}

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

// EnterChainExpression is called when production chainExpression is entered.
func (s *BaseRhumbParserListener) EnterChainExpression(ctx *ChainExpressionContext) {}

// ExitChainExpression is called when production chainExpression is exited.
func (s *BaseRhumbParserListener) ExitChainExpression(ctx *ChainExpressionContext) {}

// EnterPrefixAssignMutField is called when production prefixAssignMutField is entered.
func (s *BaseRhumbParserListener) EnterPrefixAssignMutField(ctx *PrefixAssignMutFieldContext) {}

// ExitPrefixAssignMutField is called when production prefixAssignMutField is exited.
func (s *BaseRhumbParserListener) ExitPrefixAssignMutField(ctx *PrefixAssignMutFieldContext) {}

// EnterPrefixAssignMutSubField is called when production prefixAssignMutSubField is entered.
func (s *BaseRhumbParserListener) EnterPrefixAssignMutSubField(ctx *PrefixAssignMutSubFieldContext) {}

// ExitPrefixAssignMutSubField is called when production prefixAssignMutSubField is exited.
func (s *BaseRhumbParserListener) ExitPrefixAssignMutSubField(ctx *PrefixAssignMutSubFieldContext) {}

// EnterPrefixAssignImmSubField is called when production prefixAssignImmSubField is entered.
func (s *BaseRhumbParserListener) EnterPrefixAssignImmSubField(ctx *PrefixAssignImmSubFieldContext) {}

// ExitPrefixAssignImmSubField is called when production prefixAssignImmSubField is exited.
func (s *BaseRhumbParserListener) ExitPrefixAssignImmSubField(ctx *PrefixAssignImmSubFieldContext) {}

// EnterPrefixSlurpSpread is called when production prefixSlurpSpread is entered.
func (s *BaseRhumbParserListener) EnterPrefixSlurpSpread(ctx *PrefixSlurpSpreadContext) {}

// ExitPrefixSlurpSpread is called when production prefixSlurpSpread is exited.
func (s *BaseRhumbParserListener) ExitPrefixSlurpSpread(ctx *PrefixSlurpSpreadContext) {}

// EnterAssignMutField is called when production assignMutField is entered.
func (s *BaseRhumbParserListener) EnterAssignMutField(ctx *AssignMutFieldContext) {}

// ExitAssignMutField is called when production assignMutField is exited.
func (s *BaseRhumbParserListener) ExitAssignMutField(ctx *AssignMutFieldContext) {}

// EnterAssignMutSubField is called when production assignMutSubField is entered.
func (s *BaseRhumbParserListener) EnterAssignMutSubField(ctx *AssignMutSubFieldContext) {}

// ExitAssignMutSubField is called when production assignMutSubField is exited.
func (s *BaseRhumbParserListener) ExitAssignMutSubField(ctx *AssignMutSubFieldContext) {}

// EnterAssignImmField is called when production assignImmField is entered.
func (s *BaseRhumbParserListener) EnterAssignImmField(ctx *AssignImmFieldContext) {}

// ExitAssignImmField is called when production assignImmField is exited.
func (s *BaseRhumbParserListener) ExitAssignImmField(ctx *AssignImmFieldContext) {}

// EnterAssignImmSubField is called when production assignImmSubField is entered.
func (s *BaseRhumbParserListener) EnterAssignImmSubField(ctx *AssignImmSubFieldContext) {}

// ExitAssignImmSubField is called when production assignImmSubField is exited.
func (s *BaseRhumbParserListener) ExitAssignImmSubField(ctx *AssignImmSubFieldContext) {}

// EnterSimpleMapField is called when production simpleMapField is entered.
func (s *BaseRhumbParserListener) EnterSimpleMapField(ctx *SimpleMapFieldContext) {}

// ExitSimpleMapField is called when production simpleMapField is exited.
func (s *BaseRhumbParserListener) ExitSimpleMapField(ctx *SimpleMapFieldContext) {}

// EnterSimpleField is called when production simpleField is entered.
func (s *BaseRhumbParserListener) EnterSimpleField(ctx *SimpleFieldContext) {}

// ExitSimpleField is called when production simpleField is exited.
func (s *BaseRhumbParserListener) ExitSimpleField(ctx *SimpleFieldContext) {}

// EnterAssignBreakingPattern is called when production assignBreakingPattern is entered.
func (s *BaseRhumbParserListener) EnterAssignBreakingPattern(ctx *AssignBreakingPatternContext) {}

// ExitAssignBreakingPattern is called when production assignBreakingPattern is exited.
func (s *BaseRhumbParserListener) ExitAssignBreakingPattern(ctx *AssignBreakingPatternContext) {}

// EnterAssignFallthroughPattern is called when production assignFallthroughPattern is entered.
func (s *BaseRhumbParserListener) EnterAssignFallthroughPattern(ctx *AssignFallthroughPatternContext) {
}

// ExitAssignFallthroughPattern is called when production assignFallthroughPattern is exited.
func (s *BaseRhumbParserListener) ExitAssignFallthroughPattern(ctx *AssignFallthroughPatternContext) {
}

// EnterAssignDefaultPattern is called when production assignDefaultPattern is entered.
func (s *BaseRhumbParserListener) EnterAssignDefaultPattern(ctx *AssignDefaultPatternContext) {}

// ExitAssignDefaultPattern is called when production assignDefaultPattern is exited.
func (s *BaseRhumbParserListener) ExitAssignDefaultPattern(ctx *AssignDefaultPatternContext) {}

// EnterAppend is called when production append is entered.
func (s *BaseRhumbParserListener) EnterAppend(ctx *AppendContext) {}

// ExitAppend is called when production append is exited.
func (s *BaseRhumbParserListener) ExitAppend(ctx *AppendContext) {}

// EnterUnshift is called when production unshift is entered.
func (s *BaseRhumbParserListener) EnterUnshift(ctx *UnshiftContext) {}

// ExitUnshift is called when production unshift is exited.
func (s *BaseRhumbParserListener) ExitUnshift(ctx *UnshiftContext) {}

// EnterLength is called when production length is entered.
func (s *BaseRhumbParserListener) EnterLength(ctx *LengthContext) {}

// ExitLength is called when production length is exited.
func (s *BaseRhumbParserListener) ExitLength(ctx *LengthContext) {}

// EnterEmpty is called when production empty is entered.
func (s *BaseRhumbParserListener) EnterEmpty(ctx *EmptyContext) {}

// ExitEmpty is called when production empty is exited.
func (s *BaseRhumbParserListener) ExitEmpty(ctx *EmptyContext) {}

// EnterAllSubfields is called when production allSubfields is entered.
func (s *BaseRhumbParserListener) EnterAllSubfields(ctx *AllSubfieldsContext) {}

// ExitAllSubfields is called when production allSubfields is exited.
func (s *BaseRhumbParserListener) ExitAllSubfields(ctx *AllSubfieldsContext) {}

// EnterAllFields is called when production allFields is entered.
func (s *BaseRhumbParserListener) EnterAllFields(ctx *AllFieldsContext) {}

// ExitAllFields is called when production allFields is exited.
func (s *BaseRhumbParserListener) ExitAllFields(ctx *AllFieldsContext) {}

// EnterElements is called when production elements is entered.
func (s *BaseRhumbParserListener) EnterElements(ctx *ElementsContext) {}

// ExitElements is called when production elements is exited.
func (s *BaseRhumbParserListener) ExitElements(ctx *ElementsContext) {}

// EnterFreeze is called when production freeze is entered.
func (s *BaseRhumbParserListener) EnterFreeze(ctx *FreezeContext) {}

// ExitFreeze is called when production freeze is exited.
func (s *BaseRhumbParserListener) ExitFreeze(ctx *FreezeContext) {}

// EnterCopy is called when production copy is entered.
func (s *BaseRhumbParserListener) EnterCopy(ctx *CopyContext) {}

// ExitCopy is called when production copy is exited.
func (s *BaseRhumbParserListener) ExitCopy(ctx *CopyContext) {}

// EnterToDate is called when production toDate is entered.
func (s *BaseRhumbParserListener) EnterToDate(ctx *ToDateContext) {}

// ExitToDate is called when production toDate is exited.
func (s *BaseRhumbParserListener) ExitToDate(ctx *ToDateContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseRhumbParserListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseRhumbParserListener) ExitParameters(ctx *ParametersContext) {}

// EnterConstructor is called when production constructor is entered.
func (s *BaseRhumbParserListener) EnterConstructor(ctx *ConstructorContext) {}

// ExitConstructor is called when production constructor is exited.
func (s *BaseRhumbParserListener) ExitConstructor(ctx *ConstructorContext) {}

// EnterBase is called when production base is entered.
func (s *BaseRhumbParserListener) EnterBase(ctx *BaseContext) {}

// ExitBase is called when production base is exited.
func (s *BaseRhumbParserListener) ExitBase(ctx *BaseContext) {}

// EnterToNumber is called when production toNumber is entered.
func (s *BaseRhumbParserListener) EnterToNumber(ctx *ToNumberContext) {}

// ExitToNumber is called when production toNumber is exited.
func (s *BaseRhumbParserListener) ExitToNumber(ctx *ToNumberContext) {}

// EnterNegateNumber is called when production negateNumber is entered.
func (s *BaseRhumbParserListener) EnterNegateNumber(ctx *NegateNumberContext) {}

// ExitNegateNumber is called when production negateNumber is exited.
func (s *BaseRhumbParserListener) ExitNegateNumber(ctx *NegateNumberContext) {}

// EnterToTruth is called when production toTruth is entered.
func (s *BaseRhumbParserListener) EnterToTruth(ctx *ToTruthContext) {}

// ExitToTruth is called when production toTruth is exited.
func (s *BaseRhumbParserListener) ExitToTruth(ctx *ToTruthContext) {}

// EnterNegateTruth is called when production negateTruth is entered.
func (s *BaseRhumbParserListener) EnterNegateTruth(ctx *NegateTruthContext) {}

// ExitNegateTruth is called when production negateTruth is exited.
func (s *BaseRhumbParserListener) ExitNegateTruth(ctx *NegateTruthContext) {}

// EnterVariadic is called when production variadic is entered.
func (s *BaseRhumbParserListener) EnterVariadic(ctx *VariadicContext) {}

// ExitVariadic is called when production variadic is exited.
func (s *BaseRhumbParserListener) ExitVariadic(ctx *VariadicContext) {}

// EnterToKey is called when production toKey is entered.
func (s *BaseRhumbParserListener) EnterToKey(ctx *ToKeyContext) {}

// ExitToKey is called when production toKey is exited.
func (s *BaseRhumbParserListener) ExitToKey(ctx *ToKeyContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseRhumbParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseRhumbParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseRhumbParserListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseRhumbParserListener) ExitMethod(ctx *MethodContext) {}

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

// EnterIsInner is called when production isInner is entered.
func (s *BaseRhumbParserListener) EnterIsInner(ctx *IsInnerContext) {}

// ExitIsInner is called when production isInner is exited.
func (s *BaseRhumbParserListener) ExitIsInner(ctx *IsInnerContext) {}

// EnterIsUnder is called when production isUnder is entered.
func (s *BaseRhumbParserListener) EnterIsUnder(ctx *IsUnderContext) {}

// ExitIsUnder is called when production isUnder is exited.
func (s *BaseRhumbParserListener) ExitIsUnder(ctx *IsUnderContext) {}

// EnterNotEqual is called when production notEqual is entered.
func (s *BaseRhumbParserListener) EnterNotEqual(ctx *NotEqualContext) {}

// ExitNotEqual is called when production notEqual is exited.
func (s *BaseRhumbParserListener) ExitNotEqual(ctx *NotEqualContext) {}

// EnterNotInner is called when production notInner is entered.
func (s *BaseRhumbParserListener) EnterNotInner(ctx *NotInnerContext) {}

// ExitNotInner is called when production notInner is exited.
func (s *BaseRhumbParserListener) ExitNotInner(ctx *NotInnerContext) {}

// EnterNotUnder is called when production notUnder is entered.
func (s *BaseRhumbParserListener) EnterNotUnder(ctx *NotUnderContext) {}

// ExitNotUnder is called when production notUnder is exited.
func (s *BaseRhumbParserListener) ExitNotUnder(ctx *NotUnderContext) {}

// EnterConjunctiveOp is called when production conjunctiveOp is entered.
func (s *BaseRhumbParserListener) EnterConjunctiveOp(ctx *ConjunctiveOpContext) {}

// ExitConjunctiveOp is called when production conjunctiveOp is exited.
func (s *BaseRhumbParserListener) ExitConjunctiveOp(ctx *ConjunctiveOpContext) {}

// EnterDisjunctiveOp is called when production disjunctiveOp is entered.
func (s *BaseRhumbParserListener) EnterDisjunctiveOp(ctx *DisjunctiveOpContext) {}

// ExitDisjunctiveOp is called when production disjunctiveOp is exited.
func (s *BaseRhumbParserListener) ExitDisjunctiveOp(ctx *DisjunctiveOpContext) {}

// EnterPipe is called when production pipe is entered.
func (s *BaseRhumbParserListener) EnterPipe(ctx *PipeContext) {}

// ExitPipe is called when production pipe is exited.
func (s *BaseRhumbParserListener) ExitPipe(ctx *PipeContext) {}

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

// EnterDestructure is called when production destructure is entered.
func (s *BaseRhumbParserListener) EnterDestructure(ctx *DestructureContext) {}

// ExitDestructure is called when production destructure is exited.
func (s *BaseRhumbParserListener) ExitDestructure(ctx *DestructureContext) {}

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

// EnterConcatenate is called when production concatenate is entered.
func (s *BaseRhumbParserListener) EnterConcatenate(ctx *ConcatenateContext) {}

// ExitConcatenate is called when production concatenate is exited.
func (s *BaseRhumbParserListener) ExitConcatenate(ctx *ConcatenateContext) {}

// EnterMultiplication is called when production multiplication is entered.
func (s *BaseRhumbParserListener) EnterMultiplication(ctx *MultiplicationContext) {}

// ExitMultiplication is called when production multiplication is exited.
func (s *BaseRhumbParserListener) ExitMultiplication(ctx *MultiplicationContext) {}

// EnterRationalDivision is called when production rationalDivision is entered.
func (s *BaseRhumbParserListener) EnterRationalDivision(ctx *RationalDivisionContext) {}

// ExitRationalDivision is called when production rationalDivision is exited.
func (s *BaseRhumbParserListener) ExitRationalDivision(ctx *RationalDivisionContext) {}

// EnterWholeDivision is called when production wholeDivision is entered.
func (s *BaseRhumbParserListener) EnterWholeDivision(ctx *WholeDivisionContext) {}

// ExitWholeDivision is called when production wholeDivision is exited.
func (s *BaseRhumbParserListener) ExitWholeDivision(ctx *WholeDivisionContext) {}

// EnterRemainderDivision is called when production remainderDivision is entered.
func (s *BaseRhumbParserListener) EnterRemainderDivision(ctx *RemainderDivisionContext) {}

// ExitRemainderDivision is called when production remainderDivision is exited.
func (s *BaseRhumbParserListener) ExitRemainderDivision(ctx *RemainderDivisionContext) {}

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

// EnterRange is called when production range is entered.
func (s *BaseRhumbParserListener) EnterRange(ctx *RangeContext) {}

// ExitRange is called when production range is exited.
func (s *BaseRhumbParserListener) ExitRange(ctx *RangeContext) {}

// EnterScientific is called when production scientific is entered.
func (s *BaseRhumbParserListener) EnterScientific(ctx *ScientificContext) {}

// ExitScientific is called when production scientific is exited.
func (s *BaseRhumbParserListener) ExitScientific(ctx *ScientificContext) {}

// EnterImmutableLabel is called when production immutableLabel is entered.
func (s *BaseRhumbParserListener) EnterImmutableLabel(ctx *ImmutableLabelContext) {}

// ExitImmutableLabel is called when production immutableLabel is exited.
func (s *BaseRhumbParserListener) ExitImmutableLabel(ctx *ImmutableLabelContext) {}

// EnterMutableLabel is called when production mutableLabel is entered.
func (s *BaseRhumbParserListener) EnterMutableLabel(ctx *MutableLabelContext) {}

// ExitMutableLabel is called when production mutableLabel is exited.
func (s *BaseRhumbParserListener) ExitMutableLabel(ctx *MutableLabelContext) {}

// EnterEmptyPrefix is called when production emptyPrefix is entered.
func (s *BaseRhumbParserListener) EnterEmptyPrefix(ctx *EmptyPrefixContext) {}

// ExitEmptyPrefix is called when production emptyPrefix is exited.
func (s *BaseRhumbParserListener) ExitEmptyPrefix(ctx *EmptyPrefixContext) {}

// EnterFreezePrefix is called when production freezePrefix is entered.
func (s *BaseRhumbParserListener) EnterFreezePrefix(ctx *FreezePrefixContext) {}

// ExitFreezePrefix is called when production freezePrefix is exited.
func (s *BaseRhumbParserListener) ExitFreezePrefix(ctx *FreezePrefixContext) {}

// EnterCopyPrefix is called when production copyPrefix is entered.
func (s *BaseRhumbParserListener) EnterCopyPrefix(ctx *CopyPrefixContext) {}

// ExitCopyPrefix is called when production copyPrefix is exited.
func (s *BaseRhumbParserListener) ExitCopyPrefix(ctx *CopyPrefixContext) {}

// EnterToNumberPrefix is called when production toNumberPrefix is entered.
func (s *BaseRhumbParserListener) EnterToNumberPrefix(ctx *ToNumberPrefixContext) {}

// ExitToNumberPrefix is called when production toNumberPrefix is exited.
func (s *BaseRhumbParserListener) ExitToNumberPrefix(ctx *ToNumberPrefixContext) {}

// EnterNegateNumberPrefix is called when production negateNumberPrefix is entered.
func (s *BaseRhumbParserListener) EnterNegateNumberPrefix(ctx *NegateNumberPrefixContext) {}

// ExitNegateNumberPrefix is called when production negateNumberPrefix is exited.
func (s *BaseRhumbParserListener) ExitNegateNumberPrefix(ctx *NegateNumberPrefixContext) {}

// EnterToTruthPrefix is called when production toTruthPrefix is entered.
func (s *BaseRhumbParserListener) EnterToTruthPrefix(ctx *ToTruthPrefixContext) {}

// ExitToTruthPrefix is called when production toTruthPrefix is exited.
func (s *BaseRhumbParserListener) ExitToTruthPrefix(ctx *ToTruthPrefixContext) {}

// EnterNegateTruthPrefix is called when production negateTruthPrefix is entered.
func (s *BaseRhumbParserListener) EnterNegateTruthPrefix(ctx *NegateTruthPrefixContext) {}

// ExitNegateTruthPrefix is called when production negateTruthPrefix is exited.
func (s *BaseRhumbParserListener) ExitNegateTruthPrefix(ctx *NegateTruthPrefixContext) {}

// EnterVariadicPrefix is called when production variadicPrefix is entered.
func (s *BaseRhumbParserListener) EnterVariadicPrefix(ctx *VariadicPrefixContext) {}

// ExitVariadicPrefix is called when production variadicPrefix is exited.
func (s *BaseRhumbParserListener) ExitVariadicPrefix(ctx *VariadicPrefixContext) {}

// EnterArgumentPrefix is called when production argumentPrefix is entered.
func (s *BaseRhumbParserListener) EnterArgumentPrefix(ctx *ArgumentPrefixContext) {}

// ExitArgumentPrefix is called when production argumentPrefix is exited.
func (s *BaseRhumbParserListener) ExitArgumentPrefix(ctx *ArgumentPrefixContext) {}

// EnterSignalOutwardPrefix is called when production signalOutwardPrefix is entered.
func (s *BaseRhumbParserListener) EnterSignalOutwardPrefix(ctx *SignalOutwardPrefixContext) {}

// ExitSignalOutwardPrefix is called when production signalOutwardPrefix is exited.
func (s *BaseRhumbParserListener) ExitSignalOutwardPrefix(ctx *SignalOutwardPrefixContext) {}

// EnterSignalInwardPrefix is called when production signalInwardPrefix is entered.
func (s *BaseRhumbParserListener) EnterSignalInwardPrefix(ctx *SignalInwardPrefixContext) {}

// ExitSignalInwardPrefix is called when production signalInwardPrefix is exited.
func (s *BaseRhumbParserListener) ExitSignalInwardPrefix(ctx *SignalInwardPrefixContext) {}

// EnterNestedField is called when production nestedField is entered.
func (s *BaseRhumbParserListener) EnterNestedField(ctx *NestedFieldContext) {}

// ExitNestedField is called when production nestedField is exited.
func (s *BaseRhumbParserListener) ExitNestedField(ctx *NestedFieldContext) {}

// EnterDeeplyNestedField is called when production deeplyNestedField is entered.
func (s *BaseRhumbParserListener) EnterDeeplyNestedField(ctx *DeeplyNestedFieldContext) {}

// ExitDeeplyNestedField is called when production deeplyNestedField is exited.
func (s *BaseRhumbParserListener) ExitDeeplyNestedField(ctx *DeeplyNestedFieldContext) {}

// EnterNestedSubfield is called when production nestedSubfield is entered.
func (s *BaseRhumbParserListener) EnterNestedSubfield(ctx *NestedSubfieldContext) {}

// ExitNestedSubfield is called when production nestedSubfield is exited.
func (s *BaseRhumbParserListener) ExitNestedSubfield(ctx *NestedSubfieldContext) {}

// EnterDeeplyNestedSubfield is called when production deeplyNestedSubfield is entered.
func (s *BaseRhumbParserListener) EnterDeeplyNestedSubfield(ctx *DeeplyNestedSubfieldContext) {}

// ExitDeeplyNestedSubfield is called when production deeplyNestedSubfield is exited.
func (s *BaseRhumbParserListener) ExitDeeplyNestedSubfield(ctx *DeeplyNestedSubfieldContext) {}

// EnterDatePart is called when production datePart is entered.
func (s *BaseRhumbParserListener) EnterDatePart(ctx *DatePartContext) {}

// ExitDatePart is called when production datePart is exited.
func (s *BaseRhumbParserListener) ExitDatePart(ctx *DatePartContext) {}

// EnterDate is called when production date is entered.
func (s *BaseRhumbParserListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *BaseRhumbParserListener) ExitDate(ctx *DateContext) {}

// EnterText is called when production text is entered.
func (s *BaseRhumbParserListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseRhumbParserListener) ExitText(ctx *TextContext) {}

// EnterLabelInterp is called when production labelInterp is entered.
func (s *BaseRhumbParserListener) EnterLabelInterp(ctx *LabelInterpContext) {}

// ExitLabelInterp is called when production labelInterp is exited.
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

// EnterComputedRef is called when production computedRef is entered.
func (s *BaseRhumbParserListener) EnterComputedRef(ctx *ComputedRefContext) {}

// ExitComputedRef is called when production computedRef is exited.
func (s *BaseRhumbParserListener) ExitComputedRef(ctx *ComputedRefContext) {}

// EnterFunctionRef is called when production functionRef is entered.
func (s *BaseRhumbParserListener) EnterFunctionRef(ctx *FunctionRefContext) {}

// ExitFunctionRef is called when production functionRef is exited.
func (s *BaseRhumbParserListener) ExitFunctionRef(ctx *FunctionRefContext) {}

// EnterJunctionRef is called when production junctionRef is entered.
func (s *BaseRhumbParserListener) EnterJunctionRef(ctx *JunctionRefContext) {}

// ExitJunctionRef is called when production junctionRef is exited.
func (s *BaseRhumbParserListener) ExitJunctionRef(ctx *JunctionRefContext) {}
