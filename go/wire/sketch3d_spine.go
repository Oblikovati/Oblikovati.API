// SPDX-License-Identifier: Apache-2.0

package wire

// Sketch3DArgs is the request of the per-3D-sketch methods that take only a target:
// [MethodSketch3DGet], [MethodSketch3DEdit], [MethodSketch3DExitEdit],
// [MethodSketch3DSolve], [MethodSketch3DDelete], [MethodSketch3DEntities],
// [MethodSketch3DConstraints], [MethodSketch3DDimensions],
// [MethodSketch3DConstraintStatus]. SketchIndex is the sketch's index in the active
// part's 3D-sketch collection (as returned by [MethodSketch3DCreate]/[MethodSketch3DList]).
type Sketch3DArgs struct {
	SketchIndex int `json:"sketchIndex"`
}

// CreateSketch3DArgs is the request of [MethodSketch3DCreate]: an optional name for the
// new (empty) 3D sketch. A 3D sketch has no host plane — its geometry lives directly in
// model space.
type CreateSketch3DArgs struct {
	Name string `json:"name,omitempty"`
}

// CreateSketch3DResult is the response of [MethodSketch3DCreate]: the new sketch's index.
type CreateSketch3DResult struct {
	SketchIndex int `json:"sketchIndex"`
}

// Sketch3DInfo is one row of [MethodSketch3DList] and the result of [MethodSketch3DGet]:
// a 3D sketch's identity, visibility, entity count, remaining DOF, edit state, health,
// and the display/solve overrides (dimensions-visible, color, defer).
type Sketch3DInfo struct {
	Index             int    `json:"index"`
	Name              string `json:"name"`
	Visible           bool   `json:"visible"`
	DimensionsVisible bool   `json:"dimensionsVisible"`
	EntityCount       int    `json:"entityCount"`
	DOF               int    `json:"dof"`
	Editing           bool   `json:"editing"`
	Healthy           bool   `json:"healthy"`
	// Shared reports the "share sketch" flag. (A 3D sketch's consumed/ownedBy state
	// awaits the 3D path→sketch dependency link; see the 2D [SketchInfo] for those notions.)
	Shared       bool   `json:"shared,omitempty"`
	Color        string `json:"color,omitempty"`
	DeferUpdates bool   `json:"deferUpdates,omitempty"`
}

// SetSketch3DPropertyArgs is the request of [MethodSketch3DSetProperty]: which sketch,
// which property ("name" | "visible" | "dimensionsVisible" | "color" | "deferUpdates"),
// and the new value as a string (bools as "true"/"false"). The response is the updated
// [Sketch3DInfo].
type SetSketch3DPropertyArgs struct {
	SketchIndex int    `json:"sketchIndex"`
	Property    string `json:"property"`
	Value       string `json:"value"`
}

// ListSketches3DResult is the response of [MethodSketch3DList].
type ListSketches3DResult struct {
	Sketches []Sketch3DInfo `json:"sketches"`
}

// EditSketch3DResult is the response of [MethodSketch3DEdit]/[MethodSketch3DExitEdit]:
// the sketch index and its resulting edit state.
type EditSketch3DResult struct {
	SketchIndex int  `json:"sketchIndex"`
	Editing     bool `json:"editing"`
}

// SolveSketch3DResult is the response of [MethodSketch3DSolve]: remaining DOF, a status
// string ("well" | "under" | "over"), whether the Newton/LM iteration converged, and
// the resulting health.
type SolveSketch3DResult struct {
	SketchIndex int    `json:"sketchIndex"`
	DOF         int    `json:"dof"`
	Status      string `json:"status"`
	Converged   bool   `json:"converged"`
	Healthy     bool   `json:"healthy"`
}

// Sketch3DEntityInfo is one enumerated entity from [MethodSketch3DEntities]: its index,
// session id, kind ([oblikovati.org/api/types.Sketch3DEntityKind]), construction
// flag, the defining points (each [x,y,z] in model database units, cm), and a radius for
// circular kinds (0 otherwise). MoveableStatus answers whether interactive tools may
// drag the entity ([oblikovati.org/api/types.GeometryMoveableStatus] wire spelling —
// M06-F11, Oblikovati/Oblikovati#626).
type Sketch3DEntityInfo struct {
	Index        int         `json:"index"`
	ID           uint64      `json:"id"`
	Kind         string      `json:"kind"`
	Construction bool        `json:"construction,omitempty"`
	Points       [][]float64 `json:"points,omitempty"`
	Radius       float64     `json:"radius,omitempty"`
	// ReferenceKey is the entity's persistent reference key (Oblikovati/Oblikovati#153): a
	// document-scoped UUID stable across save/load and edits, unlike the session ID. Store
	// it to refer to this entity durably; rebind it with [MethodSketchResolveReference].
	ReferenceKey   string `json:"referenceKey,omitempty"`
	MoveableStatus string `json:"moveableStatus,omitempty"`
	// CoordinateSystem reports an equation curve's coordinate system (cylindrical/spherical),
	// absent for a Cartesian curve and every other kind, so a round-trip preserves it (#1846).
	CoordinateSystem string `json:"coordinateSystem,omitempty"`
}

// EnumerateEntities3DResult is the response of [MethodSketch3DEntities].
type EnumerateEntities3DResult struct {
	Entities []Sketch3DEntityInfo `json:"entities"`
}

// Constraint3DInfo is one enumerated geometric constraint from
// [MethodSketch3DConstraints]: its index, kind
// ([oblikovati.org/api/types.Geometric3DConstraintKind]), and the session ids of
// the entities/points it relates.
type Constraint3DInfo struct {
	Index    int      `json:"index"`
	Kind     string   `json:"kind"`
	Entities []uint64 `json:"entities,omitempty"`
}

// ListConstraints3DResult is the response of [MethodSketch3DConstraints].
type ListConstraints3DResult struct {
	Constraints []Constraint3DInfo `json:"constraints"`
}

// Dimension3DInfo is one enumerated dimensional constraint from
// [MethodSketch3DDimensions]: its index, kind
// ([oblikovati.org/api/types.Dimension3DConstraintKind]), backing parameter name
// + expression, current model value (cm/rad), and whether it is driven (reports) rather
// than driving (constrains).
type Dimension3DInfo struct {
	Index      int     `json:"index"`
	Kind       string  `json:"kind"`
	Name       string  `json:"name"`
	Expression string  `json:"expression"`
	Value      float64 `json:"value"`
	Driven     bool    `json:"driven"`
}

// ListDimensions3DResult is the response of [MethodSketch3DDimensions].
type ListDimensions3DResult struct {
	Dimensions []Dimension3DInfo `json:"dimensions"`
}
