package stores

import (
	"sync"

	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/jsonc"
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
	d.mutex.Lock()
	defer d.mutex.Unlock()
	delete(d.store, uri)
}

// Get implements Store.
func (d *DummyStore) Get(key string) []core.Reference {
	panic("unimplemented")
}

// GetFrom implements Store.
func (d *DummyStore) GetFrom(uri uri.URI, key string) []core.Reference {
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
	document, err := textdocument.Open(uri)
	if err != nil {
		return err
	}
	root, _ := jsonc.ParseTree(document.GetText(), nil)
	d.store[uri] = root
	return nil
}

var _ Store = (*DummyStore)(nil)

func NewDummyStore(pattern string) *DummyStore {
	return &DummyStore{
		pattern: pattern,
		store:   make(map[uri.URI]*jsonc.Node),
	}
}

var Block = NewDummyStore(core.BlockGlob)
var FeatureRule = NewDummyStore(core.FeatureRuleGlob)
var Feature = NewDummyStore(core.FeatureGlob)
var Item = NewDummyStore(core.ItemGlob)
var Attachable = NewDummyStore(core.AttachableGlob)
var ClientAnimationControllers = NewDummyStore(core.ClientAnimationControllersGlob)
var ClientAnimations = NewDummyStore(core.ClientAnimationsGlob)
var ClientBlock = NewDummyStore(core.ClientBlockGlob)
var Geometry = NewDummyStore(core.GeometryGlob)
var ItemTexture = NewDummyStore(core.ItemTextureGlob)
var Particle = NewDummyStore(core.ParticleGlob)
var RenderController = NewDummyStore(core.RenderControllerGlob)
var SoundDefinition = NewDummyStore(core.SoundDefinitionGlob)
var TerrainTexture = NewDummyStore(core.TerrainTextureGlob)
