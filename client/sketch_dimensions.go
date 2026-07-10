// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// Dimension is the dimensional-constraint group for a sketch, reached via
// [Sketch.Dimension]. Each helper names a dimension kind; the uint64 arguments are entity
// ids and the expression is unit-bearing ("40 mm", "30 deg").
type Dimension struct {
	c     *Client
	index int
}

// Dimension returns the dimensional-constraint group for the sketch at index.
func (s Sketch) Dimension(index int) Dimension { return Dimension{s.c, index} }

// Add applies a dimension of the given kind — the escape hatch; prefer the named helpers.
//
// mcp:tool add_sketch_dimension
// mcp:summary Add a dimensional constraint: {sketchIndex, kind, entities:[ids…], expression}. kind is distance|radius|diameter|angle|arcLength|offset; expression carries units, e.g. "40 mm".
func (g Dimension) Add(kind types.DimensionConstraintKind, expression string, entities ...uint64) (wire.AddDimensionResult, error) {
	args := wire.AddDimensionArgs{SketchIndex: g.index, Kind: string(kind), Entities: entities, Expression: expression}
	return call[wire.AddDimensionResult](g.c, wire.MethodSketchAddDimension, args)
}

// AddWith is the full-control constructor — it passes the whole [wire.AddDimensionArgs]
// through, so callers can set Driven, TextPoint or LinearDiameter at create (#1875); the
// group's SketchIndex is filled in. Prefer the named helpers for the common cases.
func (g Dimension) AddWith(args wire.AddDimensionArgs) (wire.AddDimensionResult, error) {
	args.SketchIndex = g.index
	return call[wire.AddDimensionResult](g.c, wire.MethodSketchAddDimension, args)
}

// Distance dimensions the distance between two points.
func (g Dimension) Distance(p1, p2 uint64, expression string) (wire.AddDimensionResult, error) {
	return g.Add(types.DimConstraintDistance, expression, p1, p2)
}

// Angle dimensions the angle between two lines.
func (g Dimension) Angle(l1, l2 uint64, expression string) (wire.AddDimensionResult, error) {
	return g.Add(types.DimConstraintAngle, expression, l1, l2)
}

// Radius / Diameter dimension a circle.
func (g Dimension) Radius(circle uint64, expression string) (wire.AddDimensionResult, error) {
	return g.Add(types.DimConstraintRadius, expression, circle)
}

func (g Dimension) Diameter(circle uint64, expression string) (wire.AddDimensionResult, error) {
	return g.Add(types.DimConstraintDiameter, expression, circle)
}

// ArcLength dimensions an arc's length.
func (g Dimension) ArcLength(arc uint64, expression string) (wire.AddDimensionResult, error) {
	return g.Add(types.DimConstraintArcLength, expression, arc)
}

// Offset dimensions the perpendicular distance from a point to a line.
func (g Dimension) Offset(point, line uint64, expression string) (wire.AddDimensionResult, error) {
	return g.Add(types.DimConstraintOffset, expression, point, line)
}

// ThreePointAngle dimensions the angle a–vertex–b.
func (g Dimension) ThreePointAngle(vertex, a, b uint64, expression string) (wire.AddDimensionResult, error) {
	return g.Add(types.DimConstraintThreePointAngle, expression, vertex, a, b)
}

// EllipseRadius dimensions an ellipse's major radius.
func (g Dimension) EllipseRadius(ellipse uint64, expression string) (wire.AddDimensionResult, error) {
	return g.Add(types.DimConstraintEllipseRadius, expression, ellipse)
}

// OffsetSpline drives the offset distance of an offset-spline entity from its parent (#1874).
func (g Dimension) OffsetSpline(offsetSpline uint64, expression string) (wire.AddDimensionResult, error) {
	return g.Add(types.DimConstraintOffsetSpline, expression, offsetSpline)
}

// Drive edits a dimension's value (a unit-bearing expression; empty leaves it unchanged).
func (g Dimension) Drive(dimensionIndex int, expression string) (wire.OKResult, error) {
	return g.edit(wire.DriveDimensionArgs{SketchIndex: g.index, DimensionIndex: dimensionIndex, Expression: expression})
}

// SetDriven flips a dimension between driving (constrains) and driven (reports).
func (g Dimension) SetDriven(dimensionIndex int, driven bool) (wire.OKResult, error) {
	return g.edit(wire.DriveDimensionArgs{SketchIndex: g.index, DimensionIndex: dimensionIndex, SetDriven: true, Driven: driven})
}

// SetLimits clamps a dimension's value to [min, max] (model units) when driven.
func (g Dimension) SetLimits(dimensionIndex int, min, max float64) (wire.OKResult, error) {
	return g.edit(wire.DriveDimensionArgs{SketchIndex: g.index, DimensionIndex: dimensionIndex, SetLimits: true, Min: min, Max: max})
}

// mcp:tool drive_sketch_dimension
// mcp:summary Change a dimension's expression (and optionally its driven flag / animation limits) and recompute.
func (g Dimension) edit(args wire.DriveDimensionArgs) (wire.OKResult, error) {
	return call[wire.OKResult](g.c, wire.MethodSketchDriveDimension, args)
}
