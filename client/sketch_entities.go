// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// AddEntity is the general entity constructor — the escape hatch covering every kind and
// variant; prefer the typed helpers below for the common constructors.
//
// mcp:tool add_sketch_entity
// mcp:summary Add a 2D primitive to a sketch: {sketchIndex, kind, points:[[x,y],…]} with kind one of line|circle|arc|rectangle|slot|polygon|polyline|ellipse|spline|point. Optional variant (e.g. "threePoint","centerPoint"), radius (unit expr), ccw, construction. For kind "polyline" pass points:[[x,y],…] and optional closed:true (an arbitrary outline; closed ⇒ one closed profile). Coordinates are cm in sketch space.
func (s Sketch) AddEntity(args wire.AddSketchEntityArgs) (wire.AddSketchEntityResult, error) {
	var r wire.AddSketchEntityResult
	return r, s.c.call(wire.MethodSketchAddEntity, args, &r)
}

// AddLine adds a line between two points (each [x,y] in cm).
func (s Sketch) AddLine(index int, a, b []float64, construction bool) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityLine),
		Points: [][]float64{a, b}, Construction: construction,
	})
}

// AddPoint adds a standalone sketch point at [x,y] (cm).
func (s Sketch) AddPoint(index int, p []float64) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityPoint), Points: [][]float64{p},
	})
}

// AddCircleByCenterRadius adds a circle from a center [x,y] (cm) and a unit-bearing
// radius ("10 mm").
func (s Sketch) AddCircleByCenterRadius(index int, center []float64, radius string, construction bool) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityCircle), Variant: "centerRadius",
		Points: [][]float64{center}, Radius: radius, Construction: construction,
	})
}

// AddCircleByThreePoints adds the circle through three points (the circumcircle).
func (s Sketch) AddCircleByThreePoints(index int, a, b, c []float64, construction bool) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityCircle), Variant: "threePoint",
		Points: [][]float64{a, b, c}, Construction: construction,
	})
}

// AddArcByCenterStartEnd adds an arc from a center, start, and end point; ccw orients it.
func (s Sketch) AddArcByCenterStartEnd(index int, center, start, end []float64, ccw, construction bool) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityArc), Variant: "centerStartEnd",
		Points: [][]float64{center, start, end}, CCW: ccw, Construction: construction,
	})
}

// AddArcByThreePoints adds the arc through three points (start, mid, end).
func (s Sketch) AddArcByThreePoints(index int, a, b, c []float64, construction bool) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityArc), Variant: "threePoint",
		Points: [][]float64{a, b, c}, Construction: construction,
	})
}

// AddLineExpr adds a line whose endpoints are parameter expressions (#189): a and b are
// each ["x-expr","y-expr"] evaluated through the document's parameter engine ("bore_r",
// "slot_w/2 + 1 mm"), so the line is parametric at construction. Use this over [Sketch.AddLine]
// when a generated profile's vertices must track parameters (a stator slot, a linkage).
func (s Sketch) AddLineExpr(index int, a, b []string, construction bool) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityLine),
		PointExprs: [][]string{a, b}, Construction: construction,
	})
}

// AddPointExpr adds a standalone sketch point whose coordinates are parameter expressions
// (p is ["x-expr","y-expr"]); see [Sketch.AddLineExpr] (#189).
func (s Sketch) AddPointExpr(index int, p []string) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityPoint), PointExprs: [][]string{p},
	})
}

// AddArcByCenterStartEndExpr adds a center-start-end arc whose three points are parameter
// expressions (each ["x-expr","y-expr"]); ccw orients it. The expression form lets a generated
// arc's center and endpoints track parameters (a magnet arc on a rotor) (#189).
func (s Sketch) AddArcByCenterStartEndExpr(index int, center, start, end []string, ccw, construction bool) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityArc), Variant: "centerStartEnd",
		PointExprs: [][]string{center, start, end}, CCW: ccw, Construction: construction,
	})
}

// AddEllipse adds a full ellipse from a center [x,y] (cm), a major-axis direction [x,y],
// and unit-bearing major/minor radii ("20 mm").
func (s Sketch) AddEllipse(index int, center, axis []float64, majorR, minorR string, construction bool) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityEllipse),
		Points: [][]float64{center}, Axis: axis, MajorRadius: majorR, MinorRadius: minorR,
		Construction: construction,
	})
}

// EllipticalArc is the shape of a 2D elliptical arc for [Sketch.AddEllipticalArc]:
// a center [x,y] (cm) and major-axis direction [x,y], unit-bearing major/minor radii
// ("20 mm"), and unit-bearing start/end angles ("0 deg", "90 deg") in the ellipse's
// major/minor frame. Grouped into a struct so the call stays under the parameter limit
// and reads by field at the call site.
type EllipticalArc struct {
	Center, Axis             []float64
	MajorRadius, MinorRadius string
	StartAngle, EndAngle     string
	Construction             bool
}

// AddEllipticalArc adds an elliptical arc described by spec.
func (s Sketch) AddEllipticalArc(index int, spec EllipticalArc) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityEllipticalArc),
		Points: [][]float64{spec.Center}, Axis: spec.Axis,
		MajorRadius: spec.MajorRadius, MinorRadius: spec.MinorRadius,
		StartAngle: spec.StartAngle, EndAngle: spec.EndAngle, Construction: spec.Construction,
	})
}

// AddSpline adds a spline through fit points (default) or as a control-point spline when
// variant is "controlPoint". Points are [x,y] in cm; closed makes a closed loop.
func (s Sketch) AddSpline(index int, variant string, points [][]float64, closed, construction bool) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntitySpline), Variant: variant,
		Points: points, Closed: closed, Construction: construction,
	})
}

// AddEquationCurve adds a parametric curve x(t)/y(t) over t ∈ [t0, t1] (t unitless).
func (s Sketch) AddEquationCurve(index int, xExpr, yExpr string, t0, t1 float64) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityEquationCurve),
		XExpr: xExpr, YExpr: yExpr, T0: t0, T1: t1,
	})
}

// AddFixedSpline adds an immutable spline through the given fixed points ([x,y] cm each).
func (s Sketch) AddFixedSpline(index int, points [][]float64) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityFixedSpline), Points: points,
	})
}

// AddOffsetSpline adds the offset of a parent spline (by id) at a unit-bearing distance.
func (s Sketch) AddOffsetSpline(index int, parentSpline uint64, distance string) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityOffsetSpline),
		EntityRefs: []uint64{parentSpline}, Radius: distance,
	})
}

// AddRectangle adds an axis-aligned rectangle from two opposite corners (each [x,y] cm).
func (s Sketch) AddRectangle(index int, corner, opposite []float64, construction bool) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityRectangle),
		Points: [][]float64{corner, opposite}, Construction: construction,
	})
}

// AddPolygon adds a regular polygon with the given side count, centered at center with a
// vertex (inscribed) or edge-midpoint (when variant is "circumscribed") at through.
func (s Sketch) AddPolygon(index int, center, through []float64, sides int, variant string, construction bool) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityPolygon), Variant: variant,
		Points: [][]float64{center, through}, Sides: sides, Construction: construction,
	})
}

// AddSlot adds a center-to-center straight slot of the given unit-bearing width.
func (s Sketch) AddSlot(index int, c0, c1 []float64, width string, construction bool) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntitySlot),
		Points: [][]float64{c0, c1}, Width: width, Construction: construction,
	})
}

// AddFillet rounds the corner between two existing lines (by entity id) with a tangent
// arc of the given unit-bearing radius, trimming both lines.
func (s Sketch) AddFillet(index int, line1, line2 uint64, radius string) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityFillet),
		EntityRefs: []uint64{line1, line2}, Radius: radius,
	})
}

// AddChamfer bevels the corner between two existing lines with distances d1 and d2 (each
// unit-bearing); pass an empty d2 for an equal-distance chamfer.
func (s Sketch) AddChamfer(index int, line1, line2 uint64, d1, d2 string) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityChamfer),
		EntityRefs: []uint64{line1, line2}, Radius: d1, Distance2: d2,
	})
}
