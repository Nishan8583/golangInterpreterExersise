package lexer

import "../token"

// Lexer struct holds the total script input, the current position of input,
// readPosition will hold the next character to read,
// ch is the current byte that will be analyzed
type Lexer struct {
	input        string
	position     int  // current posiion of input, the value will be read next time called
	readPosition int  // current reading position in input, the value at this position will be read
	ch           byte // current char under examination
}

// New creates a new Lexer type and returns it
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // if end was reached
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition] // if EOF not reached, then update the l.ch to the position of readPosition
	}
	l.position = l.readPosition
	l.readPosition++

}

// NextToken get the next token type
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
