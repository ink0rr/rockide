package handlers

import "strings"

func ToJsonPath(path string) []string {
	return strings.Split(path, "/")
}
