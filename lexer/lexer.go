package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) nextChIsWhiteSpace() bool {
	if l.readPosition < len(l.input) {
		ch := l.input[l.readPosition]
		return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
	}
	return false
}

func (l *Lexer) nextChIsLetter() bool {
	if l.readPosition < len(l.input) {
		return isLetter(l.input[l.readPosition])
	}
	return false
}

func (l *Lexer) nextChIsNumber() bool {
	if l.readPosition < len(l.input) {
		return isDigit(l.input[l.readPosition])
	}
	return false
}

func (l *Lexer) skipWhiteSpace() {
	for l.nextChIsWhiteSpace() {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	start_position := l.position
	for l.nextChIsLetter() {
		l.readChar()
	}
	return l.input[start_position : l.position+1]
}

func (l *Lexer) readNumber() string {
	start_position := l.position
	for l.nextChIsNumber() {
		l.readChar()
	}
	return l.input[start_position : l.position+1]
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhiteSpace()
	l.readChar()

	var tokenType token.TokenType
	literal := string(l.ch)

	switch l.ch {
	case '=':
		tokenType = token.ASSIGN
	case ';':
		tokenType = token.SEMICOLON
	case '(':
		tokenType = token.LPAREN
	case ')':
		tokenType = token.RPAREN
	case ',':
		tokenType = token.COMMA
	case '+':
		tokenType = token.PLUS
	case '-':
		tokenType = token.MINUS
	case '!':
		tokenType = token.BANG
	case '/':
		tokenType = token.SLASH
	case '*':
		tokenType = token.ASTERISK
	case '<':
		tokenType = token.LT
	case '>':
		tokenType = token.GT
	case '{':
		tokenType = token.LBRACE
	case '}':
		tokenType = token.RBRACE
	case 0:
		tokenType = token.EOF
		literal = ""
	default:
		if isLetter(l.ch) {
			literal = l.readIdentifier()
			tokenType = token.LookupIdent(literal)
		} else if isDigit(l.ch) {
			tokenType = token.INT
			literal = l.readNumber()
		} else {
			tokenType = token.ILLEGAL
		}
	}

	return newToken(tokenType, literal)
}
