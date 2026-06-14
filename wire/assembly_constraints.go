// SPDX-License-Identifier: Apache-2.0

package wire

// The assembly constraint surface (M12-F01, Oblikovati/Oblikovati#358/#363): add the
// relationships that position one occurrence relative to another, solve them, and read
// the assembly's health and per-occurrence degrees of freedom. Geometry inputs are
// addressed like the assembly feature dress-ups — an occurrence session id plus a
// reference key on its component (from model.referenceKeys), resolved per placement in
// assembly space. Offsets/angles are in database units (cm) / radians.

// ConstraintGeomRef addresses one geometry input of a constraint: the Occurrence (session
// id) whose component carries the geometry, and the Entity's reference key on that
// component (a face, edge, or vertex key). The solver resolves it to assembly-space
// geometry through the occurrence's placement.
type ConstraintGeomRef struct {
	Occurrence uint64 `json:"occurrence"`
	Entity     string `json:"entity"`
}

// ConstraintLimits bounds a constraint's driven value (offset/angle). Each bound is
// optional; an absent bound (its Has* flag false) does not clamp. Resting is the value
// a drive returns to when released.
type ConstraintLimits struct {
	HasMin     bool    `json:"hasMin,omitempty"`
	Min        float64 `json:"min,omitempty"`
	HasMax     bool    `json:"hasMax,omitempty"`
	Max        float64 `json:"max,omitempty"`
	HasResting bool    `json:"hasResting,omitempty"`
	Resting    float64 `json:"resting,omitempty"`
}

// AssemblyConstraintInfo is one row of the assembly's constraint set: its session id, kind,
// name, the two geometry inputs, the driven value (offset for mate/flush/insert, angle for
// angle, ratio for the motion kinds), the solution-type discriminator where it applies,
// optional limits, health, and suppression. Default-zero/false fields are omitted.
type AssemblyConstraintInfo struct {
	ID         uint64            `json:"id"`
	Type       string            `json:"type"`
	Name       string            `json:"name"`
	A          ConstraintGeomRef `json:"a"`
	B          ConstraintGeomRef `json:"b"`
	Value      float64           `json:"value,omitempty"`
	Solution   string            `json:"solution,omitempty"`
	Limits     *ConstraintLimits `json:"limits,omitempty"`
	Health     string            `json:"health,omitempty"`
	Suppressed bool              `json:"suppressed,omitempty"`
}

// OccurrenceDOFInfo reports one occurrence's remaining free degrees of freedom after a
// solve — 0 when fully constrained (or grounded), up to 6 when free. It is the
// "this part is under-constrained (N DOF remain)" diagnostic (cross-ref #346).
type OccurrenceDOFInfo struct {
	Occurrence       uint64 `json:"occurrence"`
	DegreesOfFreedom int    `json:"degreesOfFreedom"`
}

// AssemblyHealthResult is the reply of [MethodAssemblyConstraintsSolve] and
// [MethodAssemblyConstraintsHealth]: the assembly's overall constraint health
// ([types.HealthStatus] string), counts of constraints and redundant (over-constraining)
// ones, the total remaining free DOF, and the per-occurrence DOF breakdown.
type AssemblyHealthResult struct {
	Status           string              `json:"status"`
	Constraints      int                 `json:"constraints"`
	Redundant        int                 `json:"redundant"`
	DegreesOfFreedom int                 `json:"degreesOfFreedom"`
	Occurrences      []OccurrenceDOFInfo `json:"occurrences,omitempty"`
	Converged        bool                `json:"converged"`
}

// ConstraintsResult is the reply of [MethodAssemblyConstraintsList]: the active assembly's
// constraint set in creation order.
type ConstraintsResult struct {
	Constraints []AssemblyConstraintInfo `json:"constraints"`
}

// ConstraintResult is the reply of the single-constraint add operations: the new
// constraint's info (after the assembly re-solves).
type ConstraintResult struct {
	Constraint AssemblyConstraintInfo `json:"constraint"`
}

// AddMateArgs is the request of [MethodAssemblyConstraintsAddMate]: make geometry A
// coincident with geometry B at Offset (database units). Solution selects opposed (the
// default mate) or aligned (flush) normals where the geometry is directional; "" ⇒ opposed.
type AddMateArgs struct {
	A        ConstraintGeomRef `json:"a"`
	B        ConstraintGeomRef `json:"b"`
	Offset   float64           `json:"offset,omitempty"`
	Solution string            `json:"solution,omitempty"`
}

// AddFlushArgs is the request of [MethodAssemblyConstraintsAddFlush]: make faces A and B
// co-planar (normals aligned) at Offset.
type AddFlushArgs struct {
	A      ConstraintGeomRef `json:"a"`
	B      ConstraintGeomRef `json:"b"`
	Offset float64           `json:"offset,omitempty"`
}

// AddAngleArgs is the request of [MethodAssemblyConstraintsAddAngle]: hold Angle (radians)
// between directions A and B. Solution selects undirected (default), directed, or
// reference-vector measurement; "" ⇒ undirected.
type AddAngleArgs struct {
	A        ConstraintGeomRef `json:"a"`
	B        ConstraintGeomRef `json:"b"`
	Angle    float64           `json:"angle"`
	Solution string            `json:"solution,omitempty"`
}

// AddTangentArgs is the request of [MethodAssemblyConstraintsAddTangent]: keep face A
// tangent to curved face B. Inside selects inside tangency (B wraps A); false ⇒ outside.
type AddTangentArgs struct {
	A      ConstraintGeomRef `json:"a"`
	B      ConstraintGeomRef `json:"b"`
	Inside bool              `json:"inside,omitempty"`
}

// AddInsertArgs is the request of [MethodAssemblyConstraintsAddInsert]: an insert combines
// an axis mate (A's axis collinear with B's) and a plane mate at Offset — a bolt into a
// hole. Opposed selects the opposed (default) plane sense; false ⇒ aligned.
type AddInsertArgs struct {
	A       ConstraintGeomRef `json:"a"`
	B       ConstraintGeomRef `json:"b"`
	Offset  float64           `json:"offset,omitempty"`
	Aligned bool              `json:"aligned,omitempty"`
}

// AddSymmetryArgs is the request of [MethodAssemblyConstraintsAddSymmetry]: position
// geometry A and geometry B symmetrically about the Plane (a planar face/work-plane ref).
type AddSymmetryArgs struct {
	A     ConstraintGeomRef `json:"a"`
	B     ConstraintGeomRef `json:"b"`
	Plane ConstraintGeomRef `json:"plane"`
}

// AddRotateRotateArgs is the request of [MethodAssemblyConstraintsAddRotateRotate]: couple
// rotation A to rotation B by gear Ratio (revolutions of B per revolution of A).
type AddRotateRotateArgs struct {
	A     ConstraintGeomRef `json:"a"`
	B     ConstraintGeomRef `json:"b"`
	Ratio float64           `json:"ratio"`
}

// AddRotateTranslateArgs is the request of [MethodAssemblyConstraintsAddRotateTranslate]:
// couple rotation A to translation B by Distance moved per revolution (rack and pinion).
type AddRotateTranslateArgs struct {
	A        ConstraintGeomRef `json:"a"`
	B        ConstraintGeomRef `json:"b"`
	Distance float64           `json:"distance"`
}

// AddTranslateTranslateArgs is the request of
// [MethodAssemblyConstraintsAddTranslateTranslate]: couple translation A to translation B
// by Ratio (distance of B per unit distance of A).
type AddTranslateTranslateArgs struct {
	A     ConstraintGeomRef `json:"a"`
	B     ConstraintGeomRef `json:"b"`
	Ratio float64           `json:"ratio"`
}

// AddTransitionalArgs is the request of [MethodAssemblyConstraintsAddTransitional]: keep
// face A in sliding contact with face B (the transition surface) as the component moves.
type AddTransitionalArgs struct {
	A ConstraintGeomRef `json:"a"`
	B ConstraintGeomRef `json:"b"`
}

// AddCustomArgs is the request of [MethodAssemblyConstraintsAddCustom]: register a
// relationship between A and B that the add-in named Kind solves, driven by Params. The
// built-in solver treats it as residual-free unless an add-in solver is installed.
type AddCustomArgs struct {
	A      ConstraintGeomRef `json:"a"`
	B      ConstraintGeomRef `json:"b"`
	Kind   string            `json:"kind"`
	Params []float64         `json:"params,omitempty"`
}

// DeleteAssemblyConstraintArgs is the request of [MethodAssemblyConstraintsDelete]: remove
// the constraint with id ID from the active assembly.
type DeleteAssemblyConstraintArgs struct {
	ID uint64 `json:"id"`
}

// SetConstraintLimitsArgs is the request of [MethodAssemblyConstraintsSetLimits]: set (or
// clear) the driven-value Limits of the constraint with id ID.
type SetConstraintLimitsArgs struct {
	ID     uint64           `json:"id"`
	Limits ConstraintLimits `json:"limits"`
}
