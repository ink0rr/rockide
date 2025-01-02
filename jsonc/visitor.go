package jsonc

import (
	"slices"
	"strconv"
)

type Path = []any // string | int

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

// Parses the given text and invokes the visitor functions for each object, array and literal reached.
func Visit(text string, visitor *Visitor, options *ParseOptions) any {
	if options == nil {
		options = &ParseOptions{AllowTrailingComma: false}
	}
	scanner := CreateScanner(text, false)
	// Important: Only pass copies of this to visitor functions to prevent accidental modification, and
	// to not affect visitor functions which stored a reference to a previous JSONPath
	jsonPath := Path{}
	// Depth of onXXXBegin() callbacks suppressed. onXXXEnd() decrements this if it isn't 0 already.
	// Callbacks are only called when this value is 0.
	suppressedCallbacks := 0

	toBeginVisit := func(visitFunction func(offset, length, startLine, startCharacter uint32, pathSupplier func() Path) bool) func() {
		if visitFunction == nil {
			return func() {}
		}
		return func() {
			if suppressedCallbacks > 0 {
				suppressedCallbacks++
			} else {
				cbReturn := visitFunction(scanner.GetTokenOffset(), scanner.GetTokenLength(), scanner.GetTokenStartLine(), scanner.GetTokenStartCharacter(), func() Path {
					return append(Path{}, jsonPath...)
				})
				if !cbReturn {
					suppressedCallbacks = 1
				}
			}
		}
	}

	toEndVisit := func(visitFunction func(offset, length, startLine, startCharacter uint32)) func() {
		if visitFunction == nil {
			return func() {}
		}
		return func() {
			if suppressedCallbacks > 0 {
				suppressedCallbacks--
			}
			if suppressedCallbacks == 0 {
				visitFunction(scanner.GetTokenOffset(), scanner.GetTokenLength(), scanner.GetTokenStartLine(), scanner.GetTokenStartCharacter())
			}
		}
	}

	onObjectBegin := toBeginVisit(visitor.OnObjectBegin)
	onObjectProperty := func(arg string) {
		if visitor.OnObjectProperty != nil && suppressedCallbacks == 0 {
			visitor.OnObjectProperty(arg, scanner.GetTokenOffset(), scanner.GetTokenLength(), scanner.GetTokenStartLine(), scanner.GetTokenStartCharacter(), func() Path {
				return append(Path{}, jsonPath...)
			})
		}
	}
	onObjectEnd := toEndVisit(visitor.OnObjectEnd)
	onArrayBegin := toBeginVisit(visitor.OnArrayBegin)
	onArrayEnd := toEndVisit(visitor.OnArrayEnd)
	onLiteralValue := func(arg any) {
		if visitor.OnLiteralValue != nil && suppressedCallbacks == 0 {
			visitor.OnLiteralValue(arg, scanner.GetTokenOffset(), scanner.GetTokenLength(), scanner.GetTokenStartLine(), scanner.GetTokenStartCharacter(), func() Path {
				return append(Path{}, jsonPath...)
			})
		}
	}
	onSeparator := func(arg string) {
		if visitor.OnSeparator != nil && suppressedCallbacks == 0 {
			visitor.OnSeparator(arg, scanner.GetTokenOffset(), scanner.GetTokenLength(), scanner.GetTokenStartLine(), scanner.GetTokenStartCharacter())
		}
	}
	onComment := func() {
		if visitor.OnComment != nil && suppressedCallbacks == 0 {
			visitor.OnComment(scanner.GetTokenOffset(), scanner.GetTokenLength(), scanner.GetTokenStartLine(), scanner.GetTokenStartCharacter())
		}
	}
	onError := func(arg ParseErrorCode) {
		if visitor.OnError != nil && suppressedCallbacks == 0 {
			visitor.OnError(arg, scanner.GetTokenOffset(), scanner.GetTokenLength(), scanner.GetTokenStartLine(), scanner.GetTokenStartCharacter())
		}
	}

	disallowComments := options.DisallowComments
	allowTrailingComma := options.AllowTrailingComma

	var parseValue func() bool

	handleError := func(errorCode ParseErrorCode, skipUntilAfter, skipUntil []SyntaxKind) {
		onError(errorCode)
		if len(skipUntilAfter) > 0 || len(skipUntil) > 0 {
			token := scanner.GetToken()
			for token != SyntaxKindEOF {
				if slices.Contains(skipUntilAfter, token) {
					scanner.Scan()
					break
				} else if slices.Contains(skipUntil, token) {
					break
				}
				token = scanner.Scan()
			}
		}
	}

	scanNext := func() SyntaxKind {
		for {
			token := scanner.Scan()
			switch scanner.GetTokenError() {
			case ScanErrorInvalidUnicode:
				handleError(ParseErrorCodeInvalidUnicode, nil, nil)
			case ScanErrorInvalidEscapeCharacter:
				handleError(ParseErrorCodeInvalidEscapeCharacter, nil, nil)
			case ScanErrorUnexpectedEndOfNumber:
				handleError(ParseErrorCodeUnexpectedEndOfNumber, nil, nil)
			case ScanErrorUnexpectedEndOfComment:
				if !disallowComments {
					handleError(ParseErrorCodeUnexpectedEndOfComment, nil, nil)
				}
			case ScanErrorUnexpectedEndOfString:
				handleError(ParseErrorCodeUnexpectedEndOfString, nil, nil)
			case ScanErrorInvalidCharacter:
				handleError(ParseErrorCodeInvalidCharacter, nil, nil)
			}
			switch token {
			case SyntaxKindLineCommentTrivia, SyntaxKindBlockCommentTrivia:
				if disallowComments {
					handleError(ParseErrorCodeInvalidCommentToken, nil, nil)
				} else {
					onComment()
				}
			case SyntaxKindUnknown:
				handleError(ParseErrorCodeInvalidSymbol, nil, nil)
			case SyntaxKindTrivia, SyntaxKindLineBreakTrivia:
				// Skip trivia tokens
			default:
				return token
			}
		}
	}

	parseString := func(isValue bool) bool {
		value := scanner.GetTokenValue()
		if isValue {
			onLiteralValue(value)
		} else {
			onObjectProperty(value)
			// add property name afterwards
			jsonPath = append(jsonPath, value)
		}
		scanNext()
		return true
	}

	parseLiteral := func() bool {
		switch scanner.GetToken() {
		case SyntaxKindNumericLiteral:
			tokenValue := scanner.GetTokenValue()
			value, err := strconv.ParseFloat(tokenValue, 64)
			if err != nil {
				handleError(ParseErrorCodeInvalidNumberFormat, nil, nil)
				value = 0
			}
			onLiteralValue(value)
		case SyntaxKindNullKeyword:
			onLiteralValue(nil)
		case SyntaxKindTrueKeyword:
			onLiteralValue(true)
		case SyntaxKindFalseKeyword:
			onLiteralValue(false)
		default:
			return false
		}
		scanNext()
		return true
	}

	parseProperty := func() bool {
		if scanner.GetToken() != SyntaxKindStringLiteral {
			handleError(ParseErrorCodePropertyNameExpected, nil, []SyntaxKind{SyntaxKindCloseBraceToken, SyntaxKindCommaToken})
			return false
		}
		parseString(false)
		if scanner.GetToken() == SyntaxKindColonToken {
			onSeparator(":")
			scanNext()
			if !parseValue() {
				handleError(ParseErrorCodeValueExpected, nil, []SyntaxKind{SyntaxKindCloseBraceToken, SyntaxKindCommaToken})
			}
		} else {
			handleError(ParseErrorCodeColonExpected, nil, []SyntaxKind{SyntaxKindCloseBraceToken, SyntaxKindCommaToken})
		}
		jsonPath = jsonPath[:len(jsonPath)-1]
		return true
	}

	parseObject := func() bool {
		onObjectBegin()
		scanNext()
		needsComma := false
		for scanner.GetToken() != SyntaxKindCloseBraceToken && scanner.GetToken() != SyntaxKindEOF {
			if scanner.GetToken() == SyntaxKindCommaToken {
				if !needsComma {
					handleError(ParseErrorCodeValueExpected, nil, nil)
				}
				onSeparator(",")
				scanNext()
				if scanner.GetToken() == SyntaxKindCloseBraceToken && allowTrailingComma {
					break
				}
			} else if needsComma {
				handleError(ParseErrorCodeCommaExpected, nil, nil)
			}
			if !parseProperty() {
				handleError(ParseErrorCodeValueExpected, nil, []SyntaxKind{SyntaxKindCloseBraceToken, SyntaxKindCommaToken})
			}
			needsComma = true
		}
		onObjectEnd()
		if scanner.GetToken() != SyntaxKindCloseBraceToken {
			handleError(ParseErrorCodeCloseBraceExpected, []SyntaxKind{SyntaxKindCloseBraceToken}, nil)
		} else {
			scanNext()
		}
		return true
	}

	parseArray := func() bool {
		onArrayBegin()
		scanNext()
		isFirstElement := true
		needsComma := false
		for scanner.GetToken() != SyntaxKindCloseBracketToken && scanner.GetToken() != SyntaxKindEOF {
			if scanner.GetToken() == SyntaxKindCommaToken {
				if !needsComma {
					handleError(ParseErrorCodeValueExpected, nil, nil)
				}
				onSeparator(",")
				scanNext()
				if scanner.GetToken() == SyntaxKindCloseBracketToken && allowTrailingComma {
					break
				}
			} else if needsComma {
				handleError(ParseErrorCodeCommaExpected, nil, nil)
			}
			if isFirstElement {
				jsonPath = append(jsonPath, 0)
				isFirstElement = false
			} else {
				index := len(jsonPath) - 1
				switch value := jsonPath[index].(type) {
				case int:
					jsonPath[index] = value + 1
				}
			}
			if !parseValue() {
				handleError(ParseErrorCodeValueExpected, nil, []SyntaxKind{SyntaxKindCloseBracketToken, SyntaxKindCommaToken})
			}
			needsComma = true
		}
		onArrayEnd()
		if !isFirstElement {
			jsonPath = jsonPath[:len(jsonPath)-1]
		}
		if scanner.GetToken() != SyntaxKindCloseBracketToken {
			handleError(ParseErrorCodeCloseBracketExpected, []SyntaxKind{SyntaxKindCloseBracketToken}, nil)
		} else {
			scanNext()
		}
		return true
	}

	parseValue = func() bool {
		switch scanner.GetToken() {
		case SyntaxKindOpenBracketToken:
			return parseArray()
		case SyntaxKindOpenBraceToken:
			return parseObject()
		case SyntaxKindStringLiteral:
			return parseString(true)
		default:
			return parseLiteral()
		}
	}

	scanNext()
	if scanner.GetToken() == SyntaxKindEOF {
		if options.AllowEmptyContent {
			return true
		}
		handleError(ParseErrorCodeValueExpected, nil, nil)
		return false
	}
	if !parseValue() {
		handleError(ParseErrorCodeValueExpected, nil, nil)
		return false
	}
	if scanner.GetToken() != SyntaxKindEOF {
		handleError(ParseErrorCodeEndOfFileExpected, nil, nil)
	}
	return true
}
