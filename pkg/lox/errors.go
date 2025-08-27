package lox

// Error is a Lox runtime error.
type Error struct {
	Msg  string
	Line int
}

// Error implements the [error] interface.
func (err *Error) Error() string {
	return err.Msg
}
