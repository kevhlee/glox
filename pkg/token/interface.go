// Package token defines lexical types and tokens for the Lox programming
// language.
package token

// IsDigit checks if a given character is a digit in a Lox number literal.
func IsDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

// IsIdentPart checks if a character may be part of a Lox identifier.
//
// Note that this function does not check if the character is a valid starting
// character (use [token.IsIdentStart]) instead for this case).
func IsIdentPart(ch rune) bool {
	return IsIdentStart(ch) || IsDigit(ch)
}

// IsIdentStart checks if a character is a valid starting character for a Lox
// identifier.
func IsIdentStart(ch rune) bool {
	return ch == '_' || (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}

// ParseKeyword parses a given string into a keyword lexical type, if possible.
//
// The function returns a boolean value to indicate if the string represents a
// Lox keyword.
func ParseKeyword(s string) (Type, bool) {
	t, ok := keywords[s]
	return t, ok
}
