/*
The lexer package parses the monkey lanauge and contains all the responsible function and types to produce the tokens
*/
package lexer

import (
	"../token"
)

// Lexer struct holds the total script input, the current position of input,
// readPosition will hold the next character to read,
// ch is the current byte that will be analyzed
type Lexer struct {
	input        string // the input string to parse
	position     int    // current posiion of input, the value will be read next time called
	readPosition int    // current reading position in input, the value at this position will be read
	ch           byte   // current char under examination
}

// New creates a new Lexer type and returns it
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// read the l.input, update l.ch,l.position and l.readPosition
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
	var tok token.Token // token to return

	l.skipWhiteSpace() // escaping all the whitespace
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{token.EQ, "=="}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

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
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			tok = token.Token{token.NOT_EQ, "!="}
		} else {
			tok = newToken(token.BANG, l.ch)
		}

	case '*':
		tok = newToken(token.ASTERIK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
			return tok
		}

	}
	l.readChar()
	return tok
}

// Returns a new tokenType
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Chekcinf if the passed byte is a letter
func isLetter(ch byte) bool {
	if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_' {
		return true // if the above condition is matched, then this was infact a letter, and thus we return true
	}
	return false // if the condiion was not matched, then we return false
}

// very clever and simple code here, i like it
func (l *Lexer) readIdentifier() string {
	position := l.position // getting initial postion of character to be read
	for isLetter(l.ch) {   // until l.ch is a letter
		l.readChar() // readChar() will increase the position
	}
	return l.input[position:l.position] // get the string
}

// Lets just skip the whitespace, there is also some great stuff here
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// checkinf if the byte is an integer
func isDigit(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

// reading the input and updating the char until non read postion was reached
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// see whats next in input, without updaing the positions
func (l *Lexer) peekChar() byte {
	if l.position >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
