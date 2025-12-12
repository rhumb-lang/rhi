package ast

import "fmt"

// ConstraintType defines the specificity of the version requirement
type ConstraintType int

const (
	ConstraintAny     ConstraintType = iota // "-"
	ConstraintMajor                         // "1" or "1.-"
	ConstraintPartial                       // "1.2" or "1.2.-"
	ConstraintExact                         // "1.2.0"
)

// VersionConstraint holds the parsed rules
type VersionConstraint struct {
	Type  ConstraintType
	Major int
	Minor int
	Patch int

	// Original raw value for debugging/formatting
	Raw string
}

// LibraryLiteral is the AST node for importing a library
type LibraryLiteral struct {
	Resolver   string // "!", "-", "git"
	Path       string // "math", "src/core"
	Constraint VersionConstraint

	// Metadata for warnings (e.g., using Float 1.2 instead of 1.2.-)
	Warnings []string
}

func (l *LibraryLiteral) expressionNode() {}
func (l *LibraryLiteral) String() string {
	return fmt.Sprintf("{%s|%s|%s}", l.Resolver, l.Path, l.Constraint.Raw)
}
