package server

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/handlers"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var project core.Project

var handlerList = [...]handlers.Handler{
	// BP
	handlers.AnimationController,
	handlers.Animation,
	handlers.Block,
	handlers.Entity,
	handlers.Item,
	// RP
	handlers.ClientAnimationController,
	handlers.ClientAnimation,
	handlers.ClientEntity,
}

var storeList = [...]stores.Store{
	// BP
	stores.AnimationController,
	stores.Animation,
	stores.Block,
	stores.Entity,
	stores.FeatureRule,
	stores.Feature,
	stores.Item,
	stores.LootTable,
	stores.Recipe,
	stores.TradeTable,
	// RP
	stores.Attachable,
	stores.ClientAnimationController,
	stores.ClientAnimation,
	stores.ClientBlock,
	stores.ClientEntity,
	stores.Geometry,
	stores.ItemTexture,
	stores.Particle,
	stores.RenderController,
	stores.SoundDefinition,
	stores.Sound,
	stores.TerrainTexture,
	stores.Texture,
}

func setBaseDir(dir string) error {
	project = core.Project{}
	fsys := os.DirFS(dir)

	bpPaths, err := doublestar.Glob(fsys, shared.BpGlob, doublestar.WithFailOnIOErrors())
	if bpPaths == nil || err != nil {
		return errors.New("not a minecraft project")
	}
	bp := dir + "/" + bpPaths[0]
	log.Printf("Behavior pack: %s", bp)

	rpPaths, err := doublestar.Glob(fsys, shared.RpGlob, doublestar.WithFailOnIOErrors())
	if rpPaths == nil || err != nil {
		return errors.New("not a minecraft project")
	}
	rp := dir + "/" + rpPaths[0]
	log.Printf("Resource pack: %s", rp)

	project.BP = filepath.ToSlash(filepath.Clean(bp))
	project.RP = filepath.ToSlash(filepath.Clean(rp))

	return nil
}

func indexWorkspace() {
	startTime := time.Now()
	fsys := os.DirFS(".")
	totalFiles := atomic.Uint32{}
	skippedFiles := atomic.Uint32{}

	var wg sync.WaitGroup
	wg.Add(len(storeList))
	for _, store := range storeList {
		go func() {
			defer wg.Done()
			doublestar.GlobWalk(fsys, store.GetPattern(&project), func(path string, d fs.DirEntry) error {
				if d.IsDir() {
					return nil
				}
				uri := protocol.URIFromPath(path)
				err := store.Parse(uri)
				if err != nil {
					log.Printf("Error parsing file: %s\n\t%s", path, err)
					skippedFiles.Add(1)
					return nil
				}
				totalFiles.Add(1)
				return nil
			})
		}()
	}
	wg.Wait()

	totalTime := time.Since(startTime)
	log.Printf("Scanned %d files in %s", totalFiles.Load(), totalTime)
	if count := skippedFiles.Load(); count > 0 {
		log.Printf("Skipped %d files", count)
	}
}
