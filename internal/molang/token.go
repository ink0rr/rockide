package molang

import "regexp"

type TokenKind string

const (
	NUMBER     TokenKind = "NUMBER"
	STRING     TokenKind = "STRING"
	THIS       TokenKind = "THIS"
	METHOD     TokenKind = "METHOD"
	PREFIX     TokenKind = "PREFIX"
	KEYWORD    TokenKind = "KEYWORD"
	OPERATOR   TokenKind = "OPERATOR"
	PAREN      TokenKind = "PAREN"
	COMMA      TokenKind = "COMMA"
	WHITESPACE TokenKind = "WHITESPACE"
	UNKNOWN    TokenKind = "UNKNOWN"
)

type Token struct {
	Kind   TokenKind
	Offset uint32
	Length uint32
	Value  string
}

type tokenPattern struct {
	kind    TokenKind
	pattern *regexp.Regexp
}

var tokenPatterns = []tokenPattern{
	{NUMBER, regexp.MustCompile(`^[0-9]+(\.[0-9]+)?f?`)},
	{STRING, regexp.MustCompile(`^'[^']*'`)},
	{THIS, regexp.MustCompile(`^this`)},
	{METHOD, regexp.MustCompile(`^\.([a-zA-Z_][a-zA-Z0-9_.]*)?`)},
	{PREFIX, regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*`)},
	{KEYWORD, regexp.MustCompile(`^(break|continue|for_each|loop|return)`)},
	{OPERATOR, regexp.MustCompile(`^[+\-*/%><=!&|;:?]+`)},
	{PAREN, regexp.MustCompile(`^[\(\)\{\}\[\]]`)},
	{COMMA, regexp.MustCompile(`^,`)},
	{WHITESPACE, regexp.MustCompile(`^\s+`)},
	{UNKNOWN, regexp.MustCompile(`^.`)},
}
