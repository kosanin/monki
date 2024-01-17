package lexer

import (
	"github.com/kosanin/monki/internal/token"
)

type Lexer struct {
	input                    string
	startOfTheTokenPosition  int
	currentCharacterPosition int
	ch                       byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) NextToken() token.Token {

	l.readChar()

	var tok token.Token
	switch l.ch {
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
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
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	}
	return tok
}

func newToken(tokenType token.TokenType, literal byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(literal),
	}
}

func (l *Lexer) readChar() {
	if l.isAtEnd() {
		l.ch = 0
	} else {
		l.ch = l.input[l.currentCharacterPosition]
	}
	l.startOfTheTokenPosition = l.currentCharacterPosition
	l.currentCharacterPosition += 1
}

func (l *Lexer) isAtEnd() bool {
	return l.currentCharacterPosition >= len(l.input)
}
