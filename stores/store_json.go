package stores

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/ink0rr/rockide/shared"
)

type jsonStoreEntry struct {
	Id        string
	Path      []string
	Transform func(node *jsonc.Node) *string

	// Used to cache Path splitted by '/'
	jsonPath []jsonc.Path
}

func (j *jsonStoreEntry) getJsonPath() []jsonc.Path {
	if j.jsonPath == nil {
		for _, path := range j.Path {
			j.jsonPath = append(j.jsonPath, jsonc.NewPath(path))
		}
	}
	return j.jsonPath
}

type JsonStore struct {
	pattern    shared.Pattern
	savePath   bool
	trimSuffix bool
	entries    []jsonStoreEntry
	store      map[string][]core.Reference
	mutex      sync.Mutex
}

func (j *JsonStore) Pattern() string {
	return j.pattern.ToString()
}

func (j *JsonStore) Parse(uri protocol.DocumentURI) error {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	if j.store == nil {
		j.store = make(map[string][]core.Reference)
	}
	if j.savePath {
		j.parsePath(uri)
	}
	document, err := textdocument.ReadFile(uri)
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
					if value == nil {
						return
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

func (j *JsonStore) parsePath(uri protocol.DocumentURI) {
	path, err := filepath.Rel(shared.Getwd(), uri.Path())
	if err != nil {
		panic(err)
	}
	packType := j.pattern.PackType()
	path = filepath.ToSlash(path)
	_, path, found := strings.Cut(path, packType+"/")
	if !found {
		panic("invalid project path")
	}
	if j.trimSuffix {
		path = strings.TrimSuffix(path, filepath.Ext(path))
	}
	j.store["path"] = append(j.store["path"], core.Reference{Value: path, URI: uri})
}

func (j *JsonStore) Delete(uri protocol.DocumentURI) {
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

func (j *JsonStore) Get(key string) []core.Reference {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	return j.store[key]
}

func (j *JsonStore) GetFrom(uri protocol.DocumentURI, key string) []core.Reference {
	res := []core.Reference{}
	for _, ref := range j.Get(key) {
		if ref.URI == uri {
			res = append(res, ref)
		}
	}
	return res
}
