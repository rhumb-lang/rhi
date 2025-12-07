package ast

import (
	"fmt"
	"strings"
)

// EmptyLiteral represents '___' (empty).
type EmptyLiteral struct{}

func (l *EmptyLiteral) expressionNode() {}
func (l *EmptyLiteral) String() string  { return "___" }

// LabelLiteral represents a simple identifier.
type LabelLiteral struct {
	Value string
}

func (l *LabelLiteral) expressionNode() {}
func (l *LabelLiteral) String() string  { return l.Value }

// IntegerLiteral represents whole numbers.
type IntegerLiteral struct {
	Value int64
}

func (l *IntegerLiteral) expressionNode() {}
func (l *IntegerLiteral) String() string  { return fmt.Sprintf("%d", l.Value) }

// RationalLiteral represents floating point numbers.
type RationalLiteral struct {
	Value float64
}

func (l *RationalLiteral) expressionNode() {}
func (l *RationalLiteral) String() string  { return fmt.Sprintf("%f", l.Value) }

// KeyLiteral represents a unique key like `foo`.
type KeyLiteral struct {
	Value string
}

func (l *KeyLiteral) expressionNode() {}
func (l *KeyLiteral) String() string  { return "`" + l.Value }

// DateLiteral represents a date like 2025/12/05.
type DateLiteral struct {
	Year  string // Keeping as string to support dynamic values if AST allows, though grammar suggests parts
	Month string
	Day   string
}

func (l *DateLiteral) expressionNode() {}
func (l *DateLiteral) String() string {
	return fmt.Sprintf("%s/%s/%s", l.Year, l.Month, l.Day)
}

// TextSegment is a part of an interpolated string.
type TextSegment interface {
	Node
	textSegmentNode()
}

// StringSegment is a literal string part of a text.
type StringSegment struct {
	Value string
}

func (s *StringSegment) textSegmentNode() {}
func (s *StringSegment) String() string   { return s.Value }

// InterpSegment is an interpolated expression part of a text.
type InterpSegment struct {
	Expression Expression
}

func (s *InterpSegment) textSegmentNode() {}
func (s *InterpSegment) String() string   { return "$(" + s.Expression.String() + ")" }

// TextLiteral represents a string, potentially with interpolation.
type TextLiteral struct {
	Segments []TextSegment
	IsRaw    bool
}

func (l *TextLiteral) expressionNode() {}
func (l *TextLiteral) String() string {
	if l.IsRaw {
		var sb strings.Builder
		sb.WriteString("'")
		for _, s := range l.Segments {
			sb.WriteString(s.String())
		}
		sb.WriteString("'")
		return sb.String()
	}
	var sb strings.Builder
	sb.WriteString("\"")
	for _, s := range l.Segments {
		sb.WriteString(s.String())
	}
	sb.WriteString("\"")
	return sb.String()
}
