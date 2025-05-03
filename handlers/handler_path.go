package handlers

import (
	"path/filepath"
	"strings"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/shared"
)

type Path struct {
	Pattern    shared.Pattern
	TrimSuffix bool
	store      []core.Symbol
}

func (s *Path) GetPattern() string {
	return s.Pattern.ToString()
}

func (s *Path) Parse(uri protocol.DocumentURI) error {
	path, err := filepath.Rel(shared.Getwd(), uri.Path())
	if err != nil {
		panic(err)
	}
	packType := s.Pattern.PackType()
	path = filepath.ToSlash(path)
	_, path, found := strings.Cut(path, packType+"/")
	if !found {
		panic("invalid project path")
	}
	if s.TrimSuffix {
		path = strings.TrimSuffix(path, filepath.Ext(path))
	}
	s.store = append(s.store, core.Symbol{Value: path, URI: uri})
	return nil
}

func (s *Path) GetPaths() []core.Symbol {
	return s.store
}

func (s *Path) Delete(uri protocol.DocumentURI) {
	filtered := []core.Symbol{}
	for _, ref := range s.store {
		if ref.URI != uri {
			filtered = append(filtered, ref)
		}
	}
	s.store = filtered
}
