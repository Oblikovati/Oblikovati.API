// SPDX-License-Identifier: Apache-2.0

package main

import (
	"go/ast"
	"go/types"
)

func featureKind(loader *sourceImporter, object *types.TypeName) string {
	if object.Pkg().Name() != "featureargs" {
		return ""
	}
	for _, file := range loader.files[object.Pkg().Path()] {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Name.Name != "Kind" || fn.Recv == nil || receiverName(fn) != object.Name() {
				continue
			}
			result := fn.Body.List[0].(*ast.ReturnStmt)
			return loader.info[object.Pkg().Path()].Types[result.Results[0]].Value.ExactString()
		}
	}
	return ""
}
