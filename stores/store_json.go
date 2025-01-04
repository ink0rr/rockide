package stores

import (
	"strings"
	"sync"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/jsonc"
	"github.com/ink0rr/rockide/textdocument"
	"github.com/rockide/protocol"
	"go.lsp.dev/uri"
)

type jsonStoreEntry struct {
	Id        string
	Path      []string
	jsonPath  [][]string
	Transform func(node *jsonc.Node) *string
}

type JsonStore struct {
	pattern string
	entries []jsonStoreEntry
	store   map[string][]core.Reference
	mutex   sync.Mutex
}

func newJsonStore(pattern string, entries []jsonStoreEntry) *JsonStore {
	res := JsonStore{
		pattern: pattern,
		entries: entries,
		store:   make(map[string][]core.Reference),
	}
	for i, entry := range entries {
		for _, path := range entry.Path {
			entry.jsonPath = append(entry.jsonPath, strings.Split(path, "/"))
		}
		res.entries[i] = entry
	}
	return &res
}

// GetPattern implements Store.
func (j *JsonStore) GetPattern() string {
	return j.pattern
}

// Parse implements Store.
func (j *JsonStore) Parse(uri uri.URI) error {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	document, err := textdocument.Open(uri)
	if err != nil {
		return err
	}
	root, _ := jsonc.ParseTree(document.GetText(), nil)
	for _, entry := range j.entries {
		data := j.store[entry.Id]
		for _, path := range entry.jsonPath {
			// If the last path is "*", we want to grab the values instead of keys
			// The only exception is when the path is a single "*"
			index := 0
			if len(path) > 1 && path[len(path)-1] == "*" {
				index = 1
			}
			var extract func(node *jsonc.Node)
			extract = func(node *jsonc.Node) {
				if node.Type == jsonc.NodeTypeString {
					value := node.Value
					if entry.Transform != nil {
						result := entry.Transform(node)
						if result == nil {
							return
						}
						value = *result
					}
					data = append(data, core.Reference{
						Value: value.(string),
						URI:   uri,
						Range: &protocol.Range{
							Start: document.PositionAt(node.Offset),
							End:   document.PositionAt(node.Offset + node.Length),
						},
					})
				} else if node.Type == jsonc.NodeTypeProperty && node.Children != nil && index < len(node.Children) {
					value := node.Children[index].Value
					if entry.Transform != nil {
						result := entry.Transform(node.Children[index])
						if result == nil {
							return
						}
						value = *result
					}
					targetNode := node.Children[index]
					data = append(data, core.Reference{
						Value: value.(string),
						URI:   uri,
						Range: &protocol.Range{
							Start: document.PositionAt(targetNode.Offset),
							End:   document.PositionAt(targetNode.Offset + targetNode.Length),
						},
					})
				} else if node.Children != nil {
					for _, child := range node.Children {
						extract(child)
					}
				}
			}
			for _, node := range findNodesAtPath(root, path) {
				extract(node)
			}
		}
		j.store[entry.Id] = data
	}
	return nil
}

// Get implements Store.
func (j *JsonStore) Get(key string) []core.Reference {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	return j.store[key]
}

// GetFrom implements Store.
func (j *JsonStore) GetFrom(uri uri.URI, key string) []core.Reference {
	res := []core.Reference{}
	for _, ref := range j.Get(key) {
		if uri == ref.URI {
			res = append(res, ref)
		}
	}
	return res
}

// Delete implements Store.
func (j *JsonStore) Delete(uri uri.URI) {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	for id, refs := range j.store {
		filtered := []core.Reference{}
		for _, ref := range refs {
			if ref.URI != uri {
				filtered = append(filtered, ref)
			}
		}
		j.store[id] = filtered
	}
}

var _ Store = (*JsonStore)(nil)
