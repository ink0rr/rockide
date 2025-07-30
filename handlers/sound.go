package handlers

import (
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Sound = &Path{
	Pattern: shared.SoundGlob,
	Store:   stores.SoundPath,
}
