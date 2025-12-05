package generator

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"

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

type RhumbVisitor struct {
	P.BaseRhumbParserVisitor
	VM *virtual.VirtualMachine
}

func NewRhumbVisitor(vm *virtual.VirtualMachine) *RhumbVisitor {
	visitor := new(RhumbVisitor)
	visitor.VM = vm
	return visitor
}

func (v *RhumbVisitor) Visit(tree antlr.ParseTree) (any, error) {

	return v.visit(tree)

}

func (v *RhumbVisitor) visit(tree antlr.ParseTree) (any, error) {

	viLogger.Printf("Visit[tree type: %s]\n", reflect.TypeOf(tree))

	switch t := tree.(type) {

	case *antlr.ErrorNodeImpl:

		return nil, fmt.Errorf("syntax error near '%s'", t.GetText())

	default:

		result := tree.Accept(v)

		if err, ok := result.(error); ok {

			return nil, err

		}

		return result, nil

	}

}

func (v *RhumbVisitor) VisitChildren(

	node antlr.RuleNode,

) interface{} {

	res, err := v.visitChildren(node)

	if err != nil {

		return err

	}

	return res

}

func (v *RhumbVisitor) visitChildren(

	node antlr.RuleNode,

) (any, error) {

	for _, n := range node.GetChildren() {

		viLogger.Printf(

			"VisitChildren[node type: %s]\n",

			reflect.TypeOf(n),
		)

		if _, err := v.visit(n.(antlr.ParseTree)); err != nil {

			return nil, err

		}

	}

	return nil, nil

}

// Visit a parse tree produced by RhumbParser#fields.

func (v *RhumbVisitor) VisitFields(ctx *P.FieldsContext) (any, error) {

	viLogger.Println("Fields!")

	return v.visitChildren(ctx)

}

// Visit a parse tree produced by RhumbParser#patterns.

func (v *RhumbVisitor) VisitPatterns(

	ctx *P.PatternsContext,

) (any, error) {

	viLogger.Println("Patterns!")

	viLogger.Println(ctx.GetText())

	return v.visitChildren(ctx)

}

// Visit a parse tree produced by RhumbParser#terminator.

func (v *RhumbVisitor) VisitTerminator(

	ctx *P.TerminatorContext,

) (any, error) {

	viLogger.Println("Terminator!")

	viLogger.Println(ctx.GetText())

	return v.visitChildren(ctx)

}

func (v *RhumbVisitor) visitBinaryLeaf(leaf P.IExpressionContext) {

	res, err := v.visit(leaf)

	if err != nil {

		panic(err)

	}

	switch leafTyped := res.(type) {

	case *P.LabelSymbolContext:

		labelID := v.VM.RegisterObject(

			object.NewLabel(v.VM.Memory, leafTyped.GetText()),
		)

		start := leafTyped.GetStart()

		v.VM.Write(code.NewLocal(

			start.GetLine(),

			start.GetColumn(),

			labelID,
		))

	default:

		v.visit(leaf)

	}

}

func (v *RhumbVisitor) visitMapChildren(
	children []antlr.Tree,
	kind string,
	parentCtx antlr.ParserRuleContext,
) {
	var (
		s, e antlr.Token
		op   int
	)
	viLogger.Println(fmt.Sprint(kind, "Map!"))
	op = v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, "_[[_"),
	)
	s = parentCtx.GetStart()
	switch kind {
	case "List":
		v.VM.Write(code.NewLocal(s.GetLine(), s.GetColumn(), op))
	case "Invoke":
		v.VM.Write(code.NewInner(s.GetLine(), s.GetColumn(), op))
	default:
		panic("not yet implemented")
	}

	op = v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, "_[>_"),
	)

	for _, n := range children {
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
			s = parentCtx.GetStart()
			v.VM.Write(code.NewLocal(s.GetLine(), s.GetColumn(), op))
		}
	}

	op = v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, "_]]_"),
	)
	e = parentCtx.GetStop()
	switch kind {
	case "List":
		v.VM.Write(code.NewLocal(e.GetLine(), e.GetColumn(), op))
	case "Invoke":
		v.VM.Write(code.NewInner(e.GetLine(), e.GetColumn(), op))
	default:
		panic("not yet implemented")
	}
}
