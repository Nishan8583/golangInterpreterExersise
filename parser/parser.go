package parser

import (
	"fmt"
	"strconv"
	"time"

	"../ast"
	"../lexer"
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
}

// Parser type will create AST
type Parser struct {
	l         *lexer.Lexer // Lexer type, will be used to get tokens
	curToken  token.Token  // The current token we will work with
	peekToken token.Token  // The next token we will work with, sometime the curtoken may not have complete value

	// The below are added for parsing expressions
	prefixParseFns map[token.TokenType]perfixParseFn // assosicate certain token types to prefix parsing functions
	infixParseFns  map[token.TokenType]infixParseFn  // associate certain token types to infix parsing functions
}

// New takes a lexer type and returns parser type
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	p.nextToken() // Initilization of both curtoken and peekToken
	p.nextToken()

	// Registering the token type to their parsing functions
	p.prefixParseFns = make(map[token.TokenType]perfixParseFn, 1)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	// Now registering infix operator
	p.infixParseFns = make(map[token.TokenType]infixParseFn, 1)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERIK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parses the program and returns a ast.Program type
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET: // if the token was LET type
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

// The below function parses the LetStatement
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) { // if the next statement was not a variable
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) { // if the current token now is not equals
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt

}

// now parse return statement
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
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
	fmt.Println("Parsing for ", p.curToken, "Precedence", precedence)
	prefix := p.prefixParseFns[p.curToken.Type] // getting the associated function of the tokenType
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		fmt.Println("DEBUG prefix Failure for ", p.curToken)
		time.Sleep(1 * time.Minute)
		return nil
	}

	leftExp := prefix()
	fmt.Println("DEBUG prefix Success for ", p.curToken, leftExp)
	// Now check fi seimicolon is there, and until precedencei  less then the precedence of nex token
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			fmt.Println("DEBUG Infix Failure for ", p.curToken)
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
		fmt.Println("DEBUG Infix Success for ", p.curToken, leftExp)
	}
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

	fmt.Println("*******Before nextToken() the token is ", p.curToken)
	if p.curToken.Type == token.NOT_EQ {
		expression.Right = p.parseExpression(precedence)
	} else {
		p.nextToken()
		expression.Right = p.parseExpression(precedence)
	}

	fmt.Println("------------After nextToken() the token is ", p.curToken)

	return expression
}
