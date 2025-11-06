package handlers

import (
	"log"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ink0rr/rockide/internal/molang"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/protocol/semtok"
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/internal/textdocument"
)

type MolangHandler struct{}

func (m *MolangHandler) Completions(document *textdocument.TextDocument, position protocol.Position) []protocol.CompletionItem {
	parser, err := molang.NewParser(document.GetText())
	if err != nil {
		log.Printf("Molang error: %v", err)
		return nil
	}
	offset := document.OffsetAt(position)
	index := parser.FindIndex(offset - 1)
	if index == -1 {
		editRange := protocol.Range{
			Start: position,
			End:   position,
		}
		return sliceutil.Map(molang.Prefixes, func(value string) protocol.CompletionItem {
			return protocol.CompletionItem{
				Label: value,
				Kind:  protocol.ClassCompletion,
				TextEdit: &protocol.Or_CompletionItem_textEdit{
					Value: protocol.TextEdit{
						NewText: value,
						Range:   editRange,
					},
				},
			}
		})
	}

	token := parser.Tokens[index]
	switch token.Kind {
	case molang.KindString:
		methodCall := parser.GetMethodCall(offset)
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
			// Exclude surrounding single quotes
			Start: document.PositionAt(token.Offset + 1),
			End:   document.PositionAt(token.Offset + token.Length - 1),
		}
		res := []protocol.CompletionItem{}
		if values.literals != nil {
			res = sliceutil.Map(values.literals, func(value string) protocol.CompletionItem {
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
		for _, binding := range values.bindings {
			for _, ref := range binding.Source.Get() {
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
		}
		return res
	case molang.KindPrefix, molang.KindUnknown:
		editRange := protocol.Range{
			Start: document.PositionAt(token.Offset),
			End:   document.PositionAt(token.Offset + token.Length),
		}
		return sliceutil.Map(molang.Prefixes, func(value string) protocol.CompletionItem {
			return protocol.CompletionItem{
				Label: value,
				Kind:  protocol.ClassCompletion,
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

	editRange := protocol.Range{
		Start: document.PositionAt(prefix.Offset),
		End:   document.PositionAt(token.Offset + token.Length),
	}
	return sliceutil.Map(molang.GetMethodList(prefix.Value), func(method molang.Method) protocol.CompletionItem {
		value := prefix.Value + "." + method.Name
		return protocol.CompletionItem{
			Label: value,
			Kind:  protocol.MethodCompletion,
			TextEdit: &protocol.Or_CompletionItem_textEdit{
				Value: protocol.TextEdit{
					Range:   editRange,
					NewText: value,
				},
			},
			Detail: method.Name + string(method.Signature),
			Documentation: &protocol.Or_CompletionItem_documentation{
				Value: method.Description,
			},
			Deprecated: method.Deprecated,
		}
	})
}

func (m *MolangHandler) Definitions(document *textdocument.TextDocument, position protocol.Position) []protocol.LocationLink {
	parser, err := molang.NewParser(document.GetText())
	if err != nil {
		log.Printf("Molang error: %v", err)
		return nil
	}
	offset := document.OffsetAt(position)
	index := parser.FindIndex(offset)
	if index == -1 {
		return nil
	}

	token := parser.Tokens[index]
	methodCall := parser.GetMethodCall(offset)
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
	if values.bindings == nil {
		return nil
	}
	selectionRange := protocol.Range{
		Start: document.PositionAt(token.Offset),
		End:   document.PositionAt(token.Offset + token.Length),
	}
	molangValue := token.Value[1 : len(token.Value)-1] // Exclude surrounding single quotes
	for _, binding := range values.bindings {
		for _, ref := range binding.Source.Get() {
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
	}
	return res
}

func (m *MolangHandler) Hover(document *textdocument.TextDocument, position protocol.Position) *protocol.Hover {
	parser, err := molang.NewParser(document.GetText())
	if err != nil {
		log.Printf("Molang error: %v", err)
		return nil
	}
	offset := document.OffsetAt(position)
	index := parser.FindIndex(offset)
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

func (m *MolangHandler) SignatureHelp(document *textdocument.TextDocument, position protocol.Position) *protocol.SignatureHelp {
	parser, err := molang.NewParser(document.GetText())
	if err != nil {
		log.Printf("Molang error: %v", err)
		return nil
	}
	offset := document.OffsetAt(position)
	methodCall := parser.GetMethodCall(offset - 1)
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

func (m *MolangHandler) ComputeSemanticTokens(document *textdocument.TextDocument) []semtok.Token {
	parser, err := molang.NewParser(document.GetText())
	if err != nil {
		log.Printf("Molang error: %v", err)
		return nil
	}

	tokens := []semtok.Token{}
	for _, token := range parser.Tokens {
		tokenType, ok := molangTokenMap[token.Kind]
		if !ok {
			continue
		}
		position := document.PositionAt(token.Offset)
		tokens = append(tokens,
			semtok.Token{
				Type:  tokenType,
				Line:  position.Line,
				Start: position.Character,
				Len:   token.Length,
			})
	}

	return tokens
}

func (m *MolangHandler) SemanticTokens(document *textdocument.TextDocument) *protocol.SemanticTokens {
	tokens := m.ComputeSemanticTokens(document)

	return &protocol.SemanticTokens{
		Data: semtok.Encode(tokens, tokenType, tokenModifier),
	}
}

var Molang = &MolangHandler{}
