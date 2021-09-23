package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifier + literal

	// IDENT add, foobar, x, y, ...
	IDENT = "IDENT"
	// INT 12345
	INT = "INT"

	// operator

	// ASSIGN assign operator
	ASSIGN = "="
	// PLUS addition operator
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLUSH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// delimiter

	// COMMA comma
	COMMA = ","
	// SEMICOLON semicolon
	SEMICOLON = ";"

	// parentheses

	// LPAREN left parentheses
	LPAREN = "("
	// RPAREN right parentheses
	RPAREN = ")"

	// brace

	// LBRACE left brace
	LBRACE = "{"
	// RBRACE right brace
	RBRACE = "}"

	FUNCTION = "FUNCTION"

	LET    = "LET"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
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

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
