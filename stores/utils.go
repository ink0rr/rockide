package stores

import (
	"fmt"

	"github.com/ink0rr/rockide/internal/jsonc"
)

func findNodesAtPath(root *jsonc.Node, jsonPath jsonc.Path) []*jsonc.Node {
	result := []*jsonc.Node{}
	if root == nil {
		return result
	}
	var visitNodes func(node *jsonc.Node, keys jsonc.Path)
	visitNodes = func(node *jsonc.Node, keys jsonc.Path) {
		if len(keys) == 0 {
			panic(fmt.Sprintf("unhandled empty keys: %s", jsonPath))
		}
		currentKey, remainingKeys := keys[0], keys[1:]
		if len(remainingKeys) == 0 {
			if currentKey == "**" {
				panic(fmt.Sprintf("invalid path: %s", jsonPath))
			}
			if currentKey == "*" {
				result = append(result, node.Children...)
				return
			}
			nextNode := jsonc.FindNodeAtLocation(node, jsonc.Path{currentKey})
			if nextNode != nil {
				result = append(result, nextNode)
			}
			return
		}

		if currentKey == "*" {
			for _, child := range node.Children {
				if child.Type == jsonc.NodeTypeProperty {
					// Make sure it's a valid property
					if len(child.Children) == 2 {
						visitNodes(child.Children[1], remainingKeys)
					}
				} else {
					visitNodes(child, remainingKeys)
				}
			}
			return
		}
		if currentKey == "**" {
			for _, child := range node.Children {
				if child.Type == "property" && child.Children[0].Value == remainingKeys[0] {
					visitNodes(node, remainingKeys)
				} else {
					visitNodes(child, keys)
				}
			}
			return
		}
		nextNode := jsonc.FindNodeAtLocation(node, jsonc.Path{currentKey})
		if nextNode != nil {
			visitNodes(nextNode, remainingKeys)
		}
	}
	visitNodes(root, jsonPath)
	return result
}

// Skip the keys when an entry might match both keys and values
func skipKey(node *jsonc.Node) *string {
	value, ok := node.Value.(string)
	if !ok || node.Parent != nil && node.Parent.Type == jsonc.NodeTypeProperty && len(node.Parent.Children) > 0 && node == node.Parent.Children[0] {
		return nil
	}
	return &value
}
