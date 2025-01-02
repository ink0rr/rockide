package jsonc

import (
	"fmt"
)

// The scanner object, representing a JSON scanner at a position in the input string.
type Scanner interface {
	// Sets the scan position to a new offset. A call to 'scan' is needed to get the first token.
	SetPosition(pos uint32)
	// Read the next token. Returns the token code.
	Scan() SyntaxKind
	// Returns the zero-based current scan position, which is after the last read token.
	GetPosition() uint32
	// Returns the last read token.
	GetToken() SyntaxKind
	// Returns the last read token value. The value for strings is the decoded string content. For numbers it's of type number, for boolean it's true or false.
	GetTokenValue() string
	// The zero-based start offset of the last read token.
	GetTokenOffset() uint32
	// The length of the last read token.
	GetTokenLength() uint32
	// The zero-based start line number of the last read token.
	GetTokenStartLine() uint32
	// The zero-based start character (column) of the last read token.
	GetTokenStartCharacter() uint32
	// An error code of the last scan.
	GetTokenError() ScanError
}

type scanner struct {
	text         string
	ignoreTrivia bool

	textLength               uint32
	token                    SyntaxKind
	scanError                ScanError
	pos                      uint32
	value                    string
	tokenOffset              uint32
	lineNumber               uint32
	lineStartOffset          uint32
	tokenLineStartOffset     uint32
	prevTokenLineStartOffset uint32
}

// Creates a JSON scanner on the given text.
// If ignoreTrivia is set, whitespaces or comments are ignored.
func CreateScanner(text string, ignoreTrivia bool) Scanner {
	return &scanner{
		text:         text,
		ignoreTrivia: ignoreTrivia,

		textLength: uint32(len(text)),
		token:      SyntaxKindUnknown,
		scanError:  ScanErrorNone,
	}
}

func (s *scanner) scanHexDigits(count uint32, exact bool) (byte, error) {
	digits := uint32(0)
	value := byte(0)
	for digits < count || !exact {
		ch := charAt(s.text, s.pos)
		if ch >= '0' && ch <= '9' {
			value = value*16 + ch - '0'
		} else if ch >= 'A' && ch <= 'F' {
			value = value*16 + ch - 'A' + 10
		} else if ch >= 'a' && ch <= 'f' {
			value = value*16 + ch - 'a' + 10
		} else {
			break
		}
		s.pos++
		digits++
	}
	if digits < count {
		return 0, fmt.Errorf("Expected %d digits in hex number, but got %d", count, digits)
	}
	return value, nil
}

func (s *scanner) scanNumber() string {
	start := s.pos
	if charAt(s.text, s.pos) == '0' {
		s.pos++
	} else {
		s.pos++
		for s.pos < s.textLength && isDigit(charAt(s.text, s.pos)) {
			s.pos++
		}
	}
	if s.pos < s.textLength && charAt(s.text, s.pos) == '.' {
		s.pos++
		if s.pos < s.textLength && isDigit(charAt(s.text, s.pos)) {
			s.pos++
			for s.pos < s.textLength && isDigit(charAt(s.text, s.pos)) {
				s.pos++
			}
		} else {
			s.scanError = ScanErrorUnexpectedEndOfNumber
			return substring(s.text, start, s.pos)
		}
	}
	end := s.pos
	if s.pos < s.textLength && (charAt(s.text, s.pos) == 'E' || charAt(s.text, s.pos) == 'e') {
		s.pos++
		if s.pos < s.textLength && (charAt(s.text, s.pos) == '+' || charAt(s.text, s.pos) == '-') {
			s.pos++
		}
		if s.pos < s.textLength && isDigit(charAt(s.text, s.pos)) {
			s.pos++
			for s.pos < s.textLength && isDigit(charAt(s.text, s.pos)) {
				s.pos++
			}
			end = s.pos
		} else {
			s.scanError = ScanErrorUnexpectedEndOfNumber
		}
	}
	return substring(s.text, start, end)
}

func (s *scanner) scanString() string {
	result := ""
	start := s.pos
	for {
		if s.pos >= s.textLength {
			result += substring(s.text, start, s.pos)
			s.scanError = ScanErrorUnexpectedEndOfString
			break
		}
		ch := charAt(s.text, s.pos)
		if ch == '"' {
			result += substring(s.text, start, s.pos)
			s.pos++
			break
		}
		if ch == '\\' {
			result += substring(s.text, start, s.pos)
			s.pos++
			if s.pos >= s.textLength {
				s.scanError = ScanErrorUnexpectedEndOfString
				break
			}
			ch2 := charAt(s.text, s.pos)
			s.pos++
			switch ch2 {
			case '"':
				result += "\""
			case '\\':
				result += "\\"
			case '/':
				result += "/"
			case 'b':
				result += "\b"
			case 'f':
				result += "\f"
			case 'n':
				result += "\n"
			case 'r':
				result += "\r"
			case 't':
				result += "\t"
			case 'u':
				ch3, err := s.scanHexDigits(4, true)
				if err != nil {
					s.scanError = ScanErrorInvalidUnicode
				} else {
					result += string(ch3)
				}
				break
			default:
				s.scanError = ScanErrorInvalidEscapeCharacter
			}
			start = s.pos
			continue
		}
		if ch >= 0 && ch <= 0x1f {
			if isLineBreak(ch) {
				result += substring(s.text, start, s.pos)
				s.scanError = ScanErrorUnexpectedEndOfString
				break
			} else {
				s.scanError = ScanErrorInvalidCharacter
				// mark as error but continue with string
			}
		}
		s.pos++
	}
	return result
}

func (s *scanner) scanNext() SyntaxKind {
	s.value = ""
	s.scanError = ScanErrorNone

	s.tokenOffset = s.pos
	s.lineStartOffset = s.lineNumber
	s.prevTokenLineStartOffset = s.tokenLineStartOffset

	if s.pos >= s.textLength {
		// at the end
		s.tokenOffset = s.textLength
		s.token = SyntaxKindEOF
		return s.token
	}

	code := charAt(s.text, s.pos)
	// trivia: whitespace
	if isWhiteSpace(code) {
		for {
			s.pos++
			s.value += string(code)
			code = charAt(s.text, s.pos)
			if !isWhiteSpace(code) {
				break
			}
		}
		s.token = SyntaxKindTrivia
		return s.token
	}
	// trivia: newlines
	if isLineBreak(code) {
		s.pos++
		s.value += string(code)
		if code == '\r' && charAt(s.text, s.pos) == '\n' {
			s.pos++
			s.value += "\n"
		}
		s.lineNumber++
		s.tokenLineStartOffset = s.pos
		s.token = SyntaxKindLineBreakTrivia
		return s.token
	}

	switch code {
	// tokens: []{}:,
	case '{':
		s.pos++
		s.token = SyntaxKindOpenBraceToken
		return s.token
	case '}':
		s.pos++
		s.token = SyntaxKindCloseBraceToken
		return s.token
	case '[':
		s.pos++
		s.token = SyntaxKindOpenBracketToken
		return s.token
	case ']':
		s.pos++
		s.token = SyntaxKindCloseBracketToken
		return s.token
	case ':':
		s.pos++
		s.token = SyntaxKindColonToken
		return s.token
	case ',':
		s.pos++
		s.token = SyntaxKindCommaToken
		return s.token
	// strings
	case '"':
		s.pos++
		s.value = s.scanString()
		s.token = SyntaxKindStringLiteral
		return s.token
	// comments
	case '/':
		start := s.pos - 1
		if s.pos == 0 {
			start = 0
		}
		// Single-line comment
		if charAt(s.text, s.pos+1) == '/' {
			s.pos += 2
			for s.pos < s.textLength {
				if isLineBreak(charAt(s.text, s.pos)) {
					break
				}
				s.pos++
			}
			s.value = substring(s.text, start, s.pos)
			s.token = SyntaxKindLineCommentTrivia
			return s.token
		}
		// Multi-line comment
		if charAt(s.text, s.pos+1) == '*' {
			s.pos += 2
			safeLength := s.textLength - 1 // For lookahead.
			if s.textLength == 0 {
				safeLength = 0
			}
			commentClosed := false
			for s.pos < safeLength {
				ch := charAt(s.text, s.pos)
				if ch == '*' && charAt(s.text, s.pos+1) == '/' {
					s.pos += 2
					commentClosed = true
					break
				}
				s.pos++
				if isLineBreak(ch) {
					if ch == '\r' && charAt(s.text, s.pos) == '\n' {
						s.pos++
					}
					s.lineNumber++
					s.tokenLineStartOffset = s.pos
				}
			}
			if !commentClosed {
				s.pos++
				s.scanError = ScanErrorUnexpectedEndOfComment
			}
			s.value = substring(s.text, start, s.pos)
			s.token = SyntaxKindBlockCommentTrivia
			return s.token
		}
		// just a single slash
		s.value += string(code)
		s.pos++
		s.token = SyntaxKindUnknown
		return s.token
	// numbers
	case '-':
		s.value += string(code)
		s.pos++
		if s.pos == s.textLength || !isDigit(charAt(s.text, s.pos)) {
			s.token = SyntaxKindUnknown
			return s.token
		}
		fallthrough
	// found a minus, followed by a number so
	// we fall through to proceed with scanning
	// numbers
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		s.value += s.scanNumber()
		s.token = SyntaxKindNumericLiteral
		return s.token
	// literals and unknown symbols
	default:
		// is a literal? Read the full word.
		for s.pos < s.textLength && isUnknownContentCharacter(code) {
			s.pos++
			code = charAt(s.text, s.pos)
		}
		if s.tokenOffset != s.pos {
			s.value = substring(s.text, s.tokenOffset, s.pos)
			// keywords: true, false, null
			switch s.value {
			case "true":
				s.token = SyntaxKindTrueKeyword
			case "false":
				s.token = SyntaxKindFalseKeyword
			case "null":
				s.token = SyntaxKindNullKeyword
			default:
				s.token = SyntaxKindUnknown
			}
			return s.token
		}
		// some
		s.value += string(code)
		s.pos++
		s.token = SyntaxKindUnknown
		return s.token
	}
}

func (s *scanner) scanNextNonTrivia() SyntaxKind {
	result := s.scanNext()
	for result >= SyntaxKindLineCommentTrivia && result <= SyntaxKindTrivia {
		result = s.scanNext()
	}
	return result
}

func (s *scanner) SetPosition(pos uint32) {
	s.pos = pos
	s.value = ""
	s.tokenOffset = 0
	s.token = SyntaxKindUnknown
	s.scanError = ScanErrorNone
}

func (s *scanner) GetPosition() uint32 {
	return s.pos
}

func (s *scanner) Scan() SyntaxKind {
	if s.ignoreTrivia {
		return s.scanNextNonTrivia()
	}
	return s.scanNext()
}

func (s *scanner) GetToken() SyntaxKind {
	return s.token
}

func (s *scanner) GetTokenValue() string {
	return s.value
}

func (s *scanner) GetTokenOffset() uint32 {
	return s.tokenOffset
}

func (s *scanner) GetTokenLength() uint32 {
	return s.pos - s.tokenOffset
}

func (s *scanner) GetTokenStartLine() uint32 {
	return s.lineStartOffset
}

func (s *scanner) GetTokenStartCharacter() uint32 {
	return s.tokenOffset - s.prevTokenLineStartOffset
}

func (s *scanner) GetTokenError() ScanError {
	return s.scanError
}

var _ Scanner = (*scanner)(nil)
