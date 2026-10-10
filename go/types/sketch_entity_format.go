// SPDX-License-Identifier: Apache-2.0

package types

// SketchEntityFormat is one sketch entity's formatting overrides — the line type, colour and
// stroke width the sketch Format panel's three lists set, and the values a DWG or DXF import
// carries in from the file's layer table (Oblikovati/Oblikovati#2015).
//
// Each field independently means "inherit" when unset, so an entity can override its colour while
// taking the sketch's line type. An entity with no overrides at all has no format: reading one
// back reports that, rather than returning a struct of empty fields that might be mistaken for
// explicit defaults.
//
//	f := types.SketchEntityFormat{LineType: types.SketchLineDashed}
//	f.OverrideColor = types.NewColor(255, 0, 0)
type SketchEntityFormat struct {
	// LineType is the dash pattern; the empty value inherits the sketch's line type.
	LineType SketchLineType `json:"lineType,omitempty"`
	// OverrideColor is the entity's colour. A colour whose Source is not OverrideColorSource
	// inherits instead — see [Color.IsOverride].
	OverrideColor Color `json:"overrideColor,omitempty"`
	// LineWeight is the plotted stroke width in millimetres; 0 inherits.
	LineWeight float64 `json:"lineWeight,omitempty"`
}

// IsDefault reports whether the format overrides nothing, in which case the entity draws with the
// sketch's own attributes.
func (f SketchEntityFormat) IsDefault() bool {
	return f.LineType == "" && f.LineWeight == 0 && !f.OverrideColor.IsOverride()
}

// SketchFormatModes is the Format panel's armed creation state: what newly drawn geometry becomes
// (Oblikovati/Oblikovati#2015).
//
// Each mode is what the matching panel button arms when it is pressed with nothing selected. With
// geometry selected the same button converts the selection instead, which is an action rather
// than a setting and so is not represented here.
type SketchFormatModes struct {
	// Construction makes new geometry construction geometry, excluded from profiles.
	Construction bool `json:"construction"`
	// Centerline makes new lines centerlines — an axis for revolve, mirror and symmetry.
	Centerline bool `json:"centerline"`
	// CenterPoint makes the point tool place hole-centre markers rather than plain points.
	CenterPoint bool `json:"centerPoint"`
	// DrivenDimension makes new dimensions driven by the geometry rather than driving it.
	DrivenDimension bool `json:"drivenDimension"`
	// SuppressFormatOverrides draws the sketch with DEFAULT attributes, hiding per-entity
	// formatting. It is the Show Format button, whose behaviour is the inverse of its label:
	// on shows the default format, off shows the user's.
	SuppressFormatOverrides bool `json:"suppressFormatOverrides"`
}
