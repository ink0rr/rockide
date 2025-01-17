package server

import (
	"context"

	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/sourcegraph/jsonrpc2"
)

func Definition(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.DefinitionParams) ([]protocol.LocationLink, error) {
	document := textdocument.Get(params.TextDocument.URI)
	actions := findActions(document, params.Position)
	if actions == nil || actions.Definitions == nil {
		return nil, nil
	}
	return actions.Definitions(), nil
}
