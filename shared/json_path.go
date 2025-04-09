package shared

import "github.com/ink0rr/rockide/internal/jsonc"

type JsonPath struct {
	IsKey bool
	Path  jsonc.Path
}

func JsonKey(path string) JsonPath {
	return JsonPath{IsKey: true, Path: jsonc.NewPath(path)}
}

func JsonValue(path string) JsonPath {
	return JsonPath{IsKey: false, Path: jsonc.NewPath(path)}
}
