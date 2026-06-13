// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The assembly bill-of-materials surface (M11-F05, Oblikovati/Oblikovati#730): read a
// structured (nested) or parts-only (flat) BOM view of the active assembly, and export a
// view to CSV with optional custom property columns. The BOM derives live from the current
// occurrence tree (component metadata — part number/description/structure/properties —
// comes from the placed definitions; #718 wires it from document iProperties).

// BOMRowInfo is one bill-of-materials line: its 1-based item number, the component's part
// number/description/structure, the quantity at this level, its custom properties, and (in
// the structured view) the nested child rows of an expanded sub-assembly.
type BOMRowInfo struct {
	ItemNumber  int                `json:"itemNumber"`
	PartNumber  string             `json:"partNumber,omitempty"`
	Description string             `json:"description,omitempty"`
	Structure   types.BOMStructure `json:"structure"`
	Quantity    int                `json:"quantity"`
	Properties  map[string]string  `json:"properties,omitempty"`
	Children    []BOMRowInfo       `json:"children,omitempty"`
}

// BOMViewArgs is the request of [MethodAssemblyBOMView]: read the View
// ([types.BOMStructured] | [types.BOMPartsOnly]) of the active assembly's BOM.
type BOMViewArgs struct {
	View types.BOMViewKind `json:"view"`
}

// BOMViewResult is the reply of [MethodAssemblyBOMView]: the requested view and its rows
// (the structured view's rows carry nested Children).
type BOMViewResult struct {
	View types.BOMViewKind `json:"view"`
	Rows []BOMRowInfo      `json:"rows"`
}

// BOMExportArgs is the request of [MethodAssemblyBOMExport]: export the View of the active
// assembly's BOM to CSV. Beyond the standard columns (item, part number, description, qty,
// structure) each name in Columns adds a column sourced from that component property.
type BOMExportArgs struct {
	View    types.BOMViewKind `json:"view"`
	Columns []string          `json:"columns,omitempty"`
}

// BOMExportResult is the reply of [MethodAssemblyBOMExport]: the rendered CSV document.
type BOMExportResult struct {
	CSV string `json:"csv"`
}
