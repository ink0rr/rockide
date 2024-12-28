package handler

import (
	"context"

	"go.lsp.dev/protocol"
)

func TextDocumentDidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	if len(params.ContentChanges) > 0 {
		// Do something
	}
	return nil
}
