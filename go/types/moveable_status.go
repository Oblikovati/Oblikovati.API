// SPDX-License-Identifier: Apache-2.0

package types

// GeometryMoveableStatus answers "can this sketch entity be dragged?" for
// interactive tools — derived from the solver's per-entity constraint
// analysis (M06-F11, Oblikovati/Oblikovati#626).
//
// The values are a frozen block matching the reference API's
// geometry-moveable-status enum; never renumber them.
type GeometryMoveableStatus int32

const (
	// MoveableFree geometry has free degrees of freedom and drags directly.
	MoveableFree GeometryMoveableStatus = 53505
	// MoveableByDimensionChange geometry only moves if a driving dimension
	// is relaxed — dragging it would re-solve other geometry.
	MoveableByDimensionChange GeometryMoveableStatus = 53506
	// MoveableFixed geometry is grounded/fixed (or reference) and never drags.
	MoveableFixed GeometryMoveableStatus = 53507
	// MoveableUnknown is reported when the solver cannot classify the entity.
	MoveableUnknown GeometryMoveableStatus = 53508
)

// geometryMoveableStatusNames are the frozen wire spellings.
var geometryMoveableStatusNames = map[GeometryMoveableStatus]string{
	MoveableFree:              "freeToMove",
	MoveableByDimensionChange: "byDimensionChange",
	MoveableFixed:             "fixed",
	MoveableUnknown:           "unknown",
}

// String returns the status's wire spelling.
func (m GeometryMoveableStatus) String() string {
	return enumName(geometryMoveableStatusNames, m, "enum(?)")
}

// ParseGeometryMoveableStatus resolves a wire spelling back to its status.
func ParseGeometryMoveableStatus(s string) (GeometryMoveableStatus, bool) {
	return enumFromName(geometryMoveableStatusNames, s)
}
