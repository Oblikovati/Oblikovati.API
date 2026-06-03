// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/Oblikovati/api/types"
	"github.com/Oblikovati/api/wire"
)

// AddEntity is the general entity constructor — the escape hatch covering every kind and
// variant; prefer the typed helpers below for the common constructors.
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

// AddEllipse adds a full ellipse from a center [x,y] (cm), a major-axis direction [x,y],
// and unit-bearing major/minor radii ("20 mm").
func (s Sketch) AddEllipse(index int, center, axis []float64, majorR, minorR string, construction bool) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityEllipse),
		Points: [][]float64{center}, Axis: axis, MajorRadius: majorR, MinorRadius: minorR,
		Construction: construction,
	})
}

// AddEllipticalArc adds an elliptical arc bounded by unit-bearing start/end angles
// ("0 deg", "90 deg") measured in the ellipse's major/minor frame.
func (s Sketch) AddEllipticalArc(index int, center, axis []float64, majorR, minorR, startAngle, endAngle string, construction bool) (wire.AddSketchEntityResult, error) {
	return s.AddEntity(wire.AddSketchEntityArgs{
		SketchIndex: index, Kind: string(types.SketchEntityEllipticalArc),
		Points: [][]float64{center}, Axis: axis, MajorRadius: majorR, MinorRadius: minorR,
		StartAngle: startAngle, EndAngle: endAngle, Construction: construction,
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
