package ast

import (
	"fmt"

	"../token"
)

// IndexExpression is the ast type to hold myArray[0]
type IndexExpression struct {
	Token token.Token // the [ token
	Left  Expression  // array name of literal array
	Index Expression  // the index trying to be accessed
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	return fmt.Sprintf(`(%s[%s])`, ie.Left.String(), ie.Index.String())
}
