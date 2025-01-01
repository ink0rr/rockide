package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/ink0rr/go-jsonc"
	"github.com/ink0rr/rockide/rockide"
	"github.com/ink0rr/rockide/textdocument"
	"github.com/rockide/protocol"
	"github.com/sourcegraph/jsonrpc2"
)

func main() {
	log.Print("Rockide is running!")

	ctx := context.Background()
	stream := jsonrpc2.NewBufferedStream(&stdio{}, jsonrpc2.VSCodeObjectCodec{})
	conn := jsonrpc2.NewConn(ctx, stream, jsonrpc2.AsyncHandler(&handler{}))
	<-conn.DisconnectNotify()
}

type handler struct{}

func (h *handler) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	var res any
	var err error
	switch req.Method {
	case "initialize":
		var params protocol.InitializeParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			res, err = Initialize(ctx, &params)
		}
	case "initialized":
		var params protocol.InitializedParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			err = Initialized(ctx, &params)
		}
	case "textDocument/didChange":
		var params protocol.DidChangeTextDocumentParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			err = TextDocumentDidChange(ctx, &params)
		}
	case "textDocument/completion":
		var params protocol.CompletionParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			err = Completion(ctx, &params)
		}
	default:
		log.Printf("Unhandled method: %s", req.Method)
	}
	if err != nil {
		log.Printf("Replying with error: %s", err)
		conn.ReplyWithError(ctx, req.ID, &jsonrpc2.Error{Code: jsonrpc2.CodeInternalError, Message: err.Error()})
		return
	}
	err = conn.Reply(ctx, req.ID, res)
	if err != nil {
		log.Printf("Failed to send reply: %s", err)
	}
}

func Initialize(ctx context.Context, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	log.Printf("Connected to: %s %s", params.ClientInfo.Name, params.ClientInfo.Version)
	result := protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: 1,
			CompletionProvider: &protocol.CompletionOptions{
				TriggerCharacters: strings.Split(`0123456789abcdefghijklmnopqrstuvwxyz:.,'"() `, ""),
			},
			DefinitionProvider: &protocol.DefinitionOptions{},
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
	if err := rockide.IndexWorkspaces(ctx); err != nil {
		return err
	}
	if err := rockide.WatchFiles(ctx); err != nil {
		return err
	}
	return nil
}

func TextDocumentDidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	if len(params.ContentChanges) > 0 {
		rockide.OnChange(params.TextDocument.URI)
	}
	return nil
}

func Completion(ctx context.Context, params *protocol.CompletionParams) error {
	document, err := textdocument.New(params.TextDocument.URI)
	if err != nil {
		return err
	}
	location := jsonc.GetLocation(document.GetText(), int(document.OffsetAt(params.Position)))
	_ = location
	return nil
}
