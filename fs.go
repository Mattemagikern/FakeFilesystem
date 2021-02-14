// +build !FakeFileSystem

package fs

import (
	"io/ioutil"
	"os"
)

func Mkdir(path string, permissions os.FileMode) error {
	return os.Mkdir(SanitizePath(path), permissions)
}

func MkdirAll(path string, permissions os.FileMode) error {
	return os.MkdirAll(SanitizePath(path), permissions)
}

func Stat(path string) (os.FileInfo, error) {
	return os.Stat(SanitizePath(path))
}

func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(SanitizePath(path))
}

func WriteFile(path string, data []byte, permissions os.FileMode) error {
	return ioutil.WriteFile(SanitizePath(path), data, permissions)
}

func Remove(path string) error {
	return os.Remove(SanitizePath(path))
}
