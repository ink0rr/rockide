package server

import (
	"context"

	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/sourcegraph/jsonrpc2"
)

func Completion(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.CompletionParams) ([]protocol.CompletionItem, error) {
	document := textdocument.Get(params.TextDocument.URI)
	actions := findActions(document, params.Position)
	if actions == nil || actions.Completions == nil {
		return nil, nil
	}
	return actions.Completions(), nil
}
