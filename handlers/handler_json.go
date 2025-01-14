package handlers

import (
	"log"
	"slices"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/ink0rr/rockide/stores"
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
	URI      protocol.DocumentURI
	Node     *jsonc.Node
	Location *jsonc.Location
}

func (j *jsonParams) getParentNode() *jsonc.Node {
	document, err := textdocument.Open(j.URI)
	if err != nil {
		return nil
	}
	root, _ := jsonc.ParseTree(document.GetText(), nil)
	path := j.Location.Path
	return jsonc.FindNodeAtLocation(root, path[:len(path)-1])
}

type jsonHandlerEntry struct {
	Path      []string
	MatchType string
	Actions   jsonHandlerActions
	// Filter completions to only show undeclared reference
	FilterDiff bool
	// Source for completions and definitions
	Source func(params *jsonParams) []core.Reference
	// References that uses the same source
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
	pattern core.Pattern
	entries []jsonHandlerEntry
}

func newJsonHandler(pattern core.Pattern, entries []jsonHandlerEntry) *jsonHandler {
	return &jsonHandler{pattern, entries}
}

func (j *jsonHandler) GetPattern(project *core.Project) string {
	return j.pattern.Resolve(project)
}

func (j *jsonHandler) GetActions(document *textdocument.TextDocument, position *protocol.Position) *HandlerActions {
	location := jsonc.GetLocation(document.GetText(), document.OffsetAt(position))
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
				if params.Node != nil {
					completion.TextEdit = &protocol.Or_CompletionItem_textEdit{
						Value: protocol.TextEdit{
							Range: protocol.Range{
								Start: document.PositionAt(params.Node.Offset),
								End:   document.PositionAt(params.Node.Offset + params.Node.Length),
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

	if params.Node != nil && entry.Actions.Has(definitions) {
		actions.Definitions = func() []protocol.LocationLink {
			res := []protocol.LocationLink{}
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
			return res
		}
	}

	if params.Node != nil && entry.Actions.Has(rename) {
		actions.Rename = func(newName string) *protocol.WorkspaceEdit {
			changes := make(map[protocol.DocumentURI][]protocol.TextEdit)
			for _, item := range slices.Concat(entry.Source(&params), entry.References(&params)) {
				if item.Value != params.Node.Value {
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
