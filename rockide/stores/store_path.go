package stores

import (
	"path/filepath"
	"regexp"

	"github.com/ink0rr/rockide/rockide/core"
	"go.lsp.dev/uri"
)

type BehaviorStore struct {
	refs []core.Reference
}

var bpRegex = regexp.MustCompile("(behavior_pack|[^\\/]*?bp|bp_[^\\/]*?)\\/")

func (s *BehaviorStore) Parse(uri uri.URI) error {
	path, err := filepath.Rel(".", uri.Filename())
	if err != nil {
		return err
	}
	path = bpRegex.Split(path, -1)[2]
	s.refs = append(s.refs, core.Reference{Value: path, URI: uri})
	return nil
}
