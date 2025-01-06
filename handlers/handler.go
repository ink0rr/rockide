package handlers

import (
	"github.com/ink0rr/rockide/textdocument"
	"github.com/rockide/protocol"
)

type Handler interface {
	GetPattern() string
	GetActions(document *textdocument.TextDocument, position *protocol.Position) *HandlerActions
}

type HandlerActions struct {
	Completions func() []protocol.CompletionItem
	Definitions func() []protocol.LocationLink
	Rename      func(newName string) *protocol.WorkspaceEdit
}
