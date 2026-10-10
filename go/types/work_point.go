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
	// when the axis is parallel to the plane (no intersection). This is the reference CAD API's
	// WorkPoint.AddByPlaneAndLine.
	WorkPointPlaneAxisIntersection WorkPointKind = "plane-axis-intersection"

	// Reference-model constructors on points/lines/planes (#1842). Each names its inputs in Refs.
	WorkPointOnPoint     WorkPointKind = "point"        // coincident with a referenced point; Refs = [point]
	WorkPointTwoLines    WorkPointKind = "two-lines"    // where two lines intersect; Refs = [line, line]
	WorkPointThreePlanes WorkPointKind = "three-planes" // intersection of three planes; Refs = [plane, plane, plane]

	// WorkPointFaceCenter is the centre of a spherical or toroidal face — the reference CAD API's
	// WorkPoints.AddByCenterOfSphereFace / AddByCenterOfTorusFace. Refs = [face] (a B-rep face
	// reference). It goes unhealthy for a face with no centre point (e.g. a plane or cylinder). #1842.
	WorkPointFaceCenter WorkPointKind = "face-center"

	// WorkPointMidpointOfEdge is the midpoint of a B-rep edge — the reference CAD API's WorkPoint.AddByMidpoint /
	// a curve midpoint. Refs = [edge] (a lineage-key "edge/…" reference or a geometric
	// [GeometricEdgeRef] "edge-geom/…" reference). #1842.
	WorkPointMidpointOfEdge WorkPointKind = "edge-midpoint"

	// WorkPointCurveAndEntity is where a curve meets a surface entity — the reference CAD API's
	// WorkPoints.AddByCurveAndEntity(Curve, Entity, ProximityPoint). Refs = [curve, entity]:
	// the curve (a linear/circular edge or work axis) and the entity it pierces (a face or work
	// plane). When the curve crosses the entity more than once, the create request's Proximity
	// point [x,y,z] (cm) selects the nearest intersection; omitting it takes the first solution.
	// It goes unhealthy when the curve does not meet the entity. #1842.
	WorkPointCurveAndEntity WorkPointKind = "curve-and-entity"

	// WorkPointCentroid is the centroid of a set of connected edges — the reference CAD API's
	// WorkPoints.AddAtCentroid(Entities). Refs = [edge, edge, …] (one or more edge references,
	// each a lineage-key "edge/…" or geometric "edge-geom/…" ref). The point is the length-weighted
	// mean of the referenced edges' midpoints; it goes unhealthy when no edge resolves. #1842.
	WorkPointCentroid WorkPointKind = "centroid"

	// WorkPointCloud is a datum point fixed at a captured point-cloud position — the reference CAD API's
	// WorkPoints.AddByCloudPoint. Refs = [cloudID] where the single ref is the source cloud's id;
	// the create request's At [x,y,z] (cm) gives the frozen position picked from that cloud. This
	// makes the model's AddByCloudPoint constructor reachable over the wire. #1842.
	WorkPointCloud WorkPointKind = "cloud-point"
)
