package ast

import (
	"fmt"
	"strings"

	"../token"
)

// ArrayLiteral is the ast type for arrays
type ArrayLiteral struct {
	Token    token.Token // The token is '['
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}

// TokenLiteral returns the '[' token literal lol
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }

func (al *ArrayLiteral) String() string {
	//var out bytes.Buffer

	elements := []string{}

	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	return fmt.Sprintf("[%s]", strings.Join(elements, ", "))
}
