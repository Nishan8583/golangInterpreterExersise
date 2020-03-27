package parser

import (
	"../token"

	"../ast"
)

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil // Do i need to increase this
	}
	return exp
}
