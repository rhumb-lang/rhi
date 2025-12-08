package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type RhumbErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []string
}

func NewRhumbErrorListener() *RhumbErrorListener {
	return &RhumbErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		Errors:               make([]string, 0),
	}
}

func (l *RhumbErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, fmt.Sprintf("Line %d:%d - %s", line, column, msg))
}
