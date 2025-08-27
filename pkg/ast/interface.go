// Package ast defines an AST (abstract syntax tree) for the Lox programming
// language.
package ast

import "fmt"

// Visitor is an AST visitor.
type Visitor interface {
	// Visit invokes the visitor on a given node.
	//
	// This function returns a boolean value, which is used by the [Walk]
	// function to determine whether AST traversal should continue or terminate.
	Visit(node Node) bool
}

// Children returns a node's immediate children.
func Children(node Node) (children []Node) {
	switch node := node.(type) {
	// Stmt

	case *BlockStmt:
		for _, b := range node.Body {
			children = append(children, b)
		}

	case *ExpressionStmt:
		children = append(children, node.Expression)

	case *PrintStmt:
		children = append(children, node.Value)

	case *VarStmt:
		if node.Value != nil {
			children = append(children, node.Value)
		}

	// Expr

	case *AssignExpr:
		children = append(children, node.Value)

	case *BinaryExpr:
		children = append(children, node.Left, node.Right)

	case *GroupingExpr:
		children = append(children, node.Group)

	case *LiteralExpr:
		// No children :(

	case *UnaryExpr:
		children = append(children, node.Right)

	case *VariableExpr:
		// No children :(

	default:
		panic(fmt.Errorf("Unexpected node type %T", node))
	}

	return
}

// Print returns a string representation of an AST node.
func Print(node Node) string {
	var p printer
	return p.Print(node)
}

// Walk traverses an AST using a given visitor.
func Walk(visitor Visitor, node Node) {
	if !visitor.Visit(node) {
		return
	}

	switch node := node.(type) {
	// Stmt

	case *BlockStmt:
		for _, b := range node.Body {
			Walk(visitor, b)
		}

	case *ExpressionStmt:
		Walk(visitor, node.Expression)

	case *PrintStmt:
		Walk(visitor, node.Value)

	case *VarStmt:
		if node.Value != nil {
			Walk(visitor, node.Value)
		}

	// Expr

	case *AssignExpr:
		Walk(visitor, node.Value)

	case *BinaryExpr:
		Walk(visitor, node.Left)
		Walk(visitor, node.Right)

	case *GroupingExpr:
		Walk(visitor, node.Group)

	case *LiteralExpr:
		// Do nothing

	case *UnaryExpr:
		Walk(visitor, node.Right)

	case *VariableExpr:
		// Do nothing

	default:
		panic(fmt.Errorf("Unexpected node type %T", node))
	}
}
