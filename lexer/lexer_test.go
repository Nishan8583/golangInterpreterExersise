package lexer

import (
	"log"
	"testing"

	"../token"
)

// TestNextToken will test if lexing works
func TestNextToken(t *testing.T) {

	// the input form the command line
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x,y) {
		x+y;
	}
	
	let result = add(five,ten)`

	// table testing
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ":"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.COMMA, ","},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LBRACE, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RBRACE, ")"},
		{token.COMMA, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			log.Fatalf("Wrong type in test number %d, returned Type was %s while expected type was %s", i, tok.Type, tt.expectedType)
		} else {
			log.Printf("Correct type in test number %d, returned Type was %s while expected type was %s", i, tok.Type, tt.expectedType)
		}

		if tok.Literal != tt.expectedLiteral {
			log.Fatalf("Wrong literal in test number %d, returned Literal was %s, expected Literal was %s", i, tok.Literal, tt.expectedLiteral)
		} else {
			log.Printf("Correct literal in test number %d, returned Literal was %s, expected Literal was %s", i, tok.Literal, tt.expectedLiteral)
		}

	}
}
