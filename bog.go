package bog

import (
	"os"
	"path/filepath"
	"strings"
)

type Archive struct {
	files map[string]File
	dev   bool
	root  string
}

func NewArchive(files map[string]File, dev bool, root string) *Archive {
	return &Archive{
		files: files,
		dev:   dev,
		root:  root,
	}
}

func (a *Archive) Open(name string) (File, error) {
	name = strings.TrimLeft(name, "/")
	if a.dev {
		return os.Open(filepath.Join(a.root, name))
	}
	f, ok := a.files["/" + name]
	if !ok {
		return nil, &os.PathError{"open", name, errNotFound}
	}
	return f, nil
}

func (a *Archive) Stat(name string) (fi os.FileInfo, err error) {
	name =strings.TrimLeft(name, "/")
	if a.dev {
		return os.Stat(filepath.Join(a.root, name))
	}
	f, ok := a.files["/" + name]
	if !ok {
		return nil, &os.PathError{"stat", name, errNotFound}
	}
	return f.Stat()
}
