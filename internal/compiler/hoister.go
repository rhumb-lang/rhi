package compiler

import (
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
	h.visit(node)
	return h.Locals
}

func (h *Hoister) add(name string) {
	for _, l := range h.Locals {
		if l == name {
			return // Already hoisted
		}
	}
	h.Locals = append(h.Locals, name)
}

func (h *Hoister) visit(node ast.Node) {
	if node == nil {
		return
	}
	
	switch n := node.(type) {
	case *ast.Document:
		for _, expr := range n.Expressions {
			h.visit(expr)
		}
	case *ast.RoutineExpression:
		for _, expr := range n.Expressions {
			h.visit(expr)
		}
	case *ast.BinaryExpression:
		// Check for assignment
		if n.Op == ast.OpAssignImm || n.Op == ast.OpAssignMut {
			if label, ok := n.Left.(*ast.LabelLiteral); ok {
				h.add(label.Value)
			}
		}
		
		// Do NOT recurse into Function Body (Scope Boundary)
		if n.Op == ast.OpMakeFn {
			h.visit(n.Left) // Visit params
			return
		}
		
		h.visit(n.Left)
		h.visit(n.Right)
	
	case *ast.CallExpression:
		h.visit(n.Callee)
		for _, arg := range n.Args {
			h.visit(arg)
		}
		
	// Stop at Selector boundary (creates new scope)
	case *ast.SelectorExpression:
		return
	}
}
