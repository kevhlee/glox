package lox

import (
	"fmt"
	"strconv"

	"github.com/kevhlee/glox/internal/stack"
	"github.com/kevhlee/glox/pkg/ast"
	"github.com/kevhlee/glox/pkg/token"
)

// Contains the internal state and logic of the interpreter.
type interpreter struct {
	env      *Environment
	globals  *Environment
	operands stack.Stack[any]
}

func (ip *interpreter) Interpret(globals *Environment, body []ast.Stmt) (err error) {
	ip.env = globals
	ip.globals = globals

	defer func() {
		ip.env = nil
		ip.globals = nil

		if r := recover(); r != nil {
			if r, ok := r.(*Error); ok {
				err = r
			} else {
				panic(r)
			}
		}
	}()

	for _, stmt := range body {
		ast.Walk(ip, stmt)
	}
	return
}

// Visit implements the [ast.Visitor] interface.
func (ip *interpreter) Visit(node ast.Node) bool {
	switch node := node.(type) {
	// Stmt
	case *ast.BlockStmt:
		ip.handleBlockStmt(node)
	case *ast.ExpressionStmt:
		ip.handleExprStmt(node)
	case *ast.PrintStmt:
		ip.handlePrintStmt(node)
	case *ast.VarStmt:
		ip.handleVarStmt(node)

	// Expr

	case *ast.AssignExpr:
		ip.handleAssignExpr(node)
	case *ast.BinaryExpr:
		ip.handleBinaryExpr(node)
	case *ast.GroupingExpr:
		ip.handleGroupingExpr(node)
	case *ast.LiteralExpr:
		ip.handleLiteralExpr(node)
	case *ast.UnaryExpr:
		ip.handleUnaryExpr(node)
	case *ast.VariableExpr:
		ip.handleVariableExpr(node)
	}

	return false
}

func (ip *interpreter) execute(stmt ast.Stmt) {
	ast.Walk(ip, stmt)
}

func (ip *interpreter) evaluate(expr ast.Expr) any {
	ast.Walk(ip, expr)

	if operand, ok := ip.operands.Pop(); ok {
		return operand
	}
	return nil
}

func isTruthy(value any) bool {
	if b, ok := value.(bool); ok {
		return b
	}
	return value != nil
}

//
// Stmt
//

func (ip *interpreter) handleBlockStmt(stmt *ast.BlockStmt) {
	enclosing := ip.env

	defer func() {
		ip.env = enclosing
	}()

	ip.env = newInnerEnvironment(ip.env)

	for _, b := range stmt.Body {
		ip.execute(b)
	}
}

func (ip *interpreter) handleExprStmt(stmt *ast.ExpressionStmt) {
	ip.evaluate(stmt.Expression)
}

func (ip *interpreter) handlePrintStmt(stmt *ast.PrintStmt) {
	if value := ip.evaluate(stmt.Value); value != nil {
		fmt.Printf("%v\n", value)
	} else {
		fmt.Println("nil")
	}
}

func (ip *interpreter) handleVarStmt(stmt *ast.VarStmt) {
	var value any
	if stmt.Value != nil {
		value = ip.evaluate(stmt.Value)
	}
	ip.env.Define(stmt.Name.Lexeme, value)
}

//
// Expr
//

func (ip *interpreter) handleAssignExpr(expr *ast.AssignExpr) {
	value := ip.evaluate(expr.Value)

	if !ip.env.Assign(expr.Name.Lexeme, value) {
		panic(&Error{
			Msg:  fmt.Sprintf("Undefined variable '%s'", expr.Name.Lexeme),
			Line: expr.Name.Line,
		})
	}

	ip.operands.Push(value)
}

func (ip *interpreter) handleBinaryExpr(expr *ast.BinaryExpr) {
	l, r := ip.evaluate(expr.Left), ip.evaluate(expr.Right)

	switch expr.Operator.Type {
	case token.PLUS:
		if lhs, ok := l.(float64); ok {
			if rhs, ok := r.(float64); ok {
				ip.operands.Push(lhs + rhs)
				return
			}
		}
		if lhs, ok := l.(string); ok {
			if rhs, ok := r.(string); ok {
				ip.operands.Push(lhs + rhs)
				return
			}
		}
		panic(&Error{"Operands must be two numbers or two strings", expr.Operator.Line})

	case token.BANG_EQUAL:
		ip.operands.Push(l != r)
		return

	case token.EQUAL_EQUAL:
		ip.operands.Push(l == r)
		return
	}

	lhs, lok := l.(float64)
	rhs, rok := r.(float64)

	if !(lok && rok) {
		panic(&Error{"Operands must be numbers", expr.Operator.Line})
	}

	switch expr.Operator.Type {
	case token.MINUS:
		ip.operands.Push(lhs - rhs)
	case token.STAR:
		ip.operands.Push(lhs * rhs)
	case token.SLASH:
		ip.operands.Push(lhs / rhs)
	case token.GREATER:
		ip.operands.Push(lhs > rhs)
	case token.GREATER_EQUAL:
		ip.operands.Push(lhs >= rhs)
	case token.LESS:
		ip.operands.Push(lhs < rhs)
	case token.LESS_EQUAL:
		ip.operands.Push(lhs <= rhs)
	}
}

func (ip *interpreter) handleGroupingExpr(expr *ast.GroupingExpr) {
	ast.Walk(ip, expr.Group)
}

func (ip *interpreter) handleLiteralExpr(expr *ast.LiteralExpr) {
	switch value := expr.Value; value.Type {
	case token.NIL:
		ip.operands.Push(nil)

	case token.TRUE:
		ip.operands.Push(true)

	case token.FALSE:
		ip.operands.Push(false)

	case token.STRING:
		ip.operands.Push(value.Lexeme[1 : len(value.Lexeme)-1])

	case token.NUMBER:
		number, err := strconv.ParseFloat(value.Lexeme, 64)
		if err != nil {
			panic(&Error{"Invalid number literal", value.Line})
		}
		ip.operands.Push(number)
	}
}

func (ip *interpreter) handleUnaryExpr(expr *ast.UnaryExpr) {
	r := ip.evaluate(expr.Right)

	switch expr.Operator.Type {
	case token.BANG:
		ip.operands.Push(!isTruthy(r))

	case token.MINUS:
		if rhs, ok := r.(float64); ok {
			ip.operands.Push(-rhs)
			return
		}
		panic(&Error{"Operand must be a number", expr.Operator.Line})
	}
}

func (ip *interpreter) handleVariableExpr(expr *ast.VariableExpr) {
	if value, ok := ip.env.Get(expr.Name.Lexeme); ok {
		ip.operands.Push(value)
		return
	}

	panic(&Error{
		Msg:  fmt.Sprintf("Undefined variable '%s'", expr.Name.Lexeme),
		Line: expr.Name.Line,
	})
}
