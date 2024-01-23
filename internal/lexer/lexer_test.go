package lexer

import (
	"fmt"
	"os"
	"testing"

	"github.com/kosanin/monki/internal/token"
)

var errorStringFormat string = `tests[%d] - %q wrong, expected=%v, got=%v`

func Equal[T comparable](t *testing.T, testNumber int, what string, expected T, actual T) {
	t.Helper()

	if expected != actual {
		t.Fatalf(fmt.Sprintf(errorStringFormat, testNumber, what, expected, actual))
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
		Equal(t, i, "literal", tt.expectedLiteral, tok.Literal)
		Equal(t, i, "tokenType", tt.expectedType, tok.Type)
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

func TestProperMonkiSourceCode(t *testing.T) {
	input, err := os.ReadFile("testdata/sample.mnk")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "125"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FN, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	lexer := New(string(input))
	for i, tt := range tests {
		tok := lexer.NextToken()
		Equal(t, i, "literal", tt.expectedLiteral, tok.Literal)
		Equal(t, i, "tokenType", tt.expectedType, tok.Type)
	}
}

func TestOperators(t *testing.T) {
	input, err := os.ReadFile("testdata/operators.mnk")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.PLUS, "+"},
		{token.ASTERISK, "*"},
		{token.GT, ">"},
		{token.LT, "<"},
		{token.GTE, ">="},
		{token.LTE, "<="},
		{token.EQ, "=="},
		{token.BANG_EQ, "!="},
		{token.EOF, ""},
	}
	lexer := New(string(input))
	for i, tt := range tests {
		tok := lexer.NextToken()
		Equal(t, i, "literal", tt.expectedLiteral, tok.Literal)
		Equal(t, i, "tokenType", tt.expectedType, tok.Type)
	}
}
