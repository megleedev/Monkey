package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals

	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 1 2 3 4 5 6 7 8 9

	// Operators

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// Delimiters

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{

	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {

	if tok, ok := keywords[ident]; ok {

		return tok
	}

	return IDENT
}
