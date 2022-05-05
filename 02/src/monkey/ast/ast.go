package ast

import "monkey/token"

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	
}

type Identifier struct {
	Token token.Token
	Value string
}

func (id *Identifier) expressionNode() {
}
func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}

type Node interface {
	TokenLiteral()
}
type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Statement
	Token token.Token // let
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {
}