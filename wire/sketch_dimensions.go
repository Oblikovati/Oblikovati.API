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
