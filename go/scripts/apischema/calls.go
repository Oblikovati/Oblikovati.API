// SPDX-License-Identifier: Apache-2.0

package main

import (
	"go/ast"
	"go/types"
	"sort"
	"strings"
)

type rpcShape struct {
	Group    string
	Name     string
	Method   string
	Request  typeShape
	Response typeShape
	Doc      string
}

func describeCalls(loader *sourceImporter) []rpcShape {
	const path = "oblikovati.org/api/client"
	var calls []rpcShape
	for _, file := range loader.files[path] {
		for _, decl := range file.Decls {
			function, ok := decl.(*ast.FuncDecl)
			if !ok || function.Recv == nil {
				continue
			}
			calls = append(calls, functionCalls(function, loader.info[path])...)
		}
	}
	sort.Slice(calls, func(i, j int) bool { return calls[i].Group+calls[i].Name < calls[j].Group+calls[j].Name })
	return calls
}

func functionCalls(function *ast.FuncDecl, info *types.Info) []rpcShape {
	var calls []rpcShape
	ast.Inspect(function.Body, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}
		for index, arg := range call.Args {
			selector, ok := arg.(*ast.SelectorExpr)
			if !ok || !strings.HasPrefix(selector.Sel.Name, "Method") {
				continue
			}
			if index+1 >= len(call.Args) {
				continue
			}
			rpc := rpcShape{Group: receiverName(function), Name: function.Name.Name, Method: selector.Sel.Name, Doc: function.Doc.Text()}
			rpc.Request = expressionShape(call.Args[index+1], info)
			rpc.Response = callResponse(call, info)
			calls = append(calls, rpc)
		}
		return true
	})
	return calls
}

func receiverName(function *ast.FuncDecl) string {
	typ := function.Recv.List[0].Type
	if pointer, ok := typ.(*ast.StarExpr); ok {
		typ = pointer.X
	}
	return typ.(*ast.Ident).Name
}

func expressionShape(expr ast.Expr, info *types.Info) typeShape {
	typ := info.TypeOf(expr)
	if typ == nil {
		return typeShape{Kind: "empty"}
	}
	if basic, ok := typ.(*types.Basic); ok && basic.Kind() == types.UntypedNil {
		return typeShape{Kind: "empty"}
	}
	return describeType(typ)
}

func callResponse(call *ast.CallExpr, info *types.Info) typeShape {
	if tuple, ok := info.TypeOf(call).(*types.Tuple); ok {
		return describeType(tuple.At(0).Type())
	}
	if len(call.Args) >= 3 {
		last := info.TypeOf(call.Args[len(call.Args)-1])
		if pointer, ok := last.(*types.Pointer); ok {
			return describeType(pointer.Elem())
		}
	}
	return typeShape{Kind: "empty"}
}
