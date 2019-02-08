package smock

import (
	"os"
	"path/filepath"
	"strings"
)

// stripGopath takes the directory to a package and remove the gopath to get the
// canonical package name.
//
// taken from https://github.com/ernesto-jimenez/gogen
// Copyright (c) 2015 Ernesto Jiménez
func stripGopath(p string) string {
	for _, gopath := range gopaths() {
		base := strings.TrimSuffix(gopath, string(os.PathSeparator))
		pref := strings.Join([]string{base, "src"}, string(os.PathSeparator)) + string(os.PathSeparator)
		p = strings.TrimPrefix(p, pref)
	}
	return p
}

func gopaths() []string {
	return strings.Split(os.Getenv("GOPATH"), string(filepath.ListSeparator))
}

func afterLastRune(str string, rune2 rune) string {
	idx := strings.LastIndex(str, string(rune2))
	if idx < 0 {
		return str
	}

	return str[idx+1:]
}

func afterLastSlash(str string) string {
	return afterLastRune(str, '/')
}

func afterLastDot(str string) string {
	return afterLastRune(str, '.')
}
