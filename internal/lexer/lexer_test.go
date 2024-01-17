package lexer

import (
	"fmt"
	"testing"

	"github.com/kosanin/monki/internal/token"
)

var errorStringFormat string = `tests[%d] - %q wrong, expected=%v, got=%v`

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
		if tok.Type != tt.expectedType {
			fmt.Println("literal " + tok.Literal)
			fmt.Println("expected literal " + tt.expectedLiteral)
			t.Fatalf(fmt.Sprintf(errorStringFormat, i, "tokenType", tt.expectedType, tok.Type))
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(fmt.Sprintf(errorStringFormat, i, "literal", tt.expectedLiteral, tok.Literal))
		}
	}
}
