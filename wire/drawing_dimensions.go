// SPDX-License-Identifier: Apache-2.0

package wire

// Drawing dimensions (M14-F03 PBI-141, #388): standalone linear dimensions placed on a view by
// two pick points (each snapped to the nearest projected model vertex), measuring the true model
// distance and updating with the model.

// DrawingDimensionInfo is the JSON shape of one drawing dimension.
type DrawingDimensionInfo struct {
	Name       string  `json:"name"`
	Type       string  `json:"type"` // types.DrawingDimensionType: aligned|horizontal|vertical
	ViewName   string  `json:"viewName"`
	ValueMM    float64 `json:"valueMm"` // measured model distance (mm), scale-independent
	Text       string  `json:"text"`    // displayed dimension text
	CurveCount int     `json:"curveCount"`
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

// DeleteDimensionArgs is the request of [MethodDrawingDimensionsDelete].
type DeleteDimensionArgs struct {
	Name string `json:"name"`
}

// DimensionResult is the response of [MethodDrawingDimensionsAddLinear]: the created dimension.
type DimensionResult struct {
	Dimension DrawingDimensionInfo `json:"dimension"`
}
