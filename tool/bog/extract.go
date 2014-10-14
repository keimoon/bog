package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/keimoon/bog"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strconv"
	"strings"
)

// Extract extracts content of an archived go source file to current folder
func Extract() int {
	fset := token.NewFileSet()
	sourceFile := Args[1]
	f, err := parser.ParseFile(fset, sourceFile, nil, 0)
	if err != nil {
		fmt.Println(err)
		return 2
	}
	files := make(map[string]bog.File)
	var archive *bog.Archive
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if genDecl.Tok != token.VAR {
			continue
		}
		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			if len(valueSpec.Names) <= 0 || len(valueSpec.Values) <= 0 {
				continue
			}
			varName := valueSpec.Names[0].Name
			callExpr, ok := valueSpec.Values[0].(*ast.CallExpr)
			if !ok {
				continue
			}
			selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
			if !ok {
				continue
			}
			if strings.HasPrefix(varName, "vvv") {
				var file bog.File
				if selectorExpr.Sel.Name == "NewBogFile" {
					file, err = createBogFile(callExpr.Args)
				} else {
					file, err = createBogFolder(files, callExpr.Args)
				}
				if err != nil {
					fmt.Println(err)
					return 2
				}
				files[varName] = file
			} else {
				archive, err = createArchive(files, callExpr.Args)
				if err != nil {
					fmt.Println(err)
					return 2
				}
			}
		}
	}
	err = archive.Extract()
	if err != nil {
		fmt.Println(err)
		return 2
	}
	return 0
}

func createBogFile(args []ast.Expr) (bog.File, error) {
	if len(args) != 2 {
		return nil, errors.New("malformed source file, NewBogFile has exactly 2 arguments")
	}
	firstArg, ok := args[0].(*ast.CompositeLit)
	if !ok {
		return nil, errors.New("malformed source file, first argument of NewBogFile must be a CompositeLit")
	}
	buffer := &bytes.Buffer{}
	for _, elt := range firstArg.Elts {
		basicLit, ok := elt.(*ast.BasicLit)
		if !ok {
			return nil, errors.New("malformed source file, data invalid")
		}
		val, err := strconv.ParseInt(basicLit.Value, 0, 64)
		if err != nil {
			return nil, err
		}
		err = buffer.WriteByte(byte(val))
		if err != nil {
			return nil, err
		}
	}
	stat, err := parseStat(args[1])
	if err != nil {
		return nil, err
	}
	return bog.NewBogFile(buffer.Bytes(), stat), nil
}

func createBogFolder(files map[string]bog.File, args []ast.Expr) (bog.File, error) {
	if len(args) != 2 {
		return nil, errors.New("malformed source file, NewBogFile has exactly 2 arguments")
	}
	firstArg, ok := args[0].(*ast.CompositeLit)
	if !ok {
		return nil, errors.New("malformed source file, first argument of NewBogFile must be a CompositeLit")
	}
	children := []bog.File{}
	for _, elt := range firstArg.Elts {
		ident, ok := elt.(*ast.Ident)
		if !ok {
			return nil, errors.New("malformed source file, file array invalid")
		}
		f, ok := files[ident.Name]
		if !ok {
			return nil, errors.New("malformed source file, cannot file ident " + ident.Name)
		}
		children = append(children, f)
	}
	stat, err := parseStat(args[1])
	if err != nil {
		return nil, err
	}
	return bog.NewBogFolder(children, stat), nil
}

func createArchive(files map[string]bog.File, args []ast.Expr) (*bog.Archive, error) {
	if len(args) != 4 {
		return nil, errors.New("malformed source file, NewArchive has exactly 4 arguments")
	}
	secondArgs, ok := args[1].(*ast.Ident)
	if !ok {
		return nil, errors.New("malformed source file, second argument of NewArchive must be an Ident")
	}
	if secondArgs.Name == "true" {
		return nil, errors.New("cannot extract source file in development mode")
	}
	firstArg, ok := args[0].(*ast.CompositeLit)
        if !ok {
                return nil, errors.New("malformed source file, first argument of NewArchive must be a CompositeLit")
        }
	archiveFiles := make(map[string]bog.File)
	for _, elt := range firstArg.Elts {
		keyValExpr, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			return nil, errors.New("malformed source file, key value invalid")
                }
		key, ok := keyValExpr.Key.(*ast.BasicLit)
                if !ok {
                        return nil, errors.New("malformed source file, key value invalid")
                }
		val, ok := keyValExpr.Value.(*ast.Ident)
                if !ok {
                        return nil, errors.New("malformed source file, key value invalid")
                }
		f, ok := files[val.Name]
		if !ok {
			return nil, errors.New("malformed source file, file not found: " + val.Name)
		}
		archiveFiles[strings.Trim(key.Value, "\"")] = f
	}
	var isFile bool
	thirdArgs, ok := args[2].(*ast.Ident)
        if !ok {
                return nil, errors.New("malformed source file, third argument of NewArchive must be an Ident")
        }
        if thirdArgs.Name == "true" {
		isFile = true
        } else {
		isFile = false
	}
	fourthArg, ok := args[3].(*ast.BasicLit)
        if !ok {
                return nil, errors.New("malformed source file, fourth argument of NewArchive must be a BasicLit")
        }
	return bog.NewArchive(archiveFiles, false, isFile, strings.Trim(fourthArg.Value, "\"")), nil
}

func parseStat(arg ast.Expr) (*bog.FileInfo, error) {
	var fileName string
	var fileMode os.FileMode
	secondArg, ok := arg.(*ast.UnaryExpr)
	if !ok {
		return nil, errors.New("malformed source file, second argument of NewBogFile must be a UnaryExpr")
	}
	x, ok := secondArg.X.(*ast.CompositeLit)
	if !ok {
		return nil, errors.New("malformed source file, X of second argument of NewBogFile must be a CompositeLit")
	}
	for _, elt := range x.Elts {
		keyValExpr, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			return nil, errors.New("malformed source file, key value invalid")
		}
		key, ok := keyValExpr.Key.(*ast.Ident)
		if !ok {
			return nil, errors.New("malformed source file, key value invalid")
		}
		val, ok := keyValExpr.Value.(*ast.BasicLit)
		if !ok {
			// must ignore here because of time.Unix call
			continue
		}
		if key.Name == "FileName" {
			fileName = strings.Trim(val.Value, "\"")
		}
		if key.Name == "FileMode" {
			val, err := strconv.ParseInt(val.Value, 0, 64)
			if err != nil {
				return nil, err
			}
			fileMode = os.FileMode(val)
		}
	}
	return &bog.FileInfo{
		FileName: fileName,
		FileMode: fileMode,
	}, nil
}
