package parser

import (
	"../ast"
	"../token"
)

// function for parsing hash type for the programming language
func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.curToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)

	// Loop till we get '}'
	for !p.peekTokenIs(token.RBRACE) {

		// skil token '{'
		p.nextToken()

		// Get the key out
		key := p.parseExpression(LOWEST)

		// If the next token was not ':'
		if !p.expectPeek(token.COLON) {
			return nil
		}

		// skip token ':'
		p.nextToken()

		// Getting value
		value := p.parseExpression(LOWEST)

		hash.Pairs[key] = value

		// If the next token is not either '}' or ','
		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil
		}
	}

	// IF the last token is not '}'
	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return hash
}
