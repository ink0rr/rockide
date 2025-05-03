package server

import (
	"context"

	"github.com/ink0rr/rockide/handlers"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/sourcegraph/jsonrpc2"
)

func Definition(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.DefinitionParams) ([]protocol.LocationLink, error) {
	handler, ok := handlers.Find(params.TextDocument.URI).(handlers.DefinitionProvider)
	if !ok {
		return nil, nil
	}
	document := textdocument.Get(params.TextDocument.URI)
	return handler.Definitions(document, params.Position), nil
}
