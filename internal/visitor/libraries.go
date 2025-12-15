package visitor

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/grammar"
)

// Grammar: OpenCurly libraryResolver Pipe libraryPath+ Pipe (Version | ...) CloseCurly
func (b *ASTBuilder) VisitLibrary(ctx *grammar.LibraryContext) interface{} {
	lib := &ast.LibraryLiteral{}

	lib.Resolver = ctx.LibraryResolver().GetText()

	// libraryPath+ can be multiple tokens (folders).
	var pathParts []string
	for _, p := range ctx.AllLibraryPath() {
		pathParts = append(pathParts, p.GetText())
	}
	lib.Path = strings.Join(pathParts, "/") // Canonicalize separator

	lib.Constraint, lib.Warnings = b.resolveVersionValue(ctx)

	return lib
}

// resolveVersionValue converts the polymorphic library version tokens into a constraint
func (b *ASTBuilder) resolveVersionValue(ctx *grammar.LibraryContext) (ast.VersionConstraint, []string) {
	var warnings []string

	// Case 1: Explicit Version Token (e.g., 1.0.0, 1.2.-)
	if vToken := ctx.Version(); vToken != nil {
		vLit := b.parseVersionString(vToken.GetText())
		return b.constraintFromVersionLiteral(vLit), warnings
	}

	// Case 2: Number Rule (e.g., 1, 1.-)
	// We visit the number rule which handles Integer vs VersionLiteral logic
	if nCtx := ctx.Number(); nCtx != nil {
		res := b.Visit(nCtx) // Calls VisitWholeNumber/VisitNumber

		switch val := res.(type) {
		case *ast.IntegerLiteral:
			// "1" -> Treat as Major Wildcard (1.-.-)
			return ast.VersionConstraint{
				Type:  ast.ConstraintMajor,
				Major: uint16(val.Value),
				Raw:   fmt.Sprintf("%d", val.Value),
			}, warnings

		case *ast.VersionLiteral:
			// "1.-" -> Parsed as VersionLiteral by VisitWholeNumber
			return b.constraintFromVersionLiteral(val), warnings
		}
	}

	// Case 3: Floating Point (e.g., 1.2)
	// Treated as Partial Wildcard but triggers a warning
	if fToken := ctx.FloatingPoint(); fToken != nil {
		text := fToken.GetText()
		// Normalize 1,2 -> 1.2
		normText := strings.Replace(text, ",", ".", 1)
		val, _ := strconv.ParseFloat(normText, 64)

		major := int(val)

		// Best-effort extraction of "minor" from float string
		// 1.2 -> minor 2.  1.02 -> minor 2? (Ambiguous, hence warning)
		parts := strings.Split(normText, ".")
		minor := 0
		if len(parts) > 1 {
			// Naive parse of fraction part
			m, _ := strconv.Atoi(parts[1])
			minor = m
		}

		warnings = append(warnings, fmt.Sprintf(
			"Using Rational Number '%s' as a version can be misunderstood. Did you mean '%d.%d.-'?",
			text, major, minor,
		))

		return ast.VersionConstraint{
			Type:  ast.ConstraintPatch,
			Major: uint16(major),
			Minor: uint16(minor),
			Raw:   text,
		}, warnings
	}

	// Case 4: Dash (Any Version)
	if ctx.Dash() != nil {
		return ast.VersionConstraint{
			Type: ast.ConstraintMajor,
			Raw:  "-",
		}, warnings
	}

	// Case 5: Label (Variable Reference)
	if lToken := ctx.Label(); lToken != nil {
		// e.g. {! | math | stable}
		// We treat this as Exact for now, or you could add a ConstraintVariable type.
		// The runtime resolver will look up the variable 'stable'.
		return ast.VersionConstraint{
			Type: ast.ConstraintExact,
			Raw:  lToken.GetText(),
		}, warnings
	}

	panic("should be unreachable if grammar is correct")
}

// Helper to map AST Literal -> Semantic Constraint
func (b *ASTBuilder) constraintFromVersionLiteral(v *ast.VersionLiteral) ast.VersionConstraint {
	c := ast.VersionConstraint{
		Major:      v.Major,
		Minor:      v.Minor,
		Patch:      v.Patch,
		PreRelease: v.Suffix,
		Raw:        v.String(),
	}

	if v.IsWildcard {
		if v.Minor == 0xFFFF {
			c.Type = ast.ConstraintMinor // 1.-
		} else if v.Patch == 0xFFFFFFFF {
			c.Type = ast.ConstraintPatch // 1.2.-
		} else {
			// Fallback if somehow wildcard is set but no sentinel
			c.Type = ast.ConstraintExact
		}
	} else {
		c.Type = ast.ConstraintExact
	}
	return c
}

// Helper to map AST Literal to Semantic Constraint
func constraintFromVersionLiteral(v *ast.VersionLiteral) ast.VersionConstraint {
	c := ast.VersionConstraint{
		Major:      v.Major,
		Minor:      v.Minor,
		Patch:      v.Patch,
		PreRelease: v.Suffix, // Ensure VersionLiteral has Suffix field
		Raw:        v.String(),
	}

	if v.IsWildcard {
		// Determine level based on what is missing/wild
		if v.Minor == 0xFFFF {
			c.Type = ast.ConstraintMinor // 1.-
		} else if v.Patch == 0xFFFFFFFF {
			c.Type = ast.ConstraintPatch // 1.2.-
		} else {
			c.Type = ast.ConstraintExact // Should not happen if IsWildcard is true
		}
	} else {
		c.Type = ast.ConstraintExact
	}
	return c
}
