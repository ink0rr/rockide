package jsonc_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/ink0rr/rockide/internal/jsonc"
)

func assertLocation(t *testing.T, input string, expectedSegments jsonc.Path, expectedNodeType jsonc.NodeType, expectedCompleteProperty bool) {
	offset := strings.Index(input, "|")
	if offset == -1 {
		panic("Missing offset char in input string")
	}
	text := input[:offset] + input[offset+1:]
	actual := jsonc.GetLocation(text, uint32(offset))
	if !reflect.DeepEqual(actual.Path, expectedSegments) {
		t.Errorf("Failed to match segments. Input: %s, Expected: %s, Actual: %s", input, stringify(expectedSegments), stringify(actual.Path))
	}
	if actual.PreviousNode != nil && actual.PreviousNode.Type != expectedNodeType {
		t.Errorf("Failed to match node type. Input: %s, Expected: %s, Actual: %s", input, expectedNodeType, actual.PreviousNode.Type)
	}
	if actual.IsAtPropertyKey != expectedCompleteProperty {
		t.Errorf("Expected: %v, Actual: %v", expectedCompleteProperty, actual.IsAtPropertyKey)
	}
}

func TestLocationProperties(t *testing.T) {
	assertLocation(t, `|{ "foo": "bar" }`, jsonc.Path{}, jsonc.NodeTypeNull, false)
	assertLocation(t, `{| "foo": "bar" }`, jsonc.Path{""}, jsonc.NodeTypeNull, true)
	assertLocation(t, `{ |"foo": "bar" }`, jsonc.Path{"foo"}, jsonc.NodeTypeProperty, true)
	assertLocation(t, `{ "foo|": "bar" }`, jsonc.Path{"foo"}, jsonc.NodeTypeProperty, true)
	assertLocation(t, `{ "foo"|: "bar" }`, jsonc.Path{"foo"}, jsonc.NodeTypeProperty, true)
	assertLocation(t, `{ "foo": "bar"| }`, jsonc.Path{"foo"}, jsonc.NodeTypeString, false)
	assertLocation(t, `{ "foo":| "bar" }`, jsonc.Path{"foo"}, jsonc.NodeTypeNull, false)
	assertLocation(t, `{ "foo": {"bar|": 1, "car": 2 } }`, jsonc.Path{"foo", "bar"}, jsonc.NodeTypeProperty, true)
	assertLocation(t, `{ "foo": {"bar": 1|, "car": 3 } }`, jsonc.Path{"foo", "bar"}, jsonc.NodeTypeNumber, false)
	assertLocation(t, `{ "foo": {"bar": 1,| "car": 4 } }`, jsonc.Path{"foo", ""}, jsonc.NodeTypeNull, true)
	assertLocation(t, `{ "foo": {"bar": 1, "ca|r": 5 } }`, jsonc.Path{"foo", "car"}, jsonc.NodeTypeProperty, true)
	assertLocation(t, `{ "foo": {"bar": 1, "car": 6| } }`, jsonc.Path{"foo", "car"}, jsonc.NodeTypeNumber, false)
	assertLocation(t, `{ "foo": {"bar": 1, "car": 7 }| }`, jsonc.Path{"foo"}, jsonc.NodeTypeNull, false)
	assertLocation(t, `{ "foo": {"bar": 1, "car": 8 },| "goo": {} }`, jsonc.Path{""}, jsonc.NodeTypeNull, true)
	assertLocation(t, `{ "foo": {"bar": 1, "car": 9 }, "go|o": {} }`, jsonc.Path{"goo"}, jsonc.NodeTypeProperty, true)
	assertLocation(t, `{ "dep": {"bar": 1, "car": |`, jsonc.Path{"dep", "car"}, jsonc.NodeTypeNull, false)
	assertLocation(t, `{ "dep": {"bar": 1,, "car": |`, jsonc.Path{"dep", "car"}, jsonc.NodeTypeNull, false)
	assertLocation(t, `{ "dep": {"bar": "na", "dar": "ma", "car": | } }`, jsonc.Path{"dep", "car"}, jsonc.NodeTypeNull, false)
}
