package handlers

import (
	"log"
	"regexp"
	"slices"
	"strings"

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

	if index == 0 {
		return nil
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

func (m *MolangHandler) GetHover(document *textdocument.TextDocument, offset uint32, location *jsonc.Location) *protocol.Hover {
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
	if index < 0 {
		return nil
	}
	tPrefix := parser.Tokens[index-1]
	tMethod := parser.Tokens[index]
	if tPrefix.Kind != molang.PREFIX || tMethod.Kind != molang.METHOD {
		return nil
	}
	method, found := sliceutil.Find(molang.GetMethodList(tPrefix.Value), func(method molang.Method) bool {
		return method.Name == tMethod.Value[1:]
	})
	if !found {
		return nil
	}
	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind: protocol.Markdown,
			Value: "```rockide-molang\n" +
				tPrefix.Value + tMethod.Value + string(method.Signature) +
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
			return protocol.ParameterInformation{Label: param.ToString()}
		}),
		ActiveParameter: uint32(activeParam),
	}
	return &protocol.SignatureHelp{
		Signatures: []protocol.SignatureInformation{signature},
	}
}

type tokenPattern struct {
	kind    semtok.Type
	pattern *regexp.Regexp
}

var tokenPatterns = []tokenPattern{
	{semtok.TokNumber, regexp.MustCompile(`^[0-9]+(\.[0-9]+)?f?`)},
	{semtok.TokString, regexp.MustCompile(`^'[^']*'`)},
	{semtok.TokMacro, regexp.MustCompile(`^this`)},
	{semtok.TokMethod, regexp.MustCompile(`^\.([a-zA-Z_][a-zA-Z0-9_.]*)?`)},
	{semtok.TokType, regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*`)},
	{semtok.TokKeyword, regexp.MustCompile(`^(break|continue|for_each|loop|return)`)},
	{semtok.TokOperator, regexp.MustCompile(`^[+\-*/%><=!&|;:?,]+`)},
	{semtok.TokEnumMember, regexp.MustCompile(`[\(\)\{\}\[\]]`)},
	{semtok.TokComment, regexp.MustCompile(`^\s+`)},
	{semtok.TokComment, regexp.MustCompile(`^.`)},
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
			current := text
			for len(current) > 0 {
				for _, tp := range tokenPatterns {
					match := tp.pattern.FindString(current)
					if match != "" {
						length := uint32(len(match))
						tokens = append(tokens,
							semtok.Token{
								Line:  startLine,
								Start: startCharacter + 1,
								Len:   length,
								Type:  tp.kind,
							})
						current = current[length:]
						startCharacter += length
						break
					}
				}
			}
		},
	}, nil)

	return &protocol.SemanticTokens{
		Data: semtok.Encode(tokens, tokenType, tokenModifier),
	}
}

var Molang MolangHandler
