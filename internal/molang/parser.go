package molang

import (
	"fmt"
	"slices"
	"strconv"
)

type Parser struct {
	Source string
	Tokens []Token
}

func NewParser(source string) (*Parser, error) {
	escaped := strconv.QuoteToASCII(source)
	source = escaped[1 : len(escaped)-1]
	parser := &Parser{Source: source}

	current := source
	offset := uint32(0)

	for len(current) > 0 {
		matched := false
		for _, tp := range tokenPatterns {
			match := tp.pattern.FindString(current)
			if match != "" {
				length := uint32(len(match))
				if tp.kind != KindUnknown {
					parser.Tokens = append(parser.Tokens, Token{
						Kind:   tp.kind,
						Value:  match,
						Offset: offset,
						Length: length,
					})
				}
				current = current[length:]
				offset += length
				matched = true
				break
			}
		}
		if !matched {
			return nil, fmt.Errorf("unexpected token at offset %d", offset)
		}
	}

	return parser, nil
}

func (mp *Parser) FindIndex(offset uint32) int {
	return slices.IndexFunc(mp.Tokens, func(token Token) bool {
		return offset >= token.Offset && offset < token.Offset+token.Length
	})
}

type MethodCall struct {
	Prefix     string
	Name       string
	ParamIndex int
}

func (mp *Parser) GetMethodCall(offset uint32) *MethodCall {
	index := mp.FindIndex(offset)
	if index == -1 {
		return nil
	}

	paramIndex := 0
	depth := 0
	for i := index; i >= 0; i-- {
		token := mp.Tokens[i]
		if token.Kind == KindComma && depth == 0 {
			paramIndex++
			continue
		}
		if token.Kind == KindParen && token.Value == ")" {
			depth++
			continue
		}
		if token.Kind == KindParen && token.Value == "(" {
			if depth == 0 {
				if i >= 2 {
					prefix := mp.Tokens[i-2]
					method := mp.Tokens[i-1]
					if prefix.Kind == KindPrefix && method.Kind == KindMethod {
						return &MethodCall{
							Prefix:     prefix.Value,
							Name:       method.Value,
							ParamIndex: paramIndex,
						}
					}
				}
				break
			}
			depth--
		}
	}

	return nil
}
