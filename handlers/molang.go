package handlers

import (
	"log"
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ink0rr/rockide/internal/jsonc"
	"github.com/ink0rr/rockide/internal/molang"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/protocol/semtok"
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/ink0rr/rockide/shared"
)

type MolangHandler struct{}

func (m *MolangHandler) GetActions(document *textdocument.TextDocument, offset uint32, location *jsonc.Location) *HandlerActions {
	node := location.PreviousNode
	nodeValue, ok := node.Value.(string)
	if !ok {
		return nil
	}
	molangOffset := offset - node.Offset - 2 // -2 to account for open quote and caret position
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
	if token.Kind == molang.KindString && methodCall != nil {
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
		return &HandlerActions{
			Completions: func() []protocol.CompletionItem {
				res := []protocol.CompletionItem{}
				values := getTypeValues()
				editRange := protocol.Range{
					Start: document.PositionAt(node.Offset + token.Offset + 2),
					End:   document.PositionAt(node.Offset + token.Offset + token.Length),
				}
				if values.references == nil {
					for _, value := range values.strings {
						res = append(res, protocol.CompletionItem{
							Label: value,
							TextEdit: &protocol.Or_CompletionItem_textEdit{
								Value: protocol.TextEdit{
									NewText: value,
									Range:   editRange,
								},
							},
						})
					}
				} else {
					set := mapset.NewThreadUnsafeSet[string]()
					for _, ref := range values.references {
						if set.Contains(ref.Value) {
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
			},
			Definitions: func() []protocol.LocationLink {
				res := []protocol.LocationLink{}
				values := getTypeValues()
				if values.references == nil {
					return nil
				}
				selectionRange := protocol.Range{
					Start: document.PositionAt(node.Offset + token.Offset + 1),
					End:   document.PositionAt(node.Offset + token.Offset + token.Length + 1),
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
			},
		}
	}

	if token.Kind == molang.KindPrefix || token.Kind == molang.KindUnknown {
		return &HandlerActions{
			Completions: func() []protocol.CompletionItem {
				editRange := protocol.Range{
					Start: document.PositionAt(node.Offset + token.Offset + 1),
					End:   document.PositionAt(node.Offset + token.Offset + token.Length + 1),
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
			},
		}
	}

	if index == 0 {
		return nil
	}

	prefix := parser.Tokens[index-1]
	if prefix.Kind != molang.KindPrefix || token.Kind != molang.KindMethod || strings.LastIndex(token.Value, ".") != 0 {
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
					Deprecated: method.Deprecated,
				}
			})
		},
	}
}

func (m *MolangHandler) GetHover(document *textdocument.TextDocument, offset uint32, location *jsonc.Location) *protocol.Hover {
	node := location.PreviousNode
	nodeValue, ok := node.Value.(string)
	if !ok {
		return nil
	}
	molangOffset := offset - node.Offset - 1 // -1 to account for open quote
	parser, err := molang.NewParser(nodeValue)
	if err != nil {
		log.Println(err)
		return nil
	}
	index := parser.FindIndex(molangOffset)
	if index < 0 {
		return nil
	}
	var prefix molang.Token
	var method molang.Method
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

func (m *MolangHandler) GetSignature(document *textdocument.TextDocument, offset uint32, location *jsonc.Location) *protocol.SignatureHelp {
	node := location.PreviousNode
	nodeValue, ok := node.Value.(string)
	if !ok {
		return nil
	}
	molangOffset := offset - node.Offset - 2 // -2 to offset quotes
	parser, err := molang.NewParser(nodeValue)
	if err != nil {
		log.Println(err)
		return nil
	}
	methodCall := parser.GetMethodCall(molangOffset)
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

var tokenMap = map[molang.TokenKind]semtok.Type{
	molang.KindNumber:     semtok.TokNumber,
	molang.KindString:     semtok.TokString,
	molang.KindMacro:      semtok.TokMacro,
	molang.KindMethod:     semtok.TokMethod,
	molang.KindPrefix:     semtok.TokType,
	molang.KindKeyword:    semtok.TokKeyword,
	molang.KindOperator:   semtok.TokOperator,
	molang.KindParen:      semtok.TokEnumMember,
	molang.KindComma:      semtok.TokOperator,
	molang.KindWhitespace: semtok.TokComment,
	molang.KindUnknown:    semtok.TokComment,
}

var tokenType = map[semtok.Type]bool{
	semtok.TokNumber:     true,
	semtok.TokString:     true,
	semtok.TokMacro:      true,
	semtok.TokMethod:     true,
	semtok.TokType:       true,
	semtok.TokKeyword:    true,
	semtok.TokOperator:   true,
	semtok.TokEnumMember: true,
	semtok.TokComment:    false,
}

var tokenModifier = map[semtok.Modifier]bool{}

func (m *MolangHandler) GetSemanticTokens(document *textdocument.TextDocument) *protocol.SemanticTokens {
	tokens := []semtok.Token{}

	jsonc.Visit(document.GetText(), &jsonc.Visitor{
		OnLiteralValue: func(value any, offset, length, startLine, startCharacter uint32, pathSupplier func() jsonc.Path) {
			text, ok := value.(string)
			if !ok || text == "" || text[0] == '@' || text[0] == '/' {
				return
			}
			path := pathSupplier()
			if !slices.ContainsFunc(shared.MolangSemanticLocations, func(jsonPath shared.JsonPath) bool { return path.Matches(jsonPath.Path) }) {
				return
			}
			parser, err := molang.NewParser(text)
			if err != nil {
				return
			}
			for _, token := range parser.Tokens {
				tokenType, ok := tokenMap[token.Kind]
				if !ok {
					continue
				}
				tokens = append(tokens,
					semtok.Token{
						Type:  tokenType,
						Line:  startLine,
						Start: startCharacter + token.Offset + 1,
						Len:   token.Length,
					})
			}
		},
	}, nil)

	return &protocol.SemanticTokens{
		Data: semtok.Encode(tokens, tokenType, tokenModifier),
	}
}

var Molang MolangHandler
