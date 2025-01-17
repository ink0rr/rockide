package server

import (
	"github.com/bmatcuk/doublestar/v4"
	"github.com/ink0rr/rockide/handlers"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/ink0rr/rockide/stores"
)

func findActions(document *textdocument.TextDocument, position protocol.Position) *handlers.HandlerActions {
	for _, handler := range handlerList {
		if doublestar.MatchUnvalidated("**/"+handler.GetPattern(&project), string(document.URI)) {
			return handler.GetActions(document, position)
		}
	}
	return nil
}

func findStore(uri protocol.DocumentURI) stores.Store {
	for _, store := range storeList {
		if doublestar.MatchUnvalidated("**/"+store.GetPattern(&project), string(uri)) {
			return store
		}
	}
	return nil
}
