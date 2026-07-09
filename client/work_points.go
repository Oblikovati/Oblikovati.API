// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// WorkPoints is the datum-point construction group for the active part or assembly: Create is
// the general constructor and the typed helpers wrap each datum-point constructor. A created
// point's reference can feed a work plane (e.g. a three-point plane through three points) or a
// work-plane redefine's slot re-pick.
type WorkPoints struct{ c *Client }

// WorkPoints returns the work-point construction group.
func (c *Client) WorkPoints() WorkPoints { return WorkPoints{c} }

// List returns the active model's datum points (origin centre first, then user points) with their
// position, kind, origin flag, visibility, and health (#1842).
//
// mcp:tool list_work_points
// mcp:summary List the work points of the active part or assembly (origin + user). Each point reports its ref, position [x,y,z] (database units, cm), kind, whether it is the origin centre, its visibility, and health.
func (w WorkPoints) List() (wire.ListWorkPointsResult, error) {
	return call[wire.ListWorkPointsResult](w.c, wire.MethodWorkPointsList, nil)
}

// Create adds a datum point from an explicit request — the escape hatch covering every kind
// (position, plane-axis-intersection); prefer the typed helpers below for the common
// constructors. Returns the point's index, reference, name, and health.
//
// mcp:tool create_work_point
// mcp:summary Create a datum work point. kind "position" (default): fixed at at:[x,y,z] (model units). kind "plane-axis-intersection": refs:[plane, axis] — the point where the axis pierces the plane. Returns its ref (e.g. "point/1") to use as a point input — three points make a three-points work plane, or re-point such a plane's slot via redefine_work_plane.
func (w WorkPoints) Create(args wire.CreateWorkPointArgs) (wire.CreateWorkPointResult, error) {
	return call[wire.CreateWorkPointResult](w.c, wire.MethodWorkPointsCreate, args)
}

// At adds a datum point at (x, y, z) — the common spelling of Create.
func (w WorkPoints) At(x, y, z float64) (wire.CreateWorkPointResult, error) {
	return w.Create(wire.CreateWorkPointArgs{At: []float64{x, y, z}})
}

// PlaneAxisIntersection adds the point where an axis reference pierces a plane reference. The
// result reports healthy=false when the axis is parallel to the plane (no intersection).
func (w WorkPoints) PlaneAxisIntersection(plane, axis string) (wire.CreateWorkPointResult, error) {
	return w.Create(wire.CreateWorkPointArgs{Kind: string(types.WorkPointPlaneAxisIntersection), Refs: []string{plane, axis}})
}

// OnPoint adds a datum point coincident with a referenced point (#1842).
func (w WorkPoints) OnPoint(point string) (wire.CreateWorkPointResult, error) {
	return w.Create(wire.CreateWorkPointArgs{Kind: string(types.WorkPointOnPoint), Refs: []string{point}})
}

// TwoLines adds a datum point where two line references intersect (healthy=false if parallel/skew).
func (w WorkPoints) TwoLines(line1, line2 string) (wire.CreateWorkPointResult, error) {
	return w.Create(wire.CreateWorkPointArgs{Kind: string(types.WorkPointTwoLines), Refs: []string{line1, line2}})
}

// ThreePlanes adds a datum point at the intersection of three plane references.
func (w WorkPoints) ThreePlanes(plane1, plane2, plane3 string) (wire.CreateWorkPointResult, error) {
	return w.Create(wire.CreateWorkPointArgs{Kind: string(types.WorkPointThreePlanes), Refs: []string{plane1, plane2, plane3}})
}

// FaceCenter adds a datum point at the centre of a spherical or toroidal face reference. Reports
// healthy=false for a face with no centre point (#1842).
func (w WorkPoints) FaceCenter(face string) (wire.CreateWorkPointResult, error) {
	return w.Create(wire.CreateWorkPointArgs{Kind: string(types.WorkPointFaceCenter), Refs: []string{face}})
}

// MidpointOfEdge adds a datum point at the midpoint of an edge reference — a lineage-key "edge/…"
// ref from a pick, or a geometric ref from types.GeometricEdgeRef.Ref() (ADR-0040) (#1842).
func (w WorkPoints) MidpointOfEdge(edge string) (wire.CreateWorkPointResult, error) {
	return w.Create(wire.CreateWorkPointArgs{Kind: string(types.WorkPointMidpointOfEdge), Refs: []string{edge}})
}
