package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/ink0rr/rockide/rockide"
	"github.com/ink0rr/rockide/rpc"
	"go.lsp.dev/protocol"
)

func main() {
	log.Print("Rockide is running!")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := rpc.NewServer(ctx)
	server.Listen(func(ctx context.Context, req *rpc.RequestMessage) (res any, err error) {
		switch req.Method {
		case "initialize":
			var params protocol.InitializeParams
			if err = json.Unmarshal(req.Params, &params); err == nil {
				res, err = Initialize(ctx, &params)
			}
		case "initialized":
			var params protocol.InitializedParams
			if err = json.Unmarshal(req.Params, &params); err == nil {
				err = Initialized(ctx, &params)
			}
		case "textDocument/didChange":
			var params protocol.DidChangeTextDocumentParams
			if err = json.Unmarshal(req.Params, &params); err == nil {
				err = TextDocumentDidChange(ctx, &params)
			}
		default:
			log.Printf("Unhandled method: %s", req.Method)
		}
		return
	})
}

func Initialize(ctx context.Context, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	log.Printf("Connected to: %s %s", params.ClientInfo.Name, params.ClientInfo.Version)
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
	stat, err := os.Stat("packs")
	if err == nil && stat.IsDir() {
		rockide.SetBaseDir("packs")
	}
	rockide.IndexWorkspaces(ctx)
	return nil
}

func TextDocumentDidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	if len(params.ContentChanges) > 0 {
		rockide.OnChange(params.TextDocument.URI)
	}
	return nil
}
