package ast

import (
	"fmt"
	"strings"

	"../token"
)

// function call will be like this add(2+2,3)
// in the above example add is an identifier, and arguments are expression,
// But note that add will return the funciton literal in the end, so this can be treated as expresion as well

// CallExpression will be the ast for function call
type CallExpression struct {
	Token     token.Token // will hold (
	Function  Expression  // Identifier for function literal
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	args := []string{}

	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	return fmt.Sprintf("%s(%s)", ce.Function.String(), strings.Join(args, ","))
}
