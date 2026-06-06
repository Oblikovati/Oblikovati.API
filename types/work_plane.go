// SPDX-License-Identifier: Apache-2.0

package types

// WorkPlaneKind names a datum-plane constructor, the discriminator of a work-plane
// create request. The string values are the stable
// wire vocabulary — treat them as frozen.
//
// This is the canonical, Apache-2.0 definition of the public kind names; the GPL host
// maps each to its model constructor.
type WorkPlaneKind string

const (
	// Reference-model constructors (built on planes/axes/points).
	WorkPlaneOffset         WorkPlaneKind = "plane-offset"     // parallel to a plane, offset
	WorkPlaneThreePoints    WorkPlaneKind = "three-points"     // through three points
	WorkPlaneFixed          WorkPlaneKind = "fixed-frame"      // fixed origin + X/Y axes (AddFixed)
	WorkPlanePlaneAndPoint  WorkPlaneKind = "plane-point"      // parallel to a plane, through a point
	WorkPlaneTwoPlanes      WorkPlaneKind = "two-planes"       // bisector of two planes
	WorkPlaneLinePlaneAngle WorkPlaneKind = "line-plane-angle" // through a line, at an angle to a plane
	WorkPlaneTwoLines       WorkPlaneKind = "two-lines"        // from two lines
	WorkPlaneNormalToCurve  WorkPlaneKind = "normal-to-curve"  // through a point, normal to a curve

	// Surface-tangent constructors (built on a B-rep face reference).
	WorkPlaneTorusMidPlane   WorkPlaneKind = "torus-midplane" // mid-plane of a torus face
	WorkPlanePointAndTangent WorkPlaneKind = "point-tangent"  // tangent at a point on a surface
	WorkPlanePlaneAndTangent WorkPlaneKind = "plane-tangent"  // parallel to a plane, tangent to a surface
	WorkPlaneLineAndTangent  WorkPlaneKind = "line-tangent"   // through a line, tangent to a surface
)

// Well-known work-feature references — the part's static origin coordinate frame, valid
// as entries in a work-plane create request's Refs (e.g. the base plane of an offset).
// Other references (to user work planes, or a B-rep face) come back from workPlanes.list
// and model selection.
const (
	WorkRefCenter  = "origin/point/center"
	WorkRefXAxis   = "origin/axis/x"
	WorkRefYAxis   = "origin/axis/y"
	WorkRefZAxis   = "origin/axis/z"
	WorkRefXYPlane = "origin/plane/xy"
	WorkRefXZPlane = "origin/plane/xz"
	WorkRefYZPlane = "origin/plane/yz"
)
