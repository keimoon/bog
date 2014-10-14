package main

import (
	"regexp"
	"strings"
)

var pkgRegex = regexp.MustCompile("[^a-zA-Z0-9\\-\\_/]")

func makePackageName(name string) string {
	return strings.Trim(strings.Replace(pkgRegex.ReplaceAllString(strings.ToLower(name), ""), "/", "-", -1), "-")
}

var filenameRegex = regexp.MustCompile("[^a-zA-Z0-9]")

func makeFilename(name string) string {
	return filenameRegex.ReplaceAllString(strings.ToLower(name), "")
}

var underscoreRegex = regexp.MustCompile("\\_+")

func makeVariableName(name string) string {
	return "vvv" + makePublicVariableName(underscoreRegex.ReplaceAllString(filenameRegex.ReplaceAllString(strings.ToLower(name), "_"), "_"))
}

var slugRegex = regexp.MustCompile("(\\s|-|_)+")

func makePublicVariableName(name string) string {
	return strings.Replace(strings.Title(slugRegex.ReplaceAllString(name, " ")), " ", "", -1)
}
