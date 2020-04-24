package ast

import (
	"fmt"
	"strings"

	"../token"
)

// HashLiteral is a ast for map in this programming language
type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func (h *HashLiteral) expressionNode()      {}
func (h *HashLiteral) TokenLiteral() string { return h.Token.Literal }
func (h *HashLiteral) String() string {

	pairs := []string{}
	for key, value := range h.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	return fmt.Sprintf("{%s}", strings.Join(pairs, ", "))
}
