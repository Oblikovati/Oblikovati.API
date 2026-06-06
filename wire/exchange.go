// SPDX-License-Identifier: Apache-2.0

package wire

// This file is the JSON contract for foreign-format mesh exchange (M17-F04): import a
// mesh file as model geometry, export the active part's bodies to a mesh file. The
// method-name constants live in methods.go; these are the request/response DTOs.

// ImportRequest is the request of [MethodDocumentsImport]: read the file at Path
// (interpreted as Format — an [oblikovati/api/types.ExchangeFormat] string) into the
// active part as an imported-body feature. Options carries free-form translator knobs
// (e.g. "weldTolerance"); unknown keys are ignored.
type ImportRequest struct {
	Path    string            `json:"path"`
	Format  string            `json:"format"`
	Options map[string]string `json:"options,omitempty"`
}

// ImportResponse is the reply of [MethodDocumentsImport]: how many bodies were
// imported, whether the (first) body came in as a watertight solid (vs an open
// surface body), and any non-fatal warnings (non-manifold edges, dropped attributes).
type ImportResponse struct {
	BodyCount int      `json:"bodyCount"`
	Solid     bool     `json:"solid"`
	Warnings  []string `json:"warnings,omitempty"`
}

// ExportRequest is the request of [MethodDocumentsExport]: write the active part's
// bodies to Path in Format at the given tessellation Resolution (an
// [oblikovati/api/types.MeshResolution] string; "" ⇒ medium).
type ExportRequest struct {
	Path       string `json:"path"`
	Format     string `json:"format"`
	Resolution string `json:"resolution,omitempty"`
}

// ExportResponse is the reply of [MethodDocumentsExport]: the total triangle count
// written and any non-fatal warnings.
type ExportResponse struct {
	TriangleCount int      `json:"triangleCount"`
	Warnings      []string `json:"warnings,omitempty"`
}
