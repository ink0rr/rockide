package lang

type Parser struct {
	lexer *Lexer
}

func NewParser(input []rune) *Parser {
	lexer := NewLexer(input)
	return &Parser{
		lexer: lexer,
	}
}

func (p *Parser) Parse() *Node {
	root := &Node{
		Kind: NodeFile,
	}
	entry := &Node{
		Kind:   NodeEntry,
		parent: root,
	}
	var value *Node
	var lastToken Token
	nextEntry := func() {
		if value != nil {
			last := value.children[len(value.children)-1]
			value.End = last.End
			value = nil
		}
		if entry != nil && len(entry.children) > 0 {
			last := entry.children[len(entry.children)-1]
			entry.End = last.End
			root.children = append(root.children, entry)
			entry = &Node{
				Kind:   NodeEntry,
				parent: root,
			}
		}
	}
	for token := range p.lexer.Next() {
		lastToken = token
		node := &Node{
			Start:  token.Start,
			End:    token.End,
			Value:  token.Value,
			Offset: token.Offset,
		}
		if kind, ok := tokenToNodeKind[token.Kind]; ok {
			node.Kind = kind
		}
		switch token.Kind {
		case TokenNewline:
			if entry != nil {
				nextEntry()
			}
		case TokenComment:
			if value != nil {
				node.parent = value
				node.index = len(value.children)
				value.children = append(value.children, node)
			} else if entry != nil {
				node.parent = entry
				node.index = len(entry.children)
				entry.children = append(entry.children, node)
			}
		case TokenKey, TokenAssign:
			node.parent = entry
			node.index = len(entry.children)
			entry.children = append(entry.children, node)
		case TokenText, TokenFormatSpecifier, TokenFormatCode, TokenLineBreak, TokenEmoji:
			if value == nil {
				value = &Node{
					Kind:   NodeValue,
					Start:  token.Start,
					parent: entry,
				}
				value.index = len(entry.children)
				value.parent = entry
				entry.children = append(entry.children, value)
			}
			node.parent = value
			node.index = len(value.children)
			value.children = append(value.children, node)
		}
	}
	nextEntry()
	root.End = lastToken.End
	return root
}

var tokenToNodeKind = map[TokenKind]NodeKind{
	TokenKey:             NodeKey,
	TokenAssign:          NodeAssign,
	TokenText:            NodeText,
	TokenFormatSpecifier: NodeFormatSpecifier,
	TokenFormatCode:      NodeFormatCode,
	TokenLineBreak:       NodeLineBreak,
	TokenEmoji:           NodeEmoji,
	TokenComment:         NodeComment,
}
