// Code generated from RhumbParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // RhumbParser

import "github.com/antlr4-go/antlr/v4"

type BaseRhumbParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRhumbParserVisitor) VisitExpressions(ctx *ExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitFields(ctx *FieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitPatterns(ctx *PatternsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitTerminator(ctx *TerminatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitRationalNumber(ctx *RationalNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitDateNumber(ctx *DateNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitZeroNumber(ctx *ZeroNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitWholeNumber(ctx *WholeNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitKeySymbol(ctx *KeySymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitTextSymbol(ctx *TextSymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitReferenceLiteral(ctx *ReferenceLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitLabelSymbol(ctx *LabelSymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitFieldLiteral(ctx *FieldLiteralContext) interface{} {
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

func (v *BaseRhumbParserVisitor) VisitComparative(ctx *ComparativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitSimpleExpression(ctx *SimpleExpressionContext) interface{} {
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

func (v *BaseRhumbParserVisitor) VisitLibrary(ctx *LibraryContext) interface{} {
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

func (v *BaseRhumbParserVisitor) VisitAssignLabel(ctx *AssignLabelContext) interface{} {
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

func (v *BaseRhumbParserVisitor) VisitChainExpression(ctx *ChainExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitPrefixAssignMutField(ctx *PrefixAssignMutFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitPrefixAssignMutSubField(ctx *PrefixAssignMutSubFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitPrefixAssignImmSubField(ctx *PrefixAssignImmSubFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitPrefixSlurpSpread(ctx *PrefixSlurpSpreadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAssignMutField(ctx *AssignMutFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAssignMutSubField(ctx *AssignMutSubFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAssignImmField(ctx *AssignImmFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAssignImmSubField(ctx *AssignImmSubFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitSimpleMapField(ctx *SimpleMapFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitSimpleField(ctx *SimpleFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAssignBreakingPattern(ctx *AssignBreakingPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAssignFallthroughPattern(ctx *AssignFallthroughPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAssignDefaultPattern(ctx *AssignDefaultPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAppend(ctx *AppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitUnshift(ctx *UnshiftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitLength(ctx *LengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitEmpty(ctx *EmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAllSubfields(ctx *AllSubfieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitAllFields(ctx *AllFieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitElements(ctx *ElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitFreeze(ctx *FreezeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitCopy(ctx *CopyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitToDate(ctx *ToDateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitConstructor(ctx *ConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitBase(ctx *BaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitToNumber(ctx *ToNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNegateNumber(ctx *NegateNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitToTruth(ctx *ToTruthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNegateTruth(ctx *NegateTruthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitVariadic(ctx *VariadicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitToKey(ctx *ToKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitMethod(ctx *MethodContext) interface{} {
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

func (v *BaseRhumbParserVisitor) VisitIsInner(ctx *IsInnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitIsUnder(ctx *IsUnderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNotEqual(ctx *NotEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNotInner(ctx *NotInnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNotUnder(ctx *NotUnderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitConjunctiveOp(ctx *ConjunctiveOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitDisjunctiveOp(ctx *DisjunctiveOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitPipe(ctx *PipeContext) interface{} {
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

func (v *BaseRhumbParserVisitor) VisitDestructure(ctx *DestructureContext) interface{} {
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

func (v *BaseRhumbParserVisitor) VisitConcatenate(ctx *ConcatenateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitMultiplication(ctx *MultiplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitRationalDivision(ctx *RationalDivisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitWholeDivision(ctx *WholeDivisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitRemainderDivision(ctx *RemainderDivisionContext) interface{} {
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

func (v *BaseRhumbParserVisitor) VisitRange(ctx *RangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitScientific(ctx *ScientificContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitImmutableLabel(ctx *ImmutableLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitMutableLabel(ctx *MutableLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitEmptyPrefix(ctx *EmptyPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitFreezePrefix(ctx *FreezePrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitCopyPrefix(ctx *CopyPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitToNumberPrefix(ctx *ToNumberPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNegateNumberPrefix(ctx *NegateNumberPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitToTruthPrefix(ctx *ToTruthPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNegateTruthPrefix(ctx *NegateTruthPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitVariadicPrefix(ctx *VariadicPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitArgumentPrefix(ctx *ArgumentPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitSignalOutwardPrefix(ctx *SignalOutwardPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitSignalInwardPrefix(ctx *SignalInwardPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNestedField(ctx *NestedFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitDeeplyNestedField(ctx *DeeplyNestedFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitNestedSubfield(ctx *NestedSubfieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitDeeplyNestedSubfield(ctx *DeeplyNestedSubfieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitDatePart(ctx *DatePartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitDate(ctx *DateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitText(ctx *TextContext) interface{} {
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

func (v *BaseRhumbParserVisitor) VisitComputedRef(ctx *ComputedRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitFunctionRef(ctx *FunctionRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRhumbParserVisitor) VisitJunctionRef(ctx *JunctionRefContext) interface{} {
	return v.VisitChildren(ctx)
}
