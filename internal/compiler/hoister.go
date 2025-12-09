package compiler

import (
	"slices"

	"git.sr.ht/~madcapjake/rhi/internal/ast"
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
	case *ast.BinaryExpression:
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
		// Recurse
		h.visit(n.Left)
		h.visit(n.Right)
	case *ast.CallExpression:
		h.visit(n.Callee)
		for _, arg := range n.Args {
			h.visit(arg)
		}
	case *ast.RoutineExpression:
		// Don't hoist inside routines (new scope)
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
