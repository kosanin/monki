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

	l.skipWhitespace()

	var tok token.Token
	switch l.ch {
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
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isAlpha(l.ch) {
			// read whole string and check if its keyword or identifier
			tok = l.identifier()
		}
		if isNumeric(l.ch) {
			// read whole number
			tok = l.number()
		}
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

func isAlpha(ch byte) bool {
	return ('z' >= ch && ch >= 'a') || ('Z' >= ch && ch >= 'A')
}

func isNumeric(ch byte) bool {
	return ('9' >= ch && ch >= '0')
}

func (l *Lexer) identifier() token.Token {

	// reads let five = 5;
	for isAlphanumeric(l.peek()) {
		l.currentCharacterPosition += 1
	}
	literal := l.input[l.startOfTheTokenPosition:l.currentCharacterPosition]
	var tokenType token.TokenType
	keywordType, ok := token.Keywords[literal]
	if ok {
		tokenType = keywordType
	} else {
		tokenType = token.IDENT
	}
	return token.Token{
		Type:    tokenType,
		Literal: literal,
	}
}

func (l *Lexer) number() token.Token {
	for isNumeric(l.peek()) {
		l.currentCharacterPosition += 1
	}
	literal := l.input[l.startOfTheTokenPosition:l.currentCharacterPosition]
	return token.Token{
		Type:    token.INT,
		Literal: literal,
	}
}

func isAlphanumeric(ch byte) bool {
	return isAlpha(ch) || isNumeric(ch) || ch == '_'
}

func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return 0
	}
	return l.input[l.currentCharacterPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}
