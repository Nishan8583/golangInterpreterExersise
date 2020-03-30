package parser

import (
	"../ast"
	"../token"
)

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}

	// checking if the after "if" ( is present
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()

	// parsing the expressions inside ()
	expression.Condition = p.parseExpression(LOWEST)

	// checking for )
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	// checking if {
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	// if else statement is present
	if p.peekTokenIs(token.ELSE) {

		p.nextToken()
		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()
	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}
