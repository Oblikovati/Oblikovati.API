// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"testing"
)

type rejectingWriter struct{}

func (rejectingWriter) Write([]byte) (int, error) { return 0, errors.New("write refused") }

func TestSchemaMatchesCompilerCheckedAPI(t *testing.T) {
	loader := newSourceImporter("../..")
	var output bytes.Buffer
	if err := exportSchema(loader, &output); err != nil {
		t.Fatal(err)
	}
	var schema apiSchema
	if err := json.Unmarshal(output.Bytes(), &schema); err != nil {
		t.Fatal(err)
	}
	assertSchemaSurface(t, schema)
	if err := exportSchema(loader, rejectingWriter{}); err == nil {
		t.Fatal("expected writer error")
	}
	if _, err := loader.Import("oblikovati.org/api/missing"); err == nil {
		t.Fatal("expected missing package error")
	}
}

func assertSchemaSurface(t *testing.T, schema apiSchema) {
	t.Helper()
	declarations := make(map[string]declaration)
	for _, decl := range schema.Declarations {
		declarations[decl.Package+"."+decl.Name] = decl
	}
	alias := declarations["contract.FaceShells"]
	if !alias.Alias || alias.Shape.Name != "Enumerable" {
		t.Fatalf("alias target lost: %+v", alias)
	}
	extrude := declarations["featureargs.Extrude"]
	if extrude.Kind != `"extrude"` || len(extrude.Shape.Fields) == 0 {
		t.Fatalf("invalid feature: %+v", extrude)
	}
	assertCallShape(t, schema.Calls, "Application", "ApiVersion", "empty", "ApplicationApiVersionResult")
	assertCallShape(t, schema.Calls, "Documents", "Create", "named", "DocumentInfo")
}

func assertCallShape(t *testing.T, calls []rpcShape, group, name, request, result string) {
	t.Helper()
	for _, call := range calls {
		if call.Group != group || call.Name != name {
			continue
		}
		if call.Request.Kind != request || call.Response.Name != result {
			t.Fatalf("invalid call %+v", call)
		}
		return
	}
	t.Fatalf("missing typed call %s.%s", group, name)
}

func TestSupportedCompositeShapes(t *testing.T) {
	integer := types.Typ[types.Int64]
	shapes := []struct {
		typ  types.Type
		kind string
	}{
		{integer, "basic"}, {types.NewPointer(integer), "pointer"},
		{types.NewSlice(integer), "slice"}, {types.NewArray(integer, 3), "array"},
		{types.NewMap(types.Typ[types.String], integer), "map"},
		{types.NewStruct(nil, nil), "struct"}, {types.NewInterfaceType(nil, nil), "interface"},
	}
	for _, test := range shapes {
		if got := describeType(test.typ); got.Kind != test.kind {
			t.Errorf("%s: got %+v", test.typ, got)
		}
	}
}

func TestUnsupportedTypeNamesItsShape(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("expected unsupported channel to be refused")
		}
	}()
	describeType(types.NewChan(types.SendRecv, types.Typ[types.Int]))
}

func TestExpressionAndReceiverShapes(t *testing.T) {
	expression := &ast.Ident{Name: "nil"}
	if got := expressionShape(expression, &types.Info{}); got.Kind != "empty" {
		t.Fatal(got)
	}
	fn := &ast.FuncDecl{Recv: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Group"}}}}}
	if got := receiverName(fn); got != "Group" {
		t.Fatal(got)
	}
	fn.Recv.List[0].Type = &ast.StarExpr{X: &ast.Ident{Name: "Client"}}
	if got := receiverName(fn); got != "Client" {
		t.Fatal(got)
	}
}

func TestSourceFilterAndCompilerError(t *testing.T) {
	for name, expected := range map[string]bool{"main.go": true, "main_test.go": false, "README.md": false} {
		if got := sourceFile(namedFileInfo{name: name}); got != expected {
			t.Errorf("%s: %v", name, got)
		}
	}
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "broken.go"), []byte("package broken\nvar X = Missing"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := newSourceImporter(root).Import("oblikovati.org/api"); err == nil {
		t.Fatal("expected compiler error")
	}
	if err := exportSchema(newSourceImporter(root), &bytes.Buffer{}); err == nil {
		t.Fatal("expected missing package error")
	}
}

type namedFileInfo struct {
	os.FileInfo
	name string
}

func (entry namedFileInfo) Name() string { return entry.name }

func TestFunctionWithoutTransportCall(t *testing.T) {
	file, err := parser.ParseFile(token.NewFileSet(), "fake.go", "package client; func (g Group) Local() { println(1) }", parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	fn := file.Decls[0].(*ast.FuncDecl)
	if calls := functionCalls(fn, &types.Info{}); len(calls) != 0 {
		t.Fatal(calls)
	}
}
