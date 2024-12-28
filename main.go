package main

import (
	"bufio"
	"context"
	"encoding/json"
	"os"

	"github.com/ink0rr/rockide/rockide"
	"github.com/ink0rr/rockide/rockide/handler"
	"github.com/ink0rr/rockide/rpc"
	"go.lsp.dev/protocol"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = rockide.WithLogger(ctx)
	defer cancel()

	logger := rockide.GetLogger(ctx)
	logger.Print("Rockide is running!")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		req, err := rpc.DecodeMessage(scanner.Bytes())
		if err != nil {
			logger.Printf("Error: %s", err)
			continue
		}
		handleRequest(ctx, req)
	}

}

func handleRequest(ctx context.Context, req *rpc.RequestMessage) {
	logger := rockide.GetLogger(ctx)
	id := req.Id
	switch req.Method {
	case "initialize":
		var params protocol.InitializeParams
		if err := json.Unmarshal(req.Params, &params); err != nil {
			logger.Printf("Failed to parse contents: %s", err)
		}
		result, err := handler.Initialize(ctx, &params)
		handleResponse(ctx, id, result, err)
	case "initialized":
		var params protocol.InitializedParams
		if err := json.Unmarshal(req.Params, &params); err != nil {
			logger.Printf("Failed to parse contents: %s", err)
		}
		err := handler.Initialized(ctx, &params)
		handleResponse(ctx, id, nil, err)
	case "textDocument/didChange":
		var params protocol.DidChangeTextDocumentParams
		if err := json.Unmarshal(req.Params, &params); err != nil {
			logger.Printf("Failed to parse contents: %s", err)
		}
		err := handler.TextDocumentDidChange(ctx, &params)
		handleResponse(ctx, id, nil, err)
	default:
		logger.Printf("Unhandled method: %s", req.Method)
	}

}

func handleResponse(ctx context.Context, id int, result any, err error) {
	if err != nil {
		logger := rockide.GetLogger(ctx)
		logger.Printf("Error: %s", err)
		return
	}
	if result != nil {
		msg := rpc.ResponseMessage{Id: id, Result: &result}
		reply := rpc.EncodeMessage(msg)
		os.Stdout.Write([]byte(reply))
	}
}
