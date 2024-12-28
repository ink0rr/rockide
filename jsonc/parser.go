package jsonc

import (
	"slices"
	"strconv"
)

// Parses the given text and returns a tree representation the JSON content. On invalid input, the parser tries to be as fault tolerant as possible, but still return a result.
func ParseTree(text string, options *ParseOptions) (root *Node, errors []ParseError) {
	currentParent := &Node{
		Type: NodeTypeArray, Offset: -1, Length: -1,
	}

	ensurePropertyComplete := func(endOffset int) {
		if currentParent.Type == NodeTypeProperty {
			currentParent.Length = endOffset - currentParent.Offset
			currentParent = currentParent.Parent
		}
	}

	onValue := func(valueNode *Node) *Node {
		currentParent.Children = append(currentParent.Children, valueNode)
		return valueNode
	}

	visitor := Visitor{
		OnObjectBegin: func(offset, length, startLine, startCharacter int, pathSupplier func() Path) bool {
			currentParent = onValue(&Node{Type: NodeTypeObject, Offset: offset, Length: -1, Parent: currentParent})
			return true
		},
		OnObjectProperty: func(name string, offset, length, startLine, startCharacter int, pathSupplier func() Path) {
			currentParent = onValue(&Node{Type: NodeTypeProperty, Offset: offset, Length: -1, Parent: currentParent})
			currentParent.Children = append(currentParent.Children, &Node{Type: NodeTypeString, Value: name, Offset: offset, Length: length, Parent: currentParent})
		},
		OnObjectEnd: func(offset, length, startLine, startCharacter int) {
			ensurePropertyComplete(offset + length)

			currentParent.Length = offset + length - currentParent.Offset
			currentParent = currentParent.Parent
			ensurePropertyComplete(offset + length)
		},
		OnArrayBegin: func(offset, length, startLine, startCharacter int, pathSupplier func() Path) bool {
			currentParent = onValue(&Node{Type: NodeTypeArray, Offset: offset, Length: -1, Parent: currentParent})
			return true
		},
		OnArrayEnd: func(offset, length, startLine, startCharacter int) {
			currentParent.Length = offset + length - currentParent.Offset
			currentParent = currentParent.Parent
			ensurePropertyComplete(offset + length)
		},
		OnLiteralValue: func(value any, offset, length, startLine, startCharacter int, pathSupplier func() Path) {
			onValue(&Node{Type: getNodeType(value), Offset: offset, Length: length, Parent: currentParent, Value: value})
			ensurePropertyComplete(offset + length)
		},
		OnSeparator: func(sep string, offset, length, startLine, startCharacter int) {
			if currentParent.Type == NodeTypeProperty {
				if sep == ":" {
					currentParent.ColonOffset = offset
				} else if sep == "," {
					ensurePropertyComplete(offset)
				}
			}
		},
		OnError: func(code ParseErrorCode, offset, length, startLine, startCharacter int) {
			errors = append(errors, ParseError{Error: code, Offset: offset, Length: length})
		},
	}
	Visit(text, &visitor, options)
	if len(currentParent.Children) > 0 {
		root = currentParent.Children[0]
		root.Parent = nil
	}
	return root, errors
}

// Finds the node at the given path in a JSON DOM.
func FindNodeAtLocation(root *Node, path Path) *Node {
	if root == nil {
		return nil
	}
	node := root
	for _, segment := range path {
		switch segment := segment.(type) {
		case string:
			if node.Type != NodeTypeObject || node.Children == nil {
				return nil
			}
			found := false
			for _, propertyNode := range node.Children {
				if propertyNode.Children != nil && len(propertyNode.Children) == 2 && propertyNode.Children[0].Value == segment {
					node = propertyNode.Children[1]
					found = true
					break
				}
			}
			if !found {
				return nil
			}
		case int:
			index := segment
			if node.Type != NodeTypeArray || index < 0 || node.Children == nil || index >= len(node.Children) {
				return nil
			}
			node = node.Children[index]
		}
	}
	return node
}

// Gets the JSON path of the given JSON DOM node
func GetNodePath(node *Node) Path {
	if node.Parent == nil || node.Parent.Children == nil {
		return Path{}
	}
	path := GetNodePath(node.Parent)
	if node.Parent.Type == NodeTypeProperty {
		key := node.Parent.Children[0].Value
		path = append(path, key)
	} else if node.Parent.Type == NodeTypeArray {
		index := slices.Index(node.Parent.Children, node)
		if index != -1 {
			path = append(path, index)
		}
	}
	return path
}

func contains(node *Node, offset int, includeRightBound bool) bool {
	return (offset >= node.Offset && offset < (node.Offset+node.Length)) || includeRightBound && (offset == (node.Offset+node.Length))
}

// Finds the most inner node at the given offset. If includeRightBound is set, also finds nodes that end at the given offset.
func FindNodeAtOffset(node *Node, offset int, includeRightBound bool) *Node {
	if contains(node, offset, includeRightBound) {
		children := node.Children
		if children != nil {
			for i := 0; i < len(children) && children[i].Offset <= offset; i++ {
				item := FindNodeAtOffset(children[i], offset, includeRightBound)
				if item != nil {
					return item
				}
			}

		}
		return node
	}
	return nil
}

// Parses the given text and invokes the visitor functions for each object, array and literal reached.
func Visit(text string, visitor *Visitor, options *ParseOptions) any {
	if options == nil {
		options = &ParseOptions{AllowTrailingComma: false}
	}
	scanner := CreateScanner(text, false)
	jsonPath := Path{}
	suppressedCallbacks := 0

	toBeginVisit := func(visitFunction func(offset, length, startLine, startCharacter int, pathSupplier func() Path) bool) func() {
		if visitFunction == nil {
			return func() {}
		}
		return func() {
			if suppressedCallbacks > 0 {
				suppressedCallbacks++
			} else {
				cbReturn := visitFunction(scanner.GetTokenOffset(), scanner.GetTokenLength(), scanner.GetTokenStartLine(), scanner.GetTokenStartCharacter(), func() Path {
					return append(Path{}, jsonPath...)
				})
				if !cbReturn {
					suppressedCallbacks = 1
				}
			}
		}
	}

	toEndVisit := func(visitFunction func(offset, length, startLine, startCharacter int)) func() {
		if visitFunction == nil {
			return func() {}
		}
		return func() {
			if suppressedCallbacks > 0 {
				suppressedCallbacks--
			}
			if suppressedCallbacks == 0 {
				visitFunction(scanner.GetTokenOffset(), scanner.GetTokenLength(), scanner.GetTokenStartLine(), scanner.GetTokenStartCharacter())
			}
		}
	}

	onObjectBegin := toBeginVisit(visitor.OnObjectBegin)
	onObjectProperty := func(arg string) {
		if visitor.OnObjectProperty != nil && suppressedCallbacks == 0 {
			visitor.OnObjectProperty(arg, scanner.GetTokenOffset(), scanner.GetTokenLength(), scanner.GetTokenStartLine(), scanner.GetTokenStartCharacter(), func() Path {
				return append(Path{}, jsonPath...)
			})
		}
	}
	onObjectEnd := toEndVisit(visitor.OnObjectEnd)
	onArrayBegin := toBeginVisit(visitor.OnArrayBegin)
	onArrayEnd := toEndVisit(visitor.OnArrayEnd)
	onLiteralValue := func(arg any) {
		if visitor.OnLiteralValue != nil && suppressedCallbacks == 0 {
			visitor.OnLiteralValue(arg, scanner.GetTokenOffset(), scanner.GetTokenLength(), scanner.GetTokenStartLine(), scanner.GetTokenStartCharacter(), func() Path {
				return append(Path{}, jsonPath...)
			})
		}
	}
	onSeparator := func(arg string) {
		if visitor.OnSeparator != nil && suppressedCallbacks == 0 {
			visitor.OnSeparator(arg, scanner.GetTokenOffset(), scanner.GetTokenLength(), scanner.GetTokenStartLine(), scanner.GetTokenStartCharacter())
		}
	}
	onComment := func() {
		if visitor.OnComment != nil && suppressedCallbacks == 0 {
			visitor.OnComment(scanner.GetTokenOffset(), scanner.GetTokenLength(), scanner.GetTokenStartLine(), scanner.GetTokenStartCharacter())
		}
	}
	onError := func(arg ParseErrorCode) {
		if visitor.OnError != nil && suppressedCallbacks == 0 {
			visitor.OnError(arg, scanner.GetTokenOffset(), scanner.GetTokenLength(), scanner.GetTokenStartLine(), scanner.GetTokenStartCharacter())
		}
	}

	disallowComments := options.DisallowComments
	allowTrailingComma := options.AllowTrailingComma

	var parseValue func() bool

	handleError := func(errorCode ParseErrorCode, skipUntilAfter, skipUntil []SyntaxKind) {
		onError(errorCode)
		if len(skipUntilAfter) > 0 || len(skipUntil) > 0 {
			token := scanner.GetToken()
			for token != SyntaxKindEOF {
				if slices.Contains(skipUntilAfter, token) {
					scanner.Scan()
					break
				} else if slices.Contains(skipUntil, token) {
					break
				}
				token = scanner.Scan()
			}
		}
	}

	scanNext := func() SyntaxKind {
		for {
			token := scanner.Scan()
			switch scanner.GetTokenError() {
			case ScanErrorInvalidUnicode:
				handleError(ParseErrorCodeInvalidUnicode, nil, nil)
			case ScanErrorInvalidEscapeCharacter:
				handleError(ParseErrorCodeInvalidEscapeCharacter, nil, nil)
			case ScanErrorUnexpectedEndOfNumber:
				handleError(ParseErrorCodeUnexpectedEndOfNumber, nil, nil)
			case ScanErrorUnexpectedEndOfComment:
				if !disallowComments {
					handleError(ParseErrorCodeUnexpectedEndOfComment, nil, nil)
				}
			case ScanErrorUnexpectedEndOfString:
				handleError(ParseErrorCodeUnexpectedEndOfString, nil, nil)
			case ScanErrorInvalidCharacter:
				handleError(ParseErrorCodeInvalidCharacter, nil, nil)
			}
			switch token {
			case SyntaxKindLineCommentTrivia, SyntaxKindBlockCommentTrivia:
				if disallowComments {
					handleError(ParseErrorCodeInvalidCommentToken, nil, nil)
				} else {
					onComment()
				}
			case SyntaxKindUnknown:
				handleError(ParseErrorCodeInvalidSymbol, nil, nil)
			case SyntaxKindTrivia, SyntaxKindLineBreakTrivia:
				// Skip trivia tokens
			default:
				return token
			}
		}
	}

	parseString := func(isValue bool) bool {
		value := scanner.GetTokenValue()
		if isValue {
			onLiteralValue(value)
		} else {
			onObjectProperty(value)
			jsonPath = append(jsonPath, value)
		}
		scanNext()
		return true
	}

	parseLiteral := func() bool {
		switch scanner.GetToken() {
		case SyntaxKindNumericLiteral:
			tokenValue := scanner.GetTokenValue()
			value, err := strconv.ParseFloat(tokenValue, 64)
			if err != nil {
				handleError(ParseErrorCodeInvalidNumberFormat, nil, nil)
				value = 0
			}
			onLiteralValue(value)
		case SyntaxKindNullKeyword:
			onLiteralValue(nil)
		case SyntaxKindTrueKeyword:
			onLiteralValue(true)
		case SyntaxKindFalseKeyword:
			onLiteralValue(false)
		default:
			return false
		}
		scanNext()
		return true
	}

	parseProperty := func() bool {
		if scanner.GetToken() != SyntaxKindStringLiteral {
			handleError(ParseErrorCodePropertyNameExpected, nil, []SyntaxKind{SyntaxKindCloseBraceToken, SyntaxKindCommaToken})
			return false
		}
		parseString(false)
		if scanner.GetToken() == SyntaxKindColonToken {
			onSeparator(":")
			scanNext()
			if !parseValue() {
				handleError(ParseErrorCodeValueExpected, nil, []SyntaxKind{SyntaxKindCloseBraceToken, SyntaxKindCommaToken})
			}
		} else {
			handleError(ParseErrorCodeColonExpected, nil, []SyntaxKind{SyntaxKindCloseBraceToken, SyntaxKindCommaToken})
		}
		jsonPath = jsonPath[:len(jsonPath)-1]
		return true
	}

	parseObject := func() bool {
		onObjectBegin()
		scanNext()
		needsComma := false
		for scanner.GetToken() != SyntaxKindCloseBraceToken && scanner.GetToken() != SyntaxKindEOF {
			if scanner.GetToken() == SyntaxKindCommaToken {
				if !needsComma {
					handleError(ParseErrorCodeValueExpected, nil, nil)
				}
				onSeparator(",")
				scanNext()
				if scanner.GetToken() == SyntaxKindCloseBraceToken && allowTrailingComma {
					break
				}
			} else if needsComma {
				handleError(ParseErrorCodeCommaExpected, nil, nil)
			}
			if !parseProperty() {
				handleError(ParseErrorCodeValueExpected, nil, []SyntaxKind{SyntaxKindCloseBraceToken, SyntaxKindCommaToken})
			}
			needsComma = true
		}
		onObjectEnd()
		if scanner.GetToken() != SyntaxKindCloseBraceToken {
			handleError(ParseErrorCodeCloseBraceExpected, []SyntaxKind{SyntaxKindCloseBraceToken}, nil)
		} else {
			scanNext()
		}
		return true
	}

	parseArray := func() bool {
		onArrayBegin()
		scanNext()
		isFirstElement := true
		needsComma := false
		for scanner.GetToken() != SyntaxKindCloseBracketToken && scanner.GetToken() != SyntaxKindEOF {
			if scanner.GetToken() == SyntaxKindCommaToken {
				if !needsComma {
					handleError(ParseErrorCodeValueExpected, nil, nil)
				}
				onSeparator(",")
				scanNext()
				if scanner.GetToken() == SyntaxKindCloseBracketToken && allowTrailingComma {
					break
				}
			} else if needsComma {
				handleError(ParseErrorCodeCommaExpected, nil, nil)
			}
			if isFirstElement {
				jsonPath = append(jsonPath, 0)
				isFirstElement = false
			} else {
				index := len(jsonPath) - 1
				switch value := jsonPath[index].(type) {
				case int:
					jsonPath[index] = value + 1
				}
			}
			if !parseValue() {
				handleError(ParseErrorCodeValueExpected, nil, []SyntaxKind{SyntaxKindCloseBracketToken, SyntaxKindCommaToken})
			}
			needsComma = true
		}
		onArrayEnd()
		if !isFirstElement {
			jsonPath = jsonPath[:len(jsonPath)-1]
		}
		if scanner.GetToken() != SyntaxKindCloseBracketToken {
			handleError(ParseErrorCodeCloseBracketExpected, []SyntaxKind{SyntaxKindCloseBracketToken}, nil)
		} else {
			scanNext()
		}
		return true
	}

	parseValue = func() bool {
		switch scanner.GetToken() {
		case SyntaxKindOpenBracketToken:
			return parseArray()
		case SyntaxKindOpenBraceToken:
			return parseObject()
		case SyntaxKindStringLiteral:
			return parseString(true)
		default:
			return parseLiteral()
		}
	}

	scanNext()
	if scanner.GetToken() == SyntaxKindEOF {
		if options.AllowEmptyContent {
			return true
		}
		handleError(ParseErrorCodeValueExpected, nil, nil)
		return false
	}
	if !parseValue() {
		handleError(ParseErrorCodeValueExpected, nil, nil)
		return false
	}
	if scanner.GetToken() != SyntaxKindEOF {
		handleError(ParseErrorCodeEndOfFileExpected, nil, nil)
	}
	return true
}
