package lang

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t'
}
