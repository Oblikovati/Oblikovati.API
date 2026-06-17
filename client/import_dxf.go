// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// ImportDXF imports an ASCII .dxf file into the active part's drawing geometry.
//
// mcp:tool import_dxf
// mcp:summary Import a .dxf file into the active part. A planar drawing becomes a 2D sketch on the chosen plane (plane: a name from list_work_planes, e.g. "XY Plane"; default is the first origin plane), the DXF origin mapping to the plane origin; a drawing with off-plane geometry becomes a 3D sketch. Returns whether it imported as 3D, the entity count, and any skipped-entity warnings.
func (c *Client) ImportDXF(args wire.ImportDXFArgs) (wire.ImportDXFResult, error) {
	var r wire.ImportDXFResult
	return r, c.call(wire.MethodImportDXF, args, &r)
}

// ExportDXF writes the active 2D sketch to an ASCII .dxf file.
//
// mcp:tool export_dxf
// mcp:summary Export the active 2D sketch to a .dxf file at path. version selects the DXF generation ("r2000" or "r2018"; default r2000). Returns how many sketch curves were written.
func (c *Client) ExportDXF(args wire.ExportDXFArgs) (wire.ExportDXFResult, error) {
	var r wire.ExportDXFResult
	return r, c.call(wire.MethodExportDXF, args, &r)
}
