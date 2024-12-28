package core

import (
	"os"
	"sync"

	"github.com/ink0rr/go-jsonc"
	"github.com/ink0rr/rockide/textdocument"
	"go.lsp.dev/uri"
)

type DummyStore struct {
	pattern string
	store   map[uri.URI]*jsonc.Node
	mutex   sync.Mutex
}

// Delete implements Store.
func (d *DummyStore) Delete(uri uri.URI) {
	panic("unimplemented")
}

// Get implements Store.
func (d *DummyStore) Get(key string) []Reference {
	panic("unimplemented")
}

// GetFrom implements Store.
func (d *DummyStore) GetFrom(uri uri.URI, key Store) []Reference {
	panic("unimplemented")
}

// GetPattern implements Store.
func (d *DummyStore) GetPattern() string {
	return d.pattern
}

// Parse implements Store.
func (d *DummyStore) Parse(uri uri.URI) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	txt, err := os.ReadFile(uri.Filename())
	if err != nil {
		return err
	}
	document := textdocument.New(uri, string(txt))
	root, _ := jsonc.ParseTree(document.GetText(), nil)
	d.store[uri] = root
	return nil
}

var _ Store = (*DummyStore)(nil)

func NewDummyStore(pattern string) DummyStore {
	return DummyStore{
		pattern: pattern,
		store:   make(map[uri.URI]*jsonc.Node),
	}
}

var AnimationControllerStore = NewDummyStore(AnimationControllerGlob)
var AnimationStore = NewDummyStore(AnimationGlob)
var BlockStore = NewDummyStore(BlockGlob)
var EntityStore = NewDummyStore(EntityGlob)
var FeatureRuleStore = NewDummyStore(FeatureRuleGlob)
var FeatureStore = NewDummyStore(FeatureGlob)
var ItemStore = NewDummyStore(ItemGlob)
var TradeTableStore = NewDummyStore(TradeTableGlob)
var AttachableStore = NewDummyStore(AttachableGlob)
var ClientAnimationControllersStore = NewDummyStore(ClientAnimationControllersGlob)
var ClientAnimationsStore = NewDummyStore(ClientAnimationsGlob)
var ClientBlockStore = NewDummyStore(ClientBlockGlob)
var ClientEntityStore = NewDummyStore(ClientEntityGlob)
var GeometryStore = NewDummyStore(GeometryGlob)
var ItemTextureStore = NewDummyStore(ItemTextureGlob)
var ParticleStore = NewDummyStore(ParticleGlob)
var RenderControllerStore = NewDummyStore(RenderControllerGlob)
var SoundDefinitionStore = NewDummyStore(SoundDefinitionGlob)
var TerrainTextureStore = NewDummyStore(TerrainTextureGlob)
