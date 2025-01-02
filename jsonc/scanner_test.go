package jsonc_test

import (
	"testing"

	"github.com/ink0rr/rockide/jsonc"
)

func assertKinds(t *testing.T, text string, kinds ...jsonc.SyntaxKind) {
	scanner := jsonc.CreateScanner(text, false)
	for {
		kind := scanner.Scan()
		if kind == jsonc.SyntaxKindEOF {
			break
		}
		if kind != kinds[0] {
			t.Errorf("Expected: %d, Actual: %d", kinds[0], kind)
		}
		kinds = kinds[1:]
		err := scanner.GetTokenError()
		if err != jsonc.ScanErrorNone {
			t.Errorf("Scan error: %d", err)
		}
	}
	if len(kinds) != 0 {
		t.Errorf("Unexpected EOF: %s", text)
	}
}

func assertScanError(t *testing.T, text string, scanError jsonc.ScanError, kinds ...jsonc.SyntaxKind) {
	scanner := jsonc.CreateScanner(text, false)
	scanner.Scan()
	if token := scanner.GetToken(); token != kinds[0] {
		t.Errorf("Expected: %d, Actual: %d", token, kinds[0])
	}
	kinds = kinds[1:]
	if err := scanner.GetTokenError(); err != scanError {
		t.Errorf("Expected: %d, Actual: %d", err, kinds[0])
	}
	for {
		kind := scanner.Scan()
		if kind == jsonc.SyntaxKindEOF {
			break
		}
		if kind != kinds[0] {
			t.Errorf("Expected: %d, Actual: %d", kinds[0], kind)
		}
		kinds = kinds[1:]
	}
	if len(kinds) != 0 {
		t.Errorf("Unexpected EOF: %s", text)
	}
}

func TestScanTokens(t *testing.T) {
	assertKinds(t, "{", jsonc.SyntaxKindOpenBraceToken)
	assertKinds(t, "}", jsonc.SyntaxKindCloseBraceToken)
	assertKinds(t, "[", jsonc.SyntaxKindOpenBracketToken)
	assertKinds(t, "]", jsonc.SyntaxKindCloseBracketToken)
	assertKinds(t, ":", jsonc.SyntaxKindColonToken)
	assertKinds(t, ",", jsonc.SyntaxKindCommaToken)
}

func TestScanComments(t *testing.T) {
	assertKinds(t, "// this is a comment", jsonc.SyntaxKindLineCommentTrivia)
	assertKinds(t, "// this is a comment\n", jsonc.SyntaxKindLineCommentTrivia, jsonc.SyntaxKindLineBreakTrivia)
	assertKinds(t, "/* this is a comment*/", jsonc.SyntaxKindBlockCommentTrivia)
	assertKinds(t, "/* this is a \r\ncomment*/", jsonc.SyntaxKindBlockCommentTrivia)
	assertKinds(t, "/* this is a \ncomment*/", jsonc.SyntaxKindBlockCommentTrivia)

	// unexpected end
	assertScanError(t, "/* this is a", jsonc.ScanErrorUnexpectedEndOfComment, jsonc.SyntaxKindBlockCommentTrivia)
	assertScanError(t, "/* this is a \ncomment", jsonc.ScanErrorUnexpectedEndOfComment, jsonc.SyntaxKindBlockCommentTrivia)

	// broken comment
	assertKinds(t, "/ ttt", jsonc.SyntaxKindUnknown, jsonc.SyntaxKindTrivia, jsonc.SyntaxKindUnknown)
}

func TestScanStrings(t *testing.T) {
	assertKinds(t, `"test"`, jsonc.SyntaxKindStringLiteral)
	assertKinds(t, `"\""`, jsonc.SyntaxKindStringLiteral)
	assertKinds(t, `"\/"`, jsonc.SyntaxKindStringLiteral)
	assertKinds(t, `"\b"`, jsonc.SyntaxKindStringLiteral)
	assertKinds(t, `"\f"`, jsonc.SyntaxKindStringLiteral)
	assertKinds(t, `"\n"`, jsonc.SyntaxKindStringLiteral)
	assertKinds(t, `"\r"`, jsonc.SyntaxKindStringLiteral)
	assertKinds(t, `"\t"`, jsonc.SyntaxKindStringLiteral)
	assertKinds(t, `"`+"\u88ff"+`"`, jsonc.SyntaxKindStringLiteral)
	assertKinds(t, `"`+"​\u2028"+`"`, jsonc.SyntaxKindStringLiteral)
	assertScanError(t, `"\v"`, jsonc.ScanErrorInvalidEscapeCharacter, jsonc.SyntaxKindStringLiteral)

	// unexpected end
	assertScanError(t, `"test`, jsonc.ScanErrorUnexpectedEndOfString, jsonc.SyntaxKindStringLiteral)
	assertScanError(t, `"test`+"\n"+`"`, jsonc.ScanErrorUnexpectedEndOfString, jsonc.SyntaxKindStringLiteral, jsonc.SyntaxKindLineBreakTrivia, jsonc.SyntaxKindStringLiteral)

	// invalid characters
	assertScanError(t, "\"\t\"", jsonc.ScanErrorInvalidCharacter, jsonc.SyntaxKindStringLiteral)
	assertScanError(t, "\"\t \"", jsonc.ScanErrorInvalidCharacter, jsonc.SyntaxKindStringLiteral)
}

func TestScanNumbers(t *testing.T) {
	assertKinds(t, "0", jsonc.SyntaxKindNumericLiteral)
	assertKinds(t, "0.1", jsonc.SyntaxKindNumericLiteral)
	assertKinds(t, "-0.1", jsonc.SyntaxKindNumericLiteral)
	assertKinds(t, "-1", jsonc.SyntaxKindNumericLiteral)
	assertKinds(t, "1", jsonc.SyntaxKindNumericLiteral)
	assertKinds(t, "123456789", jsonc.SyntaxKindNumericLiteral)
	assertKinds(t, "10", jsonc.SyntaxKindNumericLiteral)
	assertKinds(t, "90", jsonc.SyntaxKindNumericLiteral)
	assertKinds(t, "90E+123", jsonc.SyntaxKindNumericLiteral)
	assertKinds(t, "90e+123", jsonc.SyntaxKindNumericLiteral)
	assertKinds(t, "90e-123", jsonc.SyntaxKindNumericLiteral)
	assertKinds(t, "90E-123", jsonc.SyntaxKindNumericLiteral)
	assertKinds(t, "90E123", jsonc.SyntaxKindNumericLiteral)
	assertKinds(t, "90e123", jsonc.SyntaxKindNumericLiteral)

	// zero handling
	assertKinds(t, "01", jsonc.SyntaxKindNumericLiteral, jsonc.SyntaxKindNumericLiteral)
	assertKinds(t, "-01", jsonc.SyntaxKindNumericLiteral, jsonc.SyntaxKindNumericLiteral)

	// unexpected end
	assertKinds(t, "-", jsonc.SyntaxKindUnknown)
	assertKinds(t, ".0", jsonc.SyntaxKindUnknown)
}

func TestScanKeywords(t *testing.T) {
	assertKinds(t, "true", jsonc.SyntaxKindTrueKeyword)
	assertKinds(t, "false", jsonc.SyntaxKindFalseKeyword)
	assertKinds(t, "null", jsonc.SyntaxKindNullKeyword)

	assertKinds(t, "true false null",
		jsonc.SyntaxKindTrueKeyword,
		jsonc.SyntaxKindTrivia,
		jsonc.SyntaxKindFalseKeyword,
		jsonc.SyntaxKindTrivia,
		jsonc.SyntaxKindNullKeyword)

	// invalid words
	assertKinds(t, "nulllll", jsonc.SyntaxKindUnknown)
	assertKinds(t, "True", jsonc.SyntaxKindUnknown)
	assertKinds(t, "foo-bar", jsonc.SyntaxKindUnknown)
	assertKinds(t, "foo bar", jsonc.SyntaxKindUnknown, jsonc.SyntaxKindTrivia, jsonc.SyntaxKindUnknown)

	assertKinds(t, "false//hello", jsonc.SyntaxKindFalseKeyword, jsonc.SyntaxKindLineCommentTrivia)
}

func TestScanTrivia(t *testing.T) {
	assertKinds(t, " ", jsonc.SyntaxKindTrivia)
	assertKinds(t, "  \t  ", jsonc.SyntaxKindTrivia)
	assertKinds(t, "  \t  \n  \t  ", jsonc.SyntaxKindTrivia, jsonc.SyntaxKindLineBreakTrivia, jsonc.SyntaxKindTrivia)
	assertKinds(t, "\r\n", jsonc.SyntaxKindLineBreakTrivia)
	assertKinds(t, "\r", jsonc.SyntaxKindLineBreakTrivia)
	assertKinds(t, "\n", jsonc.SyntaxKindLineBreakTrivia)
	assertKinds(t, "\n\r", jsonc.SyntaxKindLineBreakTrivia, jsonc.SyntaxKindLineBreakTrivia)
	assertKinds(t, "\n   \n", jsonc.SyntaxKindLineBreakTrivia, jsonc.SyntaxKindTrivia, jsonc.SyntaxKindLineBreakTrivia)
}
