package lang

type State uint8

const (
	StateLineStart State = iota
	StateKey
	StateValue
)
