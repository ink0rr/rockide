package handlers

import (
	"github.com/ink0rr/rockide/shared"
	"github.com/ink0rr/rockide/stores"
)

var Structure = &Path{
	Pattern: shared.StructureGlob,
	Store:   stores.StructurePath,
}
