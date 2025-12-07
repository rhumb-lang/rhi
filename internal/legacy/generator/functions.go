package generator

import (
	"fmt"

	"git.sr.ht/~madcapjake/rhi/internal/code"
	"git.sr.ht/~madcapjake/rhi/internal/object"
	P "git.sr.ht/~madcapjake/rhi/internal/parser"
)

// Visit a parse tree produced by RhumbParser#routine.
func (v *RhumbVisitor) VisitRoutine(
	ctx *P.RoutineContext,
) (any, error) {
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

	return routine, nil
}

// Visit a parse tree produced by RhumbParser#invocation.
func (v *RhumbVisitor) VisitInvocation(
	ctx *P.InvocationContext,
) (any, error) {
	viLogger.Println("Invocation!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#function.
func (v *RhumbVisitor) VisitFunction(ctx *P.FunctionContext) (any, error) {
	viLogger.Println("Function!")
	viLogger.Println(ctx.GetText())
	label := v.VM.RegisterObject(
		object.NewLabel(v.VM.Memory, fmt.Sprint("_", ctx.GetText(), "_")),
	)
	s := ctx.GetStart()
	v.VM.Write(code.NewLocal(s.GetLine(), s.GetColumn(), label))
	return nil, nil
}

// Visit a parse tree produced by RhumbParser#method.
func (v *RhumbVisitor) VisitMethod(ctx *P.MethodContext) (any, error) {
	viLogger.Println("Method!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#parameters.
func (v *RhumbVisitor) VisitParameters(ctx *P.ParametersContext) (any, error) {
	viLogger.Println("Parameters!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}

// Visit a parse tree produced by RhumbParser#functionRef.
func (v *RhumbVisitor) VisitFunctionRef(ctx *P.FunctionRefContext) (any, error) {
	viLogger.Println("FunctionRef!")
	viLogger.Println(ctx.GetText())
	return v.visitChildren(ctx)
}
