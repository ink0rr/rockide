package lang

import (
	"iter"
)

type Lexer struct {
	src   []rune
	i     int
	line  uint32
	col   uint32
	state State
}

func NewLexer(input []rune) *Lexer {
	return &Lexer{
		src:  input,
		line: 0,
		col:  0,
	}
}

func (l *Lexer) eof() bool {
	return l.i >= len(l.src)
}

func (l *Lexer) peek() rune {
	if l.eof() {
		return 0
	}
	return l.src[l.i]
}

func (l *Lexer) hasPrefix(prefix string) bool {
	runes := []rune(prefix)
	if l.i+len(runes) > len(l.src) {
		return false
	}
	for j, r := range runes {
		if l.src[l.i+j] != r {
			return false
		}
	}
	return true
}

func (l *Lexer) pos() Position {
	return Position{
		Line:   l.line,
		Column: l.col,
	}
}

func (l *Lexer) advance(size ...int) rune {
	step := 1
	if len(size) > 0 {
		step = size[0]
	}
	var r rune
	for i := 0; i < step; i++ {
		r = l.peek()
		l.i++
		if r == '\n' {
			l.line++
			l.col = 0
		} else {
			l.col++
		}
	}
	return r
}

func (l *Lexer) emit(kind TokenKind, startPos Position, text string) Token {
	token := Token{
		Kind:   kind,
		Offset: uint32(l.i - len([]rune(text))),
		Start:  startPos,
		End:    l.pos(),
		Value:  text,
	}
	return token
}

func (l *Lexer) collectWhile(cond func(rune) bool) string {
	runes := []rune{}
	for !l.eof() && cond(l.peek()) {
		runes = append(runes, l.advance())
	}
	return string(runes)
}

func (l *Lexer) getFormatCode() string {
	if l.peek() == SectionSign {
		if l.i+1 < len(l.src) {
			next := l.src[l.i+1]
			if next != 0 && !isWhitespace(next) {
				return string([]rune{SectionSign, next})
			}
		}
	}
	return ""
}

func (l *Lexer) getFormatSpecifier() string {
	if l.peek() == '%' {
		if l.i+1 < len(l.src) {
			next := l.src[l.i+1]
			res := "%" + string(next)
			if next == 's' || next == 'd' || next == 'f' {
				// [sdf]
				return res
			} else if next >= '0' && next <= '9' {
				// [0-9]+[sdf]
				// [0-9]+$[sdf]
				i := 2
				dollar := false
				for {
					if l.i+i >= len(l.src) {
						return ""
					}
					r := l.src[l.i+i]
					if dollar && (r == 's' || r == 'd' || r == 'f') {
						res += string(r)
						return res
					} else if dollar {
						return ""
					}
					if r >= '0' && r <= '9' {
						res += string(r)
					} else if r == '$' {
						res += string(r)
						dollar = true
					} else if r == 's' || r == 'd' || r == 'f' {
						res += string(r)
						return res
					} else {
						return ""
					}
					i++
				}
			}
		}
	}
	return ""
}

func (l *Lexer) getEmoji() string {
	if l.peek() == ':' {
		i := 1
		emojis := make([]string, len(Emojis))
		copy(emojis, Emojis)
		for {
			if l.i+i >= len(l.src) || len(emojis) == 0 {
				return ""
			}
			j := 0
			r := l.src[l.i+i]
			for {
				if j >= len(emojis) {
					break
				}
				emoji := emojis[j]
				runes := []rune(emoji)
				if i >= len(runes) || runes[i] != r {
					emojis = append(emojis[:j], emojis[j+1:]...)
					continue
				}
				if runes[i] == r && i == len(runes)-1 {
					return emoji
				}
				j++
			}
			i++
		}
	}
	return ""
}

func (l *Lexer) reset() {
	l.i = 0
	l.line = 0
	l.col = 0
	l.state = StateLineStart
}

func (l *Lexer) SetInput(input string) {
	l.src = []rune(input)
	l.reset()
}

func (l *Lexer) Next() iter.Seq[Token] {
	l.reset()
	return func(yield func(Token) bool) {
		for !l.eof() {
			r := l.peek()
			if r == '\n' {
				start := l.pos()
				l.advance()
				if !yield(l.emit(TokenNewline, start, "\n")) {
					return
				}
				l.state = StateLineStart
				continue
			}
			switch l.state {
			case StateLineStart:
				if r == ' ' || r == '\t' {
					l.collectWhile(func(r rune) bool {
						return r == ' ' || r == '\t'
					})
				}
				start := l.pos()
				if l.hasPrefix("##") {
					comment := l.collectWhile(func(r rune) bool {
						return r != '\n'
					})
					if !yield(l.emit(TokenComment, start, comment)) {
						return
					}
					continue
				}
				key := l.collectWhile(func(r rune) bool {
					return r != '=' && r != '\n' && r != '\t'
				})
				if key != "" {
					if !yield(l.emit(TokenKey, start, key)) {
						return
					}
				}
				r = l.peek()
				if r == '=' {
					start := l.pos()
					l.advance()
					if !yield(l.emit(TokenAssign, start, "=")) {
						return
					}
					l.state = StateValue
				}
			case StateValue:
				switch r {
				case '\t':
					l.advance() // skip tab
					start := l.pos()
					comment := l.collectWhile(func(r rune) bool {
						return r != '\n'
					})
					if !yield(l.emit(TokenComment, start, comment)) {
						return
					}
				default:
					start := l.pos()
					if l.hasPrefix(LineBreak) {
						l.advance(len(LineBreak))
						if !yield(l.emit(TokenLineBreak, start, LineBreak)) {
							return
						}
					} else if code := l.getFormatCode(); code != "" {
						l.advance(2)
						if !yield(l.emit(TokenFormatCode, start, code)) {
							return
						}
					} else if spec := l.getFormatSpecifier(); spec != "" {
						l.advance(len(spec))
						if !yield(l.emit(TokenFormatSpecifier, start, spec)) {
							return
						}
					} else if emoji := l.getEmoji(); emoji != "" {
						l.advance(len(emoji))
						if !yield(l.emit(TokenEmoji, start, emoji)) {
							return
						}
					} else {
						value := l.collectWhile(func(r rune) bool {
							return r != '\n' && r != '\t' && !l.hasPrefix(LineBreak) && l.getFormatCode() == "" && l.getFormatSpecifier() == "" && l.getEmoji() == ""
						})
						if value != "" {
							if !yield(l.emit(TokenText, start, value)) {
								return
							}
						}
					}
				}
			}
		}
	}
}
