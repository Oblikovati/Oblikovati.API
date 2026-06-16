// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// AssemblyConstraints is the assembly constraint operation group (M12-F01,
// Oblikovati/Oblikovati#358/#363): add the relationships that position one occurrence
// relative to another, set their limits, delete them, solve the active assembly, and read
// its health and per-occurrence degrees of freedom. Each add returns the new constraint's
// info after the assembly re-solves; solve and health return the assembly health report.
type AssemblyConstraints struct{ c *Client }

// AssemblyConstraints returns the assembly constraint operation group.
func (c *Client) AssemblyConstraints() AssemblyConstraints { return AssemblyConstraints{c} }

// List returns the active assembly's constraint set in creation order.
//
// mcp:tool list_assembly_constraints
// mcp:summary List the active assembly's constraints: each with id, kind, name, its two geometry inputs (occurrence id + entity reference key), driven value, solution type, limits, and health. The ids address the delete/setLimits tools.
func (a AssemblyConstraints) List() (wire.ConstraintsResult, error) {
	var r wire.ConstraintsResult
	return r, a.c.call(wire.MethodAssemblyConstraintsList, struct{}{}, &r)
}

// AddMate makes geometry A coincident with geometry B at an offset, e.g.
// AddMate(wire.AddMateArgs{A: faceA, B: faceB, Offset: 0}).
//
// mcp:tool add_mate_constraint
// mcp:summary Mate two component geometries (each: occurrence id + entity reference key) coincident at offset (cm). solution "opposed" (default) faces normals at each other, "aligned" matches a flush. Solves and returns the constraint.
func (a AssemblyConstraints) AddMate(args wire.AddMateArgs) (wire.ConstraintResult, error) {
	var r wire.ConstraintResult
	return r, a.c.call(wire.MethodAssemblyConstraintsAddMate, args, &r)
}

// AddFlush makes faces A and B co-planar at an offset, e.g.
// AddFlush(wire.AddFlushArgs{A: faceA, B: faceB}).
//
// mcp:tool add_flush_constraint
// mcp:summary Make two component faces co-planar (normals aligned) at offset (cm). Solves and returns the constraint.
func (a AssemblyConstraints) AddFlush(args wire.AddFlushArgs) (wire.ConstraintResult, error) {
	var r wire.ConstraintResult
	return r, a.c.call(wire.MethodAssemblyConstraintsAddFlush, args, &r)
}

// AddAngle holds an angle (radians) between directions A and B, e.g.
// AddAngle(wire.AddAngleArgs{A: faceA, B: faceB, Angle: math.Pi / 2}).
//
// mcp:tool add_angle_constraint
// mcp:summary Hold angle (radians) between two component directions. solution "undirected" (default), "directed", or "reference-vector". Solves and returns the constraint.
func (a AssemblyConstraints) AddAngle(args wire.AddAngleArgs) (wire.ConstraintResult, error) {
	var r wire.ConstraintResult
	return r, a.c.call(wire.MethodAssemblyConstraintsAddAngle, args, &r)
}

// AddTangent keeps face A tangent to curved face B, e.g.
// AddTangent(wire.AddTangentArgs{A: planeA, B: cylB, Inside: false}).
//
// mcp:tool add_tangent_constraint
// mcp:summary Keep a face tangent to a curved face. inside:true wraps B around A; false is outside tangency. Solves and returns the constraint.
func (a AssemblyConstraints) AddTangent(args wire.AddTangentArgs) (wire.ConstraintResult, error) {
	var r wire.ConstraintResult
	return r, a.c.call(wire.MethodAssemblyConstraintsAddTangent, args, &r)
}

// AddInsert combines an axis mate and a plane mate at an offset (a bolt into a hole), e.g.
// AddInsert(wire.AddInsertArgs{A: holeA, B: shaftB, Offset: 0}).
//
// mcp:tool add_insert_constraint
// mcp:summary Insert: collinear axes plus a plane mate at offset (cm) — a bolt into a hole. aligned:true uses the aligned plane sense; default is opposed. Solves and returns the constraint.
func (a AssemblyConstraints) AddInsert(args wire.AddInsertArgs) (wire.ConstraintResult, error) {
	var r wire.ConstraintResult
	return r, a.c.call(wire.MethodAssemblyConstraintsAddInsert, args, &r)
}

// Snap is "grip snap": it infers the assembly constraint that snaps geometry A (on the component to
// move) onto target geometry B and re-solves, e.g. Snap(wire.SnapConstraintArgs{A: faceA, B: faceB}).
//
// mcp:tool assembly_snap_constrain
// mcp:summary Grip snap: pick a geometry A on the component to move and a target B on another component; the host infers the constraint that snaps A onto B (planar faces -> mate/flush, cylinder axes -> insert, axis pair -> mate, plane+cylinder -> tangent, point -> coincident), creates it, and re-solves so the part jumps into place. prefer ("mate"|"flush"|"insert"|"tangent") overrides the inference. Returns the created constraint (its type is what was inferred).
func (a AssemblyConstraints) Snap(args wire.SnapConstraintArgs) (wire.ConstraintResult, error) {
	var r wire.ConstraintResult
	return r, a.c.call(wire.MethodAssemblyConstraintsSnap, args, &r)
}

// AddSymmetry positions A and B symmetrically about a plane, e.g.
// AddSymmetry(wire.AddSymmetryArgs{A: faceA, B: faceB, Plane: mid}).
//
// mcp:tool add_symmetry_constraint
// mcp:summary Position two component geometries symmetrically about a plane (a planar face or work-plane reference). Solves and returns the constraint.
func (a AssemblyConstraints) AddSymmetry(args wire.AddSymmetryArgs) (wire.ConstraintResult, error) {
	var r wire.ConstraintResult
	return r, a.c.call(wire.MethodAssemblyConstraintsAddSymmetry, args, &r)
}

// AddRotateRotate couples two rotations by a gear ratio, e.g.
// AddRotateRotate(wire.AddRotateRotateArgs{A: gearA, B: gearB, Ratio: 2}).
//
// mcp:tool add_rotate_rotate_constraint
// mcp:summary Couple two rotation axes by gear ratio (revolutions of B per revolution of A). Solves and returns the constraint.
func (a AssemblyConstraints) AddRotateRotate(args wire.AddRotateRotateArgs) (wire.ConstraintResult, error) {
	var r wire.ConstraintResult
	return r, a.c.call(wire.MethodAssemblyConstraintsAddRotateRotate, args, &r)
}

// AddRotateTranslate couples a rotation to a translation (rack and pinion), e.g.
// AddRotateTranslate(wire.AddRotateTranslateArgs{A: pinion, B: rack, Distance: 6.28}).
//
// mcp:tool add_rotate_translate_constraint
// mcp:summary Couple a rotation axis to a translation axis by distance moved per revolution (cm) — rack and pinion. Solves and returns the constraint.
func (a AssemblyConstraints) AddRotateTranslate(args wire.AddRotateTranslateArgs) (wire.ConstraintResult, error) {
	var r wire.ConstraintResult
	return r, a.c.call(wire.MethodAssemblyConstraintsAddRotateTranslate, args, &r)
}

// AddTranslateTranslate couples two translations by a ratio, e.g.
// AddTranslateTranslate(wire.AddTranslateTranslateArgs{A: slideA, B: slideB, Ratio: 2}).
//
// mcp:tool add_translate_translate_constraint
// mcp:summary Couple two translation axes by ratio (distance of B per unit distance of A). Solves and returns the constraint.
func (a AssemblyConstraints) AddTranslateTranslate(args wire.AddTranslateTranslateArgs) (wire.ConstraintResult, error) {
	var r wire.ConstraintResult
	return r, a.c.call(wire.MethodAssemblyConstraintsAddTranslateTranslate, args, &r)
}

// AddTransitional keeps face A in sliding contact with face B, e.g.
// AddTransitional(wire.AddTransitionalArgs{A: pinFace, B: slotFace}).
//
// mcp:tool add_transitional_constraint
// mcp:summary Keep a face in sliding contact with a transition face as the component moves (cam/slot). Solves and returns the constraint.
func (a AssemblyConstraints) AddTransitional(args wire.AddTransitionalArgs) (wire.ConstraintResult, error) {
	var r wire.ConstraintResult
	return r, a.c.call(wire.MethodAssemblyConstraintsAddTransitional, args, &r)
}

// AddCustom registers a relationship between A and B solved by the add-in named Kind, e.g.
// AddCustom(wire.AddCustomArgs{A: a, B: b, Kind: "cam-profile", Params: []float64{1, 2}}).
//
// mcp:tool add_custom_constraint
// mcp:summary Register a relationship solved by an add-in (kind names it, params drive it). The built-in solver leaves it free unless that add-in solver is installed. Returns the constraint.
func (a AssemblyConstraints) AddCustom(args wire.AddCustomArgs) (wire.ConstraintResult, error) {
	var r wire.ConstraintResult
	return r, a.c.call(wire.MethodAssemblyConstraintsAddCustom, args, &r)
}

// Delete removes the constraint and returns the refreshed set, e.g. Delete(id).
//
// mcp:tool delete_assembly_constraint
// mcp:summary Delete an assembly constraint (id) and re-solve. Returns the remaining constraint set.
func (a AssemblyConstraints) Delete(id uint64) (wire.ConstraintsResult, error) {
	var r wire.ConstraintsResult
	return r, a.c.call(wire.MethodAssemblyConstraintsDelete, wire.DeleteAssemblyConstraintArgs{ID: id}, &r)
}

// SetLimits sets (or clears) the driven-value limits of a constraint, e.g.
// SetLimits(wire.SetConstraintLimitsArgs{ID: id, Limits: wire.ConstraintLimits{HasMax: true, Max: 1}}).
//
// mcp:tool set_constraint_limits
// mcp:summary Set or clear a constraint's driven-value limits (min/max/resting). Each bound is optional via its has* flag. Returns the updated constraint.
func (a AssemblyConstraints) SetLimits(args wire.SetConstraintLimitsArgs) (wire.ConstraintResult, error) {
	var r wire.ConstraintResult
	return r, a.c.call(wire.MethodAssemblyConstraintsSetLimits, args, &r)
}

// Solve re-positions the active assembly's occurrences to satisfy its constraints and
// returns the resulting health and DOF report.
//
// mcp:tool solve_assembly_constraints
// mcp:summary Solve the active assembly: reposition occurrences to satisfy the constraints, then report overall health, redundant-constraint count, total remaining degrees of freedom, and per-occurrence DOF.
func (a AssemblyConstraints) Solve() (wire.AssemblyHealthResult, error) {
	var r wire.AssemblyHealthResult
	return r, a.c.call(wire.MethodAssemblyConstraintsSolve, struct{}{}, &r)
}

// Health returns the active assembly's constraint health and DOF report without re-solving.
//
// mcp:tool assembly_constraint_health
// mcp:summary Report the active assembly's constraint health: overall status, redundant-constraint count, total remaining degrees of freedom, and the per-occurrence DOF breakdown — without re-solving.
func (a AssemblyConstraints) Health() (wire.AssemblyHealthResult, error) {
	var r wire.AssemblyHealthResult
	return r, a.c.call(wire.MethodAssemblyConstraintsHealth, struct{}{}, &r)
}
