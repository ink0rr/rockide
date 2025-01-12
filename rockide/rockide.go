package rockide

import (
	"context"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/arexon/fsnotify"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/handlers"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/stores"
	"github.com/ink0rr/rockide/textdocument"
)

var project core.Project

var storeList = []stores.Store{
	// BP
	stores.AnimationController,
	stores.Animation,
	stores.Block,
	stores.Entity,
	stores.FeatureRule,
	stores.Feature,
	stores.Item,
	stores.LootTable,
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
	stores.TerrainTexture,
}

var jsonHandlers = []handlers.Handler{
	// BP
	handlers.AnimationController,
	handlers.Animation,
	handlers.Block,
	handlers.Entity,
	handlers.Item,
	// RP
	handlers.ClientAnimationController,
	handlers.ClientAnimation,
}

func SetBaseDir(dir string) error {
	project = core.Project{}
	fsys := os.DirFS(dir)

	bpPaths, err := doublestar.Glob(fsys, core.BpGlob, doublestar.WithFailOnIOErrors())
	if bpPaths == nil || err != nil {
		return errors.New("not a minecraft project")
	}
	bp := dir + "/" + bpPaths[0]
	log.Printf("Behavior pack: %s", bp)

	rpPaths, err := doublestar.Glob(fsys, core.RpGlob, doublestar.WithFailOnIOErrors())
	if rpPaths == nil || err != nil {
		return errors.New("not a minecraft project")
	}
	rp := dir + "/" + rpPaths[0]
	log.Printf("Resource pack: %s", rp)

	project.BP = filepath.ToSlash(filepath.Clean(bp))
	project.RP = filepath.ToSlash(filepath.Clean(rp))

	return nil
}

func IndexWorkspaces(ctx context.Context) {
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

func WatchFiles(ctx context.Context) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	if err := watcher.Add(filepath.Join(project.BP, "...")); err != nil {
		return errors.Join(err, errors.New("failed to watch BP path"))
	}
	if err := watcher.Add(filepath.Join(project.RP, "...")); err != nil {
		return errors.Join(err, errors.New("failed to watch RP path"))
	}
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				uri := protocol.URIFromPath(event.Name)
				if event.Op.Has(fsnotify.Remove | fsnotify.Rename) {
					OnDelete(uri)
					continue
				}
				if stat, err := os.Stat(event.Name); err != nil || stat.IsDir() {
					continue
				}
				switch {
				case event.Op.Has(fsnotify.Create):
					OnCreate(uri)
				case event.Op.Has(fsnotify.Write):
					OnChange(uri)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()
	return nil
}

func OnCreate(uri protocol.DocumentURI) {
	log.Printf("create: %s", uri)
	for _, store := range storeList {
		if doublestar.MatchUnvalidated("**/"+store.GetPattern(&project), string(uri)) {
			store.Parse(uri)
			break
		}
	}
}

func OnChange(uri protocol.DocumentURI) {
	log.Printf("change: %s", uri)
	for _, store := range storeList {
		if doublestar.MatchUnvalidated("**/"+store.GetPattern(&project), string(uri)) {
			store.Delete(uri)
			store.Parse(uri)
			break
		}
	}
}

func OnDelete(uri protocol.DocumentURI) {
	log.Printf("delete: %s", uri)
	for _, store := range storeList {
		if doublestar.MatchUnvalidated("**/"+store.GetPattern(&project), string(uri)) {
			store.Delete(uri)
			break
		}
	}
}

func FindActions(document *textdocument.TextDocument, position *protocol.Position) *handlers.HandlerActions {
	for _, handler := range jsonHandlers {
		if doublestar.MatchUnvalidated("**/"+handler.GetPattern(&project), string(document.URI)) {
			return handler.GetActions(document, position)
		}
	}
	return nil
}
