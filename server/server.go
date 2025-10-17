package server

import (
	"errors"
	"io/fs"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/handlers"
	"github.com/ink0rr/rockide/internal/protocol"
)

func cleanPath(path string) string {
	u := url.URL{Path: filepath.ToSlash(filepath.Clean(path))}
	return u.String()
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
			BP: cleanPath(bp),
			RP: cleanPath(rp),
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
		BP: cleanPath(bp),
		RP: cleanPath(rp),
	}, nil
}

func indexWorkspace() {
	startTime := time.Now()
	fsys := os.DirFS(".")
	totalFiles := atomic.Uint32{}
	skippedFiles := atomic.Uint32{}

	var wg sync.WaitGroup
	wg.Add(len(handlers.GetAll()))
	for _, store := range handlers.GetAll() {
		go func() {
			defer wg.Done()
			pattern, err := url.PathUnescape(store.GetPattern())
			if err != nil {
				log.Printf("Failed to unescape pattern: %s", store.GetPattern())
				return
			}
			doublestar.GlobWalk(fsys, pattern, func(path string, d fs.DirEntry) error {
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
