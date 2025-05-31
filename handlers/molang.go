package handlers

import (
	"log"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/molang"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/protocol/semtok"
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/internal/textdocument"
)

type MolangContext struct {
	document    *textdocument.TextDocument
	text        string
	startOffset uint32
	offset      uint32
}

func NewMolangContext(document *textdocument.TextDocument, location *jsonc.Location, offset uint32) *MolangContext {
	node := location.PreviousNode
	if node == nil {
		return nil
	}
	nodeValue, ok := node.Value.(string)
	if !ok {
		return nil
	}
	startOffset := node.Offset
	return &MolangContext{
		document:    document,
		text:        nodeValue,
		startOffset: startOffset,
		offset:      offset - startOffset - 2,
	}
}

func MolangCompletions(ctx *MolangContext) []protocol.CompletionItem {
	parser, err := molang.NewParser(ctx.text)
	if err != nil {
		log.Printf("Molang error: %v", err)
		return nil
	}
	index := parser.FindIndex(ctx.offset)
	if index == -1 {
		return nil
	}

	token := parser.Tokens[index]
	switch token.Kind {
	case molang.KindString:
		methodCall := parser.GetMethodCall(ctx.offset)
		if methodCall == nil {
			return nil
		}
		method, ok := molang.GetMethod(methodCall.Prefix, methodCall.Name)
		if !ok {
			return nil
		}
		params := method.Signature.GetParams()
		param := params[len(params)-1]
		if methodCall.ParamIndex < len(params) {
			param = params[methodCall.ParamIndex]
		}
		getTypeValues := molangTypes[param.Type]
		if getTypeValues == nil {
			log.Printf("Unknown param tokenType: %s", param.Type)
			return nil
		}
		values := getTypeValues()
		editRange := protocol.Range{
			Start: ctx.document.PositionAt(ctx.startOffset + token.Offset + 2),
			End:   ctx.document.PositionAt(ctx.startOffset + token.Offset + token.Length),
		}
		res := []protocol.CompletionItem{}
		if values.strings != nil {
			res = sliceutil.Map(values.strings, func(value string) protocol.CompletionItem {
				return protocol.CompletionItem{
					Label: value,
					TextEdit: &protocol.Or_CompletionItem_textEdit{
						Value: protocol.TextEdit{
							NewText: value,
							Range:   editRange,
						},
					},
				}
			})
		}
		set := mapset.NewThreadUnsafeSet[string]()
		for _, ref := range values.references {
			if set.ContainsOne(ref.Value) {
				continue
			}
			set.Add(ref.Value)
			res = append(res, protocol.CompletionItem{
				Label: ref.Value,
				TextEdit: &protocol.Or_CompletionItem_textEdit{
					Value: protocol.TextEdit{
						NewText: ref.Value,
						Range:   editRange,
					},
				},
			})
		}
		return res
	case molang.KindPrefix, molang.KindUnknown:
		editRange := protocol.Range{
			Start: ctx.document.PositionAt(ctx.startOffset + token.Offset + 1),
			End:   ctx.document.PositionAt(ctx.startOffset + token.Offset + token.Length + 1),
		}
		return sliceutil.Map(molang.Prefixes, func(value string) protocol.CompletionItem {
			return protocol.CompletionItem{
				Label: value,
				TextEdit: &protocol.Or_CompletionItem_textEdit{
					Value: protocol.TextEdit{
						NewText: value,
						Range:   editRange,
					},
				},
			}
		})
	}

	if index == 0 {
		return nil
	}

	prefix := parser.Tokens[index-1]
	if prefix.Kind != molang.KindPrefix || token.Kind != molang.KindMethod || strings.LastIndex(token.Value, ".") != 0 {
		return nil
	}

	return sliceutil.Map(molang.GetMethodList(prefix.Value), func(method molang.Method) protocol.CompletionItem {
		return protocol.CompletionItem{
			Label:  prefix.Value + "." + method.Name,
			Kind:   protocol.MethodCompletion,
			Detail: method.Name + string(method.Signature),
			Documentation: &protocol.Or_CompletionItem_documentation{
				Value: method.Description,
			},
			Deprecated: method.Deprecated,
		}
	})
}

func MolangDefinitions(ctx *MolangContext) []protocol.LocationLink {
	parser, err := molang.NewParser(ctx.text)
	if err != nil {
		log.Printf("Molang error: %v", err)
		return nil
	}
	index := parser.FindIndex(ctx.offset)
	if index == -1 {
		return nil
	}
	token := parser.Tokens[index]
	methodCall := parser.GetMethodCall(ctx.offset)
	if token.Kind != molang.KindString || methodCall == nil {
		return nil
	}
	method, ok := molang.GetMethod(methodCall.Prefix, methodCall.Name)
	if !ok {
		return nil
	}
	params := method.Signature.GetParams()
	param := params[len(params)-1]
	if methodCall.ParamIndex < len(params) {
		param = params[methodCall.ParamIndex]
	}
	getTypeValues := molangTypes[param.Type]
	if getTypeValues == nil {
		log.Printf("Unknown param tokenType: %s", param.Type)
		return nil
	}
	res := []protocol.LocationLink{}
	values := getTypeValues()
	if values.references == nil {
		return nil
	}
	selectionRange := protocol.Range{
		Start: ctx.document.PositionAt(ctx.startOffset + token.Offset + 1),
		End:   ctx.document.PositionAt(ctx.startOffset + token.Offset + token.Length + 1),
	}
	molangValue := token.Value[1 : len(token.Value)-1] // Exclude quotes
	for _, ref := range values.references {
		if ref.Value != molangValue {
			continue
		}
		location := protocol.LocationLink{
			OriginSelectionRange: &selectionRange,
			TargetURI:            ref.URI,
		}
		if ref.Range != nil {
			location.TargetRange = *ref.Range
			location.TargetSelectionRange = *ref.Range
		}
		res = append(res, location)
	}
	return res
}

func MolangHover(ctx *MolangContext) *protocol.Hover {
	parser, err := molang.NewParser(ctx.text)
	if err != nil {
		log.Printf("Molang error: %v", err)
		return nil
	}
	index := parser.FindIndex(ctx.offset + 1)
	if index < 0 {
		return nil
	}
	var prefix molang.Token
	var method molang.Method
	var ok bool
	token := parser.Tokens[index]
	switch token.Kind {
	case molang.KindPrefix:
		if index+1 > len(parser.Tokens) {
			return nil
		}
		prefix = token
		method, ok = molang.GetMethod(prefix.Value, parser.Tokens[index+1].Value)
	case molang.KindMethod:
		if index == 0 {
			return nil
		}
		prefix = parser.Tokens[index-1]
		method, ok = molang.GetMethod(prefix.Value, parser.Tokens[index].Value)
	default:
		return nil
	}
	if !ok {
		return nil
	}
	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind: protocol.Markdown,
			Value: "```rockide-molang\n" +
				prefix.Value + "." + method.Name + string(method.Signature) +
				"\n```\n" +
				method.Description,
		},
	}
}

func MolangSignatureHelp(ctx *MolangContext) *protocol.SignatureHelp {
	parser, err := molang.NewParser(ctx.text)
	if err != nil {
		log.Printf("Molang error: %v", err)
		return nil
	}
	methodCall := parser.GetMethodCall(ctx.offset)
	if methodCall == nil {
		return nil
	}
	method, ok := molang.GetMethod(methodCall.Prefix, methodCall.Name)
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
			return protocol.ParameterInformation{Label: param.ToString()}
		}),
		ActiveParameter: uint32(activeParam),
	}
	return &protocol.SignatureHelp{
		Signatures: []protocol.SignatureInformation{signature},
	}
}

func MolangSemanticTokens(text string, startLine, startCharacter uint32) []semtok.Token {
	parser, err := molang.NewParser(text)
	if err != nil {
		log.Printf("Molang error: %v", err)
		return nil
	}
	res := []semtok.Token{}
	for _, token := range parser.Tokens {
		tokenType, ok := molangTokenMap[token.Kind]
		if !ok {
			continue
		}
		res = append(res,
			semtok.Token{
				Type:  tokenType,
				Line:  startLine,
				Start: startCharacter + token.Offset + 1,
				Len:   token.Length,
			})
	}
	return res
}
