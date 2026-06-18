// SPDX-License-Identifier: Apache-2.0

package types

// Surface texture value types (M14-F03 PBI-142, #389). A surface texture symbol states a surface's
// finish requirement (e.g. a roughness Ra value) with the ISO 1302 checkmark glyph. These are the
// canonical Apache-2.0 definitions; the GPL model draws the symbol and the roughness text.

// MaterialRemoval is the surface-texture symbol variant — whether material removal (machining) is
// allowed, required or forbidden. The zero value is MaterialRemovalAny (the basic √ symbol).
type MaterialRemoval int32

const (
	// MaterialRemovalAny is the basic surface texture symbol (the open checkmark): material removal
	// is permitted but not mandated.
	MaterialRemovalAny MaterialRemoval = iota
	// MaterialRemovalRequired adds the horizontal bar across the checkmark: machining is required.
	MaterialRemovalRequired
	// MaterialRemovalProhibited adds a circle in the checkmark's vertex: machining is forbidden.
	MaterialRemovalProhibited
)

var materialRemovalNames = map[MaterialRemoval]string{
	MaterialRemovalAny:        "any",
	MaterialRemovalRequired:   "required",
	MaterialRemovalProhibited: "prohibited",
}

// String returns the variant's wire spelling ("any", "required", "prohibited").
func (m MaterialRemoval) String() string { return enumName(materialRemovalNames, m) }

// ParseMaterialRemoval resolves a wire spelling back to its variant.
//
//	m, ok := types.ParseMaterialRemoval("required") // MaterialRemovalRequired, true
func ParseMaterialRemoval(s string) (MaterialRemoval, bool) {
	return enumFromName(materialRemovalNames, s)
}
