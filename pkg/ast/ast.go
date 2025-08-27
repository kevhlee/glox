package ast

import "github.com/kevhlee/glox/pkg/token"

// Node is an AST node.
type Node interface {
	node()
}

//
// Stmt
//

type (
	// Stmt is a statement AST node.
	Stmt interface {
		Node
		stmt()
	}

	// BlockStmt is a block statement AST node.
	BlockStmt struct {
		Body []Stmt
	}

	// ExpressionStmt is an expression statement AST node.
	ExpressionStmt struct {
		Expression Expr
	}

	// PrintStmt is a print statement AST node.
	PrintStmt struct {
		Value Expr
	}

	// VarStmt is a variable declaration statement AST node.
	VarStmt struct {
		Name  *token.Token
		Value Expr
	}
)

func (*BlockStmt) node() {}
func (*BlockStmt) stmt() {}

func (*ExpressionStmt) node() {}
func (*ExpressionStmt) stmt() {}

func (*PrintStmt) node() {}
func (*PrintStmt) stmt() {}

func (*VarStmt) node() {}
func (*VarStmt) stmt() {}

//
// Expr
//

type (
	// Expr is an expression AST node.
	Expr interface {
		Node
		expr()
	}

	// AssignExpr is an assignment expression AST node.
	AssignExpr struct {
		Name  *token.Token
		Value Expr
	}

	// BinaryExpr is a binary expression AST node.
	BinaryExpr struct {
		Left     Expr
		Operator *token.Token
		Right    Expr
	}

	// GroupingExpr is a grouped expression AST node.
	GroupingExpr struct {
		Group Expr
	}

	// LiteralExpr is a literal expression AST node.
	LiteralExpr struct {
		Value *token.Token
	}

	// UnaryExpr is a unary expression AST node.
	UnaryExpr struct {
		Operator *token.Token
		Right    Expr
	}

	// VariableExpr is a variable expression AST node.
	VariableExpr struct {
		Name *token.Token
	}
)

func (*AssignExpr) node() {}
func (*AssignExpr) expr() {}

func (*BinaryExpr) node() {}
func (*BinaryExpr) expr() {}

func (*GroupingExpr) node() {}
func (*GroupingExpr) expr() {}

func (*LiteralExpr) node() {}
func (*LiteralExpr) expr() {}

func (*UnaryExpr) node() {}
func (*UnaryExpr) expr() {}

func (*VariableExpr) node() {}
func (*VariableExpr) expr() {}
