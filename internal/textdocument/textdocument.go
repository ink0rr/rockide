package textdocument

import (
	"os"
	"sync"

	"github.com/ink0rr/rockide/internal/protocol"
)

var (
	documents    = make(map[protocol.DocumentURI]*TextDocument)
	cacheEnabled = false
	mutex        sync.Mutex
)

type TextDocument struct {
	URI         protocol.DocumentURI `json:"uri"`
	content     string
	lineOffsets []uint32
}

func Open(uri protocol.DocumentURI) (*TextDocument, error) {
	if cacheEnabled {
		mutex.Lock()
		defer mutex.Unlock()
		if document := documents[uri]; document != nil {
			return document, nil
		}
	}
	txt, err := os.ReadFile(uri.Path())
	if err != nil {
		return nil, err
	}
	document := TextDocument{URI: uri, content: string(txt)}
	if cacheEnabled {
		documents[uri] = &document
	}
	return &document, nil
}

func Update(uri protocol.DocumentURI, contentChanges []protocol.TextDocumentContentChangeEvent) {
	if len(contentChanges) == 0 {
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	document := documents[uri]
	if document == nil {
		return
	}
	for _, change := range contentChanges {
		startOffset := document.OffsetAt(change.Range.Start)
		endOffset := document.OffsetAt(change.Range.End)
		document.content = document.content[:startOffset] + change.Text + document.content[endOffset:]
		document.lineOffsets = nil
	}
}

func UpdateFull(uri protocol.DocumentURI, text *string) {
	if text == nil {
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	document := documents[uri]
	if document == nil || document.content == *text {
		return
	}
	document.content = *text
	document.lineOffsets = nil
}

func Close(uri protocol.DocumentURI) {
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

func (d *TextDocument) OffsetAt(position protocol.Position) uint32 {
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
