// SPDX-License-Identifier: Apache-2.0

package types

// B-rep operation vocabularies (M07-F05/F06/F07, Oblikovati/Oblikovati#628,
// #629, #630). All numeric blocks are FROZEN to the reference enums
// (BooleanTypeEnum / OffsetCornerClosureTypeEnum / EdgeCollectionEnum) —
// clients and saved automations depend on them; never renumber.

// BooleanType selects the transient B-rep boolean operation
// (TransientBRep.DoBoolean).
type BooleanType int32

const (
	BooleanDifference BooleanType = 74241
	BooleanUnion      BooleanType = 74242
	BooleanIntersect  BooleanType = 74243
)

var booleanTypeNames = map[BooleanType]string{
	BooleanDifference: "difference",
	BooleanUnion:      "union",
	BooleanIntersect:  "intersect",
}

// String returns the operation's wire spelling.
func (b BooleanType) String() string { return enumName(booleanTypeNames, b) }

// ParseBooleanType resolves a wire spelling back to its operation.
func ParseBooleanType(s string) (BooleanType, bool) { return enumFromName(booleanTypeNames, s) }

// OffsetCornerClosureType selects how a planar wire offset closes a gap
// corner (Wire.OffsetPlanarWire).
type OffsetCornerClosureType int32

const (
	// CircularCornerClosure rounds the gap with an arc about the corner.
	CircularCornerClosure OffsetCornerClosureType = 96257
	// LinearCornerClosure extends both sides tangentially to a miter.
	LinearCornerClosure OffsetCornerClosureType = 96258
	// ExtendCornerClosure grows the actual curves (an arc stays an arc).
	ExtendCornerClosure OffsetCornerClosureType = 96259
)

var offsetCornerClosureNames = map[OffsetCornerClosureType]string{
	CircularCornerClosure: "circular",
	LinearCornerClosure:   "linear",
	ExtendCornerClosure:   "extend",
}

// String returns the closure's wire spelling.
func (c OffsetCornerClosureType) String() string { return enumName(offsetCornerClosureNames, c) }

// ParseOffsetCornerClosureType resolves a wire spelling back to its closure.
func ParseOffsetCornerClosureType(s string) (OffsetCornerClosureType, bool) {
	return enumFromName(offsetCornerClosureNames, s)
}

// EdgeCollectionKind identifies a body's classified edge collection
// (SurfaceBody.ConvexEdges / ConcaveEdges and tangentially connected sets).
type EdgeCollectionKind int32

const (
	TangentiallyConnectedEdges EdgeCollectionKind = 27649
	AllConcaveEdges            EdgeCollectionKind = 27650
	AllConvexEdges             EdgeCollectionKind = 27651
	UndefinedEdgeCollection    EdgeCollectionKind = 27652
)

var edgeCollectionKindNames = map[EdgeCollectionKind]string{
	TangentiallyConnectedEdges: "tangentiallyConnected",
	AllConcaveEdges:            "allConcave",
	AllConvexEdges:             "allConvex",
	UndefinedEdgeCollection:    "undefined",
}

// String returns the collection's wire spelling.
func (k EdgeCollectionKind) String() string { return enumName(edgeCollectionKindNames, k) }

// ParseEdgeCollectionKind resolves a wire spelling back to its collection.
func ParseEdgeCollectionKind(s string) (EdgeCollectionKind, bool) {
	return enumFromName(edgeCollectionKindNames, s)
}
