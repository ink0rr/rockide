package handlers

import (
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Texture = &Path{
	Pattern: shared.TextureGlob,
	Store:   stores.TexturePath,
}
