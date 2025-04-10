package molang

import (
	"fmt"
	"slices"
	"strings"
)

type Parser struct {
	Source string
	Tokens []Token
}

func NewParser(source string) (*Parser, error) {
	parser := &Parser{Source: source}

	current := source
	offset := 0

	for len(current) > 0 {
		matched := false
		for _, tp := range tokenPatterns {
			match := tp.Pattern.FindString(current)
			if match != "" {
				length := len(match)
				parser.Tokens = append(parser.Tokens, Token{
					Kind:   tp.Kind,
					Value:  match,
					Offset: offset,
					Length: length,
				})
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

func (mp *Parser) FindIndex(offset int) int {
	return slices.IndexFunc(mp.Tokens, func(token Token) bool {
		return offset >= token.Offset && offset < token.Offset+token.Length
	})
}

func (mp *Parser) IsMethodCall(offset int) bool {
	index := mp.FindIndex(offset)
	if index == -1 {
		return false
	}

	depth := 0
	for i := index; i >= 0; i-- {
		token := mp.Tokens[i]
		if token.Kind == PAREN && token.Value == ")" {
			depth++
			continue
		}
		if token.Kind == PAREN && token.Value == "(" {
			if depth == 0 {
				break
			}
			depth--
		}
	}

	depth = 0
	for i := index + 1; i < len(mp.Tokens); i++ {
		token := mp.Tokens[i]
		if token.Kind == PAREN && token.Value == "(" {
			depth++
			continue
		}
		if token.Kind == PAREN && token.Value == ")" {
			if depth == 0 {
				return true
			}
			depth--
		}
	}

	return false
}

type MethodCall struct {
	Prefix     string
	Name       string
	ParamIndex int
}

func (mp *Parser) GetMethodCall(offset int) *MethodCall {
	index := mp.FindIndex(offset)
	if index == -1 {
		return nil
	}

	paramIndex := 0
	depth := 0
	for i := index; i >= 0; i-- {
		token := mp.Tokens[i]
		if token.Kind == COMMA && depth == 0 {
			paramIndex++
			continue
		}
		if token.Kind == PAREN && token.Value == ")" {
			depth++
			continue
		}
		if token.Kind == PAREN && token.Value == "(" {
			if depth == 0 {
				if i-2 >= 0 && i-1 >= 0 {
					prefix := mp.Tokens[i-2]
					method := mp.Tokens[i-1]
					if prefix.Kind == PREFIX && method.Kind == METHOD {
						return &MethodCall{
							Prefix:     prefix.Value,
							Name:       strings.TrimPrefix(method.Value, "."),
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
