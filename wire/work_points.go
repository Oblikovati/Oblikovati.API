// SPDX-License-Identifier: Apache-2.0

package wire

// CreateWorkPointArgs is the request of [MethodWorkPointsCreate]: a datum point fixed at a
// position. At is the point [x, y, z] in model units.
type CreateWorkPointArgs struct {
	At []float64 `json:"at"`
}

// CreateWorkPointResult is the response of [MethodWorkPointsCreate]: the new point's index in
// the work-points collection, its stable reference (usable as a point input to a work plane —
// e.g. a three-point plane — or to [MethodWorkPlanesRedefine]'s Repick), and its name.
type CreateWorkPointResult struct {
	Index int    `json:"index"`
	Ref   string `json:"ref"`
	Name  string `json:"name"`
}
