// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Import reads a foreign mesh file (STL/OBJ/3MF) into the active part as an
// imported-body feature, returning how many bodies came in, whether the first is a
// watertight solid, and any warnings.
//
// Example:
//
//	resp, err := c.Documents().Import(wire.ImportRequest{Path: "bolt.stl", Format: "stl"})
func (d Documents) Import(req wire.ImportRequest) (wire.ImportResponse, error) {
	var r wire.ImportResponse
	return r, d.c.call(wire.MethodDocumentsImport, req, &r)
}

// Export writes the active part's bodies to a foreign mesh file at the requested
// resolution, returning the triangle count written and any warnings.
//
// Example:
//
//	resp, err := c.Documents().Export(wire.ExportRequest{Path: "p.stl", Format: "stl", Resolution: "high"})
func (d Documents) Export(req wire.ExportRequest) (wire.ExportResponse, error) {
	var r wire.ExportResponse
	return r, d.c.call(wire.MethodDocumentsExport, req, &r)
}
