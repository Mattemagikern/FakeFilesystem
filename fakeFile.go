// +build FakeFileSystem

package fs

import (
	"os"
	"time"
)

type File struct {
	path    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
	sys     *fakeFileSystem
	Content []byte
}

type FileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
	sys     interface{}
}

func (f *File) Stat() (os.FileInfo, error) {
	fi := &FileInfo{
		size: int64(len(f.Content)),
	}
	return fi, nil
}

func (s *File) Read(p []byte) (n int, err error) {
	for _, c := range s.Content {
		p = append(p, byte(c))
	}
	return len(s.Content), nil
}

func (f *File) Close() error {
	return nil
}

func (f *File) ReadAt(p []byte, offset int64) (int, error) {
	/*Not implemented*/
	return 0, nil
}

func (f *File) Seek(offset int64, whence int) (int64, error) {
	/*Not implemented*/
	return 0, nil
}

func (fi *FileInfo) Name() string {
	return fi.name
}

func (fi *FileInfo) Size() int64 {
	return fi.size
}

func (fi *FileInfo) Mode() os.FileMode {
	return fi.mode
}

func (fi *FileInfo) ModTime() time.Time {
	return fi.modTime
}

func (fi *FileInfo) IsDir() bool {
	return fi.isDir
}

func (fi *FileInfo) Sys() interface{} {
	return fi.sys
}
