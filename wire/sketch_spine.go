// SPDX-License-Identifier: Apache-2.0

package wire

// SketchArgs is the request of the per-sketch methods that take only a target:
// [MethodSketchGet], [MethodSketchEdit], [MethodSketchExitEdit], [MethodSketchSolve],
// [MethodSketchDelete], [MethodSketchEntities], [MethodSketchConstraints],
// [MethodSketchDimensions]. SketchIndex is the sketch's index in the active part's
// sketch collection (as returned by [MethodSketchCreate] / [MethodSketchList]).
type SketchArgs struct {
	SketchIndex int `json:"sketchIndex"`
}

// SketchInfo is one row of [MethodSketchList] and the result of [MethodSketchGet]: a
// sketch's identity, host plane label, visibility, entity count, remaining DOF, edit
// state, health, and the display/solve overrides (color, line type/weight, defer).
type SketchInfo struct {
	Index        int     `json:"index"`
	Name         string  `json:"name"`
	Plane        string  `json:"plane"`
	Visible      bool    `json:"visible"`
	EntityCount  int     `json:"entityCount"`
	DOF          int     `json:"dof"`
	Editing      bool    `json:"editing"`
	Healthy      bool    `json:"healthy"`
	Color        string  `json:"color,omitempty"`
	LineType     string  `json:"lineType,omitempty"`
	LineWeight   float64 `json:"lineWeight,omitempty"`
	DeferUpdates bool    `json:"deferUpdates,omitempty"`
}

// SetSketchPropertyArgs is the request of [MethodSketchSetProperty]: which sketch, which
// property ("name" | "visible" | "color" | "lineType" | "lineWeight" | "deferUpdates"),
// and the new value as a string (bools as "true"/"false", line weight as a unit-bearing
// length like "0.5 mm"). The response is the updated [SketchInfo].
type SetSketchPropertyArgs struct {
	SketchIndex int    `json:"sketchIndex"`
	Property    string `json:"property"`
	Value       string `json:"value"`
}

// ListSketchesResult is the response of [MethodSketchList].
type ListSketchesResult struct {
	Sketches []SketchInfo `json:"sketches"`
}

// EditSketchResult is the response of [MethodSketchEdit] / [MethodSketchExitEdit]:
// the sketch index and its resulting edit state.
type EditSketchResult struct {
	SketchIndex int  `json:"sketchIndex"`
	Editing     bool `json:"editing"`
}

// SolveSketchResult is the response of [MethodSketchSolve]: the solve outcome —
// remaining DOF, a status string ("well" | "under" | "over"), whether the Newton/LM
// iteration converged, and the resulting health.
type SolveSketchResult struct {
	SketchIndex int    `json:"sketchIndex"`
	DOF         int    `json:"dof"`
	Status      string `json:"status"`
	Converged   bool   `json:"converged"`
	Healthy     bool   `json:"healthy"`
}

// ProfileInfo is one enumerated profile from [MethodSketchProfiles]: its index, enclosed
// area (sketch-plane cm², holes subtracted), whether it is closed (a solid-extrudable
// region), and the number of hole loops it contains. The Index feeds features.add's
// profileIndex.
type ProfileInfo struct {
	Index  int     `json:"index"`
	Area   float64 `json:"area"`
	Closed bool    `json:"closed"`
	Holes  int     `json:"holes"`
}

// ListProfilesResult is the response of [MethodSketchProfiles].
type ListProfilesResult struct {
	Profiles []ProfileInfo `json:"profiles"`
}

// ConstraintStatusResult is the response of [MethodSketchConstraintStatus]: the sketch's
// constraint state without moving geometry (a non-mutating DOF analysis). Status is a
// [oblikovati/api/types.ConstraintStatus]; DOF is the remaining free degrees
// of freedom; Variables/Equations are the system size; Redundant counts the dependent
// constraints (> 0 ⇒ over-constrained).
type ConstraintStatusResult struct {
	Status    string `json:"status"`
	DOF       int    `json:"dof"`
	Variables int    `json:"variables"`
	Equations int    `json:"equations"`
	Redundant int    `json:"redundant"`
}

// SketchEntityInfo is one enumerated entity from [MethodSketchEntities]: its index,
// session id, kind ([oblikovati/api/types.SketchEntityKind]), construction
// flag, the defining points (each [x,y] in sketch-plane cm), and a radius for circular
// kinds (0 otherwise).
type SketchEntityInfo struct {
	Index        int         `json:"index"`
	ID           uint64      `json:"id"`
	Kind         string      `json:"kind"`
	Construction bool        `json:"construction"`
	Points       [][]float64 `json:"points"`
	Radius       float64     `json:"radius,omitempty"`
}

// EnumerateEntitiesResult is the response of [MethodSketchEntities].
type EnumerateEntitiesResult struct {
	Entities []SketchEntityInfo `json:"entities"`
}

// ConstraintInfo is one enumerated geometric constraint from [MethodSketchConstraints]:
// its index, kind ([oblikovati/api/types.GeometricConstraintKind]), and the
// session ids of the entities it relates.
type ConstraintInfo struct {
	Index    int      `json:"index"`
	Kind     string   `json:"kind"`
	Entities []uint64 `json:"entities"`
}

// ListConstraintsResult is the response of [MethodSketchConstraints].
type ListConstraintsResult struct {
	Constraints []ConstraintInfo `json:"constraints"`
}

// DimensionInfo is one enumerated dimensional constraint from [MethodSketchDimensions]:
// its index, kind ([oblikovati/api/types.DimensionConstraintKind]), backing
// parameter name + expression, current model value (cm/rad), and whether it is driven
// (reports) rather than driving (constrains).
type DimensionInfo struct {
	Index      int     `json:"index"`
	Kind       string  `json:"kind"`
	Name       string  `json:"name"`
	Expression string  `json:"expression"`
	Value      float64 `json:"value"`
	Driven     bool    `json:"driven"`
}

// ListDimensionsResult is the response of [MethodSketchDimensions].
type ListDimensionsResult struct {
	Dimensions []DimensionInfo `json:"dimensions"`
}
