// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// AddConstraint is the general 3D geometric-constraint constructor; prefer the typed
// helpers below for the common kinds.
func (s Sketch3D) AddConstraint(args wire.AddSketch3DConstraintArgs) (wire.AddSketch3DConstraintResult, error) {
	var r wire.AddSketch3DConstraintResult
	return r, s.c.call(wire.MethodSketch3DAddConstraint, args, &r)
}

// constrain is the shared helper for the kind-specific constraint constructors.
func (s Sketch3D) constrain(index int, kind types.Geometric3DConstraintKind, entities ...uint64) (wire.AddSketch3DConstraintResult, error) {
	return s.AddConstraint(wire.AddSketch3DConstraintArgs{SketchIndex: index, Kind: string(kind), Entities: entities})
}

// Coincident pins two 3D points together; Collinear forces three onto a line; Concentric
// makes two circle/arc centers coincide.
func (s Sketch3D) Coincident(index int, a, b uint64) (wire.AddSketch3DConstraintResult, error) {
	return s.constrain(index, types.Geo3DCoincident, a, b)
}

func (s Sketch3D) Collinear(index int, a, b, c uint64) (wire.AddSketch3DConstraintResult, error) {
	return s.constrain(index, types.Geo3DCollinear, a, b, c)
}

// Parallel / Perpendicular relate two 3D lines.
func (s Sketch3D) Parallel(index int, l1, l2 uint64) (wire.AddSketch3DConstraintResult, error) {
	return s.constrain(index, types.Geo3DParallel, l1, l2)
}

func (s Sketch3D) Perpendicular(index int, l1, l2 uint64) (wire.AddSketch3DConstraintResult, error) {
	return s.constrain(index, types.Geo3DPerpendicular, l1, l2)
}

// Midpoint pins a point to the midpoint of a line; Ground fixes a point in place.
func (s Sketch3D) Midpoint(index int, point, line uint64) (wire.AddSketch3DConstraintResult, error) {
	return s.constrain(index, types.Geo3DMidpoint, point, line)
}

func (s Sketch3D) Ground(index int, point uint64) (wire.AddSketch3DConstraintResult, error) {
	return s.constrain(index, types.Geo3DGround, point)
}

// ParallelToAxis constrains a line parallel to an origin-frame axis (kind one of
// Geo3DParallelToXAxis/YAxis/ZAxis); ParallelToPlane constrains it parallel to an
// origin-frame plane (Geo3DParallelToXYPlane/XZPlane/YZPlane).
func (s Sketch3D) ParallelToAxis(index int, line uint64, kind types.Geometric3DConstraintKind) (wire.AddSketch3DConstraintResult, error) {
	return s.constrain(index, kind, line)
}

// DeleteConstraint removes the geometric constraint at the given index.
func (s Sketch3D) DeleteConstraint(index, constraintIndex int) (wire.OKResult, error) {
	var r wire.OKResult
	args := wire.DeleteSketch3DConstraintArgs{SketchIndex: index, ConstraintIndex: constraintIndex}
	return r, s.c.call(wire.MethodSketch3DDeleteConstraint, args, &r)
}
