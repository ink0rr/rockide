package handlers

import (
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ink0rr/rockide/core"
	"github.com/ink0rr/rockide/internal/lang"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/protocol/semtok"
	"github.com/ink0rr/rockide/internal/sliceutil"
	"github.com/ink0rr/rockide/internal/textdocument"
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

type LangHandler struct {
	Pattern shared.Pattern
	Client  bool
}

func (l *LangHandler) GetPattern() string {
	return l.Pattern.ToString()
}

func (l *LangHandler) Parse(uri protocol.DocumentURI) error {
	if !l.Client {
		return nil
	}
	document, err := textdocument.GetOrReadFile(uri)
	if err != nil {
		return err
	}
	content := document.GetContent()
	parser := lang.NewParser(content)
	root := parser.Parse()
	for _, entry := range root.Children() {
		if entry.Kind != lang.NodeEntry {
			continue
		}
		children := entry.Children()
		if len(children) < 2 {
			continue
		}
		keyNode := children[0]
		symbolRange := &protocol.Range{
			Start: document.PositionAt(keyNode.Offset),
			End:   document.PositionAt(keyNode.Offset + uint32(len(keyNode.Value))),
		}
		key := children[0].Value
		stores.Lang.Source.Insert(defaultScope, core.Symbol{
			Value: key,
			URI:   uri,
			Range: symbolRange,
		})
	}
	return nil
}

func (l *LangHandler) Delete(uri protocol.DocumentURI) {
	stores.Lang.Source.Delete(uri)
}

func (l *LangHandler) Completions(document *textdocument.TextDocument, position protocol.Position) []protocol.CompletionItem {
	if !l.Client {
		return nil
	}
	content := document.GetContent()
	parser := lang.NewParser(content)
	root := parser.Parse()
	node := lang.NodeAt(root, position)
	if node == nil {
		return nil
	}
	editRange := protocol.Range{
		Start: position,
		End:   position,
	}
	if parent := node.Parent(); parent != nil && parent.Kind == lang.NodeValue {
		offset := document.OffsetAt(position)
		if offset > 0 {
			if prev := content[offset-1]; prev == lang.SectionSign || prev == ' ' {
				if prev == lang.SectionSign {
					editRange.Start.Character -= 1
				}
				res := []protocol.CompletionItem{}
				for r, item := range lang.FormatCodes {
					value := "§" + string(r)
					res = append(res, protocol.CompletionItem{
						Label: value + " - " + item.Description,
						TextEdit: &protocol.Or_CompletionItem_textEdit{
							Value: protocol.TextEdit{
								NewText: value,
								Range:   editRange,
							},
						},
					})
				}
				return res
			}
		}
	}

	if node.Kind != lang.NodeKey && node.Kind != lang.NodeFile {
		return nil
	}
	if node.Kind == lang.NodeKey {
		editRange.Start = document.PositionAt(node.Offset)
		editRange.End = document.PositionAt(node.Offset + uint32(len(node.Value)))
	}
	res := []protocol.CompletionItem{}
	set := mapset.NewThreadUnsafeSet[string]()
	uriSet := mapset.NewThreadUnsafeSet[protocol.DocumentURI]()
	add := func(value string) {
		if !set.ContainsOne(value) {
			set.Add(value)
			res = append(res, protocol.CompletionItem{
				Label: value,
				TextEdit: &protocol.Or_CompletionItem_textEdit{
					Value: protocol.TextEdit{
						NewText: value,
						Range:   editRange,
					},
				},
			})
		}
	}
	for _, item := range stores.Lang.Source.GetFrom(document.URI) {
		if !set.ContainsOne(item.Value) {
			set.Add(item.Value)
		}
	}
	for _, item := range stores.Lang.Source.Get() {
		add(item.Value)
	}
	for _, item := range stores.Lang.References.Get() {
		if !uriSet.ContainsOne(item.URI) {
			uriSet.Add(item.URI)
		}
		add(item.Value)
	}
	for _, item := range stores.EntityId.Source.Get() {
		add("entity." + item.Value + ".name")
	}
	for _, item := range stores.ItemId.Source.Get("block") {
		if !uriSet.ContainsOne(item.URI) {
			uriSet.Add(item.URI)
			add("tile." + item.Value)
		}
	}
	for _, item := range stores.ItemId.Source.Get() {
		if !uriSet.ContainsOne(item.URI) {
			uriSet.Add(item.URI)
			add("item." + item.Value)
		}
	}
	return res
}

func (l *LangHandler) Definitions(document *textdocument.TextDocument, position protocol.Position) []protocol.LocationLink {
	if !l.Client {
		return nil
	}
	content := document.GetContent()
	parser := lang.NewParser(content)
	root := parser.Parse()
	node := lang.NodeAt(root, position)
	if node == nil || node.Kind != lang.NodeKey {
		return nil
	}
	res := []protocol.LocationLink{}
	for _, item := range slices.Concat(
		stores.Lang.Source.Get(),
		stores.Lang.References.Get(),
	) {
		if item.Value != node.Value {
			continue
		}
		location := protocol.LocationLink{
			TargetURI: item.URI,
			OriginSelectionRange: &protocol.Range{
				Start: document.PositionAt(node.Offset),
				End:   document.PositionAt(node.Offset + uint32(len(node.Value))),
			},
		}
		if item.Range != nil {
			location.TargetRange = *item.Range
			location.TargetSelectionRange = *item.Range
		}
		res = append(res, location)
	}
	return res
}

func (l *LangHandler) PrepareRename(document *textdocument.TextDocument, position protocol.Position) *protocol.PrepareRenamePlaceholder {
	if !l.Client {
		return nil
	}
	content := document.GetContent()
	parser := lang.NewParser(content)
	root := parser.Parse()
	node := lang.NodeAt(root, position)
	if node == nil || node.Kind != lang.NodeKey {
		return nil
	}
	start := document.PositionAt(node.Offset)
	end := document.PositionAt(node.Offset + uint32(len(node.Value)))
	return &protocol.PrepareRenamePlaceholder{
		Range: protocol.Range{
			Start: start,
			End:   end,
		},
		Placeholder: node.Value,
	}
}

func (l *LangHandler) Rename(document *textdocument.TextDocument, position protocol.Position, newName string) *protocol.WorkspaceEdit {
	if !l.Client {
		return nil
	}
	content := document.GetContent()
	parser := lang.NewParser(content)
	root := parser.Parse()
	node := lang.NodeAt(root, position)
	if node == nil || node.Kind != lang.NodeKey {
		return nil
	}
	changes := make(map[protocol.DocumentURI][]protocol.TextEdit)
	for _, item := range slices.Concat(
		stores.Lang.Source.Get(),
		stores.Lang.References.Get(),
	) {
		if item.Value != node.Value {
			continue
		}
		edit := protocol.TextEdit{
			NewText: newName,
			Range:   *item.Range,
		}
		changes[item.URI] = append(changes[item.URI], edit)
	}
	return &protocol.WorkspaceEdit{Changes: changes}
}

func (l *LangHandler) Hover(document *textdocument.TextDocument, position protocol.Position) *protocol.Hover {
	content := document.GetContent()
	parser := lang.NewParser(content)
	root := parser.Parse()
	node := lang.NodeAt(root, position)
	if node == nil {
		return nil
	}
	var entry *lang.Node
	switch node.Kind {
	case lang.NodeKey, lang.NodeAssign, lang.NodeValue:
		entry = node.Parent()
	default:
		parent := node.Parent()
		if parent != nil {
			switch parent.Kind {
			case lang.NodeValue:
				entry = parent.Parent()
			case lang.NodeEntry, lang.NodeComment:
				entry = parent
			}
		}
	}
	if entry == nil {
		return nil
	}
	var result *protocol.Hover
	switch first := entry.Children()[0]; first.Kind {
	case lang.NodeComment:
		result = &protocol.Hover{
			Contents: protocol.MarkupContent{
				Kind:  protocol.PlainText,
				Value: strings.TrimPrefix(first.Value, "##"),
			},
		}
	case lang.NodeKey:
		value := "### " + first.Value + "\n\n"
		if len(entry.Children()) > 2 {
			nodeValue := entry.Children()[2]
			children := nodeValue.Children()
			if comment, ok := sliceutil.Find(children, func(n *lang.Node) bool {
				return n.Kind == lang.NodeComment
			}); ok {
				value += strings.TrimLeft(comment.Value, "#") + "\n\n---\n\n"
			} else {
				value += "---\n\n"
			}
			arr := sliceutil.Map(children, func(node *lang.Node) string {
				switch node.Kind {
				case lang.NodeText:
					return node.Value
				case lang.NodeLineBreak:
					return "\n\n"
				case lang.NodeEmoji, lang.NodeFormatSpecifier:
					return "`" + node.Value + "`"
				}
				return ""
			})
			value += strings.Join(arr, "")
		}
		result = &protocol.Hover{
			Contents: protocol.MarkupContent{
				Kind:  protocol.Markdown,
				Value: value,
			},
		}
	}
	return result
}

func (l *LangHandler) SemanticTokens(document *textdocument.TextDocument) *protocol.SemanticTokens {
	tokens := []semtok.Token{}
	content := document.GetContent()
	lexer := lang.NewLexer(content)
	colored := false
	colorType := semtok.TokColorDefault
	modifiers := []semtok.Modifier{}
	set := mapset.NewSet[semtok.Modifier]()
	reset := func() {
		colored = false
		colorType = semtok.TokColorDefault
		modifiers = []semtok.Modifier{}
		set.Clear()
	}
	for token := range lexer.Next() {
		switch token.Kind {
		case lang.TokenNewline:
			reset()
		case lang.TokenFormatCode:
			bytes := []byte(token.Value)
			b := bytes[len(bytes)-1]
			if tokenType, ok := langColorMap[b]; ok {
				colored = true
				colorType = tokenType
			} else if mod, ok := langModifierMap[b]; ok {
				colored = true
				if !set.Contains(mod) {
					set.Add(mod)
					modifiers = append(modifiers, mod)
				}
			} else if b == 'r' {
				reset()
			}
		default:
			tokenType, ok := langTokenMap[token.Kind]
			mod := []semtok.Modifier{}
			if !ok && colored {
				tokenType = colorType
				mod = modifiers
			}
			tokens = append(tokens, semtok.Token{
				Type:      tokenType,
				Line:      token.Start.Line,
				Start:     token.Start.Column,
				Len:       token.Length(),
				Modifiers: mod,
			})
		}
	}
	return &protocol.SemanticTokens{
		Data: semtok.Encode(tokens, tokenType, map[semtok.Modifier]bool{
			semtok.ModObfuscated: true,
			semtok.ModBold:       true,
			semtok.ModItalic:     true,
		}),
	}
}

var langTokenMap = map[lang.TokenKind]semtok.Type{
	lang.TokenKey:             semtok.TokVariable,
	lang.TokenAssign:          semtok.TokOperator,
	lang.TokenComment:         semtok.TokComment,
	lang.TokenEmoji:           semtok.TokKeyword,
	lang.TokenLineBreak:       semtok.TokKeyword,
	lang.TokenFormatSpecifier: semtok.TokMacro,
}

var langColorMap = map[byte]semtok.Type{
	'a': semtok.TokColorGreen,
	'b': semtok.TokColorAqua,
	'c': semtok.TokColorRed,
	'd': semtok.TokColorLightPurple,
	'e': semtok.TokColorYellow,
	'f': semtok.TokColorWhite,
	'g': semtok.TokColorMinecoinGold,
	'h': semtok.TokColorMaterialQuartz,
	'i': semtok.TokColorMaterialIron,
	'j': semtok.TokColorMaterialNetherite,
	'm': semtok.TokColorMaterialRedstone,
	'n': semtok.TokColorMaterialCopper,
	'p': semtok.TokColorMaterialGold,
	'q': semtok.TokColorMaterialEmerald,
	's': semtok.TokColorMaterialDiamond,
	't': semtok.TokColorMaterialLapisLazuli,
	'u': semtok.TokColorMaterialAmethyst,
	'v': semtok.TokColorMaterialResin,
	'0': semtok.TokColorBlack,
	'1': semtok.TokColorDarkBlue,
	'2': semtok.TokColorDarkGreen,
	'3': semtok.TokColorDarkAqua,
	'4': semtok.TokColorDarkRed,
	'5': semtok.TokColorDarkPurple,
	'6': semtok.TokColorGold,
	'7': semtok.TokColorGray,
	'8': semtok.TokColorDarkGray,
	'9': semtok.TokColorBlue,
}

var langModifierMap = map[byte]semtok.Modifier{
	'k': semtok.ModObfuscated,
	'l': semtok.ModBold,
	'o': semtok.ModItalic,
}

var ClientLang = &LangHandler{
	Pattern: shared.ClientLangGlob,
	Client:  true,
}

var Lang = &LangHandler{
	Pattern: shared.LangGlob,
}
