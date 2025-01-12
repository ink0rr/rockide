package stores

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/protocol"
)

type Store interface {
	GetPattern(project *core.Project) string
	Parse(uri protocol.DocumentURI) error
	Get(key string) []core.Reference
	GetFrom(uri protocol.DocumentURI, key string) []core.Reference
	Delete(uri protocol.DocumentURI)
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
