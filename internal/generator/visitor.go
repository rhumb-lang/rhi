package generator

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"

	P "git.sr.ht/~madcapjake/grhumb/internal/parser"
	"git.sr.ht/~madcapjake/grhumb/internal/vm"
	"git.sr.ht/~madcapjake/grhumb/internal/word"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

var logger = log.New(io.Discard, "", log.LstdFlags)

func init() {
	if os.Getenv("RHUMB_VISITOR_DEBUG") == "1" {
		logger = log.New(os.Stdout, "{VI} ", log.LstdFlags)
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
	vm vm.VirtualMachine
}

func NewRhumbVisitor(vm vm.VirtualMachine) *RhumbVisitor {
	visitor := new(RhumbVisitor)
	visitor.vm = vm
	return visitor
}

func (v *RhumbVisitor) GetVM() *vm.VirtualMachine {
	return &v.vm
}

func (v *RhumbVisitor) Visit(tree antlr.ParseTree) interface{} {
	logger.Printf("input type: %s\n", reflect.TypeOf(tree))

	switch t := tree.(type) {
	case *antlr.ErrorNodeImpl:
		return RhumbReturn{nil, fmt.Errorf("syntax error near '%s'", t.GetText())}
	default:
		if val, ok := tree.Accept(v).(RhumbReturn); ok {
			return val
		}
	}
	return RhumbReturn{nil, fmt.Errorf("visit result not of type RhumbReturn")}
}

func (v *RhumbVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, n := range node.GetChildren() {
		logger.Println("Visit child:", reflect.TypeOf(n))
		v.Visit(n.(antlr.ParseTree))
	}
	return nil
}

func (v *RhumbVisitor) VisitSequence(ctx *P.SequenceContext) interface{} {
	logger.Println("sequence!")
	// logger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSimple(ctx *P.SimpleContext) interface{} {
	logger.Println("simple:", ctx.GetText())
	return v.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (v *RhumbVisitor) VisitMutableLabel(ctx *P.MutableLabelContext) interface{} {
	logger.Println("mutable label:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLabelLiteral(ctx *P.LabelLiteralContext) interface{} {
	var (
		text string       = ctx.GetText()
		ra   vm.RuneArray = vm.NewRuneArray(&v.vm, word.FromAddress(0), []rune(text)...)
	)
	logger.Println("label:", text)
	v.vm.WriteCodeToMain(
		ctx.GetStart().GetLine(),
		word.FromAddress(ra.Id()),
		vm.NewLocalRequest,
	)
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAssignment(ctx *P.AssignmentContext) interface{} {
	logger.Println("assignment!")
	// TODO: implement map assignment
	var (
		text string
		ra   vm.RuneArray
	)
	if addr := ctx.GetAddress(); addr != nil {
		text = addr.GetText()
		logger.Println("Address:", text)
		// TODO: check for a matching subroutine and invoke
		// TODO: check for matching outer scoped label

		ra = vm.NewRuneArray(&v.vm, word.FromAddress(0), []rune(text)...)

		v.vm.WriteCodeToMain(
			addr.GetLine(),
			word.FromAddress(ra.Id()),
			vm.NewLocalRequest,
		)
	} else {
		// FIXME: make this work as a reference
		addrRef := ctx.GetAddressRef()
		logger.Printf("addrRef.Accept(v): %v\n", addrRef.Accept(v))
		text = addrRef.GetText()
		logger.Println("AddressRef:", text)
		ra = vm.NewRuneArray(&v.vm, word.FromAddress(0), []rune(text)...)
		v.vm.WriteCodeToMain(
			addrRef.GetStart().GetLine(),
			word.FromAddress(ra.Id()),
			vm.NewLocalRequest,
		)
	}
	logger.Println("LOCAL:", text)
	logger.Println("ctx.Expression().Accept...")
	logger.Println("ctx.Expression().Accept:", ctx.Expression().Accept(v))
	logger.Println("INNER:")
	op := ctx.AssignmentOp()
	ra = vm.NewRuneArray(
		&v.vm,
		word.FromAddress(0),
		[]rune(op.GetText())...,
	)
	v.vm.WriteCodeToMain(
		op.GetStart().GetLine(),
		word.FromAddress(ra.Id()),
		vm.NewOuterRequest,
	)
	return nil
}

func (v *RhumbVisitor) VisitFloatLiteral(ctx *P.FloatLiteralContext) interface{} {
	// logger.Println("float-lit!")
	logger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIntegerLiteral(ctx *P.IntegerLiteralContext) interface{} {
	text := ctx.GetText()
	val, err := strconv.ParseInt(text, 10, 32)
	if err != nil {
		return RhumbReturn{nil, fmt.Errorf("unable to parse int")}
	}
	logger.Println("VALUE:", val)
	v.vm.WriteCodeToMain(
		ctx.GetStart().GetLine(),
		word.FromInt(uint32(val)),
		vm.NewValueLiteral,
	)
	return RhumbReturn{val, nil}
}

func (v *RhumbVisitor) VisitStringLiteral(ctx *P.StringLiteralContext) interface{} {
	logger.Println("string-lit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitReferenceLiteral(ctx *P.ReferenceLiteralContext) interface{} {
	logger.Println("ref-lit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConjunctive(ctx *P.ConjunctiveContext) interface{} {
	logger.Println("conj:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAccess(ctx *P.AccessContext) interface{} {
	logger.Println("access:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitApplicative(ctx *P.ApplicativeContext) interface{} {
	logger.Println("applicative:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConditional(ctx *P.ConditionalContext) interface{} {
	logger.Println("cond:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitPrefix(ctx *P.PrefixContext) interface{} {
	logger.Println("prefix:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitComparative(ctx *P.ComparativeContext) interface{} {
	logger.Println("compare:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMultiplicative(ctx *P.MultiplicativeContext) interface{} {
	logger.Println("multiply:", ctx.GetText())
	var (
		exprs []P.IExpressionContext     = ctx.AllExpression()
		mulOp P.IMultiplicativeOpContext = ctx.MultiplicativeOp()
		ra    vm.RuneArray               = vm.NewRuneArray(
			&v.vm,
			word.FromAddress(0),
			[]rune(mulOp.GetText())...,
		)
	)

	for i := range exprs {
		exprs[i].Accept(v)
	}

	v.vm.WriteCodeToMain(
		mulOp.GetStart().GetLine(),
		word.FromAddress(ra.Id()),
		vm.NewOuterRequest, // FIXME: re-implement as NewInnerRequest
	)
	return nil
}

func (v *RhumbVisitor) VisitAdditive(ctx *P.AdditiveContext) interface{} {
	logger.Println("additive:", ctx.GetText())
	var (
		exprs []P.IExpressionContext = ctx.AllExpression()
		addOp P.IAdditiveOpContext   = ctx.AdditiveOp()
		ra    vm.RuneArray           = vm.NewRuneArray(
			&v.vm,
			word.FromAddress(0),
			[]rune(addOp.GetText())...,
		)
	)

	for i := range exprs {
		exprs[i].Accept(v)
	}

	v.vm.WriteCodeToMain(
		addOp.GetStart().GetLine(),
		word.FromAddress(ra.Id()),
		vm.NewOuterRequest, // FIXME: re-implement as NewInnerRequest
	)
	return nil
}

func (v *RhumbVisitor) VisitInvocation(ctx *P.InvocationContext) interface{} {
	logger.Println("invoke:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRoutine(ctx *P.RoutineContext) interface{} {
	logger.Println("routine:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDisjunctive(ctx *P.DisjunctiveContext) interface{} {
	logger.Println("disjunct:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIdentity(ctx *P.IdentityContext) interface{} {
	logger.Println("identify:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitEffect(ctx *P.EffectContext) interface{} {
	logger.Println("effect:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMember(ctx *P.MemberContext) interface{} {
	logger.Println("member:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSelector(ctx *P.SelectorContext) interface{} {
	logger.Println("selector:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitPower(ctx *P.PowerContext) interface{} {
	logger.Println("power:", ctx.GetText())
	var (
		exprs []P.IExpressionContext     = ctx.AllExpression()
		powOp P.IExponentiationOpContext = ctx.ExponentiationOp()
		ra    vm.RuneArray               = vm.NewRuneArray(
			&v.vm,
			word.FromAddress(0),
			[]rune(powOp.GetText())...,
		)
	)
	for i := range exprs {
		exprs[i].Accept(v)
	}
	v.vm.WriteCodeToMain(
		powOp.GetStart().GetLine(),
		word.FromAddress(ra.Id()),
		vm.NewOuterRequest, // FIXME: re-implement as NewInnerRequest
	)
	return nil
}

func (v *RhumbVisitor) VisitMap(ctx *P.MapContext) interface{} {
	logger.Println("map:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitFreeze(ctx *P.FreezeContext) interface{} {
	logger.Println("freeze:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitInner(ctx *P.InnerContext) interface{} {
	logger.Println("inner:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLength(ctx *P.LengthContext) interface{} {
	logger.Println("length:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitFunction(ctx *P.FunctionContext) interface{} {
	logger.Println("function:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitJunction(ctx *P.JunctionContext) interface{} {
	logger.Println("junction:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitGreaterThan(ctx *P.GreaterThanContext) interface{} {
	logger.Println("greater-than:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitGreaterThanOrEqualTo(ctx *P.GreaterThanOrEqualToContext) interface{} {
	logger.Println("greater-than-or-equal-to:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLessThan(ctx *P.LessThanContext) interface{} {
	logger.Println("less-than:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLessThanOrEqualTo(ctx *P.LessThanOrEqualToContext) interface{} {
	logger.Println("less-than-or-equal-to:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsEqual(ctx *P.IsEqualContext) interface{} {
	logger.Println("is-equal:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsLike(ctx *P.IsLikeContext) interface{} {
	logger.Println("is-like:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsIn(ctx *P.IsInContext) interface{} {
	logger.Println("is-in:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsOverlayed(ctx *P.IsOverlayedContext) interface{} {
	logger.Println("is-overlayed:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsTopmost(ctx *P.IsTopmostContext) interface{} {
	logger.Println("is-topmost:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotEqual(ctx *P.NotEqualContext) interface{} {
	logger.Println("not-equal:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotLike(ctx *P.NotLikeContext) interface{} {
	logger.Println("not-like:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotIn(ctx *P.NotInContext) interface{} {
	logger.Println("not-in:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotOverlayed(ctx *P.NotOverlayedContext) interface{} {
	logger.Println("not-overlayed:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotTopmost(ctx *P.NotTopmostContext) interface{} {
	logger.Println("not-topmost:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConjunctiveOp(ctx *P.ConjunctiveOpContext) interface{} {
	logger.Println("conj-op:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDisjunctiveOp(ctx *P.DisjunctiveOpContext) interface{} {
	logger.Println("disjunct-op:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitOtherwise(ctx *P.OtherwiseContext) interface{} {
	logger.Println("otherwise:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDefault(ctx *P.DefaultContext) interface{} {
	logger.Println("default:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitForeach(ctx *P.ForeachContext) interface{} {
	logger.Println("for-each:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitWhile(ctx *P.WhileContext) interface{} {
	logger.Println("while:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitThen(ctx *P.ThenContext) interface{} {
	logger.Println("then:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitElse(ctx *P.ElseContext) interface{} {
	logger.Println("else:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAddition(ctx *P.AdditionContext) interface{} {
	logger.Println("addition:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDeviation(ctx *P.DeviationContext) interface{} {
	logger.Println("deviation:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSubtraction(ctx *P.SubtractionContext) interface{} {
	logger.Println("subtraction:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMultiplication(ctx *P.MultiplicationContext) interface{} {
	logger.Println("mupltication:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDivision(ctx *P.DivisionContext) interface{} {
	logger.Println("division:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIntegerDivision(ctx *P.IntegerDivisionContext) interface{} {
	logger.Println("int-division:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitModulo(ctx *P.ModuloContext) interface{} {
	logger.Println("modulo:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitBind(ctx *P.BindContext) interface{} {
	logger.Println("bind:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitExponent(ctx *P.ExponentContext) interface{} {
	logger.Println("exponenet:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRootExtraction(ctx *P.RootExtractionContext) interface{} {
	logger.Println("root-extract:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitScientific(ctx *P.ScientificContext) interface{} {
	logger.Println("scientific:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutablePair(ctx *P.ImmutablePairContext) interface{} {
	logger.Println("immutable-pair:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutableLabel(ctx *P.ImmutableLabelContext) interface{} {
	logger.Println("immutable-label:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMutablePair(ctx *P.MutablePairContext) interface{} {
	logger.Println("mutable-pair:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNumericalNegate(ctx *P.NumericalNegateContext) interface{} {
	logger.Println("numerical-negate:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitOuterScope(ctx *P.OuterScopeContext) interface{} {
	logger.Println("outer-scope:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLogicalNegate(ctx *P.LogicalNegateContext) interface{} {
	logger.Println("logical-negate:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAssert(ctx *P.AssertContext) interface{} {
	logger.Println("assert:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitArgument(ctx *P.ArgumentContext) interface{} {
	logger.Println("argument:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSlurpSpread(ctx *P.SlurpSpreadContext) interface{} {
	logger.Println("slurp-spread:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitBaseClone(ctx *P.BaseCloneContext) interface{} {
	logger.Println("base-clone:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNumericalPosit(ctx *P.NumericalPositContext) interface{} {
	logger.Println("numerical-posit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLogicalPosit(ctx *P.LogicalPositContext) interface{} {
	logger.Println("logical-posit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitOverlay(ctx *P.OverlayContext) interface{} {
	logger.Println("overlay:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitExistentialPosit(ctx *P.ExistentialPositContext) interface{} {
	logger.Println("existential-posit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutableDestruct(ctx *P.ImmutableDestructContext) interface{} {
	logger.Println("immutable-destruct:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMutableDestruct(ctx *P.MutableDestructContext) interface{} {
	logger.Println("mutable-destruct:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNestedLabel(ctx *P.NestedLabelContext) interface{} {
	logger.Println("nested-label:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNestedOverlay(ctx *P.NestedOverlayContext) interface{} {
	logger.Println("nested-overlay:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRange(ctx *P.RangeContext) interface{} {
	logger.Println("range:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitString(ctx *P.StringContext) interface{} {
	logger.Println("string:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLabelInterp(ctx *P.LabelInterpContext) interface{} {
	logger.Println("label-interp!")
	logger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRoutineInterp(ctx *P.RoutineInterpContext) interface{} {
	logger.Println("routine-interp:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSelectorInterp(ctx *P.SelectorInterpContext) interface{} {
	logger.Println("selector-interp:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNamedRef(ctx *P.NamedRefContext) interface{} {
	logger.Println("named-ref:", ctx.GetText())
	return ctx.Label().GetText()
}

func (v *RhumbVisitor) VisitFunctionRef(ctx *P.FunctionRefContext) interface{} {
	logger.Println("function-ref:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitComputedRef(ctx *P.ComputedRefContext) interface{} {
	logger.Println("computed-ref:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitJunctionRef(ctx *P.JunctionRefContext) interface{} {
	logger.Println("junction-ref:", ctx.GetText())
	return v.VisitChildren(ctx)
}
