// +build FakeFileSystem

package fs

import (
	"os"
)

func Mkdir(path string, permissions os.FileMode) error {
	return Fs.Mkdir(SanitizePath(path), permissions)
}

func MkdirAll(path string, permissions os.FileMode) error {
	return Fs.Mkdir(SanitizePath(path), permissions)
}

func Stat(path string) (os.FileInfo, error) {
	return Fs.Stat(SanitizePath(path))
}

func ReadFile(path string) ([]byte, error) {
	return Fs.ReadFile(SanitizePath(path))
}

func WriteFile(path string, data []byte, permissions os.FileMode) error {
	return Fs.WriteFile(SanitizePath(path), data, permissions)
}

func Remove(path string) error {
	return Fs.Remove(SanitizePath(path))
}

func FindProjectRoot() error {
	return nil
}

func Clean() {
	Fs.Clean()
}
