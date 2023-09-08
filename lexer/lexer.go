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
	return &Lexer{input: input}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readChar() {
	l.ch = l.peekChar()
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readString(judgFunc func(byte) bool) string {
	start_position := l.position
	for judgFunc(l.peekChar()) {
		l.readChar()
	}

	if l.position == len(l.input) {
		return l.input[start_position:l.position]
	}

	if l.position < len(l.input) {
		return l.input[start_position : l.position+1]
	}

	return ""
}

func (l *Lexer) NextToken() token.Token {
	l.readString(isWhiteSpace) // Delete WhiteSpace
	l.readChar()

	var tokenType token.TokenType
	literal := string(l.ch)

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tokenType = token.EQ
			literal = "=="
		} else {
			tokenType = token.ASSIGN
		}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tokenType = token.NOT_EQ
			literal = "!="
		} else {
			tokenType = token.BANG
		}
	case '"':
		tokenType = token.STRING
		l.readChar() // 初めの「"」を飛ばす
		literal = l.readString(isString)
		l.readChar() // 最後の「"」を飛ばす
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
	case '[':
		tokenType = token.LBRACKET
		literal = "["
	case ']':
		tokenType = token.RBRACKET
		literal = "]"
	case 0:
		tokenType = token.EOF
		literal = ""
	default:
		if isLetter(l.ch) {
			literal = l.readString(isLetter)
			tokenType = token.LookupIdent(literal)
		} else if isDigit(l.ch) {
			literal = l.readString(isDigit)
			tokenType = token.INT
		} else {
			tokenType = token.ILLEGAL
		}
	}

	return *token.New(tokenType, literal)
}
