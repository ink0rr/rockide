package filesystem

import (
	"fmt"
	"io/fs"
	"os"
)

func ReadDir(path string) error {
	dir := os.DirFS(".")
	matches, err := fs.Glob(dir, "")
	if err != nil {
		return err
	}
	fmt.Println(matches[0])
	return nil
}
