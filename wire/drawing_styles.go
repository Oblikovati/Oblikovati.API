// SPDX-License-Identifier: Apache-2.0

package wire

// Drawing style wire DTOs (M14-F01 PBI-138, Oblikovati/Oblikovati#385). The active drawing
// follows a drafting standard (its wire spelling, e.g. "iso") whose preset fixes the
// dimension/text/line appearance. Switching the standard returns the new active preset, so
// callers see the appearance change in one round-trip.

// DimensionStyleInfo is the JSON shape of a dimension style.
type DimensionStyleInfo struct {
	Name          string  `json:"name"`
	TextHeightMM  float64 `json:"textHeightMm"`
	ArrowSizeMM   float64 `json:"arrowSizeMm"`
	DecimalPlaces int     `json:"decimalPlaces"`
	Unit          string  `json:"unit"` // types.DimensionUnit spelling ("mm"/"in")
	LineWeightMM  float64 `json:"lineWeightMm"`
}

// TextStyleInfo is the JSON shape of an annotation text style.
type TextStyleInfo struct {
	Name     string  `json:"name"`
	FontName string  `json:"fontName"`
	HeightMM float64 `json:"heightMm"`
}

// LineStyleInfo is the JSON shape of a drawing line style.
type LineStyleInfo struct {
	Name     string  `json:"name"`
	WeightMM float64 `json:"weightMm"`
}

// StandardStyleInfo is one drafting standard's complete style preset.
type StandardStyleInfo struct {
	Standard  string             `json:"standard"` // types.DraftingStandard spelling
	Dimension DimensionStyleInfo `json:"dimension"`
	Text      TextStyleInfo      `json:"text"`
	Line      LineStyleInfo      `json:"line"`
}

// ListStandardsResult is the response of [MethodDrawingStylesListStandards]: the available
// drafting standards and the active one.
type ListStandardsResult struct {
	Standards []string `json:"standards"`
	Active    string   `json:"active"`
}

// SetStandardArgs is the request of [MethodDrawingStylesSetStandard]: the drafting standard
// to make active (a types.DraftingStandard spelling, "iso"/"ansi").
type SetStandardArgs struct {
	Standard string `json:"standard"`
}

// StandardStyleResult is the response of [MethodDrawingStylesGetActiveStyle] /
// [MethodDrawingStylesSetStandard]: the active standard's resolved style preset.
type StandardStyleResult struct {
	Style StandardStyleInfo `json:"style"`
}
