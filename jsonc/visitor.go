package jsonc

type Path = []any // string | uint32

type Visitor struct {
	OnObjectBegin    func(offset, length, startLine, startCharacter uint32, pathSupplier func() Path) bool
	OnObjectProperty func(name string, offset, length, startLine, startCharacter uint32, pathSupplier func() Path)
	OnObjectEnd      func(offset, length, startLine, startCharacter uint32)
	OnArrayBegin     func(offset, length, startLine, startCharacter uint32, pathSupplier func() Path) bool
	OnArrayEnd       func(offset, length, startLine, startCharacter uint32)
	OnLiteralValue   func(value any, offset, length, startLine, startCharacter uint32, pathSupplier func() Path)
	OnSeparator      func(sep string, offset, length, startLine, startCharacter uint32)
	OnComment        func(offset, length, startLine, startCharacter uint32)
	OnError          func(code ParseErrorCode, offset, length, startLine, startCharacter uint32)
}
