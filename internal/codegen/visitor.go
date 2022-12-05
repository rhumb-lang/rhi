package codegen

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"

	"git.sr.ht/~madcapjake/grhumb/internal/parser"
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
}

func NewRhumbVisitor() *RhumbVisitor {
	return new(RhumbVisitor)
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
		fmt.Println("child:", reflect.TypeOf(n))
		v.Visit(n.(antlr.ParseTree))
	}
	return nil
}

func (v *RhumbVisitor) VisitSequence(ctx *parser.SequenceContext) interface{} {
	fmt.Println("sequence!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSimple(ctx *parser.SimpleContext) interface{} {
	fmt.Println("simple!")
	fmt.Println(ctx.GetText())
	return v.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (v *RhumbVisitor) VisitMutableLabel(ctx *parser.MutableLabelContext) interface{} {
	fmt.Println("mutable label!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLabelLiteral(ctx *parser.LabelLiteralContext) interface{} {
	fmt.Println("label!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	fmt.Println("assignment!")
	// fmt.Println(ctx.GetText())
	label := ctx.Label()
	fmt.Println("Label:", label)
	ret := v.Visit(ctx.Expression().GetRuleContext()).(RhumbReturn)
	fmt.Println("Expr:", ret.Value)
	return nil
	// return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitFloatLiteral(ctx *parser.FloatLiteralContext) interface{} {
	fmt.Println("float-lit!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIntegerLiteral(ctx *parser.IntegerLiteralContext) interface{} {
	fmt.Println("int-lit!")
	val, err := strconv.ParseInt(ctx.GetText(), 10, 32)
	if err != nil {
		return RhumbReturn{nil, fmt.Errorf("unable to parse int")}
	}
	fmt.Println(val)
	return RhumbReturn{val, nil}
}

func (v *RhumbVisitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	fmt.Println("string-lit!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitReferenceLiteral(ctx *parser.ReferenceLiteralContext) interface{} {
	fmt.Println("ref-lit!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConjunctive(ctx *parser.ConjunctiveContext) interface{} {
	fmt.Println("conj!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAccess(ctx *parser.AccessContext) interface{} {
	fmt.Println("access!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitApplicative(ctx *parser.ApplicativeContext) interface{} {
	fmt.Println("applicative!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConditional(ctx *parser.ConditionalContext) interface{} {
	fmt.Println("cond!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitPrefix(ctx *parser.PrefixContext) interface{} {
	fmt.Println("prefix!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitComparative(ctx *parser.ComparativeContext) interface{} {
	fmt.Println("compare!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMultiplicative(ctx *parser.MultiplicativeContext) interface{} {
	fmt.Println("multiply!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAdditive(ctx *parser.AdditiveContext) interface{} {
	fmt.Println("add!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitInvocation(ctx *parser.InvocationContext) interface{} {
	fmt.Println("invoke!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRoutine(ctx *parser.RoutineContext) interface{} {
	fmt.Println("routine!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDisjunctive(ctx *parser.DisjunctiveContext) interface{} {
	fmt.Println("disjunct!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIdentity(ctx *parser.IdentityContext) interface{} {
	fmt.Println("identify!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitEffect(ctx *parser.EffectContext) interface{} {
	fmt.Println("effect!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMember(ctx *parser.MemberContext) interface{} {
	fmt.Println("member!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSelector(ctx *parser.SelectorContext) interface{} {
	fmt.Println("selector!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitPower(ctx *parser.PowerContext) interface{} {
	fmt.Println("power!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMap(ctx *parser.MapContext) interface{} {
	fmt.Println("map!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitFreeze(ctx *parser.FreezeContext) interface{} {
	fmt.Println("freeze!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitInner(ctx *parser.InnerContext) interface{} {
	fmt.Println("inner!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLength(ctx *parser.LengthContext) interface{} {
	fmt.Println("length!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	fmt.Println("function!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitJunction(ctx *parser.JunctionContext) interface{} {
	fmt.Println("junction!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitGreaterThan(ctx *parser.GreaterThanContext) interface{} {
	fmt.Println("greater-than!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitGreaterThanOrEqualTo(ctx *parser.GreaterThanOrEqualToContext) interface{} {
	fmt.Println("greater-than-or-equal-to!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLessThan(ctx *parser.LessThanContext) interface{} {
	fmt.Println("less-than!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLessThanOrEqualTo(ctx *parser.LessThanOrEqualToContext) interface{} {
	fmt.Println("less-than-or-equal-to!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsEqual(ctx *parser.IsEqualContext) interface{} {
	fmt.Println("is-equal!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsLike(ctx *parser.IsLikeContext) interface{} {
	fmt.Println("is-like!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsIn(ctx *parser.IsInContext) interface{} {
	fmt.Println("is-in!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsOverlayed(ctx *parser.IsOverlayedContext) interface{} {
	fmt.Println("is-overlayed!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIsTopmost(ctx *parser.IsTopmostContext) interface{} {
	fmt.Println("is-topmost!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotEqual(ctx *parser.NotEqualContext) interface{} {
	fmt.Println("not-equal!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotLike(ctx *parser.NotLikeContext) interface{} {
	fmt.Println("not-like!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotIn(ctx *parser.NotInContext) interface{} {
	fmt.Println("not-in!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotOverlayed(ctx *parser.NotOverlayedContext) interface{} {
	fmt.Println("not-overlayed!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNotTopmost(ctx *parser.NotTopmostContext) interface{} {
	fmt.Println("not-topmost!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitConjunctiveOp(ctx *parser.ConjunctiveOpContext) interface{} {
	fmt.Println("conj-op!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDisjunctiveOp(ctx *parser.DisjunctiveOpContext) interface{} {
	fmt.Println("disjunct-op!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitOtherwise(ctx *parser.OtherwiseContext) interface{} {
	fmt.Println("otherwise!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDefault(ctx *parser.DefaultContext) interface{} {
	fmt.Println("default!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitForeach(ctx *parser.ForeachContext) interface{} {
	fmt.Println("for-each!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitWhile(ctx *parser.WhileContext) interface{} {
	fmt.Println("while!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitThen(ctx *parser.ThenContext) interface{} {
	fmt.Println("then!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitElse(ctx *parser.ElseContext) interface{} {
	fmt.Println("else!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAddition(ctx *parser.AdditionContext) interface{} {
	fmt.Println("addition!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDeviation(ctx *parser.DeviationContext) interface{} {
	fmt.Println("deviation!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSubtraction(ctx *parser.SubtractionContext) interface{} {
	fmt.Println("subtraction!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMultiplication(ctx *parser.MultiplicationContext) interface{} {
	fmt.Println("mupltication!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitDivision(ctx *parser.DivisionContext) interface{} {
	fmt.Println("division!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitIntegerDivision(ctx *parser.IntegerDivisionContext) interface{} {
	fmt.Println("int-division!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitModulo(ctx *parser.ModuloContext) interface{} {
	fmt.Println("modulo!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitBind(ctx *parser.BindContext) interface{} {
	fmt.Println("bind!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitExponent(ctx *parser.ExponentContext) interface{} {
	fmt.Println("exponenet!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRootExtraction(ctx *parser.RootExtractionContext) interface{} {
	fmt.Println("root-extract!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitScientific(ctx *parser.ScientificContext) interface{} {
	fmt.Println("scientific!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutablePair(ctx *parser.ImmutablePairContext) interface{} {
	fmt.Println("immutable-pair!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutableLabel(ctx *parser.ImmutableLabelContext) interface{} {
	fmt.Println("immutable-label!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMutablePair(ctx *parser.MutablePairContext) interface{} {
	fmt.Println("mutable-pair!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNumericalNegate(ctx *parser.NumericalNegateContext) interface{} {
	fmt.Println("numerical-negate!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitBindBase(ctx *parser.BindBaseContext) interface{} {
	fmt.Println("bind-base!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLogicalNegate(ctx *parser.LogicalNegateContext) interface{} {
	fmt.Println("logical-negate!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitAssert(ctx *parser.AssertContext) interface{} {
	fmt.Println("assert!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitArgument(ctx *parser.ArgumentContext) interface{} {
	fmt.Println("argument!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSlurpSpread(ctx *parser.SlurpSpreadContext) interface{} {
	fmt.Println("slurp-spread!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitBaseClone(ctx *parser.BaseCloneContext) interface{} {
	fmt.Println("base-clone!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNumericalPosit(ctx *parser.NumericalPositContext) interface{} {
	fmt.Println("numerical-posit!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLogicalPosit(ctx *parser.LogicalPositContext) interface{} {
	fmt.Println("logical-posit!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitOverlay(ctx *parser.OverlayContext) interface{} {
	fmt.Println("overlay!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitExistentialPosit(ctx *parser.ExistentialPositContext) interface{} {
	fmt.Println("existential-posit!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitImmutableDestruct(ctx *parser.ImmutableDestructContext) interface{} {
	fmt.Println("immutable-destruct!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitMutableDestruct(ctx *parser.MutableDestructContext) interface{} {
	fmt.Println("mutable-destruct!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNestedLabel(ctx *parser.NestedLabelContext) interface{} {
	fmt.Println("nested-label!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNestedOverlay(ctx *parser.NestedOverlayContext) interface{} {
	fmt.Println("nested-overlay!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRange(ctx *parser.RangeContext) interface{} {
	fmt.Println("range!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitString(ctx *parser.StringContext) interface{} {
	fmt.Println("string!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitLabelInterp(ctx *parser.LabelInterpContext) interface{} {
	fmt.Println("label-interp!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitRoutineInterp(ctx *parser.RoutineInterpContext) interface{} {
	fmt.Println("routine-interp!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitSelectorInterp(ctx *parser.SelectorInterpContext) interface{} {
	fmt.Println("selector-interp!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitNamedRef(ctx *parser.NamedRefContext) interface{} {
	fmt.Println("named-ref!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitFunctionRef(ctx *parser.FunctionRefContext) interface{} {
	fmt.Println("function-ref!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitComputedRef(ctx *parser.ComputedRefContext) interface{} {
	fmt.Println("computed-ref!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *RhumbVisitor) VisitJunctionRef(ctx *parser.JunctionRefContext) interface{} {
	fmt.Println("junction-ref!")
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}
