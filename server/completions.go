package server

import (
	"context"

	"github.com/ink0rr/rockide/handlers"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/sourcegraph/jsonrpc2"
)

func Completion(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.CompletionParams) ([]protocol.CompletionItem, error) {
	handler, ok := handlers.Find(params.TextDocument.URI).(handlers.CompletionProvider)
	if !ok {
		return nil, nil
	}
	document := textdocument.Get(params.TextDocument.URI)
	return handler.Completions(document, params.Position), nil
}
