package handlers

import (
	"log"
	"slices"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

type jsonPath struct {
	isKey bool
	path  []string
}

func matchKey(path string) jsonPath {
	return jsonPath{isKey: true, path: strings.Split(path, "/")}
}

func matchValue(path string) jsonPath {
	return jsonPath{isKey: false, path: strings.Split(path, "/")}
}

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
	Matcher []jsonPath
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

func (j *jsonHandler) GetPattern(project *core.Project) string {
	return j.pattern.Resolve(project)
}

func (j *jsonHandler) GetActions(document *textdocument.TextDocument, position protocol.Position) *HandlerActions {
	location := jsonc.GetLocation(document.GetText(), document.OffsetAt(position))
	entry := j.findEntry(location)
	if entry == nil {
		log.Println("entry not found:", location.Path)
		return nil
	}

	node := location.PreviousNode
	params := jsonParams{
		URI:      document.URI,
		Location: location,
	}
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
						Start: document.PositionAt(node.Offset + 1),
						End:   document.PositionAt(node.Offset + node.Length - 1),
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

func (j *jsonHandler) findEntry(location *jsonc.Location) *jsonHandlerEntry {
	for _, entry := range j.entries {
		for _, matcher := range entry.Matcher {
			if matcher.isKey == location.IsAtPropertyKey && location.Path.Matches(matcher.path) {
				return &entry
			}
		}
	}
	return nil
}
