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
	logger.Printf("visit input type: %s\n", reflect.TypeOf(tree))

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
		fmt.Println("Visit child:", reflect.TypeOf(n))
		v.Visit(n.(antlr.ParseTree))
	}
	return nil
}

func (v *RhumbVisitor) VisitSequence(ctx *parser.SequenceContext) interface{} {
	fmt.Println("sequence!")
	// fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSimple(ctx *parser.SimpleContext) interface{} {
	fmt.Println("simple:", ctx.GetText())
	return v.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (v *RhumbVisitor) VisitMutableLabel(ctx *parser.MutableLabelContext) interface{} {
	fmt.Println("mutable label:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLabelLiteral(ctx *parser.LabelLiteralContext) interface{} {
	text := ctx.GetText()
	fmt.Println("label:", text)
	v.vm.WriteCodeToMain(
		ctx.GetStart().GetLine(),
		vm.NewInstrRef(vm.RefLabel, text, nil),
		vm.NewLocalRequest,
	)
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	fmt.Println("assignment!")
	// TODO: implement map assignment
	var text string
	if addr := ctx.GetAddress(); addr != nil {
		text = addr.GetText()
		fmt.Println("Address:", text)
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
		fmt.Printf("addrRef.Accept(v): %v\n", addrRef.Accept(v))
		text = addrRef.GetText()
		fmt.Println("AddressRef:", text)
		v.vm.WriteCodeToMain(
			addrRef.GetStart().GetLine(),
			vm.NewInstrRef(vm.RefLabel, text, nil),
			vm.NewLocalRequest,
		)
	}
	fmt.Println("LOCAL:", text)
	fmt.Println("ctx.Expression().Accept...")
	fmt.Println("ctx.Expression().Accept:", ctx.Expression().Accept(v))
	fmt.Println("INNER:")
	assignOp := ctx.AssignmentOp()
	v.vm.WriteCodeToMain(
		assignOp.GetStart().GetLine(),
		vm.NewInstrRef(vm.RefLabel, assignOp.GetText(), nil),
		vm.NewOuterRequest,
	)
	return nil
}

func (v *RhumbVisitor) VisitFloatLiteral(ctx *parser.FloatLiteralContext) interface{} {
	// fmt.Println("float-lit!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIntegerLiteral(ctx *parser.IntegerLiteralContext) interface{} {
	text := ctx.GetText()
	val, err := strconv.ParseInt(text, 10, 32)
	if err != nil {
		return RhumbReturn{nil, fmt.Errorf("unable to parse int")}
	}
	fmt.Println("VALUE:", val)
	v.vm.WriteCodeToMain(
		ctx.GetStart().GetLine(),
		vm.NewInstrRef(vm.RefInt, text, val),
		vm.NewValueLiteral,
	)
	return RhumbReturn{val, nil}
}

func (v *RhumbVisitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	fmt.Println("string-lit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitReferenceLiteral(ctx *parser.ReferenceLiteralContext) interface{} {
	fmt.Println("ref-lit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConjunctive(ctx *parser.ConjunctiveContext) interface{} {
	fmt.Println("conj:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAccess(ctx *parser.AccessContext) interface{} {
	fmt.Println("access:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitApplicative(ctx *parser.ApplicativeContext) interface{} {
	fmt.Println("applicative:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConditional(ctx *parser.ConditionalContext) interface{} {
	fmt.Println("cond:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitPrefix(ctx *parser.PrefixContext) interface{} {
	fmt.Println("prefix:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitComparative(ctx *parser.ComparativeContext) interface{} {
	fmt.Println("compare:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMultiplicative(ctx *parser.MultiplicativeContext) interface{} {
	fmt.Println("multiply:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAdditive(ctx *parser.AdditiveContext) interface{} {
	fmt.Println("additive:", ctx.GetText())
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
	fmt.Println("invoke:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRoutine(ctx *parser.RoutineContext) interface{} {
	fmt.Println("routine:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDisjunctive(ctx *parser.DisjunctiveContext) interface{} {
	fmt.Println("disjunct:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIdentity(ctx *parser.IdentityContext) interface{} {
	fmt.Println("identify:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitEffect(ctx *parser.EffectContext) interface{} {
	fmt.Println("effect:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMember(ctx *parser.MemberContext) interface{} {
	fmt.Println("member:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSelector(ctx *parser.SelectorContext) interface{} {
	fmt.Println("selector:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitPower(ctx *parser.PowerContext) interface{} {
	fmt.Println("power:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMap(ctx *parser.MapContext) interface{} {
	fmt.Println("map:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitFreeze(ctx *parser.FreezeContext) interface{} {
	fmt.Println("freeze:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitInner(ctx *parser.InnerContext) interface{} {
	fmt.Println("inner:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLength(ctx *parser.LengthContext) interface{} {
	fmt.Println("length:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	fmt.Println("function:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitJunction(ctx *parser.JunctionContext) interface{} {
	fmt.Println("junction:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitGreaterThan(ctx *parser.GreaterThanContext) interface{} {
	fmt.Println("greater-than:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitGreaterThanOrEqualTo(ctx *parser.GreaterThanOrEqualToContext) interface{} {
	fmt.Println("greater-than-or-equal-to:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLessThan(ctx *parser.LessThanContext) interface{} {
	fmt.Println("less-than:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLessThanOrEqualTo(ctx *parser.LessThanOrEqualToContext) interface{} {
	fmt.Println("less-than-or-equal-to:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsEqual(ctx *parser.IsEqualContext) interface{} {
	fmt.Println("is-equal:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsLike(ctx *parser.IsLikeContext) interface{} {
	fmt.Println("is-like:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsIn(ctx *parser.IsInContext) interface{} {
	fmt.Println("is-in:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsOverlayed(ctx *parser.IsOverlayedContext) interface{} {
	fmt.Println("is-overlayed:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsTopmost(ctx *parser.IsTopmostContext) interface{} {
	fmt.Println("is-topmost:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotEqual(ctx *parser.NotEqualContext) interface{} {
	fmt.Println("not-equal:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotLike(ctx *parser.NotLikeContext) interface{} {
	fmt.Println("not-like:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotIn(ctx *parser.NotInContext) interface{} {
	fmt.Println("not-in:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotOverlayed(ctx *parser.NotOverlayedContext) interface{} {
	fmt.Println("not-overlayed:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotTopmost(ctx *parser.NotTopmostContext) interface{} {
	fmt.Println("not-topmost:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConjunctiveOp(ctx *parser.ConjunctiveOpContext) interface{} {
	fmt.Println("conj-op:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDisjunctiveOp(ctx *parser.DisjunctiveOpContext) interface{} {
	fmt.Println("disjunct-op:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitOtherwise(ctx *parser.OtherwiseContext) interface{} {
	fmt.Println("otherwise:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDefault(ctx *parser.DefaultContext) interface{} {
	fmt.Println("default:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitForeach(ctx *parser.ForeachContext) interface{} {
	fmt.Println("for-each:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitWhile(ctx *parser.WhileContext) interface{} {
	fmt.Println("while:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitThen(ctx *parser.ThenContext) interface{} {
	fmt.Println("then:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitElse(ctx *parser.ElseContext) interface{} {
	fmt.Println("else:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAddition(ctx *parser.AdditionContext) interface{} {
	fmt.Println("addition:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDeviation(ctx *parser.DeviationContext) interface{} {
	fmt.Println("deviation:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSubtraction(ctx *parser.SubtractionContext) interface{} {
	fmt.Println("subtraction:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMultiplication(ctx *parser.MultiplicationContext) interface{} {
	fmt.Println("mupltication:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDivision(ctx *parser.DivisionContext) interface{} {
	fmt.Println("division:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIntegerDivision(ctx *parser.IntegerDivisionContext) interface{} {
	fmt.Println("int-division:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitModulo(ctx *parser.ModuloContext) interface{} {
	fmt.Println("modulo:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitBind(ctx *parser.BindContext) interface{} {
	fmt.Println("bind:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitExponent(ctx *parser.ExponentContext) interface{} {
	fmt.Println("exponenet:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRootExtraction(ctx *parser.RootExtractionContext) interface{} {
	fmt.Println("root-extract:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitScientific(ctx *parser.ScientificContext) interface{} {
	fmt.Println("scientific:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutablePair(ctx *parser.ImmutablePairContext) interface{} {
	fmt.Println("immutable-pair:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutableLabel(ctx *parser.ImmutableLabelContext) interface{} {
	fmt.Println("immutable-label:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMutablePair(ctx *parser.MutablePairContext) interface{} {
	fmt.Println("mutable-pair:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNumericalNegate(ctx *parser.NumericalNegateContext) interface{} {
	fmt.Println("numerical-negate:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitBindBase(ctx *parser.BindBaseContext) interface{} {
	fmt.Println("bind-base:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLogicalNegate(ctx *parser.LogicalNegateContext) interface{} {
	fmt.Println("logical-negate:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAssert(ctx *parser.AssertContext) interface{} {
	fmt.Println("assert:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitArgument(ctx *parser.ArgumentContext) interface{} {
	fmt.Println("argument:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSlurpSpread(ctx *parser.SlurpSpreadContext) interface{} {
	fmt.Println("slurp-spread:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitBaseClone(ctx *parser.BaseCloneContext) interface{} {
	fmt.Println("base-clone:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNumericalPosit(ctx *parser.NumericalPositContext) interface{} {
	fmt.Println("numerical-posit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLogicalPosit(ctx *parser.LogicalPositContext) interface{} {
	fmt.Println("logical-posit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitOverlay(ctx *parser.OverlayContext) interface{} {
	fmt.Println("overlay:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitExistentialPosit(ctx *parser.ExistentialPositContext) interface{} {
	fmt.Println("existential-posit:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutableDestruct(ctx *parser.ImmutableDestructContext) interface{} {
	fmt.Println("immutable-destruct:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMutableDestruct(ctx *parser.MutableDestructContext) interface{} {
	fmt.Println("mutable-destruct:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNestedLabel(ctx *parser.NestedLabelContext) interface{} {
	fmt.Println("nested-label:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNestedOverlay(ctx *parser.NestedOverlayContext) interface{} {
	fmt.Println("nested-overlay:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRange(ctx *parser.RangeContext) interface{} {
	fmt.Println("range:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitString(ctx *parser.StringContext) interface{} {
	fmt.Println("string:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLabelInterp(ctx *parser.LabelInterpContext) interface{} {
	fmt.Println("label-interp!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRoutineInterp(ctx *parser.RoutineInterpContext) interface{} {
	fmt.Println("routine-interp:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSelectorInterp(ctx *parser.SelectorInterpContext) interface{} {
	fmt.Println("selector-interp:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNamedRef(ctx *parser.NamedRefContext) interface{} {
	fmt.Println("named-ref:", ctx.GetText())
	return ctx.Label().GetText()
}

func (v *RhumbVisitor) VisitFunctionRef(ctx *parser.FunctionRefContext) interface{} {
	fmt.Println("function-ref:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitComputedRef(ctx *parser.ComputedRefContext) interface{} {
	fmt.Println("computed-ref:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitJunctionRef(ctx *parser.JunctionRefContext) interface{} {
	fmt.Println("junction-ref:", ctx.GetText())
	return v.VisitChildren(ctx)
}
