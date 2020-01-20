package parser

import (
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
	prefix := p.prefixParseFns[p.curToken.Type] // getting the associated function of the tokenType
	if prefix == nil {
		return nil
	}

	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}
