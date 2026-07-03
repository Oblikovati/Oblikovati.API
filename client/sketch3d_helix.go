// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// AddVariableHelix adds a helical curve whose pitch/diameter vary over a row
// table of stations (M06-F09, Oblikovati/Oblikovati#624). origin/axis/radius
// have the [Sketch3D.AddHelix] meaning; rows are the shape stations and args
// carries the mode and end conditions.
func (s Sketch3D) AddVariableHelix(index int, origin, axis []float64, radius string, rows []wire.HelixShapeRow, args wire.AddSketch3DEntityArgs) (wire.AddSketch3DEntityResult, error) {
	args.SketchIndex = index
	args.Kind = string(types.Sketch3DEntityHelical)
	args.Points = [][]float64{origin}
	args.Axis = axis
	args.Radius = radius
	args.Rows = rows
	return s.AddEntity(args)
}

// EditHelixDefinition redefines an existing helical curve in place — constant
// or variable shape plus end conditions — and regenerates it (M06-F09).
//
// mcp:tool sketch3d_edit_helix
// mcp:summary Redefines an existing helical curve in place — constant or variable shape plus end conditions — and regenerates it (M06-F09).
func (s Sketch3D) EditHelixDefinition(args wire.EditHelixArgs) (wire.HelixDefinitionView, error) {
	return call[wire.HelixDefinitionView](s.c, wire.MethodSketch3DEditHelix, args)
}
