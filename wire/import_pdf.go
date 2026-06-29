// SPDX-License-Identifier: Apache-2.0

package wire

// ImportPDFArgs is the request of [MethodImportPDF]: the host-side path of a vector
// .pdf file and the target work plane for its 2D geometry. Plane names a plane from
// list_work_planes (e.g. "XY Plane", or a user work plane); empty defaults to the
// first origin plane. Every page imports onto the same plane (one 2D sketch per page),
// the PDF page origin mapping onto the plane origin.
type ImportPDFArgs struct {
	Path  string `json:"path"`
	Plane string `json:"plane,omitempty"`
}

// ImportPDFResult is the response of [MethodImportPDF]: whether any page landed in a 3D
// sketch (always false for the page-flat PDFs this importer targets, but kept parallel
// to the DWG/DXF results), how many curve entities were added across all pages, and any
// per-page warnings (unsupported operators, skipped text/raster content, or pages that
// could not be decoded).
type ImportPDFResult struct {
	Is3D        bool     `json:"is3D"`
	EntityCount int      `json:"entityCount"`
	Warnings    []string `json:"warnings,omitempty"`
}
