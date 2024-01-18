package lexer

import (
	"fmt"
	"testing"

	"github.com/kosanin/monki/internal/token"
)

var errorStringFormat string = `tests[%d] - %q wrong, expected=%v, got=%v`

func Equal[T comparable](t *testing.T, testNumber int, what string, expected T, actual T) {
	t.Helper()

	if expected != actual {
		t.Fatalf(fmt.Sprintf(errorStringFormat, what, expected, actual))
	}
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
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
	lexer := New(input)
	for i, tt := range tests {
		tok := lexer.NextToken()
		Equal(t, i, "tokenType", tt.expectedType, tok.Type)
		Equal(t, i, "literal", tt.expectedLiteral, tok.Literal)
	}
}

func TestEmptyInput(t *testing.T) {
	input := ""
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.EOF, ""},
	}

	lexer := New(input)
	for i, tt := range tests {
		tok := lexer.NextToken()
		Equal(t, i, "tokenType", tt.expectedType, tok.Type)
		Equal(t, i, "literal", tt.expectedLiteral, tok.Literal)
	}

}
