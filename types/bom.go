// SPDX-License-Identifier: Apache-2.0

package types

// BOMStructure is how a placed component participates in the bill of materials (M11-F05,
// Oblikovati/Oblikovati#730). The wire spelling matches the model's bom.Structure.String().
type BOMStructure string

const (
	// BOMNormal is a counted row whose sub-assembly children are expanded.
	BOMNormal BOMStructure = "normal"
	// BOMPhantom is not a row of its own: its children are promoted into its parent.
	BOMPhantom BOMStructure = "phantom"
	// BOMReference is shown for context but never counted.
	BOMReference BOMStructure = "reference"
	// BOMPurchased is a bought item counted as one line; its children are not broken out.
	BOMPurchased BOMStructure = "purchased"
	// BOMInseparable is a welded/glued sub-assembly counted as one line; not broken out.
	BOMInseparable BOMStructure = "inseparable"
)

// BOMViewKind selects a bill-of-materials view.
type BOMViewKind string

const (
	// BOMStructured is the hierarchical view: top-level rows with sub-assembly children
	// nested, quantities counted per parent.
	BOMStructured BOMViewKind = "structured"
	// BOMPartsOnly is the flat view: every unique counted part once, with its total
	// quantity across the whole assembly.
	BOMPartsOnly BOMViewKind = "partsOnly"
)
