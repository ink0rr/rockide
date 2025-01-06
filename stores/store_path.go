package stores

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ink0rr/rockide/core"
	"go.lsp.dev/uri"
)

var cwd string

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cwd = dir + string(filepath.Separator)
}

var (
	bpRegex = regexp.MustCompile(`(?i)(behavior_pack|[^\/]*?bp|bp_[^\/]*?)[\\/]`)
	rpRegex = regexp.MustCompile(`(?i)(resource_pack|[^\/]*?rp|rp_[^\/]*?)[\\/]`)
)

type BehaviorStore struct {
	pattern string
	refs    []core.Reference
}

// GetPattern implements Store.
func (b *BehaviorStore) GetPattern() string {
	return b.pattern
}

// Parse implements Store.
func (b *BehaviorStore) Parse(uri uri.URI) error {
	path, found := strings.CutPrefix(uri.Filename(), cwd)
	path = strings.ReplaceAll(path, "\\", "/")
	if !found {
		panic("Failed to resolve path")
	}
	path = bpRegex.Split(path, -1)[1]
	b.refs = append(b.refs, core.Reference{Value: path, URI: uri})
	return nil
}

// Get implements Store.
func (b *BehaviorStore) Get(key string) []core.Reference {
	return b.refs
}

// GetFrom implements Store.
func (b *BehaviorStore) GetFrom(uri uri.URI, key string) []core.Reference {
	filtered := []core.Reference{}
	for _, ref := range b.refs {
		if ref.URI == uri {
			filtered = append(filtered, ref)
		}
	}
	return filtered
}

// Delete implements Store.
func (b *BehaviorStore) Delete(uri uri.URI) {
	filtered := []core.Reference{}
	for _, ref := range b.refs {
		if ref.URI != uri {
			filtered = append(filtered, ref)
		}
	}
	b.refs = filtered
}

type ResourceStore struct {
	pattern string
	refs    []core.Reference
}

// GetPattern implements Store.
func (r *ResourceStore) GetPattern() string {
	return r.pattern
}

// Parse implements Store.
func (r *ResourceStore) Parse(uri uri.URI) error {
	path, found := strings.CutPrefix(uri.Filename(), cwd)
	path = strings.ReplaceAll(path, "\\", "/")
	if !found {
		panic("Failed to resolve path")
	}
	path = rpRegex.Split(path, -1)[1]
	path = strings.TrimSuffix(path, filepath.Ext(path))
	r.refs = append(r.refs, core.Reference{Value: path, URI: uri})
	return nil
}

// Get implements Store.
func (r *ResourceStore) Get(key string) []core.Reference {
	return r.refs
}

// GetFrom implements Store.
func (r *ResourceStore) GetFrom(uri uri.URI, key string) []core.Reference {
	filtered := []core.Reference{}
	for _, ref := range r.refs {
		if ref.URI == uri {
			filtered = append(filtered, ref)
		}
	}
	return filtered
}

// Delete implements Store.
func (r *ResourceStore) Delete(uri uri.URI) {
	filtered := []core.Reference{}
	for _, ref := range r.refs {
		if ref.URI != uri {
			filtered = append(filtered, ref)
		}
	}
	r.refs = filtered
}
