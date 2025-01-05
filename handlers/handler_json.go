package handlers

import (
	"log"
	"slices"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/jsonc"
	"github.com/ink0rr/rockide/stores"
	"github.com/ink0rr/rockide/textdocument"
	"github.com/rockide/protocol"
)

type JsonHandlerActions int

const (
	Completions JsonHandlerActions = 1 << iota
	Definitions
	Rename
)

func (a JsonHandlerActions) Has(action JsonHandlerActions) bool {
	return a&action != 0
}

type JsonParams struct {
	URI      protocol.URI
	Node     *jsonc.Node
	Location *jsonc.Location
}

type JsonHandlerEntry struct {
	Path []string
	// Used to cache path splitted by '/'
	jsonPath   []jsonc.Path
	MatchType  string
	Actions    JsonHandlerActions
	Source     func(params *JsonParams) []core.Reference
	References func(params *JsonParams) []core.Reference
}

func (j *JsonHandlerEntry) getJsonPath() []jsonc.Path {
	if j.jsonPath == nil {
		j.jsonPath = make([]jsonc.Path, len(j.Path))
		for i, path := range j.Path {
			for _, segment := range strings.Split(path, "/") {
				j.jsonPath[i] = append(j.jsonPath[i], segment)
			}
		}
	}
	return j.jsonPath
}

type JsonHandler struct {
	pattern string
	entries []JsonHandlerEntry
}

func (j *JsonHandler) GetPattern() string {
	return j.pattern
}

func (j *JsonHandler) GetActions(document *textdocument.TextDocument, position *protocol.Position) *HandlerActions {
	location := jsonc.GetLocation(document.GetText(), document.OffsetAt(position))
	if location.PreviousNode == nil {
		log.Println("cannot find parent node", location.Path)
		return nil
	}

	entry := j.findEntry(location)
	if entry == nil {
		log.Println("entry not found:", location.Path)
		return nil
	}

	params := JsonParams{
		URI:      document.URI,
		Node:     location.PreviousNode,
		Location: location,
	}
	actions := HandlerActions{}

	if entry.Actions.Has(Completions) {
		actions.Completions = func() []protocol.CompletionItem {
			res := []protocol.CompletionItem{}
			set := make(map[string]bool)
			for _, item := range stores.Difference(entry.Source(&params), entry.References(&params)) {
				if set[item.Value] {
					continue
				}
				set[item.Value] = true
				res = append(res, protocol.CompletionItem{
					Label: item.Value,
					TextEdit: &protocol.TextEdit{
						Range: protocol.Range{
							Start: document.PositionAt(params.Node.Offset + 1),
							End:   document.PositionAt(params.Node.Offset + params.Node.Length - 1),
						},
						NewText: item.Value,
					},
					InsertTextFormat: protocol.InsertTextFormatPlainText,
					InsertTextMode:   protocol.InsertTextModeAdjustIndentation,
				})
			}
			return res
		}
	}

	if entry.Actions.Has(Definitions) {
		actions.Definitions = func() (res []protocol.LocationLink) {
			for _, item := range entry.Source(&params) {
				if item.Value != params.Node.Value {
					continue
				}
				location := protocol.LocationLink{
					OriginSelectionRange: &protocol.Range{
						Start: document.PositionAt(params.Node.Offset + 1),
						End:   document.PositionAt(params.Node.Offset + params.Node.Length - 1),
					},
					TargetURI: item.URI,
				}
				if item.Range != nil {
					location.TargetRange = *item.Range
					location.TargetSelectionRange = *item.Range
				}
				res = append(res, location)
			}
			return
		}
	}

	if entry.Actions.Has(Rename) {
		actions.Rename = func() (res []protocol.WorkspaceEdit) {
			for _, item := range slices.Concat(entry.Source(&params), entry.References(&params)) {
				if item.Value != params.Node.Value {
					continue
				}
				// TODO
			}
			return
		}
	}
	return &actions
}

func (j *JsonHandler) findEntry(location *jsonc.Location) *JsonHandlerEntry {
	for _, entry := range j.entries {
		if (entry.MatchType == "key" && !location.IsAtPropertyKey) ||
			(entry.MatchType == "value" && location.IsAtPropertyKey) {
			continue
		}
		for _, targetPath := range entry.getJsonPath() {
			if jsonc.PathMatches(location.Path, targetPath) {
				return &entry
			}
		}
	}
	return nil
}
