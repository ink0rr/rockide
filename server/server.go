package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ink0rr/rockide/handlers"
	"github.com/ink0rr/rockide/rockide"
	"github.com/ink0rr/rockide/textdocument"
	"github.com/rockide/protocol"
	"github.com/sourcegraph/jsonrpc2"
)

type Server struct{}

func New() *Server {
	return &Server{}
}

func (s *Server) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	var res any
	var err error
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
	// TextDocumentSync events
	case "textDocument/didOpen":
		var params protocol.DidOpenTextDocumentParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			_, err = textdocument.Open(params.TextDocument.URI)
		}
	case "textDocument/didChange":
		var params protocol.DidChangeTextDocumentParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			if len(params.ContentChanges) > 0 {
				textdocument.Update(params.TextDocument.URI, params.ContentChanges[0].Text)
				rockide.OnChange(params.TextDocument.URI)
			}
		}
	case "textDocument/didClose":
		var params protocol.DidCloseTextDocumentParams
		if err = json.Unmarshal(*req.Params, &params); err == nil {
			textdocument.Close(params.TextDocument.URI)
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

func Initialize(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	log.Printf("Process ID: %d", params.ProcessID)
	log.Printf("Connected to: %s %s", params.ClientInfo.Name, params.ClientInfo.Version)
	result := protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: protocol.TextDocumentSyncKindFull,
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
	if err = conn.Call(ctx, "window/workDoneProgress/create", &protocol.WorkDoneProgressCreateParams{Token: *token}, nil); err != nil {
		return err
	}
	progress := protocol.ProgressParams{
		Token: *token,
		Value: &protocol.WorkDoneProgressBegin{Kind: protocol.WorkDoneProgressKindBegin, Title: "Rockide: Indexing workspace"},
	}
	if err := conn.Notify(ctx, "$/progress", &progress); err != nil {
		return err
	}

	if err := rockide.IndexWorkspaces(ctx); err != nil {
		return err
	}
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
	handler := rockide.FindJsonHandler(document.URI)
	if handler == nil {
		return nil, nil
	}
	handlerParams := handlers.NewJsonHandlerParams(document, &params.Position)
	entry := handler.FindEntry(handlerParams.Location)
	if entry == nil || entry.Completions == nil {
		log.Println("Handler not found", handlerParams.Location.Path)
		return nil, nil
	}
	result := []protocol.CompletionItem{}
	for _, item := range entry.Completions(handlerParams) {
		result = append(result, protocol.CompletionItem{
			Label: item.Value,
			TextEdit: &protocol.TextEdit{
				Range: protocol.Range{
					Start: document.PositionAt(handlerParams.Node.Offset + 1),
					End:   document.PositionAt(handlerParams.Node.Offset + handlerParams.Node.Length - 1),
				},
				NewText: item.Value,
			},
		})
	}
	return result, nil
}

func Definition(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.DefinitionParams) ([]protocol.LocationLink, error) {
	document, err := textdocument.Open(params.TextDocument.URI)
	if err != nil {
		return nil, err
	}
	handler := rockide.FindJsonHandler(document.URI)
	if handler == nil {
		return nil, nil
	}
	handlerParams := handlers.NewJsonHandlerParams(document, &params.Position)
	entry := handler.FindEntry(handlerParams.Location)
	if entry == nil || entry.Definitions == nil {
		log.Println("Handler not found", handlerParams.Location.Path)
		return nil, nil
	}
	result := []protocol.LocationLink{}
	for _, item := range entry.Definitions(handlerParams) {
		if item.Value != handlerParams.Node.Value {
			continue
		}
		location := protocol.LocationLink{
			OriginSelectionRange: &protocol.Range{
				Start: document.PositionAt(handlerParams.Node.Offset + 1),
				End:   document.PositionAt(handlerParams.Node.Offset + handlerParams.Node.Length - 1),
			},
			TargetURI: item.URI,
		}
		if item.Range != nil {
			location.TargetRange = *item.Range
			location.TargetSelectionRange = *item.Range
		}
		result = append(result, location)
	}
	return result, nil
}
