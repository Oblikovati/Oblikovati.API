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

// FlatPatternEdgeType classifies an edge of the developed flat for manufacturing and drawing:
// a fold line where the sheet bends up or down, or a tangent line where a curved wall met the
// flat tangentially. Outer-profile (cut) edges carry no flat-pattern type.
type FlatPatternEdgeType int32

const (
	// BendUpFlatPatternEdge is a fold line whose bend lifts the material toward the front
	// (top) face when the flat is refolded.
	BendUpFlatPatternEdge FlatPatternEdgeType = iota
	// BendDownFlatPatternEdge is a fold line whose bend lifts the material toward the back.
	BendDownFlatPatternEdge
	// TangentFlatPatternEdge is the tangent line where a rolled/curved wall meets the flat.
	TangentFlatPatternEdge
)

var flatPatternEdgeTypeNames = map[FlatPatternEdgeType]string{
	BendUpFlatPatternEdge:   "bendUp",
	BendDownFlatPatternEdge: "bendDown",
	TangentFlatPatternEdge:  "tangent",
}

// String returns the edge type's wire spelling.
func (e FlatPatternEdgeType) String() string { return enumName(flatPatternEdgeTypeNames, e) }

// ParseFlatPatternEdgeType resolves a wire spelling back to its edge type.
func ParseFlatPatternEdgeType(s string) (FlatPatternEdgeType, bool) {
	return enumFromName(flatPatternEdgeTypeNames, s)
}

// FlatPatternFaceType classifies a face of the developed flat: the front (top) face the
// orientation presents, the back (bottom) face, a detail face, or unknown.
type FlatPatternFaceType int32

const (
	// UnknownFlatPatternFace is a face the flat pattern does not classify.
	UnknownFlatPatternFace FlatPatternFaceType = iota
	// FrontFlatPatternFace is the top face the active orientation presents.
	FrontFlatPatternFace
	// BackFlatPatternFace is the bottom face.
	BackFlatPatternFace
	// DetailFlatPatternFace is a face carrying punch/relief detail.
	DetailFlatPatternFace
)

var flatPatternFaceTypeNames = map[FlatPatternFaceType]string{
	UnknownFlatPatternFace: "unknown",
	FrontFlatPatternFace:   "front",
	BackFlatPatternFace:    "back",
	DetailFlatPatternFace:  "detail",
}

// String returns the face type's wire spelling.
func (f FlatPatternFaceType) String() string { return enumName(flatPatternFaceTypeNames, f) }

// ParseFlatPatternFaceType resolves a wire spelling back to its face type.
func ParseFlatPatternFaceType(s string) (FlatPatternFaceType, bool) {
	return enumFromName(flatPatternFaceTypeNames, s)
}
