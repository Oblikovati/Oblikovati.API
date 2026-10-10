// SPDX-License-Identifier: Apache-2.0

package types

// Drawing sketch value types (M14-F08, #638). A drawing sketch is 2D geometry drawn directly in
// sheet space (millimetres) — annotation linework, detail enrichment, or a boundary to hatch. These
// are the canonical Apache-2.0 definitions; the GPL model renders the entities as drawing curves.

// DrawingSketchEntityKind classifies a drawing-sketch entity. The zero value is SketchLineEntity.
type DrawingSketchEntityKind int32

const (
	// SketchLineEntity is a straight segment between two sheet points.
	SketchLineEntity DrawingSketchEntityKind = iota
	// SketchCircleEntity is a full circle from a centre point and a radius.
	SketchCircleEntity
	// SketchRectangleEntity is an axis-aligned rectangle from two opposite corner points.
	SketchRectangleEntity
)

var drawingSketchEntityKindNames = map[DrawingSketchEntityKind]string{
	SketchLineEntity:      "line",
	SketchCircleEntity:    "circle",
	SketchRectangleEntity: "rectangle",
}

// String returns the entity kind's wire spelling ("line", "circle", "rectangle").
func (k DrawingSketchEntityKind) String() string {
	return enumName(drawingSketchEntityKindNames, k, "enum(?)")
}

// ParseDrawingSketchEntityKind resolves a wire spelling back to its entity kind.
//
//	k, ok := types.ParseDrawingSketchEntityKind("circle") // SketchCircleEntity, true
func ParseDrawingSketchEntityKind(s string) (DrawingSketchEntityKind, bool) {
	return enumFromName(drawingSketchEntityKindNames, s)
}
