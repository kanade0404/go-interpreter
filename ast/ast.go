package ast

import "interpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token
	Value string
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}


type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (i *Identifier) expressionNode()  {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (s *LetStatement) statementNode()  {}
func (s *LetStatement) TokenLiteral() string {
	return s.Token.Literal
}

type ReturnStatement struct {
	Token token.Token // 'return' トークン
	ReturnValue Expression
}

func (r *ReturnStatement) statementNode()  {}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}