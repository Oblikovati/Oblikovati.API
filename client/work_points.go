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
