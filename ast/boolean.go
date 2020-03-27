package ast

import "../token"

// Boolean is the structure to handle boolean expressions
type Boolean struct {
	Token token.Token // The token type i.e token.True or token.False
	Value bool        // The actual value represented in go
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }
