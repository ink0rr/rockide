package server

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/sourcegraph/jsonrpc2"
)

func Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (res any, err error) {
	switch req.Method {
	case "initialize":
		var params protocol.InitializeParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			res, err = Initialize(ctx, conn, &params)
		}
	case "initialized":
		var params protocol.InitializedParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			err = Initialized(ctx, conn, &params)
		}

	case "textDocument/didOpen":
		var params protocol.DidOpenTextDocumentParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			textdocument.Open(params.TextDocument.URI, params.TextDocument.Text)
		}
	case "textDocument/didChange":
		var params protocol.DidChangeTextDocumentParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			textdocument.SyncIncremental(params.TextDocument.URI, params.ContentChanges)
		}
	case "textDocument/didSave":
		var params protocol.DidSaveTextDocumentParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			textdocument.SyncFull(params.TextDocument.URI, params.Text)
		}
	case "textDocument/didClose":
		var params protocol.DidCloseTextDocumentParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			textdocument.Close(params.TextDocument.URI)
		}

	case "textDocument/completion":
		var params protocol.CompletionParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			res, err = Completion(ctx, conn, &params)
		}
	case "textDocument/definition":
		var params protocol.DefinitionParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			res, err = Definition(ctx, conn, &params)
		}
	case "textDocument/prepareRename":
		var params protocol.PrepareRenameParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			res, err = PrepareRename(ctx, conn, &params)
		}
	case "textDocument/rename":
		var params protocol.RenameParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			res, err = Rename(ctx, conn, &params)
		}
	default:
		log.Printf("Unhandled method: %s", req.Method)
	}
	return res, err
}
