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
)

var drawingDimensionTypeNames = map[DrawingDimensionType]string{
	AlignedDimension:    "aligned",
	HorizontalDimension: "horizontal",
	VerticalDimension:   "vertical",
	RadiusDimension:     "radius",
	DiameterDimension:   "diameter",
}

// String returns the dimension type's wire spelling ("aligned", "horizontal", "vertical").
func (t DrawingDimensionType) String() string { return enumName(drawingDimensionTypeNames, t) }

// ParseDrawingDimensionType resolves a wire spelling back to its dimension type.
//
//	t, ok := types.ParseDrawingDimensionType("horizontal") // HorizontalDimension, true
func ParseDrawingDimensionType(s string) (DrawingDimensionType, bool) {
	return enumFromName(drawingDimensionTypeNames, s)
}
