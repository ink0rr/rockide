package jsonc

import (
	"strconv"
	"strings"
)

type Path []any // string | int

func NewPath(path string) Path {
	res := Path{}
	for _, segment := range strings.Split(path, "/") {
		if val, err := strconv.Atoi(segment); err == nil {
			res = append(res, val)
		} else {
			res = append(res, segment)
		}
	}
	return res
}

// PathMatches checks if path matches the given pattern.
// The pattern may contain wildcards "*" and "**" (match any segment and any number of segments, respectively).
func (path Path) Matches(pattern Path) bool {
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
