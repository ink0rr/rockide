package core

import (
	"reflect"

	"github.com/ink0rr/go-jsonc"
	"go.lsp.dev/uri"
)

type JsonHandler struct{}

type JsonHandlerContext struct {
	Uri      uri.URI
	Text     string
	Location *jsonc.Location
	Node     *jsonc.Node
}

func (j *JsonHandlerContext) GetParentNode() *jsonc.Node {
	root, _ := jsonc.ParseTree(j.Text, nil)
	path := j.Location.Path
	return jsonc.FindNodeAtLocation(root, path)
}

func (j *JsonHandlerContext) IsAtPropertyKeyOrArray() bool {
	return j.Location.IsAtPropertyKey || reflect.TypeOf(j.Location.Path[len(j.Location.Path)-1]).Kind() == reflect.Int
}
