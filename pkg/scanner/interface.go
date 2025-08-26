// Package scanner implements a scanner for the Lox source code into lexical
// tokens.
package scanner

import "github.com/kevhlee/glox/pkg/token"

// Scan converts Lox source code into a lexical tokens.
func ScanSource(source string) []*token.Token {
	var s scanner

	s.source = []rune(source)
	s.line = 1

	return s.scan()
}
