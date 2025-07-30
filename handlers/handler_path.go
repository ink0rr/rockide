package handlers

import (
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

type Path struct {
	Pattern shared.Pattern
	Store   *stores.PathStore
}

func (s *Path) GetPattern() string {
	return s.Pattern.ToString()
}

func (s *Path) Parse(uri protocol.DocumentURI) error {
	s.Store.Insert(s.Pattern, uri)
	return nil
}

func (s *Path) GetPaths() []core.Symbol {
	return s.Store.Get()
}

func (s *Path) Delete(uri protocol.DocumentURI) {
	s.Store.Delete(uri)
}
