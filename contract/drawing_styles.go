// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The drawing style system (M14-F01 PBI-138, Oblikovati/Oblikovati#385): a drawing follows a
// drafting standard whose style preset fixes the appearance of dimensions, text and lines.
// An in-proc consumer reads the active standard and its styles directly; switching the
// standard travels over api/wire (drawingStyles.*). The host owns the per-standard presets
// (/source, model/drawing).

// DrawingDimensionStyle is the appearance of a dimension under the active standard: the
// text and arrow sizes, the displayed precision, the measurement unit, and the line weight.
type DrawingDimensionStyle interface {
	// Name is the style's display name (e.g. "Default (ISO)").
	Name() string
	// TextHeightMM is the dimension text height in millimetres.
	TextHeightMM() float64
	// ArrowSizeMM is the dimension arrowhead size in millimetres.
	ArrowSizeMM() float64
	// DecimalPlaces is the number of fractional digits shown.
	DecimalPlaces() int
	// Unit is the measurement unit the value is displayed in.
	Unit() types.DimensionUnit
	// LineWeightMM is the dimension/extension line weight in millimetres.
	LineWeightMM() float64
}

// DrawingTextStyle is the appearance of annotation text under the active standard.
type DrawingTextStyle interface {
	// Name is the style's display name.
	Name() string
	// FontName is the text font family.
	FontName() string
	// HeightMM is the text height in millimetres.
	HeightMM() float64
}

// DrawingLineStyle is the appearance of a drawing line under the active standard.
type DrawingLineStyle interface {
	// Name is the style's display name.
	Name() string
	// WeightMM is the line weight in millimetres.
	WeightMM() float64
}

// DrawingStandardStyle is one drafting standard's complete style preset.
type DrawingStandardStyle interface {
	// Standard is the drafting standard this preset implements.
	Standard() types.DraftingStandard
	// DimensionStyle, TextStyle and LineStyle are the preset's component styles.
	DimensionStyle() DrawingDimensionStyle
	TextStyle() DrawingTextStyle
	LineStyle() DrawingLineStyle
}

// DrawingStylesManager is a drawing's style system: the active drafting standard and its
// resolved style preset. Switching the standard re-points the active preset, so every
// annotation re-renders to it.
type DrawingStylesManager interface {
	// ActiveStandard is the drawing's current drafting standard.
	ActiveStandard() types.DraftingStandard
	// ActiveStyle is the style preset of the active standard.
	ActiveStyle() DrawingStandardStyle
}
