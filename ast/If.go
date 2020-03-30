package ast

import (
	"bytes"

	"../token"
)

// IfExpression is the ast for the if () {} else {}
type IfExpression struct {
	Token       token.Token     // the if token
	Condition   Expression      // the conditon of if
	Consequence *BlockStatement // the code to execute if the if statements returns true
	Alternative *BlockStatement // else  statement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	output := "if" + ie.Condition.String() + " " + ie.Consequence.String()

	if ie.Alternative != nil {
		output = output + "else"
		output = output + ie.Alternative.String()
	}

	return output
}

// BlockStatement holds the blocks of code inside if and else condition
type BlockStatement struct {
	Token      token.Token // {
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
