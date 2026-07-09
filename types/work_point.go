// SPDX-License-Identifier: Apache-2.0

package types

// WorkPointKind names a datum-point constructor, the discriminator of a work-point
// create request. The string values are the stable wire vocabulary — treat them as
// frozen.
//
// This is the canonical, Apache-2.0 definition of the public kind names; the GPL host
// maps each to its model constructor. An empty kind means [WorkPointPosition] so the
// original position-only request (just "at") keeps working unchanged.
type WorkPointKind string

const (
	// WorkPointPosition is a datum point fixed at an absolute position — the request's
	// At [x, y, z] in model units (the model's AddByPosition constructor). This is the
	// default when Kind is omitted.
	WorkPointPosition WorkPointKind = "position" // fixed [x,y,z] (AddByPosition)

	// WorkPointPlaneAxisIntersection is the point where a referenced axis pierces a
	// referenced plane (the model's AddByPlaneAndAxisIntersection constructor). Its two
	// references are given, in order, as Refs = [plane, axis]. It reports healthy=false
	// when the axis is parallel to the plane (no intersection). This is Inventor's
	// WorkPoint.AddByPlaneAndLine.
	WorkPointPlaneAxisIntersection WorkPointKind = "plane-axis-intersection"

	// Reference-model constructors on points/lines/planes (#1842). Each names its inputs in Refs.
	WorkPointOnPoint     WorkPointKind = "point"        // coincident with a referenced point; Refs = [point]
	WorkPointTwoLines    WorkPointKind = "two-lines"    // where two lines intersect; Refs = [line, line]
	WorkPointThreePlanes WorkPointKind = "three-planes" // intersection of three planes; Refs = [plane, plane, plane]
)
