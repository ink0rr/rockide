package handlers

import (
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"go.lsp.dev/uri"
)

var handlers = []*JsonHandler{
	Entity,
}

func Find(uri uri.URI) *JsonHandler {
	name := uri.Filename()
	name = strings.ReplaceAll(name, "\\", "/")
	for _, handler := range handlers {
		if doublestar.MatchUnvalidated("**/"+handler.Pattern, name) {
			return handler
		}
	}
	return nil
}
