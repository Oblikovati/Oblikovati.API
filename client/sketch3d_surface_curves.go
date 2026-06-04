// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/Oblikovati/api/types"
	"github.com/Oblikovati/api/wire"
)

// AddSurfaceCurve is the general surface-derived curve constructor; prefer the typed
// helpers below.
func (s Sketch3D) AddSurfaceCurve(args wire.AddSketch3DSurfaceCurveArgs) (wire.AddSketch3DSurfaceCurveResult, error) {
	var r wire.AddSketch3DSurfaceCurveResult
	return r, s.c.call(wire.MethodSketch3DAddSurfaceCurve, args, &r)
}

// AddIntersectionCurve adds the intersection curve of two part faces (by reference key).
// grid windows the first face's parameter domain (needed for an unbounded planar face).
func (s Sketch3D) AddIntersectionCurve(index int, faceA, faceB string, grid wire.AddSketch3DSurfaceCurveArgs) (wire.AddSketch3DSurfaceCurveResult, error) {
	grid.SketchIndex = index
	grid.Kind = string(types.Sketch3DEntityIntersection)
	grid.FaceRefs = []string{faceA, faceB}
	return s.AddSurfaceCurve(grid)
}

// AddSilhouetteCurve adds the silhouette of a part face (by reference key) for a view
// direction [x,y,z].
func (s Sketch3D) AddSilhouetteCurve(index int, face string, viewDir []float64, grid wire.AddSketch3DSurfaceCurveArgs) (wire.AddSketch3DSurfaceCurveResult, error) {
	grid.SketchIndex = index
	grid.Kind = string(types.Sketch3DEntitySilhouette)
	grid.FaceRefs = []string{face}
	grid.ViewDir = viewDir
	return s.AddSurfaceCurve(grid)
}
