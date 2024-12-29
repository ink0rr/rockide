package textdocument

import (
	"math"
	"os"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

type TextDocument struct {
	URI     uri.URI `json:"uri"`
	content string
}

func New(uri uri.URI) (*TextDocument, error) {
	txt, err := os.ReadFile(uri.Filename())
	if err != nil {
		return nil, err
	}
	document := TextDocument{}
	document.URI = uri
	document.content = string(txt)
	return &document, nil
}

func (d *TextDocument) ensureBeforeEOL(offset uint32, lineOffset uint32) uint32 {
	for offset > lineOffset && isEOL(d.content[offset-1]) {
		offset--
	}
	return offset
}

func (d *TextDocument) PositionAt(offset uint32) protocol.Position {
	if length := uint32(len(d.content)); offset > length {
		offset = length
	} else if offset < 0 {
		offset = 0
	}
	lineOffsets := computeLineOffsets(d.content, true, 0)
	low := 0
	high := len(lineOffsets)
	if high == 0 {
		return protocol.Position{Character: offset}
	}
	for low < high {
		mid := int(math.Floor(float64((low + high) / 2)))
		if lineOffsets[mid] > offset {
			high = mid
		} else {
			low = mid + 1
		}
	}
	// low is the least x for which the line offset is larger than the current offset
	// or array.length if no line offset is larger than the current offset
	line := low - 1
	offset = d.ensureBeforeEOL(offset, lineOffsets[line])
	return protocol.Position{Line: uint32(line), Character: offset - lineOffsets[line]}
}

func (d *TextDocument) OffsetAt(position protocol.Position) uint32 {
	lineOffsets := computeLineOffsets(d.content, true, 0)
	maxLine := uint32(len(lineOffsets))
	contentLength := uint32(len(d.content))
	if position.Line >= maxLine {
		return contentLength
	}
	if position.Line < 0 {
		return 0
	}
	lineOffset := lineOffsets[position.Line]
	if position.Character <= 0 {
		return lineOffset
	}
	var nextLineOffset uint32
	if position.Line+1 < maxLine {
		nextLineOffset = lineOffsets[position.Line+1]
	} else {
		nextLineOffset = contentLength
	}
	offset := min(lineOffset+position.Character, nextLineOffset)
	return d.ensureBeforeEOL(offset, lineOffset)
}

func (d *TextDocument) GetText() string {
	return d.content
}
