package ast

import "fmt"

// ConstraintType defines the specificity of the version requirement
type ConstraintType int

const (
	ConstraintMajor ConstraintType = iota // "-"
	ConstraintMinor                       // "1" or "1.-"
	ConstraintPatch                       // "1.2" or "1.2.-"
	ConstraintExact                       // "1.2.0"
)

// VersionConstraint is the semantic representation of the requirement
type VersionConstraint struct {
	Type       ConstraintType
	Major      uint16
	Minor      uint16
	Patch      uint32
	PreRelease string // e.g., "alpha"
	Raw        string // For display/debugging
}

func (vc VersionConstraint) Matches(vMajor, vMinor uint16, vPatch uint32) bool {
	switch vc.Type {
	case ConstraintMajor:
		return true
	case ConstraintMinor:
		return vMajor == vc.Major
	case ConstraintPatch:
		return vMajor == vc.Major && vMinor == vc.Minor
	case ConstraintExact:
		return vMajor == vc.Major && vMinor == vc.Minor && vPatch == vc.Patch
	default:
		return false
	}
}

type LibraryLiteral struct {
	Resolver   string // "!", "-", "git"
	Path       string
	Constraint VersionConstraint
	Warnings   []string // Store warnings generated during parsing
}

func (l *LibraryLiteral) expressionNode() {}
func (l *LibraryLiteral) String() string {
	return fmt.Sprintf("{%s|%s|%s}", l.Resolver, l.Path, l.Constraint.Raw)
}
