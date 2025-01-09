package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ink0rr/rockide/rockide"
	"github.com/ink0rr/rockide/textdocument"
	"github.com/rockide/protocol"
	"github.com/sourcegraph/jsonrpc2"
)

func main() {
	log.Print("Rockide is running!")
	handler := jsonrpc2.HandlerWithError(Handle)
	stream := jsonrpc2.NewBufferedStream(&stdio{}, jsonrpc2.VSCodeObjectCodec{})
	conn := jsonrpc2.NewConn(context.Background(), stream, jsonrpc2.AsyncHandler(handler))
	<-conn.DisconnectNotify()
}

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
			_, err = textdocument.Open(params.TextDocument.URI)
		}
	case "textDocument/didChange":
		var params protocol.DidChangeTextDocumentParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			if textdocument.Update(params.TextDocument.URI, params.ContentChanges) {
				rockide.OnChange(params.TextDocument.URI)
			}
		}
	case "textDocument/didSave":
		var params protocol.DidSaveTextDocumentParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			if textdocument.UpdateFull(params.TextDocument.URI, params.Text) {
				rockide.OnChange(params.TextDocument.URI)
			}
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
	case "textDocument/rename":
		var params protocol.RenameParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			res, err = Rename(ctx, conn, &params)
		}
	case "textDocument/hover":
		var params protocol.HoverParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			res, err = Hover(ctx, conn, &params)
		}
	default:
		log.Printf("Unhandled method: %s", req.Method)
	}
	return
}

func Initialize(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	log.Printf("Process ID: %d", params.ProcessID)
	log.Printf("Connected to: %s %s", params.ClientInfo.Name, params.ClientInfo.Version)

	result := protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: protocol.TextDocumentSyncKindIncremental,
			CompletionProvider: &protocol.CompletionOptions{
				TriggerCharacters: strings.Split(`0123456789abcdefghijklmnopqrstuvwxyz:.'"() `, ""),
			},
			DefinitionProvider: true,
			RenameProvider:     true,
			HoverProvider:      true,
		},
		ServerInfo: &protocol.ServerInfo{
			Name:    "rockide",
			Version: "0.0.0",
		},
	}
	return &result, nil
}

func Initialized(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.InitializedParams) error {
	configParams := protocol.ConfigurationParams{
		Items: []protocol.ConfigurationItem{{Section: "rockide.baseDir"}},
	}
	configResult := []any{}
	err := conn.Call(ctx, "workspace/configuration", &configParams, &configResult)
	if err != nil {
		return err
	}
	baseDir, ok := configResult[0].(string)
	if ok && baseDir != "" {
		rockide.SetBaseDir(baseDir)
	} else if stat, err := os.Stat("packs"); err == nil && stat.IsDir() {
		rockide.SetBaseDir("packs")
	}

	if !rockide.IsMinecraftWorkspace(ctx) {
		log.Println("Not a Minecraft workspace, exiting...")
		return nil
	}

	token := protocol.NewProgressToken(fmt.Sprintf("indexing-workspace-%d", time.Now().Unix()))
	if err := conn.Call(ctx, "window/workDoneProgress/create", &protocol.WorkDoneProgressCreateParams{Token: *token}, nil); err != nil {
		return err
	}
	progress := protocol.ProgressParams{
		Token: *token,
		Value: &protocol.WorkDoneProgressBegin{Kind: protocol.WorkDoneProgressKindBegin, Title: "Rockide: Indexing workspace"},
	}
	if err := conn.Notify(ctx, "$/progress", &progress); err != nil {
		return err
	}

	rockide.IndexWorkspaces(ctx)
	textdocument.EnableCache(true)

	progress.Value = &protocol.WorkDoneProgressEnd{Kind: protocol.WorkDoneProgressKindEnd}
	if err := conn.Notify(ctx, "$/progress", &progress); err != nil {
		return err
	}

	if err := rockide.WatchFiles(ctx); err != nil {
		return err
	}
	return nil
}

func Completion(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.CompletionParams) ([]protocol.CompletionItem, error) {
	document, err := textdocument.Open(params.TextDocument.URI)
	if err != nil {
		return nil, err
	}
	handler := rockide.FindHandler(document.URI)
	if handler == nil {
		return nil, nil
	}
	actions := handler.GetActions(document, &params.Position)
	if actions == nil || actions.Completions == nil {
		return nil, nil
	}
	return actions.Completions(), nil
}

func Definition(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.DefinitionParams) ([]protocol.LocationLink, error) {
	document, err := textdocument.Open(params.TextDocument.URI)
	if err != nil {
		return nil, err
	}
	handler := rockide.FindHandler(document.URI)
	if handler == nil {
		return nil, nil
	}
	actions := handler.GetActions(document, &params.Position)
	if actions == nil || actions.Definitions == nil {
		return nil, nil
	}
	return actions.Definitions(), nil
}

func Rename(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	document, err := textdocument.Open(params.TextDocument.URI)
	if err != nil {
		return nil, err
	}
	handler := rockide.FindHandler(document.URI)
	if handler == nil {
		return nil, nil
	}
	actions := handler.GetActions(document, &params.Position)
	if actions == nil || actions.Definitions == nil {
		return nil, nil
	}
	return actions.Rename(params.NewName), nil
}

func Hover(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.HoverParams) (*protocol.Hover, error) {
	return nil, nil
}
