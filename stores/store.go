package stores

import (
	"github.com/ink0rr/rockide/core"
	"go.lsp.dev/uri"
)

type Store interface {
	Parse(uri uri.URI) error
	Get(key string) []core.Reference
	GetFrom(uri uri.URI, key string) []core.Reference
	Delete(uri uri.URI)
	GetPattern() string
}

// Returns a slice containing elements only if they're present in A but not in B.
func Difference(a []core.Reference, b []core.Reference) []core.Reference {
	result := []core.Reference{}
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
