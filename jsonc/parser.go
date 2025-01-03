package jsonc

import (
	"slices"
)

// Parses the given text and returns a tree representation the JSON content. On invalid input, the parser tries to be as fault tolerant as possible, but still return a result.
func ParseTree(text string, options *ParseOptions) (root *Node, errors []ParseError) {
	currentParent := &Node{Type: NodeTypeArray}

	ensurePropertyComplete := func(endOffset uint32) {
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
		OnObjectBegin: func(offset, length, startLine, startCharacter uint32, pathSupplier func() Path) bool {
			currentParent = onValue(&Node{Type: NodeTypeObject, Offset: offset, Parent: currentParent})
			return true
		},
		OnObjectProperty: func(name string, offset, length, startLine, startCharacter uint32, pathSupplier func() Path) {
			currentParent = onValue(&Node{Type: NodeTypeProperty, Offset: offset, Parent: currentParent})
			currentParent.Children = append(currentParent.Children, &Node{Type: NodeTypeString, Value: name, Offset: offset, Length: length, Parent: currentParent})
		},
		OnObjectEnd: func(offset, length, startLine, startCharacter uint32) {
			ensurePropertyComplete(offset + length)

			currentParent.Length = offset + length - currentParent.Offset
			currentParent = currentParent.Parent
			ensurePropertyComplete(offset + length)
		},
		OnArrayBegin: func(offset, length, startLine, startCharacter uint32, pathSupplier func() Path) bool {
			currentParent = onValue(&Node{Type: NodeTypeArray, Offset: offset, Parent: currentParent})
			return true
		},
		OnArrayEnd: func(offset, length, startLine, startCharacter uint32) {
			currentParent.Length = offset + length - currentParent.Offset
			currentParent = currentParent.Parent
			ensurePropertyComplete(offset + length)
		},
		OnLiteralValue: func(value any, offset, length, startLine, startCharacter uint32, pathSupplier func() Path) {
			onValue(&Node{Type: getNodeType(value), Offset: offset, Length: length, Parent: currentParent, Value: value})
			ensurePropertyComplete(offset + length)
		},
		OnSeparator: func(sep string, offset, length, startLine, startCharacter uint32) {
			if currentParent.Type == NodeTypeProperty {
				if sep == ":" {
					currentParent.ColonOffset = offset
				} else if sep == "," {
					ensurePropertyComplete(offset)
				}
			}
		},
		OnError: func(code ParseErrorCode, offset, length, startLine, startCharacter uint32) {
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
				if len(propertyNode.Children) == 2 && propertyNode.Children[0].Value == segment {
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

// Finds the most inner node at the given offset. If includeRightBound is set, also finds nodes that end at the given offset.
func FindNodeAtOffset(node *Node, offset uint32, includeRightBound bool) *Node {
	if (offset >= node.Offset && offset < (node.Offset+node.Length)) || includeRightBound && (offset == (node.Offset+node.Length)) {
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
