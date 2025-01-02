package handlers

import (
	"strings"

	"github.com/ink0rr/rockide/jsonc"
	"github.com/ink0rr/rockide/rockide/core"
	"github.com/ink0rr/rockide/textdocument"
	"github.com/rockide/protocol"
	"go.lsp.dev/uri"
)

type JsonHandlerEntry struct {
	Path        []string
	jsonPath    [][]string
	MatchType   string
	Completions func(params *JsonHandlerParams) []core.Reference
	Definitions func(params *JsonHandlerParams) []core.Reference
	Rename      func(params *JsonHandlerParams) []core.Reference
}

type JsonHandler struct {
	pattern string
	entries []JsonHandlerEntry
}

func NewJsonHandler(pattern string, entries []JsonHandlerEntry) *JsonHandler {
	res := JsonHandler{pattern, entries}
	for i, entry := range entries {
		for _, path := range entry.Path {
			entry.jsonPath = append(entry.jsonPath, strings.Split(path, "/"))
		}
		res.entries[i] = entry
	}
	return &res
}

func (j *JsonHandler) GetPattern() string {
	return j.pattern
}

func (j *JsonHandler) FindEntry(location *jsonc.Location) *JsonHandlerEntry {
	for _, entry := range j.entries {
		if (entry.MatchType == "key" && !location.IsAtPropertyKey) ||
			(entry.MatchType == "value" && location.IsAtPropertyKey) {
			continue
		}
		for _, targetPath := range entry.jsonPath {
			if core.IsJsonPathMatch(location.Path, targetPath) {
				return &entry
			}
		}
	}
	return nil
}

type JsonHandlerParams struct {
	URI      uri.URI
	Text     string
	Location *jsonc.Location
	Node     *jsonc.Node
}

func NewJsonHandlerParams(document *textdocument.TextDocument, position *protocol.Position) *JsonHandlerParams {
	location := jsonc.GetLocation(document.GetText(), document.OffsetAt(position))
	return &JsonHandlerParams{
		URI:      document.URI,
		Text:     document.GetText(),
		Location: location,
		Node:     location.PreviousNode,
	}
}

func (j *JsonHandlerParams) GetParentNode() *jsonc.Node {
	root, _ := jsonc.ParseTree(j.Text, nil)
	path := j.Location.Path
	return jsonc.FindNodeAtLocation(root, path)
}

func (j *JsonHandlerParams) IsAtPropertyKeyOrArray() bool {
	switch j.Location.Path[len(j.Location.Path)-1].(type) {
	case int:
		return true
	}
	return j.Location.IsAtPropertyKey
}
