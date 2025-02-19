package jsonc_test

import (
	"testing"

	"github.com/ink0rr/rockide/internal/jsonc"
)

func assertPath(t *testing.T, path jsonc.Path, pattern jsonc.Path, shouldMatch bool) {
	if shouldMatch != path.Matches(pattern) {
		t.Errorf("Failed to match path. Path: %s, Pattern: %s, Should Match: %v", path, pattern, shouldMatch)
	}
}

func TestPathMatches(t *testing.T) {
	assertPath(t, jsonc.NewPath("a/b/c"), jsonc.NewPath("a/b/c"), true)
	assertPath(t, jsonc.NewPath("a/b/c"), jsonc.NewPath("a/b/d"), false)
	assertPath(t, jsonc.NewPath("a/b/c"), jsonc.NewPath("a/*/c"), true)
	assertPath(t, jsonc.NewPath("a/d/b/c"), jsonc.NewPath("a/**/c"), true)
	assertPath(t, jsonc.NewPath("a/b/c"), jsonc.NewPath("a/**"), true)
	assertPath(t, jsonc.NewPath("a/b/c"), jsonc.NewPath("b/**"), false)
	assertPath(t, jsonc.NewPath("a/b/c"), jsonc.NewPath("**/c"), true)
	assertPath(t, jsonc.NewPath("a/b/c"), jsonc.NewPath("**/d"), false)
}
