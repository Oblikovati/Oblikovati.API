// SPDX-License-Identifier: Apache-2.0

package wire

// Drawing sketch DTOs (M14-F08 #638): 2D geometry drawn directly in sheet space (millimetres) on a
// sheet. A sketch holds entities (lines, circles, rectangles) that render as drawing curves.

// AddDrawingSketchArgs is the request of [MethodDrawingSketchesAdd]: a new empty sketch on the
// active sheet. A blank Name auto-names it.
type AddDrawingSketchArgs struct {
	Name string `json:"name,omitempty"`
}

// AddDrawingSketchEntityArgs is the request of [MethodDrawingSketchesAddEntity]: add one entity to
// the named sketch. Points are sheet-millimetre [x, y] pairs — two for a line (endpoints) or
// rectangle (opposite corners), one for a circle (centre, with Radius in millimetres).
type AddDrawingSketchEntityArgs struct {
	SketchName string       `json:"sketchName"`
	Kind       string       `json:"kind"`
	Points     [][2]float64 `json:"points"`
	Radius     float64      `json:"radiusMm,omitempty"`
}

// DrawingSketchInfo flattens a drawing sketch for the wire.
type DrawingSketchInfo struct {
	Name        string `json:"name"`
	EntityCount int    `json:"entityCount"`
	CurveCount  int    `json:"curveCount"`
}

// DrawingSketchResult is the response of the sketch add methods: the affected sketch.
type DrawingSketchResult struct {
	Sketch DrawingSketchInfo `json:"sketch"`
}

// ListDrawingSketchesResult is the response of [MethodDrawingSketchesList].
type ListDrawingSketchesResult struct {
	Sketches []DrawingSketchInfo `json:"sketches"`
}
