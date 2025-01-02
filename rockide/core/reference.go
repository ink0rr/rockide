package core

import (
	"github.com/rockide/protocol"
	"go.lsp.dev/uri"
)

type Reference struct {
	Value string
	URI   uri.URI
	Range *protocol.Range
}
