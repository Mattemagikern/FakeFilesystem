// +build FakeFileSystem

package fs

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

type fileSystem interface {
	Open(name string) (*File, error)
	Stat(name string) (os.FileInfo, error)
	ReadFile(path string) (*File, error)
	WriteFile(path string, data []byte, permissions os.FileMode) error
	Mkdir(path string, permissions os.FileMode) error
}

type fakeFileSystem struct {
	sync.Mutex
	Fs map[string]*File
}

var Fs *fakeFileSystem = &fakeFileSystem{
	Fs: make(map[string]*File),
}

func (fs *fakeFileSystem) Open(name string) (*File, error) {
	if f, ok := fs.Fs[name]; ok {
		return f, nil
	}
	if err := fs.WriteFile(name, []byte{}, 0644); err != nil {
		return nil, err
	}

	return fs.Fs[name], nil
}

func (fs *fakeFileSystem) Stat(name string) (os.FileInfo, error) {
	fs.Lock()
	defer fs.Unlock()
	if f, ok := fs.Fs[name]; ok {
		return f.Stat()
	}
	return nil, os.ErrNotExist
}

func (fs *fakeFileSystem) WriteFile(path string, data []byte, permissions os.FileMode) error {
	fs.Lock()
	defer fs.Unlock()
	fs.Fs[path] = &File{
		path:    path,
		Content: data,
		size:    int64(len(data)),
		mode:    permissions,
		modTime: time.Now(),
	}
	return nil
}

func (fs *fakeFileSystem) ReadFile(path string) ([]byte, error) {
	fs.Lock()
	defer fs.Unlock()
	if _, ok := fs.Fs[path]; !ok {
		return []byte{}, errors.New(fmt.Sprintf("%s: File Not Found", path))
	}
	return fs.Fs[path].Content, nil

}

func (fs *fakeFileSystem) Mkdir(path string, permissions os.FileMode) error {
	fs.Lock()
	defer fs.Unlock()
	if _, ok := fs.Fs[path]; ok {
		return errors.New(fmt.Sprintf("%s: File exists", path))
	}
	fs.Fs[path] = &File{
		path:    path,
		mode:    permissions,
		isDir:   true,
		modTime: time.Now(),
	}
	return nil
}

func (fs *fakeFileSystem) Remove(path string) error {
	fs.Lock()
	defer fs.Unlock()
	delete(fs.Fs, path)
	return nil
}

func (fs *fakeFileSystem) Tree() {
	for k, v := range fs.Fs {
		fmt.Println(k, v)
	}
}

func (fs *fakeFileSystem) Clean() {
	fs.Lock()
	defer fs.Unlock()
	fs.Fs = make(map[string]*File)
}
