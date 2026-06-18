// SPDX-License-Identifier: Apache-2.0

package types

// Measurement value types (M18-F01 PBI-164, #428). A measure reports a geometric quantity of one or
// two model entities. These are the canonical Apache-2.0 definitions; the GPL model computes the
// values from geometry.

// MeasureType is the quantity a measurement reports. The zero value is MeasureLength.
type MeasureType int32

const (
	// MeasureLength is the length of a single edge.
	MeasureLength MeasureType = iota
	// MeasureArea is the area of a single face.
	MeasureArea
	// MeasureDistance is the straight-line distance between two vertices.
	MeasureDistance
	// MeasureMinDistance is the minimum distance between two entities of any kind
	// (vertex, edge or face) — their closest approach. Zero when they touch or intersect.
	MeasureMinDistance
)

var measureTypeNames = map[MeasureType]string{
	MeasureLength:      "length",
	MeasureArea:        "area",
	MeasureDistance:    "distance",
	MeasureMinDistance: "minDistance",
}

// String returns the measure type's wire spelling ("length", "area", "distance", "minDistance").
func (m MeasureType) String() string { return enumName(measureTypeNames, m) }

// ParseMeasureType resolves a wire spelling back to its measure type.
func ParseMeasureType(s string) (MeasureType, bool) {
	return enumFromName(measureTypeNames, s)
}
