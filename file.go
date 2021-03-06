package bog

import (
	"bytes"
	"errors"
	"io"
	"os"
	"time"
)

var (
	errBadFileDescriptor = errors.New("bad file descriptor")
	errIsDirectory       = errors.New("is a directory")
	errInvalid           = errors.New("invalid argument")
	errNotFound          = errors.New("no such file or directory")
)

// File represents an open file in the archive.
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

// A FileInfo describes a file, and implements os.FileInfo interface.
type FileInfo struct {
	FileName    string
	FileSize    int64
	FileMode    os.FileMode
	FileModTime time.Time
}

// Name returns base name of the file.
func (fi *FileInfo) Name() string {
	return fi.FileName
}

// Size returns length in bytes for regular files; system-dependent for others.
func (fi *FileInfo) Size() int64 {
	return fi.FileSize
}

// Mode returns file mode bits
func (fi *FileInfo) Mode() os.FileMode {
	return fi.FileMode
}

// ModTime returns modification time
func (fi *FileInfo) ModTime() time.Time {
	return fi.FileModTime
}

// IsDir is abbreviation for Mode().IsDir()
func (fi *FileInfo) IsDir() bool {
	return fi.FileMode.IsDir()
}

// Sys is always nil
func (fi *FileInfo) Sys() interface{} {
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

// NewBogFile creates a File. Use internally by generator.
func NewBogFile(data []byte, info os.FileInfo) File {
	return &bogFile{
		data: data,
		r:    bytes.NewReader(data),
		stat: info,
	}
}

// NewBogFolder creates a File that represents a folder. Use internally by generator.
func NewBogFolder(children []File, info os.FileInfo) File {
	return &bogFile{
		stat:     info,
		children: children,
	}
}

func (f *bogFile) Close() error {
	f.closed = true
	f.off = 0
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
	if f.off >= len(f.children) {
		return nil, io.EOF
	}
	var children []File
	if n < 0 || f.off+n >= len(f.children) {
		children = f.children[f.off:]
		f.off = len(f.children)
	} else {
		children = f.children[f.off : f.off+n]
		f.off += n
	}
	for _, child := range children {
		stat, err := child.Stat()
		if err != nil {
			return nil, err
		}
		fi = append(fi, stat)
	}
	return fi, nil
}

func (f *bogFile) Readdirnames(n int) (names []string, err error) {
	if f.closed {
		return nil, &os.PathError{"readdirent", f.Name(), errBadFileDescriptor}
	}
	if !f.stat.IsDir() {
		return nil, &os.PathError{"readdirent", f.Name(), errInvalid}
	}
	children, err := f.Readdir(n)
	if err != nil {
		return nil, err
	}
	for _, child := range children {
		names = append(names, child.Name())
	}
	return names, nil
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
