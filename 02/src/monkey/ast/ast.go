package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type Expression interface {
	Node
	expressionNode()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (es *ExpressionStatement) statementNode() {

}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Node interface {
	TokenLiteral() string
	String() string
}
type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (i IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i IntegerLiteral) String() string {
	return i.Token.Literal
}

func (i IntegerLiteral) expressionNode() {
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
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

func (id *Identifier) String() string {
	return id.Value
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}

func (p *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Right.String())
	out.WriteString(")")
	return out.String()
}

func (p *PrefixExpression) expressionNode() {
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(i.Left.String())
	out.WriteString(" " + i.Operator + " ")
	out.WriteString(i.Right.String())
	out.WriteString(")")
	return out.String()
}

func (i *InfixExpression) expressionNode() {
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b Boolean) expressionNode() {
}

func (b Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b Boolean) String() string {
	return b.Token.Literal
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (s *BlockStatement) statementNode() {
}

func (s *BlockStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *BlockStatement) String() string {
	var out bytes.Buffer
	for _, stmt := range s.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Altenative  *BlockStatement
}

func (i *IfExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(i.Condition.String())
	out.WriteString(" ")
	out.WriteString(i.Consequence.String())

	if i.Altenative != nil {
		out.WriteString("else")
		out.WriteString("")
		out.WriteString(i.Altenative.String())
	}
	return out.String()
}

func (i *IfExpression) expressionNode() {
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (f *FunctionLiteral) TokenLiteral() string {
	return f.Token.Literal
}

func (f *FunctionLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(f.TokenLiteral())
	out.WriteString("(")
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") ")
	out.WriteString(f.Body.String())
	return out.String()
}

func (f FunctionLiteral) expressionNode() {
	//TODO implement me
	panic("implement me")
}

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (c *CallExpression) TokenLiteral() string {
	return c.Token.Literal
}

func (c *CallExpression) String() string {
	var out bytes.Buffer

	var args []string
	for _, a := range c.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(c.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ","))
	out.WriteString(")")
	return out.String()
}

func (c *CallExpression) expressionNode() {
}
