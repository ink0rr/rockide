package rockide

import (
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/ink0rr/rockide/core"
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
