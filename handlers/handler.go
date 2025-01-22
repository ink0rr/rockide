package handlers

import (
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
)

type Handler interface {
	Pattern() string
	GetActions(document *textdocument.TextDocument, position protocol.Position) *HandlerActions
}

type HandlerActions struct {
	Completions func() []protocol.CompletionItem
	Definitions func() []protocol.LocationLink
	Rename      func(newName string) *protocol.WorkspaceEdit
}
