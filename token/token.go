package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	LT        = "<"
	GT        = ">"
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
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

// {token.ILLEGAL, ""},
// {token.IDENT, "ident"},
// {token.INT, "000"},
// {token.ASSIGN, "="},
// {token.LT, "<"},
// {token.GT, ">"},
// {token.PLUS, "+"},
// {token.MINUS, "-"},
// {token.BANG, "!"},
// {token.ASTERISK, "*"},
// {token.SLASH, "/"},
// {token.COMMA, ","},
// {token.SEMICOLON, ";"},
// {token.LPAREN, "("},
// {token.RPAREN, ")"},
// {token.LBRACE, "{"},
// {token.RBRACE, "}"},
// {token.LET, "let"},
// {token.FUNCTION, "fn"},
// {token.TRUE, "true"},
// {token.FALSE, "false"},
// {token.IF, "if"},
// {token.ELSE, "else"},
// {token.RETURN, "return"},
// {token.EOF, ""},
