package scanner_test

import (
	"testing"

	"github.com/kevhlee/glox/pkg/scanner"
	"github.com/kevhlee/glox/pkg/token"
)

// TestScanIdentifiers checks to make sure the scanner handles identifier
// literals.
func TestScanIdentifiers(t *testing.T) {
	source := `andy formless fo _ _123 _abc ab123
abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_
`

	testScan(t, source, []token.Token{
		{Type: token.IDENTIFIER, Lexeme: "andy", Line: 1},
		{Type: token.IDENTIFIER, Lexeme: "formless", Line: 1},
		{Type: token.IDENTIFIER, Lexeme: "fo", Line: 1},
		{Type: token.IDENTIFIER, Lexeme: "_", Line: 1},
		{Type: token.IDENTIFIER, Lexeme: "_123", Line: 1},
		{Type: token.IDENTIFIER, Lexeme: "_abc", Line: 1},
		{Type: token.IDENTIFIER, Lexeme: "ab123", Line: 1},
		{Type: token.IDENTIFIER, Lexeme: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_", Line: 2},
		{Type: token.EOF, Lexeme: "", Line: 3},
	})
}

// TestScanKeywords checks to make sure the scanner handles keywords.
func TestScanKeywords(t *testing.T) {
	source := "and class else false for fun if nil or return super this true var while"

	testScan(t, source, []token.Token{
		{Type: token.AND, Lexeme: "and", Line: 1},
		{Type: token.CLASS, Lexeme: "class", Line: 1},
		{Type: token.ELSE, Lexeme: "else", Line: 1},
		{Type: token.FALSE, Lexeme: "false", Line: 1},
		{Type: token.FOR, Lexeme: "for", Line: 1},
		{Type: token.FUN, Lexeme: "fun", Line: 1},
		{Type: token.IF, Lexeme: "if", Line: 1},
		{Type: token.NIL, Lexeme: "nil", Line: 1},
		{Type: token.OR, Lexeme: "or", Line: 1},
		{Type: token.RETURN, Lexeme: "return", Line: 1},
		{Type: token.SUPER, Lexeme: "super", Line: 1},
		{Type: token.THIS, Lexeme: "this", Line: 1},
		{Type: token.TRUE, Lexeme: "true", Line: 1},
		{Type: token.VAR, Lexeme: "var", Line: 1},
		{Type: token.WHILE, Lexeme: "while", Line: 1},
		{Type: token.EOF, Lexeme: "", Line: 1},
	})
}

// TestScanNumbers checks to make sure the scanner handles number literals.
func TestScanNumbers(t *testing.T) {
	source := `123
123.456
.456
123.
`

	testScan(t, source, []token.Token{
		{Type: token.NUMBER, Lexeme: "123", Line: 1},
		{Type: token.NUMBER, Lexeme: "123.456", Line: 2},
		{Type: token.DOT, Lexeme: ".", Line: 3},
		{Type: token.NUMBER, Lexeme: "456", Line: 3},
		{Type: token.NUMBER, Lexeme: "123", Line: 4},
		{Type: token.DOT, Lexeme: ".", Line: 4},
		{Type: token.EOF, Lexeme: "", Line: 5},
	})
}

// TestScanPunctuators checks to make sure the scanner handles punctuation
// characters.
func TestScanPunctuators(t *testing.T) {
	source := `(){};,+-*!===<=>=!=<>/.`

	testScan(t, source, []token.Token{
		{Type: token.LEFT_PAREN, Lexeme: "(", Line: 1},
		{Type: token.RIGHT_PAREN, Lexeme: ")", Line: 1},
		{Type: token.LEFT_BRACE, Lexeme: "{", Line: 1},
		{Type: token.RIGHT_BRACE, Lexeme: "}", Line: 1},
		{Type: token.SEMICOLON, Lexeme: ";", Line: 1},
		{Type: token.COMMA, Lexeme: ",", Line: 1},
		{Type: token.PLUS, Lexeme: "+", Line: 1},
		{Type: token.MINUS, Lexeme: "-", Line: 1},
		{Type: token.STAR, Lexeme: "*", Line: 1},
		{Type: token.BANG_EQUAL, Lexeme: "!=", Line: 1},
		{Type: token.EQUAL_EQUAL, Lexeme: "==", Line: 1},
		{Type: token.LESS_EQUAL, Lexeme: "<=", Line: 1},
		{Type: token.GREATER_EQUAL, Lexeme: ">=", Line: 1},
		{Type: token.BANG_EQUAL, Lexeme: "!=", Line: 1},
		{Type: token.LESS, Lexeme: "<", Line: 1},
		{Type: token.GREATER, Lexeme: ">", Line: 1},
		{Type: token.SLASH, Lexeme: "/", Line: 1},
		{Type: token.DOT, Lexeme: ".", Line: 1},
		{Type: token.EOF, Lexeme: "", Line: 1},
	})
}

// TestScanStrings checks to make sure the scanner handles string literals.
func TestScanStrings(t *testing.T) {
	source := `""
"string"
`

	testScan(t, source, []token.Token{
		{Type: token.STRING, Lexeme: `""`, Line: 1},
		{Type: token.STRING, Lexeme: `"string"`, Line: 2},
		{Type: token.EOF, Lexeme: "", Line: 3},
	})
}

// TestScanWhitespace checks to make sure the scanner handle whitespace.
func TestScanWhitespace(t *testing.T) {
	source := `space    tabs				newlines




end
`

	testScan(t, source, []token.Token{
		{Type: token.IDENTIFIER, Lexeme: "space", Line: 1},
		{Type: token.IDENTIFIER, Lexeme: "tabs", Line: 1},
		{Type: token.IDENTIFIER, Lexeme: "newlines", Line: 1},
		{Type: token.IDENTIFIER, Lexeme: "end", Line: 6},
		{Type: token.EOF, Lexeme: "", Line: 7},
	})
}

func testScan(t *testing.T, source string, expected []token.Token) {
	actual := scanner.ScanSource(source)

	if len(actual) != len(expected) {
		t.Fatalf("Expected %d tokens, got %d instead", len(expected), len(actual))
	}

	for i := range expected {
		if actual[i].Type != expected[i].Type {
			t.Errorf("Expected type '%s', got '%s' instead", expected[i].Type, actual[i].Type)
		}
		if actual[i].Lexeme != expected[i].Lexeme {
			t.Errorf("Expected lexeme '%s', got '%s' instead", expected[i].Lexeme, actual[i].Lexeme)
		}
		if actual[i].Line != expected[i].Line {
			t.Errorf("Expected line '%d', got '%d' instead", expected[i].Line, actual[i].Line)
		}
	}
}
