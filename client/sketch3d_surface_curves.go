// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// AddSurfaceCurve is the general surface-derived curve constructor; prefer the typed
// helpers below.
//
// mcp:tool add_sketch3d_surface_curve
// mcp:summary Add a curve that lies on a part face/surface to a 3D sketch (project/intersection/silhouette).
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

// AddOnFaceCurve adds a curve drawn in a part face's parameter space (by reference key);
// uv is the flat [u0,v0,u1,v1,…] polyline mapped onto the face's surface.
func (s Sketch3D) AddOnFaceCurve(index int, face string, uv []float64) (wire.AddSketch3DSurfaceCurveResult, error) {
	return s.AddSurfaceCurve(wire.AddSketch3DSurfaceCurveArgs{
		SketchIndex: index,
		Kind:        string(types.Sketch3DEntityOnFace),
		FaceRefs:    []string{face},
		UV:          uv,
	})
}

// AddProjectToSurfaceCurve projects an in-sketch source curve (by entity id) onto a part
// face (by reference key).
func (s Sketch3D) AddProjectToSurfaceCurve(index int, sourceEntityID uint64, face string) (wire.AddSketch3DSurfaceCurveResult, error) {
	return s.AddSurfaceCurve(wire.AddSketch3DSurfaceCurveArgs{
		SketchIndex:    index,
		Kind:           string(types.Sketch3DEntityProjectToSurface),
		FaceRefs:       []string{face},
		SourceEntityID: sourceEntityID,
	})
}

// AddOffsetCurve offsets an in-sketch source curve (by entity id) by distance in the plane
// with the given normal [x,y,z] (offset direction = normal × tangent).
func (s Sketch3D) AddOffsetCurve(index int, sourceEntityID uint64, distance float64, normal []float64) (wire.AddSketch3DSurfaceCurveResult, error) {
	return s.AddSurfaceCurve(wire.AddSketch3DSurfaceCurveArgs{
		SketchIndex:    index,
		Kind:           string(types.Sketch3DEntityOffset),
		SourceEntityID: sourceEntityID,
		OffsetDistance: distance,
		Normal:         normal,
	})
}
