// Code generated from RhumbParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // RhumbParser

import "github.com/antlr4-go/antlr/v4"

// RhumbParserListener is a complete listener for a parse tree produced by RhumbParser.
type RhumbParserListener interface {
	antlr.ParseTreeListener

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterFields is called when entering the fields production.
	EnterFields(c *FieldsContext)

	// EnterPatterns is called when entering the patterns production.
	EnterPatterns(c *PatternsContext)

	// EnterTerminator is called when entering the terminator production.
	EnterTerminator(c *TerminatorContext)

	// EnterRationalNumber is called when entering the rationalNumber production.
	EnterRationalNumber(c *RationalNumberContext)

	// EnterDateNumber is called when entering the dateNumber production.
	EnterDateNumber(c *DateNumberContext)

	// EnterZeroNumber is called when entering the zeroNumber production.
	EnterZeroNumber(c *ZeroNumberContext)

	// EnterWholeNumber is called when entering the wholeNumber production.
	EnterWholeNumber(c *WholeNumberContext)

	// EnterKeySymbol is called when entering the keySymbol production.
	EnterKeySymbol(c *KeySymbolContext)

	// EnterTextSymbol is called when entering the textSymbol production.
	EnterTextSymbol(c *TextSymbolContext)

	// EnterReferenceLiteral is called when entering the referenceLiteral production.
	EnterReferenceLiteral(c *ReferenceLiteralContext)

	// EnterLabelSymbol is called when entering the labelSymbol production.
	EnterLabelSymbol(c *LabelSymbolContext)

	// EnterFieldLiteral is called when entering the fieldLiteral production.
	EnterFieldLiteral(c *FieldLiteralContext)

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

	// EnterComparative is called when entering the comparative production.
	EnterComparative(c *ComparativeContext)

	// EnterSimpleExpression is called when entering the simpleExpression production.
	EnterSimpleExpression(c *SimpleExpressionContext)

	// EnterMultiplicative is called when entering the multiplicative production.
	EnterMultiplicative(c *MultiplicativeContext)

	// EnterAdditive is called when entering the additive production.
	EnterAdditive(c *AdditiveContext)

	// EnterInvocation is called when entering the invocation production.
	EnterInvocation(c *InvocationContext)

	// EnterLibrary is called when entering the library production.
	EnterLibrary(c *LibraryContext)

	// EnterRoutine is called when entering the routine production.
	EnterRoutine(c *RoutineContext)

	// EnterDisjunctive is called when entering the disjunctive production.
	EnterDisjunctive(c *DisjunctiveContext)

	// EnterIdentity is called when entering the identity production.
	EnterIdentity(c *IdentityContext)

	// EnterAssignLabel is called when entering the assignLabel production.
	EnterAssignLabel(c *AssignLabelContext)

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

	// EnterChainExpression is called when entering the chainExpression production.
	EnterChainExpression(c *ChainExpressionContext)

	// EnterPrefixAssignMutField is called when entering the prefixAssignMutField production.
	EnterPrefixAssignMutField(c *PrefixAssignMutFieldContext)

	// EnterPrefixAssignMutSubField is called when entering the prefixAssignMutSubField production.
	EnterPrefixAssignMutSubField(c *PrefixAssignMutSubFieldContext)

	// EnterPrefixAssignImmSubField is called when entering the prefixAssignImmSubField production.
	EnterPrefixAssignImmSubField(c *PrefixAssignImmSubFieldContext)

	// EnterPrefixSlurpSpread is called when entering the prefixSlurpSpread production.
	EnterPrefixSlurpSpread(c *PrefixSlurpSpreadContext)

	// EnterAssignMutField is called when entering the assignMutField production.
	EnterAssignMutField(c *AssignMutFieldContext)

	// EnterAssignMutSubField is called when entering the assignMutSubField production.
	EnterAssignMutSubField(c *AssignMutSubFieldContext)

	// EnterAssignImmField is called when entering the assignImmField production.
	EnterAssignImmField(c *AssignImmFieldContext)

	// EnterAssignImmSubField is called when entering the assignImmSubField production.
	EnterAssignImmSubField(c *AssignImmSubFieldContext)

	// EnterSimpleMapField is called when entering the simpleMapField production.
	EnterSimpleMapField(c *SimpleMapFieldContext)

	// EnterSimpleField is called when entering the simpleField production.
	EnterSimpleField(c *SimpleFieldContext)

	// EnterAssignBreakingPattern is called when entering the assignBreakingPattern production.
	EnterAssignBreakingPattern(c *AssignBreakingPatternContext)

	// EnterAssignFallthroughPattern is called when entering the assignFallthroughPattern production.
	EnterAssignFallthroughPattern(c *AssignFallthroughPatternContext)

	// EnterAssignDefaultPattern is called when entering the assignDefaultPattern production.
	EnterAssignDefaultPattern(c *AssignDefaultPatternContext)

	// EnterAppend is called when entering the append production.
	EnterAppend(c *AppendContext)

	// EnterUnshift is called when entering the unshift production.
	EnterUnshift(c *UnshiftContext)

	// EnterLength is called when entering the length production.
	EnterLength(c *LengthContext)

	// EnterEmpty is called when entering the empty production.
	EnterEmpty(c *EmptyContext)

	// EnterAllSubfields is called when entering the allSubfields production.
	EnterAllSubfields(c *AllSubfieldsContext)

	// EnterAllFields is called when entering the allFields production.
	EnterAllFields(c *AllFieldsContext)

	// EnterElements is called when entering the elements production.
	EnterElements(c *ElementsContext)

	// EnterFreeze is called when entering the freeze production.
	EnterFreeze(c *FreezeContext)

	// EnterCopy is called when entering the copy production.
	EnterCopy(c *CopyContext)

	// EnterToDate is called when entering the toDate production.
	EnterToDate(c *ToDateContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterConstructor is called when entering the constructor production.
	EnterConstructor(c *ConstructorContext)

	// EnterBase is called when entering the base production.
	EnterBase(c *BaseContext)

	// EnterToNumber is called when entering the toNumber production.
	EnterToNumber(c *ToNumberContext)

	// EnterNegateNumber is called when entering the negateNumber production.
	EnterNegateNumber(c *NegateNumberContext)

	// EnterToTruth is called when entering the toTruth production.
	EnterToTruth(c *ToTruthContext)

	// EnterNegateTruth is called when entering the negateTruth production.
	EnterNegateTruth(c *NegateTruthContext)

	// EnterVariadic is called when entering the variadic production.
	EnterVariadic(c *VariadicContext)

	// EnterToKey is called when entering the toKey production.
	EnterToKey(c *ToKeyContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

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

	// EnterIsInner is called when entering the isInner production.
	EnterIsInner(c *IsInnerContext)

	// EnterIsUnder is called when entering the isUnder production.
	EnterIsUnder(c *IsUnderContext)

	// EnterNotEqual is called when entering the notEqual production.
	EnterNotEqual(c *NotEqualContext)

	// EnterNotInner is called when entering the notInner production.
	EnterNotInner(c *NotInnerContext)

	// EnterNotUnder is called when entering the notUnder production.
	EnterNotUnder(c *NotUnderContext)

	// EnterConjunctiveOp is called when entering the conjunctiveOp production.
	EnterConjunctiveOp(c *ConjunctiveOpContext)

	// EnterDisjunctiveOp is called when entering the disjunctiveOp production.
	EnterDisjunctiveOp(c *DisjunctiveOpContext)

	// EnterPipe is called when entering the pipe production.
	EnterPipe(c *PipeContext)

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

	// EnterDestructure is called when entering the destructure production.
	EnterDestructure(c *DestructureContext)

	// EnterAddition is called when entering the addition production.
	EnterAddition(c *AdditionContext)

	// EnterDeviation is called when entering the deviation production.
	EnterDeviation(c *DeviationContext)

	// EnterSubtraction is called when entering the subtraction production.
	EnterSubtraction(c *SubtractionContext)

	// EnterConcatenate is called when entering the concatenate production.
	EnterConcatenate(c *ConcatenateContext)

	// EnterMultiplication is called when entering the multiplication production.
	EnterMultiplication(c *MultiplicationContext)

	// EnterRationalDivision is called when entering the rationalDivision production.
	EnterRationalDivision(c *RationalDivisionContext)

	// EnterWholeDivision is called when entering the wholeDivision production.
	EnterWholeDivision(c *WholeDivisionContext)

	// EnterRemainderDivision is called when entering the remainderDivision production.
	EnterRemainderDivision(c *RemainderDivisionContext)

	// EnterBind is called when entering the bind production.
	EnterBind(c *BindContext)

	// EnterExponent is called when entering the exponent production.
	EnterExponent(c *ExponentContext)

	// EnterRootExtraction is called when entering the rootExtraction production.
	EnterRootExtraction(c *RootExtractionContext)

	// EnterRange is called when entering the range production.
	EnterRange(c *RangeContext)

	// EnterScientific is called when entering the scientific production.
	EnterScientific(c *ScientificContext)

	// EnterImmutableLabel is called when entering the immutableLabel production.
	EnterImmutableLabel(c *ImmutableLabelContext)

	// EnterMutableLabel is called when entering the mutableLabel production.
	EnterMutableLabel(c *MutableLabelContext)

	// EnterEmptyPrefix is called when entering the emptyPrefix production.
	EnterEmptyPrefix(c *EmptyPrefixContext)

	// EnterFreezePrefix is called when entering the freezePrefix production.
	EnterFreezePrefix(c *FreezePrefixContext)

	// EnterCopyPrefix is called when entering the copyPrefix production.
	EnterCopyPrefix(c *CopyPrefixContext)

	// EnterToNumberPrefix is called when entering the toNumberPrefix production.
	EnterToNumberPrefix(c *ToNumberPrefixContext)

	// EnterNegateNumberPrefix is called when entering the negateNumberPrefix production.
	EnterNegateNumberPrefix(c *NegateNumberPrefixContext)

	// EnterToTruthPrefix is called when entering the toTruthPrefix production.
	EnterToTruthPrefix(c *ToTruthPrefixContext)

	// EnterNegateTruthPrefix is called when entering the negateTruthPrefix production.
	EnterNegateTruthPrefix(c *NegateTruthPrefixContext)

	// EnterVariadicPrefix is called when entering the variadicPrefix production.
	EnterVariadicPrefix(c *VariadicPrefixContext)

	// EnterArgumentPrefix is called when entering the argumentPrefix production.
	EnterArgumentPrefix(c *ArgumentPrefixContext)

	// EnterSignalOutwardPrefix is called when entering the signalOutwardPrefix production.
	EnterSignalOutwardPrefix(c *SignalOutwardPrefixContext)

	// EnterSignalInwardPrefix is called when entering the signalInwardPrefix production.
	EnterSignalInwardPrefix(c *SignalInwardPrefixContext)

	// EnterNestedField is called when entering the nestedField production.
	EnterNestedField(c *NestedFieldContext)

	// EnterDeeplyNestedField is called when entering the deeplyNestedField production.
	EnterDeeplyNestedField(c *DeeplyNestedFieldContext)

	// EnterNestedSubfield is called when entering the nestedSubfield production.
	EnterNestedSubfield(c *NestedSubfieldContext)

	// EnterDeeplyNestedSubfield is called when entering the deeplyNestedSubfield production.
	EnterDeeplyNestedSubfield(c *DeeplyNestedSubfieldContext)

	// EnterDatePart is called when entering the datePart production.
	EnterDatePart(c *DatePartContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterLabelInterp is called when entering the labelInterp production.
	EnterLabelInterp(c *LabelInterpContext)

	// EnterRoutineInterp is called when entering the routineInterp production.
	EnterRoutineInterp(c *RoutineInterpContext)

	// EnterSelectorInterp is called when entering the selectorInterp production.
	EnterSelectorInterp(c *SelectorInterpContext)

	// EnterNamedRef is called when entering the namedRef production.
	EnterNamedRef(c *NamedRefContext)

	// EnterComputedRef is called when entering the computedRef production.
	EnterComputedRef(c *ComputedRefContext)

	// EnterFunctionRef is called when entering the functionRef production.
	EnterFunctionRef(c *FunctionRefContext)

	// EnterJunctionRef is called when entering the junctionRef production.
	EnterJunctionRef(c *JunctionRefContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitFields is called when exiting the fields production.
	ExitFields(c *FieldsContext)

	// ExitPatterns is called when exiting the patterns production.
	ExitPatterns(c *PatternsContext)

	// ExitTerminator is called when exiting the terminator production.
	ExitTerminator(c *TerminatorContext)

	// ExitRationalNumber is called when exiting the rationalNumber production.
	ExitRationalNumber(c *RationalNumberContext)

	// ExitDateNumber is called when exiting the dateNumber production.
	ExitDateNumber(c *DateNumberContext)

	// ExitZeroNumber is called when exiting the zeroNumber production.
	ExitZeroNumber(c *ZeroNumberContext)

	// ExitWholeNumber is called when exiting the wholeNumber production.
	ExitWholeNumber(c *WholeNumberContext)

	// ExitKeySymbol is called when exiting the keySymbol production.
	ExitKeySymbol(c *KeySymbolContext)

	// ExitTextSymbol is called when exiting the textSymbol production.
	ExitTextSymbol(c *TextSymbolContext)

	// ExitReferenceLiteral is called when exiting the referenceLiteral production.
	ExitReferenceLiteral(c *ReferenceLiteralContext)

	// ExitLabelSymbol is called when exiting the labelSymbol production.
	ExitLabelSymbol(c *LabelSymbolContext)

	// ExitFieldLiteral is called when exiting the fieldLiteral production.
	ExitFieldLiteral(c *FieldLiteralContext)

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

	// ExitComparative is called when exiting the comparative production.
	ExitComparative(c *ComparativeContext)

	// ExitSimpleExpression is called when exiting the simpleExpression production.
	ExitSimpleExpression(c *SimpleExpressionContext)

	// ExitMultiplicative is called when exiting the multiplicative production.
	ExitMultiplicative(c *MultiplicativeContext)

	// ExitAdditive is called when exiting the additive production.
	ExitAdditive(c *AdditiveContext)

	// ExitInvocation is called when exiting the invocation production.
	ExitInvocation(c *InvocationContext)

	// ExitLibrary is called when exiting the library production.
	ExitLibrary(c *LibraryContext)

	// ExitRoutine is called when exiting the routine production.
	ExitRoutine(c *RoutineContext)

	// ExitDisjunctive is called when exiting the disjunctive production.
	ExitDisjunctive(c *DisjunctiveContext)

	// ExitIdentity is called when exiting the identity production.
	ExitIdentity(c *IdentityContext)

	// ExitAssignLabel is called when exiting the assignLabel production.
	ExitAssignLabel(c *AssignLabelContext)

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

	// ExitChainExpression is called when exiting the chainExpression production.
	ExitChainExpression(c *ChainExpressionContext)

	// ExitPrefixAssignMutField is called when exiting the prefixAssignMutField production.
	ExitPrefixAssignMutField(c *PrefixAssignMutFieldContext)

	// ExitPrefixAssignMutSubField is called when exiting the prefixAssignMutSubField production.
	ExitPrefixAssignMutSubField(c *PrefixAssignMutSubFieldContext)

	// ExitPrefixAssignImmSubField is called when exiting the prefixAssignImmSubField production.
	ExitPrefixAssignImmSubField(c *PrefixAssignImmSubFieldContext)

	// ExitPrefixSlurpSpread is called when exiting the prefixSlurpSpread production.
	ExitPrefixSlurpSpread(c *PrefixSlurpSpreadContext)

	// ExitAssignMutField is called when exiting the assignMutField production.
	ExitAssignMutField(c *AssignMutFieldContext)

	// ExitAssignMutSubField is called when exiting the assignMutSubField production.
	ExitAssignMutSubField(c *AssignMutSubFieldContext)

	// ExitAssignImmField is called when exiting the assignImmField production.
	ExitAssignImmField(c *AssignImmFieldContext)

	// ExitAssignImmSubField is called when exiting the assignImmSubField production.
	ExitAssignImmSubField(c *AssignImmSubFieldContext)

	// ExitSimpleMapField is called when exiting the simpleMapField production.
	ExitSimpleMapField(c *SimpleMapFieldContext)

	// ExitSimpleField is called when exiting the simpleField production.
	ExitSimpleField(c *SimpleFieldContext)

	// ExitAssignBreakingPattern is called when exiting the assignBreakingPattern production.
	ExitAssignBreakingPattern(c *AssignBreakingPatternContext)

	// ExitAssignFallthroughPattern is called when exiting the assignFallthroughPattern production.
	ExitAssignFallthroughPattern(c *AssignFallthroughPatternContext)

	// ExitAssignDefaultPattern is called when exiting the assignDefaultPattern production.
	ExitAssignDefaultPattern(c *AssignDefaultPatternContext)

	// ExitAppend is called when exiting the append production.
	ExitAppend(c *AppendContext)

	// ExitUnshift is called when exiting the unshift production.
	ExitUnshift(c *UnshiftContext)

	// ExitLength is called when exiting the length production.
	ExitLength(c *LengthContext)

	// ExitEmpty is called when exiting the empty production.
	ExitEmpty(c *EmptyContext)

	// ExitAllSubfields is called when exiting the allSubfields production.
	ExitAllSubfields(c *AllSubfieldsContext)

	// ExitAllFields is called when exiting the allFields production.
	ExitAllFields(c *AllFieldsContext)

	// ExitElements is called when exiting the elements production.
	ExitElements(c *ElementsContext)

	// ExitFreeze is called when exiting the freeze production.
	ExitFreeze(c *FreezeContext)

	// ExitCopy is called when exiting the copy production.
	ExitCopy(c *CopyContext)

	// ExitToDate is called when exiting the toDate production.
	ExitToDate(c *ToDateContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitConstructor is called when exiting the constructor production.
	ExitConstructor(c *ConstructorContext)

	// ExitBase is called when exiting the base production.
	ExitBase(c *BaseContext)

	// ExitToNumber is called when exiting the toNumber production.
	ExitToNumber(c *ToNumberContext)

	// ExitNegateNumber is called when exiting the negateNumber production.
	ExitNegateNumber(c *NegateNumberContext)

	// ExitToTruth is called when exiting the toTruth production.
	ExitToTruth(c *ToTruthContext)

	// ExitNegateTruth is called when exiting the negateTruth production.
	ExitNegateTruth(c *NegateTruthContext)

	// ExitVariadic is called when exiting the variadic production.
	ExitVariadic(c *VariadicContext)

	// ExitToKey is called when exiting the toKey production.
	ExitToKey(c *ToKeyContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

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

	// ExitIsInner is called when exiting the isInner production.
	ExitIsInner(c *IsInnerContext)

	// ExitIsUnder is called when exiting the isUnder production.
	ExitIsUnder(c *IsUnderContext)

	// ExitNotEqual is called when exiting the notEqual production.
	ExitNotEqual(c *NotEqualContext)

	// ExitNotInner is called when exiting the notInner production.
	ExitNotInner(c *NotInnerContext)

	// ExitNotUnder is called when exiting the notUnder production.
	ExitNotUnder(c *NotUnderContext)

	// ExitConjunctiveOp is called when exiting the conjunctiveOp production.
	ExitConjunctiveOp(c *ConjunctiveOpContext)

	// ExitDisjunctiveOp is called when exiting the disjunctiveOp production.
	ExitDisjunctiveOp(c *DisjunctiveOpContext)

	// ExitPipe is called when exiting the pipe production.
	ExitPipe(c *PipeContext)

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

	// ExitDestructure is called when exiting the destructure production.
	ExitDestructure(c *DestructureContext)

	// ExitAddition is called when exiting the addition production.
	ExitAddition(c *AdditionContext)

	// ExitDeviation is called when exiting the deviation production.
	ExitDeviation(c *DeviationContext)

	// ExitSubtraction is called when exiting the subtraction production.
	ExitSubtraction(c *SubtractionContext)

	// ExitConcatenate is called when exiting the concatenate production.
	ExitConcatenate(c *ConcatenateContext)

	// ExitMultiplication is called when exiting the multiplication production.
	ExitMultiplication(c *MultiplicationContext)

	// ExitRationalDivision is called when exiting the rationalDivision production.
	ExitRationalDivision(c *RationalDivisionContext)

	// ExitWholeDivision is called when exiting the wholeDivision production.
	ExitWholeDivision(c *WholeDivisionContext)

	// ExitRemainderDivision is called when exiting the remainderDivision production.
	ExitRemainderDivision(c *RemainderDivisionContext)

	// ExitBind is called when exiting the bind production.
	ExitBind(c *BindContext)

	// ExitExponent is called when exiting the exponent production.
	ExitExponent(c *ExponentContext)

	// ExitRootExtraction is called when exiting the rootExtraction production.
	ExitRootExtraction(c *RootExtractionContext)

	// ExitRange is called when exiting the range production.
	ExitRange(c *RangeContext)

	// ExitScientific is called when exiting the scientific production.
	ExitScientific(c *ScientificContext)

	// ExitImmutableLabel is called when exiting the immutableLabel production.
	ExitImmutableLabel(c *ImmutableLabelContext)

	// ExitMutableLabel is called when exiting the mutableLabel production.
	ExitMutableLabel(c *MutableLabelContext)

	// ExitEmptyPrefix is called when exiting the emptyPrefix production.
	ExitEmptyPrefix(c *EmptyPrefixContext)

	// ExitFreezePrefix is called when exiting the freezePrefix production.
	ExitFreezePrefix(c *FreezePrefixContext)

	// ExitCopyPrefix is called when exiting the copyPrefix production.
	ExitCopyPrefix(c *CopyPrefixContext)

	// ExitToNumberPrefix is called when exiting the toNumberPrefix production.
	ExitToNumberPrefix(c *ToNumberPrefixContext)

	// ExitNegateNumberPrefix is called when exiting the negateNumberPrefix production.
	ExitNegateNumberPrefix(c *NegateNumberPrefixContext)

	// ExitToTruthPrefix is called when exiting the toTruthPrefix production.
	ExitToTruthPrefix(c *ToTruthPrefixContext)

	// ExitNegateTruthPrefix is called when exiting the negateTruthPrefix production.
	ExitNegateTruthPrefix(c *NegateTruthPrefixContext)

	// ExitVariadicPrefix is called when exiting the variadicPrefix production.
	ExitVariadicPrefix(c *VariadicPrefixContext)

	// ExitArgumentPrefix is called when exiting the argumentPrefix production.
	ExitArgumentPrefix(c *ArgumentPrefixContext)

	// ExitSignalOutwardPrefix is called when exiting the signalOutwardPrefix production.
	ExitSignalOutwardPrefix(c *SignalOutwardPrefixContext)

	// ExitSignalInwardPrefix is called when exiting the signalInwardPrefix production.
	ExitSignalInwardPrefix(c *SignalInwardPrefixContext)

	// ExitNestedField is called when exiting the nestedField production.
	ExitNestedField(c *NestedFieldContext)

	// ExitDeeplyNestedField is called when exiting the deeplyNestedField production.
	ExitDeeplyNestedField(c *DeeplyNestedFieldContext)

	// ExitNestedSubfield is called when exiting the nestedSubfield production.
	ExitNestedSubfield(c *NestedSubfieldContext)

	// ExitDeeplyNestedSubfield is called when exiting the deeplyNestedSubfield production.
	ExitDeeplyNestedSubfield(c *DeeplyNestedSubfieldContext)

	// ExitDatePart is called when exiting the datePart production.
	ExitDatePart(c *DatePartContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitLabelInterp is called when exiting the labelInterp production.
	ExitLabelInterp(c *LabelInterpContext)

	// ExitRoutineInterp is called when exiting the routineInterp production.
	ExitRoutineInterp(c *RoutineInterpContext)

	// ExitSelectorInterp is called when exiting the selectorInterp production.
	ExitSelectorInterp(c *SelectorInterpContext)

	// ExitNamedRef is called when exiting the namedRef production.
	ExitNamedRef(c *NamedRefContext)

	// ExitComputedRef is called when exiting the computedRef production.
	ExitComputedRef(c *ComputedRefContext)

	// ExitFunctionRef is called when exiting the functionRef production.
	ExitFunctionRef(c *FunctionRefContext)

	// ExitJunctionRef is called when exiting the junctionRef production.
	ExitJunctionRef(c *JunctionRefContext)
}
