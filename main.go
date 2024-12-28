package main

import (
	"bufio"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	initLogger()
	log.Print("Rockide is running!")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		req, err := rpc.DecodeMessage(scanner.Bytes())
		if err != nil {
			log.Printf("Error: %s", err)
			continue
		}
		handleRequest(ctx, req)
	}
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

func handleRequest(ctx context.Context, req *rpc.RequestMessage) {
	id := req.Id
	switch req.Method {
	case "initialize":
		var params protocol.InitializeParams
		if err := json.Unmarshal(req.Params, &params); err != nil {
			log.Printf("Failed to parse contents: %s", err)
		}
		result, err := handler.Initialize(ctx, &params)
		handleResponse(id, result, err)
	case "initialized":
		var params protocol.InitializedParams
		if err := json.Unmarshal(req.Params, &params); err != nil {
			log.Printf("Failed to parse contents: %s", err)
		}
		err := handler.Initialized(ctx, &params)
		handleResponse(id, nil, err)
	case "textDocument/didChange":
		var params protocol.DidChangeTextDocumentParams
		if err := json.Unmarshal(req.Params, &params); err != nil {
			log.Printf("Failed to parse contents: %s", err)
		}
		err := handler.TextDocumentDidChange(ctx, &params)
		handleResponse(id, nil, err)
	default:
		log.Printf("Unhandled method: %s", req.Method)
	}

}

func handleResponse(id int, result any, err error) {
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}
	if result != nil {
		msg := rpc.ResponseMessage{Id: id, Result: &result}
		reply := rpc.EncodeMessage(msg)
		os.Stdout.Write([]byte(reply))
	}
}
