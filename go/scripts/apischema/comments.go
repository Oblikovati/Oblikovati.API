// SPDX-License-Identifier: Apache-2.0

package main

import (
	"go/ast"
	"go/token"
	"go/types"
)

func indexComments(comments map[token.Pos]string, files []*ast.File) {
	for _, file := range files {
		ast.Inspect(file, func(node ast.Node) bool {
			names, doc := nodeComments(node)
			for _, name := range names {
				comments[name.Pos()] = doc
			}
			return true
		})
	}
}

func nodeComments(node ast.Node) ([]*ast.Ident, string) {
	switch item := node.(type) {
	case *ast.Field:
		return item.Names, item.Doc.Text() + item.Comment.Text()
	case *ast.TypeSpec:
		return []*ast.Ident{item.Name}, item.Doc.Text() + item.Comment.Text()
	case *ast.ValueSpec:
		return item.Names, item.Doc.Text() + item.Comment.Text()
	default:
		return nil, ""
	}
}

func attachMemberComments(comments map[token.Pos]string, shape *typeShape, typ types.Type) {
	switch item := typ.(type) {
	case *types.Struct:
		offset := 0
		for index := 0; index < item.NumFields(); index++ {
			if !item.Field(index).Exported() {
				continue
			}
			shape.Fields[offset].Doc = comments[item.Field(index).Pos()]
			offset++
		}
	case *types.Interface:
		for index := range shape.Methods {
			shape.Methods[index].Doc = comments[item.ExplicitMethod(index).Pos()]
		}
	}
}
