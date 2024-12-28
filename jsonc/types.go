package jsonc

type ScanError int

const (
	ScanErrorNone                   ScanError = 0
	ScanErrorUnexpectedEndOfComment ScanError = 1
	ScanErrorUnexpectedEndOfString  ScanError = 2
	ScanErrorUnexpectedEndOfNumber  ScanError = 3
	ScanErrorInvalidUnicode         ScanError = 4
	ScanErrorInvalidEscapeCharacter ScanError = 5
	ScanErrorInvalidCharacter       ScanError = 6
)

type SyntaxKind int

const (
	SyntaxKindOpenBraceToken     SyntaxKind = 1
	SyntaxKindCloseBraceToken    SyntaxKind = 2
	SyntaxKindOpenBracketToken   SyntaxKind = 3
	SyntaxKindCloseBracketToken  SyntaxKind = 4
	SyntaxKindCommaToken         SyntaxKind = 5
	SyntaxKindColonToken         SyntaxKind = 6
	SyntaxKindNullKeyword        SyntaxKind = 7
	SyntaxKindTrueKeyword        SyntaxKind = 8
	SyntaxKindFalseKeyword       SyntaxKind = 9
	SyntaxKindStringLiteral      SyntaxKind = 10
	SyntaxKindNumericLiteral     SyntaxKind = 11
	SyntaxKindLineCommentTrivia  SyntaxKind = 12
	SyntaxKindBlockCommentTrivia SyntaxKind = 13
	SyntaxKindLineBreakTrivia    SyntaxKind = 14
	SyntaxKindTrivia             SyntaxKind = 15
	SyntaxKindUnknown            SyntaxKind = 16
	SyntaxKindEOF                SyntaxKind = 17
)

type ParseErrorCode int

const (
	ParseErrorCodeInvalidSymbol          ParseErrorCode = 1
	ParseErrorCodeInvalidNumberFormat    ParseErrorCode = 2
	ParseErrorCodePropertyNameExpected   ParseErrorCode = 3
	ParseErrorCodeValueExpected          ParseErrorCode = 4
	ParseErrorCodeColonExpected          ParseErrorCode = 5
	ParseErrorCodeCommaExpected          ParseErrorCode = 6
	ParseErrorCodeCloseBraceExpected     ParseErrorCode = 7
	ParseErrorCodeCloseBracketExpected   ParseErrorCode = 8
	ParseErrorCodeEndOfFileExpected      ParseErrorCode = 9
	ParseErrorCodeInvalidCommentToken    ParseErrorCode = 10
	ParseErrorCodeUnexpectedEndOfComment ParseErrorCode = 11
	ParseErrorCodeUnexpectedEndOfString  ParseErrorCode = 12
	ParseErrorCodeUnexpectedEndOfNumber  ParseErrorCode = 13
	ParseErrorCodeInvalidUnicode         ParseErrorCode = 14
	ParseErrorCodeInvalidEscapeCharacter ParseErrorCode = 15
	ParseErrorCodeInvalidCharacter       ParseErrorCode = 16
)

type ParseError struct {
	Error  ParseErrorCode
	Offset int
	Length int
}

type ParseOptions struct {
	DisallowComments   bool
	AllowTrailingComma bool
	AllowEmptyContent  bool
}

type NodeType string

const (
	NodeTypeObject   NodeType = "object"
	NodeTypeArray    NodeType = "array"
	NodeTypeProperty NodeType = "property"
	NodeTypeString   NodeType = "string"
	NodeTypeNumber   NodeType = "number"
	NodeTypeBoolean  NodeType = "boolean"
	NodeTypeNull     NodeType = "null"
)

type Node struct {
	Type        NodeType
	Value       any
	Offset      int
	Length      int
	ColonOffset int
	Parent      *Node `json:"-"` // json.Unmarshal will stuck if not excluded
	Children    []*Node
}
