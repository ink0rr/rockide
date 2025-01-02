package handlers

import "github.com/ink0rr/rockide/jsonc"

func isJsonPathMatch(jsonPath jsonc.Path, targetPath []string) bool {
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
