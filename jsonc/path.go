package jsonc

type Path = []any // string | int

// Matches path against a pattern consisting of strings (for properties) and numbers (for array indices). '*' will match a single segment of any property name or index. '**' will match a sequence of segments of any property name or index, or no segment.
func PathMatches(path Path, pattern []string) bool {
	pathIndex := 0
	patternIndex := 0

	for pathIndex < len(path) && patternIndex < len(pattern) {
		if pattern[patternIndex] == "**" {
			// Return early if "**" at the end of the target path
			if patternIndex == len(pattern)-1 {
				return true
			}

			// Attempt to find a matching segment for the next part of the path after "**"
			for pathIndex < len(path) && path[pathIndex] != pattern[patternIndex+1] {
				pathIndex++
			}
			patternIndex++
			continue
		}
		// Match current segment
		if pattern[patternIndex] == "*" || path[pathIndex] == pattern[patternIndex] {
			pathIndex++
			patternIndex++
		} else {
			// Segment does not match
			break
		}
	}

	// Check if all jsonPath segments were matched
	return pathIndex == len(path) && patternIndex == len(pattern)
}
