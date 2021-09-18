package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// identifier + literal

	// INDENT add, foobar, x, y, ...
	INDENT = "INDENT"
	// INT 12345
	INT = "INT"

	// operator

	// ASSIGN assign operator
	ASSIGN = "="
	// PLUS addition operator
	PLUS = "+"

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
)