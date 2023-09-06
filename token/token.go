package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

func New(tokenType TokenType, literal string) *Token {
	return &Token{Type: tokenType, Literal: literal}
}

const (
	ILLEGAL   = "ILLEGAL"
	IDENT     = "IDENT"
	INT       = "INT"
	STRING    = "STRING"
	ASSIGN    = "="
	BANG      = "!"
	EQ        = "=="
	NOT_EQ    = "!="
	LT        = "<"
	GT        = ">"
	PLUS      = "+"
	MINUS     = "-"
	ASTERISK  = "*"
	SLASH     = "/"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LET       = "LET"
	FUNCTION  = "FUNCTION"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
	EOF       = "EOF"
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
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}

	return IDENT
}
