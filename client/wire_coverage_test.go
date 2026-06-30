// SPDX-License-Identifier: Apache-2.0

package client

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"strings"
	"testing"
)

// The wire↔client coverage guard: ADR-0018 requires a typed client method for
// every wire method, but documents.close / documents.closeAll shipped without
// one — an add-in could only reach them with raw JSON, which the contract
// forbids. This test parses the package sources so a wire method can never
// again be declared without typed client coverage.

// trackedUncoveredWireMethods are method constants knowingly without a typed
// client method, each requiring an open issue. Currently empty — keep it so the
// next gap must be tracked, not silently shipped.
var trackedUncoveredWireMethods = map[string]string{
	"MethodTaskPanelShow":  "task A5: TaskPanels client group deferred",
	"MethodTaskPanelClose": "task A5: TaskPanels client group deferred",
}

// TestEveryWireMethodHasTypedClientCoverage fails for each Method* constant in
// wire/methods.go that no non-test client source references.
func TestEveryWireMethodHasTypedClientCoverage(t *testing.T) {
	used := wireMethodsReferencedByClient(t)
	for _, name := range declaredWireMethodConstants(t) {
		issue, tracked := trackedUncoveredWireMethods[name]
		switch {
		case used[name] && tracked:
			t.Errorf("wire.%s is covered but still tracked-uncovered (%s) — remove the entry", name, issue)
		case !used[name] && !tracked:
			t.Errorf("wire.%s has no typed client method — add one to its operation group (ADR-0018)", name)
		}
	}
}

// declaredWireMethodConstants parses ../wire/methods.go (the test runs in the
// package directory) and returns every Method* constant name.
func declaredWireMethodConstants(t *testing.T) []string {
	t.Helper()
	file, err := parser.ParseFile(token.NewFileSet(), "../wire/methods.go", nil, 0)
	if err != nil {
		t.Fatalf("parse ../wire/methods.go: %v", err)
	}
	names := []string{}
	for _, decl := range file.Decls {
		names = append(names, methodConstNames(decl)...)
	}
	if len(names) == 0 {
		t.Fatal("no Method* constants found in ../wire/methods.go — wrong path or rename?")
	}
	return names
}

// methodConstNames returns the Method* constant names of one declaration.
func methodConstNames(decl ast.Decl) []string {
	gen, ok := decl.(*ast.GenDecl)
	if !ok || gen.Tok != token.CONST {
		return nil
	}
	names := []string{}
	for _, spec := range gen.Specs {
		for _, ident := range spec.(*ast.ValueSpec).Names {
			if strings.HasPrefix(ident.Name, "Method") {
				names = append(names, ident.Name)
			}
		}
	}
	return names
}

// wireMethodsReferencedByClient parses the non-test client sources and returns
// every wire.Method* selector they mention.
func wireMethodsReferencedByClient(t *testing.T) map[string]bool {
	t.Helper()
	notTest := func(fi fs.FileInfo) bool { return !strings.HasSuffix(fi.Name(), "_test.go") }
	pkgs, err := parser.ParseDir(token.NewFileSet(), ".", notTest, 0)
	if err != nil {
		t.Fatalf("parse client package: %v", err)
	}
	used := map[string]bool{}
	for _, pkg := range pkgs {
		ast.Inspect(pkg, func(n ast.Node) bool { return recordWireSelector(n, used) })
	}
	return used
}

// recordWireSelector records a wire.Method* selector expression into used.
func recordWireSelector(n ast.Node, used map[string]bool) bool {
	sel, ok := n.(*ast.SelectorExpr)
	if !ok {
		return true
	}
	pkg, ok := sel.X.(*ast.Ident)
	if ok && pkg.Name == "wire" && strings.HasPrefix(sel.Sel.Name, "Method") {
		used[sel.Sel.Name] = true
	}
	return true
}
