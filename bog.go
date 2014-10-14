package bog

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Archive struct {
	files  map[string]File
	dev    bool
	isFile bool
	root   string
}

func NewArchive(files map[string]File, dev bool, isFile bool, root string) *Archive {
	return &Archive{
		files:  files,
		dev:    dev,
		isFile: isFile,
		root:   root,
	}
}

func (a *Archive) SetRoot(path string) {
	if a.dev {
		a.root = path
	}
}

func (a *Archive) Open(name string) (File, error) {
	name = a.formatName(name)
	if a.dev {
		return os.Open(filepath.Join(a.root, name))
	}
	f, ok := a.files["/"+name]
	if !ok {
		return nil, &os.PathError{"open", name, errNotFound}
	}
	return f, nil
}

func (a *Archive) Stat(name string) (fi os.FileInfo, err error) {
	name = a.formatName(name)
	if a.dev {
		return os.Stat(filepath.Join(a.root, name))
	}
	f, ok := a.files["/"+name]
	if !ok {
		return nil, &os.PathError{"stat", name, errNotFound}
	}
	return f.Stat()
}

func (a *Archive) ReadDir(name string) ([]os.FileInfo, error) {
	name = a.formatName(name)
	if a.dev {
		return ioutil.ReadDir(filepath.Join(a.root, name))
	}
	f, ok := a.files["/"+name]
	if !ok {
		return nil, &os.PathError{"open", name, errNotFound}
	}
	return f.Readdir(-1)
}

func (a *Archive) ReadFile(name string) ([]byte, error) {
	f, err := a.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func (a *Archive) formatName(name string) string {
	if a.isFile {
		return ""
	}
	return strings.TrimLeft(name, "/")
}
