package textdocument

const (
	lineFeed       = 10 // The \n character
	carriageReturn = 13 // The \r character
)

func isEOL(ch byte) bool {
	return ch == lineFeed || ch == carriageReturn
}

func computeLineOffsets(text string, isAtLineStart bool, textOffset uint32) []uint32 {
	result := []uint32{}
	if isAtLineStart {
		result = []uint32{textOffset}
	}
	textLength := uint32(len(text))
	for i := uint32(0); i < textLength; i++ {
		ch := text[i]
		if isEOL(ch) {
			if ch == carriageReturn && i+1 < textLength && text[i+1] == lineFeed {
				i++
			}
			result = append(result, textOffset+i+1)
		}
	}
	return result
}
