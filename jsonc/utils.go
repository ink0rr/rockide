package jsonc

import (
	"fmt"
	"reflect"
)

func charAt(s string, pos int) byte {
	length := len(s)
	if pos >= length {
		return 0
	}
	return s[pos]
}

func substring(s string, start, end int) string {
	length := len(s)
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	return s[start:end]
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t'
}

func isLineBreak(ch byte) bool {
	return ch == '\n' || ch == '\r'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isUnknownContentCharacter(code byte) bool {
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
