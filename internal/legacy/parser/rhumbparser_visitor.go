// Code generated from RhumbParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // RhumbParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by RhumbParser.
type RhumbParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RhumbParser#expressions.
	VisitExpressions(ctx *ExpressionsContext) interface{}

	// Visit a parse tree produced by RhumbParser#fields.
	VisitFields(ctx *FieldsContext) interface{}

	// Visit a parse tree produced by RhumbParser#patterns.
	VisitPatterns(ctx *PatternsContext) interface{}

	// Visit a parse tree produced by RhumbParser#terminator.
	VisitTerminator(ctx *TerminatorContext) interface{}

	// Visit a parse tree produced by RhumbParser#rationalNumber.
	VisitRationalNumber(ctx *RationalNumberContext) interface{}

	// Visit a parse tree produced by RhumbParser#dateNumber.
	VisitDateNumber(ctx *DateNumberContext) interface{}

	// Visit a parse tree produced by RhumbParser#zeroNumber.
	VisitZeroNumber(ctx *ZeroNumberContext) interface{}

	// Visit a parse tree produced by RhumbParser#wholeNumber.
	VisitWholeNumber(ctx *WholeNumberContext) interface{}

	// Visit a parse tree produced by RhumbParser#keySymbol.
	VisitKeySymbol(ctx *KeySymbolContext) interface{}

	// Visit a parse tree produced by RhumbParser#textSymbol.
	VisitTextSymbol(ctx *TextSymbolContext) interface{}

	// Visit a parse tree produced by RhumbParser#referenceLiteral.
	VisitReferenceLiteral(ctx *ReferenceLiteralContext) interface{}

	// Visit a parse tree produced by RhumbParser#labelSymbol.
	VisitLabelSymbol(ctx *LabelSymbolContext) interface{}

	// Visit a parse tree produced by RhumbParser#fieldLiteral.
	VisitFieldLiteral(ctx *FieldLiteralContext) interface{}

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

	// Visit a parse tree produced by RhumbParser#comparative.
	VisitComparative(ctx *ComparativeContext) interface{}

	// Visit a parse tree produced by RhumbParser#simpleExpression.
	VisitSimpleExpression(ctx *SimpleExpressionContext) interface{}

	// Visit a parse tree produced by RhumbParser#multiplicative.
	VisitMultiplicative(ctx *MultiplicativeContext) interface{}

	// Visit a parse tree produced by RhumbParser#additive.
	VisitAdditive(ctx *AdditiveContext) interface{}

	// Visit a parse tree produced by RhumbParser#invocation.
	VisitInvocation(ctx *InvocationContext) interface{}

	// Visit a parse tree produced by RhumbParser#library.
	VisitLibrary(ctx *LibraryContext) interface{}

	// Visit a parse tree produced by RhumbParser#routine.
	VisitRoutine(ctx *RoutineContext) interface{}

	// Visit a parse tree produced by RhumbParser#disjunctive.
	VisitDisjunctive(ctx *DisjunctiveContext) interface{}

	// Visit a parse tree produced by RhumbParser#identity.
	VisitIdentity(ctx *IdentityContext) interface{}

	// Visit a parse tree produced by RhumbParser#assignLabel.
	VisitAssignLabel(ctx *AssignLabelContext) interface{}

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

	// Visit a parse tree produced by RhumbParser#chainExpression.
	VisitChainExpression(ctx *ChainExpressionContext) interface{}

	// Visit a parse tree produced by RhumbParser#prefixAssignMutField.
	VisitPrefixAssignMutField(ctx *PrefixAssignMutFieldContext) interface{}

	// Visit a parse tree produced by RhumbParser#prefixAssignMutSubField.
	VisitPrefixAssignMutSubField(ctx *PrefixAssignMutSubFieldContext) interface{}

	// Visit a parse tree produced by RhumbParser#prefixAssignImmSubField.
	VisitPrefixAssignImmSubField(ctx *PrefixAssignImmSubFieldContext) interface{}

	// Visit a parse tree produced by RhumbParser#prefixSlurpSpread.
	VisitPrefixSlurpSpread(ctx *PrefixSlurpSpreadContext) interface{}

	// Visit a parse tree produced by RhumbParser#assignMutField.
	VisitAssignMutField(ctx *AssignMutFieldContext) interface{}

	// Visit a parse tree produced by RhumbParser#assignMutSubField.
	VisitAssignMutSubField(ctx *AssignMutSubFieldContext) interface{}

	// Visit a parse tree produced by RhumbParser#assignImmField.
	VisitAssignImmField(ctx *AssignImmFieldContext) interface{}

	// Visit a parse tree produced by RhumbParser#assignImmSubField.
	VisitAssignImmSubField(ctx *AssignImmSubFieldContext) interface{}

	// Visit a parse tree produced by RhumbParser#simpleMapField.
	VisitSimpleMapField(ctx *SimpleMapFieldContext) interface{}

	// Visit a parse tree produced by RhumbParser#simpleField.
	VisitSimpleField(ctx *SimpleFieldContext) interface{}

	// Visit a parse tree produced by RhumbParser#assignBreakingPattern.
	VisitAssignBreakingPattern(ctx *AssignBreakingPatternContext) interface{}

	// Visit a parse tree produced by RhumbParser#assignFallthroughPattern.
	VisitAssignFallthroughPattern(ctx *AssignFallthroughPatternContext) interface{}

	// Visit a parse tree produced by RhumbParser#assignDefaultPattern.
	VisitAssignDefaultPattern(ctx *AssignDefaultPatternContext) interface{}

	// Visit a parse tree produced by RhumbParser#append.
	VisitAppend(ctx *AppendContext) interface{}

	// Visit a parse tree produced by RhumbParser#unshift.
	VisitUnshift(ctx *UnshiftContext) interface{}

	// Visit a parse tree produced by RhumbParser#length.
	VisitLength(ctx *LengthContext) interface{}

	// Visit a parse tree produced by RhumbParser#empty.
	VisitEmpty(ctx *EmptyContext) interface{}

	// Visit a parse tree produced by RhumbParser#allSubfields.
	VisitAllSubfields(ctx *AllSubfieldsContext) interface{}

	// Visit a parse tree produced by RhumbParser#allFields.
	VisitAllFields(ctx *AllFieldsContext) interface{}

	// Visit a parse tree produced by RhumbParser#elements.
	VisitElements(ctx *ElementsContext) interface{}

	// Visit a parse tree produced by RhumbParser#freeze.
	VisitFreeze(ctx *FreezeContext) interface{}

	// Visit a parse tree produced by RhumbParser#copy.
	VisitCopy(ctx *CopyContext) interface{}

	// Visit a parse tree produced by RhumbParser#toDate.
	VisitToDate(ctx *ToDateContext) interface{}

	// Visit a parse tree produced by RhumbParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by RhumbParser#constructor.
	VisitConstructor(ctx *ConstructorContext) interface{}

	// Visit a parse tree produced by RhumbParser#base.
	VisitBase(ctx *BaseContext) interface{}

	// Visit a parse tree produced by RhumbParser#toNumber.
	VisitToNumber(ctx *ToNumberContext) interface{}

	// Visit a parse tree produced by RhumbParser#negateNumber.
	VisitNegateNumber(ctx *NegateNumberContext) interface{}

	// Visit a parse tree produced by RhumbParser#toTruth.
	VisitToTruth(ctx *ToTruthContext) interface{}

	// Visit a parse tree produced by RhumbParser#negateTruth.
	VisitNegateTruth(ctx *NegateTruthContext) interface{}

	// Visit a parse tree produced by RhumbParser#variadic.
	VisitVariadic(ctx *VariadicContext) interface{}

	// Visit a parse tree produced by RhumbParser#toKey.
	VisitToKey(ctx *ToKeyContext) interface{}

	// Visit a parse tree produced by RhumbParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by RhumbParser#method.
	VisitMethod(ctx *MethodContext) interface{}

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

	// Visit a parse tree produced by RhumbParser#isInner.
	VisitIsInner(ctx *IsInnerContext) interface{}

	// Visit a parse tree produced by RhumbParser#isUnder.
	VisitIsUnder(ctx *IsUnderContext) interface{}

	// Visit a parse tree produced by RhumbParser#notEqual.
	VisitNotEqual(ctx *NotEqualContext) interface{}

	// Visit a parse tree produced by RhumbParser#notInner.
	VisitNotInner(ctx *NotInnerContext) interface{}

	// Visit a parse tree produced by RhumbParser#notUnder.
	VisitNotUnder(ctx *NotUnderContext) interface{}

	// Visit a parse tree produced by RhumbParser#conjunctiveOp.
	VisitConjunctiveOp(ctx *ConjunctiveOpContext) interface{}

	// Visit a parse tree produced by RhumbParser#disjunctiveOp.
	VisitDisjunctiveOp(ctx *DisjunctiveOpContext) interface{}

	// Visit a parse tree produced by RhumbParser#pipe.
	VisitPipe(ctx *PipeContext) interface{}

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

	// Visit a parse tree produced by RhumbParser#destructure.
	VisitDestructure(ctx *DestructureContext) interface{}

	// Visit a parse tree produced by RhumbParser#addition.
	VisitAddition(ctx *AdditionContext) interface{}

	// Visit a parse tree produced by RhumbParser#deviation.
	VisitDeviation(ctx *DeviationContext) interface{}

	// Visit a parse tree produced by RhumbParser#subtraction.
	VisitSubtraction(ctx *SubtractionContext) interface{}

	// Visit a parse tree produced by RhumbParser#concatenate.
	VisitConcatenate(ctx *ConcatenateContext) interface{}

	// Visit a parse tree produced by RhumbParser#multiplication.
	VisitMultiplication(ctx *MultiplicationContext) interface{}

	// Visit a parse tree produced by RhumbParser#rationalDivision.
	VisitRationalDivision(ctx *RationalDivisionContext) interface{}

	// Visit a parse tree produced by RhumbParser#wholeDivision.
	VisitWholeDivision(ctx *WholeDivisionContext) interface{}

	// Visit a parse tree produced by RhumbParser#remainderDivision.
	VisitRemainderDivision(ctx *RemainderDivisionContext) interface{}

	// Visit a parse tree produced by RhumbParser#bind.
	VisitBind(ctx *BindContext) interface{}

	// Visit a parse tree produced by RhumbParser#exponent.
	VisitExponent(ctx *ExponentContext) interface{}

	// Visit a parse tree produced by RhumbParser#rootExtraction.
	VisitRootExtraction(ctx *RootExtractionContext) interface{}

	// Visit a parse tree produced by RhumbParser#range.
	VisitRange(ctx *RangeContext) interface{}

	// Visit a parse tree produced by RhumbParser#scientific.
	VisitScientific(ctx *ScientificContext) interface{}

	// Visit a parse tree produced by RhumbParser#immutableLabel.
	VisitImmutableLabel(ctx *ImmutableLabelContext) interface{}

	// Visit a parse tree produced by RhumbParser#mutableLabel.
	VisitMutableLabel(ctx *MutableLabelContext) interface{}

	// Visit a parse tree produced by RhumbParser#emptyPrefix.
	VisitEmptyPrefix(ctx *EmptyPrefixContext) interface{}

	// Visit a parse tree produced by RhumbParser#freezePrefix.
	VisitFreezePrefix(ctx *FreezePrefixContext) interface{}

	// Visit a parse tree produced by RhumbParser#copyPrefix.
	VisitCopyPrefix(ctx *CopyPrefixContext) interface{}

	// Visit a parse tree produced by RhumbParser#toNumberPrefix.
	VisitToNumberPrefix(ctx *ToNumberPrefixContext) interface{}

	// Visit a parse tree produced by RhumbParser#negateNumberPrefix.
	VisitNegateNumberPrefix(ctx *NegateNumberPrefixContext) interface{}

	// Visit a parse tree produced by RhumbParser#toTruthPrefix.
	VisitToTruthPrefix(ctx *ToTruthPrefixContext) interface{}

	// Visit a parse tree produced by RhumbParser#negateTruthPrefix.
	VisitNegateTruthPrefix(ctx *NegateTruthPrefixContext) interface{}

	// Visit a parse tree produced by RhumbParser#variadicPrefix.
	VisitVariadicPrefix(ctx *VariadicPrefixContext) interface{}

	// Visit a parse tree produced by RhumbParser#argumentPrefix.
	VisitArgumentPrefix(ctx *ArgumentPrefixContext) interface{}

	// Visit a parse tree produced by RhumbParser#signalOutwardPrefix.
	VisitSignalOutwardPrefix(ctx *SignalOutwardPrefixContext) interface{}

	// Visit a parse tree produced by RhumbParser#signalInwardPrefix.
	VisitSignalInwardPrefix(ctx *SignalInwardPrefixContext) interface{}

	// Visit a parse tree produced by RhumbParser#nestedField.
	VisitNestedField(ctx *NestedFieldContext) interface{}

	// Visit a parse tree produced by RhumbParser#deeplyNestedField.
	VisitDeeplyNestedField(ctx *DeeplyNestedFieldContext) interface{}

	// Visit a parse tree produced by RhumbParser#nestedSubfield.
	VisitNestedSubfield(ctx *NestedSubfieldContext) interface{}

	// Visit a parse tree produced by RhumbParser#deeplyNestedSubfield.
	VisitDeeplyNestedSubfield(ctx *DeeplyNestedSubfieldContext) interface{}

	// Visit a parse tree produced by RhumbParser#datePart.
	VisitDatePart(ctx *DatePartContext) interface{}

	// Visit a parse tree produced by RhumbParser#date.
	VisitDate(ctx *DateContext) interface{}

	// Visit a parse tree produced by RhumbParser#text.
	VisitText(ctx *TextContext) interface{}

	// Visit a parse tree produced by RhumbParser#labelInterp.
	VisitLabelInterp(ctx *LabelInterpContext) interface{}

	// Visit a parse tree produced by RhumbParser#routineInterp.
	VisitRoutineInterp(ctx *RoutineInterpContext) interface{}

	// Visit a parse tree produced by RhumbParser#selectorInterp.
	VisitSelectorInterp(ctx *SelectorInterpContext) interface{}

	// Visit a parse tree produced by RhumbParser#namedRef.
	VisitNamedRef(ctx *NamedRefContext) interface{}

	// Visit a parse tree produced by RhumbParser#computedRef.
	VisitComputedRef(ctx *ComputedRefContext) interface{}

	// Visit a parse tree produced by RhumbParser#functionRef.
	VisitFunctionRef(ctx *FunctionRefContext) interface{}

	// Visit a parse tree produced by RhumbParser#junctionRef.
	VisitJunctionRef(ctx *JunctionRefContext) interface{}
}
