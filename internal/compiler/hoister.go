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
		// Recurse (e.g. nested assignments? or just flow)
		h.visit(n.Left)
		h.visit(n.Right)
	
	case *ast.CallExpression:
		h.visit(n.Callee)
		for _, arg := range n.Args {
			h.visit(arg)
		}
		
	case *ast.SelectorExpression:
		for _, pat := range n.Patterns {
			if p, ok := pat.(*ast.PatternDefinition); ok {
				// Check Target for bindings (LabelLiteral)
				if label, ok := p.Target.(*ast.LabelLiteral); ok {
					h.add(label.Value)
				}
				h.visit(p.Action)
			} else if def, ok := pat.(*ast.PatternDefault); ok {
				h.visit(def.Value)
			}
		}
	// For now, we only hoist from current scope boundaries.
	// Does `->` create new scope? Yes. `Hoister` should NOT descend into `OpMakeFn`'s Routine.
	// But `hoist` is called on the Body of the function by the Child Compiler.
	// So we just need to STOP recursing at function boundaries.
	}
}
