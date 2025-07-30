package server

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/protocol/semtok"
	"github.com/ink0rr/rockide/shared"
	"github.com/sourcegraph/jsonrpc2"
)

func Initialize(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	log.Printf("Process ID: %d", params.ProcessID)
	log.Printf("Connected to: %s %s", params.ClientInfo.Name, params.ClientInfo.Version)

	project, err := findProjectPaths(params.InitializationOptions)
	if err != nil {
		log.Println(err)
	} else {
		shared.SetProject(project)
	}
	triggerCharacters := strings.Split(`0123456789abcdefghijklmnopqrstuvwxyz.'"() `, "")
	result := protocol.InitializeResult{
		ServerInfo: &protocol.ServerInfo{
			Name:    "rockide",
			Version: "0.0.0",
		},
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: protocol.Incremental,
			CompletionProvider: &protocol.CompletionOptions{
				TriggerCharacters: triggerCharacters,
			},
			DefinitionProvider: &protocol.Or_ServerCapabilities_definitionProvider{Value: true},
			RenameProvider: &protocol.RenameOptions{
				PrepareProvider: true,
			},
			HoverProvider: &protocol.Or_ServerCapabilities_hoverProvider{Value: true},
			SignatureHelpProvider: &protocol.SignatureHelpOptions{
				TriggerCharacters: triggerCharacters,
			},
			SemanticTokensProvider: protocol.SemanticTokensOptions{
				Legend: protocol.SemanticTokensLegend{
					TokenTypes:     semtok.TokenTypes,
					TokenModifiers: []semtok.Modifier{},
				},
				Full: &protocol.Or_SemanticTokensOptions_full{Value: true},
			},
		},
	}
	return &result, nil
}

func Initialized(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.InitializedParams) error {
	project := shared.GetProject()
	if project == nil {
		return conn.Notify(ctx, "shutdown", nil)
	}
	fileWatcher := protocol.Registration{
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
	if err := registerCapability(ctx, conn, []protocol.Registration{fileWatcher}); err != nil {
		return err
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

func registerCapability(ctx context.Context, conn *jsonrpc2.Conn, registrations []protocol.Registration) error {
	var result any
	conn.Call(ctx, "client/registerCapability", protocol.RegistrationParams{Registrations: registrations}, &result)
	if result != nil {
		return fmt.Errorf("%v", result)
	}
	return nil
}
