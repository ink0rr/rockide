package handlers

import (
	"log"
	"strings"

	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/molang"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/internal/textdocument"
)

type MolangHandler struct{}

func (m *MolangHandler) GetActions(document *textdocument.TextDocument, offset uint32, location *jsonc.Location) *HandlerActions {
	node := location.PreviousNode
	nodeValue, ok := node.Value.(string)
	if !ok {
		return nil
	}
	molangOffset := int(offset - node.Offset - 2) // -2 to offset quotes
	parser, err := molang.NewParser(nodeValue)
	if err != nil {
		log.Println(err)
		return nil
	}
	index := parser.FindIndex(molangOffset)
	if index == -1 {
		return nil
	}

	token := parser.Tokens[index]
	methodCall := parser.GetMethodCall(molangOffset)
	if token.Kind == molang.STRING && methodCall != nil {
		method, ok := sliceutil.Find(molang.GetMethodList(methodCall.Prefix), func(method molang.Method) bool { return method.Name == methodCall.Name })
		if !ok {
			return nil
		}
		params := method.Signature.GetParams()
		param := params[len(params)-1]
		if methodCall.ParamIndex < len(params) {
			param = params[methodCall.ParamIndex]
		}
		getValues := molangLiterals[param.Type]
		if getValues == nil {
			log.Printf("Unknown param tokenType: %s", param.Type)
			return nil
		}
		return &HandlerActions{
			Completions: func() []protocol.CompletionItem {
				res := []protocol.CompletionItem{}
				set := make(map[string]bool)
				for _, value := range getValues() {
					if set[value] {
						continue
					}
					set[value] = true
					res = append(res, protocol.CompletionItem{Label: value})
				}
				return res
			},
		}
	}

	if token.Kind == molang.PREFIX {
		return &HandlerActions{
			Completions: func() []protocol.CompletionItem {
				return sliceutil.Map(molang.Prefixes, func(value string) protocol.CompletionItem {
					return protocol.CompletionItem{Label: value}
				})
			},
		}
	}

	prefix := parser.Tokens[index-1]
	if token.Kind != molang.METHOD || prefix.Kind != molang.PREFIX || strings.LastIndex(token.Value, ".") != 0 {
		return nil
	}

	return &HandlerActions{
		Completions: func() []protocol.CompletionItem {
			return sliceutil.Map(molang.GetMethodList(prefix.Value), func(method molang.Method) protocol.CompletionItem {
				return protocol.CompletionItem{
					Label:  prefix.Value + "." + method.Name,
					Kind:   protocol.MethodCompletion,
					Detail: method.Name + string(method.Signature),
					Documentation: &protocol.Or_CompletionItem_documentation{
						Value: method.Description,
					},
				}
			})
		},
	}
}

func (m *MolangHandler) GetSignature(document *textdocument.TextDocument, offset uint32, location *jsonc.Location) *protocol.SignatureHelp {
	node := location.PreviousNode
	nodeValue, ok := node.Value.(string)
	if !ok {
		return nil
	}
	molangOffset := int(offset - node.Offset - 2) // -2 to offset quotes
	parser, err := molang.NewParser(nodeValue)
	if err != nil {
		log.Println(err)
		return nil
	}
	methodCall := parser.GetMethodCall(molangOffset)
	if methodCall == nil {
		return nil
	}
	method, ok := sliceutil.Find(molang.GetMethodList(methodCall.Prefix), func(method molang.Method) bool { return method.Name == methodCall.Name })
	if !ok {
		return nil
	}
	params := method.Signature.GetParams()
	activeParam := methodCall.ParamIndex
	if lastParam := params[len(params)-1]; strings.HasPrefix(lastParam.Label, "...") {
		activeParam = min(activeParam, len(params)-1)
	}
	signature := protocol.SignatureInformation{
		Label: methodCall.Prefix + "." + method.Name + string(method.Signature),
		Documentation: &protocol.Or_SignatureInformation_documentation{
			Value: method.Description,
		},
		Parameters: sliceutil.Map(params, func(param molang.Parameter) protocol.ParameterInformation {
			return protocol.ParameterInformation{Label: param.Label}
		}),
		ActiveParameter: uint32(activeParam),
	}
	return &protocol.SignatureHelp{
		Signatures: []protocol.SignatureInformation{signature},
	}
}

var Molang MolangHandler
