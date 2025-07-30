package stores

import (
	"path/filepath"
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/vanilla"
)

type PathStore struct {
	store       []core.Symbol
	trimSuffix  bool
	VanillaData mapset.Set[string]
}

func NewPathStore(trimSuffix bool, vanillaData mapset.Set[string]) *PathStore {
	return &PathStore{trimSuffix: trimSuffix, VanillaData: vanillaData}
}

func (s *PathStore) Insert(pattern shared.Pattern, uri protocol.DocumentURI) {
	path, err := filepath.Rel(shared.Getwd(), uri.Path())
	if err != nil {
		panic(err)
	}
	packType := pattern.PackType()
	path = filepath.ToSlash(path)
	_, path, found := strings.Cut(path, packType+"/")
	if !found {
		panic("invalid project path")
	}
	if s.trimSuffix {
		path = strings.TrimSuffix(path, filepath.Ext(path))
	}
	s.store = append(s.store, core.Symbol{Value: path, URI: uri})
}

func (s *PathStore) Get() []core.Symbol {
	return s.store
}

func (s *PathStore) Delete(uri protocol.DocumentURI) {
	s.store = slices.DeleteFunc(s.store, func(symbol core.Symbol) bool {
		return symbol.URI == uri
	})
}

var (
	LootTablePath  = NewPathStore(false, vanilla.LootTablePath)
	TradeTablePath = NewPathStore(false, vanilla.TradeTablePath)
	SoundPath      = NewPathStore(true, vanilla.SoundPath)
	TexturePath    = NewPathStore(true, vanilla.TexturePath)
)
