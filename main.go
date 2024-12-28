package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ink0rr/rockide/rockide/handler"
	"github.com/ink0rr/rockide/rpc"
	"go.lsp.dev/protocol"
)

func main() {
	initLogger()
	log.Print("Rockide is running!")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := rpc.NewServer(ctx)
	server.Listen(func(ctx context.Context, req *rpc.RequestMessage) (res any, err error) {
		switch req.Method {
		case "initialize":
			var params protocol.InitializeParams
			if err = json.Unmarshal(req.Params, &params); err == nil {
				res, err = handler.Initialize(ctx, &params)
			}
		case "initialized":
			var params protocol.InitializedParams
			if err = json.Unmarshal(req.Params, &params); err == nil {
				err = handler.Initialized(ctx, &params)
			}
		case "textDocument/didChange":
			var params protocol.DidChangeTextDocumentParams
			if err = json.Unmarshal(req.Params, &params); err == nil {
				err = handler.TextDocumentDidChange(ctx, &params)
			}
		default:
			log.Printf("Unhandled method: %s", req.Method)
		}
		return
	})
}

func initLogger() {
	var home string
	if runtime.GOOS == "windows" {
		home = os.Getenv("UserProfile")
	} else {
		home = os.Getenv("HOME")
	}
	fileName := filepath.Join(home, ".rockide", "log.txt")
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Failed to open log file")
	}
	log.SetOutput(logFile)
	log.SetPrefix("[rockide]")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
