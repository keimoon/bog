package main

import (
	"fmt"
	"os"
	"path/filepath"
	"flag"
)

var Options = &struct {
	PackageName string
	Dev         bool
	isCwd       bool
}{
	isCwd: true,
}
var Args = []string{}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "%s [flags] (archive|a) folder\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "%s [flags] (extract|e) file.go\n", os.Args[0])
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	cwd = makePackageName(filepath.Base(cwd))
	flag.Usage = Usage
	flag.StringVar(&Options.PackageName, "p", cwd, "Change package name")
	
	flag.BoolVar(&Options.Dev, "d", false, "Enable development mode")
	flag.Parse()
	Args = flag.Args()
	if len(Args) == 0 {
		Usage()
		os.Exit(2)
	}
	Options.PackageName = makePackageName(Options.PackageName)
	if cwd != Options.PackageName {
		Options.isCwd = false
	}
	switch Args[0] {
	case "archive", "a":
		os.Exit(Archive())
	case "extract", "e":
		os.Exit(Extract())
	default:
		Usage()
	}
}
