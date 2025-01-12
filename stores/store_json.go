package stores

import (
	"strings"
	"sync"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/jsonc"
	"github.com/ink0rr/rockide/textdocument"
)

type jsonStoreEntry struct {
	Id        string
	Path      []string
	Transform func(node *jsonc.Node) *string

	// Used to cache Path splitted by '/'
	jsonPath [][]string
}

func (j *jsonStoreEntry) getJsonPath() [][]string {
	if j.jsonPath == nil {
		for _, path := range j.Path {
			j.jsonPath = append(j.jsonPath, strings.Split(path, "/"))
		}
	}
	return j.jsonPath
}

type jsonStore struct {
	pattern core.Pattern
	entries []jsonStoreEntry
	store   map[string][]core.Reference
	mutex   sync.Mutex
}

func newJsonStore(pattern core.Pattern, entries []jsonStoreEntry) *jsonStore {
	return &jsonStore{
		pattern: pattern,
		entries: entries,
		store:   make(map[string][]core.Reference),
	}
}

// GetPattern implements Store.
func (j *jsonStore) GetPattern(project *core.Project) string {
	return j.pattern.Resolve(project)
}

// Parse implements Store.
func (j *jsonStore) Parse(uri protocol.DocumentURI) error {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	document, err := textdocument.Open(uri)
	if err != nil {
		return err
	}
	root, _ := jsonc.ParseTree(document.GetText(), nil)
	for _, entry := range j.entries {
		data := j.store[entry.Id]
		for _, path := range entry.getJsonPath() {
			// If the last path segment is "*", we want to grab the values instead of keys
			// The only exception is when the path is just a single "*"
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
func (j *jsonStore) Get(key string) []core.Reference {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	return j.store[key]
}

// GetFrom implements Store.
func (j *jsonStore) GetFrom(uri protocol.DocumentURI, key string) []core.Reference {
	res := []core.Reference{}
	for _, ref := range j.Get(key) {
		if ref.URI == uri {
			res = append(res, ref)
		}
	}
	return res
}

// Delete implements Store.
func (j *jsonStore) Delete(uri protocol.DocumentURI) {
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
