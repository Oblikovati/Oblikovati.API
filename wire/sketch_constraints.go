// SPDX-License-Identifier: Apache-2.0

package wire

// AddConstraintArgs is the request of [MethodSketchAddConstraint] — the discriminated
// geometric-constraint constructor. Kind is a
// [github.com/Oblikovati/api/types.GeometricConstraintKind]; Entities are the session ids
// of the geometry it relates (points/lines/curves), in the kind's expected order:
//
//   - coincident/horizontal/vertical: two point ids (or one line id)
//   - parallel/perpendicular/collinear/equalLength: two line ids
//   - concentric/equalRadius: two circular-curve ids
//   - tangent: a line id + a circular-curve id, or two circular-curve ids
//   - pointOnLine/midpoint: a point id + a line id
//   - pointOnCircle: a point id + a circular-curve id
//   - symmetry: two point ids + a line id (the mirror line)
//   - fix: one point id
type AddConstraintArgs struct {
	SketchIndex int      `json:"sketchIndex"`
	Kind        string   `json:"kind"`
	Entities    []uint64 `json:"entities"`
}

// AddConstraintResult is the response of [MethodSketchAddConstraint]: the new
// constraint's index in the sketch's geometric-constraint collection, its kind, and the
// sketch's resulting degrees of freedom (so a caller sees the DOF drop).
type AddConstraintResult struct {
	Index int    `json:"index"`
	Kind  string `json:"kind"`
	DOF   int    `json:"dof"`
}

// DeleteConstraintArgs is the request of [MethodSketchDeleteConstraint]: which sketch and
// the index of the geometric constraint to remove (as listed by [MethodSketchConstraints]).
type DeleteConstraintArgs struct {
	SketchIndex     int `json:"sketchIndex"`
	ConstraintIndex int `json:"constraintIndex"`
}
