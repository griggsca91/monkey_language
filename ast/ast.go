
package ast

import (
	"monkey/token"
)

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

type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (ls *Identifier) expressionNode() {}
func (ls *Identifier) TokenLiteral() string { return ls.Token.Literal }


func (p *Program) TokenLiteral() string {
	if len(p.Statements) == 0 {
		return p.Statements[0].TokenLiteral()
	} 
	return ""
}