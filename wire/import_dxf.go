// SPDX-License-Identifier: Apache-2.0

package wire

// ImportDXFArgs is the request of [MethodImportDXF]: the host-side path of a .dxf file and
// the target work plane for a 2D drawing. Plane names a plane from list_work_planes (e.g.
// "XY Plane", or a user work plane); empty defaults to the first origin plane. A non-planar
// drawing imports into a 3D sketch and ignores Plane. The DXF world origin maps onto the
// chosen plane's origin. (DXF is DWG's open text sibling; the import behaves identically.)
type ImportDXFArgs struct {
	Path  string `json:"path"`
	Plane string `json:"plane,omitempty"`
}

// ImportDXFResult is the response of [MethodImportDXF]: whether the drawing landed in a 3D
// sketch (vs a 2D sketch on the chosen plane), how many curve entities were added, and any
// per-entity warnings (unmapped/undecodable objects that were skipped).
type ImportDXFResult struct {
	Is3D        bool     `json:"is3D"`
	EntityCount int      `json:"entityCount"`
	Warnings    []string `json:"warnings,omitempty"`
}

// ExportDXFArgs is the request of [MethodExportDXF]: the host-side path to write and the DXF
// version to target. The active 2D sketch is exported. Version is a types.DXFVersion string
// ("r2000" or "r2018"); empty defaults to r2000.
type ExportDXFArgs struct {
	Path    string `json:"path"`
	Version string `json:"version,omitempty"`
}

// ExportDXFResult is the response of [MethodExportDXF]: how many sketch curves were written.
type ExportDXFResult struct {
	EntityCount int `json:"entityCount"`
}
