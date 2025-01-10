package rockide

import (
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/handlers"
	"github.com/ink0rr/rockide/internal/protocol"
)

func getProjectPaths() (bp string, rp string, err error) {
	fs := os.DirFS(baseDir)
	bpPaths, err := doublestar.Glob(fs, core.BpGlob)
	if err != nil {
		return
	}
	if len(bpPaths) > 0 {
		bp = filepath.Join(baseDir, bpPaths[0])
	}
	rpPaths, err := doublestar.Glob(fs, core.RpGlob)
	if err != nil {
		return
	}
	if len(rpPaths) > 0 {
		rp = filepath.Join(baseDir, rpPaths[0])
	}
	return
}

func FindHandler(uri protocol.DocumentURI) handlers.Handler {
	for _, handler := range jsonHandlers {
		if doublestar.MatchUnvalidated("**/"+handler.GetPattern(), string(uri)) {
			return handler
		}
	}
	return nil
}
