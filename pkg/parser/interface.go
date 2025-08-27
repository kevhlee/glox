// Package parser implements a parser for the Lox source code into an AST.
package parser

import (
	"github.com/kevhlee/glox/pkg/ast"
	"github.com/kevhlee/glox/pkg/scanner"
	"github.com/kevhlee/glox/pkg/token"
)

// ParseSource converts Lox source code into an AST.
func ParseSource(source string) ([]ast.Stmt, error) {
	var p parser

	for _, tok := range scanner.ScanSource(source) {
		if tok.Type == token.ERROR {
			p.errors = append(p.errors, &Error{tok.Lexeme, tok})
		} else {
			p.tokens = append(p.tokens, tok)
		}
	}

	return p.parse()
}
