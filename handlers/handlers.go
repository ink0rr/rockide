package handlers

import (
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/ink0rr/rockide/internal/protocol"
	"github.com/ink0rr/rockide/internal/textdocument"
)

type Handler interface {
	GetPattern() string
	Parse(uri protocol.DocumentURI) error
	Delete(uri protocol.DocumentURI)
}

type CompletionProvider interface {
	Completions(document *textdocument.TextDocument, position protocol.Position) []protocol.CompletionItem
}

type DefinitionProvider interface {
	Definitions(document *textdocument.TextDocument, position protocol.Position) []protocol.LocationLink
}

type RenameProvider interface {
	PrepareRename(document *textdocument.TextDocument, position protocol.Position) *protocol.PrepareRenamePlaceholder
	Rename(document *textdocument.TextDocument, position protocol.Position, newName string) *protocol.WorkspaceEdit
}

type HoverProvider interface {
	Hover(document *textdocument.TextDocument, position protocol.Position) *protocol.Hover
}

type SignatureHelpProvider interface {
	SignatureHelp(document *textdocument.TextDocument, position protocol.Position) *protocol.SignatureHelp
}

type SemanticTokenProvider interface {
	SemanticTokens(document *textdocument.TextDocument) *protocol.SemanticTokens
}

var handlerList = []Handler{
	// BP
	AimAssistPreset,
	AimAssistCategory,
	Animation,
	AnimationController,
	Biome,
	Block,
	Camera,
	CraftingItemCatalog,
	Dialogue,
	Entity,
	Feature,
	FeatureRule,
	Item,
	LootTable,
	Recipe,
	TradeTable,
	WorldgenProcessor,
	WorldgenTemplatePool,
	WorldgenJigsaw,
	WorldgenStructureSet,
	// RP
	Attachable,
	Atmosphere,
	BlockCulling,
	ClientAnimation,
	ClientAnimationController,
	ClientBiome,
	ClientBlock,
	ClientEntity,
	ClientSound,
	ColorGrading,
	EntityMaterial,
	Fog,
	FlipbookTexture,
	Geometry,
	ItemTexture,
	Lighting,
	LocalLighting,
	MusicDefintion,
	Particle,
	ParticleMaterial,
	RenderController,
	Sound,
	SoundDefinition,
	TerrainTexture,
	Texture,
	TextureSet,
	Water,
}

func GetAll() []Handler {
	return handlerList
}

func Find(uri protocol.DocumentURI) Handler {
	path := filepath.ToSlash(uri.Path())
	for _, handler := range handlerList {
		if doublestar.MatchUnvalidated("**/"+handler.GetPattern(), path) {
			return handler
		}
	}
	return nil
}
