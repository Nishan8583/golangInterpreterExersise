package ast

import "../token"

// StringLiteral is the ast type for holding string data type
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}

// TokenLiteral returns the literal token
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// String method returns the string value itself
func (sl *StringLiteral) String() string { return sl.Token.Literal }
