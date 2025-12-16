package lang

import "github.com/ink0rr/rockide/internal/protocol"

type NodeKind uint8

const (
	NodeFile NodeKind = iota
	NodeEntry
	NodeKey
	NodeAssign
	NodeValue
	NodeText
	NodeFormatSpecifier
	NodeFormatCode
	NodeLineBreak
	NodeComment
	NodeEmoji
)

type Node struct {
	Kind   NodeKind
	Offset uint32
	Start  Position
	End    Position
	Value  string

	children []*Node
	parent   *Node
	index    int
}

func (n *Node) Children() []*Node {
	return n.children
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) NextSibling() *Node {
	if n.parent == nil {
		return nil
	}
	i := n.index + 1
	if i >= len(n.parent.children) {
		return nil
	}
	return n.parent.children[i]
}

func (n *Node) PrevSibling() *Node {
	if n.parent == nil {
		return nil
	}
	i := n.index - 1
	if i < 0 {
		return nil
	}
	return n.parent.children[i]
}

func NodeAt(root *Node, position protocol.Position) *Node {
	if position.Line < root.Start.Line || position.Line > root.End.Line {
		return nil
	}
	if position.Line == root.Start.Line && position.Character < root.Start.Column {
		return nil
	}
	if position.Line == root.End.Line && position.Character > root.End.Column {
		return nil
	}
	for _, child := range root.children {
		if node := NodeAt(child, position); node != nil {
			return node
		}
	}
	return root
}
