package bog

import (
	"bytes"
	"errors"
	"os"
	"time"
)

var (
	errBadFileDescriptor = errors.New("bad file descriptor")
	errIsDirectory       = errors.New("is a directory")
	errInvalid           = errors.New("invalid argument")
	errNotFound          = errors.New("no such file or directory")
)

type File interface {
	Close() error
	Name() string
	Read(b []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	Readdir(n int) (fi []os.FileInfo, err error)
	Readdirnames(n int) (names []string, err error)
	Seek(offset int64, whence int) (ret int64, err error)
	Stat() (fi os.FileInfo, err error)
}

type FileInfo struct {
	FileName    string
	FileSize    int64
	FileMode    os.FileMode
	FileModTime time.Time
}

func (fi *FileInfo) Name() string {
	return fi.FileName
}

func (fi *FileInfo) Size() int64 {
	return fi.FileSize
}

func (fi *FileInfo) Mode() os.FileMode {
	return fi.FileMode
}

func (fi *FileInfo) ModTime() time.Time {
	return fi.FileModTime
}

func (fi *FileInfo) IsDir() bool {
	return fi.FileMode.IsDir()
}

func (fu *FileInfo) Sys() interface{} {
	return nil
}

type bogFile struct {
	data     []byte
	r        *bytes.Reader
	stat     os.FileInfo
	children []File
	off      int
	closed   bool
}

func NewBogFile(data []byte, info os.FileInfo) File {
	return &bogFile{
		data: data,
		r:    bytes.NewReader(data),
		stat: info,
	}
}

func NewBogFolder(children []File, info os.FileInfo) File {
	return &bogFile{
		stat:     info,
		children: children,
	}
}

func (f *bogFile) Close() error {
	f.closed = true
	return nil
}

func (f *bogFile) Name() string {
	return f.stat.Name()
}

func (f *bogFile) Read(b []byte) (n int, err error) {
	if f.closed {
		return 0, &os.PathError{"read", f.Name(), errBadFileDescriptor}
	}
	if f.stat.IsDir() {
		return 0, &os.PathError{"read", f.Name(), errIsDirectory}
	}
	return f.r.Read(b)
}

func (f *bogFile) ReadAt(b []byte, off int64) (n int, err error) {
	if f.closed {
		return 0, &os.PathError{"read", f.Name(), errBadFileDescriptor}
	}
	if f.stat.IsDir() {
		return 0, &os.PathError{"read", f.Name(), errIsDirectory}
	}
	return f.r.ReadAt(b, off)
}

func (f *bogFile) Readdir(n int) (fi []os.FileInfo, err error) {
	if f.closed {
		return nil, &os.PathError{"readdirent", f.Name(), errBadFileDescriptor}
	}
	if !f.stat.IsDir() {
		return nil, &os.PathError{"readdirent", f.Name(), errInvalid}
	}
	return nil, nil
}

func (f *bogFile) Readdirnames(n int) (names []string, err error) {
	if f.closed {
		return nil, &os.PathError{"readdirent", f.Name(), errBadFileDescriptor}
	}
	if !f.stat.IsDir() {
		return nil, &os.PathError{"readdirent", f.Name(), errInvalid}
	}
	return nil, nil
}

func (f *bogFile) Seek(offset int64, whence int) (ret int64, err error) {
	if f.closed {
		return 0, &os.PathError{"seek", f.Name(), errBadFileDescriptor}
	}
	if f.stat.IsDir() {
		return 0, &os.PathError{"seek", f.Name(), errIsDirectory}
	}
	return f.r.Seek(offset, whence)
}

func (f *bogFile) Stat() (fi os.FileInfo, err error) {
	if f.closed {
		return nil, &os.PathError{"stat", f.Name(), errBadFileDescriptor}
	}
	return f.stat, nil
}
