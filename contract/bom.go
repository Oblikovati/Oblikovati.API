// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// BOMRow is the scalar read surface of one bill-of-materials line (M11-F05,
// Oblikovati/Oblikovati#730): its item number, the component's part number/description/
// structure, the quantity at this level, and — in a structured view — its nested child
// rows. The host adapter wraps the model's bom.Row.
type BOMRow interface {
	ItemNumber() int
	PartNumber() string
	Description() string
	Structure() types.BOMStructure
	Quantity() int
	Children() []BOMRow
}

// BOMView is the scalar read surface of one bill-of-materials view — its kind and rows.
type BOMView interface {
	Kind() types.BOMViewKind
	Rows() []BOMRow
}

// BOM is the scalar read surface of an assembly's bill of materials, derived live from its
// occurrence tree: a structured (nested) view and a parts-only (flat, totaled) view. Host:
// the bomapi adapter over model/bom.
type BOM interface {
	Structured() BOMView
	PartsOnly() BOMView
}
