// SPDX-License-Identifier: Apache-2.0

package wire

// AddSketch3DConstraintArgs is the request of [MethodSketch3DAddConstraint] — the
// discriminated 3D geometric-constraint constructor. Kind is the constraint
// ([oblikovati.org/api/types.Geometric3DConstraintKind]). Entities are the session
// ids of the geometry it relates, in the kind's expected order (parallel/perpendicular:
// two line ids; midpoint: point id + line id; ground: a point id; parallelToAxis/Plane:
// a single line id; coincident/concentric: two point ids; collinear: three point ids;
// tangent/smooth: two curve ids (line/arc/spline — the join lands on their nearest
// endpoints, and smooth needs at least one spline); equal: two circle/helix ids (their
// radius DOFs); splineFitPoints: a fit-spline id + a point id (attached to the nearest
// fit point); helical: a helix id + the circle it starts on, with parallel axes;
// bend: an arc id + the two lines it joins — the bend radius is captured from the
// arc's current geometry).
type AddSketch3DConstraintArgs struct {
	SketchIndex int      `json:"sketchIndex"`
	Kind        string   `json:"kind"`
	Entities    []uint64 `json:"entities,omitempty"`
}

// AddSketch3DConstraintResult is the response of [MethodSketch3DAddConstraint]: the new
// constraint's index in the sketch's geometric-constraint collection, its kind, and the
// sketch's resulting degree-of-freedom count.
type AddSketch3DConstraintResult struct {
	Index int    `json:"index"`
	Kind  string `json:"kind"`
	DOF   int    `json:"dof"`
}

// DeleteSketch3DConstraintArgs is the request of [MethodSketch3DDeleteConstraint]: which
// sketch and which geometric constraint (by index) to remove.
type DeleteSketch3DConstraintArgs struct {
	SketchIndex     int `json:"sketchIndex"`
	ConstraintIndex int `json:"constraintIndex"`
}
