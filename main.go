package main

import (
	"context"
	"log"

	"github.com/ink0rr/rockide/server"
	"github.com/sourcegraph/jsonrpc2"
)

func main() {
	log.Print("Rockide is running!")

	ctx := context.Background()
	server := server.New()
	stream := jsonrpc2.NewBufferedStream(&stdio{}, jsonrpc2.VSCodeObjectCodec{})
	conn := jsonrpc2.NewConn(ctx, stream, jsonrpc2.AsyncHandler(server))
	<-conn.DisconnectNotify()
}
