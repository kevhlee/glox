package token

import "fmt"

// Type is a lexical type.
type Type int

// String implements [fmt.Stringer] interface.
func (t Type) String() string {
	return names[t]
}

const (
	__START__ Type = iota

	// Special

	EOF   // End-of-file
	ERROR // Error

	// Single character

	LEFT_PAREN  // (
	RIGHT_PAREN // )
	LEFT_BRACE  // {
	RIGHT_BRACE // }
	COMMA       // ,
	DOT         // .
	PLUS        // +
	MINUS       // -
	SEMICOLON   // ;
	SLASH       // /
	STAR        // *

	// Single or double characters

	BANG          // !
	BANG_EQUAL    // !=
	EQUAL         // =
	EQUAL_EQUAL   // ==
	GREATER       // >
	GREATER_EQUAL // >=
	LESS          // <
	LESS_EQUAL    // <=

	// Literals

	IDENTIFIER // IDENTIFIER
	STRING     // STRING
	NUMBER     // NUMBER

	// Keywords

	AND    // and
	CLASS  // class
	ELSE   // else
	FALSE  // false
	FOR    // for
	FUN    // fun
	IF     // if
	NIL    // nil
	OR     // or
	PRINT  // print
	RETURN // return
	SUPER  // super
	THIS   // this
	TRUE   // true
	VAR    // var
	WHILE  // while

	__END__
)

var (
	names = map[Type]string{
		EOF:   "EOF",
		ERROR: "ERROR",

		LEFT_PAREN:  "(",
		RIGHT_PAREN: ")",
		LEFT_BRACE:  "{",
		RIGHT_BRACE: "}",
		COMMA:       ",",
		DOT:         ".",
		PLUS:        "+",
		MINUS:       "-",
		SEMICOLON:   ";",
		SLASH:       "/",
		STAR:        "*",

		BANG:          "!",
		BANG_EQUAL:    "!=",
		EQUAL:         "=",
		EQUAL_EQUAL:   "==",
		GREATER:       ">",
		GREATER_EQUAL: ">=",
		LESS:          "<",
		LESS_EQUAL:    "<=",

		IDENTIFIER: "IDENTIFIER",
		STRING:     "STRING",
		NUMBER:     "NUMBER",

		AND:    "and",
		CLASS:  "class",
		ELSE:   "else",
		FALSE:  "false",
		FOR:    "for",
		FUN:    "fun",
		IF:     "if",
		NIL:    "nil",
		OR:     "or",
		PRINT:  "print",
		RETURN: "return",
		SUPER:  "super",
		THIS:   "this",
		TRUE:   "true",
		VAR:    "var",
		WHILE:  "while",
	}

	keywords = map[string]Type{
		"and":    AND,
		"class":  CLASS,
		"else":   ELSE,
		"false":  FALSE,
		"for":    FOR,
		"fun":    FUN,
		"if":     IF,
		"nil":    NIL,
		"or":     OR,
		"print":  PRINT,
		"return": RETURN,
		"super":  SUPER,
		"this":   THIS,
		"true":   TRUE,
		"var":    VAR,
		"while":  WHILE,
	}
)

// Token is a lexical token.
type Token struct {
	Type
	Lexeme string
	Line   int
}

// String implements the [fmt.Stringer] interface.
func (t *Token) String() string {
	return fmt.Sprintf("<%s %s %d>", t.Type, t.Lexeme, t.Line)
}
