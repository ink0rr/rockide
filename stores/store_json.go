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
	Path      []shared.JsonPath
	Transform func(node *jsonc.Node) *string
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
	document := textdocument.Get(uri)
	if document == nil {
		doc, err := textdocument.ReadFile(uri)
		if err != nil {
			return err
		}
		document = doc
	}
	root, _ := jsonc.ParseTree(document.GetText(), nil)
	for _, entry := range j.entries {
		data := j.store[entry.Id]
		for _, jsonPath := range entry.Path {
			index := 1
			if jsonPath.IsKey {
				index = 0
			}
			extract := func(node *jsonc.Node) {
				if node.Type == jsonc.NodeTypeProperty && len(node.Children) > index {
					targetNode := node.Children[index]
					value := targetNode.Value
					if entry.Transform != nil {
						result := entry.Transform(targetNode)
						if result == nil {
							return
						}
						value = *result
					}
					nodeValue, ok := value.(string)
					if !ok {
						return
					}
					data = append(data, core.Reference{
						Value: nodeValue,
						URI:   uri,
						Range: &protocol.Range{
							Start: document.PositionAt(targetNode.Offset),
							End:   document.PositionAt(targetNode.Offset + targetNode.Length),
						},
					})
				} else if !jsonPath.IsKey && node.Type == jsonc.NodeTypeString {
					value := node.Value
					if entry.Transform != nil {
						result := entry.Transform(node)
						if result == nil {
							return
						}
						value = *result
					}
					nodeValue, ok := value.(string)
					if !ok {
						return
					}
					data = append(data, core.Reference{
						Value: nodeValue,
						URI:   uri,
						Range: &protocol.Range{
							Start: document.PositionAt(node.Offset),
							End:   document.PositionAt(node.Offset + node.Length),
						},
					})
				}
			}
			for _, node := range findNodesAtPath(root, jsonPath.Path) {
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
