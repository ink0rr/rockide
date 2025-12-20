package lang

const (
	LineBreak   string = "~LINEBREAK~"
	SectionSign rune   = '§'
)

var (
	lineBreakRunes = []rune(LineBreak)
	commentRunes   = []rune("##")
)
