package jsonc

import (
	"fmt"
	"reflect"
)

func charAt(s []rune, pos uint32) rune {
	length := uint32(len(s))
	if pos >= length {
		return 0
	}
	return s[pos]
}

func substring(s []rune, start, end uint32) []rune {
	length := uint32(len(s))
	if end > length {
		end = length
	}
	return s[start:end]
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func isLineBreak(ch rune) bool {
	return ch == '\n' || ch == '\r'
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isUnknownContentCharacter(code rune) bool {
	if isWhiteSpace(code) || isLineBreak(code) {
		return false
	}
	switch code {
	case '{', '}', '[', ']', '"', ':', ',', '/':
		return false
	}
	return true
}

func getNodeType(value any) NodeType {
	switch value.(type) {
	case bool:
		return NodeTypeBoolean
	case float64:
		return NodeTypeNumber
	case string:
		return NodeTypeString
	case nil:
		return NodeTypeNull
	default:
		panic(fmt.Sprintf("Unhandled type %s", reflect.TypeOf(value).String()))
	}
}
