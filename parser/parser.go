package parser

import (
	"../ast"
	"../lexer"
	"../token"
)

// Parser type will create AST
type Parser struct {
	l         *lexer.Lexer // Lexer type, will be used to get tokens
	curToken  token.Token  // The current token we will work with
	peekToken token.Token  // The next token we will work with, sometime the curtoken may not have complete value
}

func New(l *lexer.Lexer) *Parser {
	p := Parser{
		l: l,
	}

	p.nextToken() // Initilization of both curtoken and peekToken
	p.nextToken()

	return &p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
