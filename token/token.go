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
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	EOF       = "EOF"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
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
// {token.FUNCTION, "fn"},
// {token.LET, "let"},
// {token.EOF, ""},
