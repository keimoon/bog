package main

import (
	"bufio"
	"fmt"
	"github.com/keimoon/bog"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func Archive() int {
	if len(Args) <= 1 {
		Usage()
		return 2
	}
	folder := Args[1]
	stat, err := os.Stat(folder)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	filePackageName := makePackageName(folder)
	var outputFileName string
	var varName string
	if len(filePackageName) > 0 {
		outputFileName = filePackageName + "-archive.go"
		varName = makePublicVariableName(filePackageName) + "Archive"
	} else {
		outputFileName = Options.PackageName + "-archive.go"
		varName = makePublicVariableName(Options.PackageName) + "Archive"
	}
	outputFolder := "."
	if !Options.isCwd && Options.PackageName != "main" {
		outputFolder = Options.PackageName
		err = os.RemoveAll(outputFolder)
		if err != nil {
			fmt.Println(err)
			return 1
		}
	} else {
		err = os.RemoveAll(outputFileName)
		if err != nil {
			fmt.Println(err)
			return 1
		}
	}
	fileVars := []*FileVar{}
	now := strconv.FormatInt(time.Now().Unix(), 10)
	if !Options.Dev {
		if stat.IsDir() {
			err = walk(folder, func(path string, info os.FileInfo, children []string) error {
				fileVar := &FileVar{
					VarName: makeVariableName(now + "_" + path),
					Path:    "/" + strings.TrimLeft(strings.TrimPrefix(path, folder), "/"),
					Stat: &bog.FileInfo{
						FileName:    info.Name(),
						FileSize:    info.Size(),
						FileMode:    info.Mode(),
						FileModTime: info.ModTime(),
					},
				}
				if info.IsDir() {
					fileVar.IsDir = true
					for _, child := range children {
						fileVar.Children = append(fileVar.Children, makeVariableName(now+"_"+path+"_"+child))
					}
				} else {
					b, err := ioutil.ReadFile(path)
					if err != nil {
						return err
					}
					fileVar.Data = b
				}
				fileVars = append(fileVars, fileVar)
				return nil
			}, ignoreRules{})
			if err != nil {
				fmt.Println(err)
				return 1
			}
		} else {
			fileVar := &FileVar{
				VarName: makeVariableName(now + "_" + folder),
				Path:    "/",
				Stat:    stat,
			}
			b, err := ioutil.ReadFile(folder)
			if err != nil {
				return 2
			}
			fileVar.Data = b
			fileVars = append(fileVars, fileVar)
		}
	}
	if !Options.isCwd {
		err = os.MkdirAll(outputFolder, 0755)
	}
	f, err := os.Create(filepath.Join(outputFolder, outputFileName))
	if err != nil {
		fmt.Println(err)
		return 1
	}
	defer f.Close()
	mainTmpl, err := loadTemplate("main.go.tmpl")
	if err != nil {
		fmt.Println(err)
		return 1
	}
	tmplData := &struct {
		PackageName string
		Files       []*FileVar
		Root        string
		VarName     string
		IsFile      bool
		Dev         bool
	}{
		PackageName: Options.PackageName,
		Files:       fileVars,
		Root:        folder,
		VarName:     varName,
		IsFile:      !stat.IsDir(),
		Dev:         Options.Dev,
	}
	err = mainTmpl.Execute(f, tmplData)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

type FileVar struct {
	VarName  string
	Path     string
	IsDir    bool
	Stat     os.FileInfo
	Data     []byte
	Children []string
}

type ignoreRules []string

func (r ignoreRules) ignore(name string) bool {
	if name == ".bogignore" {
		return true
	}
	for _, rule := range r {
		if matched, _ := filepath.Match(rule, name); matched {
			return true
		}
	}
	return false
}

type walkFunc func(path string, info os.FileInfo, children []string) error

func walk(root string, walkFn walkFunc, rules ignoreRules) error {
	info, err := os.Stat(root)
	if err != nil {
		return err
	}
	children := []string{}
	if info.IsDir() {
		f, err := os.Open(filepath.Join(root, ".bogignore"))
		if err == nil {
			defer f.Close()
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				rules = append(rules, scanner.Text())
			}
			if err = scanner.Err(); err != nil {
				return err
			}
		}
		childrenInfos, err := ioutil.ReadDir(root)
		if err != nil {
			return err
		}
		for _, childInfo := range childrenInfos {
			if rules.ignore(childInfo.Name()) {
				continue
			}
			err = walk(filepath.Join(root, childInfo.Name()), walkFn, rules)
			if err != nil {
				return err
			}
			children = append(children, childInfo.Name())
		}
	}
	return walkFn(root, info, children)
}

func loadTemplate(name string) (*template.Template, error) {
	f, err := TemplatesArchive.Open("main.go.tmpl")
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	t := template.New(name)
	return t.Parse(string(b))
}
