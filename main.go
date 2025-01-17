package main

import (
	"context"

	"github.com/ink0rr/rockide/server"
	"github.com/sourcegraph/jsonrpc2"
)

func main() {
	ctx := context.Background()
	handler := jsonrpc2.HandlerWithError(server.Handle)
	stream := jsonrpc2.NewBufferedStream(&stdio{}, jsonrpc2.VSCodeObjectCodec{})
	<-jsonrpc2.NewConn(ctx, stream, jsonrpc2.AsyncHandler(handler)).DisconnectNotify()
}
