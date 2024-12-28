package molang_test

import (
	"testing"

	"github.com/ink0rr/rockide/rockide/molang"
)

func TestParser(t *testing.T) {
	parser, err := molang.NewParser("q.life_time && q.item_any_tags('asd', 'bca', 'qwe')")
	if err != nil {
		t.Fatal(err)
	}
	index := parser.FindIndex(13)
	token := parser.Tokens[index]
	if token.Value != "&&" {
		t.Fatal("Not equals!")
	}
}
