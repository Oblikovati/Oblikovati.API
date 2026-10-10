// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// ImportPDF imports a vector .pdf (a CAD drawing plotted to PDF) into the active part's
// drawing geometry.
//
// mcp:tool import_pdf
// mcp:summary Import a vector .pdf file (a CAD drawing plotted to PDF, e.g. from AutoCAD) into the active part. Each page's vector paths become a 2D sketch on the chosen plane (plane: a name from list_work_planes, e.g. "XY Plane"; default is the first origin plane), the page origin mapping to the plane origin. Text and raster images in the page are skipped. Returns the entity count and any per-page warnings.
func (c *Client) ImportPDF(args wire.ImportPDFArgs) (wire.ImportPDFResult, error) {
	return call[wire.ImportPDFResult](c, wire.MethodImportPDF, args)
}
