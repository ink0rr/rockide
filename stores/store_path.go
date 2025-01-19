package stores

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/shared"
)

var cwd string

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cwd = dir + string(filepath.Separator)
}

type PathStore struct {
	pattern    shared.Pattern
	refs       []core.Reference
	trimSuffix bool
}

func (s *PathStore) GetPattern() string {
	return s.pattern.ToString()
}

func (s *PathStore) Parse(uri protocol.DocumentURI) error {
	path, err := filepath.Rel(cwd, uri.Path())
	if err != nil {
		panic(err)
	}
	packType := s.pattern.PackType()
	path = filepath.ToSlash(path)
	_, path, found := strings.Cut(path, packType+"/")
	if !found {
		panic("invalid project path")
	}
	if s.trimSuffix {
		path = strings.TrimSuffix(path, filepath.Ext(path))
	}
	s.refs = append(s.refs, core.Reference{Value: path, URI: uri})
	return nil
}

func (s *PathStore) Delete(uri protocol.DocumentURI) {
	filtered := []core.Reference{}
	for _, ref := range s.refs {
		if ref.URI != uri {
			filtered = append(filtered, ref)
		}
	}
	s.refs = filtered
}

func (s *PathStore) GetPaths() []core.Reference {
	return s.refs
}
