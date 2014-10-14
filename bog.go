package bog

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Archive is representation os archived folders/files in Go code.
type Archive struct {
	files  map[string]File
	dev    bool
	isFile bool
	root   string
}

// NewArchive creates new Archive. This function is called by the generator.
func NewArchive(files map[string]File, dev bool, isFile bool, root string) *Archive {
	return &Archive{
		files:  files,
		dev:    dev,
		isFile: isFile,
		root:   root,
	}
}

// SetRoot sets the root folder when using development mode.
func (a *Archive) SetRoot(path string) {
	if a.dev {
		a.root = path
	}
}

// Open opens the named file for reading. If successful, methods on the returned file can be used for reading.
// If there is an error, it will be of type *PathError
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

// Stat returns a FileInfo describing the named file. If there is an error, it will be of type *PathError.
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

// ReadDir reads the directory named by dirname and returns a list of directory entries.
func (a *Archive) ReadDir(dirname string) ([]os.FileInfo, error) {
	dirname = a.formatName(dirname)
	if a.dev {
		return ioutil.ReadDir(filepath.Join(a.root, dirname))
	}
	f, ok := a.files["/"+dirname]
	if !ok {
		return nil, &os.PathError{"open", dirname, errNotFound}
	}
	return f.Readdir(-1)
}

// ReadFile reads the file named by filename and returns the contents. A successful call returns err == nil, not err == EOF. 
// Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported.
func (a *Archive) ReadFile(name string) ([]byte, error) {
	f, err := a.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

// Extract extracts the content of the archive to current folder.
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
