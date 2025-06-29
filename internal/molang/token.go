package molang

import "regexp"

type TokenKind string

const (
	KindNumber   TokenKind = "NUMBER"
	KindString   TokenKind = "STRING"
	KindMacro    TokenKind = "THIS"
	KindMethod   TokenKind = "METHOD"
	KindPrefix   TokenKind = "PREFIX"
	KindKeyword  TokenKind = "KEYWORD"
	KindOperator TokenKind = "OPERATOR"
	KindParen    TokenKind = "PAREN"
	KindComma    TokenKind = "COMMA"
	KindUnknown  TokenKind = "UNKNOWN"
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
	{KindNumber, regexp.MustCompile(`^[0-9]+(\.[0-9]+)?f?`)},
	{KindString, regexp.MustCompile(`^'[^']*'`)},
	{KindMacro, regexp.MustCompile(`(?i)^(this|true|false)`)},
	{KindMethod, regexp.MustCompile(`^\.([a-zA-Z_][a-zA-Z0-9_.]*)?`)},
	{KindPrefix, regexp.MustCompile(`(?i)^(q|v|t|c|query|variable|temp|context|math|array|geometry|material|texture)\b`)},
	{KindKeyword, regexp.MustCompile(`(?i)^(break|continue|for_each|loop|return)`)},
	{KindOperator, regexp.MustCompile(`^[+\-*/%><=!&|;:?]+`)},
	{KindParen, regexp.MustCompile(`^[\(\)\{\}\[\]]`)},
	{KindComma, regexp.MustCompile(`^,`)},
	{KindUnknown, regexp.MustCompile(`^\s+`)},
	{KindUnknown, regexp.MustCompile(`^\w+`)},
	{KindUnknown, regexp.MustCompile(`^.`)},
}
