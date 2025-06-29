package molang_test

import (
	"reflect"
	"testing"

	"github.com/ink0rr/rockide/internal/molang"
)

func assertTokens(t *testing.T, input string, expected []molang.Token) {
	parser, err := molang.NewParser(input)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(parser.Tokens, expected) {
		t.Errorf("Expected: %v, Actual: %v", expected, parser.Tokens)
	}
}

func TestParser(t *testing.T) {
	assertTokens(t, "123", []molang.Token{{molang.KindNumber, 0, 3, "123"}})
	assertTokens(t, "&&", []molang.Token{{molang.KindOperator, 0, 2, "&&"}})
	assertTokens(t, "q.is_item_name_any('slot.weapon.mainhand', 0, 'minecraft:iron_sword')", []molang.Token{
		{molang.KindPrefix, 0, 1, "q"},
		{molang.KindMethod, 1, 17, ".is_item_name_any"},
		{molang.KindParen, 18, 1, "("},
		{molang.KindString, 19, 22, "'slot.weapon.mainhand'"},
		{molang.KindComma, 41, 1, ","},
		{molang.KindNumber, 43, 1, "0"},
		{molang.KindComma, 44, 1, ","},
		{molang.KindString, 46, 22, "'minecraft:iron_sword'"},
		{molang.KindParen, 68, 1, ")"},
	})
	assertTokens(t, "q.life_time && q.item_any_tags('asd', 'bca', 'qwe')", []molang.Token{
		{molang.KindPrefix, 0, 1, "q"},
		{molang.KindMethod, 1, 10, ".life_time"},
		{molang.KindOperator, 12, 2, "&&"},
		{molang.KindPrefix, 15, 1, "q"},
		{molang.KindMethod, 16, 14, ".item_any_tags"},
		{molang.KindParen, 30, 1, "("},
		{molang.KindString, 31, 5, "'asd'"},
		{molang.KindComma, 36, 1, ","},
		{molang.KindString, 38, 5, "'bca'"},
		{molang.KindComma, 43, 1, ","},
		{molang.KindString, 45, 5, "'qwe'"},
		{molang.KindParen, 50, 1, ")"},
	})
}
