// SPDX-License-Identifier: Apache-2.0

// Command apischema exports the checked Go declarations used by the Rust binding generator.
package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type sourceImporter struct {
	root     string
	fset     *token.FileSet
	packages map[string]*types.Package
	files    map[string][]*ast.File
	info     map[string]*types.Info
	standard types.Importer
	comments map[token.Pos]string
}

func (loader *sourceImporter) Import(path string) (*types.Package, error) {
	if !strings.HasPrefix(path, "oblikovati.org/api") {
		return loader.standard.Import(path)
	}
	if pkg := loader.packages[path]; pkg != nil {
		return pkg, nil
	}
	return loader.load(path)
}

func (loader *sourceImporter) load(path string) (*types.Package, error) {
	dir := filepath.Join(loader.root, strings.TrimPrefix(path, "oblikovati.org/api"))
	parsed, err := parser.ParseDir(loader.fset, dir, sourceFile, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	var files []*ast.File
	for _, pkg := range parsed {
		for _, file := range pkg.Files {
			files = append(files, file)
		}
	}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue), Defs: make(map[*ast.Ident]types.Object)}
	config := types.Config{Importer: loader}
	pkg, err := config.Check(path, loader.fset, files, info)
	if err != nil {
		return nil, err
	}
	loader.packages[path], loader.files[path], loader.info[path] = pkg, files, info
	indexComments(loader.comments, files)
	return pkg, nil
}

func sourceFile(info os.FileInfo) bool {
	return strings.HasSuffix(info.Name(), ".go") && !strings.HasSuffix(info.Name(), "_test.go")
}

func main() {
	loader := newSourceImporter(".")
	if err := exportSchema(loader, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newSourceImporter(root string) *sourceImporter {
	return &sourceImporter{root: root, fset: token.NewFileSet(), packages: make(map[string]*types.Package), files: make(map[string][]*ast.File), info: make(map[string]*types.Info), standard: importer.Default(), comments: make(map[token.Pos]string)}
}

func exportSchema(loader *sourceImporter, output io.Writer) error {
	for _, name := range []string{"types", "contract", "wire", "wire/featureargs", "client"} {
		if _, err := loader.Import("oblikovati.org/api/" + name); err != nil {
			return err
		}
	}
	schema := buildSchema(loader)
	encoder := json.NewEncoder(output)
	encoder.SetIndent("", "  ")
	return encoder.Encode(schema)
}
