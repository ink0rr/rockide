package core

import "github.com/ink0rr/rockide/jsonc"

func FindNodesAtPath(root *jsonc.Node, jsonPath []string) []*jsonc.Node {
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
				for _, child := range node.Children {
					result = append(result, child)
				}
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

func IsJsonPathMatch(jsonPath jsonc.Path, targetPath []string) bool {
	jsonIndex := 0
	pathIndex := 0

	for jsonIndex < len(jsonPath) && pathIndex < len(targetPath) {
		if targetPath[pathIndex] == "**" {
			// Return early if "**" at the end of the target path
			if pathIndex == len(targetPath)-1 {
				return true
			}

			// Attempt to find a matching segment for the next part of the path after "**"
			for jsonIndex < len(jsonPath) && jsonPath[jsonIndex] != targetPath[pathIndex+1] {
				jsonIndex++
			}
			pathIndex++
			continue
		}
		// Match current segment
		if targetPath[pathIndex] == "*" || jsonPath[jsonIndex] == targetPath[pathIndex] {
			jsonIndex++
			pathIndex++
		} else {
			// Segment does not match
			break
		}
	}

	// Check if all jsonPath segments were matched
	return jsonIndex == len(jsonPath) && pathIndex == len(targetPath)
}
