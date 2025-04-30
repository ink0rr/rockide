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
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/ink0rr/rockide/stores"
)

var handlerList = [...]handlers.Handler{
	// BP
	handlers.AnimationController,
	handlers.Animation,
	handlers.Block,
	handlers.Entity,
	handlers.FeatureRule,
	handlers.Feature,
	handlers.Item,
	handlers.LootTable,
	// RP
	handlers.Attachable,
	handlers.ClientAnimationController,
	handlers.ClientAnimation,
	handlers.ClientBlock,
	handlers.ClientEntity,
	handlers.ClientSound,
	handlers.Geometry,
	handlers.ItemTexture,
	handlers.Particle,
	handlers.RenderController,
	handlers.SoundDefinition,
	handlers.TerrainTexture,
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
	stores.ClientSound,
	stores.Geometry,
	stores.ItemTexture,
	stores.Particle,
	stores.RenderController,
	stores.SoundDefinition,
	stores.Sound,
	stores.TerrainTexture,
	stores.Texture,
}

func findProjectPaths(params any) (*core.Project, error) {
	// try to get project paths from user settings
	options, ok := params.(map[string]any)
	if ok {
		bp, ok := options["behaviorPack"].(string)
		if !ok {
			return nil, errors.New("invalid initialization options")
		}
		rp, ok := options["resourcePack"].(string)
		if !ok {
			return nil, errors.New("invalid initialization options")
		}
		return &core.Project{
			BP: filepath.ToSlash(filepath.Clean(bp)),
			RP: filepath.ToSlash(filepath.Clean(rp)),
		}, nil
	}

	// if not found, search the current dir and the 'packs' dir
	dir := "."
	if stat, err := os.Stat("packs"); err == nil && stat.IsDir() {
		dir = "packs"
	}
	fsys := os.DirFS(dir)

	bpPaths, err := doublestar.Glob(fsys, "{behavior_pack,*BP,BP_*,*bp,bp_*}", doublestar.WithFailOnIOErrors())
	if bpPaths == nil || err != nil {
		return nil, errors.New("not a minecraft project")
	}
	bp := dir + "/" + bpPaths[0]
	log.Printf("Behavior pack: %s", bp)

	rpPaths, err := doublestar.Glob(fsys, "{resource_pack,*RP,RP_*,*rp,rp_*}", doublestar.WithFailOnIOErrors())
	if rpPaths == nil || err != nil {
		return nil, errors.New("not a minecraft project")
	}
	rp := dir + "/" + rpPaths[0]
	log.Printf("Resource pack: %s", rp)

	return &core.Project{
		BP: filepath.ToSlash(filepath.Clean(bp)),
		RP: filepath.ToSlash(filepath.Clean(rp)),
	}, nil
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
			doublestar.GlobWalk(fsys, store.Pattern(), func(path string, d fs.DirEntry) error {
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

func findActions(document *textdocument.TextDocument, position protocol.Position) *handlers.HandlerActions {
	for _, handler := range handlerList {
		if doublestar.MatchUnvalidated("**/"+handler.Pattern(), string(document.URI)) {
			return handler.GetActions(document, position)
		}
	}
	return nil
}

func findStore(uri protocol.DocumentURI) stores.Store {
	for _, store := range storeList {
		if doublestar.MatchUnvalidated("**/"+store.Pattern(), string(uri)) {
			return store
		}
	}
	return nil
}
