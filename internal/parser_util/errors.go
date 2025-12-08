package parser_util

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

// RhumbErrorListener collects syntax errors during parsing.
type RhumbErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []string
}

// NewRhumbErrorListener creates a new instance of RhumbErrorListener.
func NewRhumbErrorListener() *RhumbErrorListener {
	return &RhumbErrorListener{Errors: make([]string, 0)}
}

// SyntaxError captures syntax errors.
func (l *RhumbErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, fmt.Sprintf("line %d:%d %s", line, column, msg))
}
