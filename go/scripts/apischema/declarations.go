// SPDX-License-Identifier: Apache-2.0

package main

import (
	"go/ast"
	"go/types"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

type typeShape struct {
	Kind       string
	Name       string        `json:",omitempty"`
	Package    string        `json:",omitempty"`
	Underlying string        `json:",omitempty"`
	Length     int64         `json:",omitempty"`
	Element    *typeShape    `json:",omitempty"`
	Key        *typeShape    `json:",omitempty"`
	Fields     []fieldShape  `json:",omitempty"`
	Methods    []methodShape `json:",omitempty"`
	Embedded   []typeShape   `json:",omitempty"`
	Parameters []typeShape   `json:",omitempty"`
}

type fieldShape struct {
	Name     string
	Type     typeShape
	Tag      string
	Embedded bool
	Doc      string `json:",omitempty"`
}
type methodShape struct {
	Name       string
	Parameters []fieldShape
	Results    []fieldShape
	Doc        string `json:",omitempty"`
}
type declaration struct {
	Name     string
	Package  string
	File     string
	Doc      string
	Shape    typeShape
	Constant bool
	Alias    bool
	Value    string `json:",omitempty"`
	Kind     string `json:",omitempty"`
}
type apiSchema struct {
	Declarations []declaration
	Calls        []rpcShape
}

func buildSchema(loader *sourceImporter) apiSchema {
	var schema apiSchema
	for path, pkg := range loader.packages {
		if strings.HasSuffix(path, "/client") {
			continue
		}
		for _, name := range pkg.Scope().Names() {
			object := pkg.Scope().Lookup(name)
			if !object.Exported() {
				continue
			}
			if decl, ok := describeDeclaration(loader, object); ok {
				schema.Declarations = append(schema.Declarations, decl)
			}
		}
	}
	sort.Slice(schema.Declarations, func(i, j int) bool {
		return schema.Declarations[i].Package+schema.Declarations[i].Name < schema.Declarations[j].Package+schema.Declarations[j].Name
	})
	schema.Calls = describeCalls(loader)
	return schema
}

func describeDeclaration(loader *sourceImporter, object types.Object) (declaration, bool) {
	decl := declaration{Name: object.Name(), Package: object.Pkg().Name(), File: filepath.Base(loader.fset.Position(object.Pos()).Filename)}
	decl.Doc = declarationDoc(loader, object)
	switch named := object.(type) {
	case *types.TypeName:
		decl.Shape = describeType(named.Type().Underlying())
		decl.Alias = named.IsAlias()
		if decl.Alias {
			decl.Shape = describeType(types.Unalias(named.Type()))
		}
		if generic, ok := named.Type().(*types.Named); ok {
			decl.Shape.Parameters = typeParameters(generic.TypeParams())
		}
		decl.Kind = featureKind(loader, named)
		if !decl.Alias {
			attachMemberComments(loader.comments, &decl.Shape, named.Type().Underlying())
		}
	case *types.Const:
		decl.Shape, decl.Value = describeType(named.Type()), named.Val().ExactString()
		decl.Constant = true
	default:
		return decl, false
	}
	return decl, true
}

func declarationDoc(loader *sourceImporter, object types.Object) string {
	if doc := loader.comments[object.Pos()]; doc != "" {
		return doc
	}
	for _, file := range loader.files[object.Pkg().Path()] {
		for _, decl := range file.Decls {
			group, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			for _, spec := range group.Specs {
				if spec.Pos() <= object.Pos() && object.Pos() <= spec.End() {
					return group.Doc.Text()
				}
			}
		}
	}
	return ""
}

func typeParameters(parameters *types.TypeParamList) []typeShape {
	var result []typeShape
	for index := 0; index < parameters.Len(); index++ {
		result = append(result, describeType(parameters.At(index)))
	}
	return result
}

func describeType(typ types.Type) typeShape {
	switch item := typ.(type) {
	case *types.Basic:
		return typeShape{Kind: "basic", Name: item.Name()}
	case *types.Named:
		return describeNamed(item)
	case *types.Alias:
		return describeType(types.Unalias(item))
	case *types.Pointer:
		child := describeType(item.Elem())
		return typeShape{Kind: "pointer", Element: &child}
	case *types.Slice:
		child := describeType(item.Elem())
		return typeShape{Kind: "slice", Element: &child}
	case *types.Array:
		child := describeType(item.Elem())
		return typeShape{Kind: "array", Element: &child, Length: item.Len()}
	case *types.Map:
		key, child := describeType(item.Key()), describeType(item.Elem())
		return typeShape{Kind: "map", Element: &child, Key: &key}
	case *types.Struct:
		return describeStruct(item)
	case *types.Interface:
		return describeInterface(item)
	case *types.TypeParam:
		return typeShape{Kind: "parameter", Name: item.Obj().Name()}
	default:
		panic("unsupported Go declaration type " + typ.String() + "; expected value or interface")
	}
}

func describeNamed(named *types.Named) typeShape {
	shape := typeShape{Kind: "named", Name: named.Obj().Name()}
	switch underlying := named.Underlying().(type) {
	case *types.Basic:
		shape.Underlying = underlying.Name()
	case *types.Struct:
		shape.Underlying = "struct"
	case *types.Slice:
		shape.Underlying = "slice"
	}
	if named.Obj().Pkg() != nil {
		shape.Package = named.Obj().Pkg().Name()
	}
	for index := 0; index < named.TypeArgs().Len(); index++ {
		shape.Parameters = append(shape.Parameters, describeType(named.TypeArgs().At(index)))
	}
	return shape
}

func describeStruct(value *types.Struct) typeShape {
	shape := typeShape{Kind: "struct"}
	for index := 0; index < value.NumFields(); index++ {
		field := value.Field(index)
		if !field.Exported() {
			continue
		}
		shape.Fields = append(shape.Fields, fieldShape{Name: field.Name(), Type: describeType(field.Type()), Tag: reflect.StructTag(value.Tag(index)).Get("json"), Embedded: field.Embedded()})
	}
	return shape
}

func describeInterface(value *types.Interface) typeShape {
	shape := typeShape{Kind: "interface"}
	for index := 0; index < value.NumExplicitMethods(); index++ {
		method := value.ExplicitMethod(index)
		signature := method.Type().(*types.Signature)
		shape.Methods = append(shape.Methods, methodShape{Name: method.Name(), Parameters: describeTuple(signature.Params()), Results: describeTuple(signature.Results())})
	}
	for index := 0; index < value.NumEmbeddeds(); index++ {
		shape.Embedded = append(shape.Embedded, describeType(value.EmbeddedType(index)))
	}
	return shape
}

func describeTuple(tuple *types.Tuple) []fieldShape {
	var result []fieldShape
	for index := 0; index < tuple.Len(); index++ {
		item := tuple.At(index)
		result = append(result, fieldShape{Name: item.Name(), Type: describeType(item.Type())})
	}
	return result
}
