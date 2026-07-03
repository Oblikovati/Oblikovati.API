// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// The assembly bill-of-materials operations (M11-F05, Oblikovati/Oblikovati#730) extend
// the Assembly group: read a structured or parts-only BOM view of the active assembly, and
// export a view to CSV with optional custom property columns.

// BOMView reads the given view of the active assembly's BOM, e.g.
// BOMView(types.BOMPartsOnly).
//
// mcp:tool assembly_bom_view
// mcp:summary Reads the given view of the active assembly's BOM, e.g.
func (a Assembly) BOMView(view types.BOMViewKind) (wire.BOMViewResult, error) {
	return call[wire.BOMViewResult](a.c, wire.MethodAssemblyBOMView, wire.BOMViewArgs{View: view})
}

// BOMExport exports the given view to CSV, adding a column for each named component
// property beyond the standard set, e.g.
// BOMExport(wire.BOMExportArgs{View: types.BOMStructured, Columns: []string{"Material"}}).
//
// mcp:tool assembly_bom_export
// mcp:summary Exports the given view to CSV, adding a column for each named component property beyond the standard set, e.g.
func (a Assembly) BOMExport(args wire.BOMExportArgs) (wire.BOMExportResult, error) {
	return call[wire.BOMExportResult](a.c, wire.MethodAssemblyBOMExport, args)
}
