package parser

import (
	"slices"

	"github.com/kevhlee/glox/pkg/ast"
	"github.com/kevhlee/glox/pkg/token"
)

// Contains the internal state and logic of the parser.
type parser struct {
	tokens  []*token.Token
	errors  ErrorList
	current int
}

func (p *parser) parse() ([]ast.Stmt, error) {
	var result []ast.Stmt

	for p.isParsing() {
		if decl := p.declaration(); decl != nil {
			result = append(result, decl)
		}
	}

	return result, p.errors.Err()
}

func (p *parser) peek() *token.Token {
	return p.tokens[p.current]
}

func (p *parser) previous() *token.Token {
	return p.tokens[p.current-1]
}

func (p *parser) isParsing() bool {
	return p.peek().Type != token.EOF
}

func (p *parser) advance() *token.Token {
	if p.isParsing() {
		p.current++
	}
	return p.previous()
}

func (p *parser) check(expected token.Type) bool {
	return p.isParsing() && p.peek().Type == expected
}

func (p *parser) match(expected ...token.Type) bool {
	if slices.ContainsFunc(expected, p.check) {
		p.advance()
		return true
	}
	return false
}

func (p *parser) expect(expected token.Type, msg string) *token.Token {
	if p.check(expected) {
		return p.advance()
	}
	panic(&Error{msg, p.peek()})
}

func (p *parser) synchronize() {
	p.advance()

	for p.isParsing() {
		if p.previous().Type == token.SEMICOLON {
			break
		}

		switch p.peek().Type {
		case token.CLASS, token.FUN, token.VAR, token.FOR, token.IF, token.WHILE, token.PRINT, token.RETURN:
			return
		default:
			p.advance()
		}
	}
}

//
// Stmt
//

func (p *parser) declaration() ast.Stmt {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(*Error); ok {
				p.errors = append(p.errors, err)
				p.synchronize()
			} else {
				panic(r)
			}
		}
	}()

	if p.match(token.VAR) {
		return p.varDeclaration()
	}

	return p.statement()
}

func (p *parser) varDeclaration() *ast.VarStmt {
	name := p.expect(token.IDENTIFIER, "Expect variable name")

	var value ast.Expr
	if p.match(token.EQUAL) {
		value = p.expression()
	}
	p.expect(token.SEMICOLON, "Expect ';' after variable declaration")

	return &ast.VarStmt{Name: name, Value: value}
}

func (p *parser) statement() ast.Stmt {
	if p.match(token.PRINT) {
		return p.printStatement()
	}

	if p.match(token.LEFT_BRACE) {
		return &ast.BlockStmt{Body: p.block()}
	}

	return p.expressionStatement()
}

func (p *parser) block() (body []ast.Stmt) {
	for p.isParsing() && !p.check(token.RIGHT_BRACE) {
		body = append(body, p.declaration())
	}
	p.expect(token.RIGHT_BRACE, "Expect '}' after block")

	return
}

func (p *parser) printStatement() *ast.PrintStmt {
	value := p.expression()
	p.expect(token.SEMICOLON, "Expect ';' after value")
	return &ast.PrintStmt{Value: value}
}

func (p *parser) expressionStatement() *ast.ExpressionStmt {
	expression := p.expression()
	p.expect(token.SEMICOLON, "Expect ';' after expression")
	return &ast.ExpressionStmt{Expression: expression}
}

//
// Expr
//

func (p *parser) expression() ast.Expr {
	return p.assignment()
}

func (p *parser) assignment() ast.Expr {
	expr := p.equality()

	if p.match(token.EQUAL) {
		equal := p.previous()
		if variable, ok := expr.(*ast.VariableExpr); ok {
			return &ast.AssignExpr{Name: variable.Name, Value: p.expression()}
		}
		panic(&Error{"Invalid assignment target", equal})
	}

	return expr
}

func (p *parser) equality() ast.Expr {
	expr := p.comparison()
	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		expr = &ast.BinaryExpr{Left: expr, Operator: p.previous(), Right: p.comparison()}
	}
	return expr
}

func (p *parser) comparison() ast.Expr {
	expr := p.term()
	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		expr = &ast.BinaryExpr{Left: expr, Operator: p.previous(), Right: p.term()}
	}
	return expr
}

func (p *parser) term() ast.Expr {
	expr := p.factor()
	for p.match(token.PLUS, token.MINUS) {
		expr = &ast.BinaryExpr{Left: expr, Operator: p.previous(), Right: p.factor()}
	}
	return expr
}

func (p *parser) factor() ast.Expr {
	expr := p.unary()
	for p.match(token.STAR, token.SLASH) {
		expr = &ast.BinaryExpr{Left: expr, Operator: p.previous(), Right: p.unary()}
	}
	return expr
}

func (p *parser) unary() ast.Expr {
	if p.match(token.BANG, token.MINUS) {
		return &ast.UnaryExpr{Operator: p.previous(), Right: p.unary()}
	}
	return p.primary()
}

func (p *parser) primary() ast.Expr {
	if p.match(token.NIL, token.TRUE, token.FALSE, token.STRING, token.NUMBER) {
		return &ast.LiteralExpr{Value: p.previous()}
	}

	if p.match(token.IDENTIFIER) {
		return &ast.VariableExpr{Name: p.previous()}
	}

	if p.match(token.LEFT_PAREN) {
		group := p.expression()
		p.expect(token.RIGHT_PAREN, "Expect ')' after expression")
		return &ast.GroupingExpr{Group: group}
	}

	panic(&Error{"Expect expression", p.peek()})
}
