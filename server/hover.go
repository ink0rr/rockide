package server

import (
	"context"

	"github.com/ink0rr/rockide/handlers"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/sourcegraph/jsonrpc2"
)

func Hover(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.HoverParams) (*protocol.Hover, error) {
	handler, ok := handlers.Find(params.TextDocument.URI).(handlers.HoverProvider)
	if !ok {
		return nil, nil
	}
	document := textdocument.Get(params.TextDocument.URI)
	return handler.Hover(document, params.Position), nil
}
