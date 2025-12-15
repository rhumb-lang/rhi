package visitor

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cockroachdb/apd/v3"
	"github.com/rhumb-lang/rhi/internal/ast"
	"github.com/rhumb-lang/rhi/internal/grammar"
)

func (b *ASTBuilder) VisitRationalNumber(ctx *grammar.RationalNumberContext) interface{} {
	text := ctx.GetText()
	// Replace comma with dot for Go parsing if user used comma
	text = strings.Replace(text, ",", ".", 1)

	// Check for Decimal (Leading Zero distinction)
	// e.g. 01.5 (Decimal) vs 1.5 (Float) vs 0.5 (Float)
	// We check the part before the dot.
	parts := strings.Split(text, ".")
	integerPart := parts[0]
	isDecimal := len(integerPart) > 1 && strings.HasPrefix(integerPart, "0")

	if isDecimal {
		// Remove ".-" suffix if present for clean parsing
		cleanText := strings.TrimSuffix(text, ".-")

		d, _, err := apd.NewFromString(cleanText)
		if err != nil {
			panic(fmt.Errorf("error creating apd.Decimal from string: %s", err))
		}

		return &ast.DecimalLiteral{Value: d}
	}

	// Handle 10.- format (10.0) for standard floats
	if strings.HasSuffix(text, ".-") || strings.HasSuffix(text, ",-") {
		text = text[:len(text)-1] + "0"
	}

	val, _ := strconv.ParseFloat(text, 64)
	return &ast.RationalLiteral{Value: val}
}

func (b *ASTBuilder) VisitWholeNumber(ctx *grammar.WholeNumberContext) interface{} {
	text := ctx.GetText()

	// Check for ".-" suffix (Wildcard or Decimal Precision)
	if strings.HasSuffix(text, ".-") || strings.HasSuffix(text, ",-") {
		baseText := text[:len(text)-2]

		// Case 1: Decimal (Leading Zero) -> 01.-
		if len(baseText) > 1 && strings.HasPrefix(baseText, "0") {

			d, _, err := apd.NewFromString(baseText)
			if err != nil {
				panic(fmt.Errorf("error creating apd.Decimal from string: %s", err))
			}
			return &ast.DecimalLiteral{Value: d}
		}

		// Case 2: Version Wildcard -> 1.-
		// We treat this as a VersionLiteral with IsWildcard=true
		// 1.- means 1.*.*. Set Minor and Patch to MAX to indicate wildcard at Minor level (and below).
		val, _ := strconv.ParseUint(baseText, 10, 16)
		return &ast.VersionLiteral{
			Major:      uint16(val),
			Minor:      0xFFFF,     // MaxUint16
			Patch:      0xFFFFFFFF, // MaxUint32
			IsWildcard: true,
		}
	}

	val, _ := strconv.ParseInt(text, 10, 64)
	return &ast.IntegerLiteral{Value: val}
}

func (b *ASTBuilder) VisitNumber(ctx *grammar.NumberContext) interface{} {
	text := ctx.GetText()

	// Check for ".-" suffix (Wildcard or Decimal Precision)
	if strings.HasSuffix(text, ".-") || strings.HasSuffix(text, ",-") {
		baseText := text[:len(text)-2]

		// Case 1: Decimal (Leading Zero) -> 01.-
		if len(baseText) > 1 && strings.HasPrefix(baseText, "0") {

			d, _, err := apd.NewFromString(baseText)
			if err != nil {
				panic(fmt.Errorf("error creating apd.Decimal from string: %s", err))
			}
			return &ast.DecimalLiteral{Value: d}
		}

		// Case 2: Version Wildcard -> 1.-
		val, _ := strconv.ParseUint(baseText, 10, 16)
		return &ast.VersionLiteral{
			Major:      uint16(val),
			Minor:      0xFFFF,     // MaxUint16
			Patch:      0xFFFFFFFF, // MaxUint32
			IsWildcard: true,
		}
	}

	val, _ := strconv.ParseInt(text, 10, 64)
	return &ast.IntegerLiteral{Value: val}
}

// Add this helper method to ASTBuilder (can be in numbers.go)
func (b *ASTBuilder) parseVersionString(text string) *ast.VersionLiteral {
	suffix := ""
	base := text

	// Detect suffix (starts with - or +)
	// Need to be careful not to confuse with ".-" wildcard
	// Regex would be easiest, but manual scan is fine.
	// Version structure is roughly: N.N.N-Suffix
	// But ".-" is valid at the end.
	if idx := strings.IndexAny(text, "-+"); idx != -1 {
		// Check if it is the ".-" wildcard at the end
		if idx == len(text)-1 && text[idx] == '-' && text[idx-1] == '.' {
			// It is ".-" wildcard, not a suffix
		} else if idx > 0 && text[idx] == '-' && text[idx-1] == '.' {
			// It matches ".-" pattern
		} else {
			// It is a real suffix
			suffix = text[idx:]
			base = text[:idx]
		}
	}

	isWildcard := strings.HasSuffix(base, ".-")
	if isWildcard {
		base = strings.TrimSuffix(base, ".-")
	}

	parts := strings.Split(strings.ReplaceAll(base, ",", "."), ".")
	v := &ast.VersionLiteral{IsWildcard: isWildcard, Suffix: suffix}

	if len(parts) > 0 {
		i, _ := strconv.ParseUint(parts[0], 10, 16)
		v.Major = uint16(i)
	} else {
		panic("should not have a version with zero parts")
	}

	if len(parts) > 1 {
		i, _ := strconv.ParseUint(parts[1], 10, 16)
		v.Minor = uint16(i)
	} else {
		panic("should not have a version with one part")
	}

	if len(parts) > 2 {
		i, _ := strconv.ParseUint(parts[2], 10, 32)
		v.Patch = uint32(i)
	} else if isWildcard {
		// 1.2.- -> Patch is Wild.
		v.Patch = 0xFFFFFFFF
	} else {
		panic("should not have only two parts and no wildcard")
	}

	// Returns *ast.VersionLiteral populated with Major, Minor, Patch, IsWildcard
	return v
}

func (b *ASTBuilder) VisitVersionNumber(ctx *grammar.VersionNumberContext) interface{} {
	return b.parseVersionString(ctx.GetText())
}

func (b *ASTBuilder) VisitDateNumber(ctx *grammar.DateNumberContext) interface{} {
	text := ctx.GetText()
	// Formats: 2025/01/01, 2025/01/01@12:00:00, 2025/01/01@12:00:00.000

	if strings.Contains(text, "/00") { // Crude check for 00 month/day
		// Manual Parse
		parts := strings.Split(text, "@")
		datePart := parts[0]
		timePart := ""
		if len(parts) > 1 {
			timePart = parts[1]
		}

		dateParts := strings.Split(datePart, "/")
		if len(dateParts) == 3 {
			y, _ := strconv.ParseInt(dateParts[0], 10, 64)
			m, _ := strconv.ParseInt(dateParts[1], 10, 64)
			d, _ := strconv.ParseInt(dateParts[2], 10, 64)

			ms := int64(0)
			ms += y * 365 * 24 * 3600 * 1000
			ms += m * 30 * 24 * 3600 * 1000
			ms += d * 24 * 3600 * 1000

			if timePart != "" {
				t, err := time.Parse("15:04:05", timePart)
				if err == nil {
					// Add hours/min/sec from t (relative to 0000-01-01 00:00:00)
					ms += int64(t.Hour()) * 3600 * 1000
					ms += int64(t.Minute()) * 60 * 1000
					ms += int64(t.Second()) * 1000
					ms += int64(t.Nanosecond()) / 1000000
				}
			}

			return &ast.DateTimeLiteral{Value: ms}
		}
	}

	layouts := []string{
		"2006/01/02@15:04:05.000",
		"2006/01/02@15:04:05",
		"2006/01/02",
		"15:04:05",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, text)
		if err == nil {
			return &ast.DateTimeLiteral{Value: t.UnixMilli()}
		}
	}

	panic(fmt.Errorf("not able to convert value '%s' to DateTime", text))
}

func (b *ASTBuilder) VisitDurationNumber(ctx *grammar.DurationNumberContext) interface{} {
	text := ctx.GetText()
	// "Time Only" literal (e.g. 00:05:00)
	// Defaults to Unix Epoch Date: 1970/01/01
	layouts := []string{
		"15:04:05.000",
		"15:04:05",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, text)
		if err == nil {
			// Force date to 1970-01-01
			t = time.Date(1970, 1, 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
			return &ast.DateTimeLiteral{Value: t.UnixMilli()}
		}
	}

	return &ast.DateTimeLiteral{Value: 0}
}

func (b *ASTBuilder) VisitZeroNumber(ctx *grammar.ZeroNumberContext) interface{} {
	return &ast.IntegerLiteral{Value: 0}
}
