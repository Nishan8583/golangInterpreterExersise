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
		return p.Statements[0].TokeLiteral()
	}
	return ""
}

//Node ... Each thing in the code is a node ?
type Node interface {
	TokeLiteral() string // for debugging purpose
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

// LetStatement Abstract structure for let
type LetStatement struct {
	Token token.Token // The tokent let
	Name  *Identifier
	Value Expression
}

// Identifier is structure for variable type
type Identifier struct {
	Token token.Token // token.IDENT type
	valye string      // The identifier value i.e varabile name
}

func (lex *LetStatement) statementNode() {}

//TokenLiteral returns the token.Literal value
func (lex *LetStatement) TokenLiteral() string {
	return lex.Token.Literal
}

func (id *Identifier) statementNode() {}

// TokenLiteral returns string
func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}
