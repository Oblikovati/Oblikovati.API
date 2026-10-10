// SPDX-License-Identifier: Apache-2.0

package wire

// AddSketch3DDimensionArgs is the request of [MethodSketch3DAddDimension] — the
// discriminated 3D dimensional-constraint constructor. Kind is the dimension
// ([oblikovati.org/api/types.Dimension3DConstraintKind]). Entities are the session
// ids of the dimensioned geometry (distance: two point ids; lineLength/radius: one entity
// id; pointPlaneDistance: one point id; twoLineAngle: two line ids). Expression is the
// unit-bearing value ("10 mm", "30 deg"). Plane selects the reference origin plane for
// pointPlaneDistance ("XY" | "XZ" | "YZ").
type AddSketch3DDimensionArgs struct {
	SketchIndex int      `json:"sketchIndex"`
	Kind        string   `json:"kind"`
	Entities    []uint64 `json:"entities,omitempty"`
	Expression  string   `json:"expression"`
	Plane       string   `json:"plane,omitempty"`
}

// AddSketch3DDimensionResult is the response of [MethodSketch3DAddDimension]: the
// dimension's index, its kind, the backing parameter name, the current measured value
// (cm/rad), and the sketch's resulting degree-of-freedom count.
type AddSketch3DDimensionResult struct {
	Index     int     `json:"index"`
	Kind      string  `json:"kind"`
	Parameter string  `json:"parameter"`
	Value     float64 `json:"value"`
	DOF       int     `json:"dof"`
}

// DriveSketch3DDimensionArgs is the request of [MethodSketch3DDriveDimension]: which
// dimension to edit, an optional new value expression, and an optional driving/driven
// toggle (SetDriven gates whether Driven is applied).
type DriveSketch3DDimensionArgs struct {
	SketchIndex    int    `json:"sketchIndex"`
	DimensionIndex int    `json:"dimensionIndex"`
	Expression     string `json:"expression,omitempty"`
	SetDriven      bool   `json:"setDriven,omitempty"`
	Driven         bool   `json:"driven,omitempty"`
}
