package server

import (
	"context"

	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/sourcegraph/jsonrpc2"
)

type PrepareRenameResult struct {
	DefaultBehavior bool `json:"defaultBehavior"`
}

func PrepareRename(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.PrepareRenameParams) (*PrepareRenameResult, error) {
	document := textdocument.Get(params.TextDocument.URI)
	actions := findActions(document, params.Position)
	if actions == nil || actions.Rename == nil {
		return nil, nil
	}
	return &PrepareRenameResult{DefaultBehavior: true}, nil
}

func Rename(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	document := textdocument.Get(params.TextDocument.URI)
	actions := findActions(document, params.Position)
	if actions == nil || actions.Rename == nil {
		return nil, nil
	}
	return actions.Rename(params.NewName), nil
}
