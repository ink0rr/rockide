package rockide

import (
	"context"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/arexon/fsnotify"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/ink0rr/rockide/handlers"
	"github.com/ink0rr/rockide/stores"
	"go.lsp.dev/uri"
)

var baseDir = "."

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

var jsonHandlers = []*handlers.JsonHandler{
	handlers.Entity,
}

func SetBaseDir(dir string) {
	baseDir = dir
}

func IndexWorkspaces(ctx context.Context) error {
	startTime := time.Now()
	fsys := os.DirFS(baseDir)
	totalFiles := atomic.Uint32{}
	skippedFiles := atomic.Uint32{}

	var wg sync.WaitGroup
	for _, store := range storeList {
		go func() {
			defer wg.Done()
			wg.Add(1)
			doublestar.GlobWalk(fsys, store.GetPattern(), func(path string, d fs.DirEntry) error {
				if d.IsDir() {
					return nil
				}
				uri, err := toURI(filepath.Join(baseDir, path))
				if err != nil {
					log.Printf("Error parsing URI: %s\n\t%s", err, path)
					skippedFiles.Add(1)
					return nil
				}
				err = store.Parse(uri)
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
	return nil
}

func WatchFiles(ctx context.Context) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	bp, rp, err := getProjectPaths()
	if err != nil {
		return errors.Join(errors.New("failed to get project paths"), err)
	}
	if err := watcher.Add(filepath.Join(bp, "...")); err != nil {
		return errors.Join(errors.New("failed to watch BP path"), err)
	}
	if err := watcher.Add(filepath.Join(rp, "...")); err != nil {
		return errors.Join(errors.New("failed to watch RP path"), err)
	}
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				stat, err := os.Stat(event.Name)
				if err != nil || stat.IsDir() {
					continue
				}
				uri, err := toURI(event.Name)
				if err != nil {
					log.Println(err)
					continue
				}
				if event.Op.Has(fsnotify.Create) {
					OnCreate(uri)
					continue
				}
				if event.Op.Has(fsnotify.Write) {
					OnChange(uri)
					continue
				}
				if event.Op.Has(fsnotify.Remove | fsnotify.Rename) {
					OnDelete(uri)
					continue
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

func OnCreate(uri uri.URI) {
	name := uri.Filename()
	name = strings.ReplaceAll(name, "\\", "/")
	log.Printf("create: %s", name)
	for _, store := range storeList {
		if doublestar.MatchUnvalidated("**/"+store.GetPattern(), name) {
			store.Parse(uri)
			break
		}
	}
}

func OnChange(uri uri.URI) {
	name := uri.Filename()
	name = strings.ReplaceAll(name, "\\", "/")
	log.Printf("change: %s", name)
	for _, store := range storeList {
		if doublestar.MatchUnvalidated("**/"+store.GetPattern(), name) {
			store.Delete(uri)
			store.Parse(uri)
			break
		}
	}
}

func OnDelete(uri uri.URI) {
	name := uri.Filename()
	name = strings.ReplaceAll(name, "\\", "/")
	log.Printf("delete: %s", name)
	for _, store := range storeList {
		if doublestar.MatchUnvalidated("**/"+store.GetPattern(), name) {
			store.Delete(uri)
			break
		}
	}
}
