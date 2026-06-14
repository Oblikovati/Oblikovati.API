// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// AddEntity is the general 3D entity constructor — the escape hatch covering every kind;
// prefer the typed helpers below for the common constructors.
//
// mcp:tool add_sketch3d_entity
// mcp:summary Add a 3D curve to a sketch: {sketchIndex, kind, points:[[x,y,z],…]} with kind line|arc|circle|spline|point|helix and friends. Coordinates are cm.
func (s Sketch3D) AddEntity(args wire.AddSketch3DEntityArgs) (wire.AddSketch3DEntityResult, error) {
	var r wire.AddSketch3DEntityResult
	return r, s.c.call(wire.MethodSketch3DAddEntity, args, &r)
}

// AddPoint adds a standalone 3D sketch point at [x,y,z] (cm).
func (s Sketch3D) AddPoint(index int, p []float64) (wire.AddSketch3DEntityResult, error) {
	return s.AddEntity(wire.AddSketch3DEntityArgs{
		SketchIndex: index, Kind: string(types.Sketch3DEntityPoint), Points: [][]float64{p},
	})
}

// AddLine adds a straight 3D segment between two points (each [x,y,z] in cm).
func (s Sketch3D) AddLine(index int, a, b []float64, construction bool) (wire.AddSketch3DEntityResult, error) {
	return s.AddEntity(wire.AddSketch3DEntityArgs{
		SketchIndex: index, Kind: string(types.Sketch3DEntityLine),
		Points: [][]float64{a, b}, Construction: construction,
	})
}

// AddCircle adds a circle from a center [x,y,z] (cm), a plane-normal axis [x,y,z] (empty
// ⇒ +Z), and a unit-bearing radius ("10 mm").
func (s Sketch3D) AddCircle(index int, center, axis []float64, radius string, construction bool) (wire.AddSketch3DEntityResult, error) {
	return s.AddEntity(wire.AddSketch3DEntityArgs{
		SketchIndex: index, Kind: string(types.Sketch3DEntityCircle),
		Points: [][]float64{center}, Axis: axis, Radius: radius, Construction: construction,
	})
}

// AddArc adds a circular arc from a center, start and end point (each [x,y,z] in cm);
// ccw orients the sweep.
func (s Sketch3D) AddArc(index int, center, start, end []float64, ccw, construction bool) (wire.AddSketch3DEntityResult, error) {
	return s.AddEntity(wire.AddSketch3DEntityArgs{
		SketchIndex: index, Kind: string(types.Sketch3DEntityArc),
		Points: [][]float64{center, start, end}, CCW: ccw, Construction: construction,
	})
}

// AddEllipse adds a full ellipse from a center [x,y,z] (cm), a plane normal axis [x,y,z]
// (empty ⇒ +Z), an in-plane major-axis direction (empty ⇒ +X), and unit-bearing major/
// minor radii.
func (s Sketch3D) AddEllipse(index int, center, axis, majorAxis []float64, majorR, minorR string, construction bool) (wire.AddSketch3DEntityResult, error) {
	return s.AddEntity(wire.AddSketch3DEntityArgs{
		SketchIndex: index, Kind: string(types.Sketch3DEntityEllipse),
		Points: [][]float64{center}, Axis: axis, MajorAxis: majorAxis,
		MajorRadius: majorR, MinorRadius: minorR, Construction: construction,
	})
}

// EllipticalArc3D is the shape of a 3D elliptical arc for [Sketch3D.AddEllipticalArc]:
// a center [x,y,z] (cm), plane normal Axis [x,y,z] (empty ⇒ +Z), in-plane MajorAxis
// direction (empty ⇒ +X), unit-bearing major/minor radii, and unit-bearing start/sweep
// angles spanning StartAngle..StartAngle+SweepAngle. Grouped into a struct so the call
// stays under the parameter limit and reads by field at the call site.
type EllipticalArc3D struct {
	Center, Axis, MajorAxis  []float64
	MajorRadius, MinorRadius string
	StartAngle, SweepAngle   string
	Construction             bool
}

// AddEllipticalArc adds a bounded ellipse described by spec.
func (s Sketch3D) AddEllipticalArc(index int, spec EllipticalArc3D) (wire.AddSketch3DEntityResult, error) {
	return s.AddEntity(wire.AddSketch3DEntityArgs{
		SketchIndex: index, Kind: string(types.Sketch3DEntityEllipticalArc),
		Points: [][]float64{spec.Center}, Axis: spec.Axis, MajorAxis: spec.MajorAxis,
		MajorRadius: spec.MajorRadius, MinorRadius: spec.MinorRadius,
		StartAngle: spec.StartAngle, SweepAngle: spec.SweepAngle,
		Construction: spec.Construction,
	})
}

// AddSpline adds an interpolation (fit=true) or control-point (fit=false) spline through
// the given points (each [x,y,z] in cm); closed marks a loop.
func (s Sketch3D) AddSpline(index int, points [][]float64, closed, fit bool) (wire.AddSketch3DEntityResult, error) {
	kind := types.Sketch3DEntitySpline
	if !fit {
		kind = types.Sketch3DEntityControlPointSpline
	}
	return s.AddEntity(wire.AddSketch3DEntityArgs{
		SketchIndex: index, Kind: string(kind), Points: points, Closed: closed,
	})
}

// AddFixedSpline adds an immutable spline through the given points.
func (s Sketch3D) AddFixedSpline(index int, points [][]float64, closed bool) (wire.AddSketch3DEntityResult, error) {
	return s.AddEntity(wire.AddSketch3DEntityArgs{
		SketchIndex: index, Kind: string(types.Sketch3DEntityFixedSpline), Points: points, Closed: closed,
	})
}

// AddEquationCurve adds a parametric curve from x(t)/y(t)/z(t) expressions over [t0,t1].
func (s Sketch3D) AddEquationCurve(index int, xExpr, yExpr, zExpr string, t0, t1 float64) (wire.AddSketch3DEntityResult, error) {
	return s.AddEntity(wire.AddSketch3DEntityArgs{
		SketchIndex: index, Kind: string(types.Sketch3DEntityEquationCurve),
		XExpr: xExpr, YExpr: yExpr, ZExpr: zExpr, T0: t0, T1: t1,
	})
}

// AddHelix adds a helical curve. origin [x,y,z] (cm) is the axis base, axis [x,y,z] the
// winding direction (empty ⇒ +Z), radius a unit-bearing start radius. mode selects which
// two of pitch/height/revolutions define the helix; pass the matching unit-bearing
// pitch/height and/or a revolution count. See [wire.AddSketch3DEntityArgs] for the modes.
func (s Sketch3D) AddHelix(index int, origin, axis []float64, radius string, args wire.AddSketch3DEntityArgs) (wire.AddSketch3DEntityResult, error) {
	args.SketchIndex = index
	args.Kind = string(types.Sketch3DEntityHelical)
	args.Points = [][]float64{origin}
	args.Axis = axis
	args.Radius = radius
	return s.AddEntity(args)
}
