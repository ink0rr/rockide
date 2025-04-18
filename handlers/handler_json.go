package handlers

import (
	"log"
	"slices"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

type jsonHandlerActions uint16

const (
	completions jsonHandlerActions = 1 << iota
	definitions
	rename
)

func (a jsonHandlerActions) Has(action jsonHandlerActions) bool {
	return a&action != 0
}

type jsonParams struct {
	URI      protocol.DocumentURI
	Location *jsonc.Location
}

func (j *jsonParams) getParentNode() *jsonc.Node {
	document := textdocument.Get(j.URI)
	root, _ := jsonc.ParseTree(document.GetText(), nil)
	path := j.Location.Path
	return jsonc.FindNodeAtLocation(root, path[:len(path)-1])
}

type jsonHandlerEntry struct {
	Path    []shared.JsonPath
	Matcher func(params *jsonParams) bool
	Actions jsonHandlerActions
	// Filter completions to only show undeclared reference
	FilterDiff bool
	// Source for completions and definitions
	Source func(params *jsonParams) []core.Reference
	// References that uses the same source
	References func(params *jsonParams) []core.Reference
}

type jsonHandler struct {
	pattern shared.Pattern
	entries []jsonHandlerEntry
}

func newJsonHandler(pattern shared.Pattern, entries []jsonHandlerEntry) *jsonHandler {
	return &jsonHandler{pattern, entries}
}

func (j *jsonHandler) Pattern() string {
	return j.pattern.ToString()
}

func (j *jsonHandler) GetActions(document *textdocument.TextDocument, position protocol.Position) *HandlerActions {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	if shared.IsMolangLocation(location) {
		return Molang.GetActions(document, offset, location)
	}
	params := jsonParams{
		URI:      document.URI,
		Location: location,
	}
	entry := j.findEntry(&params)
	if entry == nil {
		log.Println("Entry not found:", location.Path)
		return nil
	}

	node := location.PreviousNode
	actions := HandlerActions{}

	if entry.Actions.Has(completions) {
		actions.Completions = func() []protocol.CompletionItem {
			res := []protocol.CompletionItem{}
			set := make(map[string]bool)
			var items []core.Reference
			if entry.FilterDiff {
				items = stores.Difference(entry.Source(&params), entry.References(&params))
			} else {
				items = entry.Source(&params)
			}
			for _, item := range items {
				if set[item.Value] {
					continue
				}
				set[item.Value] = true
				value := `"` + item.Value + `"`
				completion := protocol.CompletionItem{
					Label: value,
				}
				if node != nil {
					completion.TextEdit = &protocol.Or_CompletionItem_textEdit{
						Value: protocol.TextEdit{
							Range: protocol.Range{
								Start: document.PositionAt(node.Offset),
								End:   document.PositionAt(node.Offset + node.Length),
							},
							NewText: value,
						},
					}
				}
				res = append(res, completion)
			}
			return res
		}
	}

	if node != nil && entry.Actions.Has(definitions) {
		actions.Definitions = func() []protocol.LocationLink {
			res := []protocol.LocationLink{}
			for _, item := range entry.Source(&params) {
				if item.Value != node.Value {
					continue
				}
				location := protocol.LocationLink{
					OriginSelectionRange: &protocol.Range{
						Start: document.PositionAt(node.Offset),
						End:   document.PositionAt(node.Offset + node.Length),
					},
					TargetURI: item.URI,
				}
				if item.Range != nil {
					location.TargetRange = *item.Range
					location.TargetSelectionRange = *item.Range
				}
				res = append(res, location)
			}
			return res
		}
	}

	if node != nil && entry.Actions.Has(rename) {
		actions.Rename = func(newName string) *protocol.WorkspaceEdit {
			changes := make(map[protocol.DocumentURI][]protocol.TextEdit)
			for _, item := range slices.Concat(entry.Source(&params), entry.References(&params)) {
				if item.Value != node.Value {
					continue
				}
				edit := protocol.TextEdit{
					NewText: newName,
					Range:   *item.Range,
				}
				// Exclude quotation marks
				edit.Range.Start.Character++
				edit.Range.End.Character--

				changes[item.URI] = append(changes[item.URI], edit)
			}
			return &protocol.WorkspaceEdit{Changes: changes}
		}
	}
	return &actions
}

func (j *jsonHandler) findEntry(params *jsonParams) *jsonHandlerEntry {
	for _, entry := range j.entries {
		for _, jsonPath := range entry.Path {
			if jsonPath.IsKey == params.Location.IsAtPropertyKey && params.Location.Path.Matches(jsonPath.Path) {
				if entry.Matcher == nil || entry.Matcher(params) {
					return &entry
				}
			}
		}
	}
	return nil
}
