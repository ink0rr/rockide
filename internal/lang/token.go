package lang

type TokenKind uint8

const (
	TokenKey TokenKind = iota
	TokenAssign
	TokenText
	TokenFormatSpecifier // %s %2d %.3f etc.
	TokenFormatCode      // §a, §b, etc.
	TokenLineBreak       // ~LINEBREAK~
	TokenEmoji           // :_input_key.jump:, :wood_pickaxe:, etc.
	TokenComment
	TokenNewline
)

type Token struct {
	Kind   TokenKind
	Offset uint32
	Start  Position
	End    Position
	Value  string
}

type Position struct {
	Line   uint32
	Column uint32
}

func (t Token) Length() uint32 {
	return t.End.Column - t.Start.Column
}
