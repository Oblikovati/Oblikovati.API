// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Drawing dimensions (M14-F03 PBI-141, #388): standalone linear dimensions placed on a view by
// two pick points (each snapped to the nearest projected model vertex), measuring the true model
// distance and updating with the model.

// DrawingDimensionInfo is the JSON shape of one drawing dimension.
type DrawingDimensionInfo struct {
	Name       string  `json:"name"`
	Type       string  `json:"type"` // types.DrawingDimensionType: aligned|horizontal|vertical|radius|diameter|angular|ordinate|arcLength
	ViewName   string  `json:"viewName"`
	ValueMM    float64 `json:"valueMm"`            // measured model distance (mm), scale-independent; 0 for angular
	ValueDeg   float64 `json:"valueDeg,omitempty"` // measured angle (degrees) for an angular dimension
	Text       string  `json:"text"`               // displayed dimension text (with the overrides applied)
	CurveCount int     `json:"curveCount"`
	// Text overrides (#1992/#1993). Prefix/Suffix wrap the value; OverrideText replaces the whole
	// label; HideValue drops the value; DualUnit appends the inch value in brackets.
	Prefix       string `json:"prefix,omitempty"`
	Suffix       string `json:"suffix,omitempty"`
	OverrideText string `json:"overrideText,omitempty"`
	HideValue    bool   `json:"hideValue,omitempty"`
	DualUnit     bool   `json:"dualUnit,omitempty"`
	// Tolerance is the dimension's engineering tolerance (#1990); nil ⇒ none.
	Tolerance *types.DimensionTolerance `json:"tolerance,omitempty"`
	// Inspection is the dimension's inspection annotation (#1996); nil ⇒ not an inspection dimension.
	Inspection *types.InspectionDimension `json:"inspection,omitempty"`
	// Retrieved reports whether the dimension was retrieved from a model (parametric) dimension, and
	// RetrievedFrom names that source parameter — the model↔drawing association (#1991).
	Retrieved     bool   `json:"retrieved,omitempty"`
	RetrievedFrom string `json:"retrievedFrom,omitempty"`
}

// RetrievableDimensionInfo is one candidate model dimension a view can retrieve: its source parameter
// name, current value (mm) and the sheet position of its midpoint (#1991).
type RetrievableDimensionInfo struct {
	Name    string  `json:"name"`
	ValueMM float64 `json:"valueMm"`
	SheetX  float64 `json:"sheetX"`
	SheetY  float64 `json:"sheetY"`
}

// ListRetrievableDimensionsArgs is the request of [MethodDrawingDimensionsListRetrievable]: the base
// view whose referenced model's parametric dimensions to list (#1991).
type ListRetrievableDimensionsArgs struct {
	ViewName string `json:"viewName"`
}

// RetrievableDimensionsResult is the reply of [MethodDrawingDimensionsListRetrievable]: the model's
// retrievable dimensions projected onto the view.
type RetrievableDimensionsResult struct {
	Dimensions []RetrievableDimensionInfo `json:"dimensions"`
}

// RetrieveDimensionsArgs is the request of [MethodDrawingDimensionsRetrieve]: materialise the named
// model dimensions on the base view ViewName as retrieved drawing dimensions; an empty Names retrieves
// every model dimension. OffsetMM stands the dimension lines off the geometry (#1991).
type RetrieveDimensionsArgs struct {
	ViewName string   `json:"viewName"`
	Names    []string `json:"names,omitempty"`
	OffsetMM float64  `json:"offsetMm,omitempty"`
}

// RetrievedDimensionsResult is the reply of [MethodDrawingDimensionsRetrieve]: the drawing dimensions
// created, each flagged Retrieved with its RetrievedFrom back-reference.
type RetrievedDimensionsResult struct {
	Dimensions []DrawingDimensionInfo `json:"dimensions"`
}

// SetDimensionToleranceArgs is the request of [MethodDrawingDimensionsSetTolerance]: set the named
// dimension's engineering tolerance (#1990). A zero-value (none) Tolerance clears it.
type SetDimensionToleranceArgs struct {
	Name      string                   `json:"name"`
	Tolerance types.DimensionTolerance `json:"tolerance"`
}

// SetDimensionInspectionArgs is the request of [MethodDrawingDimensionsSetInspection]: flag the
// named dimension as an inspection dimension with a border shape, label and sampling rate (#1996).
// A NoInspectionBorder shape clears the inspection annotation.
type SetDimensionInspectionArgs struct {
	Name       string                    `json:"name"`
	Inspection types.InspectionDimension `json:"inspection"`
}

// SetDimensionTextStyleArgs is the request of [MethodDrawingDimensionsSetTextStyle]: change any
// subset of the named dimension's text overrides (#1992/#1993). Each pointer field is applied only
// when present, so one call can set a prefix without clearing an override.
type SetDimensionTextStyleArgs struct {
	Name         string  `json:"name"`
	Prefix       *string `json:"prefix,omitempty"`
	Suffix       *string `json:"suffix,omitempty"`
	OverrideText *string `json:"overrideText,omitempty"`
	HideValue    *bool   `json:"hideValue,omitempty"`
	DualUnit     *bool   `json:"dualUnit,omitempty"`
}

// ListDrawingDimensionsResult is the response of [MethodDrawingDimensionsList].
type ListDrawingDimensionsResult struct {
	Dimensions []DrawingDimensionInfo `json:"dimensions"`
}

// AddLinearDimensionArgs is the request of [MethodDrawingDimensionsAddLinear]: a linear dimension
// on ViewName between two pick points (sheet millimetres). Each pick is snapped to the nearest
// projected model vertex, so the dimension re-measures when the model changes. Type selects the
// measured component (aligned = true distance, default; horizontal/vertical = the view-X/Y
// component). OffsetMM stands the dimension line off the measured points (signed, sheet mm).
type AddLinearDimensionArgs struct {
	Name     string  `json:"name,omitempty"`
	ViewName string  `json:"viewName"`
	Type     string  `json:"type,omitempty"`
	X1       float64 `json:"x1"`
	Y1       float64 `json:"y1"`
	X2       float64 `json:"x2"`
	Y2       float64 `json:"y2"`
	OffsetMM float64 `json:"offsetMm,omitempty"`
}

// AddRadialDimensionArgs is the request of [MethodDrawingDimensionsAddRadial]: a radius or
// diameter dimension on ViewName, attached to the circular model edge nearest the pick point
// (sheet mm). Type is "radius" or "diameter". The dimension re-measures when the model changes.
type AddRadialDimensionArgs struct {
	Name     string  `json:"name,omitempty"`
	ViewName string  `json:"viewName"`
	Type     string  `json:"type,omitempty"` // radius (default) | diameter
	PickXMM  float64 `json:"pickXmm"`
	PickYMM  float64 `json:"pickYmm"`
}

// AddAngularDimensionArgs is the request of [MethodDrawingDimensionsAddAngular]: an angular
// dimension on ViewName between the two straight model edges nearest the pick points (sheet mm).
// The measured angle re-derives when the model changes.
type AddAngularDimensionArgs struct {
	Name     string  `json:"name,omitempty"`
	ViewName string  `json:"viewName"`
	X1       float64 `json:"x1"`
	Y1       float64 `json:"y1"`
	X2       float64 `json:"x2"`
	Y2       float64 `json:"y2"`
}

// AddDimensionSetArgs is the request of [MethodDrawingDimensionsAddBaseline] /
// [MethodDrawingDimensionsAddChain]: a set of linear dimensions on ViewName from a list of pick
// points (each [x,y] sheet mm, snapped to the nearest projected model vertex). A baseline set
// measures from the first point to each of the others (stacked); a chain set measures between
// consecutive points (in a line). Type selects the measured component (aligned/horizontal/vertical).
type AddDimensionSetArgs struct {
	ViewName string      `json:"viewName"`
	Type     string      `json:"type,omitempty"`
	Points   [][]float64 `json:"points"`
}

// DimensionSetResult is the response of the dimension-set methods: the created dimensions.
type DimensionSetResult struct {
	Dimensions []DrawingDimensionInfo `json:"dimensions"`
}

// AddOrdinateDimensionsArgs is the request of [MethodDrawingDimensionsAddOrdinate]: an ordinate
// dimension for each point in Points, each measuring that point's offset from the common Datum
// ([x,y] sheet mm) along Axis. Datum and every point are snapped to the nearest projected model
// vertex, so the values stay associative. Axis is "horizontal" (the view-X offset) or "vertical"
// (the view-Y offset). Each ordinate is drawn as a leader to its value with no dimension line.
type AddOrdinateDimensionsArgs struct {
	ViewName string      `json:"viewName"`
	Axis     string      `json:"axis,omitempty"` // horizontal (default) | vertical
	Datum    []float64   `json:"datum"`          // [x,y] sheet mm — the common origin
	Points   [][]float64 `json:"points"`         // each [x,y] sheet mm
}

// AddArcLengthDimensionArgs is the request of [MethodDrawingDimensionsAddArcLength]: an arc-length
// dimension on ViewName, attached to the circular/arc model edge nearest the pick point (sheet mm).
// It measures the edge's swept length (a full circle's circumference) with the dimension line
// following the arc; the value re-measures when the model changes.
type AddArcLengthDimensionArgs struct {
	Name     string  `json:"name,omitempty"`
	ViewName string  `json:"viewName"`
	PickXMM  float64 `json:"pickXmm"`
	PickYMM  float64 `json:"pickYmm"`
}

// DeleteDimensionArgs is the request of [MethodDrawingDimensionsDelete].
type DeleteDimensionArgs struct {
	Name string `json:"name"`
}

// DimensionResult is the response of [MethodDrawingDimensionsAddLinear]: the created dimension.
type DimensionResult struct {
	Dimension DrawingDimensionInfo `json:"dimension"`
}
