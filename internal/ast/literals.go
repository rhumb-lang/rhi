package ast

import (
	"fmt"
	"strings"

	"github.com/cockroachdb/apd/v3"
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

// DateTimeLiteral represents a date/time as Unix Milliseconds.
type DateTimeLiteral struct {
	Value int64
}

func (l *DateTimeLiteral) expressionNode() {}
func (l *DateTimeLiteral) String() string {
	return fmt.Sprintf("DateTime<%d>", l.Value)
}

// VersionLiteral represents a semantic version.
type VersionLiteral struct {
	Major, Minor uint16
	Patch        uint32
	Suffix       string
	IsWildcard   bool
}

func (l *VersionLiteral) expressionNode() {}
func (l *VersionLiteral) String() string {
	s := fmt.Sprintf("%d.%d.%d", l.Major, l.Minor, l.Patch)
	if l.Suffix != "" {
		s += l.Suffix
	}
	if l.IsWildcard {
		s += ".-"
	}
	return s
}

// DecimalLiteral represents an arbitrary precision decimal.
type DecimalLiteral struct {
	Value *apd.Decimal
}

func (l *DecimalLiteral) expressionNode() {}
func (l *DecimalLiteral) String() string {
	s := "00.0"

	if l.Value != nil {
		s = l.Value.Text('f')
	}

	if s[0] != '0' {
		var sb strings.Builder
		sb.WriteRune('0')
		sb.WriteString(s)
		s = sb.String()
	}

	return s
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
