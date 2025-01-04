package jsonc_test

import (
	"testing"

	"github.com/ink0rr/rockide/jsonc"
)

func assertPath(t *testing.T, path, pattern jsonc.Path, shouldMatch bool) {
	if shouldMatch != jsonc.PathMatches(path, pattern) {
		t.Errorf("Failed to match path. Path: %s, Pattern: %s, Should Match: %v", path, pattern, shouldMatch)
	}
}

func TestPathMatches(t *testing.T) {
	assertPath(t, jsonc.Path{"a", "b", "c"}, jsonc.Path{"a", "b", "c"}, true)
	assertPath(t, jsonc.Path{"a", "b", "c"}, jsonc.Path{"a", "b", "d"}, false)
	assertPath(t, jsonc.Path{"a", "b", "c"}, jsonc.Path{"a", "*", "c"}, true)
	assertPath(t, jsonc.Path{"a", "d", "b", "c"}, jsonc.Path{"a", "**", "c"}, true)
	assertPath(t, jsonc.Path{"a", "b", "c"}, jsonc.Path{"a", "**"}, true)
	assertPath(t, jsonc.Path{"a", "b", "c"}, jsonc.Path{"b", "**"}, false)
	assertPath(t, jsonc.Path{"a", "b", "c"}, jsonc.Path{"**", "c"}, true)
	assertPath(t, jsonc.Path{"a", "b", "c"}, jsonc.Path{"**", "d"}, false)
}
