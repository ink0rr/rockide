package jsonc

type Path = []any // string | int

type Visitor struct {
	OnObjectBegin    func(offset int, length int, startLine int, startCharacter int, pathSupplier func() Path) bool
	OnObjectProperty func(name string, offset int, length int, startLine int, startCharacter int, pathSupplier func() Path)
	OnObjectEnd      func(offset int, length int, startLine int, startCharacter int)
	OnArrayBegin     func(offset int, length int, startLine int, startCharacter int, pathSupplier func() Path) bool
	OnArrayEnd       func(offset int, length int, startLine int, startCharacter int)
	OnLiteralValue   func(value any, offset int, length int, startLine int, startCharacter int, pathSupplier func() Path)
	OnSeparator      func(sep string, offset int, length int, startLine int, startCharacter int)
	OnComment        func(offset int, length int, startLine int, startCharacter int)
	OnError          func(code ParseErrorCode, offset int, length int, startLine int, startCharacter int)
}
