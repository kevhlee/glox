// Package lox implements an interpreter for the Lox programming language.
package lox

import (
	"fmt"
	"os"

	"github.com/kevhlee/glox/pkg/parser"
	"github.com/kevhlee/glox/pkg/token"
)

const (
	// ExitOK is the exit status code returned when the Lox interpreter
	// successfully executed.
	ExitOK = 0

	// ExitCompileErr is the exit status code returned when the Lox interpreter
	// failed to execute due to one or more compile errors in the source code.
	ExitCompileErr = 65

	// ExitRuntimeErr is the exit status code returned when the Lox interpreter
	// failed to execute due to a runtime error.
	ExitRuntimeErr = 70
)

// RunSource executes Lox source code.
//
// The function returns a exit status code that can be passed to [os.Exit] as an
// argument.
func RunSource(globals *Environment, source string) int {
	parsed, err := parser.ParseSource(source)

	if err != nil {
		for _, err := range err.(parser.ErrorList) {
			switch err.Token.Type {
			case token.EOF:
				reportCompileError(err.Token.Line, " at end", err.Error())
			case token.ERROR:
				reportCompileError(err.Token.Line, "", err.Error())
			default:
				reportCompileError(err.Token.Line, fmt.Sprintf(" at '%s'", err.Token.Lexeme), err.Error())
			}
		}

		return ExitCompileErr
	}

	var ip interpreter

	if err := ip.Interpret(globals, parsed); err != nil {
		if err, ok := err.(*Error); ok {
			reportRuntimeError(err.Line, err.Error())
		}
		return ExitRuntimeErr
	}

	return ExitOK
}

func reportCompileError(line int, where, msg string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s.\n", line, where, msg)
}

func reportRuntimeError(line int, msg string) {
	fmt.Fprintf(os.Stderr, "%s.\n[line %d]\n", msg, line)
}
