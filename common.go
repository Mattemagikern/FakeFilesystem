package fs

import (
	"os"
	"path/filepath"
)

func SanitizePath(path string) string {
	return filepath.Clean(filepath.ToSlash(ExpandHome(path)))
}

func ExpandHome(path string) string {
	if len(path) == 0 || path[0] != '~' {
		return path
	}
	home := os.Getenv("HOME")
	return filepath.Join(home, path[1:])
}
