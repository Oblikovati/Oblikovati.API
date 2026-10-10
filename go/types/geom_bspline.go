// SPDX-License-Identifier: Apache-2.0

package types

// BSplineCurveDef is the complete NURBS curve recipe — the
// BSplineCurveDefinition equivalent as pure data, usable locally by an add-in
// and as the wire encoding of a spline crossing the boundary. Weights may be nil
// for a non-rational curve (all 1).
type BSplineCurveDef struct {
	Degree  int       `json:"degree"`
	Poles   []Point   `json:"poles"`
	Weights []float64 `json:"weights,omitempty"`
	Knots   []float64 `json:"knots"`
}

// BSplineCurve2dDef is the 2D NURBS curve recipe.
type BSplineCurve2dDef struct {
	Degree  int       `json:"degree"`
	Poles   []Point2d `json:"poles"`
	Weights []float64 `json:"weights,omitempty"`
	Knots   []float64 `json:"knots"`
}

// BSplineSurfaceDef is the NURBS surface recipe: poles in row-major (u-major)
// order, PolesU×PolesV of them, with optional weights in the same layout.
type BSplineSurfaceDef struct {
	DegreeU int       `json:"degreeU"`
	DegreeV int       `json:"degreeV"`
	PolesU  int       `json:"polesU"`
	PolesV  int       `json:"polesV"`
	Poles   []Point   `json:"poles"`
	Weights []float64 `json:"weights,omitempty"`
	KnotsU  []float64 `json:"knotsU"`
	KnotsV  []float64 `json:"knotsV"`
}
