package lexer

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isString(ch byte) bool {
	return ch != '"' && ch != 0
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
