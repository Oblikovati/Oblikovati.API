// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// WorkPoints is the datum-point construction group for the active part. A created point's
// reference can feed a work plane (e.g. a three-point plane through three points) or a
// work-plane redefine's slot re-pick.
type WorkPoints struct{ c *Client }

// WorkPoints returns the work-point construction group.
func (c *Client) WorkPoints() WorkPoints { return WorkPoints{c} }

// Create adds a datum point fixed at position [x, y, z] (model units) and returns its index,
// reference, and name.
//
// mcp:tool create_work_point
// mcp:summary Create a datum work point fixed at a position (at:[x,y,z], model units). Returns its ref (e.g. "point/1") to use as a point input — three points make a three-points work plane, or re-point such a plane's slot via redefine_work_plane.
func (w WorkPoints) Create(args wire.CreateWorkPointArgs) (wire.CreateWorkPointResult, error) {
	var r wire.CreateWorkPointResult
	return r, w.c.call(wire.MethodWorkPointsCreate, args, &r)
}

// At adds a datum point at (x, y, z) — the common spelling of Create.
func (w WorkPoints) At(x, y, z float64) (wire.CreateWorkPointResult, error) {
	return w.Create(wire.CreateWorkPointArgs{At: []float64{x, y, z}})
}
