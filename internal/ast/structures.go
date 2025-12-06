package ast

import "strings"

// MapExpression represents a map definition [key: val; ...].
type MapExpression struct {
	Fields []Field
}

func (m *MapExpression) expressionNode() {}
func (m *MapExpression) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, f := range m.Fields {
		if i > 0 {
			sb.WriteString("; ")
		}
		sb.WriteString(f.String())
	}
	sb.WriteString("]")
	return sb.String()
}

// CallExpression represents a function call: Callee(Args).
type CallExpression struct {
	Callee Expression
	Args   []Expression
}

func (c *CallExpression) expressionNode() {}
func (c *CallExpression) String() string {
	var sb strings.Builder
	sb.WriteString(c.Callee.String())
	sb.WriteString("(")
	for i, arg := range c.Args {
		if i > 0 {
			sb.WriteString("; ")
		}
		sb.WriteString(arg.String())
	}
	sb.WriteString(")")
	return sb.String()
}

// RoutineExpression represents a subroutine ( ... ).
type RoutineExpression struct {
	Expressions []Expression
}

func (r *RoutineExpression) expressionNode() {}
func (r *RoutineExpression) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	for i, e := range r.Expressions {
		if i > 0 {
			sb.WriteString("; ")
		}
		sb.WriteString(e.String())
	}
	sb.WriteString(")")
	return sb.String()
}

// SelectorExpression represents a selector block { pattern .. action; ... }.
type SelectorExpression struct {
	Patterns []Pattern
}

func (s *SelectorExpression) expressionNode() {}
func (s *SelectorExpression) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	for i, p := range s.Patterns {
		if i > 0 {
			sb.WriteString("; ")
		}
		sb.WriteString(p.String())
	}
	sb.WriteString("}")
	return sb.String()
}

// ChildRealmLiteral represents <$>.
type ChildRealmLiteral struct{}

func (r *ChildRealmLiteral) expressionNode() {}
func (r *ChildRealmLiteral) String() string  { return "<$>" }

// DetachedRealmLiteral represents <|>.
type DetachedRealmLiteral struct{}

func (r *DetachedRealmLiteral) expressionNode() {}
func (r *DetachedRealmLiteral) String() string  { return "<|>" }

// LibraryLiteral represents a library import { name | "path" | 1.0.0 }.
type LibraryLiteral struct {
	Alias   string // Optional
	Path    *TextLiteral
	Version string // Keeping simpler for now
}

func (l *LibraryLiteral) expressionNode() {}
func (l *LibraryLiteral) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	if l.Alias != "" {
		sb.WriteString(l.Alias + " | ")
	}
	sb.WriteString(l.Path.String() + " | " + l.Version)
	sb.WriteString("}")
	return sb.String()
}

// NamedReference represents <Label>.
type NamedReference struct {
	Label string
}

func (r *NamedReference) expressionNode() {}
func (r *NamedReference) String() string  { return "<" + r.Label + ">" }

// ComputedTextReference represents <"text">.
type ComputedTextReference struct {
	Text *TextLiteral
}

func (r *ComputedTextReference) expressionNode() {}
func (r *ComputedTextReference) String() string  { return "<" + r.Text.String() + ">" }

// RoutineReference represents <( ... )>.
type RoutineReference struct {
	Routine *RoutineExpression
}

func (r *RoutineReference) expressionNode() {}
func (r *RoutineReference) String() string  { return "<" + r.Routine.String() + ">" }

// ComputedExpressionReference represents <[ ... ]>.
type ComputedExpressionReference struct {
	Expression Expression
}

func (r *ComputedExpressionReference) expressionNode() {}
func (r *ComputedExpressionReference) String() string  { return "<" + r.Expression.String() + ">" }

// VassalReference represents <{ ... }>.
type VassalReference struct {
	Selector *SelectorExpression
}

func (r *VassalReference) expressionNode() {}
func (r *VassalReference) String() string  { return "<" + r.Selector.String() + ">" }