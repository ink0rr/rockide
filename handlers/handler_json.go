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

type jsonHandlerActions int

const (
	completions jsonHandlerActions = 1 << iota
	definitions
	rename
)

func (a jsonHandlerActions) Has(action jsonHandlerActions) bool {
	return a&action != 0
}

type jsonParams struct {
	URI      protocol.URI
	Node     *jsonc.Node
	Location *jsonc.Location
}

type jsonHandlerEntry struct {
	Path       []string
	MatchType  string
	Actions    jsonHandlerActions
	Source     func(params *jsonParams) []core.Reference
	References func(params *jsonParams) []core.Reference

	// Used to cache Path splitted by '/'
	jsonPath [][]string
}

func (j *jsonHandlerEntry) getJsonPath() [][]string {
	if j.jsonPath == nil {
		for _, path := range j.Path {
			j.jsonPath = append(j.jsonPath, strings.Split(path, "/"))
		}
	}
	return j.jsonPath
}

type jsonHandler struct {
	pattern string
	entries []jsonHandlerEntry
}

func (j *jsonHandler) GetPattern() string {
	return j.pattern
}

func (j *jsonHandler) GetActions(document *textdocument.TextDocument, position *protocol.Position) *HandlerActions {
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

	params := jsonParams{
		URI:      document.URI,
		Node:     location.PreviousNode,
		Location: location,
	}
	actions := HandlerActions{}

	if entry.Actions.Has(completions) {
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

	if entry.Actions.Has(definitions) {
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

	if entry.Actions.Has(rename) {
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

func (j *jsonHandler) findEntry(location *jsonc.Location) *jsonHandlerEntry {
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
