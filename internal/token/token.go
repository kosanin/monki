package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = iota
	EOF

	IDENT
	INT

	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	FN
	LET

	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH

	LT
	GT
	LTE
	GTE
	EQ
	BANG_EQ
)

var Keywords = map[string]TokenType{
	"let": LET,
	"fn":  FN,
}
