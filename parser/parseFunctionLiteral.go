package parser

import (
	"../ast"
	"../token"
)

// parseFunctionLiteral() will parse the funciton and then return
func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) { // If the next token is not (, then return
		return nil
	}

	lit.Parameters = p.parseFunctionParameters() // Parse dem parameters

	if !p.expectPeek(token.LBRACE) { // parsing parameters funciton will have update the next token to be {, and if we do not get this, there will have be an error
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(token.RPAREN) { // if the next token is ), then there was no parameters passed
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal} // getting the first identifier
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(token.RPAREN) { // if after the parameter we do not have ), syntax error obviously
		return nil
	}
	return identifiers
}
