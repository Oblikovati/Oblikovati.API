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

// AddHatchRegionArgs is the request of [MethodDrawingSketchesAddHatch]: fill the rectangle at
// (XMM, YMM) of size WidthMM×HeightMM (sheet millimetres) with a hatch pattern. Pattern is the
// built-in pattern name ("general", "cross", "ansi31"); ScaleMm overrides the line spacing (0 ⇒ the
// pattern default). The region is added to the named sketch (created if SketchName is blank).
type AddHatchRegionArgs struct {
	SketchName string  `json:"sketchName,omitempty"`
	XMM        float64 `json:"xmm"`
	YMM        float64 `json:"ymm"`
	WidthMM    float64 `json:"widthMm"`
	HeightMM   float64 `json:"heightMm"`
	Pattern    string  `json:"pattern,omitempty"`
	ScaleMM    float64 `json:"scaleMm,omitempty"`
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
