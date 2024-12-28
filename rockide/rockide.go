package rockide

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/ink0rr/rockide/rockide/core"
	"go.lsp.dev/uri"
)

type Rockide struct {
	stores []core.Store
}

func NewRockide() *Rockide {
	return &Rockide{
		stores: []core.Store{
			&core.AnimationControllerStore,
			&core.AnimationStore,
			&core.BlockStore,
			&core.EntityStore,
			&core.FeatureRuleStore,
			&core.FeatureStore,
			&core.ItemStore,
			&core.TradeTableStore,
			&core.AttachableStore,
			&core.ClientAnimationControllersStore,
			&core.ClientAnimationsStore,
			&core.ClientBlockStore,
			&core.ClientEntityStore,
			&core.GeometryStore,
			&core.ItemTextureStore,
			&core.ParticleStore,
			&core.RenderControllerStore,
			&core.SoundDefinitionStore,
			&core.TerrainTextureStore,
		},
	}
}

func (r *Rockide) IndexWorkspaces(ctx context.Context) error {
	startTime := time.Now()
	logger := GetLogger(ctx)
	fsys := os.DirFS(".")
	totalFiles := atomic.Int32{}

	var wg sync.WaitGroup
	for _, store := range r.stores {
		go func() {
			defer wg.Done()
			wg.Add(1)
			matches, err := doublestar.Glob(fsys, store.GetPattern(), doublestar.WithFilesOnly())
			if err != nil {
				return
			}
			for _, match := range matches {
				uri, err := toURI(match)
				if err != nil {
					return
				}
				totalFiles.Add(1)
				store.Parse(uri)
			}
		}()
	}
	wg.Wait()
	totalTime := time.Now().Sub(startTime)
	logger.Printf("Scanned %d files in %s", totalFiles.Load(), totalTime)
	return nil
}

func toURI(path string) (uri.URI, error) {
	var result uri.URI
	abs, err := filepath.Abs(path)
	if err != nil {
		return result, errors.New("Failed to resolve absolute path")
	}
	result, err = uri.Parse("file:///" + abs)
	if err != nil {
		return result, errors.New(fmt.Sprintf("Failed to parse uri: %s", abs))
	}
	return result, nil
}
