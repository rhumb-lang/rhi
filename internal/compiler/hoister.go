package compiler

import (
	"slices"

	"github.com/rhumb-lang/rhi/internal/ast"
)

// Hoister finds all variable declarations in a block.
type Hoister struct {
	Locals []string
}

func NewHoister() *Hoister {
	return &Hoister{
		Locals: make([]string, 0),
	}
}

func (h *Hoister) Hoist(node ast.Node) []string {
	// If the node itself is a RoutineExpression (function body),
	// we want to hoist its contents, but NOT recurse into nested routines found within.
	if routine, ok := node.(*ast.RoutineExpression); ok {
		for _, expr := range routine.Expressions {
			h.visit(expr)
		}
	} else {
		h.visit(node)
	}
	return h.Locals
}

func (h *Hoister) add(name string) {
	if slices.Contains(h.Locals, name) {
		return
	}
	h.Locals = append(h.Locals, name)
}

func (h *Hoister) visit(node ast.Node) {
	switch n := node.(type) {
	case *ast.Document:
		for _, expr := range n.Expressions {
			h.visit(expr)
		}
	case *ast.AssertionWrapper:
		h.visit(n.Actual)
		h.visit(n.Expected)
	case *ast.InspectionWrapper:
		h.visit(n.Expr)
	case *ast.UnaryExpression:
		h.visit(n.Expr)
	case *ast.EffectExpression:
		h.visit(n.Target)
		// Selector has its own scope?
		// Selector patterns bind new variables, they don't declare in outer scope.
		// h.visit(n.Selector) // No.
	case *ast.ChainExpression:
		h.visit(n.Base)
		// Steps are just identifiers.
	case *ast.BinaryExpression:
		// 1. Handle Assignments (Hoist the LHS)
		if n.Op == ast.OpAssignImm || n.Op == ast.OpAssignMut || n.Op == ast.OpDestruct {
			if label, ok := n.Left.(*ast.LabelLiteral); ok {
				h.add(label.Value)
			}
			// Destructuring
			if mapExpr, ok := n.Left.(*ast.MapExpression); ok {
				for _, f := range mapExpr.Fields {
					if pun, ok := f.(*ast.FieldPun); ok {
						if label, ok := pun.Key.(*ast.LabelLiteral); ok {
							h.add(label.Value)
						}
					}
				}
			}
		}

		// 2. Recurse (BUT STOP at Function Boundaries)
		h.visit(n.Left)

		// CRITICAL FIX: Do not hoist from the body of a function definition.
		// The compiler handles function bodies with a NEW Hoister.
		if n.Op == ast.OpMakeFn || n.Op == ast.OpBindFn || n.Op == ast.OpLetFn {
			return
		}

		h.visit(n.Right)

	case *ast.CallExpression:
		h.visit(n.Callee)
		for _, arg := range n.Args {
			h.visit(arg)
		}
	case *ast.RoutineExpression:
		// CRITICAL FIX: Recurse into routines!
		// Previously this returned immediately, blocking visibility into blocks like ( a := 1 ).
		for _, expr := range n.Expressions {
			h.visit(expr)
		}
	case *ast.MapExpression:
		for _, f := range n.Fields {
			switch field := f.(type) {
			case *ast.FieldDefinition:
				h.visit(field.Value)
			case *ast.FieldElement:
				h.visit(field.Value)
			}
		}
	}
}
