package stores

import (
	"strings"
	"sync"

	"github.com/ink0rr/rockide/jsonc"
	"github.com/ink0rr/rockide/rockide/core"
	"github.com/ink0rr/rockide/textdocument"
	"github.com/rockide/protocol"
	"go.lsp.dev/uri"
)

type transformResult struct {
	Node  *jsonc.Node
	Value string
	Skip  bool
}

type jsonStoreEntry struct {
	Id             string
	path           []string
	JsonPath       [][]string
	TransformValue func(node *jsonc.Node) transformResult
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
		jsonPath := []string{}
		for _, path := range entry.path {
			jsonPath = append(jsonPath, strings.Split(path, "/")...)
		}
		entries[i].JsonPath = append(entry.JsonPath, jsonPath)
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
	document, err := textdocument.New(uri)
	if err != nil {
		return err
	}
	root, _ := jsonc.ParseTree(document.GetText(), nil)
	for _, entry := range j.entries {
		data := j.store[entry.Id]
		for _, path := range entry.JsonPath {
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
					if entry.TransformValue != nil {
						result := entry.TransformValue(node)
						if result.Skip {
							return
						}
						value = result.Value
					}
					data = append(data, core.Reference{
						Value: value.(string),
						URI:   uri,
						Range: &protocol.Range{
							Start: document.PositionAt(node.Offset),
							End:   document.PositionAt(node.Offset + node.Length),
						}})
				} else if node.Type == jsonc.NodeTypeProperty && node.Children != nil && index < len(node.Children) {
					value := node.Children[index].Value
					if entry.TransformValue != nil {
						result := entry.TransformValue(node.Children[index])
						if result.Skip {
							return
						}
						value = result.Value
					}
					targetNode := node.Children[index]
					data = append(data, core.Reference{
						Value: value.(string),
						URI:   uri,
						Range: &protocol.Range{
							Start: document.PositionAt(targetNode.Offset),
							End:   document.PositionAt(targetNode.Offset + targetNode.Length),
						}})
				} else if node.Children != nil {
					for _, child := range node.Children {
						extract(child)
					}
				}
			}
			for _, node := range core.FindNodesAtPath(root, path) {
				extract(node)
			}
		}
		j.store[entry.Id] = data
	}
	_ = root
	return nil
}

func jsonReference(document textdocument.TextDocument, value string, node *jsonc.Node) *core.Reference {
	res := core.Reference{Value: value, URI: document.URI, Range: &protocol.Range{}}
	if node != nil {
		res.Value = node.Value.(string)
	}
	return &res
}

// Get implements Store.
func (j *JsonStore) Get(key string) []core.Reference {
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
			if uri != ref.URI {
				filtered = append(filtered, ref)
			}
		}
		j.store[id] = refs
	}
}

var _ Store = (*JsonStore)(nil)
