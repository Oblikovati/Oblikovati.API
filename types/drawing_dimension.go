// SPDX-License-Identifier: Apache-2.0

package types

// Drawing dimension value types (M14-F03 PBI-141, #388). A drawing dimension measures and
// annotates a distance on a drawing view. It is associative: its two attachment points re-bind
// to model geometry, so the measured value tracks the model. These are the canonical Apache-2.0
// definitions; the GPL model implements the measurement and the dimension geometry.

// DrawingDimensionType names how a linear dimension is measured between its two attachment
// points. The zero value is AlignedDimension (the true point-to-point distance).
type DrawingDimensionType int32

const (
	// AlignedDimension measures the straight distance between the two points.
	AlignedDimension DrawingDimensionType = iota
	// HorizontalDimension measures the horizontal (view-X) component of the distance.
	HorizontalDimension
	// VerticalDimension measures the vertical (view-Y) component of the distance.
	VerticalDimension
	// RadiusDimension measures the radius of a circular edge (annotated "R<value>").
	RadiusDimension
	// DiameterDimension measures the diameter of a circular edge (annotated "⌀<value>").
	DiameterDimension
	// AngularDimension measures the angle between two straight edges, in degrees.
	AngularDimension
	// OrdinateDimension measures a point's view-X or view-Y offset from a common datum, shown as a
	// leader to the value with no dimension line (the running-coordinate callout). Which axis it
	// measures is chosen when the dimension is created (the "axis" request field).
	OrdinateDimension
	// ArcLengthDimension measures the length along a circular edge — the arc's swept length, or a
	// full circle's circumference — with the dimension line following the arc.
	ArcLengthDimension
	// ForeshortenedDimension is a radius dimension whose centre is off-sheet, drawn with a jogged
	// (broken) dimension line so the value still reads (Inventor's foreshortened radius, #1994).
	ForeshortenedDimension
	// SymmetricDimension dimensions one side of a feature symmetric about a centre line, labelling
	// the full size from the half measurement (#1994).
	SymmetricDimension
	// SumDimension shows the running sum of a chain of dimensions from a common origin (#1994).
	SumDimension
)

var drawingDimensionTypeNames = map[DrawingDimensionType]string{
	AlignedDimension:       "aligned",
	HorizontalDimension:    "horizontal",
	VerticalDimension:      "vertical",
	RadiusDimension:        "radius",
	DiameterDimension:      "diameter",
	AngularDimension:       "angular",
	OrdinateDimension:      "ordinate",
	ArcLengthDimension:     "arcLength",
	ForeshortenedDimension: "foreshortened",
	SymmetricDimension:     "symmetric",
	SumDimension:           "sum",
}

// String returns the dimension type's wire spelling ("aligned", "horizontal", "vertical").
func (t DrawingDimensionType) String() string {
	return enumName(drawingDimensionTypeNames, t, "enum(?)")
}

// ParseDrawingDimensionType resolves a wire spelling back to its dimension type.
//
//	t, ok := types.ParseDrawingDimensionType("horizontal") // HorizontalDimension, true
func ParseDrawingDimensionType(s string) (DrawingDimensionType, bool) {
	return enumFromName(drawingDimensionTypeNames, s)
}

// DimensionToleranceType selects how a dimension's engineering tolerance is shown — none, a
// symmetric ±, an asymmetric deviation (+plus/−minus), stacked max/min limits, or an ISO
// limits-and-fits class such as H7 (Inventor's tolerance methods, #1990).
type DimensionToleranceType int32

const (
	// NoTolerance shows the nominal value alone (the default).
	NoTolerance DimensionToleranceType = iota
	// SymmetricTolerance shows a single ± deviation.
	SymmetricTolerance
	// DeviationTolerance shows an upper (+) and lower (−) deviation.
	DeviationTolerance
	// LimitsTolerance shows the max and min sizes stacked (nominal + deviations resolved).
	LimitsTolerance
	// FitsTolerance shows an ISO limits-and-fits class (e.g. "H7") after the value.
	FitsTolerance
)

var dimensionToleranceTypeNames = map[DimensionToleranceType]string{
	NoTolerance:        "none",
	SymmetricTolerance: "symmetric",
	DeviationTolerance: "deviation",
	LimitsTolerance:    "limits",
	FitsTolerance:      "fits",
}

// String returns the tolerance type's wire spelling.
func (t DimensionToleranceType) String() string {
	return enumName(dimensionToleranceTypeNames, t, "enum(?)")
}

// ParseDimensionToleranceType resolves a wire spelling back to its tolerance type; "" ⇒ none.
func ParseDimensionToleranceType(s string) (DimensionToleranceType, bool) {
	if s == "" {
		return NoTolerance, true
	}
	return enumFromName(dimensionToleranceTypeNames, s)
}

// DimensionTolerance is one dimension's engineering tolerance (#1990): its method, the upper (Plus) and
// lower (Minus) deviations in millimetres, the ISO fit class for the fits method, and the decimal
// precision the tolerance values render at.
type DimensionTolerance struct {
	Type      DimensionToleranceType `json:"type"`
	Plus      float64                `json:"plus,omitempty"`
	Minus     float64                `json:"minus,omitempty"`
	Fit       string                 `json:"fit,omitempty"`
	Precision int                    `json:"precision,omitempty"`
}

// InspectionShape selects the border an inspection dimension wraps its text in, matching
// Inventor's InspectionDimensionShapeEnum. The zero value is NoInspectionBorder — the dimension
// is not an inspection dimension. An inspection dimension additionally carries a label and a
// sampling rate for QA / first-article drawings (#1996).
type InspectionShape int32

const (
	// NoInspectionBorder is a plain dimension (not an inspection dimension) — the default.
	NoInspectionBorder InspectionShape = iota
	// AngularEndsInspectionBorder wraps the text in a border with angular (chevron) ends.
	AngularEndsInspectionBorder
	// RoundedEndsInspectionBorder wraps the text in a border with rounded (stadium) ends.
	RoundedEndsInspectionBorder
)

var inspectionShapeNames = map[InspectionShape]string{
	NoInspectionBorder:          "none",
	AngularEndsInspectionBorder: "angular",
	RoundedEndsInspectionBorder: "rounded",
}

// String returns the inspection shape's wire spelling ("none", "angular", "rounded").
func (s InspectionShape) String() string { return enumName(inspectionShapeNames, s, "enum(?)") }

// ParseInspectionShape resolves a wire spelling back to its inspection shape; "" ⇒ none.
func ParseInspectionShape(s string) (InspectionShape, bool) {
	if s == "" {
		return NoInspectionBorder, true
	}
	return enumFromName(inspectionShapeNames, s)
}

// InspectionDimension is a dimension's inspection annotation (#1996): the border shape, and the
// QA label and sampling rate shown with it. A NoInspectionBorder shape means the dimension is not
// an inspection dimension (Label and Rate are then ignored).
type InspectionDimension struct {
	Shape InspectionShape `json:"shape"`
	Label string          `json:"label,omitempty"`
	Rate  string          `json:"rate,omitempty"`
}
