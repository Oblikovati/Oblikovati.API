// SPDX-License-Identifier: Apache-2.0

package types

// DimensionDisplayType controls how dimensions tied to parameters render in
// sketches: the evaluated value, the parameter name, the authored expression,
// the value with its tolerance, or the precise (unrounded) value (parity:
// DimensionDisplayTypeEnum, M02-F07, Oblikovati/Oblikovati#606). The numeric
// values are frozen at the reference API's ids and must never be renumbered.
type DimensionDisplayType int32

const (
	DimensionDisplayValue        DimensionDisplayType = 34817
	DimensionDisplayName         DimensionDisplayType = 34818
	DimensionDisplayExpression   DimensionDisplayType = 34819
	DimensionDisplayTolerance    DimensionDisplayType = 34820
	DimensionDisplayPreciseValue DimensionDisplayType = 34821
)

// dimensionDisplayNames are the wire spellings
// (wire.ParameterSettingsInfo.DimensionDisplayType).
var dimensionDisplayNames = map[DimensionDisplayType]string{
	DimensionDisplayValue:        "value",
	DimensionDisplayName:         "name",
	DimensionDisplayExpression:   "expression",
	DimensionDisplayTolerance:    "tolerance",
	DimensionDisplayPreciseValue: "preciseValue",
}

// String returns the display type's wire spelling.
func (d DimensionDisplayType) String() string { return enumName(dimensionDisplayNames, d, "enum(?)") }

// ParseDimensionDisplayType resolves a wire spelling back to its
// DimensionDisplayType.
func ParseDimensionDisplayType(s string) (DimensionDisplayType, bool) {
	return enumFromName(dimensionDisplayNames, s)
}
