package rockide

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/handlers"
	"go.lsp.dev/uri"
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

func toURI(path string) (uri.URI, error) {
	var result uri.URI
	abs, err := filepath.Abs(path)
	if err != nil {
		return result, errors.New("failed to resolve absolute path")
	}
	result = uri.File(abs)
	return result, nil
}

func FindHandler(uri uri.URI) handlers.Handler {
	name := uri.Filename()
	name = strings.ReplaceAll(name, "\\", "/")
	for _, handler := range jsonHandlers {
		if doublestar.MatchUnvalidated("**/"+handler.GetPattern(), name) {
			return handler
		}
	}
	return nil
}
