// contains all the tokens that will be used int he monkey programming langauge
package token

// TokenType is an alias for string
type TokenType string

// Token is a struct that holds the TokenType and the literal value
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// comments
	ILLEGEAL = "ILLEGAL" // if illegal
	EOF      = "EOF"

	IDENT = "IDENT" // identifiers i.e function, variable names
	INT   = "int"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent will check the the identifier is a keyword
func LookupIdent(indent string) TokenType {
	if tok, ok := keywords[indent]; ok {
		return tok
	}
	return IDENT
}
