// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// Constrain is the geometric-constraint group for a sketch, reached via [Sketch.Constrain].
// Each helper names a constraint kind; the uint64 arguments are entity ids returned by the
// addEntity helpers (a point id, a line id, or a circle/arc id, per the kind).
type Constrain struct {
	c     *Client
	index int
}

// Constrain returns the geometric-constraint group for the sketch at index.
func (s Sketch) Constrain(index int) Constrain { return Constrain{s.c, index} }

// Add applies a constraint of the given kind to the referenced entities — the escape
// hatch covering every kind; prefer the named helpers below.
//
// mcp:tool add_sketch_constraint
// mcp:summary Add a geometric constraint: {sketchIndex, kind, entities:[ids…]}. kind is coincident|horizontal|vertical|parallel|perpendicular|collinear|concentric|tangent|equalLength|equalRadius|pointOnLine|midpoint|pointOnCircle|fix|ground|symmetric|smooth. Entity arity depends on the kind.
func (g Constrain) Add(kind types.GeometricConstraintKind, entities ...uint64) (wire.AddConstraintResult, error) {
	var r wire.AddConstraintResult
	args := wire.AddConstraintArgs{SketchIndex: g.index, Kind: string(kind), Entities: entities}
	return r, g.c.call(wire.MethodSketchAddConstraint, args, &r)
}

// Coincident makes two points coincident.
func (g Constrain) Coincident(p1, p2 uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintCoincident, p1, p2)
}

// Horizontal aligns two points horizontally; Vertical aligns them vertically.
func (g Constrain) Horizontal(p1, p2 uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintHorizontal, p1, p2)
}

func (g Constrain) Vertical(p1, p2 uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintVertical, p1, p2)
}

// Parallel / Perpendicular / Collinear / EqualLength relate two lines.
func (g Constrain) Parallel(l1, l2 uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintParallel, l1, l2)
}

func (g Constrain) Perpendicular(l1, l2 uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintPerpendicular, l1, l2)
}

func (g Constrain) Collinear(l1, l2 uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintCollinear, l1, l2)
}

func (g Constrain) EqualLength(l1, l2 uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintEqualLength, l1, l2)
}

// Concentric / EqualRadius relate two circular curves (circles/arcs).
func (g Constrain) Concentric(c1, c2 uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintConcentric, c1, c2)
}

func (g Constrain) EqualRadius(c1, c2 uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintEqualRadius, c1, c2)
}

// Tangent makes a line tangent to a circular curve (line ref first, curve ref second).
func (g Constrain) Tangent(line, curve uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintTangent, line, curve)
}

// PointOnLine / Midpoint constrain a point onto / to the midpoint of a line.
func (g Constrain) PointOnLine(point, line uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintPointOnLine, point, line)
}

func (g Constrain) Midpoint(point, line uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintMidpoint, point, line)
}

// PointOnCircle constrains a point onto a circular curve.
func (g Constrain) PointOnCircle(point, curve uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintPointOnCircle, point, curve)
}

// Fix grounds a point in place.
func (g Constrain) Fix(point uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintFix, point)
}

// Ground fixes every point of an entity (the whole geometry) in place.
func (g Constrain) Ground(entity uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintGround, entity)
}

// Offset holds two lines parallel at their current perpendicular distance.
func (g Constrain) Offset(line1, line2 uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintOffset, line1, line2)
}

// PatternLink rigidly links a member point to a seed point at their current offset.
func (g Constrain) PatternLink(seed, member uint64) (wire.AddConstraintResult, error) {
	return g.Add(types.GeoConstraintPattern, seed, member)
}

// Delete removes the geometric constraint at the given collection index.
//
// mcp:tool delete_sketch_constraint
// mcp:summary Delete a geometric constraint by its index (see list_sketch_constraints).
func (g Constrain) Delete(constraintIndex int) (wire.OKResult, error) {
	var r wire.OKResult
	args := wire.DeleteConstraintArgs{SketchIndex: g.index, ConstraintIndex: constraintIndex}
	return r, g.c.call(wire.MethodSketchDeleteConstraint, args, &r)
}
