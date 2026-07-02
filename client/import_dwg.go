// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// ImportDWG imports an AutoCAD .dwg file into the active part's drawing geometry.
//
// mcp:tool import_dwg
// mcp:summary Import a .dwg file into the active part. A planar drawing becomes a 2D sketch on the chosen plane (plane: a name from list_work_planes, e.g. "XY Plane"; default is the first origin plane), the DWG origin mapping to the plane origin; a drawing with off-plane geometry becomes a 3D sketch. Returns whether it imported as 3D, the entity count, and any skipped-entity warnings.
func (c *Client) ImportDWG(args wire.ImportDWGArgs) (wire.ImportDWGResult, error) {
	return call[wire.ImportDWGResult](c, wire.MethodImportDWG, args)
}
