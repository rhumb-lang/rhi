// Code generated from RhumbParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // RhumbParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type RhumbParser struct {
	*antlr.BaseParser
}

var RhumbParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rhumbparserParserInit() {
	staticData := &RhumbParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'&&'", "'&'", "'@'", "'@@'", "'\\'", "'\\\\'", "'\\/'", "'`'",
		"'!'", "'!!'", "'!>'", "'^'", "'^^'", "'^='", "'^/'", "':'", "'::'",
		"':='", "','", "',-'", "'-'", "'--'", "'-/'", "'->'", "'$'", "'.'",
		"'.-'", "'..'", "'.='", "'='", "'=@'", "'=\\'", "'=='", "'=>'", "'/'",
		"'/\\'", "'//'", "'>='", "'>>'", "'#'", "'<='", "'<>'", "'<<'", "'|'",
		"'|>'", "'||'", "'+'", "'+-'", "'+/'", "'++'", "'?'", "'??'", "';'",
		"'*'", "'*^'", "'**'", "'~'", "'~@'", "'~\\'", "'~~'", "'~>'", "", "",
		"'0'", "", "", "", "", "'('", "')'", "'['", "']'", "'{'", "'}'", "'<'",
		"'>'", "", "", "", "", "", "", "", "", "'$('", "'${'", "", "'`$'", "'`\"'",
	}
	staticData.SymbolicNames = []string{
		"", "AmpAmp", "Ampersand", "At", "AtAt", "BSlash", "BSlashBSlash", "BSlashFSlash",
		"Backtick", "Bang", "BangBang", "BangGreater", "Caret", "CaretCaret",
		"CaretEqual", "CaretFSlash", "Colon", "ColonColon", "ColonEqual", "Comma",
		"CommaDash", "Dash", "DashDash", "DashFSlash", "DashGreater", "Dollar",
		"Dot", "DotDash", "DotDot", "DotEqual", "Equal", "EqualAt", "EqualBSlash",
		"EqualEqual", "EqualGreater", "FSlash", "FSlashBSlash", "FSlashFSlash",
		"GreaterEqual", "GreaterGreater", "Hash", "LesserEqual", "LesserGreater",
		"LesserLesser", "Pipe", "PipeGreater", "PipePipe", "Plus", "PlusDash",
		"PlusFSlash", "PlusPlus", "QMark", "QMarkQMark", "Semicolon", "Star",
		"StarCaret", "StarStar", "Tilde", "TildeAt", "TildeBSlash", "TildeTilde",
		"TildeGreater", "Version", "FloatingPoint", "Zero", "NumberPart", "NumberStart",
		"RawText", "OpenInterpString", "OpenParen", "CloseParen", "OpenBracket",
		"CloseBracket", "OpenCurly", "CloseCurly", "OpenAnglet", "CloseAnglet",
		"NL", "WS", "BlockComment", "LineComment", "Key", "Path", "Label", "Letter",
		"EnterRoutineInterp", "EnterSelectorInterp", "InterpLabel", "EscapedDollar",
		"EscapedQuote", "InnerText", "CloseInterpString",
	}
	staticData.RuleNames = []string{
		"expressions", "fields", "patterns", "terminator", "literal", "fieldLiteral",
		"expression", "chainExpression", "field", "pattern", "accessOp", "applicativeOp",
		"comparativeOp", "identityOp", "conjunctiveOp", "disjunctiveOp", "conditionalOp",
		"additiveOp", "multiplicativeOp", "exponentiationOp", "assignmentOp",
		"prefixOp", "chainOp", "datePart", "date", "text", "interpolation",
		"reference",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 91, 693, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 5, 0, 58, 8, 0, 10, 0, 12, 0, 61, 9, 0, 1, 0,
		1, 0, 1, 0, 3, 0, 66, 8, 0, 5, 0, 68, 8, 0, 10, 0, 12, 0, 71, 9, 0, 1,
		1, 5, 1, 74, 8, 1, 10, 1, 12, 1, 77, 9, 1, 1, 1, 1, 1, 1, 1, 3, 1, 82,
		8, 1, 5, 1, 84, 8, 1, 10, 1, 12, 1, 87, 9, 1, 1, 2, 5, 2, 90, 8, 2, 10,
		2, 12, 2, 93, 9, 2, 1, 2, 1, 2, 1, 2, 3, 2, 98, 8, 2, 5, 2, 100, 8, 2,
		10, 2, 12, 2, 103, 9, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 3, 4, 116, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 3, 5, 125, 8, 5, 1, 6, 1, 6, 1, 6, 3, 6, 130, 8, 6, 1, 6, 1, 6, 1,
		6, 3, 6, 135, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 142, 8, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 147, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 156, 8, 6, 1, 6, 5, 6, 159, 8, 6, 10, 6, 12, 6, 162, 9, 6, 1, 6,
		1, 6, 5, 6, 166, 8, 6, 10, 6, 12, 6, 169, 9, 6, 1, 6, 1, 6, 1, 6, 3, 6,
		174, 8, 6, 1, 6, 1, 6, 5, 6, 178, 8, 6, 10, 6, 12, 6, 181, 9, 6, 1, 6,
		1, 6, 5, 6, 185, 8, 6, 10, 6, 12, 6, 188, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		5, 6, 194, 8, 6, 10, 6, 12, 6, 197, 9, 6, 1, 6, 1, 6, 5, 6, 201, 8, 6,
		10, 6, 12, 6, 204, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 210, 8, 6, 10, 6,
		12, 6, 213, 9, 6, 1, 6, 1, 6, 5, 6, 217, 8, 6, 10, 6, 12, 6, 220, 9, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 226, 8, 6, 10, 6, 12, 6, 229, 9, 6, 1, 6,
		1, 6, 5, 6, 233, 8, 6, 10, 6, 12, 6, 236, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		5, 6, 242, 8, 6, 10, 6, 12, 6, 245, 9, 6, 1, 6, 1, 6, 5, 6, 249, 8, 6,
		10, 6, 12, 6, 252, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 258, 8, 6, 10, 6,
		12, 6, 261, 9, 6, 1, 6, 1, 6, 5, 6, 265, 8, 6, 10, 6, 12, 6, 268, 9, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 274, 8, 6, 10, 6, 12, 6, 277, 9, 6, 1, 6,
		1, 6, 5, 6, 281, 8, 6, 10, 6, 12, 6, 284, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		5, 6, 290, 8, 6, 10, 6, 12, 6, 293, 9, 6, 1, 6, 1, 6, 5, 6, 297, 8, 6,
		10, 6, 12, 6, 300, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 306, 8, 6, 10, 6,
		12, 6, 309, 9, 6, 1, 6, 1, 6, 5, 6, 313, 8, 6, 10, 6, 12, 6, 316, 9, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 323, 8, 6, 10, 6, 12, 6, 326, 9, 6,
		1, 6, 3, 6, 329, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 335, 8, 6, 10, 6,
		12, 6, 338, 9, 6, 1, 6, 1, 6, 3, 6, 342, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		5, 6, 348, 8, 6, 10, 6, 12, 6, 351, 9, 6, 1, 6, 3, 6, 354, 8, 6, 1, 6,
		5, 6, 357, 8, 6, 10, 6, 12, 6, 360, 9, 6, 1, 7, 1, 7, 1, 7, 3, 7, 365,
		8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 371, 8, 7, 10, 7, 12, 7, 374, 9, 7,
		1, 7, 1, 7, 3, 7, 378, 8, 7, 1, 7, 1, 7, 1, 7, 5, 7, 383, 8, 7, 10, 7,
		12, 7, 386, 9, 7, 1, 7, 3, 7, 389, 8, 7, 1, 7, 4, 7, 392, 8, 7, 11, 7,
		12, 7, 393, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 410, 8, 8, 10, 8, 12, 8, 413, 9, 8, 1, 8,
		1, 8, 5, 8, 417, 8, 8, 10, 8, 12, 8, 420, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 5, 8, 427, 8, 8, 10, 8, 12, 8, 430, 9, 8, 1, 8, 1, 8, 5, 8, 434,
		8, 8, 10, 8, 12, 8, 437, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 443, 8, 8,
		10, 8, 12, 8, 446, 9, 8, 1, 8, 1, 8, 5, 8, 450, 8, 8, 10, 8, 12, 8, 453,
		9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 460, 8, 8, 10, 8, 12, 8, 463,
		9, 8, 1, 8, 1, 8, 5, 8, 467, 8, 8, 10, 8, 12, 8, 470, 9, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 476, 8, 8, 1, 8, 1, 8, 3, 8, 480, 8, 8, 1, 9, 1, 9, 5,
		9, 484, 8, 9, 10, 9, 12, 9, 487, 9, 9, 1, 9, 1, 9, 5, 9, 491, 8, 9, 10,
		9, 12, 9, 494, 9, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 500, 8, 9, 10, 9, 12,
		9, 503, 9, 9, 1, 9, 1, 9, 5, 9, 507, 8, 9, 10, 9, 12, 9, 510, 9, 9, 1,
		9, 1, 9, 1, 9, 3, 9, 515, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 3, 10, 536, 8, 10, 1, 11, 1, 11, 3, 11, 540, 8, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 3, 12, 546, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 3, 13, 554, 8, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 567, 8, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 3, 17, 573, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3,
		18, 580, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 586, 8, 19, 1, 20, 1,
		20, 3, 20, 590, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 603, 8, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 3, 22, 609, 8, 22, 1, 23, 1, 23, 1, 23, 3, 23, 614, 8, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 3, 25, 623, 8, 25, 1, 25, 1, 25,
		1, 25, 5, 25, 628, 8, 25, 10, 25, 12, 25, 631, 9, 25, 1, 25, 1, 25, 3,
		25, 635, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 3, 26, 646, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 5, 27, 658, 8, 27, 10, 27, 12, 27, 661, 9, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 670, 8, 27, 10,
		27, 12, 27, 673, 9, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		5, 27, 682, 8, 27, 10, 27, 12, 27, 685, 9, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 3, 27, 691, 8, 27, 1, 27, 0, 1, 12, 28, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 0, 1, 2, 0, 53, 53, 77, 77, 843, 0, 59, 1, 0, 0, 0, 2, 75, 1, 0,
		0, 0, 4, 91, 1, 0, 0, 0, 6, 104, 1, 0, 0, 0, 8, 115, 1, 0, 0, 0, 10, 124,
		1, 0, 0, 0, 12, 173, 1, 0, 0, 0, 14, 361, 1, 0, 0, 0, 16, 479, 1, 0, 0,
		0, 18, 514, 1, 0, 0, 0, 20, 535, 1, 0, 0, 0, 22, 539, 1, 0, 0, 0, 24, 545,
		1, 0, 0, 0, 26, 553, 1, 0, 0, 0, 28, 555, 1, 0, 0, 0, 30, 557, 1, 0, 0,
		0, 32, 566, 1, 0, 0, 0, 34, 572, 1, 0, 0, 0, 36, 579, 1, 0, 0, 0, 38, 585,
		1, 0, 0, 0, 40, 589, 1, 0, 0, 0, 42, 602, 1, 0, 0, 0, 44, 608, 1, 0, 0,
		0, 46, 613, 1, 0, 0, 0, 48, 615, 1, 0, 0, 0, 50, 634, 1, 0, 0, 0, 52, 645,
		1, 0, 0, 0, 54, 690, 1, 0, 0, 0, 56, 58, 3, 6, 3, 0, 57, 56, 1, 0, 0, 0,
		58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 62, 1,
		0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 69, 3, 12, 6, 0, 63, 65, 3, 6, 3, 0, 64,
		66, 3, 12, 6, 0, 65, 64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1, 0,
		0, 0, 67, 63, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70,
		1, 0, 0, 0, 70, 1, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 74, 3, 6, 3, 0,
		73, 72, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1,
		0, 0, 0, 76, 78, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 85, 3, 16, 8, 0, 79,
		81, 3, 6, 3, 0, 80, 82, 3, 16, 8, 0, 81, 80, 1, 0, 0, 0, 81, 82, 1, 0,
		0, 0, 82, 84, 1, 0, 0, 0, 83, 79, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83,
		1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 3, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0,
		88, 90, 3, 6, 3, 0, 89, 88, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1,
		0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94,
		101, 3, 18, 9, 0, 95, 97, 3, 6, 3, 0, 96, 98, 3, 18, 9, 0, 97, 96, 1, 0,
		0, 0, 97, 98, 1, 0, 0, 0, 98, 100, 1, 0, 0, 0, 99, 95, 1, 0, 0, 0, 100,
		103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 5, 1,
		0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 105, 7, 0, 0, 0, 105, 7, 1, 0, 0, 0,
		106, 116, 5, 63, 0, 0, 107, 116, 3, 48, 24, 0, 108, 116, 5, 64, 0, 0, 109,
		116, 5, 65, 0, 0, 110, 116, 5, 81, 0, 0, 111, 116, 3, 50, 25, 0, 112, 116,
		3, 54, 27, 0, 113, 116, 5, 9, 0, 0, 114, 116, 5, 83, 0, 0, 115, 106, 1,
		0, 0, 0, 115, 107, 1, 0, 0, 0, 115, 108, 1, 0, 0, 0, 115, 109, 1, 0, 0,
		0, 115, 110, 1, 0, 0, 0, 115, 111, 1, 0, 0, 0, 115, 112, 1, 0, 0, 0, 115,
		113, 1, 0, 0, 0, 115, 114, 1, 0, 0, 0, 116, 9, 1, 0, 0, 0, 117, 125, 5,
		64, 0, 0, 118, 125, 5, 65, 0, 0, 119, 125, 5, 81, 0, 0, 120, 125, 3, 50,
		25, 0, 121, 125, 3, 54, 27, 0, 122, 125, 5, 9, 0, 0, 123, 125, 5, 83, 0,
		0, 124, 117, 1, 0, 0, 0, 124, 118, 1, 0, 0, 0, 124, 119, 1, 0, 0, 0, 124,
		120, 1, 0, 0, 0, 124, 121, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 123,
		1, 0, 0, 0, 125, 11, 1, 0, 0, 0, 126, 127, 6, 6, -1, 0, 127, 129, 5, 71,
		0, 0, 128, 130, 3, 2, 1, 0, 129, 128, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0,
		130, 131, 1, 0, 0, 0, 131, 174, 5, 72, 0, 0, 132, 134, 5, 73, 0, 0, 133,
		135, 5, 83, 0, 0, 134, 133, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 136,
		1, 0, 0, 0, 136, 137, 5, 82, 0, 0, 137, 138, 5, 62, 0, 0, 138, 174, 5,
		74, 0, 0, 139, 141, 5, 73, 0, 0, 140, 142, 3, 4, 2, 0, 141, 140, 1, 0,
		0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 174, 5, 74, 0, 0,
		144, 146, 5, 69, 0, 0, 145, 147, 3, 0, 0, 0, 146, 145, 1, 0, 0, 0, 146,
		147, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 174, 5, 70, 0, 0, 149, 174,
		3, 14, 7, 0, 150, 151, 3, 42, 21, 0, 151, 152, 3, 12, 6, 11, 152, 174,
		1, 0, 0, 0, 153, 156, 3, 14, 7, 0, 154, 156, 3, 10, 5, 0, 155, 153, 1,
		0, 0, 0, 155, 154, 1, 0, 0, 0, 156, 160, 1, 0, 0, 0, 157, 159, 5, 77, 0,
		0, 158, 157, 1, 0, 0, 0, 159, 162, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160,
		161, 1, 0, 0, 0, 161, 163, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 163, 167,
		3, 40, 20, 0, 164, 166, 5, 77, 0, 0, 165, 164, 1, 0, 0, 0, 166, 169, 1,
		0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 170, 1, 0, 0,
		0, 169, 167, 1, 0, 0, 0, 170, 171, 3, 12, 6, 2, 171, 174, 1, 0, 0, 0, 172,
		174, 3, 8, 4, 0, 173, 126, 1, 0, 0, 0, 173, 132, 1, 0, 0, 0, 173, 139,
		1, 0, 0, 0, 173, 144, 1, 0, 0, 0, 173, 149, 1, 0, 0, 0, 173, 150, 1, 0,
		0, 0, 173, 155, 1, 0, 0, 0, 173, 172, 1, 0, 0, 0, 174, 358, 1, 0, 0, 0,
		175, 179, 10, 12, 0, 0, 176, 178, 5, 77, 0, 0, 177, 176, 1, 0, 0, 0, 178,
		181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 182,
		1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182, 186, 3, 38, 19, 0, 183, 185, 5,
		77, 0, 0, 184, 183, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0,
		0, 186, 187, 1, 0, 0, 0, 187, 189, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189,
		190, 3, 12, 6, 12, 190, 357, 1, 0, 0, 0, 191, 195, 10, 10, 0, 0, 192, 194,
		5, 77, 0, 0, 193, 192, 1, 0, 0, 0, 194, 197, 1, 0, 0, 0, 195, 193, 1, 0,
		0, 0, 195, 196, 1, 0, 0, 0, 196, 198, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0,
		198, 202, 3, 36, 18, 0, 199, 201, 5, 77, 0, 0, 200, 199, 1, 0, 0, 0, 201,
		204, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 205,
		1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 205, 206, 3, 12, 6, 11, 206, 357, 1,
		0, 0, 0, 207, 211, 10, 9, 0, 0, 208, 210, 5, 77, 0, 0, 209, 208, 1, 0,
		0, 0, 210, 213, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0,
		212, 214, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 214, 218, 3, 34, 17, 0, 215,
		217, 5, 77, 0, 0, 216, 215, 1, 0, 0, 0, 217, 220, 1, 0, 0, 0, 218, 216,
		1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 221, 1, 0, 0, 0, 220, 218, 1, 0,
		0, 0, 221, 222, 3, 12, 6, 10, 222, 357, 1, 0, 0, 0, 223, 227, 10, 8, 0,
		0, 224, 226, 5, 77, 0, 0, 225, 224, 1, 0, 0, 0, 226, 229, 1, 0, 0, 0, 227,
		225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 230, 1, 0, 0, 0, 229, 227,
		1, 0, 0, 0, 230, 234, 3, 24, 12, 0, 231, 233, 5, 77, 0, 0, 232, 231, 1,
		0, 0, 0, 233, 236, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0,
		0, 235, 237, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 237, 238, 3, 12, 6, 9, 238,
		357, 1, 0, 0, 0, 239, 243, 10, 7, 0, 0, 240, 242, 5, 77, 0, 0, 241, 240,
		1, 0, 0, 0, 242, 245, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0,
		0, 0, 244, 246, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 246, 250, 3, 26, 13,
		0, 247, 249, 5, 77, 0, 0, 248, 247, 1, 0, 0, 0, 249, 252, 1, 0, 0, 0, 250,
		248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 253, 1, 0, 0, 0, 252, 250,
		1, 0, 0, 0, 253, 254, 3, 12, 6, 8, 254, 357, 1, 0, 0, 0, 255, 259, 10,
		6, 0, 0, 256, 258, 5, 77, 0, 0, 257, 256, 1, 0, 0, 0, 258, 261, 1, 0, 0,
		0, 259, 257, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 262, 1, 0, 0, 0, 261,
		259, 1, 0, 0, 0, 262, 266, 3, 28, 14, 0, 263, 265, 5, 77, 0, 0, 264, 263,
		1, 0, 0, 0, 265, 268, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0,
		0, 0, 267, 269, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 269, 270, 3, 12, 6, 7,
		270, 357, 1, 0, 0, 0, 271, 275, 10, 5, 0, 0, 272, 274, 5, 77, 0, 0, 273,
		272, 1, 0, 0, 0, 274, 277, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276,
		1, 0, 0, 0, 276, 278, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 278, 282, 3, 30,
		15, 0, 279, 281, 5, 77, 0, 0, 280, 279, 1, 0, 0, 0, 281, 284, 1, 0, 0,
		0, 282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 285, 1, 0, 0, 0, 284,
		282, 1, 0, 0, 0, 285, 286, 3, 12, 6, 6, 286, 357, 1, 0, 0, 0, 287, 291,
		10, 4, 0, 0, 288, 290, 5, 77, 0, 0, 289, 288, 1, 0, 0, 0, 290, 293, 1,
		0, 0, 0, 291, 289, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 294, 1, 0, 0,
		0, 293, 291, 1, 0, 0, 0, 294, 298, 3, 22, 11, 0, 295, 297, 5, 77, 0, 0,
		296, 295, 1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298,
		299, 1, 0, 0, 0, 299, 301, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 301, 302,
		3, 12, 6, 4, 302, 357, 1, 0, 0, 0, 303, 307, 10, 3, 0, 0, 304, 306, 5,
		77, 0, 0, 305, 304, 1, 0, 0, 0, 306, 309, 1, 0, 0, 0, 307, 305, 1, 0, 0,
		0, 307, 308, 1, 0, 0, 0, 308, 310, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 310,
		314, 3, 32, 16, 0, 311, 313, 5, 77, 0, 0, 312, 311, 1, 0, 0, 0, 313, 316,
		1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 317, 1, 0,
		0, 0, 316, 314, 1, 0, 0, 0, 317, 318, 3, 12, 6, 3, 318, 357, 1, 0, 0, 0,
		319, 320, 10, 15, 0, 0, 320, 324, 5, 69, 0, 0, 321, 323, 3, 6, 3, 0, 322,
		321, 1, 0, 0, 0, 323, 326, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 324, 325,
		1, 0, 0, 0, 325, 328, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 327, 329, 3, 0,
		0, 0, 328, 327, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0,
		330, 357, 5, 70, 0, 0, 331, 332, 10, 14, 0, 0, 332, 336, 5, 71, 0, 0, 333,
		335, 3, 6, 3, 0, 334, 333, 1, 0, 0, 0, 335, 338, 1, 0, 0, 0, 336, 334,
		1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 341, 1, 0, 0, 0, 338, 336, 1, 0,
		0, 0, 339, 342, 3, 0, 0, 0, 340, 342, 3, 20, 10, 0, 341, 339, 1, 0, 0,
		0, 341, 340, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343,
		357, 5, 72, 0, 0, 344, 345, 10, 13, 0, 0, 345, 349, 5, 73, 0, 0, 346, 348,
		3, 6, 3, 0, 347, 346, 1, 0, 0, 0, 348, 351, 1, 0, 0, 0, 349, 347, 1, 0,
		0, 0, 349, 350, 1, 0, 0, 0, 350, 353, 1, 0, 0, 0, 351, 349, 1, 0, 0, 0,
		352, 354, 3, 4, 2, 0, 353, 352, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354,
		355, 1, 0, 0, 0, 355, 357, 5, 74, 0, 0, 356, 175, 1, 0, 0, 0, 356, 191,
		1, 0, 0, 0, 356, 207, 1, 0, 0, 0, 356, 223, 1, 0, 0, 0, 356, 239, 1, 0,
		0, 0, 356, 255, 1, 0, 0, 0, 356, 271, 1, 0, 0, 0, 356, 287, 1, 0, 0, 0,
		356, 303, 1, 0, 0, 0, 356, 319, 1, 0, 0, 0, 356, 331, 1, 0, 0, 0, 356,
		344, 1, 0, 0, 0, 357, 360, 1, 0, 0, 0, 358, 356, 1, 0, 0, 0, 358, 359,
		1, 0, 0, 0, 359, 13, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 361, 391, 3, 10,
		5, 0, 362, 364, 3, 44, 22, 0, 363, 365, 3, 42, 21, 0, 364, 363, 1, 0, 0,
		0, 364, 365, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 367, 3, 10, 5, 0, 367,
		392, 1, 0, 0, 0, 368, 372, 5, 71, 0, 0, 369, 371, 3, 6, 3, 0, 370, 369,
		1, 0, 0, 0, 371, 374, 1, 0, 0, 0, 372, 370, 1, 0, 0, 0, 372, 373, 1, 0,
		0, 0, 373, 377, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 375, 378, 3, 12, 6, 0,
		376, 378, 3, 20, 10, 0, 377, 375, 1, 0, 0, 0, 377, 376, 1, 0, 0, 0, 377,
		378, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 392, 5, 72, 0, 0, 380, 384,
		5, 69, 0, 0, 381, 383, 3, 6, 3, 0, 382, 381, 1, 0, 0, 0, 383, 386, 1, 0,
		0, 0, 384, 382, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 388, 1, 0, 0, 0,
		386, 384, 1, 0, 0, 0, 387, 389, 3, 0, 0, 0, 388, 387, 1, 0, 0, 0, 388,
		389, 1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 392, 5, 70, 0, 0, 391, 362,
		1, 0, 0, 0, 391, 368, 1, 0, 0, 0, 391, 380, 1, 0, 0, 0, 392, 393, 1, 0,
		0, 0, 393, 391, 1, 0, 0, 0, 393, 394, 1, 0, 0, 0, 394, 15, 1, 0, 0, 0,
		395, 396, 5, 26, 0, 0, 396, 480, 3, 10, 5, 0, 397, 398, 5, 26, 0, 0, 398,
		399, 5, 3, 0, 0, 399, 480, 3, 10, 5, 0, 400, 401, 5, 16, 0, 0, 401, 480,
		3, 10, 5, 0, 402, 403, 5, 16, 0, 0, 403, 404, 5, 3, 0, 0, 404, 480, 3,
		10, 5, 0, 405, 406, 5, 2, 0, 0, 406, 480, 3, 10, 5, 0, 407, 411, 3, 10,
		5, 0, 408, 410, 5, 77, 0, 0, 409, 408, 1, 0, 0, 0, 410, 413, 1, 0, 0, 0,
		411, 409, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412, 414, 1, 0, 0, 0, 413,
		411, 1, 0, 0, 0, 414, 418, 5, 17, 0, 0, 415, 417, 5, 77, 0, 0, 416, 415,
		1, 0, 0, 0, 417, 420, 1, 0, 0, 0, 418, 416, 1, 0, 0, 0, 418, 419, 1, 0,
		0, 0, 419, 421, 1, 0, 0, 0, 420, 418, 1, 0, 0, 0, 421, 422, 3, 12, 6, 0,
		422, 480, 1, 0, 0, 0, 423, 424, 5, 3, 0, 0, 424, 428, 3, 10, 5, 0, 425,
		427, 5, 77, 0, 0, 426, 425, 1, 0, 0, 0, 427, 430, 1, 0, 0, 0, 428, 426,
		1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 431, 1, 0, 0, 0, 430, 428, 1, 0,
		0, 0, 431, 435, 5, 17, 0, 0, 432, 434, 5, 77, 0, 0, 433, 432, 1, 0, 0,
		0, 434, 437, 1, 0, 0, 0, 435, 433, 1, 0, 0, 0, 435, 436, 1, 0, 0, 0, 436,
		438, 1, 0, 0, 0, 437, 435, 1, 0, 0, 0, 438, 439, 3, 12, 6, 0, 439, 480,
		1, 0, 0, 0, 440, 444, 3, 10, 5, 0, 441, 443, 5, 77, 0, 0, 442, 441, 1,
		0, 0, 0, 443, 446, 1, 0, 0, 0, 444, 442, 1, 0, 0, 0, 444, 445, 1, 0, 0,
		0, 445, 447, 1, 0, 0, 0, 446, 444, 1, 0, 0, 0, 447, 451, 5, 28, 0, 0, 448,
		450, 5, 77, 0, 0, 449, 448, 1, 0, 0, 0, 450, 453, 1, 0, 0, 0, 451, 449,
		1, 0, 0, 0, 451, 452, 1, 0, 0, 0, 452, 454, 1, 0, 0, 0, 453, 451, 1, 0,
		0, 0, 454, 455, 3, 12, 6, 0, 455, 480, 1, 0, 0, 0, 456, 457, 5, 3, 0, 0,
		457, 461, 3, 10, 5, 0, 458, 460, 5, 77, 0, 0, 459, 458, 1, 0, 0, 0, 460,
		463, 1, 0, 0, 0, 461, 459, 1, 0, 0, 0, 461, 462, 1, 0, 0, 0, 462, 464,
		1, 0, 0, 0, 463, 461, 1, 0, 0, 0, 464, 468, 5, 28, 0, 0, 465, 467, 5, 77,
		0, 0, 466, 465, 1, 0, 0, 0, 467, 470, 1, 0, 0, 0, 468, 466, 1, 0, 0, 0,
		468, 469, 1, 0, 0, 0, 469, 471, 1, 0, 0, 0, 470, 468, 1, 0, 0, 0, 471,
		472, 3, 12, 6, 0, 472, 480, 1, 0, 0, 0, 473, 475, 5, 71, 0, 0, 474, 476,
		3, 2, 1, 0, 475, 474, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 477, 1, 0,
		0, 0, 477, 480, 5, 72, 0, 0, 478, 480, 3, 12, 6, 0, 479, 395, 1, 0, 0,
		0, 479, 397, 1, 0, 0, 0, 479, 400, 1, 0, 0, 0, 479, 402, 1, 0, 0, 0, 479,
		405, 1, 0, 0, 0, 479, 407, 1, 0, 0, 0, 479, 423, 1, 0, 0, 0, 479, 440,
		1, 0, 0, 0, 479, 456, 1, 0, 0, 0, 479, 473, 1, 0, 0, 0, 479, 478, 1, 0,
		0, 0, 480, 17, 1, 0, 0, 0, 481, 485, 3, 12, 6, 0, 482, 484, 5, 77, 0, 0,
		483, 482, 1, 0, 0, 0, 484, 487, 1, 0, 0, 0, 485, 483, 1, 0, 0, 0, 485,
		486, 1, 0, 0, 0, 486, 488, 1, 0, 0, 0, 487, 485, 1, 0, 0, 0, 488, 492,
		5, 28, 0, 0, 489, 491, 5, 77, 0, 0, 490, 489, 1, 0, 0, 0, 491, 494, 1,
		0, 0, 0, 492, 490, 1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 495, 1, 0, 0,
		0, 494, 492, 1, 0, 0, 0, 495, 496, 3, 12, 6, 0, 496, 515, 1, 0, 0, 0, 497,
		501, 3, 12, 6, 0, 498, 500, 5, 77, 0, 0, 499, 498, 1, 0, 0, 0, 500, 503,
		1, 0, 0, 0, 501, 499, 1, 0, 0, 0, 501, 502, 1, 0, 0, 0, 502, 504, 1, 0,
		0, 0, 503, 501, 1, 0, 0, 0, 504, 508, 5, 17, 0, 0, 505, 507, 5, 77, 0,
		0, 506, 505, 1, 0, 0, 0, 507, 510, 1, 0, 0, 0, 508, 506, 1, 0, 0, 0, 508,
		509, 1, 0, 0, 0, 509, 511, 1, 0, 0, 0, 510, 508, 1, 0, 0, 0, 511, 512,
		3, 12, 6, 0, 512, 515, 1, 0, 0, 0, 513, 515, 3, 12, 6, 0, 514, 481, 1,
		0, 0, 0, 514, 497, 1, 0, 0, 0, 514, 513, 1, 0, 0, 0, 515, 19, 1, 0, 0,
		0, 516, 536, 5, 76, 0, 0, 517, 536, 5, 75, 0, 0, 518, 536, 5, 40, 0, 0,
		519, 536, 5, 51, 0, 0, 520, 536, 5, 3, 0, 0, 521, 536, 5, 54, 0, 0, 522,
		536, 5, 64, 0, 0, 523, 536, 5, 26, 0, 0, 524, 536, 5, 16, 0, 0, 525, 536,
		5, 35, 0, 0, 526, 536, 5, 25, 0, 0, 527, 536, 5, 12, 0, 0, 528, 536, 5,
		9, 0, 0, 529, 536, 5, 47, 0, 0, 530, 536, 5, 21, 0, 0, 531, 536, 5, 30,
		0, 0, 532, 536, 5, 57, 0, 0, 533, 536, 5, 2, 0, 0, 534, 536, 5, 8, 0, 0,
		535, 516, 1, 0, 0, 0, 535, 517, 1, 0, 0, 0, 535, 518, 1, 0, 0, 0, 535,
		519, 1, 0, 0, 0, 535, 520, 1, 0, 0, 0, 535, 521, 1, 0, 0, 0, 535, 522,
		1, 0, 0, 0, 535, 523, 1, 0, 0, 0, 535, 524, 1, 0, 0, 0, 535, 525, 1, 0,
		0, 0, 535, 526, 1, 0, 0, 0, 535, 527, 1, 0, 0, 0, 535, 528, 1, 0, 0, 0,
		535, 529, 1, 0, 0, 0, 535, 530, 1, 0, 0, 0, 535, 531, 1, 0, 0, 0, 535,
		532, 1, 0, 0, 0, 535, 533, 1, 0, 0, 0, 535, 534, 1, 0, 0, 0, 536, 21, 1,
		0, 0, 0, 537, 540, 5, 24, 0, 0, 538, 540, 5, 11, 0, 0, 539, 537, 1, 0,
		0, 0, 539, 538, 1, 0, 0, 0, 540, 23, 1, 0, 0, 0, 541, 546, 5, 39, 0, 0,
		542, 546, 5, 38, 0, 0, 543, 546, 5, 43, 0, 0, 544, 546, 5, 41, 0, 0, 545,
		541, 1, 0, 0, 0, 545, 542, 1, 0, 0, 0, 545, 543, 1, 0, 0, 0, 545, 544,
		1, 0, 0, 0, 546, 25, 1, 0, 0, 0, 547, 554, 5, 33, 0, 0, 548, 554, 5, 32,
		0, 0, 549, 554, 5, 31, 0, 0, 550, 554, 5, 60, 0, 0, 551, 554, 5, 59, 0,
		0, 552, 554, 5, 58, 0, 0, 553, 547, 1, 0, 0, 0, 553, 548, 1, 0, 0, 0, 553,
		549, 1, 0, 0, 0, 553, 550, 1, 0, 0, 0, 553, 551, 1, 0, 0, 0, 553, 552,
		1, 0, 0, 0, 554, 27, 1, 0, 0, 0, 555, 556, 5, 36, 0, 0, 556, 29, 1, 0,
		0, 0, 557, 558, 5, 7, 0, 0, 558, 31, 1, 0, 0, 0, 559, 567, 5, 46, 0, 0,
		560, 567, 5, 52, 0, 0, 561, 567, 5, 42, 0, 0, 562, 567, 5, 45, 0, 0, 563,
		567, 5, 34, 0, 0, 564, 567, 5, 61, 0, 0, 565, 567, 5, 14, 0, 0, 566, 559,
		1, 0, 0, 0, 566, 560, 1, 0, 0, 0, 566, 561, 1, 0, 0, 0, 566, 562, 1, 0,
		0, 0, 566, 563, 1, 0, 0, 0, 566, 564, 1, 0, 0, 0, 566, 565, 1, 0, 0, 0,
		567, 33, 1, 0, 0, 0, 568, 573, 5, 50, 0, 0, 569, 573, 5, 48, 0, 0, 570,
		573, 5, 22, 0, 0, 571, 573, 5, 1, 0, 0, 572, 568, 1, 0, 0, 0, 572, 569,
		1, 0, 0, 0, 572, 570, 1, 0, 0, 0, 572, 571, 1, 0, 0, 0, 573, 35, 1, 0,
		0, 0, 574, 580, 5, 56, 0, 0, 575, 580, 5, 37, 0, 0, 576, 580, 5, 49, 0,
		0, 577, 580, 5, 23, 0, 0, 578, 580, 5, 10, 0, 0, 579, 574, 1, 0, 0, 0,
		579, 575, 1, 0, 0, 0, 579, 576, 1, 0, 0, 0, 579, 577, 1, 0, 0, 0, 579,
		578, 1, 0, 0, 0, 580, 37, 1, 0, 0, 0, 581, 586, 5, 13, 0, 0, 582, 586,
		5, 15, 0, 0, 583, 586, 5, 44, 0, 0, 584, 586, 5, 55, 0, 0, 585, 581, 1,
		0, 0, 0, 585, 582, 1, 0, 0, 0, 585, 583, 1, 0, 0, 0, 585, 584, 1, 0, 0,
		0, 586, 39, 1, 0, 0, 0, 587, 590, 5, 29, 0, 0, 588, 590, 5, 18, 0, 0, 589,
		587, 1, 0, 0, 0, 589, 588, 1, 0, 0, 0, 590, 41, 1, 0, 0, 0, 591, 603, 5,
		51, 0, 0, 592, 603, 5, 26, 0, 0, 593, 603, 5, 16, 0, 0, 594, 603, 5, 47,
		0, 0, 595, 603, 5, 21, 0, 0, 596, 603, 5, 30, 0, 0, 597, 603, 5, 57, 0,
		0, 598, 603, 5, 2, 0, 0, 599, 603, 5, 25, 0, 0, 600, 603, 5, 40, 0, 0,
		601, 603, 5, 12, 0, 0, 602, 591, 1, 0, 0, 0, 602, 592, 1, 0, 0, 0, 602,
		593, 1, 0, 0, 0, 602, 594, 1, 0, 0, 0, 602, 595, 1, 0, 0, 0, 602, 596,
		1, 0, 0, 0, 602, 597, 1, 0, 0, 0, 602, 598, 1, 0, 0, 0, 602, 599, 1, 0,
		0, 0, 602, 600, 1, 0, 0, 0, 602, 601, 1, 0, 0, 0, 603, 43, 1, 0, 0, 0,
		604, 609, 5, 5, 0, 0, 605, 609, 5, 6, 0, 0, 606, 609, 5, 3, 0, 0, 607,
		609, 5, 4, 0, 0, 608, 604, 1, 0, 0, 0, 608, 605, 1, 0, 0, 0, 608, 606,
		1, 0, 0, 0, 608, 607, 1, 0, 0, 0, 609, 45, 1, 0, 0, 0, 610, 614, 5, 65,
		0, 0, 611, 612, 5, 25, 0, 0, 612, 614, 5, 83, 0, 0, 613, 610, 1, 0, 0,
		0, 613, 611, 1, 0, 0, 0, 614, 47, 1, 0, 0, 0, 615, 616, 3, 46, 23, 0, 616,
		617, 5, 35, 0, 0, 617, 618, 3, 46, 23, 0, 618, 619, 5, 35, 0, 0, 619, 620,
		3, 46, 23, 0, 620, 49, 1, 0, 0, 0, 621, 623, 5, 9, 0, 0, 622, 621, 1, 0,
		0, 0, 622, 623, 1, 0, 0, 0, 623, 624, 1, 0, 0, 0, 624, 629, 5, 68, 0, 0,
		625, 628, 5, 90, 0, 0, 626, 628, 3, 52, 26, 0, 627, 625, 1, 0, 0, 0, 627,
		626, 1, 0, 0, 0, 628, 631, 1, 0, 0, 0, 629, 627, 1, 0, 0, 0, 629, 630,
		1, 0, 0, 0, 630, 632, 1, 0, 0, 0, 631, 629, 1, 0, 0, 0, 632, 635, 5, 91,
		0, 0, 633, 635, 5, 67, 0, 0, 634, 622, 1, 0, 0, 0, 634, 633, 1, 0, 0, 0,
		635, 51, 1, 0, 0, 0, 636, 646, 5, 87, 0, 0, 637, 638, 5, 85, 0, 0, 638,
		639, 3, 0, 0, 0, 639, 640, 5, 70, 0, 0, 640, 646, 1, 0, 0, 0, 641, 642,
		5, 86, 0, 0, 642, 643, 3, 0, 0, 0, 643, 644, 5, 74, 0, 0, 644, 646, 1,
		0, 0, 0, 645, 636, 1, 0, 0, 0, 645, 637, 1, 0, 0, 0, 645, 641, 1, 0, 0,
		0, 646, 53, 1, 0, 0, 0, 647, 648, 5, 75, 0, 0, 648, 649, 5, 83, 0, 0, 649,
		691, 5, 76, 0, 0, 650, 651, 5, 75, 0, 0, 651, 652, 3, 50, 25, 0, 652, 653,
		5, 76, 0, 0, 653, 691, 1, 0, 0, 0, 654, 655, 5, 75, 0, 0, 655, 659, 5,
		69, 0, 0, 656, 658, 3, 6, 3, 0, 657, 656, 1, 0, 0, 0, 658, 661, 1, 0, 0,
		0, 659, 657, 1, 0, 0, 0, 659, 660, 1, 0, 0, 0, 660, 662, 1, 0, 0, 0, 661,
		659, 1, 0, 0, 0, 662, 663, 3, 0, 0, 0, 663, 664, 5, 70, 0, 0, 664, 665,
		5, 76, 0, 0, 665, 691, 1, 0, 0, 0, 666, 667, 5, 75, 0, 0, 667, 671, 5,
		71, 0, 0, 668, 670, 3, 6, 3, 0, 669, 668, 1, 0, 0, 0, 670, 673, 1, 0, 0,
		0, 671, 669, 1, 0, 0, 0, 671, 672, 1, 0, 0, 0, 672, 674, 1, 0, 0, 0, 673,
		671, 1, 0, 0, 0, 674, 675, 3, 0, 0, 0, 675, 676, 5, 72, 0, 0, 676, 677,
		5, 76, 0, 0, 677, 691, 1, 0, 0, 0, 678, 679, 5, 75, 0, 0, 679, 683, 5,
		73, 0, 0, 680, 682, 3, 6, 3, 0, 681, 680, 1, 0, 0, 0, 682, 685, 1, 0, 0,
		0, 683, 681, 1, 0, 0, 0, 683, 684, 1, 0, 0, 0, 684, 686, 1, 0, 0, 0, 685,
		683, 1, 0, 0, 0, 686, 687, 3, 0, 0, 0, 687, 688, 5, 74, 0, 0, 688, 689,
		5, 76, 0, 0, 689, 691, 1, 0, 0, 0, 690, 647, 1, 0, 0, 0, 690, 650, 1, 0,
		0, 0, 690, 654, 1, 0, 0, 0, 690, 666, 1, 0, 0, 0, 690, 678, 1, 0, 0, 0,
		691, 55, 1, 0, 0, 0, 88, 59, 65, 69, 75, 81, 85, 91, 97, 101, 115, 124,
		129, 134, 141, 146, 155, 160, 167, 173, 179, 186, 195, 202, 211, 218, 227,
		234, 243, 250, 259, 266, 275, 282, 291, 298, 307, 314, 324, 328, 336, 341,
		349, 353, 356, 358, 364, 372, 377, 384, 388, 391, 393, 411, 418, 428, 435,
		444, 451, 461, 468, 475, 479, 485, 492, 501, 508, 514, 535, 539, 545, 553,
		566, 572, 579, 585, 589, 602, 608, 613, 622, 627, 629, 634, 645, 659, 671,
		683, 690,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// RhumbParserInit initializes any static state used to implement RhumbParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRhumbParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RhumbParserInit() {
	staticData := &RhumbParserParserStaticData
	staticData.once.Do(rhumbparserParserInit)
}

// NewRhumbParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRhumbParser(input antlr.TokenStream) *RhumbParser {
	RhumbParserInit()
	this := new(RhumbParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &RhumbParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "RhumbParser.g4"

	return this
}

// RhumbParser tokens.
const (
	RhumbParserEOF                 = antlr.TokenEOF
	RhumbParserAmpAmp              = 1
	RhumbParserAmpersand           = 2
	RhumbParserAt                  = 3
	RhumbParserAtAt                = 4
	RhumbParserBSlash              = 5
	RhumbParserBSlashBSlash        = 6
	RhumbParserBSlashFSlash        = 7
	RhumbParserBacktick            = 8
	RhumbParserBang                = 9
	RhumbParserBangBang            = 10
	RhumbParserBangGreater         = 11
	RhumbParserCaret               = 12
	RhumbParserCaretCaret          = 13
	RhumbParserCaretEqual          = 14
	RhumbParserCaretFSlash         = 15
	RhumbParserColon               = 16
	RhumbParserColonColon          = 17
	RhumbParserColonEqual          = 18
	RhumbParserComma               = 19
	RhumbParserCommaDash           = 20
	RhumbParserDash                = 21
	RhumbParserDashDash            = 22
	RhumbParserDashFSlash          = 23
	RhumbParserDashGreater         = 24
	RhumbParserDollar              = 25
	RhumbParserDot                 = 26
	RhumbParserDotDash             = 27
	RhumbParserDotDot              = 28
	RhumbParserDotEqual            = 29
	RhumbParserEqual               = 30
	RhumbParserEqualAt             = 31
	RhumbParserEqualBSlash         = 32
	RhumbParserEqualEqual          = 33
	RhumbParserEqualGreater        = 34
	RhumbParserFSlash              = 35
	RhumbParserFSlashBSlash        = 36
	RhumbParserFSlashFSlash        = 37
	RhumbParserGreaterEqual        = 38
	RhumbParserGreaterGreater      = 39
	RhumbParserHash                = 40
	RhumbParserLesserEqual         = 41
	RhumbParserLesserGreater       = 42
	RhumbParserLesserLesser        = 43
	RhumbParserPipe                = 44
	RhumbParserPipeGreater         = 45
	RhumbParserPipePipe            = 46
	RhumbParserPlus                = 47
	RhumbParserPlusDash            = 48
	RhumbParserPlusFSlash          = 49
	RhumbParserPlusPlus            = 50
	RhumbParserQMark               = 51
	RhumbParserQMarkQMark          = 52
	RhumbParserSemicolon           = 53
	RhumbParserStar                = 54
	RhumbParserStarCaret           = 55
	RhumbParserStarStar            = 56
	RhumbParserTilde               = 57
	RhumbParserTildeAt             = 58
	RhumbParserTildeBSlash         = 59
	RhumbParserTildeTilde          = 60
	RhumbParserTildeGreater        = 61
	RhumbParserVersion             = 62
	RhumbParserFloatingPoint       = 63
	RhumbParserZero                = 64
	RhumbParserNumberPart          = 65
	RhumbParserNumberStart         = 66
	RhumbParserRawText             = 67
	RhumbParserOpenInterpString    = 68
	RhumbParserOpenParen           = 69
	RhumbParserCloseParen          = 70
	RhumbParserOpenBracket         = 71
	RhumbParserCloseBracket        = 72
	RhumbParserOpenCurly           = 73
	RhumbParserCloseCurly          = 74
	RhumbParserOpenAnglet          = 75
	RhumbParserCloseAnglet         = 76
	RhumbParserNL                  = 77
	RhumbParserWS                  = 78
	RhumbParserBlockComment        = 79
	RhumbParserLineComment         = 80
	RhumbParserKey                 = 81
	RhumbParserPath                = 82
	RhumbParserLabel               = 83
	RhumbParserLetter              = 84
	RhumbParserEnterRoutineInterp  = 85
	RhumbParserEnterSelectorInterp = 86
	RhumbParserInterpLabel         = 87
	RhumbParserEscapedDollar       = 88
	RhumbParserEscapedQuote        = 89
	RhumbParserInnerText           = 90
	RhumbParserCloseInterpString   = 91
)

// RhumbParser rules.
const (
	RhumbParserRULE_expressions      = 0
	RhumbParserRULE_fields           = 1
	RhumbParserRULE_patterns         = 2
	RhumbParserRULE_terminator       = 3
	RhumbParserRULE_literal          = 4
	RhumbParserRULE_fieldLiteral     = 5
	RhumbParserRULE_expression       = 6
	RhumbParserRULE_chainExpression  = 7
	RhumbParserRULE_field            = 8
	RhumbParserRULE_pattern          = 9
	RhumbParserRULE_accessOp         = 10
	RhumbParserRULE_applicativeOp    = 11
	RhumbParserRULE_comparativeOp    = 12
	RhumbParserRULE_identityOp       = 13
	RhumbParserRULE_conjunctiveOp    = 14
	RhumbParserRULE_disjunctiveOp    = 15
	RhumbParserRULE_conditionalOp    = 16
	RhumbParserRULE_additiveOp       = 17
	RhumbParserRULE_multiplicativeOp = 18
	RhumbParserRULE_exponentiationOp = 19
	RhumbParserRULE_assignmentOp     = 20
	RhumbParserRULE_prefixOp         = 21
	RhumbParserRULE_chainOp          = 22
	RhumbParserRULE_datePart         = 23
	RhumbParserRULE_date             = 24
	RhumbParserRULE_text             = 25
	RhumbParserRULE_interpolation    = 26
	RhumbParserRULE_reference        = 27
)

// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_expressions
	return p
}

func InitEmptyExpressionsContext(p *ExpressionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_expressions
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionsContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionsContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (s *ExpressionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitExpressions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RhumbParserRULE_expressions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RhumbParserSemicolon || _la == RhumbParserNL {
		{
			p.SetState(56)
			p.Terminator()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(62)
		p.expression(0)
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RhumbParserSemicolon || _la == RhumbParserNL {
		{
			p.SetState(63)
			p.Terminator()
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9076863210788679164) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&658107) != 0) {
			{
				p.SetState(64)
				p.expression(0)
			}

		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldsContext is an interface to support dynamic dispatch.
type IFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext

	// IsFieldsContext differentiates from other interfaces.
	IsFieldsContext()
}

type FieldsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsContext() *FieldsContext {
	var p = new(FieldsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_fields
	return p
}

func InitEmptyFieldsContext(p *FieldsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_fields
}

func (*FieldsContext) IsFieldsContext() {}

func NewFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsContext {
	var p = new(FieldsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_fields

	return p
}

func (s *FieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *FieldsContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldsContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *FieldsContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *FieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterFields(s)
	}
}

func (s *FieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitFields(s)
	}
}

func (s *FieldsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitFields(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) Fields() (localctx IFieldsContext) {
	localctx = NewFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RhumbParserRULE_fields)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RhumbParserSemicolon || _la == RhumbParserNL {
		{
			p.SetState(72)
			p.Terminator()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
		p.Field()
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RhumbParserSemicolon || _la == RhumbParserNL {
		{
			p.SetState(79)
			p.Terminator()
		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9076863210788679156) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&658107) != 0) {
			{
				p.SetState(80)
				p.Field()
			}

		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPatternsContext is an interface to support dynamic dispatch.
type IPatternsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPattern() []IPatternContext
	Pattern(i int) IPatternContext
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext

	// IsPatternsContext differentiates from other interfaces.
	IsPatternsContext()
}

type PatternsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternsContext() *PatternsContext {
	var p = new(PatternsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_patterns
	return p
}

func InitEmptyPatternsContext(p *PatternsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_patterns
}

func (*PatternsContext) IsPatternsContext() {}

func NewPatternsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternsContext {
	var p = new(PatternsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_patterns

	return p
}

func (s *PatternsContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternsContext) AllPattern() []IPatternContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPatternContext); ok {
			len++
		}
	}

	tst := make([]IPatternContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPatternContext); ok {
			tst[i] = t.(IPatternContext)
			i++
		}
	}

	return tst
}

func (s *PatternsContext) Pattern(i int) IPatternContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *PatternsContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *PatternsContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *PatternsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterPatterns(s)
	}
}

func (s *PatternsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitPatterns(s)
	}
}

func (s *PatternsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitPatterns(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) Patterns() (localctx IPatternsContext) {
	localctx = NewPatternsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RhumbParserRULE_patterns)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RhumbParserSemicolon || _la == RhumbParserNL {
		{
			p.SetState(88)
			p.Terminator()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
		p.Pattern()
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RhumbParserSemicolon || _la == RhumbParserNL {
		{
			p.SetState(95)
			p.Terminator()
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9076863210788679164) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&658107) != 0) {
			{
				p.SetState(96)
				p.Pattern()
			}

		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITerminatorContext is an interface to support dynamic dispatch.
type ITerminatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Semicolon() antlr.TerminalNode
	NL() antlr.TerminalNode

	// IsTerminatorContext differentiates from other interfaces.
	IsTerminatorContext()
}

type TerminatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerminatorContext() *TerminatorContext {
	var p = new(TerminatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_terminator
	return p
}

func InitEmptyTerminatorContext(p *TerminatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_terminator
}

func (*TerminatorContext) IsTerminatorContext() {}

func NewTerminatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TerminatorContext {
	var p = new(TerminatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_terminator

	return p
}

func (s *TerminatorContext) GetParser() antlr.Parser { return s.parser }

func (s *TerminatorContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(RhumbParserSemicolon, 0)
}

func (s *TerminatorContext) NL() antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, 0)
}

func (s *TerminatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerminatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TerminatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterTerminator(s)
	}
}

func (s *TerminatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitTerminator(s)
	}
}

func (s *TerminatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitTerminator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) Terminator() (localctx ITerminatorContext) {
	localctx = NewTerminatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RhumbParserRULE_terminator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RhumbParserSemicolon || _la == RhumbParserNL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TextSymbolContext struct {
	LiteralContext
}

func NewTextSymbolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TextSymbolContext {
	var p = new(TextSymbolContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *TextSymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextSymbolContext) Text() ITextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *TextSymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterTextSymbol(s)
	}
}

func (s *TextSymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitTextSymbol(s)
	}
}

func (s *TextSymbolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitTextSymbol(s)

	default:
		return t.VisitChildren(s)
	}
}

type RationalNumberContext struct {
	LiteralContext
}

func NewRationalNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RationalNumberContext {
	var p = new(RationalNumberContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *RationalNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RationalNumberContext) FloatingPoint() antlr.TerminalNode {
	return s.GetToken(RhumbParserFloatingPoint, 0)
}

func (s *RationalNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterRationalNumber(s)
	}
}

func (s *RationalNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitRationalNumber(s)
	}
}

func (s *RationalNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitRationalNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type WholeNumberContext struct {
	LiteralContext
}

func NewWholeNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WholeNumberContext {
	var p = new(WholeNumberContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *WholeNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WholeNumberContext) NumberPart() antlr.TerminalNode {
	return s.GetToken(RhumbParserNumberPart, 0)
}

func (s *WholeNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterWholeNumber(s)
	}
}

func (s *WholeNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitWholeNumber(s)
	}
}

func (s *WholeNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitWholeNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReferenceLiteralContext struct {
	LiteralContext
}

func NewReferenceLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReferenceLiteralContext {
	var p = new(ReferenceLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *ReferenceLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceLiteralContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *ReferenceLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterReferenceLiteral(s)
	}
}

func (s *ReferenceLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitReferenceLiteral(s)
	}
}

func (s *ReferenceLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitReferenceLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type ZeroNumberContext struct {
	LiteralContext
}

func NewZeroNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ZeroNumberContext {
	var p = new(ZeroNumberContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *ZeroNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZeroNumberContext) Zero() antlr.TerminalNode {
	return s.GetToken(RhumbParserZero, 0)
}

func (s *ZeroNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterZeroNumber(s)
	}
}

func (s *ZeroNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitZeroNumber(s)
	}
}

func (s *ZeroNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitZeroNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type KeySymbolContext struct {
	LiteralContext
}

func NewKeySymbolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KeySymbolContext {
	var p = new(KeySymbolContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *KeySymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeySymbolContext) Key() antlr.TerminalNode {
	return s.GetToken(RhumbParserKey, 0)
}

func (s *KeySymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterKeySymbol(s)
	}
}

func (s *KeySymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitKeySymbol(s)
	}
}

func (s *KeySymbolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitKeySymbol(s)

	default:
		return t.VisitChildren(s)
	}
}

type LabelSymbolContext struct {
	LiteralContext
}

func NewLabelSymbolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LabelSymbolContext {
	var p = new(LabelSymbolContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LabelSymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelSymbolContext) Bang() antlr.TerminalNode {
	return s.GetToken(RhumbParserBang, 0)
}

func (s *LabelSymbolContext) Label() antlr.TerminalNode {
	return s.GetToken(RhumbParserLabel, 0)
}

func (s *LabelSymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterLabelSymbol(s)
	}
}

func (s *LabelSymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitLabelSymbol(s)
	}
}

func (s *LabelSymbolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitLabelSymbol(s)

	default:
		return t.VisitChildren(s)
	}
}

type DateNumberContext struct {
	LiteralContext
}

func NewDateNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateNumberContext {
	var p = new(DateNumberContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *DateNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateNumberContext) Date() IDateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateContext)
}

func (s *DateNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterDateNumber(s)
	}
}

func (s *DateNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitDateNumber(s)
	}
}

func (s *DateNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitDateNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RhumbParserRULE_literal)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRationalNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Match(RhumbParserFloatingPoint)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDateNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)
			p.Date()
		}

	case 3:
		localctx = NewZeroNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(108)
			p.Match(RhumbParserZero)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewWholeNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(109)
			p.Match(RhumbParserNumberPart)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewKeySymbolContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(110)
			p.Match(RhumbParserKey)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewTextSymbolContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(111)
			p.Text()
		}

	case 7:
		localctx = NewReferenceLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(112)
			p.Reference()
		}

	case 8:
		localctx = NewLabelSymbolContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(113)
			p.Match(RhumbParserBang)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewLabelSymbolContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(114)
			p.Match(RhumbParserLabel)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldLiteralContext is an interface to support dynamic dispatch.
type IFieldLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Zero() antlr.TerminalNode
	NumberPart() antlr.TerminalNode
	Key() antlr.TerminalNode
	Text() ITextContext
	Reference() IReferenceContext
	Bang() antlr.TerminalNode
	Label() antlr.TerminalNode

	// IsFieldLiteralContext differentiates from other interfaces.
	IsFieldLiteralContext()
}

type FieldLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldLiteralContext() *FieldLiteralContext {
	var p = new(FieldLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_fieldLiteral
	return p
}

func InitEmptyFieldLiteralContext(p *FieldLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_fieldLiteral
}

func (*FieldLiteralContext) IsFieldLiteralContext() {}

func NewFieldLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldLiteralContext {
	var p = new(FieldLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_fieldLiteral

	return p
}

func (s *FieldLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldLiteralContext) Zero() antlr.TerminalNode {
	return s.GetToken(RhumbParserZero, 0)
}

func (s *FieldLiteralContext) NumberPart() antlr.TerminalNode {
	return s.GetToken(RhumbParserNumberPart, 0)
}

func (s *FieldLiteralContext) Key() antlr.TerminalNode {
	return s.GetToken(RhumbParserKey, 0)
}

func (s *FieldLiteralContext) Text() ITextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *FieldLiteralContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *FieldLiteralContext) Bang() antlr.TerminalNode {
	return s.GetToken(RhumbParserBang, 0)
}

func (s *FieldLiteralContext) Label() antlr.TerminalNode {
	return s.GetToken(RhumbParserLabel, 0)
}

func (s *FieldLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterFieldLiteral(s)
	}
}

func (s *FieldLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitFieldLiteral(s)
	}
}

func (s *FieldLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitFieldLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) FieldLiteral() (localctx IFieldLiteralContext) {
	localctx = NewFieldLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RhumbParserRULE_fieldLiteral)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Match(RhumbParserZero)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Match(RhumbParserNumberPart)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(119)
			p.Match(RhumbParserKey)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(120)
			p.Text()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(121)
			p.Reference()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(122)
			p.Match(RhumbParserBang)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(123)
			p.Match(RhumbParserLabel)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConjunctiveContext struct {
	ExpressionContext
}

func NewConjunctiveContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConjunctiveContext {
	var p = new(ConjunctiveContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ConjunctiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConjunctiveContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConjunctiveContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConjunctiveContext) ConjunctiveOp() IConjunctiveOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConjunctiveOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConjunctiveOpContext)
}

func (s *ConjunctiveContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *ConjunctiveContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *ConjunctiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterConjunctive(s)
	}
}

func (s *ConjunctiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitConjunctive(s)
	}
}

func (s *ConjunctiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitConjunctive(s)

	default:
		return t.VisitChildren(s)
	}
}

type AccessContext struct {
	ExpressionContext
}

func NewAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccessContext {
	var p = new(AccessContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AccessContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenBracket, 0)
}

func (s *AccessContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseBracket, 0)
}

func (s *AccessContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *AccessContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *AccessContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *AccessContext) AccessOp() IAccessOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessOpContext)
}

func (s *AccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAccess(s)
	}
}

func (s *AccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAccess(s)
	}
}

func (s *AccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type ApplicativeContext struct {
	ExpressionContext
}

func NewApplicativeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ApplicativeContext {
	var p = new(ApplicativeContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ApplicativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplicativeContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ApplicativeContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ApplicativeContext) ApplicativeOp() IApplicativeOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApplicativeOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApplicativeOpContext)
}

func (s *ApplicativeContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *ApplicativeContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *ApplicativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterApplicative(s)
	}
}

func (s *ApplicativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitApplicative(s)
	}
}

func (s *ApplicativeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitApplicative(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConditionalContext struct {
	ExpressionContext
}

func NewConditionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionalContext {
	var p = new(ConditionalContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalContext) ConditionalOp() IConditionalOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalOpContext)
}

func (s *ConditionalContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *ConditionalContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (s *ConditionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitConditional(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrefixContext struct {
	ExpressionContext
}

func NewPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrefixContext {
	var p = new(PrefixContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixContext) PrefixOp() IPrefixOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefixOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefixOpContext)
}

func (s *PrefixContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterPrefix(s)
	}
}

func (s *PrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitPrefix(s)
	}
}

func (s *PrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparativeContext struct {
	ExpressionContext
}

func NewComparativeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparativeContext {
	var p = new(ComparativeContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ComparativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparativeContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ComparativeContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ComparativeContext) ComparativeOp() IComparativeOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparativeOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparativeOpContext)
}

func (s *ComparativeContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *ComparativeContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *ComparativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterComparative(s)
	}
}

func (s *ComparativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitComparative(s)
	}
}

func (s *ComparativeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitComparative(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpleExpressionContext struct {
	ExpressionContext
}

func NewSimpleExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleExpressionContext {
	var p = new(SimpleExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SimpleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleExpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *SimpleExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterSimpleExpression(s)
	}
}

func (s *SimpleExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitSimpleExpression(s)
	}
}

func (s *SimpleExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitSimpleExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiplicativeContext struct {
	ExpressionContext
}

func NewMultiplicativeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeContext {
	var p = new(MultiplicativeContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MultiplicativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultiplicativeContext) MultiplicativeOp() IMultiplicativeOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeOpContext)
}

func (s *MultiplicativeContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *MultiplicativeContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *MultiplicativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterMultiplicative(s)
	}
}

func (s *MultiplicativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitMultiplicative(s)
	}
}

func (s *MultiplicativeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitMultiplicative(s)

	default:
		return t.VisitChildren(s)
	}
}

type AdditiveContext struct {
	ExpressionContext
}

func NewAdditiveContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveContext {
	var p = new(AdditiveContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AdditiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AdditiveContext) AdditiveOp() IAdditiveOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveOpContext)
}

func (s *AdditiveContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *AdditiveContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *AdditiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAdditive(s)
	}
}

func (s *AdditiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAdditive(s)
	}
}

func (s *AdditiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAdditive(s)

	default:
		return t.VisitChildren(s)
	}
}

type InvocationContext struct {
	ExpressionContext
}

func NewInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvocationContext {
	var p = new(InvocationContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InvocationContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenParen, 0)
}

func (s *InvocationContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseParen, 0)
}

func (s *InvocationContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *InvocationContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *InvocationContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *InvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterInvocation(s)
	}
}

func (s *InvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitInvocation(s)
	}
}

func (s *InvocationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitInvocation(s)

	default:
		return t.VisitChildren(s)
	}
}

type LibraryContext struct {
	ExpressionContext
}

func NewLibraryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LibraryContext {
	var p = new(LibraryContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LibraryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LibraryContext) OpenCurly() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenCurly, 0)
}

func (s *LibraryContext) Path() antlr.TerminalNode {
	return s.GetToken(RhumbParserPath, 0)
}

func (s *LibraryContext) Version() antlr.TerminalNode {
	return s.GetToken(RhumbParserVersion, 0)
}

func (s *LibraryContext) CloseCurly() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseCurly, 0)
}

func (s *LibraryContext) Label() antlr.TerminalNode {
	return s.GetToken(RhumbParserLabel, 0)
}

func (s *LibraryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterLibrary(s)
	}
}

func (s *LibraryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitLibrary(s)
	}
}

func (s *LibraryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitLibrary(s)

	default:
		return t.VisitChildren(s)
	}
}

type RoutineContext struct {
	ExpressionContext
}

func NewRoutineContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoutineContext {
	var p = new(RoutineContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RoutineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoutineContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenParen, 0)
}

func (s *RoutineContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseParen, 0)
}

func (s *RoutineContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *RoutineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterRoutine(s)
	}
}

func (s *RoutineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitRoutine(s)
	}
}

func (s *RoutineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitRoutine(s)

	default:
		return t.VisitChildren(s)
	}
}

type DisjunctiveContext struct {
	ExpressionContext
}

func NewDisjunctiveContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DisjunctiveContext {
	var p = new(DisjunctiveContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DisjunctiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctiveContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *DisjunctiveContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DisjunctiveContext) DisjunctiveOp() IDisjunctiveOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDisjunctiveOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDisjunctiveOpContext)
}

func (s *DisjunctiveContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *DisjunctiveContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *DisjunctiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterDisjunctive(s)
	}
}

func (s *DisjunctiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitDisjunctive(s)
	}
}

func (s *DisjunctiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitDisjunctive(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentityContext struct {
	ExpressionContext
}

func NewIdentityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentityContext {
	var p = new(IdentityContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdentityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentityContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IdentityContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IdentityContext) IdentityOp() IIdentityOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentityOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentityOpContext)
}

func (s *IdentityContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *IdentityContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *IdentityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterIdentity(s)
	}
}

func (s *IdentityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitIdentity(s)
	}
}

func (s *IdentityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitIdentity(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignLabelContext struct {
	ExpressionContext
}

func NewAssignLabelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignLabelContext {
	var p = new(AssignLabelContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AssignLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignLabelContext) AssignmentOp() IAssignmentOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentOpContext)
}

func (s *AssignLabelContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignLabelContext) ChainExpression() IChainExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChainExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChainExpressionContext)
}

func (s *AssignLabelContext) FieldLiteral() IFieldLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLiteralContext)
}

func (s *AssignLabelContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *AssignLabelContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *AssignLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAssignLabel(s)
	}
}

func (s *AssignLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAssignLabel(s)
	}
}

func (s *AssignLabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAssignLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

type EffectContext struct {
	ExpressionContext
}

func NewEffectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EffectContext {
	var p = new(EffectContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *EffectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EffectContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EffectContext) OpenCurly() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenCurly, 0)
}

func (s *EffectContext) CloseCurly() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseCurly, 0)
}

func (s *EffectContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *EffectContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *EffectContext) Patterns() IPatternsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternsContext)
}

func (s *EffectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterEffect(s)
	}
}

func (s *EffectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitEffect(s)
	}
}

func (s *EffectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitEffect(s)

	default:
		return t.VisitChildren(s)
	}
}

type MemberContext struct {
	ExpressionContext
}

func NewMemberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberContext {
	var p = new(MemberContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberContext) ChainExpression() IChainExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChainExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChainExpressionContext)
}

func (s *MemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterMember(s)
	}
}

func (s *MemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitMember(s)
	}
}

func (s *MemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitMember(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelectorContext struct {
	ExpressionContext
}

func NewSelectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectorContext {
	var p = new(SelectorContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) OpenCurly() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenCurly, 0)
}

func (s *SelectorContext) CloseCurly() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseCurly, 0)
}

func (s *SelectorContext) Patterns() IPatternsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternsContext)
}

func (s *SelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (s *SelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowerContext struct {
	ExpressionContext
}

func NewPowerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerContext {
	var p = new(PowerContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PowerContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PowerContext) ExponentiationOp() IExponentiationOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExponentiationOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExponentiationOpContext)
}

func (s *PowerContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *PowerContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *PowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterPower(s)
	}
}

func (s *PowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitPower(s)
	}
}

func (s *PowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitPower(s)

	default:
		return t.VisitChildren(s)
	}
}

type MapContext struct {
	ExpressionContext
}

func NewMapContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapContext {
	var p = new(MapContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenBracket, 0)
}

func (s *MapContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseBracket, 0)
}

func (s *MapContext) Fields() IFieldsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsContext)
}

func (s *MapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterMap(s)
	}
}

func (s *MapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitMap(s)
	}
}

func (s *MapContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitMap(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *RhumbParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, RhumbParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMapContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(127)
			p.Match(RhumbParserOpenBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9067856011533938164) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&666299) != 0) {
			{
				p.SetState(128)
				p.Fields()
			}

		}
		{
			p.SetState(131)
			p.Match(RhumbParserCloseBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewLibraryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(132)
			p.Match(RhumbParserOpenCurly)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RhumbParserLabel {
			{
				p.SetState(133)
				p.Match(RhumbParserLabel)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(136)
			p.Match(RhumbParserPath)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Match(RhumbParserVersion)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(RhumbParserCloseCurly)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewSelectorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(139)
			p.Match(RhumbParserOpenCurly)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9067856011533938172) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&666299) != 0) {
			{
				p.SetState(140)
				p.Patterns()
			}

		}
		{
			p.SetState(143)
			p.Match(RhumbParserCloseCurly)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewRoutineContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(144)
			p.Match(RhumbParserOpenParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9067856011533938172) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&666299) != 0) {
			{
				p.SetState(145)
				p.Expressions()
			}

		}
		{
			p.SetState(148)
			p.Match(RhumbParserCloseParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewMemberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(149)
			p.ChainExpression()
		}

	case 6:
		localctx = NewPrefixContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(150)
			p.PrefixOp()
		}
		{
			p.SetState(151)
			p.expression(11)
		}

	case 7:
		localctx = NewAssignLabelContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(153)
				p.ChainExpression()
			}

		case 2:
			{
				p.SetState(154)
				p.FieldLiteral()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserNL {
			{
				p.SetState(157)
				p.Match(RhumbParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(162)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(163)
			p.AssignmentOp()
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserNL {
			{
				p.SetState(164)
				p.Match(RhumbParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(169)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(170)
			p.expression(2)
		}

	case 8:
		localctx = NewSimpleExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(172)
			p.Literal()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(356)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowerContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(175)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				p.SetState(179)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(176)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(181)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(182)
					p.ExponentiationOp()
				}
				p.SetState(186)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(183)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(188)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(189)
					p.expression(12)
				}

			case 2:
				localctx = NewMultiplicativeContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(191)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				p.SetState(195)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(192)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(197)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(198)
					p.MultiplicativeOp()
				}
				p.SetState(202)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(199)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(204)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(205)
					p.expression(11)
				}

			case 3:
				localctx = NewAdditiveContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(207)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				p.SetState(211)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(208)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(213)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(214)
					p.AdditiveOp()
				}
				p.SetState(218)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(215)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(220)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(221)
					p.expression(10)
				}

			case 4:
				localctx = NewComparativeContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(223)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				p.SetState(227)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(224)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(229)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(230)
					p.ComparativeOp()
				}
				p.SetState(234)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(231)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(236)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(237)
					p.expression(9)
				}

			case 5:
				localctx = NewIdentityContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(239)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				p.SetState(243)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(240)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(245)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(246)
					p.IdentityOp()
				}
				p.SetState(250)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(247)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(252)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(253)
					p.expression(8)
				}

			case 6:
				localctx = NewConjunctiveContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(255)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				p.SetState(259)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(256)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(261)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(262)
					p.ConjunctiveOp()
				}
				p.SetState(266)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(263)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(268)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(269)
					p.expression(7)
				}

			case 7:
				localctx = NewDisjunctiveContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(271)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				p.SetState(275)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(272)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(277)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(278)
					p.DisjunctiveOp()
				}
				p.SetState(282)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(279)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(284)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(285)
					p.expression(6)
				}

			case 8:
				localctx = NewApplicativeContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(287)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				p.SetState(291)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(288)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(293)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(294)
					p.ApplicativeOp()
				}
				p.SetState(298)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(295)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(300)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(301)
					p.expression(4)
				}

			case 9:
				localctx = NewConditionalContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(303)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				p.SetState(307)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(304)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(309)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(310)
					p.ConditionalOp()
				}
				p.SetState(314)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserNL {
					{
						p.SetState(311)
						p.Match(RhumbParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(316)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(317)
					p.expression(3)
				}

			case 10:
				localctx = NewInvocationContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(319)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(320)
					p.Match(RhumbParserOpenParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(324)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
				for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					if _alt == 1 {
						{
							p.SetState(321)
							p.Terminator()
						}

					}
					p.SetState(326)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
					if p.HasError() {
						goto errorExit
					}
				}
				p.SetState(328)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9067856011533938172) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&666299) != 0) {
					{
						p.SetState(327)
						p.Expressions()
					}

				}
				{
					p.SetState(330)
					p.Match(RhumbParserCloseParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 11:
				localctx = NewAccessContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(331)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(332)
					p.Match(RhumbParserOpenBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(336)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
				for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					if _alt == 1 {
						{
							p.SetState(333)
							p.Terminator()
						}

					}
					p.SetState(338)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
					if p.HasError() {
						goto errorExit
					}
				}
				p.SetState(341)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(339)
						p.Expressions()
					}

				} else if p.HasError() { // JIM
					goto errorExit
				} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) == 2 {
					{
						p.SetState(340)
						p.AccessOp()
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}
				{
					p.SetState(343)
					p.Match(RhumbParserCloseBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 12:
				localctx = NewEffectContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(344)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(345)
					p.Match(RhumbParserOpenCurly)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(349)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
				for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					if _alt == 1 {
						{
							p.SetState(346)
							p.Terminator()
						}

					}
					p.SetState(351)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
					if p.HasError() {
						goto errorExit
					}
				}
				p.SetState(353)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9067856011533938172) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&666299) != 0) {
					{
						p.SetState(352)
						p.Patterns()
					}

				}
				{
					p.SetState(355)
					p.Match(RhumbParserCloseCurly)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IChainExpressionContext is an interface to support dynamic dispatch.
type IChainExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFieldLiteral() []IFieldLiteralContext
	FieldLiteral(i int) IFieldLiteralContext
	AllChainOp() []IChainOpContext
	ChainOp(i int) IChainOpContext
	AllOpenBracket() []antlr.TerminalNode
	OpenBracket(i int) antlr.TerminalNode
	AllCloseBracket() []antlr.TerminalNode
	CloseBracket(i int) antlr.TerminalNode
	AllOpenParen() []antlr.TerminalNode
	OpenParen(i int) antlr.TerminalNode
	AllCloseParen() []antlr.TerminalNode
	CloseParen(i int) antlr.TerminalNode
	AllPrefixOp() []IPrefixOpContext
	PrefixOp(i int) IPrefixOpContext
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllAccessOp() []IAccessOpContext
	AccessOp(i int) IAccessOpContext
	AllExpressions() []IExpressionsContext
	Expressions(i int) IExpressionsContext

	// IsChainExpressionContext differentiates from other interfaces.
	IsChainExpressionContext()
}

type ChainExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChainExpressionContext() *ChainExpressionContext {
	var p = new(ChainExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_chainExpression
	return p
}

func InitEmptyChainExpressionContext(p *ChainExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_chainExpression
}

func (*ChainExpressionContext) IsChainExpressionContext() {}

func NewChainExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChainExpressionContext {
	var p = new(ChainExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_chainExpression

	return p
}

func (s *ChainExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ChainExpressionContext) AllFieldLiteral() []IFieldLiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldLiteralContext); ok {
			len++
		}
	}

	tst := make([]IFieldLiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldLiteralContext); ok {
			tst[i] = t.(IFieldLiteralContext)
			i++
		}
	}

	return tst
}

func (s *ChainExpressionContext) FieldLiteral(i int) IFieldLiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLiteralContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLiteralContext)
}

func (s *ChainExpressionContext) AllChainOp() []IChainOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IChainOpContext); ok {
			len++
		}
	}

	tst := make([]IChainOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IChainOpContext); ok {
			tst[i] = t.(IChainOpContext)
			i++
		}
	}

	return tst
}

func (s *ChainExpressionContext) ChainOp(i int) IChainOpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChainOpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChainOpContext)
}

func (s *ChainExpressionContext) AllOpenBracket() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserOpenBracket)
}

func (s *ChainExpressionContext) OpenBracket(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenBracket, i)
}

func (s *ChainExpressionContext) AllCloseBracket() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserCloseBracket)
}

func (s *ChainExpressionContext) CloseBracket(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseBracket, i)
}

func (s *ChainExpressionContext) AllOpenParen() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserOpenParen)
}

func (s *ChainExpressionContext) OpenParen(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenParen, i)
}

func (s *ChainExpressionContext) AllCloseParen() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserCloseParen)
}

func (s *ChainExpressionContext) CloseParen(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseParen, i)
}

func (s *ChainExpressionContext) AllPrefixOp() []IPrefixOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrefixOpContext); ok {
			len++
		}
	}

	tst := make([]IPrefixOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrefixOpContext); ok {
			tst[i] = t.(IPrefixOpContext)
			i++
		}
	}

	return tst
}

func (s *ChainExpressionContext) PrefixOp(i int) IPrefixOpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefixOpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefixOpContext)
}

func (s *ChainExpressionContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *ChainExpressionContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *ChainExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ChainExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ChainExpressionContext) AllAccessOp() []IAccessOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAccessOpContext); ok {
			len++
		}
	}

	tst := make([]IAccessOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAccessOpContext); ok {
			tst[i] = t.(IAccessOpContext)
			i++
		}
	}

	return tst
}

func (s *ChainExpressionContext) AccessOp(i int) IAccessOpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessOpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessOpContext)
}

func (s *ChainExpressionContext) AllExpressions() []IExpressionsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionsContext); ok {
			len++
		}
	}

	tst := make([]IExpressionsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionsContext); ok {
			tst[i] = t.(IExpressionsContext)
			i++
		}
	}

	return tst
}

func (s *ChainExpressionContext) Expressions(i int) IExpressionsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *ChainExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChainExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterChainExpression(s)
	}
}

func (s *ChainExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitChainExpression(s)
	}
}

func (s *ChainExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitChainExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) ChainExpression() (localctx IChainExpressionContext) {
	localctx = NewChainExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RhumbParserRULE_chainExpression)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		p.FieldLiteral()
	}
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(391)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case RhumbParserAt, RhumbParserAtAt, RhumbParserBSlash, RhumbParserBSlashBSlash:
				{
					p.SetState(362)
					p.ChainOp()
				}
				p.SetState(364)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&146508826066096132) != 0 {
					{
						p.SetState(363)
						p.PrefixOp()
					}

				}
				{
					p.SetState(366)
					p.FieldLiteral()
				}

			case RhumbParserOpenBracket:
				{
					p.SetState(368)
					p.Match(RhumbParserOpenBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(372)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserSemicolon || _la == RhumbParserNL {
					{
						p.SetState(369)
						p.Terminator()
					}

					p.SetState(374)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				p.SetState(377)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(375)
						p.expression(0)
					}

				} else if p.HasError() { // JIM
					goto errorExit
				} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) == 2 {
					{
						p.SetState(376)
						p.AccessOp()
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}
				{
					p.SetState(379)
					p.Match(RhumbParserCloseBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case RhumbParserOpenParen:
				{
					p.SetState(380)
					p.Match(RhumbParserOpenParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(384)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
				for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					if _alt == 1 {
						{
							p.SetState(381)
							p.Terminator()
						}

					}
					p.SetState(386)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
					if p.HasError() {
						goto errorExit
					}
				}
				p.SetState(388)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9067856011533938172) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&666299) != 0) {
					{
						p.SetState(387)
						p.Expressions()
					}

				}
				{
					p.SetState(390)
					p.Match(RhumbParserCloseParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) CopyAll(ctx *FieldContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SimpleFieldContext struct {
	FieldContext
}

func NewSimpleFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleFieldContext {
	var p = new(SimpleFieldContext)

	InitEmptyFieldContext(&p.FieldContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldContext))

	return p
}

func (s *SimpleFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleFieldContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SimpleFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterSimpleField(s)
	}
}

func (s *SimpleFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitSimpleField(s)
	}
}

func (s *SimpleFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitSimpleField(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignMutSubFieldContext struct {
	FieldContext
}

func NewAssignMutSubFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignMutSubFieldContext {
	var p = new(AssignMutSubFieldContext)

	InitEmptyFieldContext(&p.FieldContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldContext))

	return p
}

func (s *AssignMutSubFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignMutSubFieldContext) At() antlr.TerminalNode {
	return s.GetToken(RhumbParserAt, 0)
}

func (s *AssignMutSubFieldContext) FieldLiteral() IFieldLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLiteralContext)
}

func (s *AssignMutSubFieldContext) ColonColon() antlr.TerminalNode {
	return s.GetToken(RhumbParserColonColon, 0)
}

func (s *AssignMutSubFieldContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignMutSubFieldContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *AssignMutSubFieldContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *AssignMutSubFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAssignMutSubField(s)
	}
}

func (s *AssignMutSubFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAssignMutSubField(s)
	}
}

func (s *AssignMutSubFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAssignMutSubField(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrefixAssignImmSubFieldContext struct {
	FieldContext
}

func NewPrefixAssignImmSubFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrefixAssignImmSubFieldContext {
	var p = new(PrefixAssignImmSubFieldContext)

	InitEmptyFieldContext(&p.FieldContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldContext))

	return p
}

func (s *PrefixAssignImmSubFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixAssignImmSubFieldContext) Colon() antlr.TerminalNode {
	return s.GetToken(RhumbParserColon, 0)
}

func (s *PrefixAssignImmSubFieldContext) FieldLiteral() IFieldLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLiteralContext)
}

func (s *PrefixAssignImmSubFieldContext) At() antlr.TerminalNode {
	return s.GetToken(RhumbParserAt, 0)
}

func (s *PrefixAssignImmSubFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterPrefixAssignImmSubField(s)
	}
}

func (s *PrefixAssignImmSubFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitPrefixAssignImmSubField(s)
	}
}

func (s *PrefixAssignImmSubFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitPrefixAssignImmSubField(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrefixAssignMutSubFieldContext struct {
	FieldContext
}

func NewPrefixAssignMutSubFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrefixAssignMutSubFieldContext {
	var p = new(PrefixAssignMutSubFieldContext)

	InitEmptyFieldContext(&p.FieldContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldContext))

	return p
}

func (s *PrefixAssignMutSubFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixAssignMutSubFieldContext) Dot() antlr.TerminalNode {
	return s.GetToken(RhumbParserDot, 0)
}

func (s *PrefixAssignMutSubFieldContext) At() antlr.TerminalNode {
	return s.GetToken(RhumbParserAt, 0)
}

func (s *PrefixAssignMutSubFieldContext) FieldLiteral() IFieldLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLiteralContext)
}

func (s *PrefixAssignMutSubFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterPrefixAssignMutSubField(s)
	}
}

func (s *PrefixAssignMutSubFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitPrefixAssignMutSubField(s)
	}
}

func (s *PrefixAssignMutSubFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitPrefixAssignMutSubField(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignMutFieldContext struct {
	FieldContext
}

func NewAssignMutFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignMutFieldContext {
	var p = new(AssignMutFieldContext)

	InitEmptyFieldContext(&p.FieldContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldContext))

	return p
}

func (s *AssignMutFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignMutFieldContext) FieldLiteral() IFieldLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLiteralContext)
}

func (s *AssignMutFieldContext) ColonColon() antlr.TerminalNode {
	return s.GetToken(RhumbParserColonColon, 0)
}

func (s *AssignMutFieldContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignMutFieldContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *AssignMutFieldContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *AssignMutFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAssignMutField(s)
	}
}

func (s *AssignMutFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAssignMutField(s)
	}
}

func (s *AssignMutFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAssignMutField(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrefixSlurpSpreadContext struct {
	FieldContext
}

func NewPrefixSlurpSpreadContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrefixSlurpSpreadContext {
	var p = new(PrefixSlurpSpreadContext)

	InitEmptyFieldContext(&p.FieldContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldContext))

	return p
}

func (s *PrefixSlurpSpreadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixSlurpSpreadContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(RhumbParserAmpersand, 0)
}

func (s *PrefixSlurpSpreadContext) FieldLiteral() IFieldLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLiteralContext)
}

func (s *PrefixSlurpSpreadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterPrefixSlurpSpread(s)
	}
}

func (s *PrefixSlurpSpreadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitPrefixSlurpSpread(s)
	}
}

func (s *PrefixSlurpSpreadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitPrefixSlurpSpread(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrefixAssignMutFieldContext struct {
	FieldContext
}

func NewPrefixAssignMutFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrefixAssignMutFieldContext {
	var p = new(PrefixAssignMutFieldContext)

	InitEmptyFieldContext(&p.FieldContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldContext))

	return p
}

func (s *PrefixAssignMutFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixAssignMutFieldContext) Dot() antlr.TerminalNode {
	return s.GetToken(RhumbParserDot, 0)
}

func (s *PrefixAssignMutFieldContext) FieldLiteral() IFieldLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLiteralContext)
}

func (s *PrefixAssignMutFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterPrefixAssignMutField(s)
	}
}

func (s *PrefixAssignMutFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitPrefixAssignMutField(s)
	}
}

func (s *PrefixAssignMutFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitPrefixAssignMutField(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignImmFieldContext struct {
	FieldContext
}

func NewAssignImmFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignImmFieldContext {
	var p = new(AssignImmFieldContext)

	InitEmptyFieldContext(&p.FieldContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldContext))

	return p
}

func (s *AssignImmFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignImmFieldContext) FieldLiteral() IFieldLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLiteralContext)
}

func (s *AssignImmFieldContext) DotDot() antlr.TerminalNode {
	return s.GetToken(RhumbParserDotDot, 0)
}

func (s *AssignImmFieldContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignImmFieldContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *AssignImmFieldContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *AssignImmFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAssignImmField(s)
	}
}

func (s *AssignImmFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAssignImmField(s)
	}
}

func (s *AssignImmFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAssignImmField(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignImmSubFieldContext struct {
	FieldContext
}

func NewAssignImmSubFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignImmSubFieldContext {
	var p = new(AssignImmSubFieldContext)

	InitEmptyFieldContext(&p.FieldContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldContext))

	return p
}

func (s *AssignImmSubFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignImmSubFieldContext) At() antlr.TerminalNode {
	return s.GetToken(RhumbParserAt, 0)
}

func (s *AssignImmSubFieldContext) FieldLiteral() IFieldLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLiteralContext)
}

func (s *AssignImmSubFieldContext) DotDot() antlr.TerminalNode {
	return s.GetToken(RhumbParserDotDot, 0)
}

func (s *AssignImmSubFieldContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignImmSubFieldContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *AssignImmSubFieldContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *AssignImmSubFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAssignImmSubField(s)
	}
}

func (s *AssignImmSubFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAssignImmSubField(s)
	}
}

func (s *AssignImmSubFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAssignImmSubField(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpleMapFieldContext struct {
	FieldContext
}

func NewSimpleMapFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleMapFieldContext {
	var p = new(SimpleMapFieldContext)

	InitEmptyFieldContext(&p.FieldContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldContext))

	return p
}

func (s *SimpleMapFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleMapFieldContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenBracket, 0)
}

func (s *SimpleMapFieldContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseBracket, 0)
}

func (s *SimpleMapFieldContext) Fields() IFieldsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsContext)
}

func (s *SimpleMapFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterSimpleMapField(s)
	}
}

func (s *SimpleMapFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitSimpleMapField(s)
	}
}

func (s *SimpleMapFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitSimpleMapField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RhumbParserRULE_field)
	var _la int

	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPrefixAssignMutFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(395)
			p.Match(RhumbParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.FieldLiteral()
		}

	case 2:
		localctx = NewPrefixAssignMutSubFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(397)
			p.Match(RhumbParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.Match(RhumbParserAt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)
			p.FieldLiteral()
		}

	case 3:
		localctx = NewPrefixAssignImmSubFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(400)
			p.Match(RhumbParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
			p.FieldLiteral()
		}

	case 4:
		localctx = NewPrefixAssignImmSubFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(402)
			p.Match(RhumbParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.Match(RhumbParserAt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.FieldLiteral()
		}

	case 5:
		localctx = NewPrefixSlurpSpreadContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(405)
			p.Match(RhumbParserAmpersand)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.FieldLiteral()
		}

	case 6:
		localctx = NewAssignMutFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(407)
			p.FieldLiteral()
		}
		p.SetState(411)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserNL {
			{
				p.SetState(408)
				p.Match(RhumbParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(413)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(414)
			p.Match(RhumbParserColonColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(418)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserNL {
			{
				p.SetState(415)
				p.Match(RhumbParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(420)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(421)
			p.expression(0)
		}

	case 7:
		localctx = NewAssignMutSubFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(423)
			p.Match(RhumbParserAt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.FieldLiteral()
		}
		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserNL {
			{
				p.SetState(425)
				p.Match(RhumbParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(430)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(431)
			p.Match(RhumbParserColonColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(435)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserNL {
			{
				p.SetState(432)
				p.Match(RhumbParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(437)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(438)
			p.expression(0)
		}

	case 8:
		localctx = NewAssignImmFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(440)
			p.FieldLiteral()
		}
		p.SetState(444)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserNL {
			{
				p.SetState(441)
				p.Match(RhumbParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(446)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(447)
			p.Match(RhumbParserDotDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(451)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserNL {
			{
				p.SetState(448)
				p.Match(RhumbParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(453)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(454)
			p.expression(0)
		}

	case 9:
		localctx = NewAssignImmSubFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(456)
			p.Match(RhumbParserAt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.FieldLiteral()
		}
		p.SetState(461)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserNL {
			{
				p.SetState(458)
				p.Match(RhumbParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(463)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(464)
			p.Match(RhumbParserDotDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserNL {
			{
				p.SetState(465)
				p.Match(RhumbParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(470)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(471)
			p.expression(0)
		}

	case 10:
		localctx = NewSimpleMapFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(473)
			p.Match(RhumbParserOpenBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9067856011533938164) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&666299) != 0) {
			{
				p.SetState(474)
				p.Fields()
			}

		}
		{
			p.SetState(477)
			p.Match(RhumbParserCloseBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewSimpleFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(478)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPatternContext is an interface to support dynamic dispatch.
type IPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPatternContext differentiates from other interfaces.
	IsPatternContext()
}

type PatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternContext() *PatternContext {
	var p = new(PatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_pattern
	return p
}

func InitEmptyPatternContext(p *PatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_pattern
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) CopyAll(ctx *PatternContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignBreakingPatternContext struct {
	PatternContext
}

func NewAssignBreakingPatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignBreakingPatternContext {
	var p = new(AssignBreakingPatternContext)

	InitEmptyPatternContext(&p.PatternContext)
	p.parser = parser
	p.CopyAll(ctx.(*PatternContext))

	return p
}

func (s *AssignBreakingPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignBreakingPatternContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignBreakingPatternContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignBreakingPatternContext) DotDot() antlr.TerminalNode {
	return s.GetToken(RhumbParserDotDot, 0)
}

func (s *AssignBreakingPatternContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *AssignBreakingPatternContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *AssignBreakingPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAssignBreakingPattern(s)
	}
}

func (s *AssignBreakingPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAssignBreakingPattern(s)
	}
}

func (s *AssignBreakingPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAssignBreakingPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignDefaultPatternContext struct {
	PatternContext
}

func NewAssignDefaultPatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignDefaultPatternContext {
	var p = new(AssignDefaultPatternContext)

	InitEmptyPatternContext(&p.PatternContext)
	p.parser = parser
	p.CopyAll(ctx.(*PatternContext))

	return p
}

func (s *AssignDefaultPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignDefaultPatternContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignDefaultPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAssignDefaultPattern(s)
	}
}

func (s *AssignDefaultPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAssignDefaultPattern(s)
	}
}

func (s *AssignDefaultPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAssignDefaultPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignFallthroughPatternContext struct {
	PatternContext
}

func NewAssignFallthroughPatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignFallthroughPatternContext {
	var p = new(AssignFallthroughPatternContext)

	InitEmptyPatternContext(&p.PatternContext)
	p.parser = parser
	p.CopyAll(ctx.(*PatternContext))

	return p
}

func (s *AssignFallthroughPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignFallthroughPatternContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignFallthroughPatternContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignFallthroughPatternContext) ColonColon() antlr.TerminalNode {
	return s.GetToken(RhumbParserColonColon, 0)
}

func (s *AssignFallthroughPatternContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserNL)
}

func (s *AssignFallthroughPatternContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserNL, i)
}

func (s *AssignFallthroughPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAssignFallthroughPattern(s)
	}
}

func (s *AssignFallthroughPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAssignFallthroughPattern(s)
	}
}

func (s *AssignFallthroughPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAssignFallthroughPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) Pattern() (localctx IPatternContext) {
	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RhumbParserRULE_pattern)
	var _la int

	p.SetState(514)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignBreakingPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(481)
			p.expression(0)
		}
		p.SetState(485)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserNL {
			{
				p.SetState(482)
				p.Match(RhumbParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(487)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(488)
			p.Match(RhumbParserDotDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(492)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserNL {
			{
				p.SetState(489)
				p.Match(RhumbParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(494)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(495)
			p.expression(0)
		}

	case 2:
		localctx = NewAssignFallthroughPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(497)
			p.expression(0)
		}
		p.SetState(501)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserNL {
			{
				p.SetState(498)
				p.Match(RhumbParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(503)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(504)
			p.Match(RhumbParserColonColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(508)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserNL {
			{
				p.SetState(505)
				p.Match(RhumbParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(510)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(511)
			p.expression(0)
		}

	case 3:
		localctx = NewAssignDefaultPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(513)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccessOpContext is an interface to support dynamic dispatch.
type IAccessOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAccessOpContext differentiates from other interfaces.
	IsAccessOpContext()
}

type AccessOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessOpContext() *AccessOpContext {
	var p = new(AccessOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_accessOp
	return p
}

func InitEmptyAccessOpContext(p *AccessOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_accessOp
}

func (*AccessOpContext) IsAccessOpContext() {}

func NewAccessOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessOpContext {
	var p = new(AccessOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_accessOp

	return p
}

func (s *AccessOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessOpContext) CopyAll(ctx *AccessOpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AccessOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AllFieldsContext struct {
	AccessOpContext
}

func NewAllFieldsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllFieldsContext {
	var p = new(AllFieldsContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *AllFieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllFieldsContext) Star() antlr.TerminalNode {
	return s.GetToken(RhumbParserStar, 0)
}

func (s *AllFieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAllFields(s)
	}
}

func (s *AllFieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAllFields(s)
	}
}

func (s *AllFieldsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAllFields(s)

	default:
		return t.VisitChildren(s)
	}
}

type ToTruthContext struct {
	AccessOpContext
}

func NewToTruthContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToTruthContext {
	var p = new(ToTruthContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *ToTruthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToTruthContext) Equal() antlr.TerminalNode {
	return s.GetToken(RhumbParserEqual, 0)
}

func (s *ToTruthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterToTruth(s)
	}
}

func (s *ToTruthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitToTruth(s)
	}
}

func (s *ToTruthContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitToTruth(s)

	default:
		return t.VisitChildren(s)
	}
}

type ToDateContext struct {
	AccessOpContext
}

func NewToDateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToDateContext {
	var p = new(ToDateContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *ToDateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToDateContext) FSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserFSlash, 0)
}

func (s *ToDateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterToDate(s)
	}
}

func (s *ToDateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitToDate(s)
	}
}

func (s *ToDateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitToDate(s)

	default:
		return t.VisitChildren(s)
	}
}

type LengthContext struct {
	AccessOpContext
}

func NewLengthContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LengthContext {
	var p = new(LengthContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *LengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthContext) Hash() antlr.TerminalNode {
	return s.GetToken(RhumbParserHash, 0)
}

func (s *LengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterLength(s)
	}
}

func (s *LengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitLength(s)
	}
}

func (s *LengthContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitLength(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstructorContext struct {
	AccessOpContext
}

func NewConstructorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstructorContext {
	var p = new(ConstructorContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *ConstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructorContext) Caret() antlr.TerminalNode {
	return s.GetToken(RhumbParserCaret, 0)
}

func (s *ConstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterConstructor(s)
	}
}

func (s *ConstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitConstructor(s)
	}
}

func (s *ConstructorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitConstructor(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegateNumberContext struct {
	AccessOpContext
}

func NewNegateNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegateNumberContext {
	var p = new(NegateNumberContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *NegateNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegateNumberContext) Dash() antlr.TerminalNode {
	return s.GetToken(RhumbParserDash, 0)
}

func (s *NegateNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNegateNumber(s)
	}
}

func (s *NegateNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNegateNumber(s)
	}
}

func (s *NegateNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNegateNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type EmptyContext struct {
	AccessOpContext
}

func NewEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyContext {
	var p = new(EmptyContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *EmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyContext) QMark() antlr.TerminalNode {
	return s.GetToken(RhumbParserQMark, 0)
}

func (s *EmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterEmpty(s)
	}
}

func (s *EmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitEmpty(s)
	}
}

func (s *EmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

type AllSubfieldsContext struct {
	AccessOpContext
}

func NewAllSubfieldsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllSubfieldsContext {
	var p = new(AllSubfieldsContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *AllSubfieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllSubfieldsContext) At() antlr.TerminalNode {
	return s.GetToken(RhumbParserAt, 0)
}

func (s *AllSubfieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAllSubfields(s)
	}
}

func (s *AllSubfieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAllSubfields(s)
	}
}

func (s *AllSubfieldsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAllSubfields(s)

	default:
		return t.VisitChildren(s)
	}
}

type ToKeyContext struct {
	AccessOpContext
}

func NewToKeyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToKeyContext {
	var p = new(ToKeyContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *ToKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToKeyContext) Backtick() antlr.TerminalNode {
	return s.GetToken(RhumbParserBacktick, 0)
}

func (s *ToKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterToKey(s)
	}
}

func (s *ToKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitToKey(s)
	}
}

func (s *ToKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitToKey(s)

	default:
		return t.VisitChildren(s)
	}
}

type FreezeContext struct {
	AccessOpContext
}

func NewFreezeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FreezeContext {
	var p = new(FreezeContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *FreezeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FreezeContext) Dot() antlr.TerminalNode {
	return s.GetToken(RhumbParserDot, 0)
}

func (s *FreezeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterFreeze(s)
	}
}

func (s *FreezeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitFreeze(s)
	}
}

func (s *FreezeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitFreeze(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegateTruthContext struct {
	AccessOpContext
}

func NewNegateTruthContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegateTruthContext {
	var p = new(NegateTruthContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *NegateTruthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegateTruthContext) Tilde() antlr.TerminalNode {
	return s.GetToken(RhumbParserTilde, 0)
}

func (s *NegateTruthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNegateTruth(s)
	}
}

func (s *NegateTruthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNegateTruth(s)
	}
}

func (s *NegateTruthContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNegateTruth(s)

	default:
		return t.VisitChildren(s)
	}
}

type ElementsContext struct {
	AccessOpContext
}

func NewElementsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElementsContext {
	var p = new(ElementsContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *ElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementsContext) Zero() antlr.TerminalNode {
	return s.GetToken(RhumbParserZero, 0)
}

func (s *ElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterElements(s)
	}
}

func (s *ElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitElements(s)
	}
}

func (s *ElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitElements(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnshiftContext struct {
	AccessOpContext
}

func NewUnshiftContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnshiftContext {
	var p = new(UnshiftContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *UnshiftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnshiftContext) OpenAnglet() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenAnglet, 0)
}

func (s *UnshiftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterUnshift(s)
	}
}

func (s *UnshiftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitUnshift(s)
	}
}

func (s *UnshiftContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitUnshift(s)

	default:
		return t.VisitChildren(s)
	}
}

type CopyContext struct {
	AccessOpContext
}

func NewCopyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CopyContext {
	var p = new(CopyContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *CopyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyContext) Colon() antlr.TerminalNode {
	return s.GetToken(RhumbParserColon, 0)
}

func (s *CopyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterCopy(s)
	}
}

func (s *CopyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitCopy(s)
	}
}

func (s *CopyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitCopy(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariadicContext struct {
	AccessOpContext
}

func NewVariadicContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariadicContext {
	var p = new(VariadicContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *VariadicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariadicContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(RhumbParserAmpersand, 0)
}

func (s *VariadicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterVariadic(s)
	}
}

func (s *VariadicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitVariadic(s)
	}
}

func (s *VariadicContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitVariadic(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParametersContext struct {
	AccessOpContext
}

func NewParametersContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParametersContext {
	var p = new(ParametersContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) Dollar() antlr.TerminalNode {
	return s.GetToken(RhumbParserDollar, 0)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (s *ParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

type ToNumberContext struct {
	AccessOpContext
}

func NewToNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToNumberContext {
	var p = new(ToNumberContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *ToNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToNumberContext) Plus() antlr.TerminalNode {
	return s.GetToken(RhumbParserPlus, 0)
}

func (s *ToNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterToNumber(s)
	}
}

func (s *ToNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitToNumber(s)
	}
}

func (s *ToNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitToNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppendContext struct {
	AccessOpContext
}

func NewAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppendContext {
	var p = new(AppendContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *AppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendContext) CloseAnglet() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseAnglet, 0)
}

func (s *AppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAppend(s)
	}
}

func (s *AppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAppend(s)
	}
}

func (s *AppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

type BaseContext struct {
	AccessOpContext
}

func NewBaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BaseContext {
	var p = new(BaseContext)

	InitEmptyAccessOpContext(&p.AccessOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessOpContext))

	return p
}

func (s *BaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseContext) Bang() antlr.TerminalNode {
	return s.GetToken(RhumbParserBang, 0)
}

func (s *BaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterBase(s)
	}
}

func (s *BaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitBase(s)
	}
}

func (s *BaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitBase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) AccessOp() (localctx IAccessOpContext) {
	localctx = NewAccessOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RhumbParserRULE_accessOp)
	p.SetState(535)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RhumbParserCloseAnglet:
		localctx = NewAppendContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(516)
			p.Match(RhumbParserCloseAnglet)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserOpenAnglet:
		localctx = NewUnshiftContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(517)
			p.Match(RhumbParserOpenAnglet)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserHash:
		localctx = NewLengthContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(518)
			p.Match(RhumbParserHash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserQMark:
		localctx = NewEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(519)
			p.Match(RhumbParserQMark)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserAt:
		localctx = NewAllSubfieldsContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(520)
			p.Match(RhumbParserAt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserStar:
		localctx = NewAllFieldsContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(521)
			p.Match(RhumbParserStar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserZero:
		localctx = NewElementsContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(522)
			p.Match(RhumbParserZero)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserDot:
		localctx = NewFreezeContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(523)
			p.Match(RhumbParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserColon:
		localctx = NewCopyContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(524)
			p.Match(RhumbParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserFSlash:
		localctx = NewToDateContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(525)
			p.Match(RhumbParserFSlash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserDollar:
		localctx = NewParametersContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(526)
			p.Match(RhumbParserDollar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserCaret:
		localctx = NewConstructorContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(527)
			p.Match(RhumbParserCaret)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserBang:
		localctx = NewBaseContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(528)
			p.Match(RhumbParserBang)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserPlus:
		localctx = NewToNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(529)
			p.Match(RhumbParserPlus)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserDash:
		localctx = NewNegateNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(530)
			p.Match(RhumbParserDash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserEqual:
		localctx = NewToTruthContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(531)
			p.Match(RhumbParserEqual)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserTilde:
		localctx = NewNegateTruthContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(532)
			p.Match(RhumbParserTilde)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserAmpersand:
		localctx = NewVariadicContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(533)
			p.Match(RhumbParserAmpersand)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserBacktick:
		localctx = NewToKeyContext(p, localctx)
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(534)
			p.Match(RhumbParserBacktick)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IApplicativeOpContext is an interface to support dynamic dispatch.
type IApplicativeOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsApplicativeOpContext differentiates from other interfaces.
	IsApplicativeOpContext()
}

type ApplicativeOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApplicativeOpContext() *ApplicativeOpContext {
	var p = new(ApplicativeOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_applicativeOp
	return p
}

func InitEmptyApplicativeOpContext(p *ApplicativeOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_applicativeOp
}

func (*ApplicativeOpContext) IsApplicativeOpContext() {}

func NewApplicativeOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApplicativeOpContext {
	var p = new(ApplicativeOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_applicativeOp

	return p
}

func (s *ApplicativeOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ApplicativeOpContext) CopyAll(ctx *ApplicativeOpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ApplicativeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplicativeOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MethodContext struct {
	ApplicativeOpContext
}

func NewMethodContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethodContext {
	var p = new(MethodContext)

	InitEmptyApplicativeOpContext(&p.ApplicativeOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ApplicativeOpContext))

	return p
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) BangGreater() antlr.TerminalNode {
	return s.GetToken(RhumbParserBangGreater, 0)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (s *MethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionContext struct {
	ApplicativeOpContext
}

func NewFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionContext {
	var p = new(FunctionContext)

	InitEmptyApplicativeOpContext(&p.ApplicativeOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ApplicativeOpContext))

	return p
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) DashGreater() antlr.TerminalNode {
	return s.GetToken(RhumbParserDashGreater, 0)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) ApplicativeOp() (localctx IApplicativeOpContext) {
	localctx = NewApplicativeOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RhumbParserRULE_applicativeOp)
	p.SetState(539)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RhumbParserDashGreater:
		localctx = NewFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(537)
			p.Match(RhumbParserDashGreater)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserBangGreater:
		localctx = NewMethodContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(538)
			p.Match(RhumbParserBangGreater)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparativeOpContext is an interface to support dynamic dispatch.
type IComparativeOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsComparativeOpContext differentiates from other interfaces.
	IsComparativeOpContext()
}

type ComparativeOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparativeOpContext() *ComparativeOpContext {
	var p = new(ComparativeOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_comparativeOp
	return p
}

func InitEmptyComparativeOpContext(p *ComparativeOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_comparativeOp
}

func (*ComparativeOpContext) IsComparativeOpContext() {}

func NewComparativeOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparativeOpContext {
	var p = new(ComparativeOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_comparativeOp

	return p
}

func (s *ComparativeOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparativeOpContext) CopyAll(ctx *ComparativeOpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ComparativeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparativeOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LessThanContext struct {
	ComparativeOpContext
}

func NewLessThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanContext {
	var p = new(LessThanContext)

	InitEmptyComparativeOpContext(&p.ComparativeOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ComparativeOpContext))

	return p
}

func (s *LessThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanContext) LesserLesser() antlr.TerminalNode {
	return s.GetToken(RhumbParserLesserLesser, 0)
}

func (s *LessThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterLessThan(s)
	}
}

func (s *LessThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitLessThan(s)
	}
}

func (s *LessThanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitLessThan(s)

	default:
		return t.VisitChildren(s)
	}
}

type GreaterThanOrEqualToContext struct {
	ComparativeOpContext
}

func NewGreaterThanOrEqualToContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanOrEqualToContext {
	var p = new(GreaterThanOrEqualToContext)

	InitEmptyComparativeOpContext(&p.ComparativeOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ComparativeOpContext))

	return p
}

func (s *GreaterThanOrEqualToContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanOrEqualToContext) GreaterEqual() antlr.TerminalNode {
	return s.GetToken(RhumbParserGreaterEqual, 0)
}

func (s *GreaterThanOrEqualToContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterGreaterThanOrEqualTo(s)
	}
}

func (s *GreaterThanOrEqualToContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitGreaterThanOrEqualTo(s)
	}
}

func (s *GreaterThanOrEqualToContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitGreaterThanOrEqualTo(s)

	default:
		return t.VisitChildren(s)
	}
}

type LessThanOrEqualToContext struct {
	ComparativeOpContext
}

func NewLessThanOrEqualToContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanOrEqualToContext {
	var p = new(LessThanOrEqualToContext)

	InitEmptyComparativeOpContext(&p.ComparativeOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ComparativeOpContext))

	return p
}

func (s *LessThanOrEqualToContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanOrEqualToContext) LesserEqual() antlr.TerminalNode {
	return s.GetToken(RhumbParserLesserEqual, 0)
}

func (s *LessThanOrEqualToContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterLessThanOrEqualTo(s)
	}
}

func (s *LessThanOrEqualToContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitLessThanOrEqualTo(s)
	}
}

func (s *LessThanOrEqualToContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitLessThanOrEqualTo(s)

	default:
		return t.VisitChildren(s)
	}
}

type GreaterThanContext struct {
	ComparativeOpContext
}

func NewGreaterThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanContext {
	var p = new(GreaterThanContext)

	InitEmptyComparativeOpContext(&p.ComparativeOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ComparativeOpContext))

	return p
}

func (s *GreaterThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanContext) GreaterGreater() antlr.TerminalNode {
	return s.GetToken(RhumbParserGreaterGreater, 0)
}

func (s *GreaterThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterGreaterThan(s)
	}
}

func (s *GreaterThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitGreaterThan(s)
	}
}

func (s *GreaterThanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitGreaterThan(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) ComparativeOp() (localctx IComparativeOpContext) {
	localctx = NewComparativeOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RhumbParserRULE_comparativeOp)
	p.SetState(545)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RhumbParserGreaterGreater:
		localctx = NewGreaterThanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(541)
			p.Match(RhumbParserGreaterGreater)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserGreaterEqual:
		localctx = NewGreaterThanOrEqualToContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(542)
			p.Match(RhumbParserGreaterEqual)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserLesserLesser:
		localctx = NewLessThanContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(543)
			p.Match(RhumbParserLesserLesser)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserLesserEqual:
		localctx = NewLessThanOrEqualToContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(544)
			p.Match(RhumbParserLesserEqual)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentityOpContext is an interface to support dynamic dispatch.
type IIdentityOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIdentityOpContext differentiates from other interfaces.
	IsIdentityOpContext()
}

type IdentityOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentityOpContext() *IdentityOpContext {
	var p = new(IdentityOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_identityOp
	return p
}

func InitEmptyIdentityOpContext(p *IdentityOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_identityOp
}

func (*IdentityOpContext) IsIdentityOpContext() {}

func NewIdentityOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentityOpContext {
	var p = new(IdentityOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_identityOp

	return p
}

func (s *IdentityOpContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentityOpContext) CopyAll(ctx *IdentityOpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IdentityOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentityOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IsUnderContext struct {
	IdentityOpContext
}

func NewIsUnderContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsUnderContext {
	var p = new(IsUnderContext)

	InitEmptyIdentityOpContext(&p.IdentityOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdentityOpContext))

	return p
}

func (s *IsUnderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsUnderContext) EqualAt() antlr.TerminalNode {
	return s.GetToken(RhumbParserEqualAt, 0)
}

func (s *IsUnderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterIsUnder(s)
	}
}

func (s *IsUnderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitIsUnder(s)
	}
}

func (s *IsUnderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitIsUnder(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsEqualContext struct {
	IdentityOpContext
}

func NewIsEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsEqualContext {
	var p = new(IsEqualContext)

	InitEmptyIdentityOpContext(&p.IdentityOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdentityOpContext))

	return p
}

func (s *IsEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsEqualContext) EqualEqual() antlr.TerminalNode {
	return s.GetToken(RhumbParserEqualEqual, 0)
}

func (s *IsEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterIsEqual(s)
	}
}

func (s *IsEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitIsEqual(s)
	}
}

func (s *IsEqualContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitIsEqual(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotEqualContext struct {
	IdentityOpContext
}

func NewNotEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotEqualContext {
	var p = new(NotEqualContext)

	InitEmptyIdentityOpContext(&p.IdentityOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdentityOpContext))

	return p
}

func (s *NotEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqualContext) TildeTilde() antlr.TerminalNode {
	return s.GetToken(RhumbParserTildeTilde, 0)
}

func (s *NotEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNotEqual(s)
	}
}

func (s *NotEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNotEqual(s)
	}
}

func (s *NotEqualContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNotEqual(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotInnerContext struct {
	IdentityOpContext
}

func NewNotInnerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotInnerContext {
	var p = new(NotInnerContext)

	InitEmptyIdentityOpContext(&p.IdentityOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdentityOpContext))

	return p
}

func (s *NotInnerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotInnerContext) TildeBSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserTildeBSlash, 0)
}

func (s *NotInnerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNotInner(s)
	}
}

func (s *NotInnerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNotInner(s)
	}
}

func (s *NotInnerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNotInner(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotUnderContext struct {
	IdentityOpContext
}

func NewNotUnderContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotUnderContext {
	var p = new(NotUnderContext)

	InitEmptyIdentityOpContext(&p.IdentityOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdentityOpContext))

	return p
}

func (s *NotUnderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotUnderContext) TildeAt() antlr.TerminalNode {
	return s.GetToken(RhumbParserTildeAt, 0)
}

func (s *NotUnderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNotUnder(s)
	}
}

func (s *NotUnderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNotUnder(s)
	}
}

func (s *NotUnderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNotUnder(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsInnerContext struct {
	IdentityOpContext
}

func NewIsInnerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsInnerContext {
	var p = new(IsInnerContext)

	InitEmptyIdentityOpContext(&p.IdentityOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdentityOpContext))

	return p
}

func (s *IsInnerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsInnerContext) EqualBSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserEqualBSlash, 0)
}

func (s *IsInnerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterIsInner(s)
	}
}

func (s *IsInnerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitIsInner(s)
	}
}

func (s *IsInnerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitIsInner(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) IdentityOp() (localctx IIdentityOpContext) {
	localctx = NewIdentityOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RhumbParserRULE_identityOp)
	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RhumbParserEqualEqual:
		localctx = NewIsEqualContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(547)
			p.Match(RhumbParserEqualEqual)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserEqualBSlash:
		localctx = NewIsInnerContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(548)
			p.Match(RhumbParserEqualBSlash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserEqualAt:
		localctx = NewIsUnderContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(549)
			p.Match(RhumbParserEqualAt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserTildeTilde:
		localctx = NewNotEqualContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(550)
			p.Match(RhumbParserTildeTilde)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserTildeBSlash:
		localctx = NewNotInnerContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(551)
			p.Match(RhumbParserTildeBSlash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserTildeAt:
		localctx = NewNotUnderContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(552)
			p.Match(RhumbParserTildeAt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConjunctiveOpContext is an interface to support dynamic dispatch.
type IConjunctiveOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FSlashBSlash() antlr.TerminalNode

	// IsConjunctiveOpContext differentiates from other interfaces.
	IsConjunctiveOpContext()
}

type ConjunctiveOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConjunctiveOpContext() *ConjunctiveOpContext {
	var p = new(ConjunctiveOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_conjunctiveOp
	return p
}

func InitEmptyConjunctiveOpContext(p *ConjunctiveOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_conjunctiveOp
}

func (*ConjunctiveOpContext) IsConjunctiveOpContext() {}

func NewConjunctiveOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConjunctiveOpContext {
	var p = new(ConjunctiveOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_conjunctiveOp

	return p
}

func (s *ConjunctiveOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ConjunctiveOpContext) FSlashBSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserFSlashBSlash, 0)
}

func (s *ConjunctiveOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConjunctiveOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConjunctiveOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterConjunctiveOp(s)
	}
}

func (s *ConjunctiveOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitConjunctiveOp(s)
	}
}

func (s *ConjunctiveOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitConjunctiveOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) ConjunctiveOp() (localctx IConjunctiveOpContext) {
	localctx = NewConjunctiveOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RhumbParserRULE_conjunctiveOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(555)
		p.Match(RhumbParserFSlashBSlash)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDisjunctiveOpContext is an interface to support dynamic dispatch.
type IDisjunctiveOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BSlashFSlash() antlr.TerminalNode

	// IsDisjunctiveOpContext differentiates from other interfaces.
	IsDisjunctiveOpContext()
}

type DisjunctiveOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisjunctiveOpContext() *DisjunctiveOpContext {
	var p = new(DisjunctiveOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_disjunctiveOp
	return p
}

func InitEmptyDisjunctiveOpContext(p *DisjunctiveOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_disjunctiveOp
}

func (*DisjunctiveOpContext) IsDisjunctiveOpContext() {}

func NewDisjunctiveOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisjunctiveOpContext {
	var p = new(DisjunctiveOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_disjunctiveOp

	return p
}

func (s *DisjunctiveOpContext) GetParser() antlr.Parser { return s.parser }

func (s *DisjunctiveOpContext) BSlashFSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserBSlashFSlash, 0)
}

func (s *DisjunctiveOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctiveOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisjunctiveOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterDisjunctiveOp(s)
	}
}

func (s *DisjunctiveOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitDisjunctiveOp(s)
	}
}

func (s *DisjunctiveOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitDisjunctiveOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) DisjunctiveOp() (localctx IDisjunctiveOpContext) {
	localctx = NewDisjunctiveOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RhumbParserRULE_disjunctiveOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(557)
		p.Match(RhumbParserBSlashFSlash)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalOpContext is an interface to support dynamic dispatch.
type IConditionalOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConditionalOpContext differentiates from other interfaces.
	IsConditionalOpContext()
}

type ConditionalOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalOpContext() *ConditionalOpContext {
	var p = new(ConditionalOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_conditionalOp
	return p
}

func InitEmptyConditionalOpContext(p *ConditionalOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_conditionalOp
}

func (*ConditionalOpContext) IsConditionalOpContext() {}

func NewConditionalOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalOpContext {
	var p = new(ConditionalOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_conditionalOp

	return p
}

func (s *ConditionalOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalOpContext) CopyAll(ctx *ConditionalOpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ConditionalOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForeachContext struct {
	ConditionalOpContext
}

func NewForeachContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForeachContext {
	var p = new(ForeachContext)

	InitEmptyConditionalOpContext(&p.ConditionalOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionalOpContext))

	return p
}

func (s *ForeachContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForeachContext) LesserGreater() antlr.TerminalNode {
	return s.GetToken(RhumbParserLesserGreater, 0)
}

func (s *ForeachContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterForeach(s)
	}
}

func (s *ForeachContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitForeach(s)
	}
}

func (s *ForeachContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitForeach(s)

	default:
		return t.VisitChildren(s)
	}
}

type DefaultContext struct {
	ConditionalOpContext
}

func NewDefaultContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultContext {
	var p = new(DefaultContext)

	InitEmptyConditionalOpContext(&p.ConditionalOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionalOpContext))

	return p
}

func (s *DefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultContext) QMarkQMark() antlr.TerminalNode {
	return s.GetToken(RhumbParserQMarkQMark, 0)
}

func (s *DefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterDefault(s)
	}
}

func (s *DefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitDefault(s)
	}
}

func (s *DefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

type ElseContext struct {
	ConditionalOpContext
}

func NewElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseContext {
	var p = new(ElseContext)

	InitEmptyConditionalOpContext(&p.ConditionalOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionalOpContext))

	return p
}

func (s *ElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseContext) TildeGreater() antlr.TerminalNode {
	return s.GetToken(RhumbParserTildeGreater, 0)
}

func (s *ElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterElse(s)
	}
}

func (s *ElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitElse(s)
	}
}

func (s *ElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitElse(s)

	default:
		return t.VisitChildren(s)
	}
}

type PipeContext struct {
	ConditionalOpContext
}

func NewPipeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PipeContext {
	var p = new(PipeContext)

	InitEmptyConditionalOpContext(&p.ConditionalOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionalOpContext))

	return p
}

func (s *PipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeContext) PipePipe() antlr.TerminalNode {
	return s.GetToken(RhumbParserPipePipe, 0)
}

func (s *PipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterPipe(s)
	}
}

func (s *PipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitPipe(s)
	}
}

func (s *PipeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitPipe(s)

	default:
		return t.VisitChildren(s)
	}
}

type ThenContext struct {
	ConditionalOpContext
}

func NewThenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThenContext {
	var p = new(ThenContext)

	InitEmptyConditionalOpContext(&p.ConditionalOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionalOpContext))

	return p
}

func (s *ThenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThenContext) EqualGreater() antlr.TerminalNode {
	return s.GetToken(RhumbParserEqualGreater, 0)
}

func (s *ThenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterThen(s)
	}
}

func (s *ThenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitThen(s)
	}
}

func (s *ThenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitThen(s)

	default:
		return t.VisitChildren(s)
	}
}

type WhileContext struct {
	ConditionalOpContext
}

func NewWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileContext {
	var p = new(WhileContext)

	InitEmptyConditionalOpContext(&p.ConditionalOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionalOpContext))

	return p
}

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) PipeGreater() antlr.TerminalNode {
	return s.GetToken(RhumbParserPipeGreater, 0)
}

func (s *WhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterWhile(s)
	}
}

func (s *WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitWhile(s)
	}
}

func (s *WhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

type DestructureContext struct {
	ConditionalOpContext
}

func NewDestructureContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DestructureContext {
	var p = new(DestructureContext)

	InitEmptyConditionalOpContext(&p.ConditionalOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionalOpContext))

	return p
}

func (s *DestructureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestructureContext) CaretEqual() antlr.TerminalNode {
	return s.GetToken(RhumbParserCaretEqual, 0)
}

func (s *DestructureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterDestructure(s)
	}
}

func (s *DestructureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitDestructure(s)
	}
}

func (s *DestructureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitDestructure(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) ConditionalOp() (localctx IConditionalOpContext) {
	localctx = NewConditionalOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RhumbParserRULE_conditionalOp)
	p.SetState(566)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RhumbParserPipePipe:
		localctx = NewPipeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(559)
			p.Match(RhumbParserPipePipe)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserQMarkQMark:
		localctx = NewDefaultContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(560)
			p.Match(RhumbParserQMarkQMark)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserLesserGreater:
		localctx = NewForeachContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(561)
			p.Match(RhumbParserLesserGreater)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserPipeGreater:
		localctx = NewWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(562)
			p.Match(RhumbParserPipeGreater)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserEqualGreater:
		localctx = NewThenContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(563)
			p.Match(RhumbParserEqualGreater)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserTildeGreater:
		localctx = NewElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(564)
			p.Match(RhumbParserTildeGreater)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserCaretEqual:
		localctx = NewDestructureContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(565)
			p.Match(RhumbParserCaretEqual)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdditiveOpContext is an interface to support dynamic dispatch.
type IAdditiveOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAdditiveOpContext differentiates from other interfaces.
	IsAdditiveOpContext()
}

type AdditiveOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveOpContext() *AdditiveOpContext {
	var p = new(AdditiveOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_additiveOp
	return p
}

func InitEmptyAdditiveOpContext(p *AdditiveOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_additiveOp
}

func (*AdditiveOpContext) IsAdditiveOpContext() {}

func NewAdditiveOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveOpContext {
	var p = new(AdditiveOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_additiveOp

	return p
}

func (s *AdditiveOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveOpContext) CopyAll(ctx *AdditiveOpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AdditiveOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConcatenateContext struct {
	AdditiveOpContext
}

func NewConcatenateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConcatenateContext {
	var p = new(ConcatenateContext)

	InitEmptyAdditiveOpContext(&p.AdditiveOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AdditiveOpContext))

	return p
}

func (s *ConcatenateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatenateContext) AmpAmp() antlr.TerminalNode {
	return s.GetToken(RhumbParserAmpAmp, 0)
}

func (s *ConcatenateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterConcatenate(s)
	}
}

func (s *ConcatenateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitConcatenate(s)
	}
}

func (s *ConcatenateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitConcatenate(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubtractionContext struct {
	AdditiveOpContext
}

func NewSubtractionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubtractionContext {
	var p = new(SubtractionContext)

	InitEmptyAdditiveOpContext(&p.AdditiveOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AdditiveOpContext))

	return p
}

func (s *SubtractionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtractionContext) DashDash() antlr.TerminalNode {
	return s.GetToken(RhumbParserDashDash, 0)
}

func (s *SubtractionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterSubtraction(s)
	}
}

func (s *SubtractionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitSubtraction(s)
	}
}

func (s *SubtractionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitSubtraction(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeviationContext struct {
	AdditiveOpContext
}

func NewDeviationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeviationContext {
	var p = new(DeviationContext)

	InitEmptyAdditiveOpContext(&p.AdditiveOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AdditiveOpContext))

	return p
}

func (s *DeviationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeviationContext) PlusDash() antlr.TerminalNode {
	return s.GetToken(RhumbParserPlusDash, 0)
}

func (s *DeviationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterDeviation(s)
	}
}

func (s *DeviationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitDeviation(s)
	}
}

func (s *DeviationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitDeviation(s)

	default:
		return t.VisitChildren(s)
	}
}

type AdditionContext struct {
	AdditiveOpContext
}

func NewAdditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditionContext {
	var p = new(AdditionContext)

	InitEmptyAdditiveOpContext(&p.AdditiveOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AdditiveOpContext))

	return p
}

func (s *AdditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionContext) PlusPlus() antlr.TerminalNode {
	return s.GetToken(RhumbParserPlusPlus, 0)
}

func (s *AdditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAddition(s)
	}
}

func (s *AdditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAddition(s)
	}
}

func (s *AdditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAddition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) AdditiveOp() (localctx IAdditiveOpContext) {
	localctx = NewAdditiveOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RhumbParserRULE_additiveOp)
	p.SetState(572)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RhumbParserPlusPlus:
		localctx = NewAdditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(568)
			p.Match(RhumbParserPlusPlus)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserPlusDash:
		localctx = NewDeviationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(569)
			p.Match(RhumbParserPlusDash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserDashDash:
		localctx = NewSubtractionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(570)
			p.Match(RhumbParserDashDash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserAmpAmp:
		localctx = NewConcatenateContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(571)
			p.Match(RhumbParserAmpAmp)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplicativeOpContext is an interface to support dynamic dispatch.
type IMultiplicativeOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMultiplicativeOpContext differentiates from other interfaces.
	IsMultiplicativeOpContext()
}

type MultiplicativeOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeOpContext() *MultiplicativeOpContext {
	var p = new(MultiplicativeOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_multiplicativeOp
	return p
}

func InitEmptyMultiplicativeOpContext(p *MultiplicativeOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_multiplicativeOp
}

func (*MultiplicativeOpContext) IsMultiplicativeOpContext() {}

func NewMultiplicativeOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeOpContext {
	var p = new(MultiplicativeOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_multiplicativeOp

	return p
}

func (s *MultiplicativeOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeOpContext) CopyAll(ctx *MultiplicativeOpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MultiplicativeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RationalDivisionContext struct {
	MultiplicativeOpContext
}

func NewRationalDivisionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RationalDivisionContext {
	var p = new(RationalDivisionContext)

	InitEmptyMultiplicativeOpContext(&p.MultiplicativeOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultiplicativeOpContext))

	return p
}

func (s *RationalDivisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RationalDivisionContext) FSlashFSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserFSlashFSlash, 0)
}

func (s *RationalDivisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterRationalDivision(s)
	}
}

func (s *RationalDivisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitRationalDivision(s)
	}
}

func (s *RationalDivisionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitRationalDivision(s)

	default:
		return t.VisitChildren(s)
	}
}

type BindContext struct {
	MultiplicativeOpContext
}

func NewBindContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BindContext {
	var p = new(BindContext)

	InitEmptyMultiplicativeOpContext(&p.MultiplicativeOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultiplicativeOpContext))

	return p
}

func (s *BindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindContext) BangBang() antlr.TerminalNode {
	return s.GetToken(RhumbParserBangBang, 0)
}

func (s *BindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterBind(s)
	}
}

func (s *BindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitBind(s)
	}
}

func (s *BindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitBind(s)

	default:
		return t.VisitChildren(s)
	}
}

type RemainderDivisionContext struct {
	MultiplicativeOpContext
}

func NewRemainderDivisionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RemainderDivisionContext {
	var p = new(RemainderDivisionContext)

	InitEmptyMultiplicativeOpContext(&p.MultiplicativeOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultiplicativeOpContext))

	return p
}

func (s *RemainderDivisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemainderDivisionContext) DashFSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserDashFSlash, 0)
}

func (s *RemainderDivisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterRemainderDivision(s)
	}
}

func (s *RemainderDivisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitRemainderDivision(s)
	}
}

func (s *RemainderDivisionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitRemainderDivision(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiplicationContext struct {
	MultiplicativeOpContext
}

func NewMultiplicationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicationContext {
	var p = new(MultiplicationContext)

	InitEmptyMultiplicativeOpContext(&p.MultiplicativeOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultiplicativeOpContext))

	return p
}

func (s *MultiplicationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationContext) StarStar() antlr.TerminalNode {
	return s.GetToken(RhumbParserStarStar, 0)
}

func (s *MultiplicationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterMultiplication(s)
	}
}

func (s *MultiplicationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitMultiplication(s)
	}
}

func (s *MultiplicationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitMultiplication(s)

	default:
		return t.VisitChildren(s)
	}
}

type WholeDivisionContext struct {
	MultiplicativeOpContext
}

func NewWholeDivisionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WholeDivisionContext {
	var p = new(WholeDivisionContext)

	InitEmptyMultiplicativeOpContext(&p.MultiplicativeOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultiplicativeOpContext))

	return p
}

func (s *WholeDivisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WholeDivisionContext) PlusFSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserPlusFSlash, 0)
}

func (s *WholeDivisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterWholeDivision(s)
	}
}

func (s *WholeDivisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitWholeDivision(s)
	}
}

func (s *WholeDivisionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitWholeDivision(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) MultiplicativeOp() (localctx IMultiplicativeOpContext) {
	localctx = NewMultiplicativeOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RhumbParserRULE_multiplicativeOp)
	p.SetState(579)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RhumbParserStarStar:
		localctx = NewMultiplicationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(574)
			p.Match(RhumbParserStarStar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserFSlashFSlash:
		localctx = NewRationalDivisionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(575)
			p.Match(RhumbParserFSlashFSlash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserPlusFSlash:
		localctx = NewWholeDivisionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(576)
			p.Match(RhumbParserPlusFSlash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserDashFSlash:
		localctx = NewRemainderDivisionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(577)
			p.Match(RhumbParserDashFSlash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserBangBang:
		localctx = NewBindContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(578)
			p.Match(RhumbParserBangBang)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExponentiationOpContext is an interface to support dynamic dispatch.
type IExponentiationOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExponentiationOpContext differentiates from other interfaces.
	IsExponentiationOpContext()
}

type ExponentiationOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExponentiationOpContext() *ExponentiationOpContext {
	var p = new(ExponentiationOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_exponentiationOp
	return p
}

func InitEmptyExponentiationOpContext(p *ExponentiationOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_exponentiationOp
}

func (*ExponentiationOpContext) IsExponentiationOpContext() {}

func NewExponentiationOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExponentiationOpContext {
	var p = new(ExponentiationOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_exponentiationOp

	return p
}

func (s *ExponentiationOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExponentiationOpContext) CopyAll(ctx *ExponentiationOpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExponentiationOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentiationOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RootExtractionContext struct {
	ExponentiationOpContext
}

func NewRootExtractionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RootExtractionContext {
	var p = new(RootExtractionContext)

	InitEmptyExponentiationOpContext(&p.ExponentiationOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExponentiationOpContext))

	return p
}

func (s *RootExtractionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootExtractionContext) CaretFSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserCaretFSlash, 0)
}

func (s *RootExtractionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterRootExtraction(s)
	}
}

func (s *RootExtractionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitRootExtraction(s)
	}
}

func (s *RootExtractionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitRootExtraction(s)

	default:
		return t.VisitChildren(s)
	}
}

type ScientificContext struct {
	ExponentiationOpContext
}

func NewScientificContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScientificContext {
	var p = new(ScientificContext)

	InitEmptyExponentiationOpContext(&p.ExponentiationOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExponentiationOpContext))

	return p
}

func (s *ScientificContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScientificContext) StarCaret() antlr.TerminalNode {
	return s.GetToken(RhumbParserStarCaret, 0)
}

func (s *ScientificContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterScientific(s)
	}
}

func (s *ScientificContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitScientific(s)
	}
}

func (s *ScientificContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitScientific(s)

	default:
		return t.VisitChildren(s)
	}
}

type RangeContext struct {
	ExponentiationOpContext
}

func NewRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangeContext {
	var p = new(RangeContext)

	InitEmptyExponentiationOpContext(&p.ExponentiationOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExponentiationOpContext))

	return p
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) Pipe() antlr.TerminalNode {
	return s.GetToken(RhumbParserPipe, 0)
}

func (s *RangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterRange(s)
	}
}

func (s *RangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitRange(s)
	}
}

func (s *RangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitRange(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExponentContext struct {
	ExponentiationOpContext
}

func NewExponentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExponentContext {
	var p = new(ExponentContext)

	InitEmptyExponentiationOpContext(&p.ExponentiationOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExponentiationOpContext))

	return p
}

func (s *ExponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentContext) CaretCaret() antlr.TerminalNode {
	return s.GetToken(RhumbParserCaretCaret, 0)
}

func (s *ExponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterExponent(s)
	}
}

func (s *ExponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitExponent(s)
	}
}

func (s *ExponentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitExponent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) ExponentiationOp() (localctx IExponentiationOpContext) {
	localctx = NewExponentiationOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RhumbParserRULE_exponentiationOp)
	p.SetState(585)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RhumbParserCaretCaret:
		localctx = NewExponentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(581)
			p.Match(RhumbParserCaretCaret)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserCaretFSlash:
		localctx = NewRootExtractionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(582)
			p.Match(RhumbParserCaretFSlash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserPipe:
		localctx = NewRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(583)
			p.Match(RhumbParserPipe)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserStarCaret:
		localctx = NewScientificContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(584)
			p.Match(RhumbParserStarCaret)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentOpContext is an interface to support dynamic dispatch.
type IAssignmentOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignmentOpContext differentiates from other interfaces.
	IsAssignmentOpContext()
}

type AssignmentOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOpContext() *AssignmentOpContext {
	var p = new(AssignmentOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_assignmentOp
	return p
}

func InitEmptyAssignmentOpContext(p *AssignmentOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_assignmentOp
}

func (*AssignmentOpContext) IsAssignmentOpContext() {}

func NewAssignmentOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOpContext {
	var p = new(AssignmentOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_assignmentOp

	return p
}

func (s *AssignmentOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentOpContext) CopyAll(ctx *AssignmentOpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssignmentOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MutableLabelContext struct {
	AssignmentOpContext
}

func NewMutableLabelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MutableLabelContext {
	var p = new(MutableLabelContext)

	InitEmptyAssignmentOpContext(&p.AssignmentOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentOpContext))

	return p
}

func (s *MutableLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutableLabelContext) ColonEqual() antlr.TerminalNode {
	return s.GetToken(RhumbParserColonEqual, 0)
}

func (s *MutableLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterMutableLabel(s)
	}
}

func (s *MutableLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitMutableLabel(s)
	}
}

func (s *MutableLabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitMutableLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

type ImmutableLabelContext struct {
	AssignmentOpContext
}

func NewImmutableLabelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImmutableLabelContext {
	var p = new(ImmutableLabelContext)

	InitEmptyAssignmentOpContext(&p.AssignmentOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentOpContext))

	return p
}

func (s *ImmutableLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImmutableLabelContext) DotEqual() antlr.TerminalNode {
	return s.GetToken(RhumbParserDotEqual, 0)
}

func (s *ImmutableLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterImmutableLabel(s)
	}
}

func (s *ImmutableLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitImmutableLabel(s)
	}
}

func (s *ImmutableLabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitImmutableLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) AssignmentOp() (localctx IAssignmentOpContext) {
	localctx = NewAssignmentOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RhumbParserRULE_assignmentOp)
	p.SetState(589)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RhumbParserDotEqual:
		localctx = NewImmutableLabelContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(587)
			p.Match(RhumbParserDotEqual)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserColonEqual:
		localctx = NewMutableLabelContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(588)
			p.Match(RhumbParserColonEqual)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrefixOpContext is an interface to support dynamic dispatch.
type IPrefixOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrefixOpContext differentiates from other interfaces.
	IsPrefixOpContext()
}

type PrefixOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixOpContext() *PrefixOpContext {
	var p = new(PrefixOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_prefixOp
	return p
}

func InitEmptyPrefixOpContext(p *PrefixOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_prefixOp
}

func (*PrefixOpContext) IsPrefixOpContext() {}

func NewPrefixOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixOpContext {
	var p = new(PrefixOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_prefixOp

	return p
}

func (s *PrefixOpContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixOpContext) CopyAll(ctx *PrefixOpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrefixOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NegateNumberPrefixContext struct {
	PrefixOpContext
}

func NewNegateNumberPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegateNumberPrefixContext {
	var p = new(NegateNumberPrefixContext)

	InitEmptyPrefixOpContext(&p.PrefixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrefixOpContext))

	return p
}

func (s *NegateNumberPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegateNumberPrefixContext) Dash() antlr.TerminalNode {
	return s.GetToken(RhumbParserDash, 0)
}

func (s *NegateNumberPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNegateNumberPrefix(s)
	}
}

func (s *NegateNumberPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNegateNumberPrefix(s)
	}
}

func (s *NegateNumberPrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNegateNumberPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type SignalInwardPrefixContext struct {
	PrefixOpContext
}

func NewSignalInwardPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SignalInwardPrefixContext {
	var p = new(SignalInwardPrefixContext)

	InitEmptyPrefixOpContext(&p.PrefixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrefixOpContext))

	return p
}

func (s *SignalInwardPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignalInwardPrefixContext) Caret() antlr.TerminalNode {
	return s.GetToken(RhumbParserCaret, 0)
}

func (s *SignalInwardPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterSignalInwardPrefix(s)
	}
}

func (s *SignalInwardPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitSignalInwardPrefix(s)
	}
}

func (s *SignalInwardPrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitSignalInwardPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type CopyPrefixContext struct {
	PrefixOpContext
}

func NewCopyPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CopyPrefixContext {
	var p = new(CopyPrefixContext)

	InitEmptyPrefixOpContext(&p.PrefixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrefixOpContext))

	return p
}

func (s *CopyPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyPrefixContext) Colon() antlr.TerminalNode {
	return s.GetToken(RhumbParserColon, 0)
}

func (s *CopyPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterCopyPrefix(s)
	}
}

func (s *CopyPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitCopyPrefix(s)
	}
}

func (s *CopyPrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitCopyPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type ToNumberPrefixContext struct {
	PrefixOpContext
}

func NewToNumberPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToNumberPrefixContext {
	var p = new(ToNumberPrefixContext)

	InitEmptyPrefixOpContext(&p.PrefixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrefixOpContext))

	return p
}

func (s *ToNumberPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToNumberPrefixContext) Plus() antlr.TerminalNode {
	return s.GetToken(RhumbParserPlus, 0)
}

func (s *ToNumberPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterToNumberPrefix(s)
	}
}

func (s *ToNumberPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitToNumberPrefix(s)
	}
}

func (s *ToNumberPrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitToNumberPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type ToTruthPrefixContext struct {
	PrefixOpContext
}

func NewToTruthPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToTruthPrefixContext {
	var p = new(ToTruthPrefixContext)

	InitEmptyPrefixOpContext(&p.PrefixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrefixOpContext))

	return p
}

func (s *ToTruthPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToTruthPrefixContext) Equal() antlr.TerminalNode {
	return s.GetToken(RhumbParserEqual, 0)
}

func (s *ToTruthPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterToTruthPrefix(s)
	}
}

func (s *ToTruthPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitToTruthPrefix(s)
	}
}

func (s *ToTruthPrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitToTruthPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariadicPrefixContext struct {
	PrefixOpContext
}

func NewVariadicPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariadicPrefixContext {
	var p = new(VariadicPrefixContext)

	InitEmptyPrefixOpContext(&p.PrefixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrefixOpContext))

	return p
}

func (s *VariadicPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariadicPrefixContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(RhumbParserAmpersand, 0)
}

func (s *VariadicPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterVariadicPrefix(s)
	}
}

func (s *VariadicPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitVariadicPrefix(s)
	}
}

func (s *VariadicPrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitVariadicPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type EmptyPrefixContext struct {
	PrefixOpContext
}

func NewEmptyPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyPrefixContext {
	var p = new(EmptyPrefixContext)

	InitEmptyPrefixOpContext(&p.PrefixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrefixOpContext))

	return p
}

func (s *EmptyPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyPrefixContext) QMark() antlr.TerminalNode {
	return s.GetToken(RhumbParserQMark, 0)
}

func (s *EmptyPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterEmptyPrefix(s)
	}
}

func (s *EmptyPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitEmptyPrefix(s)
	}
}

func (s *EmptyPrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitEmptyPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegateTruthPrefixContext struct {
	PrefixOpContext
}

func NewNegateTruthPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegateTruthPrefixContext {
	var p = new(NegateTruthPrefixContext)

	InitEmptyPrefixOpContext(&p.PrefixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrefixOpContext))

	return p
}

func (s *NegateTruthPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegateTruthPrefixContext) Tilde() antlr.TerminalNode {
	return s.GetToken(RhumbParserTilde, 0)
}

func (s *NegateTruthPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNegateTruthPrefix(s)
	}
}

func (s *NegateTruthPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNegateTruthPrefix(s)
	}
}

func (s *NegateTruthPrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNegateTruthPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArgumentPrefixContext struct {
	PrefixOpContext
}

func NewArgumentPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgumentPrefixContext {
	var p = new(ArgumentPrefixContext)

	InitEmptyPrefixOpContext(&p.PrefixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrefixOpContext))

	return p
}

func (s *ArgumentPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentPrefixContext) Dollar() antlr.TerminalNode {
	return s.GetToken(RhumbParserDollar, 0)
}

func (s *ArgumentPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterArgumentPrefix(s)
	}
}

func (s *ArgumentPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitArgumentPrefix(s)
	}
}

func (s *ArgumentPrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitArgumentPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type FreezePrefixContext struct {
	PrefixOpContext
}

func NewFreezePrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FreezePrefixContext {
	var p = new(FreezePrefixContext)

	InitEmptyPrefixOpContext(&p.PrefixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrefixOpContext))

	return p
}

func (s *FreezePrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FreezePrefixContext) Dot() antlr.TerminalNode {
	return s.GetToken(RhumbParserDot, 0)
}

func (s *FreezePrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterFreezePrefix(s)
	}
}

func (s *FreezePrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitFreezePrefix(s)
	}
}

func (s *FreezePrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitFreezePrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type SignalOutwardPrefixContext struct {
	PrefixOpContext
}

func NewSignalOutwardPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SignalOutwardPrefixContext {
	var p = new(SignalOutwardPrefixContext)

	InitEmptyPrefixOpContext(&p.PrefixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrefixOpContext))

	return p
}

func (s *SignalOutwardPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignalOutwardPrefixContext) Hash() antlr.TerminalNode {
	return s.GetToken(RhumbParserHash, 0)
}

func (s *SignalOutwardPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterSignalOutwardPrefix(s)
	}
}

func (s *SignalOutwardPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitSignalOutwardPrefix(s)
	}
}

func (s *SignalOutwardPrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitSignalOutwardPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) PrefixOp() (localctx IPrefixOpContext) {
	localctx = NewPrefixOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RhumbParserRULE_prefixOp)
	p.SetState(602)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RhumbParserQMark:
		localctx = NewEmptyPrefixContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(591)
			p.Match(RhumbParserQMark)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserDot:
		localctx = NewFreezePrefixContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(592)
			p.Match(RhumbParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserColon:
		localctx = NewCopyPrefixContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(593)
			p.Match(RhumbParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserPlus:
		localctx = NewToNumberPrefixContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(594)
			p.Match(RhumbParserPlus)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserDash:
		localctx = NewNegateNumberPrefixContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(595)
			p.Match(RhumbParserDash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserEqual:
		localctx = NewToTruthPrefixContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(596)
			p.Match(RhumbParserEqual)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserTilde:
		localctx = NewNegateTruthPrefixContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(597)
			p.Match(RhumbParserTilde)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserAmpersand:
		localctx = NewVariadicPrefixContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(598)
			p.Match(RhumbParserAmpersand)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserDollar:
		localctx = NewArgumentPrefixContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(599)
			p.Match(RhumbParserDollar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserHash:
		localctx = NewSignalOutwardPrefixContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(600)
			p.Match(RhumbParserHash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserCaret:
		localctx = NewSignalInwardPrefixContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(601)
			p.Match(RhumbParserCaret)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IChainOpContext is an interface to support dynamic dispatch.
type IChainOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsChainOpContext differentiates from other interfaces.
	IsChainOpContext()
}

type ChainOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChainOpContext() *ChainOpContext {
	var p = new(ChainOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_chainOp
	return p
}

func InitEmptyChainOpContext(p *ChainOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_chainOp
}

func (*ChainOpContext) IsChainOpContext() {}

func NewChainOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChainOpContext {
	var p = new(ChainOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_chainOp

	return p
}

func (s *ChainOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ChainOpContext) CopyAll(ctx *ChainOpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ChainOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeeplyNestedSubfieldContext struct {
	ChainOpContext
}

func NewDeeplyNestedSubfieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeeplyNestedSubfieldContext {
	var p = new(DeeplyNestedSubfieldContext)

	InitEmptyChainOpContext(&p.ChainOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ChainOpContext))

	return p
}

func (s *DeeplyNestedSubfieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeeplyNestedSubfieldContext) AtAt() antlr.TerminalNode {
	return s.GetToken(RhumbParserAtAt, 0)
}

func (s *DeeplyNestedSubfieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterDeeplyNestedSubfield(s)
	}
}

func (s *DeeplyNestedSubfieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitDeeplyNestedSubfield(s)
	}
}

func (s *DeeplyNestedSubfieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitDeeplyNestedSubfield(s)

	default:
		return t.VisitChildren(s)
	}
}

type NestedSubfieldContext struct {
	ChainOpContext
}

func NewNestedSubfieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedSubfieldContext {
	var p = new(NestedSubfieldContext)

	InitEmptyChainOpContext(&p.ChainOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ChainOpContext))

	return p
}

func (s *NestedSubfieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedSubfieldContext) At() antlr.TerminalNode {
	return s.GetToken(RhumbParserAt, 0)
}

func (s *NestedSubfieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNestedSubfield(s)
	}
}

func (s *NestedSubfieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNestedSubfield(s)
	}
}

func (s *NestedSubfieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNestedSubfield(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeeplyNestedFieldContext struct {
	ChainOpContext
}

func NewDeeplyNestedFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeeplyNestedFieldContext {
	var p = new(DeeplyNestedFieldContext)

	InitEmptyChainOpContext(&p.ChainOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ChainOpContext))

	return p
}

func (s *DeeplyNestedFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeeplyNestedFieldContext) BSlashBSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserBSlashBSlash, 0)
}

func (s *DeeplyNestedFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterDeeplyNestedField(s)
	}
}

func (s *DeeplyNestedFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitDeeplyNestedField(s)
	}
}

func (s *DeeplyNestedFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitDeeplyNestedField(s)

	default:
		return t.VisitChildren(s)
	}
}

type NestedFieldContext struct {
	ChainOpContext
}

func NewNestedFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedFieldContext {
	var p = new(NestedFieldContext)

	InitEmptyChainOpContext(&p.ChainOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ChainOpContext))

	return p
}

func (s *NestedFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedFieldContext) BSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserBSlash, 0)
}

func (s *NestedFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNestedField(s)
	}
}

func (s *NestedFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNestedField(s)
	}
}

func (s *NestedFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNestedField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) ChainOp() (localctx IChainOpContext) {
	localctx = NewChainOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, RhumbParserRULE_chainOp)
	p.SetState(608)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RhumbParserBSlash:
		localctx = NewNestedFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(604)
			p.Match(RhumbParserBSlash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserBSlashBSlash:
		localctx = NewDeeplyNestedFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(605)
			p.Match(RhumbParserBSlashBSlash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserAt:
		localctx = NewNestedSubfieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(606)
			p.Match(RhumbParserAt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserAtAt:
		localctx = NewDeeplyNestedSubfieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(607)
			p.Match(RhumbParserAtAt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDatePartContext is an interface to support dynamic dispatch.
type IDatePartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NumberPart() antlr.TerminalNode
	Dollar() antlr.TerminalNode
	Label() antlr.TerminalNode

	// IsDatePartContext differentiates from other interfaces.
	IsDatePartContext()
}

type DatePartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatePartContext() *DatePartContext {
	var p = new(DatePartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_datePart
	return p
}

func InitEmptyDatePartContext(p *DatePartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_datePart
}

func (*DatePartContext) IsDatePartContext() {}

func NewDatePartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatePartContext {
	var p = new(DatePartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_datePart

	return p
}

func (s *DatePartContext) GetParser() antlr.Parser { return s.parser }

func (s *DatePartContext) NumberPart() antlr.TerminalNode {
	return s.GetToken(RhumbParserNumberPart, 0)
}

func (s *DatePartContext) Dollar() antlr.TerminalNode {
	return s.GetToken(RhumbParserDollar, 0)
}

func (s *DatePartContext) Label() antlr.TerminalNode {
	return s.GetToken(RhumbParserLabel, 0)
}

func (s *DatePartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatePartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatePartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterDatePart(s)
	}
}

func (s *DatePartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitDatePart(s)
	}
}

func (s *DatePartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitDatePart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) DatePart() (localctx IDatePartContext) {
	localctx = NewDatePartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, RhumbParserRULE_datePart)
	p.SetState(613)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RhumbParserNumberPart:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(610)
			p.Match(RhumbParserNumberPart)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserDollar:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(611)
			p.Match(RhumbParserDollar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(612)
			p.Match(RhumbParserLabel)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDateContext is an interface to support dynamic dispatch.
type IDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDatePart() []IDatePartContext
	DatePart(i int) IDatePartContext
	AllFSlash() []antlr.TerminalNode
	FSlash(i int) antlr.TerminalNode

	// IsDateContext differentiates from other interfaces.
	IsDateContext()
}

type DateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateContext() *DateContext {
	var p = new(DateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_date
	return p
}

func InitEmptyDateContext(p *DateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_date
}

func (*DateContext) IsDateContext() {}

func NewDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateContext {
	var p = new(DateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_date

	return p
}

func (s *DateContext) GetParser() antlr.Parser { return s.parser }

func (s *DateContext) AllDatePart() []IDatePartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDatePartContext); ok {
			len++
		}
	}

	tst := make([]IDatePartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDatePartContext); ok {
			tst[i] = t.(IDatePartContext)
			i++
		}
	}

	return tst
}

func (s *DateContext) DatePart(i int) IDatePartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatePartContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatePartContext)
}

func (s *DateContext) AllFSlash() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserFSlash)
}

func (s *DateContext) FSlash(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserFSlash, i)
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitDate(s)
	}
}

func (s *DateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitDate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) Date() (localctx IDateContext) {
	localctx = NewDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, RhumbParserRULE_date)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(615)
		p.DatePart()
	}
	{
		p.SetState(616)
		p.Match(RhumbParserFSlash)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(617)
		p.DatePart()
	}
	{
		p.SetState(618)
		p.Match(RhumbParserFSlash)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(619)
		p.DatePart()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpenInterpString() antlr.TerminalNode
	CloseInterpString() antlr.TerminalNode
	Bang() antlr.TerminalNode
	AllInnerText() []antlr.TerminalNode
	InnerText(i int) antlr.TerminalNode
	AllInterpolation() []IInterpolationContext
	Interpolation(i int) IInterpolationContext
	RawText() antlr.TerminalNode

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_text
	return p
}

func InitEmptyTextContext(p *TextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_text
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) OpenInterpString() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenInterpString, 0)
}

func (s *TextContext) CloseInterpString() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseInterpString, 0)
}

func (s *TextContext) Bang() antlr.TerminalNode {
	return s.GetToken(RhumbParserBang, 0)
}

func (s *TextContext) AllInnerText() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserInnerText)
}

func (s *TextContext) InnerText(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserInnerText, i)
}

func (s *TextContext) AllInterpolation() []IInterpolationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInterpolationContext); ok {
			len++
		}
	}

	tst := make([]IInterpolationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInterpolationContext); ok {
			tst[i] = t.(IInterpolationContext)
			i++
		}
	}

	return tst
}

func (s *TextContext) Interpolation(i int) IInterpolationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterpolationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterpolationContext)
}

func (s *TextContext) RawText() antlr.TerminalNode {
	return s.GetToken(RhumbParserRawText, 0)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitText(s)
	}
}

func (s *TextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, RhumbParserRULE_text)
	var _la int

	p.SetState(634)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RhumbParserBang, RhumbParserOpenInterpString:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(622)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RhumbParserBang {
			{
				p.SetState(621)
				p.Match(RhumbParserBang)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(624)
			p.Match(RhumbParserOpenInterpString)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(629)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-85)) & ^0x3f) == 0 && ((int64(1)<<(_la-85))&39) != 0 {
			p.SetState(627)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case RhumbParserInnerText:
				{
					p.SetState(625)
					p.Match(RhumbParserInnerText)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case RhumbParserEnterRoutineInterp, RhumbParserEnterSelectorInterp, RhumbParserInterpLabel:
				{
					p.SetState(626)
					p.Interpolation()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(631)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(632)
			p.Match(RhumbParserCloseInterpString)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserRawText:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(633)
			p.Match(RhumbParserRawText)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInterpolationContext is an interface to support dynamic dispatch.
type IInterpolationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsInterpolationContext differentiates from other interfaces.
	IsInterpolationContext()
}

type InterpolationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterpolationContext() *InterpolationContext {
	var p = new(InterpolationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_interpolation
	return p
}

func InitEmptyInterpolationContext(p *InterpolationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_interpolation
}

func (*InterpolationContext) IsInterpolationContext() {}

func NewInterpolationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterpolationContext {
	var p = new(InterpolationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_interpolation

	return p
}

func (s *InterpolationContext) GetParser() antlr.Parser { return s.parser }

func (s *InterpolationContext) CopyAll(ctx *InterpolationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *InterpolationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterpolationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RoutineInterpContext struct {
	InterpolationContext
}

func NewRoutineInterpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoutineInterpContext {
	var p = new(RoutineInterpContext)

	InitEmptyInterpolationContext(&p.InterpolationContext)
	p.parser = parser
	p.CopyAll(ctx.(*InterpolationContext))

	return p
}

func (s *RoutineInterpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoutineInterpContext) EnterRoutineInterp() antlr.TerminalNode {
	return s.GetToken(RhumbParserEnterRoutineInterp, 0)
}

func (s *RoutineInterpContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *RoutineInterpContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseParen, 0)
}

func (s *RoutineInterpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterRoutineInterp(s)
	}
}

func (s *RoutineInterpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitRoutineInterp(s)
	}
}

func (s *RoutineInterpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitRoutineInterp(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelectorInterpContext struct {
	InterpolationContext
}

func NewSelectorInterpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectorInterpContext {
	var p = new(SelectorInterpContext)

	InitEmptyInterpolationContext(&p.InterpolationContext)
	p.parser = parser
	p.CopyAll(ctx.(*InterpolationContext))

	return p
}

func (s *SelectorInterpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorInterpContext) EnterSelectorInterp() antlr.TerminalNode {
	return s.GetToken(RhumbParserEnterSelectorInterp, 0)
}

func (s *SelectorInterpContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *SelectorInterpContext) CloseCurly() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseCurly, 0)
}

func (s *SelectorInterpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterSelectorInterp(s)
	}
}

func (s *SelectorInterpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitSelectorInterp(s)
	}
}

func (s *SelectorInterpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitSelectorInterp(s)

	default:
		return t.VisitChildren(s)
	}
}

type LabelInterpContext struct {
	InterpolationContext
}

func NewLabelInterpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LabelInterpContext {
	var p = new(LabelInterpContext)

	InitEmptyInterpolationContext(&p.InterpolationContext)
	p.parser = parser
	p.CopyAll(ctx.(*InterpolationContext))

	return p
}

func (s *LabelInterpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelInterpContext) InterpLabel() antlr.TerminalNode {
	return s.GetToken(RhumbParserInterpLabel, 0)
}

func (s *LabelInterpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterLabelInterp(s)
	}
}

func (s *LabelInterpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitLabelInterp(s)
	}
}

func (s *LabelInterpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitLabelInterp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) Interpolation() (localctx IInterpolationContext) {
	localctx = NewInterpolationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, RhumbParserRULE_interpolation)
	p.SetState(645)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RhumbParserInterpLabel:
		localctx = NewLabelInterpContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(636)
			p.Match(RhumbParserInterpLabel)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserEnterRoutineInterp:
		localctx = NewRoutineInterpContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(637)
			p.Match(RhumbParserEnterRoutineInterp)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(638)
			p.Expressions()
		}
		{
			p.SetState(639)
			p.Match(RhumbParserCloseParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RhumbParserEnterSelectorInterp:
		localctx = NewSelectorInterpContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(641)
			p.Match(RhumbParserEnterSelectorInterp)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(642)
			p.Expressions()
		}
		{
			p.SetState(643)
			p.Match(RhumbParserCloseCurly)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsReferenceContext differentiates from other interfaces.
	IsReferenceContext()
}

type ReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_reference
	return p
}

func InitEmptyReferenceContext(p *ReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RhumbParserRULE_reference
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) CopyAll(ctx *ReferenceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionRefContext struct {
	ReferenceContext
}

func NewFunctionRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionRefContext {
	var p = new(FunctionRefContext)

	InitEmptyReferenceContext(&p.ReferenceContext)
	p.parser = parser
	p.CopyAll(ctx.(*ReferenceContext))

	return p
}

func (s *FunctionRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionRefContext) OpenAnglet() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenAnglet, 0)
}

func (s *FunctionRefContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenParen, 0)
}

func (s *FunctionRefContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *FunctionRefContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseParen, 0)
}

func (s *FunctionRefContext) CloseAnglet() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseAnglet, 0)
}

func (s *FunctionRefContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *FunctionRefContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *FunctionRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterFunctionRef(s)
	}
}

func (s *FunctionRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitFunctionRef(s)
	}
}

func (s *FunctionRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitFunctionRef(s)

	default:
		return t.VisitChildren(s)
	}
}

type NamedRefContext struct {
	ReferenceContext
}

func NewNamedRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NamedRefContext {
	var p = new(NamedRefContext)

	InitEmptyReferenceContext(&p.ReferenceContext)
	p.parser = parser
	p.CopyAll(ctx.(*ReferenceContext))

	return p
}

func (s *NamedRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedRefContext) OpenAnglet() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenAnglet, 0)
}

func (s *NamedRefContext) Label() antlr.TerminalNode {
	return s.GetToken(RhumbParserLabel, 0)
}

func (s *NamedRefContext) CloseAnglet() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseAnglet, 0)
}

func (s *NamedRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNamedRef(s)
	}
}

func (s *NamedRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNamedRef(s)
	}
}

func (s *NamedRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNamedRef(s)

	default:
		return t.VisitChildren(s)
	}
}

type JunctionRefContext struct {
	ReferenceContext
}

func NewJunctionRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JunctionRefContext {
	var p = new(JunctionRefContext)

	InitEmptyReferenceContext(&p.ReferenceContext)
	p.parser = parser
	p.CopyAll(ctx.(*ReferenceContext))

	return p
}

func (s *JunctionRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JunctionRefContext) OpenAnglet() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenAnglet, 0)
}

func (s *JunctionRefContext) OpenCurly() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenCurly, 0)
}

func (s *JunctionRefContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *JunctionRefContext) CloseCurly() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseCurly, 0)
}

func (s *JunctionRefContext) CloseAnglet() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseAnglet, 0)
}

func (s *JunctionRefContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *JunctionRefContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *JunctionRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterJunctionRef(s)
	}
}

func (s *JunctionRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitJunctionRef(s)
	}
}

func (s *JunctionRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitJunctionRef(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComputedRefContext struct {
	ReferenceContext
}

func NewComputedRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComputedRefContext {
	var p = new(ComputedRefContext)

	InitEmptyReferenceContext(&p.ReferenceContext)
	p.parser = parser
	p.CopyAll(ctx.(*ReferenceContext))

	return p
}

func (s *ComputedRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComputedRefContext) OpenAnglet() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenAnglet, 0)
}

func (s *ComputedRefContext) Text() ITextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *ComputedRefContext) CloseAnglet() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseAnglet, 0)
}

func (s *ComputedRefContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenBracket, 0)
}

func (s *ComputedRefContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *ComputedRefContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseBracket, 0)
}

func (s *ComputedRefContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *ComputedRefContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *ComputedRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterComputedRef(s)
	}
}

func (s *ComputedRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitComputedRef(s)
	}
}

func (s *ComputedRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitComputedRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) Reference() (localctx IReferenceContext) {
	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, RhumbParserRULE_reference)
	var _alt int

	p.SetState(690)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 87, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNamedRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(647)
			p.Match(RhumbParserOpenAnglet)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(648)
			p.Match(RhumbParserLabel)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(649)
			p.Match(RhumbParserCloseAnglet)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewComputedRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(650)
			p.Match(RhumbParserOpenAnglet)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(651)
			p.Text()
		}
		{
			p.SetState(652)
			p.Match(RhumbParserCloseAnglet)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewFunctionRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(654)
			p.Match(RhumbParserOpenAnglet)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(655)
			p.Match(RhumbParserOpenParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(659)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 84, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(656)
					p.Terminator()
				}

			}
			p.SetState(661)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 84, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(662)
			p.Expressions()
		}
		{
			p.SetState(663)
			p.Match(RhumbParserCloseParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(664)
			p.Match(RhumbParserCloseAnglet)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewComputedRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(666)
			p.Match(RhumbParserOpenAnglet)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(667)
			p.Match(RhumbParserOpenBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(671)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 85, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(668)
					p.Terminator()
				}

			}
			p.SetState(673)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 85, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(674)
			p.Expressions()
		}
		{
			p.SetState(675)
			p.Match(RhumbParserCloseBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(676)
			p.Match(RhumbParserCloseAnglet)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewJunctionRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(678)
			p.Match(RhumbParserOpenAnglet)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(679)
			p.Match(RhumbParserOpenCurly)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(683)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 86, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(680)
					p.Terminator()
				}

			}
			p.SetState(685)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 86, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(686)
			p.Expressions()
		}
		{
			p.SetState(687)
			p.Match(RhumbParserCloseCurly)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(688)
			p.Match(RhumbParserCloseAnglet)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *RhumbParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RhumbParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 13)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
