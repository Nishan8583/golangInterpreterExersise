// contains all the abstract syntax structures
package ast

import (
	"../token"

	"bytes"
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

//
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

//Node ... Each thing in the code is a node ?
type Node interface {
	TokenLiteral() string // for debugging purpose
	String() string       // This was added later, used only for debugging purpose
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

//String debug helper funciton for let statement
func (let *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(let.TokenLiteral() + " ")
	out.WriteString(let.Name.String())
	out.WriteString(" = ")

	if let.Value != nil {
		out.WriteString(let.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (id *Identifier) expressionNode() {}

// TokenLiteral returns string
func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}

func (id *Identifier) String() string { return id.Value }

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

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

/*-------------Expression AST-------------*/

//Ast for expressions such as 4+5 ...
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression  // The expression
}

func (es *ExpressionStatement) statementNode() {}

// TokebLiteral fullfills the Node interface requirement
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {

	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// Starting for integer literal
type IntegerLiteral struct {
	Token token.Token
	Value int64 // will hold the actual integer value
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// Now starts the prefix expression node
type PrefixExpression struct {
	Token    token.Token
	Operator string     // will contain the operator in left such as - or !
	Right    Expression // will contain the expression in the right side
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
