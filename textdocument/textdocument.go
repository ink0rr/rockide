package textdocument

import (
	"os"
	"sync"

	"github.com/rockide/protocol"
	"go.lsp.dev/uri"
)

var documents = make(map[protocol.URI]*TextDocument)
var cacheEnabled = false
var mutex sync.Mutex

type TextDocument struct {
	URI         uri.URI `json:"uri"`
	content     string
	lineOffsets []uint32
}

func Open(uri uri.URI) (*TextDocument, error) {
	if cacheEnabled {
		mutex.Lock()
		defer mutex.Unlock()
		if document := documents[uri]; document != nil {
			return document, nil
		}
	}
	txt, err := os.ReadFile(uri.Filename())
	if err != nil {
		return nil, err
	}
	document := TextDocument{URI: uri, content: string(txt)}
	if cacheEnabled {
		documents[uri] = &document
	}
	return &document, nil
}

func Update(uri protocol.URI, content string) {
	mutex.Lock()
	defer mutex.Unlock()
	if document := documents[uri]; document != nil {
		document.content = content
	}
}

func Close(uri protocol.URI) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(documents, uri)
}

func EnableCache(flag bool) {
	cacheEnabled = flag
}

func (d *TextDocument) ensureBeforeEOL(offset uint32, lineOffset uint32) uint32 {
	for offset > lineOffset && isEOL(d.content[offset-1]) {
		offset--
	}
	return offset
}

func (d *TextDocument) getLineOffsets() []uint32 {
	if d.lineOffsets == nil {
		d.lineOffsets = computeLineOffsets(d.content, true, 0)
	}
	return d.lineOffsets
}

func (d *TextDocument) PositionAt(offset uint32) protocol.Position {
	offset = min(offset, uint32(len(d.content)))
	lineOffsets := d.getLineOffsets()
	low := 0
	high := len(lineOffsets)
	if high == 0 {
		return protocol.Position{Character: offset}
	}
	for low < high {
		mid := (low + high) / 2
		if lineOffsets[mid] > offset {
			high = mid
		} else {
			low = mid + 1
		}
	}
	// low is the least x for which the line offset is larger than the current offset
	// or array.length if no line offset is larger than the current offset
	line := low - 1
	if low == 0 {
		line = 0
	}
	offset = d.ensureBeforeEOL(offset, lineOffsets[line])
	return protocol.Position{Line: uint32(line), Character: offset - lineOffsets[line]}
}

func (d *TextDocument) OffsetAt(position *protocol.Position) uint32 {
	lineOffsets := d.getLineOffsets()
	maxLine := uint32(len(lineOffsets))
	contentLength := uint32(len(d.content))
	if position.Line >= maxLine {
		return contentLength
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
