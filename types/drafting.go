// SPDX-License-Identifier: Apache-2.0

package types

// Drafting standard value types (M14-F01 PBI-138). A drawing document follows a
// drafting standard (ISO or ANSI) that governs the appearance of its dimensions, text
// and lines. Switching the standard re-points the active style preset, so every
// annotation re-renders to that standard. These are the canonical, Apache-2.0
// definitions; the GPL model aliases them (ADR-0018) and owns the per-standard presets.

// DraftingStandard names the drawing's drafting convention. The zero value is
// DraftingISO (metric, the default for a new drawing).
type DraftingStandard int32

const (
	// DraftingISO is the ISO (metric) drafting standard — millimetres, the default.
	DraftingISO DraftingStandard = iota
	// DraftingANSI is the ANSI/ASME (imperial) drafting standard — inches.
	DraftingANSI
)

var draftingStandardNames = map[DraftingStandard]string{
	DraftingISO:  "iso",
	DraftingANSI: "ansi",
}

// String returns the standard's wire spelling ("iso" / "ansi").
func (d DraftingStandard) String() string { return enumName(draftingStandardNames, d) }

// ParseDraftingStandard resolves a wire spelling back to its drafting standard.
//
//	s, ok := types.ParseDraftingStandard("ansi") // DraftingANSI, true
func ParseDraftingStandard(s string) (DraftingStandard, bool) {
	return enumFromName(draftingStandardNames, s)
}

// DimensionUnit names the unit a dimension is measured and displayed in. The zero
// value is DimensionMillimeter (the ISO default).
type DimensionUnit int32

const (
	// DimensionMillimeter measures dimensions in millimetres (ISO).
	DimensionMillimeter DimensionUnit = iota
	// DimensionInch measures dimensions in inches (ANSI).
	DimensionInch
)

var dimensionUnitNames = map[DimensionUnit]string{
	DimensionMillimeter: "mm",
	DimensionInch:       "in",
}

// String returns the unit's wire spelling ("mm" / "in").
func (u DimensionUnit) String() string { return enumName(dimensionUnitNames, u) }

// ParseDimensionUnit resolves a wire spelling back to its dimension unit.
func ParseDimensionUnit(s string) (DimensionUnit, bool) {
	return enumFromName(dimensionUnitNames, s)
}
