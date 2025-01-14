package jsonc_test

import (
	"testing"

	"github.com/ink0rr/rockide/internal/jsonc"
)

func assertPath(t *testing.T, path jsonc.Path, pattern []string, shouldMatch bool) {
	if shouldMatch != path.Matches(pattern) {
		t.Errorf("Failed to match path. Path: %s, Pattern: %s, Should Match: %v", path, pattern, shouldMatch)
	}
}

func TestPathMatches(t *testing.T) {
	assertPath(t, jsonc.Path{"a", "b", "c"}, []string{"a", "b", "c"}, true)
	assertPath(t, jsonc.Path{"a", "b", "c"}, []string{"a", "b", "d"}, false)
	assertPath(t, jsonc.Path{"a", "b", "c"}, []string{"a", "*", "c"}, true)
	assertPath(t, jsonc.Path{"a", "d", "b", "c"}, []string{"a", "**", "c"}, true)
	assertPath(t, jsonc.Path{"a", "b", "c"}, []string{"a", "**"}, true)
	assertPath(t, jsonc.Path{"a", "b", "c"}, []string{"b", "**"}, false)
	assertPath(t, jsonc.Path{"a", "b", "c"}, []string{"**", "c"}, true)
	assertPath(t, jsonc.Path{"a", "b", "c"}, []string{"**", "d"}, false)
}
