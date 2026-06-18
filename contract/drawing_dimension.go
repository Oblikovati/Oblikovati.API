// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// DrawingDimension is one dimension on a drawing sheet (M14-F03 PBI-141, #388): a measured,
// annotated distance between two attachment points on a view. It is associative — its value
// re-measures when the referenced model changes. The GPL model implements it; an add-in reads it
// through this surface.
type DrawingDimension interface {
	// Name is the dimension's name, unique within the sheet.
	Name() string
	// Type is how the distance is measured (aligned/horizontal/vertical).
	Type() types.DrawingDimensionType
	// ViewName is the drawing view the dimension is attached to.
	ViewName() string
	// ValueMM is the measured model distance in millimetres — the true size, independent of the
	// view scale. It is 0 for an angular dimension (see ValueDeg).
	ValueMM() float64
	// ValueDeg is the measured angle in degrees for an angular dimension, and 0 otherwise.
	ValueDeg() float64
	// Text is the displayed dimension text (the formatted value).
	Text() string
	// CurveCount is the number of drawing curves the dimension renders (its extension lines,
	// dimension line and arrowheads).
	CurveCount() int
}
