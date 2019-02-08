package smock

import (
	"bytes"
	"fmt"
	"go/importer"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Generator struct {
	Type types.Type
}

func loadType(pkg, typeName string) (types.Type, error) {
	var err error
	// Resolve package path
	if pkg == "" || pkg[0] == '.' {
		cleanPath := filepath.Clean(pkg)
		pkg, err = filepath.Abs(cleanPath)
		if err != nil {
			return nil, err
		}

		pkg = stripGopath(pkg)
	}

	targetPackage, err := importer.For("source", nil).Import(pkg)
	if err != nil {
		return nil, err
	}

	obj := targetPackage.Scope().Lookup(typeName)
	if obj == nil {
		return nil, fmt.Errorf("struct %s missing", typeName)
	}

	var st *types.Named
	switch obj.Type().Underlying().(type) {
	case *types.Interface:
		st = obj.Type().(*types.Named)
	default:
		return nil, fmt.Errorf("%s should be an struct, was %s", typeName, obj.Type().Underlying())
	}

	return st, nil
}

func (g *Generator) GenerateMockOnMethods(buf *bytes.Buffer) error {
	fs := NewFuncSet(g.Type.(*types.Named))
	return onFuncTemplate.Execute(buf, fs)
}

func (g *Generator) WriteMockOnMethodsToFile(fileName string) error {
	var buf bytes.Buffer
	if err := g.GenerateMockOnMethods(&buf); err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, buf.Bytes(), os.ModePerm)
}

func NewGeneratorForType(pkg, typeName string) (*Generator, error) {
	st, err := loadType(pkg, typeName)
	if err != nil {
		return nil, err
	}

	return &Generator{Type: st}, nil
}
