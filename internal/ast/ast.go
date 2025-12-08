package ast

import "strings"

// Node is the base interface for all AST nodes.
type Node interface {
	String() string
}

// Expression represents a value-producing node.
type Expression interface {
	Node
	expressionNode()
}

// Field represents a key-value pair or definition in a map.
type Field interface {
	Node
	fieldNode()
}

// Pattern represents a matching rule in a selector.
type Pattern interface {
	Node
	patternNode()
}

// Document represents the root of a parsed file.
type Document struct {
	Expressions []Expression
}

func (d *Document) String() string {
	var sb strings.Builder
	for _, expr := range d.Expressions {
		sb.WriteString(expr.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

// AssertionWrapper wraps an expression that has an attached test assertion.
type AssertionWrapper struct {
	Actual   Expression
	Expected Expression
}

func (a *AssertionWrapper) expressionNode() {}
func (a *AssertionWrapper) String() string {
	return a.Actual.String() + " \"%=\" " + a.Expected.String()
}