package lexer

import (
	"log"
	"testing"

	"../token"
)

// TestNextToken will test if lexing works
func TestNextToken(t *testing.T) {

	// the input form the command line
	input := `=+(){},;`

	// table testing
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
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
