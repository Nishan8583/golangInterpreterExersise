package ast

import (
	"bytes"
	"fmt"
	"strings"

	"../token"
)

// FunctionLiteral is the ast Type for storing function body type
type FunctionLiteral struct {
	Token      token.Token     // The token fn
	Parameters []*Identifier   // The parameters passed
	Body       *BlockStatement // The body of the function
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {

	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	return fmt.Sprintf("%s (%s) %s", out.String(), strings.Join(params, ","), fl.Body.String())
}
