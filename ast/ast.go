// contains all the abstract syntax structures
package ast

import (
	"../token"
)

type Program struct {
	Statements []Statement // root node this is
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

//Node ... Each thing in the code is a node ?
type Node interface {
	TokenLiteral() string // for debugging purpose
}

// Statement is something that does not produce any value
type Statement interface {
	Node            // statement is a node
	statementNode() // For debugging as well
}

// Expression is something that produces value
type Expression interface {
	Node
	expressionNode()
}

/*-------------LET STATEMENT AST-------------*/

// LetStatement Abstract structure for let
type LetStatement struct {
	Token token.Token // The tokent let
	Name  *Identifier
	Value Expression
}

// Identifier is structure for variable type
type Identifier struct {
	Token token.Token // token.IDENT type
	Value string      // The identifier value i.e varabile name
}

func (let *LetStatement) statementNode() {}

//TokenLiteral returns the token.Literal value
func (let *LetStatement) TokenLiteral() string {
	return let.Token.Literal
}

func (id *Identifier) statementNode() {}

// TokenLiteral returns string
func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}

/*-------------Return STATEMENT AST-------------*/

// ReturnStatement structure as the ast for return <expression>
type ReturnStatement struct {
	Token       token.Token // holds the return token
	ReturnValue Expression  // will hold the returned value
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral will the literal token itself
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

/*-------------Expression AST-------------*/

//Ast for expressions such as 4+5 ...
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression  // The expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
