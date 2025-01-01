package core

import (
	"github.com/rockide/protocol"
	"go.lsp.dev/uri"
)

type Reference struct {
	Value string
	Uri   uri.URI
	Range *protocol.Range
}

type Store interface {
	Parse(uri uri.URI) error
	Get(key string) []Reference
	GetFrom(uri uri.URI, key Store) []Reference
	Delete(uri uri.URI)
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
