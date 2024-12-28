package core

import (
	"path/filepath"
	"regexp"

	"go.lsp.dev/protocol"
)

type BehaviorStore struct {
	refs []Reference
}

var bpRegex = regexp.MustCompile("(behavior_pack|[^\\/]*?bp|bp_[^\\/]*?)\\/")

func (s *BehaviorStore) Parse(uri protocol.URI) error {
	path, err := filepath.Rel(".", uri.Filename())
	if err != nil {
		return err
	}
	path = bpRegex.Split(path, -1)[2]
	s.refs = append(s.refs, Reference{Value: path, Uri: uri})
	return nil
}
