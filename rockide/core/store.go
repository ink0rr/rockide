package core

import (
	"go.lsp.dev/protocol"
)

type Reference struct {
	Value string
	Uri   protocol.URI
	Range *protocol.Range
}

type Store interface {
	Parse(uri protocol.URI) error
	Get(key string) []Reference
	GetFrom(uri protocol.URI, key Store) []Reference
	Delete(uri protocol.URI)
	GetPattern() string
}

func Difference(a []Reference, b []Reference) []Reference {
	result := []Reference{}
	set := map[string]bool{}
	for _, ref := range b {
		set[ref.Value] = true
	}
	for _, ref := range a {
		if !set[ref.Value] {
			result = append(result, ref)
		}
	}
	return result
}
