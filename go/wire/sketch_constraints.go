// SPDX-License-Identifier: Apache-2.0

package wire

// AddConstraintArgs is the request of [MethodSketchAddConstraint] — the discriminated
// geometric-constraint constructor. Kind is a
// [oblikovati.org/api/types.GeometricConstraintKind]; Entities are the session ids
// of the geometry it relates (points/lines/curves), in the kind's expected order:
//
//   - horizontal/vertical: ONE line id (or one ellipse id) makes it horizontal/
//     vertical; TWO point ids level the points (the align form, also reachable
//     as horizontalAlign/verticalAlign) (#1871)
//   - horizontalAlign/verticalAlign: two point ids
//   - coincident: two point ids
//   - parallel/perpendicular/collinear/equalLength: two line ids; parallel/
//     perpendicular/collinear also accept an ellipse id per operand, its
//     constrained direction selected by the UseEllipse*MajorAxis flags (#1879)
//   - concentric/equalRadius: two circular-curve ids
//   - tangent: a line id + a circular-curve id, or two circular-curve ids
//   - pointOnLine: a point id + a line id
//   - midpoint: a point id + a line id, or a point id + an arc id (#1872)
//   - pointOnCircle: a point id + a circular-curve id
//   - symmetry: two entity ids (both points, both lines, or both circular
//     curves) + a line id (the mirror line) (#1870)
//   - fix: one point id
//   - custom: any entity ids to tag; ClientID is the owning add-in id
//     (required) and Name the record's name — an attribute-carrying marker,
//     not a solver constraint (M06-F11, Oblikovati/Oblikovati#626)
//
// The UseEllipse*MajorAxis flags select which axis of an ellipse operand is the
// constrained direction (major when nil/true, minor when false) — the reference CAD API's
// UseEllipseMajorAxis / UseEllipseOneMajorAxis / UseEllipseTwoMajorAxis (#1879).
// UseEllipseMajorAxis applies to the single-operand horizontal/vertical form;
// UseEllipseOneMajorAxis / UseEllipseTwoMajorAxis to the first / second operand
// of parallel/perpendicular/collinear. They are ignored for non-ellipse operands.
type AddConstraintArgs struct {
	SketchIndex            int      `json:"sketchIndex"`
	Kind                   string   `json:"kind"`
	Entities               []uint64 `json:"entities"`
	ClientID               string   `json:"clientId,omitempty"`
	Name                   string   `json:"name,omitempty"`
	UseEllipseMajorAxis    *bool    `json:"useEllipseMajorAxis,omitempty"`
	UseEllipseOneMajorAxis *bool    `json:"useEllipseOneMajorAxis,omitempty"`
	UseEllipseTwoMajorAxis *bool    `json:"useEllipseTwoMajorAxis,omitempty"`
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
