package server

import (
	"context"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/ink0rr/rockide/handlers"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/ink0rr/rockide/shared"
	"github.com/sourcegraph/jsonrpc2"
)

func SemanticTokens(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.SemanticTokensParams) (*protocol.SemanticTokens, error) {
	if !doublestar.MatchUnvalidated("**/"+shared.ProjectGlob+"/**/*.json", string(params.TextDocument.URI)) {
		return nil, nil
	}
	document := textdocument.Get(params.TextDocument.URI)
	return handlers.Molang.GetSemanticTokens(document), nil
}
