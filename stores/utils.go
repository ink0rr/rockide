package stores

import "github.com/ink0rr/rockide/jsonc"

func findNodesAtPath(root *jsonc.Node, jsonPath []string) []*jsonc.Node {
	result := []*jsonc.Node{}
	var visitNodes func(node *jsonc.Node, keys []string)
	visitNodes = func(node *jsonc.Node, keys []string) {
		if len(keys) == 0 {
			panic(`Unhandled empty keys: ${jsonPath}`)
		}
		currentKey, remainingKeys := keys[0], keys[1:]
		if len(remainingKeys) == 0 {
			if currentKey == "**" {
				panic(`Invalid JSON path: ${jsonPath}`)
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
					visitNodes(child.Children[1], remainingKeys)
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

func skipKey(node *jsonc.Node) transformResult {
	value, ok := node.Value.(string)
	if !ok || node.Parent != nil && node.Parent.Type == jsonc.NodeTypeProperty && len(node.Parent.Children) > 0 {
		return transformResult{Skip: true}
	}
	return transformResult{Value: value}
}
