// SPDX-License-Identifier: Apache-2.0

package wire

// ImportDWGArgs is the request of [MethodImportDWG]: the host-side path of a .dwg
// file and the target work plane for a 2D drawing. Plane names a plane from
// list_work_planes (e.g. "XY Plane", or a user work plane); empty defaults to the
// first origin plane. A non-planar drawing imports into a 3D sketch and ignores
// Plane. The DWG world origin maps onto the chosen plane's origin.
type ImportDWGArgs struct {
	Path  string `json:"path"`
	Plane string `json:"plane,omitempty"`
}

// ImportDWGResult is the response of [MethodImportDWG]: whether the drawing landed
// in a 3D sketch (vs a 2D sketch on the chosen plane), how many curve entities were
// added, and any per-entity warnings (unmapped/undecodable objects that were
// skipped).
type ImportDWGResult struct {
	Is3D        bool     `json:"is3D"`
	EntityCount int      `json:"entityCount"`
	Warnings    []string `json:"warnings,omitempty"`
}
