package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/ink0rr/rockide/server"
	"github.com/sourcegraph/jsonrpc2"
)

var version = "dev"

func main() {
	printVersion := flag.Bool("version", false, "Print version")
	flag.Parse()
	if *printVersion {
		fmt.Println(version)
		return
	}

	ctx := context.Background()
	handler := jsonrpc2.HandlerWithError(server.Handle)
	stream := jsonrpc2.NewBufferedStream(&stdio{}, jsonrpc2.VSCodeObjectCodec{})
	<-jsonrpc2.NewConn(ctx, stream, jsonrpc2.AsyncHandler(handler)).DisconnectNotify()
}
