package shared

import (
	"fmt"

	"github.com/ink0rr/rockide/internal/jsonc"
)

type JsonPath struct {
	IsKey bool
	Path  jsonc.Path
}

func JsonKey(path string) JsonPath {
	return JsonPath{IsKey: true, Path: jsonc.NewPath(path)}
}

func JsonValue(path string) JsonPath {
	return JsonPath{IsKey: false, Path: jsonc.NewPath(path)}
}

func (j *JsonPath) GetNodes(root *jsonc.Node) []*jsonc.Node {
	res := []*jsonc.Node{}
	if root == nil {
		return res
	}

	var visitNodes func(node *jsonc.Node, keys jsonc.Path)
	visitNodes = func(node *jsonc.Node, keys jsonc.Path) {
		currentKey, remainingKeys := keys[0], keys[1:]

		if len(remainingKeys) == 0 {
			switch currentKey {
			case "**":
				panic(fmt.Sprintf("invalid path: %s", j.Path))
			case "*":
				for _, child := range node.Children {
					if target := j.getTarget(child); target != nil {
						res = append(res, target)
					}
				}
			default:
				if nextNode := jsonc.FindNodeAtLocation(node, jsonc.Path{currentKey}); nextNode != nil {
					if target := j.getTarget(nextNode); target != nil {
						res = append(res, target)
					}
				}
			}
			return
		}

		switch currentKey {
		case "**":
			for _, child := range node.Children {
				if child.Type == jsonc.NodeTypeProperty && child.Children[0].Value == remainingKeys[0] {
					visitNodes(node, remainingKeys)
				} else {
					visitNodes(child, keys)
				}
			}
		case "*":
			for _, child := range node.Children {
				if child.Type == jsonc.NodeTypeProperty && len(child.Children) == 2 {
					visitNodes(child.Children[1], remainingKeys)
				} else {
					visitNodes(child, remainingKeys)
				}
			}
		default:
			if nextNode := jsonc.FindNodeAtLocation(node, jsonc.Path{currentKey}); nextNode != nil {
				visitNodes(nextNode, remainingKeys)
			}
		}
	}

	visitNodes(root, j.Path)

	return res
}

func (j *JsonPath) getTarget(node *jsonc.Node) *jsonc.Node {
	index := 1
	if j.IsKey {
		index = 0
	}

	if node.Type == jsonc.NodeTypeProperty && len(node.Children) > index {
		child := node.Children[index]
		if child.Type == jsonc.NodeTypeString {
			return child
		}
	} else if !j.IsKey && node.Type == jsonc.NodeTypeString {
		return node
	}
	return nil
}
