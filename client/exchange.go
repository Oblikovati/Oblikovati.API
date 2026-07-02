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
//
// mcp:tool import_file
// mcp:summary Import a CAD file into the active part as an imported-body feature. Format: step|stl|obj|3mf. Path is on the host filesystem (e.g. "/path/EDF.STEP"). Returns the body count and whether the first body came in as a watertight solid. Pair with capture_viewport to SEE the imported geometry.
func (d Documents) Import(req wire.ImportRequest) (wire.ImportResponse, error) {
	return call[wire.ImportResponse](d.c, wire.MethodDocumentsImport, req)
}

// Export writes the active part's bodies to a foreign mesh file at the requested
// resolution, returning the triangle count written and any warnings.
//
// Example:
//
//	resp, err := c.Documents().Export(wire.ExportRequest{Path: "p.stl", Format: "stl", Resolution: "high"})
//
// mcp:tool documents_export
// mcp:summary Writes the active part's bodies to a foreign mesh file at the requested resolution, returning the triangle count written and any warnings.
func (d Documents) Export(req wire.ExportRequest) (wire.ExportResponse, error) {
	return call[wire.ExportResponse](d.c, wire.MethodDocumentsExport, req)
}
