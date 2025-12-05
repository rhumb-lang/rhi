package ast

import "strings"

// FieldDefinition represents a named field in a map (e.g., `x .. 1`).
type FieldDefinition struct {
	Key       Node // Literal or Reference
	Value     Expression
	IsMutable bool // true for ::, false for ..
	IsSub     bool // true if @ present
}

func (f *FieldDefinition) fieldNode() {}
func (f *FieldDefinition) String() string {
	var sb strings.Builder
	if f.IsSub {
		sb.WriteString("@")
	}
	sb.WriteString(f.Key.String())
	if f.IsMutable {
		sb.WriteString(" :: ")
	} else {
		sb.WriteString(" .. ")
	}
	sb.WriteString(f.Value.String())
	return sb.String()
}

// FieldPun represents a shorthand field (e.g., `.x`).
type FieldPun struct {
	Key       Node
	IsMutable bool // .x vs :x (Wait, grammar has Dot for mut? Check visitor.)
	IsSub     bool
}

func (f *FieldPun) fieldNode() {}
func (f *FieldPun) String() string {
	var sb strings.Builder
	// TODO: Align with grammar semantics in Visitor
	if f.IsMutable {
		sb.WriteString(":") // Placeholder
	} else {
		sb.WriteString(".")
	}
	if f.IsSub {
		sb.WriteString("@")
	}
	sb.WriteString(f.Key.String())
	return sb.String()
}

// FieldSpread represents `&expr`.
type FieldSpread struct {
	Key Node
}

func (f *FieldSpread) fieldNode() {}
func (f *FieldSpread) String() string {
	return "&" + f.Key.String()
}

// FieldElement represents a positional element (just an expression).
type FieldElement struct {
	Value Expression
}

func (f *FieldElement) fieldNode() {}
func (f *FieldElement) String() string {
	return f.Value.String()
}

// PatternDefinition represents `expr .. expr` or `expr :: expr`.
type PatternDefinition struct {
	Target    Expression
	Action    Expression
	IsConsume bool // true for .. (Consume/Stop), false for :: (Peek/Fallthrough)
}

func (p *PatternDefinition) patternNode() {}
func (p *PatternDefinition) String() string {
	op := " :: "
	if p.IsConsume {
		op = " .. "
	}
	return p.Target.String() + op + p.Action.String()
}

// PatternDefault represents just `expr`.
type PatternDefault struct {
	Value Expression
}

func (p *PatternDefault) patternNode() {}
func (p *PatternDefault) String() string {
	return p.Value.String()
}
