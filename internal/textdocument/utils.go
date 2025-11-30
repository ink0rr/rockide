package textdocument

const (
	lineFeed       = '\n'
	carriageReturn = '\r'
)

func isEOL(ch rune) bool {
	return ch == lineFeed || ch == carriageReturn
}

func computeLineOffsets(text []rune, isAtLineStart bool, textOffset uint32) []uint32 {
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

// count UTF-16 code units in a rune slice
func utf16Len(runes []rune) uint32 {
	var n uint32
	for _, r := range runes {
		if r >= 0x10000 {
			n += 2 // surrogate pair
		} else {
			n++
		}
	}
	return n
}

// convert a UTF-16 position (char offset) to UTF-32 rune offset within a line
func utf16ToRuneOffset(runes []rune, utf16Pos uint32) uint32 {
	var ru, u16 uint32
	for ru < uint32(len(runes)) && u16 < utf16Pos {
		if runes[ru] >= 0x10000 {
			u16 += 2
		} else {
			u16++
		}
		ru++
	}
	return ru
}
