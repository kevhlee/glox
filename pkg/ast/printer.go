package ast

import (
	"fmt"
	"strings"

	"github.com/kevhlee/glox/internal/stack"
	"github.com/kevhlee/glox/pkg/token"
)

// Contains the internal state and logic of the AST pretty-printer.
type printer struct {
	strings.Builder
	paddings  stack.Stack[string]
	lastChild bool
}

func (p *printer) Print(node Node) string {
	p.lastChild = true
	Walk(p, node)
	return p.String()
}

// Visit implements the [Visitor] interface.
func (p *printer) Visit(node Node) bool {
	originalLastChild := p.lastChild

	defer func() {
		p.lastChild = originalLastChild
	}()

	if !p.paddings.Empty() {
		for i := range p.paddings.Len() - 1 {
			p.WriteString(p.paddings.At(i))
		}

		if p.lastChild {
			p.WriteString("└── ")
		} else {
			p.WriteString("├── ")
		}
	}

	switch node := node.(type) {
	// Stmt

	case *BlockStmt:
		p.WriteString("BLOCK\n")

	case *ExpressionStmt:
		p.WriteString("EXPRESSION\n")

	case *PrintStmt:
		p.WriteString("PRINT\n")

	case *VarStmt:
		p.WriteString("VAR(")
		p.WriteString(node.Name.Lexeme)
		p.WriteString(")\n")

	// Expr

	case *AssignExpr:
		p.WriteString("ASSIGN(")
		p.WriteString(node.Name.Lexeme)
		p.WriteString(")\n")

	case *BinaryExpr:
		p.WriteString("BINARY(")
		p.WriteString(node.Operator.Lexeme)
		p.WriteString(")\n")

	case *GroupingExpr:
		p.WriteString("GROUP\n")

	case *LiteralExpr:
		switch value := node.Value; value.Type {
		case token.NIL:
			p.WriteString("NIL\n")
		case token.TRUE:
			p.WriteString("TRUE\n")
		case token.FALSE:
			p.WriteString("FALSE\n")
		default:
			p.WriteString(value.Type.String())
			p.WriteString("(")
			p.WriteString(value.Lexeme)
			p.WriteString(")\n")
		}

	case *UnaryExpr:
		p.WriteString("UNARY(")
		p.WriteString(node.Operator.Lexeme)
		p.WriteString(")\n")

	case *VariableExpr:
		p.WriteString("VARIABLE(")
		p.WriteString(node.Name.Lexeme)
		p.WriteString(")\n")

	default:
		panic(fmt.Errorf("Unexpected node type %T", node))
	}

	if children := Children(node); len(children) > 0 {
		for i, child := range children {
			if p.lastChild = i >= len(children)-1; p.lastChild {
				p.paddings.Push("    ")
			} else {
				p.paddings.Push("│   ")
			}

			Walk(p, child)

			p.paddings.Pop()
		}
	}

	return false
}
