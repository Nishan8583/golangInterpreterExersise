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
	ILLEGAL = "ILLEGAL" // if illegal
	EOF     = "EOF"

	IDENT = "IDENT" // identifiers i.e function, variable names
	INT   = "int"

	// operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERIK   = "*"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"
	EQ        = "=="
	NOT_EQ    = "!="
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywpords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	// extending parser, chapter 4 section
	STRING = "STRING"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent will check the the identifier is a keyword
func LookupIdent(indent string) TokenType {
	if tok, ok := keywords[indent]; ok {
		return tok
	}
	return IDENT
}
