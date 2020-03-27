package parser

import (
	"../ast"
	"../token"
)

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: p.curToken,
		Value: p.curTokenIs(token.TRUE), // if the current token is True, true is returned, else false, beautiful
	}
}
