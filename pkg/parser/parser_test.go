package parser_test

import (
	"strings"
	"testing"

	"github.com/kevhlee/glox/pkg/ast"
	"github.com/kevhlee/glox/pkg/parser"
)

// TestParseExpr checks that the parser can correctly parse expressions.
func TestParseExpr(t *testing.T) {
	res, err := parser.ParseSource("(5 - (3 - 1)) + -1;")
	if err != nil {
		t.Fatal(err)
	}

	expected := `EXPRESSION
└── BINARY(+)
    ├── GROUP
    │   └── BINARY(-)
    │       ├── NUMBER(5)
    │       └── GROUP
    │           └── BINARY(-)
    │               ├── NUMBER(3)
    │               └── NUMBER(1)
    └── UNARY(-)
        └── NUMBER(1)
`

	var sb strings.Builder

	for _, stmt := range res {
		sb.WriteString(ast.Print(stmt))
	}

	if actual := sb.String(); actual != expected {
		t.Errorf("\nExpected:\n%s\nActual:\n%s\n", expected, actual)
	}
}
