package core

import "github.com/ink0rr/rockide/internal/protocol"

type Project struct {
	BP string
	RP string
}

type Symbol struct {
	Value string
	URI   protocol.DocumentURI
	Range *protocol.Range
}
