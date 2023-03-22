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
	viLogger.Printf("Visit[tree type: %s]\n", reflect.TypeOf(tree))

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
		viLogger.Printf("VisitChildren[node type: %s]\n", reflect.TypeOf(n))
		v.Visit(n.(antlr.ParseTree))
	}
	return nil
}

func (v *RhumbVisitor) VisitSequence(ctx *P.SequenceContext) interface{} {
	viLogger.Println("sequence!")
	// logger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSimple(ctx *P.SimpleContext) interface{} {
	viLogger.Println("simple:", ctx.GetText())
	return v.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (v *RhumbVisitor) VisitMutableLabel(ctx *P.MutableLabelContext) interface{} {
	viLogger.Println("mutable label:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLabelLiteral(ctx *P.LabelLiteralContext) interface{} {
	var (
		text       string       = ctx.GetText()
		lits       vm.WordArray = v.vm.CurrentChunk.ReviveLits(&v.vm)
		lblIdx     uint64
		lblFindErr error
		ra         vm.RuneArray
	)
	viLogger.Println("label:", text)
	lblIdx, lblFindErr = lits.Find(&v.vm, text)
	if lblFindErr != nil {
		ra = vm.NewRuneArray(&v.vm, word.FromAddress(0), []rune(text)...)
		lblIdx = ra.Id()
	}
	v.vm.WriteCodeToCurrentChunk(
		ctx.GetStart().GetLine(),
		word.FromAddress(lblIdx),
		vm.NewLocalRequest,
	)
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAssignment(ctx *P.AssignmentContext) interface{} {
	viLogger.Println("assignment!")
	// TODO: implement map assignment
	var (
		text                  string
		ra                    vm.RuneArray
		lblIdx, opIdx         uint64
		lblFindErr, opFindErr error
		lits                  vm.WordArray = v.vm.CurrentChunk.ReviveLits(&v.vm)
	)
	if addr := ctx.GetAddress(); addr != nil {
		text = addr.GetText()
		viLogger.Println("Address:", text)
		// TODO: check for a matching subroutine and invoke
		// TODO: check for matching outer scoped label
		lblIdx, lblFindErr = lits.Find(&v.vm, text)
		if lblFindErr != nil {
			ra = vm.NewRuneArray(&v.vm, word.FromAddress(0), []rune(text)...)
			lblIdx = ra.Id()
		}

		v.vm.WriteCodeToCurrentChunk(
			addr.GetLine(),
			word.FromAddress(lblIdx),
			vm.NewLocalRequest,
		)
	} else {
		// FIXME: make this work as a reference
		addrRef := ctx.GetAddressRef()
		viLogger.Printf("addrRef.Accept(v): %v\n", addrRef.Accept(v))
		text = addrRef.GetText()
		viLogger.Println("AddressRef:", text)
		lblIdx, lblFindErr = lits.Find(&v.vm, text)
		if lblFindErr != nil {
			ra = vm.NewRuneArray(&v.vm, word.FromAddress(0), []rune(text)...)
			lblIdx = ra.Id()
		}
		v.vm.WriteCodeToCurrentChunk(
			addrRef.GetStart().GetLine(),
			word.FromAddress(lblIdx),
			vm.NewLocalRequest,
		)
	}
	viLogger.Println("LOCAL:", text)
	viLogger.Println("ctx.Expression().Accept...")
	viLogger.Println("ctx.Expression().Accept:", ctx.Expression().Accept(v))
	viLogger.Println("INNER:")
	op := ctx.AssignmentOp()
	opIdx, opFindErr = lits.Find(&v.vm, op.GetText())
	if opFindErr != nil {
		ra = vm.NewRuneArray(&v.vm, word.FromAddress(0), []rune(op.GetText())...)
		opIdx = ra.Id()
	}
	v.vm.WriteCodeToCurrentChunk(
		op.GetStart().GetLine(),
		word.FromAddress(opIdx),
		vm.NewOuterRequest,
	)
	return nil
}

func (v *RhumbVisitor) VisitFloatLiteral(ctx *P.FloatLiteralContext) interface{} {
	// logger.Println("float-lit!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIntegerLiteral(ctx *P.IntegerLiteralContext) interface{} {
	text := ctx.GetText()
	val, err := strconv.ParseInt(text, 10, 32)
	if err != nil {
		return RhumbReturn{nil, fmt.Errorf("unable to parse int")}
	}
	viLogger.Println("VALUE:", val)
	v.vm.WriteCodeToCurrentChunk(
		ctx.GetStart().GetLine(),
		word.FromInt(uint32(val)),
		vm.NewValueLiteral,
	)
	return RhumbReturn{val, nil}
}

func (v *RhumbVisitor) VisitStringLiteral(ctx *P.StringLiteralContext) interface{} {
	viLogger.Println("string-lit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitReferenceLiteral(ctx *P.ReferenceLiteralContext) interface{} {
	viLogger.Println("ref-lit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConjunctive(ctx *P.ConjunctiveContext) interface{} {
	viLogger.Println("conj:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAccess(ctx *P.AccessContext) interface{} {
	viLogger.Println("access:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitApplicative(ctx *P.ApplicativeContext) interface{} {
	viLogger.Println("applicative:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConditional(ctx *P.ConditionalContext) interface{} {
	viLogger.Println("cond:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitPrefix(ctx *P.PrefixContext) interface{} {
	viLogger.Println("prefix:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitComparative(ctx *P.ComparativeContext) interface{} {
	viLogger.Println("compare:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMultiplicative(ctx *P.MultiplicativeContext) interface{} {
	viLogger.Println("multiply:", ctx.GetText())
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

	v.vm.WriteCodeToCurrentChunk(
		mulOp.GetStart().GetLine(),
		word.FromAddress(ra.Id()),
		vm.NewOuterRequest, // FIXME: re-implement as NewInnerRequest
	)
	return nil
}

func (v *RhumbVisitor) VisitAdditive(ctx *P.AdditiveContext) interface{} {
	viLogger.Println("additive:", ctx.GetText())
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

	v.vm.WriteCodeToCurrentChunk(
		addOp.GetStart().GetLine(),
		word.FromAddress(ra.Id()),
		vm.NewOuterRequest, // FIXME: re-implement as NewInnerRequest
	)
	return nil
}

func (v *RhumbVisitor) VisitInvocation(ctx *P.InvocationContext) interface{} {
	viLogger.Println("invoke:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRoutine(ctx *P.RoutineContext) interface{} {
	viLogger.Println("routine:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDisjunctive(ctx *P.DisjunctiveContext) interface{} {
	viLogger.Println("disjunct:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIdentity(ctx *P.IdentityContext) interface{} {
	viLogger.Println("identify:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitEffect(ctx *P.EffectContext) interface{} {
	viLogger.Println("effect:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMember(ctx *P.MemberContext) interface{} {
	viLogger.Println("member:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSelector(ctx *P.SelectorContext) interface{} {
	viLogger.Println("selector:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitPower(ctx *P.PowerContext) interface{} {
	viLogger.Println("power:", ctx.GetText())
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
	v.vm.WriteCodeToCurrentChunk(
		powOp.GetStart().GetLine(),
		word.FromAddress(ra.Id()),
		vm.NewOuterRequest, // FIXME: re-implement as NewInnerRequest
	)
	return nil
}

func (v *RhumbVisitor) VisitMap(ctx *P.MapContext) interface{} {
	viLogger.Println("map:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitFreeze(ctx *P.FreezeContext) interface{} {
	viLogger.Println("freeze:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitInner(ctx *P.InnerContext) interface{} {
	viLogger.Println("inner:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLength(ctx *P.LengthContext) interface{} {
	viLogger.Println("length:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitFunction(ctx *P.FunctionContext) interface{} {
	viLogger.Println("function:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitJunction(ctx *P.JunctionContext) interface{} {
	viLogger.Println("junction:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitGreaterThan(ctx *P.GreaterThanContext) interface{} {
	viLogger.Println("greater-than:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitGreaterThanOrEqualTo(ctx *P.GreaterThanOrEqualToContext) interface{} {
	viLogger.Println("greater-than-or-equal-to:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLessThan(ctx *P.LessThanContext) interface{} {
	viLogger.Println("less-than:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLessThanOrEqualTo(ctx *P.LessThanOrEqualToContext) interface{} {
	viLogger.Println("less-than-or-equal-to:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsEqual(ctx *P.IsEqualContext) interface{} {
	viLogger.Println("is-equal:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsLike(ctx *P.IsLikeContext) interface{} {
	viLogger.Println("is-like:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsIn(ctx *P.IsInContext) interface{} {
	viLogger.Println("is-in:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsOverlayed(ctx *P.IsOverlayedContext) interface{} {
	viLogger.Println("is-overlayed:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsTopmost(ctx *P.IsTopmostContext) interface{} {
	viLogger.Println("is-topmost:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotEqual(ctx *P.NotEqualContext) interface{} {
	viLogger.Println("not-equal:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotLike(ctx *P.NotLikeContext) interface{} {
	viLogger.Println("not-like:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotIn(ctx *P.NotInContext) interface{} {
	viLogger.Println("not-in:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotOverlayed(ctx *P.NotOverlayedContext) interface{} {
	viLogger.Println("not-overlayed:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotTopmost(ctx *P.NotTopmostContext) interface{} {
	viLogger.Println("not-topmost:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConjunctiveOp(ctx *P.ConjunctiveOpContext) interface{} {
	viLogger.Println("conj-op:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDisjunctiveOp(ctx *P.DisjunctiveOpContext) interface{} {
	viLogger.Println("disjunct-op:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitOtherwise(ctx *P.OtherwiseContext) interface{} {
	viLogger.Println("otherwise:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDefault(ctx *P.DefaultContext) interface{} {
	viLogger.Println("default:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitForeach(ctx *P.ForeachContext) interface{} {
	viLogger.Println("for-each:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitWhile(ctx *P.WhileContext) interface{} {
	viLogger.Println("while:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitThen(ctx *P.ThenContext) interface{} {
	viLogger.Println("then:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitElse(ctx *P.ElseContext) interface{} {
	viLogger.Println("else:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAddition(ctx *P.AdditionContext) interface{} {
	viLogger.Println("addition:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDeviation(ctx *P.DeviationContext) interface{} {
	viLogger.Println("deviation:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSubtraction(ctx *P.SubtractionContext) interface{} {
	viLogger.Println("subtraction:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMultiplication(ctx *P.MultiplicationContext) interface{} {
	viLogger.Println("mupltication:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDivision(ctx *P.DivisionContext) interface{} {
	viLogger.Println("division:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIntegerDivision(ctx *P.IntegerDivisionContext) interface{} {
	viLogger.Println("int-division:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitModulo(ctx *P.ModuloContext) interface{} {
	viLogger.Println("modulo:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitBind(ctx *P.BindContext) interface{} {
	viLogger.Println("bind:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitExponent(ctx *P.ExponentContext) interface{} {
	viLogger.Println("exponenet:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRootExtraction(ctx *P.RootExtractionContext) interface{} {
	viLogger.Println("root-extract:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitScientific(ctx *P.ScientificContext) interface{} {
	viLogger.Println("scientific:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutablePair(ctx *P.ImmutablePairContext) interface{} {
	viLogger.Println("immutable-pair:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutableLabel(ctx *P.ImmutableLabelContext) interface{} {
	viLogger.Println("immutable-label:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMutablePair(ctx *P.MutablePairContext) interface{} {
	viLogger.Println("mutable-pair:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNumericalNegate(ctx *P.NumericalNegateContext) interface{} {
	viLogger.Println("numerical-negate:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitOuterScope(ctx *P.OuterScopeContext) interface{} {
	viLogger.Println("outer-scope:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLogicalNegate(ctx *P.LogicalNegateContext) interface{} {
	viLogger.Println("logical-negate:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAssert(ctx *P.AssertContext) interface{} {
	viLogger.Println("assert:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitArgument(ctx *P.ArgumentContext) interface{} {
	viLogger.Println("argument:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSlurpSpread(ctx *P.SlurpSpreadContext) interface{} {
	viLogger.Println("slurp-spread:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitBaseClone(ctx *P.BaseCloneContext) interface{} {
	viLogger.Println("base-clone:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNumericalPosit(ctx *P.NumericalPositContext) interface{} {
	viLogger.Println("numerical-posit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLogicalPosit(ctx *P.LogicalPositContext) interface{} {
	viLogger.Println("logical-posit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitOverlay(ctx *P.OverlayContext) interface{} {
	viLogger.Println("overlay:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitExistentialPosit(ctx *P.ExistentialPositContext) interface{} {
	viLogger.Println("existential-posit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutableDestruct(ctx *P.ImmutableDestructContext) interface{} {
	viLogger.Println("immutable-destruct:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMutableDestruct(ctx *P.MutableDestructContext) interface{} {
	viLogger.Println("mutable-destruct:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNestedLabel(ctx *P.NestedLabelContext) interface{} {
	viLogger.Println("nested-label:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNestedOverlay(ctx *P.NestedOverlayContext) interface{} {
	viLogger.Println("nested-overlay:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRange(ctx *P.RangeContext) interface{} {
	viLogger.Println("range:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitString(ctx *P.StringContext) interface{} {
	viLogger.Println("string:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLabelInterp(ctx *P.LabelInterpContext) interface{} {
	viLogger.Println("label-interp!")
	viLogger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRoutineInterp(ctx *P.RoutineInterpContext) interface{} {
	viLogger.Println("routine-interp:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSelectorInterp(ctx *P.SelectorInterpContext) interface{} {
	viLogger.Println("selector-interp:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNamedRef(ctx *P.NamedRefContext) interface{} {
	viLogger.Println("named-ref:", ctx.GetText())
	return ctx.Label().GetText()
}

func (v *RhumbVisitor) VisitFunctionRef(ctx *P.FunctionRefContext) interface{} {
	viLogger.Println("function-ref:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitComputedRef(ctx *P.ComputedRefContext) interface{} {
	viLogger.Println("computed-ref:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitJunctionRef(ctx *P.JunctionRefContext) interface{} {
	viLogger.Println("junction-ref:", ctx.GetText())
	return v.VisitChildren(ctx)
}
