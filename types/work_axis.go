// SPDX-License-Identifier: Apache-2.0

package types

// WorkAxisKind names a datum-axis constructor, the discriminator of a work-axis
// create request. The string values are the stable wire vocabulary — treat them as
// frozen.
//
// This is the canonical, Apache-2.0 definition of the public kind names; the GPL host
// maps each to its model constructor.
type WorkAxisKind string

const (
	// WorkAxisLine is a raw grounded axis: fixed geometry from an origin point and a
	// direction vector, no references (the model's grounded-axis constructor). It is the
	// axis a revolve or sweep spins about when the axis is not one of the origin axes —
	// e.g. an axis matching a sketch line.
	WorkAxisLine WorkAxisKind = "line" // grounded origin + direction (fixedAxisDef)

	// Reference-model constructors (built on points/planes/lines).
	WorkAxisTwoPoints         WorkAxisKind = "two-points"         // through two points (AddByTwoPoints)
	WorkAxisPlaneIntersection WorkAxisKind = "plane-intersection" // where two planes meet (AddByPlaneIntersection)
	WorkAxisPointAndPlane     WorkAxisKind = "point-and-plane"    // through a point, normal to a plane (#1840)
	WorkAxisLineAndPoint      WorkAxisKind = "line-and-point"     // through a point, parallel to a line (#1840)
	WorkAxisLineAndPlane      WorkAxisKind = "line-and-plane"     // a line projected onto a plane (#1840)

	// WorkAxisRevolvedFace is the axis of revolution of a cylindrical, conical, or toroidal face —
	// Inventor's WorkAxes.AddByRevolvedFace. Refs = [face] (a B-rep face reference). It goes
	// unhealthy for a face with no axis of revolution (e.g. a plane or sphere). #1840.
	WorkAxisRevolvedFace WorkAxisKind = "revolved-face"
)

// The origin coordinate axes (WorkRefXAxis / WorkRefYAxis / WorkRefZAxis, defined
// alongside the other origin-frame refs in work_plane.go) are grounded axis references
// that already exist — valid wherever an axis input is accepted (a revolve axis, a
// normal-to-curve work plane). They are referenced, never created. User axis references
// come back from workAxes.list.
