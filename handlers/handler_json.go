package handlers

import (
	"path/filepath"
	"slices"
	"strings"
	"sync"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/protocol/semtok"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/ink0rr/rockide/shared"
)

const defaultScope = "__default__"

type JsonContext struct {
	URI           protocol.DocumentURI
	NodeValue     string
	GetPath       func() jsonc.Path
	GetParentNode func() *jsonc.Node
}

type JsonEntry struct {
	Id            string
	Path          []shared.JsonPath
	Matcher       func(ctx *JsonContext) bool
	Transform     func(value string) string
	DisableRename bool
	// Filter completions to only show undeclared reference
	FilterDiff bool
	// Function to set the scope key. If omitted will use the `defaultScope` value
	ScopeKey func(ctx *JsonContext) string
	// Source for completions and definitions
	Source func(ctx *JsonContext) []core.Symbol
	// References that uses the same source
	References  func(ctx *JsonContext) []core.Symbol
	VanillaData mapset.Set[string]
}

type JsonHandler struct {
	Pattern                 shared.Pattern
	SavePath                bool
	Entries                 []JsonEntry
	MolangLocations         []shared.JsonPath
	MolangSemanticLocations []shared.JsonPath
	storeList               map[string]map[string][]core.Symbol
	mu                      sync.Mutex
}

func (j *JsonHandler) GetPattern() string {
	return j.Pattern.ToString()
}

func (j *JsonHandler) Parse(uri protocol.DocumentURI) error {
	j.mu.Lock()
	defer j.mu.Unlock()
	if j.storeList == nil {
		j.storeList = make(map[string]map[string][]core.Symbol)
	}
	if j.SavePath {
		j.parsePath(uri)
	}
	document, err := textdocument.GetOrReadFile(uri)
	if err != nil {
		return err
	}
	root, _ := jsonc.ParseTree(document.GetText(), nil)
	for _, entry := range j.Entries {
		if j.storeList[entry.Id] == nil {
			j.storeList[entry.Id] = make(map[string][]core.Symbol)
		}
		for _, jsonPath := range entry.Path {
			for _, node := range jsonPath.GetNodes(root) {
				nodeValue, ok := node.Value.(string)
				if !ok {
					continue
				}
				ctx := JsonContext{
					URI:       uri,
					NodeValue: nodeValue,
					GetPath: func() jsonc.Path {
						return jsonc.GetNodePath(node)
					},
					GetParentNode: func() *jsonc.Node {
						path := jsonc.GetNodePath(node)
						return jsonc.FindNodeAtLocation(root, path[:len(path)-1])
					},
				}
				if entry.Matcher != nil && !entry.Matcher(&ctx) {
					continue
				}
				if entry.Transform != nil {
					nodeValue = entry.Transform(nodeValue)
				}
				var scope string
				if entry.ScopeKey != nil {
					scope = entry.ScopeKey(&ctx)
				}
				j.storeList[entry.Id][scope] = append(j.storeList[entry.Id][scope], core.Symbol{
					Value: nodeValue,
					URI:   uri,
					Range: &protocol.Range{
						Start: document.PositionAt(node.Offset),
						End:   document.PositionAt(node.Offset + node.Length),
					},
				})
			}
		}
	}
	return nil
}

func (j *JsonHandler) parsePath(uri protocol.DocumentURI) {
	path, err := filepath.Rel(shared.Getwd(), uri.Path())
	if err != nil {
		panic(err)
	}
	packType := j.Pattern.PackType()
	path = filepath.ToSlash(path)
	_, path, found := strings.Cut(path, packType+"/")
	if !found {
		panic("invalid project path")
	}
	if j.storeList["path"] == nil {
		j.storeList["path"] = make(map[string][]core.Symbol)
	}
	j.storeList["path"][defaultScope] = append(j.storeList["path"][defaultScope], core.Symbol{Value: path, URI: uri})
}

func (j *JsonHandler) Get(id string, scopes ...string) []core.Symbol {
	j.mu.Lock()
	defer j.mu.Unlock()

	res := []core.Symbol{}
	if len(scopes) > 0 {
		for _, scope := range scopes {
			res = append(res, j.storeList[id][scope]...)
		}
	} else {
		for _, symbols := range j.storeList[id] {
			res = append(res, symbols...)
		}
	}
	return res
}

func (j *JsonHandler) GetFrom(uri protocol.DocumentURI, id string, scopes ...string) []core.Symbol {
	res := []core.Symbol{}
	for _, symbol := range j.Get(id, scopes...) {
		if symbol.URI == uri {
			res = append(res, symbol)
		}
	}
	return res
}

func (j *JsonHandler) Delete(uri protocol.DocumentURI) {
	j.mu.Lock()
	defer j.mu.Unlock()
	for _, store := range j.storeList {
		for scope, symbols := range store {
			store[scope] = slices.DeleteFunc(symbols, func(symbol core.Symbol) bool {
				return symbol.URI == uri
			})
		}
	}
}

func (j *JsonHandler) prepareContext(document *textdocument.TextDocument, location *jsonc.Location) (*JsonEntry, *JsonContext) {
	var nodeValue string
	if node := location.PreviousNode; node != nil {
		nodeValue, _ = node.Value.(string)
	}
	params := JsonContext{
		URI:       document.URI,
		NodeValue: nodeValue,
		GetPath: func() jsonc.Path {
			return location.Path
		},
		GetParentNode: func() *jsonc.Node {
			root, _ := jsonc.ParseTree(document.GetText(), nil)
			path := location.Path
			return jsonc.FindNodeAtLocation(root, path[:len(path)-1])
		},
	}
	for _, entry := range j.Entries {
		for _, jsonPath := range entry.Path {
			if jsonPath.IsKey == location.IsAtPropertyKey && location.Path.Matches(jsonPath.Path) {
				if entry.Matcher == nil || entry.Matcher(&params) {
					return &entry, &params
				}
			}
		}
	}
	return nil, nil
}

func (j *JsonHandler) isMolangLocation(location *jsonc.Location) bool {
	if location.IsAtPropertyKey {
		return false
	}
	if j.MolangLocations != nil {
		for _, jsonPath := range j.MolangLocations {
			if location.Path.Matches(jsonPath.Path) {
				return true
			}
		}
	}
	return false
}

func (j *JsonHandler) isMolangSemanticLocation(location *jsonc.Location) bool {
	if location.IsAtPropertyKey {
		return false
	}
	if j.MolangSemanticLocations != nil {
		for _, jsonPath := range j.MolangSemanticLocations {
			if location.Path.Matches(jsonPath.Path) {
				return true
			}
		}
	}
	return j.isMolangLocation(location)
}

func (j *JsonHandler) Completions(document *textdocument.TextDocument, position protocol.Position) []protocol.CompletionItem {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	node := location.PreviousNode

	res := []protocol.CompletionItem{}
	if j.isMolangLocation(location) {
		if ctx := NewMolangContext(document, location, offset); ctx != nil {
			res = MolangCompletions(ctx)
		}
	}
	entry, ctx := j.prepareContext(document, location)
	if entry == nil || entry.Source == nil || entry.References == nil {
		return res
	}

	var items []core.Symbol
	if entry.FilterDiff {
		items = difference(entry.Source(ctx), entry.References(ctx))
	} else {
		items = entry.Source(ctx)
	}

	set := mapset.NewThreadUnsafeSet[string]()
	if entry.VanillaData != nil {
		set = entry.VanillaData.Clone()
	}

	for _, item := range items {
		if set.ContainsOne(item.Value) {
			continue
		}
		set.Add(item.Value)
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

func (j *JsonHandler) Definitions(document *textdocument.TextDocument, position protocol.Position) []protocol.LocationLink {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	node := location.PreviousNode
	if node == nil {
		return nil
	}

	res := []protocol.LocationLink{}
	if j.isMolangLocation(location) {
		if ctx := NewMolangContext(document, location, offset); ctx != nil {
			res = MolangDefinitions(ctx)
		}
	}
	entry, ctx := j.prepareContext(document, location)
	if entry == nil || entry.Source == nil || entry.References == nil {
		return res
	}

	nodeValue, ok := node.Value.(string)
	if !ok {
		return res
	}
	if entry.Transform != nil {
		nodeValue = entry.Transform(nodeValue)
	}

	for _, item := range entry.Source(ctx) {
		if item.Value != nodeValue {
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

func (j *JsonHandler) PrepareRename(document *textdocument.TextDocument, position protocol.Position) *protocol.PrepareRenamePlaceholder {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	node := location.PreviousNode
	if node == nil {
		return nil
	}
	entry, _ := j.prepareContext(document, location)
	if entry == nil || entry.Source == nil || entry.References == nil || entry.DisableRename {
		return nil
	}
	// TODO: Support renaming entry that uses transform
	if entry.Transform != nil {
		return nil
	}

	start := node.Offset + 1
	end := start + node.Length - 2
	return &protocol.PrepareRenamePlaceholder{
		Range: protocol.Range{
			Start: document.PositionAt(start),
			End:   document.PositionAt(end),
		},
		Placeholder: document.GetText()[start:end],
	}
}

func (j *JsonHandler) Rename(document *textdocument.TextDocument, position protocol.Position, newName string) *protocol.WorkspaceEdit {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	node := location.PreviousNode
	if node == nil {
		return nil
	}
	entry, ctx := j.prepareContext(document, location)
	if entry == nil || entry.Source == nil || entry.References == nil || entry.DisableRename {
		return nil
	}

	changes := make(map[protocol.DocumentURI][]protocol.TextEdit)
	for _, item := range slices.Concat(entry.Source(ctx), entry.References(ctx)) {
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

func (j *JsonHandler) Hover(document *textdocument.TextDocument, position protocol.Position) *protocol.Hover {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	if j.isMolangLocation(location) {
		if ctx := NewMolangContext(document, location, offset); ctx != nil {
			return MolangHover(ctx)
		}
	}
	return nil
}

func (j *JsonHandler) SignatureHelp(document *textdocument.TextDocument, position protocol.Position) *protocol.SignatureHelp {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	if j.isMolangLocation(location) {
		if ctx := NewMolangContext(document, location, offset); ctx != nil {
			return MolangSignatureHelp(ctx)
		}
	}
	return nil
}

func (j *JsonHandler) SemanticTokens(document *textdocument.TextDocument) *protocol.SemanticTokens {
	tokens := []semtok.Token{}

	jsonc.Visit(document.GetText(), &jsonc.Visitor{
		OnLiteralValue: func(value any, offset, length, startLine, startCharacter uint32, pathSupplier func() jsonc.Path) {
			text, ok := value.(string)
			if !ok || text == "" || text[0] == '@' || text[0] == '/' {
				return
			}
			location := jsonc.Location{Path: pathSupplier()}
			if j.isMolangSemanticLocation(&location) {
				tokens = append(tokens, MolangSemanticTokens(text, startLine, startCharacter)...)
			}
		},
	}, nil)

	return &protocol.SemanticTokens{
		Data: semtok.Encode(tokens, tokenType, tokenModifier),
	}
}
