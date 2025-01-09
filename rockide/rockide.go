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
	"github.com/ink0rr/rockide/jsonc"
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

func SetBaseDir(dir string) {
	baseDir = dir
}

func IsMinecraftWorkspace(ctx context.Context) bool {
	fsys := os.DirFS(baseDir)
	hasManifest := false
	err := doublestar.GlobWalk(fsys, core.ProjectGlob+"/manifest.json", func(path string, d fs.DirEntry) error {
		if d.IsDir() {
			return nil
		}
		path = filepath.Join(baseDir, path)
		log.Printf("Found manifest: %s", path)
		txt, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		root, _ := jsonc.ParseTree(string(txt), nil)
		formatVersion := jsonc.FindNodeAtLocation(root, jsonc.Path{"format_version"}) != nil
		header := jsonc.FindNodeAtLocation(root, jsonc.Path{"header"}) != nil
		modules := jsonc.FindNodeAtLocation(root, jsonc.Path{"modules"}) != nil
		if formatVersion && header && modules {
			hasManifest = true
		}
		return nil
	})
	return err == nil && hasManifest
}

func IndexWorkspaces(ctx context.Context) {
	startTime := time.Now()
	fsys := os.DirFS(baseDir)
	totalFiles := atomic.Uint32{}
	skippedFiles := atomic.Uint32{}

	var wg sync.WaitGroup
	wg.Add(len(storeList))
	for _, store := range storeList {
		go func() {
			defer wg.Done()
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
}

func WatchFiles(ctx context.Context) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	bp, rp, err := getProjectPaths()
	if err != nil {
		return errors.Join(err, errors.New("failed to get project paths"))
	}
	if err := watcher.Add(filepath.Join(bp, "...")); err != nil {
		return errors.Join(err, errors.New("failed to watch BP path"))
	}
	if err := watcher.Add(filepath.Join(rp, "...")); err != nil {
		return errors.Join(err, errors.New("failed to watch RP path"))
	}
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				uri, err := toURI(event.Name)
				if err != nil {
					log.Println(err)
					continue
				}
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

func OnCreate(uri uri.URI) {
	log.Printf("create: %s", uri)
	for _, store := range storeList {
		if doublestar.MatchUnvalidated("**/"+store.GetPattern(), string(uri)) {
			store.Parse(uri)
			break
		}
	}
}

func OnChange(uri uri.URI) {
	log.Printf("change: %s", uri)
	for _, store := range storeList {
		if doublestar.MatchUnvalidated("**/"+store.GetPattern(), string(uri)) {
			store.Delete(uri)
			store.Parse(uri)
			break
		}
	}
}

func OnDelete(uri uri.URI) {
	log.Printf("delete: %s", uri)
	for _, store := range storeList {
		if doublestar.MatchUnvalidated("**/"+store.GetPattern(), string(uri)) {
			store.Delete(uri)
			break
		}
	}
}
