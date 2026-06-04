// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/Oblikovati/api/types"
	"github.com/Oblikovati/api/wire"
)

// AddEntity is the general 3D entity constructor — the escape hatch covering every kind;
// prefer the typed helpers below for the common constructors.
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

// AddEllipticalArc adds a bounded ellipse spanning startAngle..startAngle+sweepAngle
// (unit-bearing angles).
func (s Sketch3D) AddEllipticalArc(index int, center, axis, majorAxis []float64, majorR, minorR, startAngle, sweepAngle string, construction bool) (wire.AddSketch3DEntityResult, error) {
	return s.AddEntity(wire.AddSketch3DEntityArgs{
		SketchIndex: index, Kind: string(types.Sketch3DEntityEllipticalArc),
		Points: [][]float64{center}, Axis: axis, MajorAxis: majorAxis,
		MajorRadius: majorR, MinorRadius: minorR, StartAngle: startAngle, SweepAngle: sweepAngle,
		Construction: construction,
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
