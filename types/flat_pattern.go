// SPDX-License-Identifier: Apache-2.0

package types

// Flat-pattern value types (M13-F05). The flat pattern is more than a developed body: it is a
// container with named orientations (saved alignment states that frame the flat for drawing
// views and export), per-region plates, and edge/face classification. These are the canonical
// Apache-2.0 enums; the GPL implementation aliases them (ADR-0018).

// AlignmentType names how a flat-pattern orientation aligns its reference axis to the page:
// the chosen edge/axis is laid horizontal or vertical, then rotated by the orientation's
// alignment angle.
type AlignmentType int32

const (
	// HorizontalAlignment lays the orientation's alignment axis along the horizontal (the
	// default), so the flat's reported length runs left-to-right.
	HorizontalAlignment AlignmentType = iota
	// VerticalAlignment lays the alignment axis along the vertical.
	VerticalAlignment
)

var alignmentTypeNames = map[AlignmentType]string{
	HorizontalAlignment: "horizontal",
	VerticalAlignment:   "vertical",
}

// String returns the alignment type's wire spelling.
func (a AlignmentType) String() string { return enumName(alignmentTypeNames, a) }

// ParseAlignmentType resolves a wire spelling back to its alignment type.
func ParseAlignmentType(s string) (AlignmentType, bool) {
	return enumFromName(alignmentTypeNames, s)
}
