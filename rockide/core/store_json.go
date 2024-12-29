package core

import (
	"github.com/ink0rr/go-jsonc"
	"github.com/ink0rr/rockide/textdocument"
	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

type TransformResult struct {
	Value string
	Skip  bool
}

type JsonStoreEntry struct {
	Id             string
	Path           []jsonc.Path
	TransformValue func(node *jsonc.Node) TransformResult
}

type JsonStore struct {
	pattern string
	entries []JsonStoreEntry
	store   map[string][]Reference
}

// GetPattern implements Store.
func (j *JsonStore) GetPattern() string {
	return j.pattern
}

// Parse implements Store.
func (j *JsonStore) Parse(uri uri.URI) error {
	document, err := textdocument.New(uri)
	if err != nil {
		return err
	}
	root, _ := jsonc.ParseTree(document.GetText(), nil)
	for _, entry := range j.entries {
		data := j.store[entry.Id]
		for _, path := range entry.Path {
			// If the last path is "*", we want to grab the values instead of keys
			index := 0
			if path[len(path)-1] == "*" {
				index = 1
			}
			var extract func(node *jsonc.Node)
			extract = func(node *jsonc.Node) {
				if node.Type == jsonc.NodeTypeString {
					value := node.Value
					if entry.TransformValue != nil {
						result := entry.TransformValue(node)
						if result.Skip {
							return
						}
						value = result.Value
					}

					data = append(data, Reference{
						Value: value.(string),
						Uri:   uri,
						Range: &protocol.Range{
							Start: document.PositionAt(uint32(node.Offset)),
							End:   document.PositionAt(uint32(node.Offset) + uint32(node.Length)),
						}})
				} else if node.Type == jsonc.NodeTypeProperty && node.Children != nil && index < len(node.Children) {
					// TODO
				} else if node.Children != nil {
					for _, child := range node.Children {
						extract(child)
					}
				}
			}
		}
	}
	_ = root
	return nil
}

// Get implements Store.
func (j *JsonStore) Get(key string) []Reference {
	panic("unimplemented")
}

// GetFrom implements Store.
func (j *JsonStore) GetFrom(uri uri.URI, key Store) []Reference {
	panic("unimplemented")
}

// Delete implements Store.
func (j *JsonStore) Delete(uri uri.URI) {
	panic("unimplemented")
}

var _ Store = (*JsonStore)(nil)

func NewJsonStore(pattern string, entries []JsonStoreEntry) *JsonStore {
	store := make(map[string][]Reference)
	return &JsonStore{store: store, entries: entries}
}
