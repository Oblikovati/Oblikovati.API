// SPDX-License-Identifier: Apache-2.0

package client

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
	"testing"
)

// TestMCPWireAnnotations requires each wire method to declare how the bridge
// should handle it. A documented skip is for methods intentionally kept out of MCP.
func TestMCPWireAnnotations(t *testing.T) {
	annotated := make(map[string]bool)
	for _, file := range parseMCPFiles(t, ".") {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Body == nil || !hasMCPAnnotation(fn.Doc.Text()) {
				continue
			}
			if method := clientMCPMethod(fn.Body); method != "" {
				annotated[method] = true
			}
		}
	}
	for _, file := range parseMCPFiles(t, "../wire") {
		for _, spec := range file.Decls {
			checkWireAnnotations(t, spec, annotated)
		}
	}
}

func clientMCPMethod(body *ast.BlockStmt) string {
	var method string
	ast.Inspect(body, func(node ast.Node) bool {
		if call, ok := node.(*ast.CallExpr); ok {
			if name := mcpWireMethod(call); name != "" {
				method = name
			}
		}
		return true
	})
	return method
}

func parseMCPFiles(t *testing.T, dir string) []*ast.File {
	t.Helper()
	paths, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil || len(paths) == 0 {
		t.Fatalf("find Go files in %s: %d paths, error %v", dir, len(paths), err)
	}
	fset := token.NewFileSet()
	var files []*ast.File
	for _, path := range paths {
		if strings.HasSuffix(path, "_test.go") {
			continue
		}
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			t.Fatal(err)
		}
		files = append(files, file)
	}
	return files
}

func hasMCPAnnotation(doc string) bool {
	var tool, summary bool
	for line := range strings.Lines(doc) {
		key, value, _ := strings.Cut(strings.TrimSpace(line), " ")
		if strings.TrimSpace(value) == "" {
			continue
		}
		switch key {
		case "mcp:skip":
			return true
		case "mcp:tool":
			tool = true
		case "mcp:summary":
			summary = true
		}
	}
	return tool && summary
}

// Match the same dispatch forms as MCP Bridge's generator.
func mcpWireMethod(call *ast.CallExpr) string {
	var index int
	switch fn := call.Fun.(type) {
	case *ast.SelectorExpr:
		if fn.Sel.Name != "invoke" && fn.Sel.Name != "call" {
			return ""
		}
	case *ast.IndexExpr:
		name, ok := fn.X.(*ast.Ident)
		if !ok || name.Name != "call" {
			return ""
		}
		index = 1
	default:
		return ""
	}
	if len(call.Args) <= index+1 {
		return ""
	}
	method, ok := call.Args[index].(*ast.SelectorExpr)
	if !ok {
		return ""
	}
	pkg, ok := method.X.(*ast.Ident)
	if !ok || pkg.Name != "wire" {
		return ""
	}
	return method.Sel.Name
}

func checkWireAnnotations(t *testing.T, decl ast.Decl, annotated map[string]bool) {
	t.Helper()
	group, ok := decl.(*ast.GenDecl)
	if !ok || group.Tok != token.CONST {
		return
	}
	for _, spec := range group.Specs {
		for _, name := range spec.(*ast.ValueSpec).Names {
			if strings.HasPrefix(name.Name, "Method") && !annotated[name.Name] {
				t.Errorf("wire.%s: add mcp:tool and mcp:summary, or mcp:skip with a reason, to its Go client call", name.Name)
			}
		}
	}
}
