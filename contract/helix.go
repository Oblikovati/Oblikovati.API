// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// HelicalCurveDefinition is the scalar view of a 3D sketch helical curve's
// definition (M06-F09, Oblikovati/Oblikovati#624). The host's
// model/sketch.HelixDefinition satisfies this via a compile-time assertion.
type HelicalCurveDefinition interface {
	// ShapeKind is how the shape is specified (pitch/height/revolutions/spiral).
	ShapeKind() types.HelicalShapeDefinitionKind
	// Variable reports a row-table (variable-shape) definition.
	Variable() bool
	// RowCount is the number of shape stations (0 for a constant shape).
	RowCount() int
	// Clockwise is the winding handedness.
	Clockwise() bool
	// StartEnd and EndEnd are the end transition conditions.
	StartEnd() types.HelixEndKind
	EndEnd() types.HelixEndKind
}
