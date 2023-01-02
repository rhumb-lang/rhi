package generator

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"

	"git.sr.ht/~madcapjake/grhumb/internal/parser"
	"git.sr.ht/~madcapjake/grhumb/internal/vm"
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
	parser.BaseRhumbParserVisitor
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

func (v *RhumbVisitor) VisitSequence(ctx *parser.SequenceContext) interface{} {
	logger.Println("sequence!")
	// logger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSimple(ctx *parser.SimpleContext) interface{} {
	logger.Println("simple:", ctx.GetText())
	return v.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (v *RhumbVisitor) VisitMutableLabel(ctx *parser.MutableLabelContext) interface{} {
	logger.Println("mutable label:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLabelLiteral(ctx *parser.LabelLiteralContext) interface{} {
	text := ctx.GetText()
	logger.Println("label:", text)
	v.vm.WriteCodeToMain(
		ctx.GetStart().GetLine(),
		vm.NewInstrRef(vm.RefLabel, text, nil),
		vm.NewLocalRequest,
	)
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	logger.Println("assignment!")
	// TODO: implement map assignment
	var text string
	if addr := ctx.GetAddress(); addr != nil {
		text = addr.GetText()
		logger.Println("Address:", text)
		// TODO: check for a matching subroutine and invoke
		// TODO: check for matching outer scoped label
		v.vm.WriteCodeToMain(
			addr.GetLine(),
			vm.NewInstrRef(vm.RefLabel, text, nil),
			vm.NewLocalRequest,
		)
	} else {
		// FIXME: make this work as a reference
		addrRef := ctx.GetAddressRef()
		logger.Printf("addrRef.Accept(v): %v\n", addrRef.Accept(v))
		text = addrRef.GetText()
		logger.Println("AddressRef:", text)
		v.vm.WriteCodeToMain(
			addrRef.GetStart().GetLine(),
			vm.NewInstrRef(vm.RefLabel, text, nil),
			vm.NewLocalRequest,
		)
	}
	logger.Println("LOCAL:", text)
	logger.Println("ctx.Expression().Accept...")
	logger.Println("ctx.Expression().Accept:", ctx.Expression().Accept(v))
	logger.Println("INNER:")
	assignOp := ctx.AssignmentOp()
	v.vm.WriteCodeToMain(
		assignOp.GetStart().GetLine(),
		vm.NewInstrRef(vm.RefLabel, assignOp.GetText(), nil),
		vm.NewOuterRequest,
	)
	return nil
}

func (v *RhumbVisitor) VisitFloatLiteral(ctx *parser.FloatLiteralContext) interface{} {
	// logger.Println("float-lit!")
	logger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIntegerLiteral(ctx *parser.IntegerLiteralContext) interface{} {
	text := ctx.GetText()
	val, err := strconv.ParseInt(text, 10, 32)
	if err != nil {
		return RhumbReturn{nil, fmt.Errorf("unable to parse int")}
	}
	logger.Println("VALUE:", val)
	v.vm.WriteCodeToMain(
		ctx.GetStart().GetLine(),
		vm.NewInstrRef(vm.RefInt, text, val),
		vm.NewValueLiteral,
	)
	return RhumbReturn{val, nil}
}

func (v *RhumbVisitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	logger.Println("string-lit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitReferenceLiteral(ctx *parser.ReferenceLiteralContext) interface{} {
	logger.Println("ref-lit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConjunctive(ctx *parser.ConjunctiveContext) interface{} {
	logger.Println("conj:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAccess(ctx *parser.AccessContext) interface{} {
	logger.Println("access:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitApplicative(ctx *parser.ApplicativeContext) interface{} {
	logger.Println("applicative:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConditional(ctx *parser.ConditionalContext) interface{} {
	logger.Println("cond:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitPrefix(ctx *parser.PrefixContext) interface{} {
	logger.Println("prefix:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitComparative(ctx *parser.ComparativeContext) interface{} {
	logger.Println("compare:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMultiplicative(ctx *parser.MultiplicativeContext) interface{} {
	logger.Println("multiply:", ctx.GetText())
	exprs := ctx.AllExpression()
	for i := range exprs {
		exprs[i].Accept(v)
	}
	mulOp := ctx.MultiplicativeOp()
	v.vm.WriteCodeToMain(
		mulOp.GetStart().GetLine(),
		vm.NewInstrRef(vm.RefLabel, mulOp.GetText(), nil),
		// FIXME: re-implement as NewInnerRequest
		vm.NewOuterRequest,
	)
	return nil
}

func (v *RhumbVisitor) VisitAdditive(ctx *parser.AdditiveContext) interface{} {
	logger.Println("additive:", ctx.GetText())
	exprs := ctx.AllExpression()
	for i := range exprs {
		exprs[i].Accept(v)
	}
	addOp := ctx.AdditiveOp()
	v.vm.WriteCodeToMain(
		addOp.GetStart().GetLine(),
		vm.NewInstrRef(vm.RefLabel, addOp.GetText(), nil),
		// FIXME: re-implement as NewInnerRequest
		vm.NewOuterRequest,
	)
	return nil
}

func (v *RhumbVisitor) VisitInvocation(ctx *parser.InvocationContext) interface{} {
	logger.Println("invoke:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRoutine(ctx *parser.RoutineContext) interface{} {
	logger.Println("routine:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDisjunctive(ctx *parser.DisjunctiveContext) interface{} {
	logger.Println("disjunct:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIdentity(ctx *parser.IdentityContext) interface{} {
	logger.Println("identify:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitEffect(ctx *parser.EffectContext) interface{} {
	logger.Println("effect:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMember(ctx *parser.MemberContext) interface{} {
	logger.Println("member:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSelector(ctx *parser.SelectorContext) interface{} {
	logger.Println("selector:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitPower(ctx *parser.PowerContext) interface{} {
	logger.Println("power:", ctx.GetText())
	exprs := ctx.AllExpression()
	for i := range exprs {
		exprs[i].Accept(v)
	}
	powOp := ctx.ExponentiationOp()
	v.vm.WriteCodeToMain(
		powOp.GetStart().GetLine(),
		vm.NewInstrRef(vm.RefLabel, powOp.GetText(), nil),
		// FIXME: re-implement as NewInnerRequest
		vm.NewOuterRequest,
	)
	return nil
}

func (v *RhumbVisitor) VisitMap(ctx *parser.MapContext) interface{} {
	logger.Println("map:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitFreeze(ctx *parser.FreezeContext) interface{} {
	logger.Println("freeze:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitInner(ctx *parser.InnerContext) interface{} {
	logger.Println("inner:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLength(ctx *parser.LengthContext) interface{} {
	logger.Println("length:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	logger.Println("function:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitJunction(ctx *parser.JunctionContext) interface{} {
	logger.Println("junction:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitGreaterThan(ctx *parser.GreaterThanContext) interface{} {
	logger.Println("greater-than:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitGreaterThanOrEqualTo(ctx *parser.GreaterThanOrEqualToContext) interface{} {
	logger.Println("greater-than-or-equal-to:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLessThan(ctx *parser.LessThanContext) interface{} {
	logger.Println("less-than:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLessThanOrEqualTo(ctx *parser.LessThanOrEqualToContext) interface{} {
	logger.Println("less-than-or-equal-to:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsEqual(ctx *parser.IsEqualContext) interface{} {
	logger.Println("is-equal:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsLike(ctx *parser.IsLikeContext) interface{} {
	logger.Println("is-like:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsIn(ctx *parser.IsInContext) interface{} {
	logger.Println("is-in:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsOverlayed(ctx *parser.IsOverlayedContext) interface{} {
	logger.Println("is-overlayed:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsTopmost(ctx *parser.IsTopmostContext) interface{} {
	logger.Println("is-topmost:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotEqual(ctx *parser.NotEqualContext) interface{} {
	logger.Println("not-equal:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotLike(ctx *parser.NotLikeContext) interface{} {
	logger.Println("not-like:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotIn(ctx *parser.NotInContext) interface{} {
	logger.Println("not-in:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotOverlayed(ctx *parser.NotOverlayedContext) interface{} {
	logger.Println("not-overlayed:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotTopmost(ctx *parser.NotTopmostContext) interface{} {
	logger.Println("not-topmost:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConjunctiveOp(ctx *parser.ConjunctiveOpContext) interface{} {
	logger.Println("conj-op:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDisjunctiveOp(ctx *parser.DisjunctiveOpContext) interface{} {
	logger.Println("disjunct-op:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitOtherwise(ctx *parser.OtherwiseContext) interface{} {
	logger.Println("otherwise:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDefault(ctx *parser.DefaultContext) interface{} {
	logger.Println("default:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitForeach(ctx *parser.ForeachContext) interface{} {
	logger.Println("for-each:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitWhile(ctx *parser.WhileContext) interface{} {
	logger.Println("while:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitThen(ctx *parser.ThenContext) interface{} {
	logger.Println("then:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitElse(ctx *parser.ElseContext) interface{} {
	logger.Println("else:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAddition(ctx *parser.AdditionContext) interface{} {
	logger.Println("addition:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDeviation(ctx *parser.DeviationContext) interface{} {
	logger.Println("deviation:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSubtraction(ctx *parser.SubtractionContext) interface{} {
	logger.Println("subtraction:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMultiplication(ctx *parser.MultiplicationContext) interface{} {
	logger.Println("mupltication:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDivision(ctx *parser.DivisionContext) interface{} {
	logger.Println("division:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIntegerDivision(ctx *parser.IntegerDivisionContext) interface{} {
	logger.Println("int-division:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitModulo(ctx *parser.ModuloContext) interface{} {
	logger.Println("modulo:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitBind(ctx *parser.BindContext) interface{} {
	logger.Println("bind:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitExponent(ctx *parser.ExponentContext) interface{} {
	logger.Println("exponenet:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRootExtraction(ctx *parser.RootExtractionContext) interface{} {
	logger.Println("root-extract:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitScientific(ctx *parser.ScientificContext) interface{} {
	logger.Println("scientific:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutablePair(ctx *parser.ImmutablePairContext) interface{} {
	logger.Println("immutable-pair:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutableLabel(ctx *parser.ImmutableLabelContext) interface{} {
	logger.Println("immutable-label:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMutablePair(ctx *parser.MutablePairContext) interface{} {
	logger.Println("mutable-pair:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNumericalNegate(ctx *parser.NumericalNegateContext) interface{} {
	logger.Println("numerical-negate:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitOuterScope(ctx *parser.OuterScopeContext) interface{} {
	logger.Println("outer-scope:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLogicalNegate(ctx *parser.LogicalNegateContext) interface{} {
	logger.Println("logical-negate:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAssert(ctx *parser.AssertContext) interface{} {
	logger.Println("assert:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitArgument(ctx *parser.ArgumentContext) interface{} {
	logger.Println("argument:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSlurpSpread(ctx *parser.SlurpSpreadContext) interface{} {
	logger.Println("slurp-spread:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitBaseClone(ctx *parser.BaseCloneContext) interface{} {
	logger.Println("base-clone:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNumericalPosit(ctx *parser.NumericalPositContext) interface{} {
	logger.Println("numerical-posit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLogicalPosit(ctx *parser.LogicalPositContext) interface{} {
	logger.Println("logical-posit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitOverlay(ctx *parser.OverlayContext) interface{} {
	logger.Println("overlay:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitExistentialPosit(ctx *parser.ExistentialPositContext) interface{} {
	logger.Println("existential-posit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutableDestruct(ctx *parser.ImmutableDestructContext) interface{} {
	logger.Println("immutable-destruct:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMutableDestruct(ctx *parser.MutableDestructContext) interface{} {
	logger.Println("mutable-destruct:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNestedLabel(ctx *parser.NestedLabelContext) interface{} {
	logger.Println("nested-label:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNestedOverlay(ctx *parser.NestedOverlayContext) interface{} {
	logger.Println("nested-overlay:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRange(ctx *parser.RangeContext) interface{} {
	logger.Println("range:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitString(ctx *parser.StringContext) interface{} {
	logger.Println("string:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLabelInterp(ctx *parser.LabelInterpContext) interface{} {
	logger.Println("label-interp!")
	logger.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRoutineInterp(ctx *parser.RoutineInterpContext) interface{} {
	logger.Println("routine-interp:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSelectorInterp(ctx *parser.SelectorInterpContext) interface{} {
	logger.Println("selector-interp:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNamedRef(ctx *parser.NamedRefContext) interface{} {
	logger.Println("named-ref:", ctx.GetText())
	return ctx.Label().GetText()
}

func (v *RhumbVisitor) VisitFunctionRef(ctx *parser.FunctionRefContext) interface{} {
	logger.Println("function-ref:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitComputedRef(ctx *parser.ComputedRefContext) interface{} {
	logger.Println("computed-ref:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitJunctionRef(ctx *parser.JunctionRefContext) interface{} {
	logger.Println("junction-ref:", ctx.GetText())
	return v.VisitChildren(ctx)
}
