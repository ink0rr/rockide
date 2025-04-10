package server

import (
	"context"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/ink0rr/rockide/handlers"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/ink0rr/rockide/shared"
	"github.com/sourcegraph/jsonrpc2"
)

func SignatureHelp(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	if !doublestar.MatchUnvalidated("**/"+shared.ProjectGlob+"/**/*.json", string(params.TextDocument.URI)) {
		return nil, nil
	}
	document := textdocument.Get(params.TextDocument.URI)
	offset := document.OffsetAt(params.Position)
	location := jsonc.GetLocation(document.GetText(), offset)
	if shared.IsMolangLocation(location) {
		return handlers.Molang.GetSignature(document, offset, location), nil
	}
	return nil, nil
}
