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
	// MeasureAngle is the angle in degrees between two entities (an edge's direction or a planar
	// face's normal), or — with a third vertex — the angle at the apex of three vertices.
	MeasureAngle
	// MeasureLoopLength is the length of a face's outer boundary loop — its perimeter.
	MeasureLoopLength
)

var measureTypeNames = map[MeasureType]string{
	MeasureLength:      "length",
	MeasureArea:        "area",
	MeasureDistance:    "distance",
	MeasureMinDistance: "minDistance",
	MeasureAngle:       "angle",
	MeasureLoopLength:  "loopLength",
}

// String returns the measure type's wire spelling ("length", "area", "distance", "minDistance",
// "angle", "loopLength").
func (m MeasureType) String() string { return enumName(measureTypeNames, m) }

// ParseMeasureType resolves a wire spelling back to its measure type.
func ParseMeasureType(s string) (MeasureType, bool) {
	return enumFromName(measureTypeNames, s)
}
