package handler

import (
	"context"

	"github.com/ink0rr/rockide/rockide"
	"go.lsp.dev/protocol"
)

func Initialize(ctx context.Context, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	logger := rockide.GetLogger(ctx)
	logger.Printf("Connected to: %s %s", params.ClientInfo.Name, params.ClientInfo.Version)
	result := protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: 1, // Full
		},
		ServerInfo: &protocol.ServerInfo{
			Name:    "rockide",
			Version: "0.0.0",
		},
	}
	return &result, nil
}

func Initialized(ctx context.Context, params *protocol.InitializedParams) error {
	server, err := rockide.NewServer()
	if err != nil {
		return err
	}
	server.Rockide.IndexWorkspaces(ctx)
	return nil
}
