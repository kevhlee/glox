package parser

import (
	"fmt"

	"github.com/kevhlee/glox/pkg/token"
)

// Error is a syntax error detected by the parser.
type Error struct {
	Msg   string
	Token *token.Token
}

// Error implements the [error] interface.
func (err *Error) Error() string {
	return err.Msg
}

// ErrorList is a list of parser errors.
type ErrorList []*Error

// Err returns an error representative of the list if it is not empty.
func (l ErrorList) Err() error {
	if len(l) > 0 {
		return l
	}
	return nil
}

// Error implements the [error] interface.
func (l ErrorList) Error() string {
	switch len(l) {
	case 0:
		return "no errors"
	case 1:
		return l[0].Error()
	case 2:
		return fmt.Sprintf("%s (and 1 more error)", l[0].Error())
	default:
		return fmt.Sprintf("%s (and %d more errors)", l[0].Error(), len(l)-1)
	}
}
