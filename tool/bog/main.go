package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"regexp"
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

var pkgRegex = regexp.MustCompile("[^a-zA-Z0-9\\-\\_]")
func makePackageName(name string) string {
	return pkgRegex.ReplaceAllString(strings.ToLower(name), "")
}

var filenameRegex = regexp.MustCompile("[^a-zA-Z0-9]")
func makeFilename(name string) string {
        return filenameRegex.ReplaceAllString(strings.ToLower(name), "")
}

func makeVariableName(name string) string {
	return "v_" + filenameRegex.ReplaceAllString(strings.ToLower(name), "_")
}

var slugRegex = regexp.MustCompile("(\\s|-|_)+")
func makePublicVariableName(name string) string {
	return strings.Replace(strings.Title(slugRegex.ReplaceAllString(name, " ")), " ", "", -1)
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	cwd = makePackageName(filepath.Base(cwd))
	flag.Usage = Usage
	flag.StringVar(&Options.PackageName, "p", cwd, "Package name")
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
		Extract()
	default:
		Usage()
	}
}
