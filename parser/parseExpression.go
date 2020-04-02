package parser

import (
	"fmt"
	"strconv"
	"time"

	"../ast"
	"../token"
)

// constants used for operator precedence
const (
	_ = iota // this will be zero, from below will start from 1
	LOWEST
	EQAULS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or +x
	CALL        // myfunc()

)

// Setting up operator precendence
var precedences = map[token.TokenType]int{
	token.EQ:      EQAULS,
	token.NOT_EQ:  EQAULS,
	token.LT:      LESSGREATER,
	token.GT:      LESSGREATER,
	token.PLUS:    SUM,
	token.MINUS:   SUM,
	token.SLASH:   PRODUCT,
	token.ASTERIK: PRODUCT,
	token.LPAREN:  CALL,
}

/*Code added after expression parsing section was started*/
type (
	perfixParseFn func() ast.Expression               // this funciton is called for prefix operator parsing conditions
	infixParseFn  func(ast.Expression) ast.Expression // this function is called when infix parsing function is encoutered
)

//helper functions that will register the the token type to their parsing funciton
func (p *Parser) registerPrefix(tokenType token.TokenType, fn perfixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {

	// First try and parse the prefix
	prefix := p.prefixParseFns[p.curToken.Type] // getting the associated function of the tokenType
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		time.Sleep(1 * time.Minute)
		return nil
	}

	leftExp := prefix()
	// Now check fi seimicolon is there, and until precedencei  less then the precedence of nex token
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()
		leftExp = infix(leftExp)

	}
	fmt.Printf("\n\n\n\n\n")
	// finished adding for the infix parsing section
	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		fmt.Println("ERROR while trying to generate integer literal", err)
		return nil
	}

	lit.Value = value
	return lit
}

// For printing better error messages
func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("No prefix funciton associated to this token type %s ", t)
	fmt.Println(msg)
}

// now writing function that actually parses the prefixExpression
func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken() // to get to the integer

	expression.Right = p.parseExpression(PREFIX) // now the current token type is int. so parseIntegerLiteral is callsed
	return expression
}

// The two function below will give the precedence value for various operator
func (p *Parser) peekPrecedence() int {
	prec, ok := precedences[p.peekToken.Type]
	if ok {
		return prec
	}

	return LOWEST
}

func (p *Parser) curPrecedence() int {
	prec, ok := precedences[p.curToken.Type]
	if ok {
		return prec
	}

	return LOWEST
}

// function that actually parses the infix expressions
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left, // the previous left expression
	}

	precedence := p.curPrecedence() // Get the precedence of the current token
	//if p.curToken.Type == token.NOT_EQ {
	//	expression.Right = p.parseExpression(precedence)
	//} else {
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	//}

	return expression
}
