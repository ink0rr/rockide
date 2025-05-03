package handlers

import (
	"github.com/ink0rr/rockide/internal/molang"
	"github.com/ink0rr/rockide/internal/protocol/semtok"
)

var tokenType = map[semtok.Type]bool{
	semtok.TokNumber:     true,
	semtok.TokString:     true,
	semtok.TokMacro:      true,
	semtok.TokMethod:     true,
	semtok.TokType:       true,
	semtok.TokKeyword:    true,
	semtok.TokOperator:   true,
	semtok.TokEnumMember: true,
	semtok.TokComment:    false,
}

var tokenModifier = map[semtok.Modifier]bool{}

var molangTokenMap = map[molang.TokenKind]semtok.Type{
	molang.KindNumber:     semtok.TokNumber,
	molang.KindString:     semtok.TokString,
	molang.KindMacro:      semtok.TokMacro,
	molang.KindMethod:     semtok.TokMethod,
	molang.KindPrefix:     semtok.TokType,
	molang.KindKeyword:    semtok.TokKeyword,
	molang.KindOperator:   semtok.TokOperator,
	molang.KindParen:      semtok.TokEnumMember,
	molang.KindComma:      semtok.TokOperator,
	molang.KindWhitespace: semtok.TokComment,
	molang.KindUnknown:    semtok.TokComment,
}
