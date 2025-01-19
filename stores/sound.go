package stores

import "github.com/ink0rr/rockide/shared"

var Sound = &PathStore{pattern: shared.SoundGlob, trimSuffix: true}
