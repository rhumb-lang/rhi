package generator

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"

	"git.sr.ht/~madcapjake/rhi/internal/code"
	"git.sr.ht/~madcapjake/rhi/internal/object"
	P "git.sr.ht/~madcapjake/rhi/internal/parser"
	virtual "git.sr.ht/~madcapjake/rhi/internal/vm"
	"github.com/antlr4-go/antlr/v4"
)

type codeKindContextKey int

const CodeKind codeKindContextKey = iota

var viLogger = log.New(io.Discard, "", log.LstdFlags)

func init() {
	if os.Getenv("RHUMB_VISITOR_DEBUG") == "1" {
		viLogger = log.New(os.Stdout, "{VI} ", log.Ltime)
	}
}

type RhumbErrorListener struct {
	*antlr.DefaultErrorListener
}

func (el *RhumbErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException) {
	fmt.Fprintf(os.Stderr, "ERROR: line %d:%d %s\n", line, column, msg)
}

var ErrorBlank = errors.New("blank line")

type RhumbReturn struct {
	Value any
	Error error
}

type RhumbVisitor struct {
	P.BaseRhumbParserVisitor
	VM *virtual.VirtualMachine
}

func NewRhumbVisitor(vm *virtual.VirtualMachine) *RhumbVisitor {
	visitor := new(RhumbVisitor)
	visitor.VM = vm
	return visitor
}

func (v *RhumbVisitor) Visit(tree antlr.ParseTree) interface{} {
	viLogger.Printf("Visit[tree type: %s]\n", reflect.TypeOf(tree))

	switch t := tree.(type) {
	case *antlr.ErrorNodeImpl:
		return RhumbReturn{
			nil,
			fmt.Errorf("syntax error near '%s'", t.GetText()),
		}
	default:
		if val, ok := tree.Accept(v).(RhumbReturn); ok {
			return val
		}
	}
	return RhumbReturn{
		nil,
		fmt.Errorf("visit result not of type RhumbReturn"),
	}
}

func (v *RhumbVisitor) VisitChildren(
	node antlr.RuleNode,
) interface{} {
	for _, n := range node.GetChildren() {
		viLogger.Printf(
			"VisitChildren[node type: %s]\n",
			reflect.TypeOf(n),
		)
		v.Visit(n.(antlr.ParseTree))
	}
	return nil
}

func (v *RhumbVisitor) VisitExpressions(
	ctx *P.ExpressionsContext,
) interface{} {
	viLogger.Println("expressions!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#fields.
func (v *RhumbVisitor) VisitFields(ctx *P.FieldsContext) interface{} {
	viLogger.Println("Fields!")
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#patterns.
func (v *RhumbVisitor) VisitPatterns(
	ctx *P.PatternsContext,
) interface{} {
	viLogger.Println("Patterns!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#terminator.
func (v *RhumbVisitor) VisitTerminator(
	ctx *P.TerminatorContext,
) interface{} {
	viLogger.Println("Terminator!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#rationalNumber.
func (v *RhumbVisitor) VisitRationalNumber(
	ctx *P.RationalNumberContext,
) interface{} {
	viLogger.Println("RationalNumber!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#dateNumber.
func (v *RhumbVisitor) VisitDateNumber(
	ctx *P.DateNumberContext,
) interface{} {
	viLogger.Println("DateNumber!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#zeroNumber.
func (v *RhumbVisitor) VisitZeroNumber(
	ctx *P.ZeroNumberContext,
) interface{} {
	viLogger.Println("ZeroNumber!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#wholeNumber.
func (v *RhumbVisitor) VisitWholeNumber(
	ctx *P.WholeNumberContext,
) interface{} {
	viLogger.Println("WholeNumber!")
	viLogger.Println(ctx.GetText())
	num, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		panic("Integer conv error")
	}
	obj := object.NewNumber(v.VM.Memory, num)
	objID := v.VM.RegisterObject(obj)
	start := ctx.GetStart()
	v.VM.Write(code.NewValue(
		start.GetLine(),
		start.GetColumn(),
		objID,
	))
	return nil
}

// Visit a parse tree produced by RhumbParser#keySymbol.
func (v *RhumbVisitor) VisitKeySymbol(
	ctx *P.KeySymbolContext,
) interface{} {
	viLogger.Println("KeySymbol!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#textSymbol.
func (v *RhumbVisitor) VisitTextSymbol(
	ctx *P.TextSymbolContext,
) interface{} {
	viLogger.Println("TextSymbol!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#referenceLiteral.
func (v *RhumbVisitor) VisitReferenceLiteral(
	ctx *P.ReferenceLiteralContext,
) interface{} {
	viLogger.Println("ReferenceLiteral!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#labelSymbol.
func (v *RhumbVisitor) VisitLabelSymbol(
	ctx *P.LabelSymbolContext,
) interface{} {
	viLogger.Println("LabelSymbol!")
	labelID := v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, ctx.GetText()),
	)
	start := ctx.GetStart()
	v.VM.Write(code.NewValue(
		start.GetLine(),
		start.GetColumn(),
		labelID,
	))
	return labelID
}

// Visit a parse tree produced by RhumbParser#fieldLiteral.
func (v *RhumbVisitor) VisitFieldLiteral(
	ctx *P.FieldLiteralContext,
) interface{} {
	viLogger.Println("FieldLiteral!")
	viLogger.Println(ctx.GetText())
	fieldId := v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, ctx.GetText()),
	)
	start := ctx.GetStart()
	v.VM.Write(code.NewValue(
		start.GetLine(),
		start.GetColumn(),
		fieldId,
	))
	return nil
}

// Visit a parse tree produced by RhumbParser#conjunctive.
func (v *RhumbVisitor) VisitConjunctive(
	ctx *P.ConjunctiveContext,
) interface{} {
	viLogger.Println("Conjunctive!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil
}

// Visit a parse tree produced by RhumbParser#access.
func (v *RhumbVisitor) VisitAccess(
	ctx *P.AccessContext,
) interface{} {
	viLogger.Println("Access!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#applicative.
func (v *RhumbVisitor) VisitApplicative(
	ctx *P.ApplicativeContext,
) interface{} {
	viLogger.Println("Applicative!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil
}

// Visit a parse tree produced by RhumbParser#conditional.
func (v *RhumbVisitor) VisitConditional(
	ctx *P.ConditionalContext,
) interface{} {
	viLogger.Println("Conditional!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil
}

// Visit a parse tree produced by RhumbParser#prefix.
func (v *RhumbVisitor) VisitPrefix(
	ctx *P.PrefixContext,
) interface{} {
	viLogger.Println("Prefix!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#comparative.
func (v *RhumbVisitor) VisitComparative(
	ctx *P.ComparativeContext,
) interface{} {
	viLogger.Println("Comparative!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil
}

// Visit a parse tree produced by RhumbParser#simpleExpression.
func (v *RhumbVisitor) VisitSimpleExpression(
	ctx *P.SimpleExpressionContext,
) interface{} {
	viLogger.Println("SimpleExpression!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#multiplicative.
func (v *RhumbVisitor) VisitMultiplicative(
	ctx *P.MultiplicativeContext,
) interface{} {
	viLogger.Println("Multiplicative!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil
}

// Visit a parse tree produced by RhumbParser#additive.
func (v *RhumbVisitor) VisitAdditive(
	ctx *P.AdditiveContext,
) interface{} {
	viLogger.Println("Additive!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil
}

// Visit a parse tree produced by RhumbParser#invocation.
func (v *RhumbVisitor) VisitInvocation(
	ctx *P.InvocationContext,
) interface{} {
	viLogger.Println("Invocation!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#library.
func (v *RhumbVisitor) VisitLibrary(
	ctx *P.LibraryContext,
) interface{} {
	viLogger.Println("Library!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#routine.
func (v *RhumbVisitor) VisitRoutine(
	ctx *P.RoutineContext,
) interface{} {
	viLogger.Println("Routine!")
	viLogger.Println(ctx.GetText())
	v.VM.PushRoutine()
	v.VisitChildren(ctx.Expressions())
	routine := v.VM.PopRoutine()
	routineID := v.VM.RegisterObject(routine)
	start := ctx.GetStart()
	v.VM.Write(code.NewValue(
		start.GetLine(),
		start.GetColumn(),
		routineID,
	))

	return RhumbReturn{
		routine,
		nil,
	}
}

// Visit a parse tree produced by RhumbParser#disjunctive.
func (v *RhumbVisitor) VisitDisjunctive(
	ctx *P.DisjunctiveContext,
) interface{} {
	viLogger.Println("Disjunctive!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil
}

// Visit a parse tree produced by RhumbParser#identity.
func (v *RhumbVisitor) VisitIdentity(
	ctx *P.IdentityContext,
) interface{} {
	viLogger.Println("Identity!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#assignLabel.
func (v *RhumbVisitor) VisitAssignLabel(
	ctx *P.AssignLabelContext,
) interface{} {
	viLogger.Println("AssignLabel!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil
}

// Visit a parse tree produced by RhumbParser#effect.
func (v *RhumbVisitor) VisitEffect(
	ctx *P.EffectContext,
) interface{} {
	viLogger.Println("Effect!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#member.
func (v *RhumbVisitor) VisitMember(
	ctx *P.MemberContext,
) interface{} {
	viLogger.Println("Member!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#selector.
func (v *RhumbVisitor) VisitSelector(
	ctx *P.SelectorContext,
) interface{} {
	viLogger.Println("Selector!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#power.
func (v *RhumbVisitor) VisitPower(
	ctx *P.PowerContext,
) interface{} {
	viLogger.Println("Power!")
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	v.Visit(ctx.GetChild(1).(antlr.ParseTree))
	return nil
}

func (v *RhumbVisitor) visitMapChildren(
	kind string,
	ctx antlr.ParserRuleContext,
) {
	var (
		s, e antlr.Token
		op   int
	)
	viLogger.Println(fmt.Sprint(kind, "Map!"))
	op = v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, "_[[_"),
	)
	s = ctx.GetStart()
	v.VM.Write(code.NewLocal(s.GetLine(), s.GetColumn(), op))

	op = v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, "_[>_"),
	)

	for _, n := range ctx.GetChildren() {
		viLogger.Printf(
			"Visit%sMapChild[node type: %s]\n",
			kind,
			reflect.TypeOf(n),
		)

		switch nType := n.(type) {
		case *antlr.TerminalNodeImpl:
			continue
		default:
			v.Visit(nType.(antlr.ParseTree))
			s = ctx.GetStart()
			v.VM.Write(code.NewLocal(s.GetLine(), s.GetColumn(), op))
		}
	}

	op = v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, "_]]_"),
	)
	e = ctx.GetStop()
	v.VM.Write(code.NewLocal(e.GetLine(), e.GetColumn(), op))
}

// Visit a parse tree produced by RhumbParser#map.
func (v *RhumbVisitor) VisitMap(
	ctx *P.MapContext,
) interface{} {
	viLogger.Println("Map!")
	v.visitMapChildren("List", ctx)
	return nil
}

// Visit a parse tree produced by RhumbParser#chainExpression.
func (v *RhumbVisitor) VisitChainExpression(
	ctx *P.ChainExpressionContext,
) interface{} {
	viLogger.Println("ChainExpression!")
	viLogger.Println(ctx.GetText())

	var (
		s antlr.Token
	)

	for _, n := range ctx.GetChildren() {
		viLogger.Printf(
			"VisitChainChild[node type: %s]\n",
			reflect.TypeOf(n),
		)

		s = ctx.GetStart()
		switch nTyped := n.(type) {
		case *P.FieldLiteralContext:
			ref := v.VM.RegisterObject(
				object.NewReference(
					v.VM.Memory,
					object.NewLabel(v.VM.Memory, nTyped.GetText()),
				),
			)
			v.VM.Write(code.NewLocal(s.GetLine(), s.GetColumn(), ref))
		case *P.ExpressionsContext:
			v.visitMapChildren("Invoke", ctx)
		case *antlr.TerminalNodeImpl:
			continue
		default:
			panic("unknown chain expression child")
			// v.Visit(nType.(antlr.ParseTree))
			// s = ctx.GetStart()
			// v.VM.Write(code.NewLocal(s.GetLine(), s.GetColumn(), op))
		}
	}
	return nil
	// return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#prefixAssignMutField.
func (v *RhumbVisitor) VisitPrefixAssignMutField(
	ctx *P.PrefixAssignMutFieldContext,
) interface{} {
	viLogger.Println("PrefixAssignMutField!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#prefixAssignMutSubField.
func (v *RhumbVisitor) VisitPrefixAssignMutSubField(
	ctx *P.PrefixAssignMutSubFieldContext,
) interface{} {
	viLogger.Println("PrefixAssignMutSubField!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#prefixAssignImmSubField.
func (v *RhumbVisitor) VisitPrefixAssignImmSubField(
	ctx *P.PrefixAssignImmSubFieldContext,
) interface{} {
	viLogger.Println("PrefixAssignImmSubField!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#prefixSlurpSpread.
func (v *RhumbVisitor) VisitPrefixSlurpSpread(
	ctx *P.PrefixSlurpSpreadContext,
) interface{} {
	viLogger.Println("PrefixSlurpSpread!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#assignMutField.
func (v *RhumbVisitor) VisitAssignMutField(
	ctx *P.AssignMutFieldContext,
) interface{} {
	viLogger.Println("AssignMutField!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#assignMutSubField.
func (v *RhumbVisitor) VisitAssignMutSubField(
	ctx *P.AssignMutSubFieldContext,
) interface{} {
	viLogger.Println("AssignMutSubField!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#assignImmField.
func (v *RhumbVisitor) VisitAssignImmField(
	ctx *P.AssignImmFieldContext,
) interface{} {
	viLogger.Println("AssignImmField!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#assignImmSubField.
func (v *RhumbVisitor) VisitAssignImmSubField(
	ctx *P.AssignImmSubFieldContext,
) interface{} {
	viLogger.Println("AssignImmSubField!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#simpleMapField.
func (v *RhumbVisitor) VisitSimpleMapField(
	ctx *P.SimpleMapFieldContext,
) interface{} {
	viLogger.Println("SimpleMapField!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#simpleField.
func (v *RhumbVisitor) VisitSimpleField(ctx *P.SimpleFieldContext) interface{} {
	viLogger.Println("SimpleField!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#assignBreakingPattern.
func (v *RhumbVisitor) VisitAssignBreakingPattern(
	ctx *P.AssignBreakingPatternContext,
) interface{} {
	viLogger.Println("AssignBreakingPattern!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#assignFallthroughPattern.
func (v *RhumbVisitor) VisitAssignFallthroughPattern(
	ctx *P.AssignFallthroughPatternContext,
) interface{} {
	viLogger.Println("AssignFallthroughPattern!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#assignDefaultPattern.
func (v *RhumbVisitor) VisitAssignDefaultPattern(
	ctx *P.AssignDefaultPatternContext,
) interface{} {
	viLogger.Println("AssignDefaultPattern!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#append.
func (v *RhumbVisitor) VisitAppend(ctx *P.AppendContext) interface{} {
	viLogger.Println("Append!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#unshift.
func (v *RhumbVisitor) VisitUnshift(ctx *P.UnshiftContext) interface{} {
	viLogger.Println("Unshift!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#length.
func (v *RhumbVisitor) VisitLength(ctx *P.LengthContext) interface{} {
	viLogger.Println("Length!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#empty.
func (v *RhumbVisitor) VisitEmpty(ctx *P.EmptyContext) interface{} {
	viLogger.Println("Empty!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#allSubfields.
func (v *RhumbVisitor) VisitAllSubfields(ctx *P.AllSubfieldsContext) interface{} {
	viLogger.Println("AllSubfields!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#allFields.
func (v *RhumbVisitor) VisitAllFields(ctx *P.AllFieldsContext) interface{} {
	viLogger.Println("AllFields!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#elements.
func (v *RhumbVisitor) VisitElements(ctx *P.ElementsContext) interface{} {
	viLogger.Println("Elements!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#freeze.
func (v *RhumbVisitor) VisitFreeze(ctx *P.FreezeContext) interface{} {
	viLogger.Println("Freeze!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#copy.
func (v *RhumbVisitor) VisitCopy(ctx *P.CopyContext) interface{} {
	viLogger.Println("Copy!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#toDate.
func (v *RhumbVisitor) VisitToDate(ctx *P.ToDateContext) interface{} {
	viLogger.Println("ToDate!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#parameters.
func (v *RhumbVisitor) VisitParameters(ctx *P.ParametersContext) interface{} {
	viLogger.Println("Parameters!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#constructor.
func (v *RhumbVisitor) VisitConstructor(ctx *P.ConstructorContext) interface{} {
	viLogger.Println("Constructor!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#base.
func (v *RhumbVisitor) VisitBase(ctx *P.BaseContext) interface{} {
	viLogger.Println("Base!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#toNumber.
func (v *RhumbVisitor) VisitToNumber(ctx *P.ToNumberContext) interface{} {
	viLogger.Println("ToNumber!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#negateNumber.
func (v *RhumbVisitor) VisitNegateNumber(ctx *P.NegateNumberContext) interface{} {
	viLogger.Println("NegateNumber!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#toTruth.
func (v *RhumbVisitor) VisitToTruth(ctx *P.ToTruthContext) interface{} {
	viLogger.Println("ToTruth!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#negateTruth.
func (v *RhumbVisitor) VisitNegateTruth(ctx *P.NegateTruthContext) interface{} {
	viLogger.Println("NegateTruth!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#variadic.
func (v *RhumbVisitor) VisitVariadic(ctx *P.VariadicContext) interface{} {
	viLogger.Println("Variadic!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#toKey.
func (v *RhumbVisitor) VisitToKey(ctx *P.ToKeyContext) interface{} {
	viLogger.Println("ToKey!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#function.
func (v *RhumbVisitor) VisitFunction(ctx *P.FunctionContext) interface{} {
	viLogger.Println("Function!")
	viLogger.Println(ctx.GetText())
	label := v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, fmt.Sprint("_", ctx.GetText(), "_")),
	)
	s := ctx.GetStart()
	v.VM.Write(code.NewLocal(s.GetLine(), s.GetColumn(), label))
	return nil
}

// Visit a parse tree produced by RhumbParser#method.
func (v *RhumbVisitor) VisitMethod(ctx *P.MethodContext) interface{} {
	viLogger.Println("Method!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#greaterThan.
func (v *RhumbVisitor) VisitGreaterThan(ctx *P.GreaterThanContext) interface{} {
	viLogger.Println("GreaterThan!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#greaterThanOrEqualTo.
func (v *RhumbVisitor) VisitGreaterThanOrEqualTo(ctx *P.GreaterThanOrEqualToContext) interface{} {
	viLogger.Println("GreaterThanOrEqualTo!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#lessThan.
func (v *RhumbVisitor) VisitLessThan(ctx *P.LessThanContext) interface{} {
	viLogger.Println("LessThan!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#lessThanOrEqualTo.
func (v *RhumbVisitor) VisitLessThanOrEqualTo(ctx *P.LessThanOrEqualToContext) interface{} {
	viLogger.Println("LessThanOrEqualTo!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#isEqual.
func (v *RhumbVisitor) VisitIsEqual(ctx *P.IsEqualContext) interface{} {
	viLogger.Println("IsEqual!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#isInner.
func (v *RhumbVisitor) VisitIsInner(ctx *P.IsInnerContext) interface{} {
	viLogger.Println("IsInner!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#isUnder.
func (v *RhumbVisitor) VisitIsUnder(ctx *P.IsUnderContext) interface{} {
	viLogger.Println("IsUnder!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#notEqual.
func (v *RhumbVisitor) VisitNotEqual(ctx *P.NotEqualContext) interface{} {
	viLogger.Println("NotEqual!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#notInner.
func (v *RhumbVisitor) VisitNotInner(ctx *P.NotInnerContext) interface{} {
	viLogger.Println("NotInner!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#notUnder.
func (v *RhumbVisitor) VisitNotUnder(ctx *P.NotUnderContext) interface{} {
	viLogger.Println("NotUnder!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#conjunctiveOp.
func (v *RhumbVisitor) VisitConjunctiveOp(ctx *P.ConjunctiveOpContext) interface{} {
	viLogger.Println("ConjunctiveOp!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#disjunctiveOp.
func (v *RhumbVisitor) VisitDisjunctiveOp(ctx *P.DisjunctiveOpContext) interface{} {
	viLogger.Println("DisjunctiveOp!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#pipe.
func (v *RhumbVisitor) VisitPipe(ctx *P.PipeContext) interface{} {
	viLogger.Println("Pipe!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#default.
func (v *RhumbVisitor) VisitDefault(ctx *P.DefaultContext) interface{} {
	viLogger.Println("Default!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#foreach.
func (v *RhumbVisitor) VisitForeach(ctx *P.ForeachContext) interface{} {
	viLogger.Println("Foreach!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#while.
func (v *RhumbVisitor) VisitWhile(ctx *P.WhileContext) interface{} {
	viLogger.Println("While!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#then.
func (v *RhumbVisitor) VisitThen(ctx *P.ThenContext) interface{} {
	viLogger.Println("Then!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#else.
func (v *RhumbVisitor) VisitElse(ctx *P.ElseContext) interface{} {
	viLogger.Println("Else!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#destructure.
func (v *RhumbVisitor) VisitDestructure(ctx *P.DestructureContext) interface{} {
	viLogger.Println("Destructure!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#addition.
func (v *RhumbVisitor) VisitAddition(ctx *P.AdditionContext) interface{} {
	viLogger.Println("Addition!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#deviation.
func (v *RhumbVisitor) VisitDeviation(ctx *P.DeviationContext) interface{} {
	viLogger.Println("Deviation!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#subtraction.
func (v *RhumbVisitor) VisitSubtraction(ctx *P.SubtractionContext) interface{} {
	viLogger.Println("Subtraction!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#concatenate.
func (v *RhumbVisitor) VisitConcatenate(ctx *P.ConcatenateContext) interface{} {
	viLogger.Println("Concatenate!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#multiplication.
func (v *RhumbVisitor) VisitMultiplication(ctx *P.MultiplicationContext) interface{} {
	viLogger.Println("Multiplication!")
	viLogger.Println(ctx.GetText())

	op := v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, "_**_"),
	)
	s := ctx.GetStart()
	v.VM.Write(code.NewLocal(s.GetLine(), s.GetColumn(), op))
	return nil
}

// Visit a parse tree produced by RhumbParser#rationalDivision.
func (v *RhumbVisitor) VisitRationalDivision(ctx *P.RationalDivisionContext) interface{} {
	viLogger.Println("RationalDivision!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#wholeDivision.
func (v *RhumbVisitor) VisitWholeDivision(ctx *P.WholeDivisionContext) interface{} {
	viLogger.Println("WholeDivision!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#remainderDivision.
func (v *RhumbVisitor) VisitRemainderDivision(ctx *P.RemainderDivisionContext) interface{} {
	viLogger.Println("RemainderDivision!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#bind.
func (v *RhumbVisitor) VisitBind(ctx *P.BindContext) interface{} {
	viLogger.Println("Bind!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#exponent.
func (v *RhumbVisitor) VisitExponent(ctx *P.ExponentContext) interface{} {
	viLogger.Println("Exponent!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#rootExtraction.
func (v *RhumbVisitor) VisitRootExtraction(ctx *P.RootExtractionContext) interface{} {
	viLogger.Println("RootExtraction!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#range.
func (v *RhumbVisitor) VisitRange(ctx *P.RangeContext) interface{} {
	viLogger.Println("Range!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#scientific.
func (v *RhumbVisitor) VisitScientific(ctx *P.ScientificContext) interface{} {
	viLogger.Println("Scientific!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#immutableLabel.
func (v *RhumbVisitor) VisitImmutableLabel(ctx *P.ImmutableLabelContext) interface{} {
	viLogger.Println("ImmutableLabel!")
	op := v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, fmt.Sprint("_", ctx.GetText(), "_")),
	)
	s := ctx.GetStart()
	v.VM.Write(code.NewLocal(s.GetLine(), s.GetColumn(), op))
	return nil
}

// Visit a parse tree produced by RhumbParser#mutableLabel.
func (v *RhumbVisitor) VisitMutableLabel(ctx *P.MutableLabelContext) interface{} {
	viLogger.Println("MutableLabel!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#emptyPrefix.
func (v *RhumbVisitor) VisitEmptyPrefix(ctx *P.EmptyPrefixContext) interface{} {
	viLogger.Println("EmptyPrefix!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#freezePrefix.
func (v *RhumbVisitor) VisitFreezePrefix(ctx *P.FreezePrefixContext) interface{} {
	viLogger.Println("FreezePrefix!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#copyPrefix.
func (v *RhumbVisitor) VisitCopyPrefix(ctx *P.CopyPrefixContext) interface{} {
	viLogger.Println("CopyPrefix!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#toNumberPrefix.
func (v *RhumbVisitor) VisitToNumberPrefix(ctx *P.ToNumberPrefixContext) interface{} {
	viLogger.Println("ToNumberPrefix!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#negateNumberPrefix.
func (v *RhumbVisitor) VisitNegateNumberPrefix(ctx *P.NegateNumberPrefixContext) interface{} {
	viLogger.Println("NegateNumberPrefix!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#toTruthPrefix.
func (v *RhumbVisitor) VisitToTruthPrefix(ctx *P.ToTruthPrefixContext) interface{} {
	viLogger.Println("ToTruthPrefix!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#negateTruthPrefix.
func (v *RhumbVisitor) VisitNegateTruthPrefix(ctx *P.NegateTruthPrefixContext) interface{} {
	viLogger.Println("NegateTruthPrefix!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#variadicPrefix.
func (v *RhumbVisitor) VisitVariadicPrefix(ctx *P.VariadicPrefixContext) interface{} {
	viLogger.Println("VariadicPrefix!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#argumentPrefix.
func (v *RhumbVisitor) VisitArgumentPrefix(ctx *P.ArgumentPrefixContext) interface{} {
	viLogger.Println("ArgumentPrefix!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#signalOutwardPrefix.
func (v *RhumbVisitor) VisitSignalOutwardPrefix(ctx *P.SignalOutwardPrefixContext) interface{} {
	viLogger.Println("SignalOutwardPrefix!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#signalInwardPrefix.
func (v *RhumbVisitor) VisitSignalInwardPrefix(ctx *P.SignalInwardPrefixContext) interface{} {
	viLogger.Println("SignalInwardPrefix!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#nestedField.
func (v *RhumbVisitor) VisitNestedField(ctx *P.NestedFieldContext) interface{} {
	viLogger.Println("NestedField!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#deeplyNestedField.
func (v *RhumbVisitor) VisitDeeplyNestedField(ctx *P.DeeplyNestedFieldContext) interface{} {
	viLogger.Println("DeeplyNestedField!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#nestedSubfield.
func (v *RhumbVisitor) VisitNestedSubfield(ctx *P.NestedSubfieldContext) interface{} {
	viLogger.Println("NestedSubfield!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#deeplyNestedSubfield.
func (v *RhumbVisitor) VisitDeeplyNestedSubfield(ctx *P.DeeplyNestedSubfieldContext) interface{} {
	viLogger.Println("DeeplyNestedSubfield!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#datePart.
func (v *RhumbVisitor) VisitDatePart(ctx *P.DatePartContext) interface{} {
	viLogger.Println("DatePart!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#date.
func (v *RhumbVisitor) VisitDate(ctx *P.DateContext) interface{} {
	viLogger.Println("Date!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#text.
func (v *RhumbVisitor) VisitText(ctx *P.TextContext) interface{} {
	viLogger.Println("Text!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#labelInterp.
func (v *RhumbVisitor) VisitLabelInterp(ctx *P.LabelInterpContext) interface{} {
	viLogger.Println("LabelInterp!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#routineInterp.
func (v *RhumbVisitor) VisitRoutineInterp(ctx *P.RoutineInterpContext) interface{} {
	viLogger.Println("RoutineInterp!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#selectorInterp.
func (v *RhumbVisitor) VisitSelectorInterp(ctx *P.SelectorInterpContext) interface{} {
	viLogger.Println("SelectorInterp!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#namedRef.
func (v *RhumbVisitor) VisitNamedRef(ctx *P.NamedRefContext) interface{} {
	viLogger.Println("NamedRef!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#computedRef.
func (v *RhumbVisitor) VisitComputedRef(ctx *P.ComputedRefContext) interface{} {
	viLogger.Println("ComputedRef!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#functionRef.
func (v *RhumbVisitor) VisitFunctionRef(ctx *P.FunctionRefContext) interface{} {
	viLogger.Println("FunctionRef!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#junctionRef.
func (v *RhumbVisitor) VisitJunctionRef(ctx *P.JunctionRefContext) interface{} {
	viLogger.Println("JunctionRef!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}
