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

// SetBOMStructure sets a placed component's BOM structure as a per-occurrence override — "phantom"
// promotes its children, "default" clears the override so it inherits the definition (#1978).
//
// mcp:tool assembly_set_bom_structure
// mcp:summary Set an occurrence's BOM structure (normal/phantom/reference/purchased/inseparable, or "default" to inherit the definition). Reflected in both BOM views.
func (a Assembly) SetBOMStructure(args wire.SetBOMStructureArgs) (wire.SetBOMStructureResult, error) {
	return call[wire.SetBOMStructureResult](a.c, wire.MethodAssemblySetBOMStructure, args)
}
