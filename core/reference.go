package core

import "github.com/ink0rr/rockide/internal/protocol"

type Reference struct {
	Value string
	URI   protocol.DocumentURI
	Range *protocol.Range
}
