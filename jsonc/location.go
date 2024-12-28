package jsonc

type Location struct {
	Path            Path
	PreviousNode    *Node
	IsAtPropertyKey bool
}

func (l *Location) matches(pattern Path) bool {
	k := 0
	for i := 0; k < len(pattern) && i < len(l.Path); i++ {
		if pattern[k] == l.Path[i] || pattern[k] == "*" {
			k++
		} else if pattern[k] != "**" {
			return false
		}
	}
	return k == len(pattern)
}

// For a given offset, evaluate the location in the JSON document. Each segment in the location path is either a property name or an array index.
func GetLocation(text string, position int) *Location {
	isError := false
	segments := Path{}
	previousNode := &Node{Type: NodeTypeNull}
	previousNodeInst := &Node{Type: NodeTypeNull}
	isAtPropertyKey := false

	setPreviousNode := func(value any, offset int, length int, nodeType NodeType) {
		previousNodeInst.Value = value
		previousNodeInst.Offset = offset
		previousNodeInst.Length = length
		previousNodeInst.Type = nodeType
		previousNodeInst.ColonOffset = 0
		previousNode = previousNodeInst
	}

	Visit(text, &Visitor{
		OnObjectBegin: func(offset, length, startLine, startCharacter int, pathSupplier func() Path) bool {
			if isError {
				return false
			}
			if position <= offset {
				isError = true
				return false
			}
			previousNode = nil
			isAtPropertyKey = position > offset
			segments = append(segments, "")
			return true
		},
		OnObjectProperty: func(name string, offset, length, startLine, startCharacter int, pathSupplier func() Path) {
			if isError {
				return
			}
			if position < offset {
				isError = true
				return
			}
			setPreviousNode(name, offset, length, NodeTypeProperty)
			segments[len(segments)-1] = name
			if position <= offset+length {
				isError = true
				return
			}
		},
		OnObjectEnd: func(offset, length, startLine, startCharacter int) {
			if isError {
				return
			}
			if position <= offset {
				isError = true
				return
			}
			previousNode = nil
			if len(segments) > 0 {
				segments = segments[:len(segments)-1]
			}
		},
		OnArrayBegin: func(offset, length, startLine, startCharacter int, pathSupplier func() Path) bool {
			if isError {
				return false
			}
			if position <= offset {
				isError = true
				return false
			}
			previousNode = nil
			segments = append(segments, 0)
			return true
		},
		OnArrayEnd: func(offset, length, startLine, startCharacter int) {
			if isError {
				return
			}
			if position <= offset {
				isError = true
				return
			}
			previousNode = nil
			segments = append(segments, 0)
		},
		OnLiteralValue: func(value any, offset, length, startLine, startCharacter int, pathSupplier func() Path) {
			if isError {
				return
			}
			if position < offset {
				isError = true
				return
			}
			setPreviousNode(value, offset, length, getNodeType(value))
			if position <= offset+length {
				isError = true
				return
			}
		},
		OnSeparator: func(sep string, offset, length, startLine, startCharacter int) {
			if isError {
				return
			}
			if position <= offset {
				isError = true
				return
			}
			if sep == ":" && previousNode != nil && previousNode.Type == NodeTypeProperty {
				previousNode.ColonOffset = offset
				isAtPropertyKey = false
				previousNode = nil
			} else if sep == "," {
				last := segments[len(segments)-1]
				switch last := last.(type) {
				case int:
					segments[len(segments)-1] = last + 1
				default:
					isAtPropertyKey = true
					segments[len(segments)-1] = ""
				}
				previousNode = nil
			}
		},
	}, nil)
	return &Location{
		Path:            segments,
		PreviousNode:    previousNode,
		IsAtPropertyKey: isAtPropertyKey,
	}
}
