// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // RhumbParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type RhumbParser struct {
	*antlr.BaseParser
}

var rhumbparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rhumbparserParserInit() {
	staticData := &rhumbparserParserStaticData
	staticData.literalNames = []string{
		"", "'&&'", "'&'", "'@'", "'\\'", "'`'", "'!'", "'!@'", "'!\\'", "'!!'",
		"'!^'", "'!='", "'!>'", "'!~'", "'^'", "'^^'", "'^/'", "':'", "'::'",
		"':='", "','", "',-'", "'-'", "'--'", "'-/'", "'->'", "'.'", "'.-'",
		"'..'", "'.='", "'='", "'=@'", "'=\\'", "'=^'", "'=='", "'=>'", "'=~'",
		"'/'", "'//'", "'>='", "'>>'", "'#'", "'<='", "'<>'", "'<<'", "'|'",
		"'|>'", "'||'", "'+'", "'+-'", "'+/'", "'++'", "'?'", "'??'", "';'",
		"'*'", "'*^'", "'**'", "'~'", "'~~'", "'~>'", "", "", "", "", "", "'('",
		"')'", "'['", "']'", "'{'", "'}'", "'<'", "'>'", "", "", "", "", "",
		"'#('", "'#{'", "", "'`#'", "'`\"'",
	}
	staticData.symbolicNames = []string{
		"", "AmpAmp", "Ampersand", "At", "BSlash", "Backtick", "Bang", "BangAt",
		"BangBSlash", "BangBang", "BangCaret", "BangEqual", "BangGreater", "BangTilde",
		"Caret", "CaretCaret", "CaretFSlash", "Colon", "ColonColon", "ColonEqual",
		"Comma", "CommaDash", "Dash", "DashDash", "DashFSlash", "DashGreater",
		"Dot", "DotDash", "DotDot", "DotEqual", "Equal", "EqualAt", "EqualBSlash",
		"EqualCaret", "EqualEqual", "EqualGreater", "EqualTilde", "FSlash",
		"FSlashFSlash", "GreaterEqual", "GreaterGreater", "Hash", "LesserEqual",
		"LesserGreater", "LesserLesser", "Pipe", "PipeGreater", "PipePipe",
		"Plus", "PlusDash", "PlusFSlash", "PlusPlus", "QMark", "QMarkQMark",
		"Semicolon", "Star", "StarCaret", "StarStar", "Tilde", "TildeTilde",
		"TildeGreater", "FloatingPoint", "NumberPart", "NumberStart", "RawString",
		"OpenInterpString", "OpenParen", "CloseParen", "OpenBracket", "CloseBracket",
		"OpenCurly", "CloseCurly", "OpenAnglet", "CloseAnglet", "NL", "WS",
		"BlockComment", "LineComment", "Label", "EnterRoutineInterp", "EnterSelectorInterp",
		"InterpLabel", "EscapedHash", "EscapedQuote", "RawText", "CloseInterpString",
	}
	staticData.ruleNames = []string{
		"sequence", "terminator", "literal", "expression", "accessOp", "applicativeOp",
		"comparativeOp", "identityOp", "conjunctiveOp", "disjunctiveOp", "conditionalOp",
		"additiveOp", "multiplicativeOp", "exponentiationOp", "assignmentOp",
		"prefixOp", "chainOp", "string", "interpolation", "reference",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 85, 361, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 0, 1,
		0, 3, 0, 44, 8, 0, 5, 0, 46, 8, 0, 10, 0, 12, 0, 49, 9, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 59, 8, 2, 1, 3, 1, 3, 1, 3, 5,
		3, 64, 8, 3, 10, 3, 12, 3, 67, 9, 3, 1, 3, 3, 3, 70, 8, 3, 1, 3, 1, 3,
		1, 3, 5, 3, 75, 8, 3, 10, 3, 12, 3, 78, 9, 3, 1, 3, 3, 3, 81, 8, 3, 1,
		3, 1, 3, 1, 3, 5, 3, 86, 8, 3, 10, 3, 12, 3, 89, 9, 3, 1, 3, 3, 3, 92,
		8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 100, 8, 3, 10, 3, 12, 3,
		103, 9, 3, 1, 3, 3, 3, 106, 8, 3, 1, 3, 1, 3, 3, 3, 110, 8, 3, 1, 3, 1,
		3, 1, 3, 5, 3, 115, 8, 3, 10, 3, 12, 3, 118, 9, 3, 1, 3, 3, 3, 121, 8,
		3, 1, 3, 1, 3, 3, 3, 125, 8, 3, 1, 3, 1, 3, 3, 3, 129, 8, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 135, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 176, 8, 3, 10, 3,
		12, 3, 179, 9, 3, 1, 3, 3, 3, 182, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3,
		188, 8, 3, 10, 3, 12, 3, 191, 9, 3, 1, 3, 1, 3, 3, 3, 195, 8, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 5, 3, 201, 8, 3, 10, 3, 12, 3, 204, 9, 3, 1, 3, 3, 3,
		207, 8, 3, 1, 3, 5, 3, 210, 8, 3, 10, 3, 12, 3, 213, 9, 3, 1, 4, 1, 4,
		1, 4, 3, 4, 218, 8, 4, 1, 5, 1, 5, 3, 5, 222, 8, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 3, 6, 228, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 3, 7, 240, 8, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 3, 10, 252, 8, 10, 1, 11, 1, 11, 1, 11, 3, 11, 257,
		8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 264, 8, 12, 1, 13, 1,
		13, 1, 13, 3, 13, 269, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 275, 8,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 3, 15, 290, 8, 15, 1, 16, 1, 16, 1, 16, 3, 16, 295,
		8, 16, 1, 17, 1, 17, 1, 17, 5, 17, 300, 8, 17, 10, 17, 12, 17, 303, 9,
		17, 1, 17, 1, 17, 3, 17, 307, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 318, 8, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 5, 19, 326, 8, 19, 10, 19, 12, 19, 329, 9, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 338, 8, 19, 10, 19, 12,
		19, 341, 9, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19,
		350, 8, 19, 10, 19, 12, 19, 353, 9, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3,
		19, 359, 8, 19, 1, 19, 0, 1, 6, 20, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 0, 1, 2, 0, 54, 54, 74, 74, 441,
		0, 40, 1, 0, 0, 0, 2, 50, 1, 0, 0, 0, 4, 58, 1, 0, 0, 0, 6, 134, 1, 0,
		0, 0, 8, 217, 1, 0, 0, 0, 10, 221, 1, 0, 0, 0, 12, 227, 1, 0, 0, 0, 14,
		239, 1, 0, 0, 0, 16, 241, 1, 0, 0, 0, 18, 243, 1, 0, 0, 0, 20, 251, 1,
		0, 0, 0, 22, 256, 1, 0, 0, 0, 24, 263, 1, 0, 0, 0, 26, 268, 1, 0, 0, 0,
		28, 274, 1, 0, 0, 0, 30, 289, 1, 0, 0, 0, 32, 294, 1, 0, 0, 0, 34, 306,
		1, 0, 0, 0, 36, 317, 1, 0, 0, 0, 38, 358, 1, 0, 0, 0, 40, 47, 3, 6, 3,
		0, 41, 43, 3, 2, 1, 0, 42, 44, 3, 6, 3, 0, 43, 42, 1, 0, 0, 0, 43, 44,
		1, 0, 0, 0, 44, 46, 1, 0, 0, 0, 45, 41, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0,
		47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 1, 1, 0, 0, 0, 49, 47, 1, 0,
		0, 0, 50, 51, 7, 0, 0, 0, 51, 3, 1, 0, 0, 0, 52, 59, 5, 61, 0, 0, 53, 59,
		5, 62, 0, 0, 54, 59, 3, 34, 17, 0, 55, 59, 3, 38, 19, 0, 56, 59, 5, 58,
		0, 0, 57, 59, 5, 78, 0, 0, 58, 52, 1, 0, 0, 0, 58, 53, 1, 0, 0, 0, 58,
		54, 1, 0, 0, 0, 58, 55, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 57, 1, 0, 0,
		0, 59, 5, 1, 0, 0, 0, 60, 61, 6, 3, -1, 0, 61, 65, 5, 68, 0, 0, 62, 64,
		3, 2, 1, 0, 63, 62, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0,
		65, 66, 1, 0, 0, 0, 66, 69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 70, 3,
		0, 0, 0, 69, 68, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71,
		135, 5, 69, 0, 0, 72, 76, 5, 70, 0, 0, 73, 75, 3, 2, 1, 0, 74, 73, 1, 0,
		0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 80,
		1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 81, 3, 0, 0, 0, 80, 79, 1, 0, 0, 0,
		80, 81, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 135, 5, 71, 0, 0, 83, 87, 5,
		66, 0, 0, 84, 86, 3, 2, 1, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87,
		85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0,
		0, 90, 92, 3, 0, 0, 0, 91, 90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93,
		1, 0, 0, 0, 93, 135, 5, 67, 0, 0, 94, 95, 3, 30, 15, 0, 95, 96, 3, 6, 3,
		11, 96, 135, 1, 0, 0, 0, 97, 101, 5, 66, 0, 0, 98, 100, 3, 2, 1, 0, 99,
		98, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1,
		0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 106, 3, 0, 0,
		0, 105, 104, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107,
		110, 5, 67, 0, 0, 108, 110, 3, 38, 19, 0, 109, 97, 1, 0, 0, 0, 109, 108,
		1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 124, 3, 10, 5, 0, 112, 116, 5, 66,
		0, 0, 113, 115, 3, 2, 1, 0, 114, 113, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0,
		116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118,
		116, 1, 0, 0, 0, 119, 121, 3, 0, 0, 0, 120, 119, 1, 0, 0, 0, 120, 121,
		1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 125, 5, 67, 0, 0, 123, 125, 3, 38,
		19, 0, 124, 112, 1, 0, 0, 0, 124, 123, 1, 0, 0, 0, 125, 135, 1, 0, 0, 0,
		126, 129, 3, 38, 19, 0, 127, 129, 5, 78, 0, 0, 128, 126, 1, 0, 0, 0, 128,
		127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 3, 28, 14, 0, 131, 132,
		3, 6, 3, 2, 132, 135, 1, 0, 0, 0, 133, 135, 3, 4, 2, 0, 134, 60, 1, 0,
		0, 0, 134, 72, 1, 0, 0, 0, 134, 83, 1, 0, 0, 0, 134, 94, 1, 0, 0, 0, 134,
		109, 1, 0, 0, 0, 134, 128, 1, 0, 0, 0, 134, 133, 1, 0, 0, 0, 135, 211,
		1, 0, 0, 0, 136, 137, 10, 16, 0, 0, 137, 138, 3, 32, 16, 0, 138, 139, 3,
		6, 3, 17, 139, 210, 1, 0, 0, 0, 140, 141, 10, 12, 0, 0, 141, 142, 3, 26,
		13, 0, 142, 143, 3, 6, 3, 12, 143, 210, 1, 0, 0, 0, 144, 145, 10, 10, 0,
		0, 145, 146, 3, 24, 12, 0, 146, 147, 3, 6, 3, 11, 147, 210, 1, 0, 0, 0,
		148, 149, 10, 9, 0, 0, 149, 150, 3, 22, 11, 0, 150, 151, 3, 6, 3, 10, 151,
		210, 1, 0, 0, 0, 152, 153, 10, 8, 0, 0, 153, 154, 3, 12, 6, 0, 154, 155,
		3, 6, 3, 9, 155, 210, 1, 0, 0, 0, 156, 157, 10, 7, 0, 0, 157, 158, 3, 16,
		8, 0, 158, 159, 3, 6, 3, 8, 159, 210, 1, 0, 0, 0, 160, 161, 10, 6, 0, 0,
		161, 162, 3, 18, 9, 0, 162, 163, 3, 6, 3, 7, 163, 210, 1, 0, 0, 0, 164,
		165, 10, 5, 0, 0, 165, 166, 3, 14, 7, 0, 166, 167, 3, 6, 3, 6, 167, 210,
		1, 0, 0, 0, 168, 169, 10, 4, 0, 0, 169, 170, 3, 20, 10, 0, 170, 171, 3,
		6, 3, 4, 171, 210, 1, 0, 0, 0, 172, 173, 10, 15, 0, 0, 173, 177, 5, 66,
		0, 0, 174, 176, 3, 2, 1, 0, 175, 174, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0,
		177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179,
		177, 1, 0, 0, 0, 180, 182, 3, 0, 0, 0, 181, 180, 1, 0, 0, 0, 181, 182,
		1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 210, 5, 67, 0, 0, 184, 185, 10,
		14, 0, 0, 185, 189, 5, 68, 0, 0, 186, 188, 3, 2, 1, 0, 187, 186, 1, 0,
		0, 0, 188, 191, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0,
		190, 194, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 192, 195, 3, 0, 0, 0, 193,
		195, 3, 8, 4, 0, 194, 192, 1, 0, 0, 0, 194, 193, 1, 0, 0, 0, 194, 195,
		1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 210, 5, 69, 0, 0, 197, 198, 10,
		13, 0, 0, 198, 202, 5, 70, 0, 0, 199, 201, 3, 2, 1, 0, 200, 199, 1, 0,
		0, 0, 201, 204, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0,
		203, 206, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 205, 207, 3, 0, 0, 0, 206,
		205, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 210,
		5, 71, 0, 0, 209, 136, 1, 0, 0, 0, 209, 140, 1, 0, 0, 0, 209, 144, 1, 0,
		0, 0, 209, 148, 1, 0, 0, 0, 209, 152, 1, 0, 0, 0, 209, 156, 1, 0, 0, 0,
		209, 160, 1, 0, 0, 0, 209, 164, 1, 0, 0, 0, 209, 168, 1, 0, 0, 0, 209,
		172, 1, 0, 0, 0, 209, 184, 1, 0, 0, 0, 209, 197, 1, 0, 0, 0, 210, 213,
		1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 7, 1, 0, 0,
		0, 213, 211, 1, 0, 0, 0, 214, 218, 5, 26, 0, 0, 215, 218, 5, 3, 0, 0, 216,
		218, 5, 55, 0, 0, 217, 214, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 217, 216,
		1, 0, 0, 0, 218, 9, 1, 0, 0, 0, 219, 222, 5, 25, 0, 0, 220, 222, 5, 60,
		0, 0, 221, 219, 1, 0, 0, 0, 221, 220, 1, 0, 0, 0, 222, 11, 1, 0, 0, 0,
		223, 228, 5, 40, 0, 0, 224, 228, 5, 39, 0, 0, 225, 228, 5, 44, 0, 0, 226,
		228, 5, 42, 0, 0, 227, 223, 1, 0, 0, 0, 227, 224, 1, 0, 0, 0, 227, 225,
		1, 0, 0, 0, 227, 226, 1, 0, 0, 0, 228, 13, 1, 0, 0, 0, 229, 240, 5, 34,
		0, 0, 230, 240, 5, 36, 0, 0, 231, 240, 5, 32, 0, 0, 232, 240, 5, 31, 0,
		0, 233, 240, 5, 33, 0, 0, 234, 240, 5, 11, 0, 0, 235, 240, 5, 13, 0, 0,
		236, 240, 5, 8, 0, 0, 237, 240, 5, 7, 0, 0, 238, 240, 5, 10, 0, 0, 239,
		229, 1, 0, 0, 0, 239, 230, 1, 0, 0, 0, 239, 231, 1, 0, 0, 0, 239, 232,
		1, 0, 0, 0, 239, 233, 1, 0, 0, 0, 239, 234, 1, 0, 0, 0, 239, 235, 1, 0,
		0, 0, 239, 236, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 239, 238, 1, 0, 0, 0,
		240, 15, 1, 0, 0, 0, 241, 242, 5, 1, 0, 0, 242, 17, 1, 0, 0, 0, 243, 244,
		5, 47, 0, 0, 244, 19, 1, 0, 0, 0, 245, 252, 5, 9, 0, 0, 246, 252, 5, 53,
		0, 0, 247, 252, 5, 43, 0, 0, 248, 252, 5, 46, 0, 0, 249, 252, 5, 35, 0,
		0, 250, 252, 5, 12, 0, 0, 251, 245, 1, 0, 0, 0, 251, 246, 1, 0, 0, 0, 251,
		247, 1, 0, 0, 0, 251, 248, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 251, 250,
		1, 0, 0, 0, 252, 21, 1, 0, 0, 0, 253, 257, 5, 51, 0, 0, 254, 257, 5, 49,
		0, 0, 255, 257, 5, 23, 0, 0, 256, 253, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0,
		256, 255, 1, 0, 0, 0, 257, 23, 1, 0, 0, 0, 258, 264, 5, 57, 0, 0, 259,
		264, 5, 38, 0, 0, 260, 264, 5, 50, 0, 0, 261, 264, 5, 24, 0, 0, 262, 264,
		5, 59, 0, 0, 263, 258, 1, 0, 0, 0, 263, 259, 1, 0, 0, 0, 263, 260, 1, 0,
		0, 0, 263, 261, 1, 0, 0, 0, 263, 262, 1, 0, 0, 0, 264, 25, 1, 0, 0, 0,
		265, 269, 5, 15, 0, 0, 266, 269, 5, 16, 0, 0, 267, 269, 5, 56, 0, 0, 268,
		265, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 268, 267, 1, 0, 0, 0, 269, 27, 1,
		0, 0, 0, 270, 275, 5, 28, 0, 0, 271, 275, 5, 29, 0, 0, 272, 275, 5, 18,
		0, 0, 273, 275, 5, 19, 0, 0, 274, 270, 1, 0, 0, 0, 274, 271, 1, 0, 0, 0,
		274, 272, 1, 0, 0, 0, 274, 273, 1, 0, 0, 0, 275, 29, 1, 0, 0, 0, 276, 290,
		5, 22, 0, 0, 277, 290, 5, 4, 0, 0, 278, 290, 5, 6, 0, 0, 279, 290, 5, 41,
		0, 0, 280, 290, 5, 14, 0, 0, 281, 290, 5, 2, 0, 0, 282, 290, 5, 55, 0,
		0, 283, 290, 5, 48, 0, 0, 284, 290, 5, 30, 0, 0, 285, 290, 5, 3, 0, 0,
		286, 290, 5, 52, 0, 0, 287, 290, 5, 26, 0, 0, 288, 290, 5, 17, 0, 0, 289,
		276, 1, 0, 0, 0, 289, 277, 1, 0, 0, 0, 289, 278, 1, 0, 0, 0, 289, 279,
		1, 0, 0, 0, 289, 280, 1, 0, 0, 0, 289, 281, 1, 0, 0, 0, 289, 282, 1, 0,
		0, 0, 289, 283, 1, 0, 0, 0, 289, 284, 1, 0, 0, 0, 289, 285, 1, 0, 0, 0,
		289, 286, 1, 0, 0, 0, 289, 287, 1, 0, 0, 0, 289, 288, 1, 0, 0, 0, 290,
		31, 1, 0, 0, 0, 291, 295, 5, 4, 0, 0, 292, 295, 5, 3, 0, 0, 293, 295, 5,
		45, 0, 0, 294, 291, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 294, 293, 1, 0, 0,
		0, 295, 33, 1, 0, 0, 0, 296, 301, 5, 65, 0, 0, 297, 300, 5, 84, 0, 0, 298,
		300, 3, 36, 18, 0, 299, 297, 1, 0, 0, 0, 299, 298, 1, 0, 0, 0, 300, 303,
		1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 304, 1, 0,
		0, 0, 303, 301, 1, 0, 0, 0, 304, 307, 5, 85, 0, 0, 305, 307, 5, 64, 0,
		0, 306, 296, 1, 0, 0, 0, 306, 305, 1, 0, 0, 0, 307, 35, 1, 0, 0, 0, 308,
		318, 5, 81, 0, 0, 309, 310, 5, 79, 0, 0, 310, 311, 3, 0, 0, 0, 311, 312,
		5, 67, 0, 0, 312, 318, 1, 0, 0, 0, 313, 314, 5, 80, 0, 0, 314, 315, 3,
		0, 0, 0, 315, 316, 5, 71, 0, 0, 316, 318, 1, 0, 0, 0, 317, 308, 1, 0, 0,
		0, 317, 309, 1, 0, 0, 0, 317, 313, 1, 0, 0, 0, 318, 37, 1, 0, 0, 0, 319,
		320, 5, 72, 0, 0, 320, 321, 5, 78, 0, 0, 321, 359, 5, 73, 0, 0, 322, 323,
		5, 72, 0, 0, 323, 327, 5, 66, 0, 0, 324, 326, 3, 2, 1, 0, 325, 324, 1,
		0, 0, 0, 326, 329, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 327, 328, 1, 0, 0,
		0, 328, 330, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 330, 331, 3, 0, 0, 0, 331,
		332, 5, 67, 0, 0, 332, 333, 5, 73, 0, 0, 333, 359, 1, 0, 0, 0, 334, 335,
		5, 72, 0, 0, 335, 339, 5, 68, 0, 0, 336, 338, 3, 2, 1, 0, 337, 336, 1,
		0, 0, 0, 338, 341, 1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 339, 340, 1, 0, 0,
		0, 340, 342, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 342, 343, 3, 0, 0, 0, 343,
		344, 5, 69, 0, 0, 344, 345, 5, 73, 0, 0, 345, 359, 1, 0, 0, 0, 346, 347,
		5, 72, 0, 0, 347, 351, 5, 70, 0, 0, 348, 350, 3, 2, 1, 0, 349, 348, 1,
		0, 0, 0, 350, 353, 1, 0, 0, 0, 351, 349, 1, 0, 0, 0, 351, 352, 1, 0, 0,
		0, 352, 354, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 354, 355, 3, 0, 0, 0, 355,
		356, 5, 71, 0, 0, 356, 357, 5, 73, 0, 0, 357, 359, 1, 0, 0, 0, 358, 319,
		1, 0, 0, 0, 358, 322, 1, 0, 0, 0, 358, 334, 1, 0, 0, 0, 358, 346, 1, 0,
		0, 0, 359, 39, 1, 0, 0, 0, 44, 43, 47, 58, 65, 69, 76, 80, 87, 91, 101,
		105, 109, 116, 120, 124, 128, 134, 177, 181, 189, 194, 202, 206, 209, 211,
		217, 221, 227, 239, 251, 256, 263, 268, 274, 289, 294, 299, 301, 306, 317,
		327, 339, 351, 358,
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
	staticData := &rhumbparserParserStaticData
	staticData.once.Do(rhumbparserParserInit)
}

// NewRhumbParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRhumbParser(input antlr.TokenStream) *RhumbParser {
	RhumbParserInit()
	this := new(RhumbParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rhumbparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// RhumbParser tokens.
const (
	RhumbParserEOF                 = antlr.TokenEOF
	RhumbParserAmpAmp              = 1
	RhumbParserAmpersand           = 2
	RhumbParserAt                  = 3
	RhumbParserBSlash              = 4
	RhumbParserBacktick            = 5
	RhumbParserBang                = 6
	RhumbParserBangAt              = 7
	RhumbParserBangBSlash          = 8
	RhumbParserBangBang            = 9
	RhumbParserBangCaret           = 10
	RhumbParserBangEqual           = 11
	RhumbParserBangGreater         = 12
	RhumbParserBangTilde           = 13
	RhumbParserCaret               = 14
	RhumbParserCaretCaret          = 15
	RhumbParserCaretFSlash         = 16
	RhumbParserColon               = 17
	RhumbParserColonColon          = 18
	RhumbParserColonEqual          = 19
	RhumbParserComma               = 20
	RhumbParserCommaDash           = 21
	RhumbParserDash                = 22
	RhumbParserDashDash            = 23
	RhumbParserDashFSlash          = 24
	RhumbParserDashGreater         = 25
	RhumbParserDot                 = 26
	RhumbParserDotDash             = 27
	RhumbParserDotDot              = 28
	RhumbParserDotEqual            = 29
	RhumbParserEqual               = 30
	RhumbParserEqualAt             = 31
	RhumbParserEqualBSlash         = 32
	RhumbParserEqualCaret          = 33
	RhumbParserEqualEqual          = 34
	RhumbParserEqualGreater        = 35
	RhumbParserEqualTilde          = 36
	RhumbParserFSlash              = 37
	RhumbParserFSlashFSlash        = 38
	RhumbParserGreaterEqual        = 39
	RhumbParserGreaterGreater      = 40
	RhumbParserHash                = 41
	RhumbParserLesserEqual         = 42
	RhumbParserLesserGreater       = 43
	RhumbParserLesserLesser        = 44
	RhumbParserPipe                = 45
	RhumbParserPipeGreater         = 46
	RhumbParserPipePipe            = 47
	RhumbParserPlus                = 48
	RhumbParserPlusDash            = 49
	RhumbParserPlusFSlash          = 50
	RhumbParserPlusPlus            = 51
	RhumbParserQMark               = 52
	RhumbParserQMarkQMark          = 53
	RhumbParserSemicolon           = 54
	RhumbParserStar                = 55
	RhumbParserStarCaret           = 56
	RhumbParserStarStar            = 57
	RhumbParserTilde               = 58
	RhumbParserTildeTilde          = 59
	RhumbParserTildeGreater        = 60
	RhumbParserFloatingPoint       = 61
	RhumbParserNumberPart          = 62
	RhumbParserNumberStart         = 63
	RhumbParserRawString           = 64
	RhumbParserOpenInterpString    = 65
	RhumbParserOpenParen           = 66
	RhumbParserCloseParen          = 67
	RhumbParserOpenBracket         = 68
	RhumbParserCloseBracket        = 69
	RhumbParserOpenCurly           = 70
	RhumbParserCloseCurly          = 71
	RhumbParserOpenAnglet          = 72
	RhumbParserCloseAnglet         = 73
	RhumbParserNL                  = 74
	RhumbParserWS                  = 75
	RhumbParserBlockComment        = 76
	RhumbParserLineComment         = 77
	RhumbParserLabel               = 78
	RhumbParserEnterRoutineInterp  = 79
	RhumbParserEnterSelectorInterp = 80
	RhumbParserInterpLabel         = 81
	RhumbParserEscapedHash         = 82
	RhumbParserEscapedQuote        = 83
	RhumbParserRawText             = 84
	RhumbParserCloseInterpString   = 85
)

// RhumbParser rules.
const (
	RhumbParserRULE_sequence         = 0
	RhumbParserRULE_terminator       = 1
	RhumbParserRULE_literal          = 2
	RhumbParserRULE_expression       = 3
	RhumbParserRULE_accessOp         = 4
	RhumbParserRULE_applicativeOp    = 5
	RhumbParserRULE_comparativeOp    = 6
	RhumbParserRULE_identityOp       = 7
	RhumbParserRULE_conjunctiveOp    = 8
	RhumbParserRULE_disjunctiveOp    = 9
	RhumbParserRULE_conditionalOp    = 10
	RhumbParserRULE_additiveOp       = 11
	RhumbParserRULE_multiplicativeOp = 12
	RhumbParserRULE_exponentiationOp = 13
	RhumbParserRULE_assignmentOp     = 14
	RhumbParserRULE_prefixOp         = 15
	RhumbParserRULE_chainOp          = 16
	RhumbParserRULE_string           = 17
	RhumbParserRULE_interpolation    = 18
	RhumbParserRULE_reference        = 19
)

// ISequenceContext is an interface to support dynamic dispatch.
type ISequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequenceContext differentiates from other interfaces.
	IsSequenceContext()
}

type SequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequenceContext() *SequenceContext {
	var p = new(SequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_sequence
	return p
}

func (*SequenceContext) IsSequenceContext() {}

func NewSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequenceContext {
	var p = new(SequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_sequence

	return p
}

func (s *SequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SequenceContext) AllExpression() []IExpressionContext {
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

func (s *SequenceContext) Expression(i int) IExpressionContext {
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

func (s *SequenceContext) AllTerminator() []ITerminatorContext {
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

func (s *SequenceContext) Terminator(i int) ITerminatorContext {
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

func (s *SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterSequence(s)
	}
}

func (s *SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitSequence(s)
	}
}

func (s *SequenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitSequence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) Sequence() (localctx ISequenceContext) {
	this := p
	_ = this

	localctx = NewSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RhumbParserRULE_sequence)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.expression(0)
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RhumbParserSemicolon || _la == RhumbParserNL {
		{
			p.SetState(41)
			p.Terminator()
		}
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7246575475584286812) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&16727) != 0 {
			{
				p.SetState(42)
				p.expression(0)
			}

		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITerminatorContext is an interface to support dynamic dispatch.
type ITerminatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerminatorContext differentiates from other interfaces.
	IsTerminatorContext()
}

type TerminatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerminatorContext() *TerminatorContext {
	var p = new(TerminatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_terminator
	return p
}

func (*TerminatorContext) IsTerminatorContext() {}

func NewTerminatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TerminatorContext {
	var p = new(TerminatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewTerminatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RhumbParserRULE_terminator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RhumbParserSemicolon || _la == RhumbParserNL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyFrom(ctx *LiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ReferenceLiteralContext struct {
	*LiteralContext
}

func NewReferenceLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReferenceLiteralContext {
	var p = new(ReferenceLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

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

type LabelLiteralContext struct {
	*LiteralContext
}

func NewLabelLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LabelLiteralContext {
	var p = new(LabelLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *LabelLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelLiteralContext) Tilde() antlr.TerminalNode {
	return s.GetToken(RhumbParserTilde, 0)
}

func (s *LabelLiteralContext) Label() antlr.TerminalNode {
	return s.GetToken(RhumbParserLabel, 0)
}

func (s *LabelLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterLabelLiteral(s)
	}
}

func (s *LabelLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitLabelLiteral(s)
	}
}

func (s *LabelLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitLabelLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringLiteralContext struct {
	*LiteralContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatLiteralContext struct {
	*LiteralContext
}

func NewFloatLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) FloatingPoint() antlr.TerminalNode {
	return s.GetToken(RhumbParserFloatingPoint, 0)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntegerLiteralContext struct {
	*LiteralContext
}

func NewIntegerLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) NumberPart() antlr.TerminalNode {
	return s.GetToken(RhumbParserNumberPart, 0)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RhumbParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(58)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RhumbParserFloatingPoint:
		localctx = NewFloatLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(RhumbParserFloatingPoint)
		}

	case RhumbParserNumberPart:
		localctx = NewIntegerLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Match(RhumbParserNumberPart)
		}

	case RhumbParserRawString, RhumbParserOpenInterpString:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.String_()
		}

	case RhumbParserOpenAnglet:
		localctx = NewReferenceLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(55)
			p.Reference()
		}

	case RhumbParserTilde:
		localctx = NewLabelLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(56)
			p.Match(RhumbParserTilde)
		}

	case RhumbParserLabel:
		localctx = NewLabelLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(57)
			p.Match(RhumbParserLabel)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConjunctiveContext struct {
	*ExpressionContext
}

func NewConjunctiveContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConjunctiveContext {
	var p = new(ConjunctiveContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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
	*ExpressionContext
}

func NewAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccessContext {
	var p = new(AccessContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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

func (s *AccessContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
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
	*ExpressionContext
	params    ISequenceContext
	paramsRef IReferenceContext
	body      ISequenceContext
	bodyRef   IReferenceContext
}

func NewApplicativeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ApplicativeContext {
	var p = new(ApplicativeContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ApplicativeContext) GetParams() ISequenceContext { return s.params }

func (s *ApplicativeContext) GetParamsRef() IReferenceContext { return s.paramsRef }

func (s *ApplicativeContext) GetBody() ISequenceContext { return s.body }

func (s *ApplicativeContext) GetBodyRef() IReferenceContext { return s.bodyRef }

func (s *ApplicativeContext) SetParams(v ISequenceContext) { s.params = v }

func (s *ApplicativeContext) SetParamsRef(v IReferenceContext) { s.paramsRef = v }

func (s *ApplicativeContext) SetBody(v ISequenceContext) { s.body = v }

func (s *ApplicativeContext) SetBodyRef(v IReferenceContext) { s.bodyRef = v }

func (s *ApplicativeContext) GetRuleContext() antlr.RuleContext {
	return s
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

func (s *ApplicativeContext) AllOpenParen() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserOpenParen)
}

func (s *ApplicativeContext) OpenParen(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenParen, i)
}

func (s *ApplicativeContext) AllCloseParen() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserCloseParen)
}

func (s *ApplicativeContext) CloseParen(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseParen, i)
}

func (s *ApplicativeContext) AllReference() []IReferenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReferenceContext); ok {
			len++
		}
	}

	tst := make([]IReferenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReferenceContext); ok {
			tst[i] = t.(IReferenceContext)
			i++
		}
	}

	return tst
}

func (s *ApplicativeContext) Reference(i int) IReferenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
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

	return t.(IReferenceContext)
}

func (s *ApplicativeContext) AllTerminator() []ITerminatorContext {
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

func (s *ApplicativeContext) Terminator(i int) ITerminatorContext {
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

func (s *ApplicativeContext) AllSequence() []ISequenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISequenceContext); ok {
			len++
		}
	}

	tst := make([]ISequenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISequenceContext); ok {
			tst[i] = t.(ISequenceContext)
			i++
		}
	}

	return tst
}

func (s *ApplicativeContext) Sequence(i int) ISequenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
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

	return t.(ISequenceContext)
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
	*ExpressionContext
}

func NewConditionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionalContext {
	var p = new(ConditionalContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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
	*ExpressionContext
}

func NewPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrefixContext {
	var p = new(PrefixContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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

type AssignmentContext struct {
	*ExpressionContext
	addressRef IReferenceContext
	address    antlr.Token
}

func NewAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentContext {
	var p = new(AssignmentContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AssignmentContext) GetAddress() antlr.Token { return s.address }

func (s *AssignmentContext) SetAddress(v antlr.Token) { s.address = v }

func (s *AssignmentContext) GetAddressRef() IReferenceContext { return s.addressRef }

func (s *AssignmentContext) SetAddressRef(v IReferenceContext) { s.addressRef = v }

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) AssignmentOp() IAssignmentOpContext {
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

func (s *AssignmentContext) Expression() IExpressionContext {
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

func (s *AssignmentContext) Reference() IReferenceContext {
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

func (s *AssignmentContext) Label() antlr.TerminalNode {
	return s.GetToken(RhumbParserLabel, 0)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparativeContext struct {
	*ExpressionContext
}

func NewComparativeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparativeContext {
	var p = new(ComparativeContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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

type SimpleContext struct {
	*ExpressionContext
}

func NewSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleContext {
	var p = new(SimpleContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleContext) Literal() ILiteralContext {
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

func (s *SimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterSimple(s)
	}
}

func (s *SimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitSimple(s)
	}
}

func (s *SimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitSimple(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiplicativeContext struct {
	*ExpressionContext
}

func NewMultiplicativeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeContext {
	var p = new(MultiplicativeContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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
	*ExpressionContext
}

func NewAdditiveContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveContext {
	var p = new(AdditiveContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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
	*ExpressionContext
}

func NewInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvocationContext {
	var p = new(InvocationContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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

func (s *InvocationContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
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

type RoutineContext struct {
	*ExpressionContext
}

func NewRoutineContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoutineContext {
	var p = new(RoutineContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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

func (s *RoutineContext) AllTerminator() []ITerminatorContext {
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

func (s *RoutineContext) Terminator(i int) ITerminatorContext {
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

func (s *RoutineContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
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
	*ExpressionContext
}

func NewDisjunctiveContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DisjunctiveContext {
	var p = new(DisjunctiveContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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
	*ExpressionContext
}

func NewIdentityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentityContext {
	var p = new(IdentityContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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

type EffectContext struct {
	*ExpressionContext
}

func NewEffectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EffectContext {
	var p = new(EffectContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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

func (s *EffectContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
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
	*ExpressionContext
}

func NewMemberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberContext {
	var p = new(MemberContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberContext) AllExpression() []IExpressionContext {
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

func (s *MemberContext) Expression(i int) IExpressionContext {
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

func (s *MemberContext) ChainOp() IChainOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChainOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChainOpContext)
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
	*ExpressionContext
}

func NewSelectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectorContext {
	var p = new(SelectorContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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

func (s *SelectorContext) AllTerminator() []ITerminatorContext {
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

func (s *SelectorContext) Terminator(i int) ITerminatorContext {
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

func (s *SelectorContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
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
	*ExpressionContext
}

func NewPowerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerContext {
	var p = new(PowerContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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
	*ExpressionContext
}

func NewMapContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapContext {
	var p = new(MapContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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

func (s *MapContext) AllTerminator() []ITerminatorContext {
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

func (s *MapContext) Terminator(i int) ITerminatorContext {
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

func (s *MapContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
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
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, RhumbParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMapContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(61)
			p.Match(RhumbParserOpenBracket)
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserSemicolon || _la == RhumbParserNL {
			{
				p.SetState(62)
				p.Terminator()
			}

			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7246575475584286812) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&16727) != 0 {
			{
				p.SetState(68)
				p.Sequence()
			}

		}
		{
			p.SetState(71)
			p.Match(RhumbParserCloseBracket)
		}

	case 2:
		localctx = NewSelectorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(72)
			p.Match(RhumbParserOpenCurly)
		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserSemicolon || _la == RhumbParserNL {
			{
				p.SetState(73)
				p.Terminator()
			}

			p.SetState(78)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7246575475584286812) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&16727) != 0 {
			{
				p.SetState(79)
				p.Sequence()
			}

		}
		{
			p.SetState(82)
			p.Match(RhumbParserCloseCurly)
		}

	case 3:
		localctx = NewRoutineContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(83)
			p.Match(RhumbParserOpenParen)
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserSemicolon || _la == RhumbParserNL {
			{
				p.SetState(84)
				p.Terminator()
			}

			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7246575475584286812) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&16727) != 0 {
			{
				p.SetState(90)
				p.Sequence()
			}

		}
		{
			p.SetState(93)
			p.Match(RhumbParserCloseParen)
		}

	case 4:
		localctx = NewPrefixContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(94)
			p.PrefixOp()
		}
		{
			p.SetState(95)
			p.expression(11)
		}

	case 5:
		localctx = NewApplicativeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(109)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case RhumbParserOpenParen:
			{
				p.SetState(97)
				p.Match(RhumbParserOpenParen)
			}
			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == RhumbParserSemicolon || _la == RhumbParserNL {
				{
					p.SetState(98)
					p.Terminator()
				}

				p.SetState(103)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7246575475584286812) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&16727) != 0 {
				{
					p.SetState(104)

					var _x = p.Sequence()

					localctx.(*ApplicativeContext).params = _x
				}

			}
			{
				p.SetState(107)
				p.Match(RhumbParserCloseParen)
			}

		case RhumbParserOpenAnglet:
			{
				p.SetState(108)

				var _x = p.Reference()

				localctx.(*ApplicativeContext).paramsRef = _x
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(111)
			p.ApplicativeOp()
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case RhumbParserOpenParen:
			{
				p.SetState(112)
				p.Match(RhumbParserOpenParen)
			}
			p.SetState(116)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == RhumbParserSemicolon || _la == RhumbParserNL {
				{
					p.SetState(113)
					p.Terminator()
				}

				p.SetState(118)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7246575475584286812) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&16727) != 0 {
				{
					p.SetState(119)

					var _x = p.Sequence()

					localctx.(*ApplicativeContext).body = _x
				}

			}
			{
				p.SetState(122)
				p.Match(RhumbParserCloseParen)
			}

		case RhumbParserOpenAnglet:
			{
				p.SetState(123)

				var _x = p.Reference()

				localctx.(*ApplicativeContext).bodyRef = _x
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case 6:
		localctx = NewAssignmentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(128)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case RhumbParserOpenAnglet:
			{
				p.SetState(126)

				var _x = p.Reference()

				localctx.(*AssignmentContext).addressRef = _x
			}

		case RhumbParserLabel:
			{
				p.SetState(127)

				var _m = p.Match(RhumbParserLabel)

				localctx.(*AssignmentContext).address = _m
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(130)
			p.AssignmentOp()
		}
		{
			p.SetState(131)
			p.expression(2)
		}

	case 7:
		localctx = NewSimpleContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(133)
			p.Literal()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMemberContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(136)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(137)
					p.ChainOp()
				}
				{
					p.SetState(138)
					p.expression(17)
				}

			case 2:
				localctx = NewPowerContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(140)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(141)
					p.ExponentiationOp()
				}
				{
					p.SetState(142)
					p.expression(12)
				}

			case 3:
				localctx = NewMultiplicativeContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(144)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(145)
					p.MultiplicativeOp()
				}
				{
					p.SetState(146)
					p.expression(11)
				}

			case 4:
				localctx = NewAdditiveContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(148)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(149)
					p.AdditiveOp()
				}
				{
					p.SetState(150)
					p.expression(10)
				}

			case 5:
				localctx = NewComparativeContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(153)
					p.ComparativeOp()
				}
				{
					p.SetState(154)
					p.expression(9)
				}

			case 6:
				localctx = NewConjunctiveContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(156)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(157)
					p.ConjunctiveOp()
				}
				{
					p.SetState(158)
					p.expression(8)
				}

			case 7:
				localctx = NewDisjunctiveContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(160)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(161)
					p.DisjunctiveOp()
				}
				{
					p.SetState(162)
					p.expression(7)
				}

			case 8:
				localctx = NewIdentityContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(164)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(165)
					p.IdentityOp()
				}
				{
					p.SetState(166)
					p.expression(6)
				}

			case 9:
				localctx = NewConditionalContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(168)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(169)
					p.ConditionalOp()
				}
				{
					p.SetState(170)
					p.expression(4)
				}

			case 10:
				localctx = NewInvocationContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(172)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(173)
					p.Match(RhumbParserOpenParen)
				}
				p.SetState(177)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserSemicolon || _la == RhumbParserNL {
					{
						p.SetState(174)
						p.Terminator()
					}

					p.SetState(179)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				p.SetState(181)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7246575475584286812) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&16727) != 0 {
					{
						p.SetState(180)
						p.Sequence()
					}

				}
				{
					p.SetState(183)
					p.Match(RhumbParserCloseParen)
				}

			case 11:
				localctx = NewAccessContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(184)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(185)
					p.Match(RhumbParserOpenBracket)
				}
				p.SetState(189)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserSemicolon || _la == RhumbParserNL {
					{
						p.SetState(186)
						p.Terminator()
					}

					p.SetState(191)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				p.SetState(194)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(192)
						p.Sequence()
					}

				} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 2 {
					{
						p.SetState(193)
						p.AccessOp()
					}

				}
				{
					p.SetState(196)
					p.Match(RhumbParserCloseBracket)
				}

			case 12:
				localctx = NewEffectContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RhumbParserRULE_expression)
				p.SetState(197)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(198)
					p.Match(RhumbParserOpenCurly)
				}
				p.SetState(202)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == RhumbParserSemicolon || _la == RhumbParserNL {
					{
						p.SetState(199)
						p.Terminator()
					}

					p.SetState(204)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				p.SetState(206)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7246575475584286812) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&16727) != 0 {
					{
						p.SetState(205)
						p.Sequence()
					}

				}
				{
					p.SetState(208)
					p.Match(RhumbParserCloseCurly)
				}

			}

		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessOpContext() *AccessOpContext {
	var p = new(AccessOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_accessOp
	return p
}

func (*AccessOpContext) IsAccessOpContext() {}

func NewAccessOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessOpContext {
	var p = new(AccessOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_accessOp

	return p
}

func (s *AccessOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessOpContext) CopyFrom(ctx *AccessOpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AccessOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FreezeContext struct {
	*AccessOpContext
}

func NewFreezeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FreezeContext {
	var p = new(FreezeContext)

	p.AccessOpContext = NewEmptyAccessOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AccessOpContext))

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

type LengthContext struct {
	*AccessOpContext
}

func NewLengthContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LengthContext {
	var p = new(LengthContext)

	p.AccessOpContext = NewEmptyAccessOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AccessOpContext))

	return p
}

func (s *LengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthContext) Star() antlr.TerminalNode {
	return s.GetToken(RhumbParserStar, 0)
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

type InnerContext struct {
	*AccessOpContext
}

func NewInnerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InnerContext {
	var p = new(InnerContext)

	p.AccessOpContext = NewEmptyAccessOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AccessOpContext))

	return p
}

func (s *InnerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerContext) At() antlr.TerminalNode {
	return s.GetToken(RhumbParserAt, 0)
}

func (s *InnerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterInner(s)
	}
}

func (s *InnerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitInner(s)
	}
}

func (s *InnerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitInner(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) AccessOp() (localctx IAccessOpContext) {
	this := p
	_ = this

	localctx = NewAccessOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RhumbParserRULE_accessOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(217)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RhumbParserDot:
		localctx = NewFreezeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Match(RhumbParserDot)
		}

	case RhumbParserAt:
		localctx = NewInnerContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Match(RhumbParserAt)
		}

	case RhumbParserStar:
		localctx = NewLengthContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)
			p.Match(RhumbParserStar)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApplicativeOpContext() *ApplicativeOpContext {
	var p = new(ApplicativeOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_applicativeOp
	return p
}

func (*ApplicativeOpContext) IsApplicativeOpContext() {}

func NewApplicativeOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApplicativeOpContext {
	var p = new(ApplicativeOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_applicativeOp

	return p
}

func (s *ApplicativeOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ApplicativeOpContext) CopyFrom(ctx *ApplicativeOpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ApplicativeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplicativeOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type JunctionContext struct {
	*ApplicativeOpContext
}

func NewJunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JunctionContext {
	var p = new(JunctionContext)

	p.ApplicativeOpContext = NewEmptyApplicativeOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ApplicativeOpContext))

	return p
}

func (s *JunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JunctionContext) TildeGreater() antlr.TerminalNode {
	return s.GetToken(RhumbParserTildeGreater, 0)
}

func (s *JunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterJunction(s)
	}
}

func (s *JunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitJunction(s)
	}
}

func (s *JunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitJunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionContext struct {
	*ApplicativeOpContext
}

func NewFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionContext {
	var p = new(FunctionContext)

	p.ApplicativeOpContext = NewEmptyApplicativeOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ApplicativeOpContext))

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
	this := p
	_ = this

	localctx = NewApplicativeOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RhumbParserRULE_applicativeOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(221)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RhumbParserDashGreater:
		localctx = NewFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)
			p.Match(RhumbParserDashGreater)
		}

	case RhumbParserTildeGreater:
		localctx = NewJunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.Match(RhumbParserTildeGreater)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparativeOpContext() *ComparativeOpContext {
	var p = new(ComparativeOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_comparativeOp
	return p
}

func (*ComparativeOpContext) IsComparativeOpContext() {}

func NewComparativeOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparativeOpContext {
	var p = new(ComparativeOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_comparativeOp

	return p
}

func (s *ComparativeOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparativeOpContext) CopyFrom(ctx *ComparativeOpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ComparativeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparativeOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LessThanContext struct {
	*ComparativeOpContext
}

func NewLessThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanContext {
	var p = new(LessThanContext)

	p.ComparativeOpContext = NewEmptyComparativeOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparativeOpContext))

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
	*ComparativeOpContext
}

func NewGreaterThanOrEqualToContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanOrEqualToContext {
	var p = new(GreaterThanOrEqualToContext)

	p.ComparativeOpContext = NewEmptyComparativeOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparativeOpContext))

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
	*ComparativeOpContext
}

func NewLessThanOrEqualToContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanOrEqualToContext {
	var p = new(LessThanOrEqualToContext)

	p.ComparativeOpContext = NewEmptyComparativeOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparativeOpContext))

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
	*ComparativeOpContext
}

func NewGreaterThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanContext {
	var p = new(GreaterThanContext)

	p.ComparativeOpContext = NewEmptyComparativeOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparativeOpContext))

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
	this := p
	_ = this

	localctx = NewComparativeOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RhumbParserRULE_comparativeOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(227)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RhumbParserGreaterGreater:
		localctx = NewGreaterThanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.Match(RhumbParserGreaterGreater)
		}

	case RhumbParserGreaterEqual:
		localctx = NewGreaterThanOrEqualToContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.Match(RhumbParserGreaterEqual)
		}

	case RhumbParserLesserLesser:
		localctx = NewLessThanContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(225)
			p.Match(RhumbParserLesserLesser)
		}

	case RhumbParserLesserEqual:
		localctx = NewLessThanOrEqualToContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(226)
			p.Match(RhumbParserLesserEqual)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentityOpContext() *IdentityOpContext {
	var p = new(IdentityOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_identityOp
	return p
}

func (*IdentityOpContext) IsIdentityOpContext() {}

func NewIdentityOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentityOpContext {
	var p = new(IdentityOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_identityOp

	return p
}

func (s *IdentityOpContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentityOpContext) CopyFrom(ctx *IdentityOpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *IdentityOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentityOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NotOverlayedContext struct {
	*IdentityOpContext
}

func NewNotOverlayedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotOverlayedContext {
	var p = new(NotOverlayedContext)

	p.IdentityOpContext = NewEmptyIdentityOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentityOpContext))

	return p
}

func (s *NotOverlayedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotOverlayedContext) BangAt() antlr.TerminalNode {
	return s.GetToken(RhumbParserBangAt, 0)
}

func (s *NotOverlayedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNotOverlayed(s)
	}
}

func (s *NotOverlayedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNotOverlayed(s)
	}
}

func (s *NotOverlayedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNotOverlayed(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotTopmostContext struct {
	*IdentityOpContext
}

func NewNotTopmostContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotTopmostContext {
	var p = new(NotTopmostContext)

	p.IdentityOpContext = NewEmptyIdentityOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentityOpContext))

	return p
}

func (s *NotTopmostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotTopmostContext) BangCaret() antlr.TerminalNode {
	return s.GetToken(RhumbParserBangCaret, 0)
}

func (s *NotTopmostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNotTopmost(s)
	}
}

func (s *NotTopmostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNotTopmost(s)
	}
}

func (s *NotTopmostContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNotTopmost(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsLikeContext struct {
	*IdentityOpContext
}

func NewIsLikeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsLikeContext {
	var p = new(IsLikeContext)

	p.IdentityOpContext = NewEmptyIdentityOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentityOpContext))

	return p
}

func (s *IsLikeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsLikeContext) EqualTilde() antlr.TerminalNode {
	return s.GetToken(RhumbParserEqualTilde, 0)
}

func (s *IsLikeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterIsLike(s)
	}
}

func (s *IsLikeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitIsLike(s)
	}
}

func (s *IsLikeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitIsLike(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsOverlayedContext struct {
	*IdentityOpContext
}

func NewIsOverlayedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsOverlayedContext {
	var p = new(IsOverlayedContext)

	p.IdentityOpContext = NewEmptyIdentityOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentityOpContext))

	return p
}

func (s *IsOverlayedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsOverlayedContext) EqualAt() antlr.TerminalNode {
	return s.GetToken(RhumbParserEqualAt, 0)
}

func (s *IsOverlayedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterIsOverlayed(s)
	}
}

func (s *IsOverlayedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitIsOverlayed(s)
	}
}

func (s *IsOverlayedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitIsOverlayed(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotLikeContext struct {
	*IdentityOpContext
}

func NewNotLikeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotLikeContext {
	var p = new(NotLikeContext)

	p.IdentityOpContext = NewEmptyIdentityOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentityOpContext))

	return p
}

func (s *NotLikeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotLikeContext) BangTilde() antlr.TerminalNode {
	return s.GetToken(RhumbParserBangTilde, 0)
}

func (s *NotLikeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNotLike(s)
	}
}

func (s *NotLikeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNotLike(s)
	}
}

func (s *NotLikeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNotLike(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsTopmostContext struct {
	*IdentityOpContext
}

func NewIsTopmostContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsTopmostContext {
	var p = new(IsTopmostContext)

	p.IdentityOpContext = NewEmptyIdentityOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentityOpContext))

	return p
}

func (s *IsTopmostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsTopmostContext) EqualCaret() antlr.TerminalNode {
	return s.GetToken(RhumbParserEqualCaret, 0)
}

func (s *IsTopmostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterIsTopmost(s)
	}
}

func (s *IsTopmostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitIsTopmost(s)
	}
}

func (s *IsTopmostContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitIsTopmost(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsEqualContext struct {
	*IdentityOpContext
}

func NewIsEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsEqualContext {
	var p = new(IsEqualContext)

	p.IdentityOpContext = NewEmptyIdentityOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentityOpContext))

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
	*IdentityOpContext
}

func NewNotEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotEqualContext {
	var p = new(NotEqualContext)

	p.IdentityOpContext = NewEmptyIdentityOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentityOpContext))

	return p
}

func (s *NotEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqualContext) BangEqual() antlr.TerminalNode {
	return s.GetToken(RhumbParserBangEqual, 0)
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

type NotInContext struct {
	*IdentityOpContext
}

func NewNotInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotInContext {
	var p = new(NotInContext)

	p.IdentityOpContext = NewEmptyIdentityOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentityOpContext))

	return p
}

func (s *NotInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotInContext) BangBSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserBangBSlash, 0)
}

func (s *NotInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNotIn(s)
	}
}

func (s *NotInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNotIn(s)
	}
}

func (s *NotInContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNotIn(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsInContext struct {
	*IdentityOpContext
}

func NewIsInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsInContext {
	var p = new(IsInContext)

	p.IdentityOpContext = NewEmptyIdentityOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentityOpContext))

	return p
}

func (s *IsInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsInContext) EqualBSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserEqualBSlash, 0)
}

func (s *IsInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterIsIn(s)
	}
}

func (s *IsInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitIsIn(s)
	}
}

func (s *IsInContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitIsIn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) IdentityOp() (localctx IIdentityOpContext) {
	this := p
	_ = this

	localctx = NewIdentityOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RhumbParserRULE_identityOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(239)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RhumbParserEqualEqual:
		localctx = NewIsEqualContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Match(RhumbParserEqualEqual)
		}

	case RhumbParserEqualTilde:
		localctx = NewIsLikeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.Match(RhumbParserEqualTilde)
		}

	case RhumbParserEqualBSlash:
		localctx = NewIsInContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(231)
			p.Match(RhumbParserEqualBSlash)
		}

	case RhumbParserEqualAt:
		localctx = NewIsOverlayedContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(232)
			p.Match(RhumbParserEqualAt)
		}

	case RhumbParserEqualCaret:
		localctx = NewIsTopmostContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(233)
			p.Match(RhumbParserEqualCaret)
		}

	case RhumbParserBangEqual:
		localctx = NewNotEqualContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(234)
			p.Match(RhumbParserBangEqual)
		}

	case RhumbParserBangTilde:
		localctx = NewNotLikeContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(235)
			p.Match(RhumbParserBangTilde)
		}

	case RhumbParserBangBSlash:
		localctx = NewNotInContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(236)
			p.Match(RhumbParserBangBSlash)
		}

	case RhumbParserBangAt:
		localctx = NewNotOverlayedContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(237)
			p.Match(RhumbParserBangAt)
		}

	case RhumbParserBangCaret:
		localctx = NewNotTopmostContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(238)
			p.Match(RhumbParserBangCaret)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConjunctiveOpContext is an interface to support dynamic dispatch.
type IConjunctiveOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConjunctiveOpContext differentiates from other interfaces.
	IsConjunctiveOpContext()
}

type ConjunctiveOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConjunctiveOpContext() *ConjunctiveOpContext {
	var p = new(ConjunctiveOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_conjunctiveOp
	return p
}

func (*ConjunctiveOpContext) IsConjunctiveOpContext() {}

func NewConjunctiveOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConjunctiveOpContext {
	var p = new(ConjunctiveOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_conjunctiveOp

	return p
}

func (s *ConjunctiveOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ConjunctiveOpContext) AmpAmp() antlr.TerminalNode {
	return s.GetToken(RhumbParserAmpAmp, 0)
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
	this := p
	_ = this

	localctx = NewConjunctiveOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RhumbParserRULE_conjunctiveOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(RhumbParserAmpAmp)
	}

	return localctx
}

// IDisjunctiveOpContext is an interface to support dynamic dispatch.
type IDisjunctiveOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDisjunctiveOpContext differentiates from other interfaces.
	IsDisjunctiveOpContext()
}

type DisjunctiveOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisjunctiveOpContext() *DisjunctiveOpContext {
	var p = new(DisjunctiveOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_disjunctiveOp
	return p
}

func (*DisjunctiveOpContext) IsDisjunctiveOpContext() {}

func NewDisjunctiveOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisjunctiveOpContext {
	var p = new(DisjunctiveOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_disjunctiveOp

	return p
}

func (s *DisjunctiveOpContext) GetParser() antlr.Parser { return s.parser }

func (s *DisjunctiveOpContext) PipePipe() antlr.TerminalNode {
	return s.GetToken(RhumbParserPipePipe, 0)
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
	this := p
	_ = this

	localctx = NewDisjunctiveOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RhumbParserRULE_disjunctiveOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(RhumbParserPipePipe)
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalOpContext() *ConditionalOpContext {
	var p = new(ConditionalOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_conditionalOp
	return p
}

func (*ConditionalOpContext) IsConditionalOpContext() {}

func NewConditionalOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalOpContext {
	var p = new(ConditionalOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_conditionalOp

	return p
}

func (s *ConditionalOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalOpContext) CopyFrom(ctx *ConditionalOpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConditionalOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OtherwiseContext struct {
	*ConditionalOpContext
}

func NewOtherwiseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OtherwiseContext {
	var p = new(OtherwiseContext)

	p.ConditionalOpContext = NewEmptyConditionalOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionalOpContext))

	return p
}

func (s *OtherwiseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OtherwiseContext) BangBang() antlr.TerminalNode {
	return s.GetToken(RhumbParserBangBang, 0)
}

func (s *OtherwiseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterOtherwise(s)
	}
}

func (s *OtherwiseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitOtherwise(s)
	}
}

func (s *OtherwiseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitOtherwise(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForeachContext struct {
	*ConditionalOpContext
}

func NewForeachContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForeachContext {
	var p = new(ForeachContext)

	p.ConditionalOpContext = NewEmptyConditionalOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionalOpContext))

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
	*ConditionalOpContext
}

func NewDefaultContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultContext {
	var p = new(DefaultContext)

	p.ConditionalOpContext = NewEmptyConditionalOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionalOpContext))

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
	*ConditionalOpContext
}

func NewElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseContext {
	var p = new(ElseContext)

	p.ConditionalOpContext = NewEmptyConditionalOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionalOpContext))

	return p
}

func (s *ElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseContext) BangGreater() antlr.TerminalNode {
	return s.GetToken(RhumbParserBangGreater, 0)
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

type ThenContext struct {
	*ConditionalOpContext
}

func NewThenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThenContext {
	var p = new(ThenContext)

	p.ConditionalOpContext = NewEmptyConditionalOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionalOpContext))

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
	*ConditionalOpContext
}

func NewWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileContext {
	var p = new(WhileContext)

	p.ConditionalOpContext = NewEmptyConditionalOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionalOpContext))

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

func (p *RhumbParser) ConditionalOp() (localctx IConditionalOpContext) {
	this := p
	_ = this

	localctx = NewConditionalOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RhumbParserRULE_conditionalOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(251)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RhumbParserBangBang:
		localctx = NewOtherwiseContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(245)
			p.Match(RhumbParserBangBang)
		}

	case RhumbParserQMarkQMark:
		localctx = NewDefaultContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(246)
			p.Match(RhumbParserQMarkQMark)
		}

	case RhumbParserLesserGreater:
		localctx = NewForeachContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(247)
			p.Match(RhumbParserLesserGreater)
		}

	case RhumbParserPipeGreater:
		localctx = NewWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(248)
			p.Match(RhumbParserPipeGreater)
		}

	case RhumbParserEqualGreater:
		localctx = NewThenContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(249)
			p.Match(RhumbParserEqualGreater)
		}

	case RhumbParserBangGreater:
		localctx = NewElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(250)
			p.Match(RhumbParserBangGreater)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveOpContext() *AdditiveOpContext {
	var p = new(AdditiveOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_additiveOp
	return p
}

func (*AdditiveOpContext) IsAdditiveOpContext() {}

func NewAdditiveOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveOpContext {
	var p = new(AdditiveOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_additiveOp

	return p
}

func (s *AdditiveOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveOpContext) CopyFrom(ctx *AdditiveOpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AdditiveOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SubtractionContext struct {
	*AdditiveOpContext
}

func NewSubtractionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubtractionContext {
	var p = new(SubtractionContext)

	p.AdditiveOpContext = NewEmptyAdditiveOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AdditiveOpContext))

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
	*AdditiveOpContext
}

func NewDeviationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeviationContext {
	var p = new(DeviationContext)

	p.AdditiveOpContext = NewEmptyAdditiveOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AdditiveOpContext))

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
	*AdditiveOpContext
}

func NewAdditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditionContext {
	var p = new(AdditionContext)

	p.AdditiveOpContext = NewEmptyAdditiveOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AdditiveOpContext))

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
	this := p
	_ = this

	localctx = NewAdditiveOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RhumbParserRULE_additiveOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(256)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RhumbParserPlusPlus:
		localctx = NewAdditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.Match(RhumbParserPlusPlus)
		}

	case RhumbParserPlusDash:
		localctx = NewDeviationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.Match(RhumbParserPlusDash)
		}

	case RhumbParserDashDash:
		localctx = NewSubtractionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(255)
			p.Match(RhumbParserDashDash)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeOpContext() *MultiplicativeOpContext {
	var p = new(MultiplicativeOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_multiplicativeOp
	return p
}

func (*MultiplicativeOpContext) IsMultiplicativeOpContext() {}

func NewMultiplicativeOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeOpContext {
	var p = new(MultiplicativeOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_multiplicativeOp

	return p
}

func (s *MultiplicativeOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeOpContext) CopyFrom(ctx *MultiplicativeOpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MultiplicativeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DivisionContext struct {
	*MultiplicativeOpContext
}

func NewDivisionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivisionContext {
	var p = new(DivisionContext)

	p.MultiplicativeOpContext = NewEmptyMultiplicativeOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MultiplicativeOpContext))

	return p
}

func (s *DivisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivisionContext) FSlashFSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserFSlashFSlash, 0)
}

func (s *DivisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterDivision(s)
	}
}

func (s *DivisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitDivision(s)
	}
}

func (s *DivisionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitDivision(s)

	default:
		return t.VisitChildren(s)
	}
}

type BindContext struct {
	*MultiplicativeOpContext
}

func NewBindContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BindContext {
	var p = new(BindContext)

	p.MultiplicativeOpContext = NewEmptyMultiplicativeOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MultiplicativeOpContext))

	return p
}

func (s *BindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindContext) TildeTilde() antlr.TerminalNode {
	return s.GetToken(RhumbParserTildeTilde, 0)
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

type IntegerDivisionContext struct {
	*MultiplicativeOpContext
}

func NewIntegerDivisionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerDivisionContext {
	var p = new(IntegerDivisionContext)

	p.MultiplicativeOpContext = NewEmptyMultiplicativeOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MultiplicativeOpContext))

	return p
}

func (s *IntegerDivisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerDivisionContext) PlusFSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserPlusFSlash, 0)
}

func (s *IntegerDivisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterIntegerDivision(s)
	}
}

func (s *IntegerDivisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitIntegerDivision(s)
	}
}

func (s *IntegerDivisionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitIntegerDivision(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiplicationContext struct {
	*MultiplicativeOpContext
}

func NewMultiplicationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicationContext {
	var p = new(MultiplicationContext)

	p.MultiplicativeOpContext = NewEmptyMultiplicativeOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MultiplicativeOpContext))

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

type ModuloContext struct {
	*MultiplicativeOpContext
}

func NewModuloContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModuloContext {
	var p = new(ModuloContext)

	p.MultiplicativeOpContext = NewEmptyMultiplicativeOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MultiplicativeOpContext))

	return p
}

func (s *ModuloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuloContext) DashFSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserDashFSlash, 0)
}

func (s *ModuloContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterModulo(s)
	}
}

func (s *ModuloContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitModulo(s)
	}
}

func (s *ModuloContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitModulo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) MultiplicativeOp() (localctx IMultiplicativeOpContext) {
	this := p
	_ = this

	localctx = NewMultiplicativeOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RhumbParserRULE_multiplicativeOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(263)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RhumbParserStarStar:
		localctx = NewMultiplicationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)
			p.Match(RhumbParserStarStar)
		}

	case RhumbParserFSlashFSlash:
		localctx = NewDivisionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(259)
			p.Match(RhumbParserFSlashFSlash)
		}

	case RhumbParserPlusFSlash:
		localctx = NewIntegerDivisionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(260)
			p.Match(RhumbParserPlusFSlash)
		}

	case RhumbParserDashFSlash:
		localctx = NewModuloContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(261)
			p.Match(RhumbParserDashFSlash)
		}

	case RhumbParserTildeTilde:
		localctx = NewBindContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(262)
			p.Match(RhumbParserTildeTilde)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExponentiationOpContext() *ExponentiationOpContext {
	var p = new(ExponentiationOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_exponentiationOp
	return p
}

func (*ExponentiationOpContext) IsExponentiationOpContext() {}

func NewExponentiationOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExponentiationOpContext {
	var p = new(ExponentiationOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_exponentiationOp

	return p
}

func (s *ExponentiationOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExponentiationOpContext) CopyFrom(ctx *ExponentiationOpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExponentiationOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentiationOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RootExtractionContext struct {
	*ExponentiationOpContext
}

func NewRootExtractionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RootExtractionContext {
	var p = new(RootExtractionContext)

	p.ExponentiationOpContext = NewEmptyExponentiationOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExponentiationOpContext))

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
	*ExponentiationOpContext
}

func NewScientificContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScientificContext {
	var p = new(ScientificContext)

	p.ExponentiationOpContext = NewEmptyExponentiationOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExponentiationOpContext))

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

type ExponentContext struct {
	*ExponentiationOpContext
}

func NewExponentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExponentContext {
	var p = new(ExponentContext)

	p.ExponentiationOpContext = NewEmptyExponentiationOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExponentiationOpContext))

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
	this := p
	_ = this

	localctx = NewExponentiationOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RhumbParserRULE_exponentiationOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(268)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RhumbParserCaretCaret:
		localctx = NewExponentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
			p.Match(RhumbParserCaretCaret)
		}

	case RhumbParserCaretFSlash:
		localctx = NewRootExtractionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.Match(RhumbParserCaretFSlash)
		}

	case RhumbParserStarCaret:
		localctx = NewScientificContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(267)
			p.Match(RhumbParserStarCaret)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOpContext() *AssignmentOpContext {
	var p = new(AssignmentOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_assignmentOp
	return p
}

func (*AssignmentOpContext) IsAssignmentOpContext() {}

func NewAssignmentOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOpContext {
	var p = new(AssignmentOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_assignmentOp

	return p
}

func (s *AssignmentOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentOpContext) CopyFrom(ctx *AssignmentOpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AssignmentOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ImmutablePairContext struct {
	*AssignmentOpContext
}

func NewImmutablePairContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImmutablePairContext {
	var p = new(ImmutablePairContext)

	p.AssignmentOpContext = NewEmptyAssignmentOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignmentOpContext))

	return p
}

func (s *ImmutablePairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImmutablePairContext) DotDot() antlr.TerminalNode {
	return s.GetToken(RhumbParserDotDot, 0)
}

func (s *ImmutablePairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterImmutablePair(s)
	}
}

func (s *ImmutablePairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitImmutablePair(s)
	}
}

func (s *ImmutablePairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitImmutablePair(s)

	default:
		return t.VisitChildren(s)
	}
}

type MutableLabelContext struct {
	*AssignmentOpContext
}

func NewMutableLabelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MutableLabelContext {
	var p = new(MutableLabelContext)

	p.AssignmentOpContext = NewEmptyAssignmentOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignmentOpContext))

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
	*AssignmentOpContext
}

func NewImmutableLabelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImmutableLabelContext {
	var p = new(ImmutableLabelContext)

	p.AssignmentOpContext = NewEmptyAssignmentOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignmentOpContext))

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

type MutablePairContext struct {
	*AssignmentOpContext
}

func NewMutablePairContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MutablePairContext {
	var p = new(MutablePairContext)

	p.AssignmentOpContext = NewEmptyAssignmentOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignmentOpContext))

	return p
}

func (s *MutablePairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutablePairContext) ColonColon() antlr.TerminalNode {
	return s.GetToken(RhumbParserColonColon, 0)
}

func (s *MutablePairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterMutablePair(s)
	}
}

func (s *MutablePairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitMutablePair(s)
	}
}

func (s *MutablePairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitMutablePair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) AssignmentOp() (localctx IAssignmentOpContext) {
	this := p
	_ = this

	localctx = NewAssignmentOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RhumbParserRULE_assignmentOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(274)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RhumbParserDotDot:
		localctx = NewImmutablePairContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(270)
			p.Match(RhumbParserDotDot)
		}

	case RhumbParserDotEqual:
		localctx = NewImmutableLabelContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(271)
			p.Match(RhumbParserDotEqual)
		}

	case RhumbParserColonColon:
		localctx = NewMutablePairContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(272)
			p.Match(RhumbParserColonColon)
		}

	case RhumbParserColonEqual:
		localctx = NewMutableLabelContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(273)
			p.Match(RhumbParserColonEqual)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixOpContext() *PrefixOpContext {
	var p = new(PrefixOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_prefixOp
	return p
}

func (*PrefixOpContext) IsPrefixOpContext() {}

func NewPrefixOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixOpContext {
	var p = new(PrefixOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_prefixOp

	return p
}

func (s *PrefixOpContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixOpContext) CopyFrom(ctx *PrefixOpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PrefixOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BaseCloneContext struct {
	*PrefixOpContext
}

func NewBaseCloneContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BaseCloneContext {
	var p = new(BaseCloneContext)

	p.PrefixOpContext = NewEmptyPrefixOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrefixOpContext))

	return p
}

func (s *BaseCloneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseCloneContext) Star() antlr.TerminalNode {
	return s.GetToken(RhumbParserStar, 0)
}

func (s *BaseCloneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterBaseClone(s)
	}
}

func (s *BaseCloneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitBaseClone(s)
	}
}

func (s *BaseCloneContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitBaseClone(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArgumentContext struct {
	*PrefixOpContext
}

func NewArgumentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgumentContext {
	var p = new(ArgumentContext)

	p.PrefixOpContext = NewEmptyPrefixOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrefixOpContext))

	return p
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) Caret() antlr.TerminalNode {
	return s.GetToken(RhumbParserCaret, 0)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

type OverlayContext struct {
	*PrefixOpContext
}

func NewOverlayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OverlayContext {
	var p = new(OverlayContext)

	p.PrefixOpContext = NewEmptyPrefixOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrefixOpContext))

	return p
}

func (s *OverlayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OverlayContext) At() antlr.TerminalNode {
	return s.GetToken(RhumbParserAt, 0)
}

func (s *OverlayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterOverlay(s)
	}
}

func (s *OverlayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitOverlay(s)
	}
}

func (s *OverlayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitOverlay(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumericalNegateContext struct {
	*PrefixOpContext
}

func NewNumericalNegateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumericalNegateContext {
	var p = new(NumericalNegateContext)

	p.PrefixOpContext = NewEmptyPrefixOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrefixOpContext))

	return p
}

func (s *NumericalNegateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericalNegateContext) Dash() antlr.TerminalNode {
	return s.GetToken(RhumbParserDash, 0)
}

func (s *NumericalNegateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNumericalNegate(s)
	}
}

func (s *NumericalNegateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNumericalNegate(s)
	}
}

func (s *NumericalNegateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNumericalNegate(s)

	default:
		return t.VisitChildren(s)
	}
}

type MutableDestructContext struct {
	*PrefixOpContext
}

func NewMutableDestructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MutableDestructContext {
	var p = new(MutableDestructContext)

	p.PrefixOpContext = NewEmptyPrefixOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrefixOpContext))

	return p
}

func (s *MutableDestructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutableDestructContext) Colon() antlr.TerminalNode {
	return s.GetToken(RhumbParserColon, 0)
}

func (s *MutableDestructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterMutableDestruct(s)
	}
}

func (s *MutableDestructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitMutableDestruct(s)
	}
}

func (s *MutableDestructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitMutableDestruct(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalPositContext struct {
	*PrefixOpContext
}

func NewLogicalPositContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalPositContext {
	var p = new(LogicalPositContext)

	p.PrefixOpContext = NewEmptyPrefixOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrefixOpContext))

	return p
}

func (s *LogicalPositContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalPositContext) Equal() antlr.TerminalNode {
	return s.GetToken(RhumbParserEqual, 0)
}

func (s *LogicalPositContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterLogicalPosit(s)
	}
}

func (s *LogicalPositContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitLogicalPosit(s)
	}
}

func (s *LogicalPositContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitLogicalPosit(s)

	default:
		return t.VisitChildren(s)
	}
}

type SlurpSpreadContext struct {
	*PrefixOpContext
}

func NewSlurpSpreadContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SlurpSpreadContext {
	var p = new(SlurpSpreadContext)

	p.PrefixOpContext = NewEmptyPrefixOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrefixOpContext))

	return p
}

func (s *SlurpSpreadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlurpSpreadContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(RhumbParserAmpersand, 0)
}

func (s *SlurpSpreadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterSlurpSpread(s)
	}
}

func (s *SlurpSpreadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitSlurpSpread(s)
	}
}

func (s *SlurpSpreadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitSlurpSpread(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumericalPositContext struct {
	*PrefixOpContext
}

func NewNumericalPositContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumericalPositContext {
	var p = new(NumericalPositContext)

	p.PrefixOpContext = NewEmptyPrefixOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrefixOpContext))

	return p
}

func (s *NumericalPositContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericalPositContext) Plus() antlr.TerminalNode {
	return s.GetToken(RhumbParserPlus, 0)
}

func (s *NumericalPositContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNumericalPosit(s)
	}
}

func (s *NumericalPositContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNumericalPosit(s)
	}
}

func (s *NumericalPositContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNumericalPosit(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssertContext struct {
	*PrefixOpContext
}

func NewAssertContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssertContext {
	var p = new(AssertContext)

	p.PrefixOpContext = NewEmptyPrefixOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrefixOpContext))

	return p
}

func (s *AssertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssertContext) Hash() antlr.TerminalNode {
	return s.GetToken(RhumbParserHash, 0)
}

func (s *AssertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterAssert(s)
	}
}

func (s *AssertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitAssert(s)
	}
}

func (s *AssertContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitAssert(s)

	default:
		return t.VisitChildren(s)
	}
}

type ImmutableDestructContext struct {
	*PrefixOpContext
}

func NewImmutableDestructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImmutableDestructContext {
	var p = new(ImmutableDestructContext)

	p.PrefixOpContext = NewEmptyPrefixOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrefixOpContext))

	return p
}

func (s *ImmutableDestructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImmutableDestructContext) Dot() antlr.TerminalNode {
	return s.GetToken(RhumbParserDot, 0)
}

func (s *ImmutableDestructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterImmutableDestruct(s)
	}
}

func (s *ImmutableDestructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitImmutableDestruct(s)
	}
}

func (s *ImmutableDestructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitImmutableDestruct(s)

	default:
		return t.VisitChildren(s)
	}
}

type OuterScopeContext struct {
	*PrefixOpContext
}

func NewOuterScopeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OuterScopeContext {
	var p = new(OuterScopeContext)

	p.PrefixOpContext = NewEmptyPrefixOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrefixOpContext))

	return p
}

func (s *OuterScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OuterScopeContext) BSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserBSlash, 0)
}

func (s *OuterScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterOuterScope(s)
	}
}

func (s *OuterScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitOuterScope(s)
	}
}

func (s *OuterScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitOuterScope(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExistentialPositContext struct {
	*PrefixOpContext
}

func NewExistentialPositContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExistentialPositContext {
	var p = new(ExistentialPositContext)

	p.PrefixOpContext = NewEmptyPrefixOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrefixOpContext))

	return p
}

func (s *ExistentialPositContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExistentialPositContext) QMark() antlr.TerminalNode {
	return s.GetToken(RhumbParserQMark, 0)
}

func (s *ExistentialPositContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterExistentialPosit(s)
	}
}

func (s *ExistentialPositContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitExistentialPosit(s)
	}
}

func (s *ExistentialPositContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitExistentialPosit(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalNegateContext struct {
	*PrefixOpContext
}

func NewLogicalNegateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalNegateContext {
	var p = new(LogicalNegateContext)

	p.PrefixOpContext = NewEmptyPrefixOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrefixOpContext))

	return p
}

func (s *LogicalNegateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalNegateContext) Bang() antlr.TerminalNode {
	return s.GetToken(RhumbParserBang, 0)
}

func (s *LogicalNegateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterLogicalNegate(s)
	}
}

func (s *LogicalNegateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitLogicalNegate(s)
	}
}

func (s *LogicalNegateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitLogicalNegate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) PrefixOp() (localctx IPrefixOpContext) {
	this := p
	_ = this

	localctx = NewPrefixOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RhumbParserRULE_prefixOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(289)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RhumbParserDash:
		localctx = NewNumericalNegateContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)
			p.Match(RhumbParserDash)
		}

	case RhumbParserBSlash:
		localctx = NewOuterScopeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.Match(RhumbParserBSlash)
		}

	case RhumbParserBang:
		localctx = NewLogicalNegateContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(278)
			p.Match(RhumbParserBang)
		}

	case RhumbParserHash:
		localctx = NewAssertContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(279)
			p.Match(RhumbParserHash)
		}

	case RhumbParserCaret:
		localctx = NewArgumentContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(280)
			p.Match(RhumbParserCaret)
		}

	case RhumbParserAmpersand:
		localctx = NewSlurpSpreadContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(281)
			p.Match(RhumbParserAmpersand)
		}

	case RhumbParserStar:
		localctx = NewBaseCloneContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(282)
			p.Match(RhumbParserStar)
		}

	case RhumbParserPlus:
		localctx = NewNumericalPositContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(283)
			p.Match(RhumbParserPlus)
		}

	case RhumbParserEqual:
		localctx = NewLogicalPositContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(284)
			p.Match(RhumbParserEqual)
		}

	case RhumbParserAt:
		localctx = NewOverlayContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(285)
			p.Match(RhumbParserAt)
		}

	case RhumbParserQMark:
		localctx = NewExistentialPositContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(286)
			p.Match(RhumbParserQMark)
		}

	case RhumbParserDot:
		localctx = NewImmutableDestructContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(287)
			p.Match(RhumbParserDot)
		}

	case RhumbParserColon:
		localctx = NewMutableDestructContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(288)
			p.Match(RhumbParserColon)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChainOpContext() *ChainOpContext {
	var p = new(ChainOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_chainOp
	return p
}

func (*ChainOpContext) IsChainOpContext() {}

func NewChainOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChainOpContext {
	var p = new(ChainOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_chainOp

	return p
}

func (s *ChainOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ChainOpContext) CopyFrom(ctx *ChainOpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ChainOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NestedOverlayContext struct {
	*ChainOpContext
}

func NewNestedOverlayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedOverlayContext {
	var p = new(NestedOverlayContext)

	p.ChainOpContext = NewEmptyChainOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ChainOpContext))

	return p
}

func (s *NestedOverlayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedOverlayContext) At() antlr.TerminalNode {
	return s.GetToken(RhumbParserAt, 0)
}

func (s *NestedOverlayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNestedOverlay(s)
	}
}

func (s *NestedOverlayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNestedOverlay(s)
	}
}

func (s *NestedOverlayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNestedOverlay(s)

	default:
		return t.VisitChildren(s)
	}
}

type NestedLabelContext struct {
	*ChainOpContext
}

func NewNestedLabelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedLabelContext {
	var p = new(NestedLabelContext)

	p.ChainOpContext = NewEmptyChainOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ChainOpContext))

	return p
}

func (s *NestedLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedLabelContext) BSlash() antlr.TerminalNode {
	return s.GetToken(RhumbParserBSlash, 0)
}

func (s *NestedLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterNestedLabel(s)
	}
}

func (s *NestedLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitNestedLabel(s)
	}
}

func (s *NestedLabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitNestedLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

type RangeContext struct {
	*ChainOpContext
}

func NewRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangeContext {
	var p = new(RangeContext)

	p.ChainOpContext = NewEmptyChainOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ChainOpContext))

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

func (p *RhumbParser) ChainOp() (localctx IChainOpContext) {
	this := p
	_ = this

	localctx = NewChainOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RhumbParserRULE_chainOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(294)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RhumbParserBSlash:
		localctx = NewNestedLabelContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(291)
			p.Match(RhumbParserBSlash)
		}

	case RhumbParserAt:
		localctx = NewNestedOverlayContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)
			p.Match(RhumbParserAt)
		}

	case RhumbParserPipe:
		localctx = NewRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(293)
			p.Match(RhumbParserPipe)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_string
	return p
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) OpenInterpString() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenInterpString, 0)
}

func (s *StringContext) CloseInterpString() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseInterpString, 0)
}

func (s *StringContext) AllRawText() []antlr.TerminalNode {
	return s.GetTokens(RhumbParserRawText)
}

func (s *StringContext) RawText(i int) antlr.TerminalNode {
	return s.GetToken(RhumbParserRawText, i)
}

func (s *StringContext) AllInterpolation() []IInterpolationContext {
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

func (s *StringContext) Interpolation(i int) IInterpolationContext {
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

func (s *StringContext) RawString() antlr.TerminalNode {
	return s.GetToken(RhumbParserRawString, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RhumbParserListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RhumbParserVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RhumbParser) String_() (localctx IStringContext) {
	this := p
	_ = this

	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RhumbParserRULE_string)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(306)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RhumbParserOpenInterpString:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(296)
			p.Match(RhumbParserOpenInterpString)
		}
		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-79)) & ^0x3f) == 0 && ((int64(1)<<(_la-79))&39) != 0 {
			p.SetState(299)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case RhumbParserRawText:
				{
					p.SetState(297)
					p.Match(RhumbParserRawText)
				}

			case RhumbParserEnterRoutineInterp, RhumbParserEnterSelectorInterp, RhumbParserInterpLabel:
				{
					p.SetState(298)
					p.Interpolation()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(303)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(304)
			p.Match(RhumbParserCloseInterpString)
		}

	case RhumbParserRawString:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)
			p.Match(RhumbParserRawString)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterpolationContext() *InterpolationContext {
	var p = new(InterpolationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_interpolation
	return p
}

func (*InterpolationContext) IsInterpolationContext() {}

func NewInterpolationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterpolationContext {
	var p = new(InterpolationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_interpolation

	return p
}

func (s *InterpolationContext) GetParser() antlr.Parser { return s.parser }

func (s *InterpolationContext) CopyFrom(ctx *InterpolationContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *InterpolationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterpolationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LabelInterpContext struct {
	*InterpolationContext
}

func NewLabelInterpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LabelInterpContext {
	var p = new(LabelInterpContext)

	p.InterpolationContext = NewEmptyInterpolationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InterpolationContext))

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

type RoutineInterpContext struct {
	*InterpolationContext
}

func NewRoutineInterpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoutineInterpContext {
	var p = new(RoutineInterpContext)

	p.InterpolationContext = NewEmptyInterpolationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InterpolationContext))

	return p
}

func (s *RoutineInterpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoutineInterpContext) EnterRoutineInterp() antlr.TerminalNode {
	return s.GetToken(RhumbParserEnterRoutineInterp, 0)
}

func (s *RoutineInterpContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
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
	*InterpolationContext
}

func NewSelectorInterpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectorInterpContext {
	var p = new(SelectorInterpContext)

	p.InterpolationContext = NewEmptyInterpolationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InterpolationContext))

	return p
}

func (s *SelectorInterpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorInterpContext) EnterSelectorInterp() antlr.TerminalNode {
	return s.GetToken(RhumbParserEnterSelectorInterp, 0)
}

func (s *SelectorInterpContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
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

func (p *RhumbParser) Interpolation() (localctx IInterpolationContext) {
	this := p
	_ = this

	localctx = NewInterpolationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RhumbParserRULE_interpolation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(317)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RhumbParserInterpLabel:
		localctx = NewLabelInterpContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(308)
			p.Match(RhumbParserInterpLabel)
		}

	case RhumbParserEnterRoutineInterp:
		localctx = NewRoutineInterpContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(309)
			p.Match(RhumbParserEnterRoutineInterp)
		}
		{
			p.SetState(310)
			p.Sequence()
		}
		{
			p.SetState(311)
			p.Match(RhumbParserCloseParen)
		}

	case RhumbParserEnterSelectorInterp:
		localctx = NewSelectorInterpContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(313)
			p.Match(RhumbParserEnterSelectorInterp)
		}
		{
			p.SetState(314)
			p.Sequence()
		}
		{
			p.SetState(315)
			p.Match(RhumbParserCloseCurly)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RhumbParserRULE_reference
	return p
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RhumbParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) CopyFrom(ctx *ReferenceContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionRefContext struct {
	*ReferenceContext
}

func NewFunctionRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionRefContext {
	var p = new(FunctionRefContext)

	p.ReferenceContext = NewEmptyReferenceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ReferenceContext))

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

func (s *FunctionRefContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
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
	*ReferenceContext
}

func NewNamedRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NamedRefContext {
	var p = new(NamedRefContext)

	p.ReferenceContext = NewEmptyReferenceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ReferenceContext))

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
	*ReferenceContext
}

func NewJunctionRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JunctionRefContext {
	var p = new(JunctionRefContext)

	p.ReferenceContext = NewEmptyReferenceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ReferenceContext))

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

func (s *JunctionRefContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
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
	*ReferenceContext
}

func NewComputedRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComputedRefContext {
	var p = new(ComputedRefContext)

	p.ReferenceContext = NewEmptyReferenceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ReferenceContext))

	return p
}

func (s *ComputedRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComputedRefContext) OpenAnglet() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenAnglet, 0)
}

func (s *ComputedRefContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(RhumbParserOpenBracket, 0)
}

func (s *ComputedRefContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *ComputedRefContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseBracket, 0)
}

func (s *ComputedRefContext) CloseAnglet() antlr.TerminalNode {
	return s.GetToken(RhumbParserCloseAnglet, 0)
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
	this := p
	_ = this

	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RhumbParserRULE_reference)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNamedRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(319)
			p.Match(RhumbParserOpenAnglet)
		}
		{
			p.SetState(320)
			p.Match(RhumbParserLabel)
		}
		{
			p.SetState(321)
			p.Match(RhumbParserCloseAnglet)
		}

	case 2:
		localctx = NewFunctionRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.Match(RhumbParserOpenAnglet)
		}
		{
			p.SetState(323)
			p.Match(RhumbParserOpenParen)
		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserSemicolon || _la == RhumbParserNL {
			{
				p.SetState(324)
				p.Terminator()
			}

			p.SetState(329)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(330)
			p.Sequence()
		}
		{
			p.SetState(331)
			p.Match(RhumbParserCloseParen)
		}
		{
			p.SetState(332)
			p.Match(RhumbParserCloseAnglet)
		}

	case 3:
		localctx = NewComputedRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(334)
			p.Match(RhumbParserOpenAnglet)
		}
		{
			p.SetState(335)
			p.Match(RhumbParserOpenBracket)
		}
		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserSemicolon || _la == RhumbParserNL {
			{
				p.SetState(336)
				p.Terminator()
			}

			p.SetState(341)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(342)
			p.Sequence()
		}
		{
			p.SetState(343)
			p.Match(RhumbParserCloseBracket)
		}
		{
			p.SetState(344)
			p.Match(RhumbParserCloseAnglet)
		}

	case 4:
		localctx = NewJunctionRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(346)
			p.Match(RhumbParserOpenAnglet)
		}
		{
			p.SetState(347)
			p.Match(RhumbParserOpenCurly)
		}
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RhumbParserSemicolon || _la == RhumbParserNL {
			{
				p.SetState(348)
				p.Terminator()
			}

			p.SetState(353)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(354)
			p.Sequence()
		}
		{
			p.SetState(355)
			p.Match(RhumbParserCloseCurly)
		}
		{
			p.SetState(356)
			p.Match(RhumbParserCloseAnglet)
		}

	}

	return localctx
}

func (p *RhumbParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
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
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 4)

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
