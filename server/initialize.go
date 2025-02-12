package server

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/shared"
	"github.com/sourcegraph/jsonrpc2"
)

func Initialize(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	log.Printf("Process ID: %d", params.ProcessID)
	log.Printf("Connected to: %s %s", params.ClientInfo.Name, params.ClientInfo.Version)

	if err := findProjectPaths(params.InitializationOptions); err != nil {
		return nil, err
	}

	result := protocol.InitializeResult{
		ServerInfo: &protocol.ServerInfo{
			Name:    "rockide",
			Version: "0.0.0",
		},
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: protocol.Incremental,
			CompletionProvider: &protocol.CompletionOptions{
				TriggerCharacters: strings.Split(`0123456789abcdefghijklmnopqrstuvwxyz.'"() `, ""),
			},
			DefinitionProvider: &protocol.Or_ServerCapabilities_definitionProvider{Value: true},
			RenameProvider: &protocol.RenameOptions{
				PrepareProvider: true,
			},
			HoverProvider: &protocol.Or_ServerCapabilities_hoverProvider{Value: true},
		},
	}
	return &result, nil
}

func Initialized(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.InitializedParams) error {
	project := shared.GetProject()
	registration := protocol.Registration{
		ID:     "fileWatcher",
		Method: "workspace/didChangeWatchedFiles",
		RegisterOptions: protocol.DidChangeWatchedFilesRegistrationOptions{
			Watchers: []protocol.FileSystemWatcher{{
				GlobPattern: protocol.GlobPattern{Value: protocol.RelativePattern{
					BaseURI: protocol.URIFromPath(shared.Getwd()),
					Pattern: fmt.Sprintf("{%s,%s}/**/*", project.BP, project.RP),
				}},
			}},
		},
	}
	var registrationError any
	conn.Call(ctx, "client/registerCapability", protocol.RegistrationParams{
		Registrations: []protocol.Registration{registration},
	}, &registrationError)
	if registrationError != nil {
		return fmt.Errorf("%v", registrationError)
	}

	token := protocol.ProgressToken(fmt.Sprintf("indexing-workspace-%d", time.Now().Unix()))
	if err := conn.Call(ctx, "window/workDoneProgress/create", &protocol.WorkDoneProgressCreateParams{Token: token}, nil); err != nil {
		return err
	}
	progress := protocol.ProgressParams{
		Token: token,
		Value: &protocol.WorkDoneProgressBegin{Kind: "begin", Title: "Rockide: Indexing workspace"},
	}
	if err := conn.Notify(ctx, "$/progress", &progress); err != nil {
		return err
	}

	indexWorkspace()

	progress.Value = &protocol.WorkDoneProgressEnd{Kind: "end"}
	if err := conn.Notify(ctx, "$/progress", &progress); err != nil {
		return err
	}

	return nil
}
