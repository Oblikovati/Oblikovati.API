// SPDX-License-Identifier: Apache-2.0

package wire

// AddDimensionArgs is the request of [MethodSketchAddDimension] — the discriminated
// dimensional-constraint constructor. Kind is a
// [oblikovati.org/api/types.DimensionConstraintKind]; Entities are the session ids
// of the geometry being dimensioned (in the kind's order); Expression is the unit-bearing
// value ("40 mm", "30 deg"):
//
//   - distance: two point ids
//   - angle: two line ids
//   - radius/diameter: one circle id
//   - arcLength: one arc id
type AddDimensionArgs struct {
	SketchIndex int      `json:"sketchIndex"`
	Kind        string   `json:"kind"`
	Entities    []uint64 `json:"entities"`
	Expression  string   `json:"expression"`
	// FarSide selects the far tangent point for a "tangentDistance" dimension (line→circle/arc);
	// the default (false) dimensions to the near side. Ignored by other kinds (#152).
	FarSide bool `json:"farSide,omitempty"`
	// Orientation selects what a "distance" dimension between two points measures — the reference CAD API's
	// DimensionOrientationEnum: "aligned" (default; Euclidean |P2−P1|), "horizontal" (the X
	// separation only, leaving the pair free to slide vertically), or "vertical" (the Y
	// separation only). Empty ⇒ aligned. Ignored by other kinds. #1869.
	Orientation string `json:"orientation,omitempty"`
	// Driven creates the dimension as driven (reference) — it measures but does not constrain —
	// in one call, matching the reference CAD API's Add*(…, bool? Driven=false). The default (false) creates a
	// driving dimension. Setting it here avoids the transient over-constraint of the two-step
	// create-then-SetDriven path (#1875).
	Driven bool `json:"driven,omitempty"`
	// TextPoint is the [x,y] sketch-plane placement (cm) of the dimension's annotation text —
	// the reference CAD API's Point2d TextPoint. Stored on the dimension and reported on enumeration; omitted
	// leaves it unset. #1875.
	TextPoint []float64 `json:"textPoint,omitempty"`
	// LinearDiameter makes an "offsetDim" or "tangentDistance" dimension read as a diameter: its
	// value is twice the measured linear distance (the reference CAD API's bool LinearDiameter). Ignored by
	// other kinds. #1875.
	LinearDiameter bool `json:"linearDiameter,omitempty"`
}

// AddDimensionResult is the response of [MethodSketchAddDimension]: the new dimension's
// index in the sketch's dimensional-constraint collection, its kind, the backing
// parameter name, the current measured value (model units), and the sketch's resulting DOF.
type AddDimensionResult struct {
	Index     int     `json:"index"`
	Kind      string  `json:"kind"`
	Parameter string  `json:"parameter"`
	Value     float64 `json:"value"`
	DOF       int     `json:"dof"`
}

// DeleteSketchDimensionArgs is the request of [MethodSketchDeleteDimension]: which dimension, by its
// index in the sketch's dimensional-constraint collection, to remove. Deleting a dimension frees
// the degree of freedom it held and drops its backing parameter with it, so any expression
// referring to that parameter must be rewritten first. #2017.
type DeleteSketchDimensionArgs struct {
	SketchIndex    int `json:"sketchIndex"`
	DimensionIndex int `json:"dimensionIndex"`
}

// DeleteSketchDimensionResult is the response of [MethodSketchDeleteDimension]: the sketch's degrees of
// freedom after the removal, so a caller can confirm the constraint it dropped was the one holding
// the DOF it meant to free.
type DeleteSketchDimensionResult struct {
	DOF int `json:"dof"`
}

// MoveSketchDimensionArgs is the request of [MethodSketchMoveDimension]: where to place a dimension's
// annotation text, as an [x,y] sketch-plane point in centimetres. This is the same placement
// [AddDimensionArgs.TextPoint] sets at create time and that dragging the label in the sketch
// editor writes; it moves only the annotation, never the geometry being measured. #2017.
type MoveSketchDimensionArgs struct {
	SketchIndex    int       `json:"sketchIndex"`
	DimensionIndex int       `json:"dimensionIndex"`
	TextPoint      []float64 `json:"textPoint"`
}

// DriveDimensionArgs is the request of [MethodSketchDriveDimension]: which dimension (by
// collection index) to edit, an optional new value (a unit-bearing expression; empty
// leaves it), whether to set its driven flag (SetDriven + Driven), and optional limits.
type DriveDimensionArgs struct {
	SketchIndex    int     `json:"sketchIndex"`
	DimensionIndex int     `json:"dimensionIndex"`
	Expression     string  `json:"expression,omitempty"`
	SetDriven      bool    `json:"setDriven,omitempty"`
	Driven         bool    `json:"driven,omitempty"`
	SetLimits      bool    `json:"setLimits,omitempty"`
	Min            float64 `json:"min,omitempty"`
	Max            float64 `json:"max,omitempty"`
}
