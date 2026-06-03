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
