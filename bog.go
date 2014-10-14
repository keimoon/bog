package bog

import (
	"io"
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

func (a *Archive) Extract() error {
	var outputPrefix = "."
	if !a.isFile {
		root := filepath.Base(a.root)
		if root != "." && root != ".." {
			err := os.RemoveAll(root)
			if err != nil {
				return err
			}
			outputPrefix = root
		}
	} else {
		outputPrefix = filepath.Base(a.root)
	}
	for path, f := range a.files {
		stat, err := f.Stat()
		if err != nil {
			return err
		}
		if !stat.IsDir() {
			outputFilePath := filepath.Join(outputPrefix, path)
			outputFolder := filepath.Dir(outputFilePath)
			if outputFolder != "." {
				err = os.MkdirAll(outputFolder, 0755)
				if err != nil {
					return err
				}
			}
			newFile, err := os.Create(outputFilePath)
			if err != nil {
				return err
			}
			defer newFile.Close()
			_, err = io.Copy(newFile, f)
			if err != nil {
				return err
			}
			err = newFile.Chmod(stat.Mode())
			if err != nil {
				return err
			}
		} else {
			folder := filepath.Join(outputPrefix, path)
			err = os.MkdirAll(folder, 0755)
			if err != nil {
				return err
			}
			err = os.Chmod(folder, stat.Mode())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *Archive) formatName(name string) string {
	if a.isFile {
		return ""
	}
	return strings.TrimLeft(name, "/")
}
