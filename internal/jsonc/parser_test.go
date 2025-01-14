package jsonc_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ink0rr/rockide/internal/jsonc"
)

func stringify(value any) string {
	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func stringifyIndent(value any) string {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func checkParent(t *testing.T, node *jsonc.Node) {
	if node == nil {
		return
	}
	if node.Children != nil {
		for _, child := range node.Children {
			if node != child.Parent {
				t.Error("Mismatched parent")
			}
			child.Parent = nil
			checkParent(t, child)
		}
	}
}

func assertTree(t *testing.T, input string, expected *jsonc.Node, expectedErrors []jsonc.ParseError) {
	actual, errors := jsonc.ParseTree(input, nil)
	if !reflect.DeepEqual(errors, expectedErrors) {
		t.Errorf("Expected: %s, Actual: %s", stringifyIndent(expectedErrors), stringifyIndent(errors))
	}
	checkParent(t, actual)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %s, Actual: %s", stringifyIndent(expected), stringifyIndent(actual))
	}
}

func TestTreeLiterals(t *testing.T) {
	assertTree(t, "true", &jsonc.Node{Type: jsonc.NodeTypeBoolean, Offset: 0, Length: 4, Value: true}, nil)
	assertTree(t, "false", &jsonc.Node{Type: jsonc.NodeTypeBoolean, Offset: 0, Length: 5, Value: false}, nil)
	assertTree(t, "null", &jsonc.Node{Type: jsonc.NodeTypeNull, Offset: 0, Length: 4, Value: nil}, nil)
	assertTree(t, "23", &jsonc.Node{Type: jsonc.NodeTypeNumber, Offset: 0, Length: 2, Value: float64(23)}, nil)
	assertTree(t, "-1.93e-19", &jsonc.Node{Type: jsonc.NodeTypeNumber, Offset: 0, Length: 9, Value: float64(-1.93e-19)}, nil)
	assertTree(t, `"hello"`, &jsonc.Node{Type: jsonc.NodeTypeString, Offset: 0, Length: 7, Value: "hello"}, nil)
}

func TestTreeArrays(t *testing.T) {
	assertTree(t, "[]", &jsonc.Node{Type: jsonc.NodeTypeArray, Offset: 0, Length: 2}, nil)
	assertTree(t, "[ 1 ]", &jsonc.Node{
		Type:     jsonc.NodeTypeArray,
		Offset:   0,
		Length:   5,
		Children: []*jsonc.Node{{Type: jsonc.NodeTypeNumber, Offset: 2, Length: 1, Value: float64(1)}},
	}, nil)
	assertTree(t, `[ 1,"x"]`, &jsonc.Node{
		Type:   jsonc.NodeTypeArray,
		Offset: 0,
		Length: 8,
		Children: []*jsonc.Node{
			{Type: jsonc.NodeTypeNumber, Offset: 2, Length: 1, Value: float64(1)},
			{Type: jsonc.NodeTypeString, Offset: 4, Length: 3, Value: "x"},
		},
	}, nil)
}

func TestTreeObjects(t *testing.T) {
	assertTree(t, "{ }", &jsonc.Node{Type: jsonc.NodeTypeObject, Offset: 0, Length: 3}, nil)
	assertTree(t, `{ "val": 1 }`, &jsonc.Node{
		Type:   jsonc.NodeTypeObject,
		Offset: 0,
		Length: 12,
		Children: []*jsonc.Node{
			{
				Type:        jsonc.NodeTypeProperty,
				Offset:      2,
				Length:      8,
				ColonOffset: 7,
				Children: []*jsonc.Node{
					{Type: jsonc.NodeTypeString, Offset: 2, Length: 5, Value: "val"},
					{Type: jsonc.NodeTypeNumber, Offset: 9, Length: 1, Value: float64(1)},
				},
			},
		},
	}, nil)
	assertTree(t, `{"id": "$", "v": [ null, null] }`, &jsonc.Node{
		Type:   jsonc.NodeTypeObject,
		Offset: 0,
		Length: 32,
		Children: []*jsonc.Node{
			{
				Type:        jsonc.NodeTypeProperty,
				Offset:      1,
				Length:      9,
				ColonOffset: 5,
				Children: []*jsonc.Node{
					{Type: jsonc.NodeTypeString, Offset: 1, Length: 4, Value: "id"},
					{Type: jsonc.NodeTypeString, Offset: 7, Length: 3, Value: "$"},
				},
			},
			{
				Type:        jsonc.NodeTypeProperty,
				Offset:      12,
				Length:      18,
				ColonOffset: 15,
				Children: []*jsonc.Node{
					{Type: jsonc.NodeTypeString, Offset: 12, Length: 3, Value: "v"},
					{
						Type:   jsonc.NodeTypeArray,
						Offset: 17,
						Length: 13,
						Children: []*jsonc.Node{
							{Type: jsonc.NodeTypeNull, Offset: 19, Length: 4, Value: nil},
							{Type: jsonc.NodeTypeNull, Offset: 25, Length: 4, Value: nil},
						},
					},
				},
			},
		},
	}, nil)
	assertTree(t, `{  "id": { "foo": { } } , }`,
		&jsonc.Node{
			Type:   jsonc.NodeTypeObject,
			Offset: 0,
			Length: 27,
			Children: []*jsonc.Node{
				{
					Type:        jsonc.NodeTypeProperty,
					Offset:      3,
					Length:      20,
					ColonOffset: 7,
					Children: []*jsonc.Node{
						{Type: jsonc.NodeTypeString, Offset: 3, Length: 4, Value: "id"},
						{
							Type:   jsonc.NodeTypeObject,
							Offset: 9,
							Length: 14,
							Children: []*jsonc.Node{
								{
									Type:        jsonc.NodeTypeProperty,
									Offset:      11,
									Length:      10,
									ColonOffset: 16,
									Children: []*jsonc.Node{
										{Type: jsonc.NodeTypeString, Offset: 11, Length: 5, Value: "foo"},
										{Type: jsonc.NodeTypeObject, Offset: 18, Length: 3},
									},
								},
							},
						},
					},
				},
			},
		},
		[]jsonc.ParseError{{jsonc.ParseErrorCodePropertyNameExpected, 26, 1}, {jsonc.ParseErrorCodeValueExpected, 26, 1}},
	)
}
