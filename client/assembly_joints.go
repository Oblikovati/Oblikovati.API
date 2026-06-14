// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// AssemblyJoints is the assembly joint operation group (M12-F02, Oblikovati/Oblikovati#359/
// #364): author the simplified joints that establish a degree-of-freedom set between two
// occurrences, set their limits and flip, and delete them. Joints solve together with
// constraints — use solve_assembly_constraints / assembly_constraint_health to position and
// report. Each add returns the new joint's info after the assembly re-solves.
type AssemblyJoints struct{ c *Client }

// AssemblyJoints returns the assembly joint operation group.
func (c *Client) AssemblyJoints() AssemblyJoints { return AssemblyJoints{c} }

// List returns the active assembly's joint set in creation order.
//
// mcp:tool list_assembly_joints
// mcp:summary List the active assembly's joints: each with id, kind, name, its two joint-origin geometry inputs (occurrence id + entity reference key), flip, free DOF, and limits.
func (a AssemblyJoints) List() (wire.AssemblyJointsResult, error) {
	var r wire.AssemblyJointsResult
	return r, a.c.call(wire.MethodAssemblyJointsList, struct{}{}, &r)
}

// AddRigid fixes two components together (0 DOF).
//
// mcp:tool add_rigid_joint
// mcp:summary Add a rigid joint fixing two components together (0 DOF) at their joint origins (each: occurrence id + entity reference key). Solves and returns the joint.
func (a AssemblyJoints) AddRigid(args wire.AddJointArgs) (wire.AssemblyJointResult, error) {
	var r wire.AssemblyJointResult
	return r, a.c.call(wire.MethodAssemblyJointsAddRigid, args, &r)
}

// AddRotational allows one rotation about the joint axis (1 DOF).
//
// mcp:tool add_rotational_joint
// mcp:summary Add a rotational joint — one rotation about the joint axis (1 DOF), a hinge — between two component joint origins. flip reverses the facing sense. Solves and returns the joint.
func (a AssemblyJoints) AddRotational(args wire.AddJointArgs) (wire.AssemblyJointResult, error) {
	var r wire.AssemblyJointResult
	return r, a.c.call(wire.MethodAssemblyJointsAddRotational, args, &r)
}

// AddSlider allows one translation along the joint axis (1 DOF).
//
// mcp:tool add_slider_joint
// mcp:summary Add a slider joint — one translation along the joint axis (1 DOF) — between two component joint origins. Solves and returns the joint.
func (a AssemblyJoints) AddSlider(args wire.AddJointArgs) (wire.AssemblyJointResult, error) {
	var r wire.AssemblyJointResult
	return r, a.c.call(wire.MethodAssemblyJointsAddSlider, args, &r)
}

// AddCylindrical allows translation along and rotation about the axis (2 DOF).
//
// mcp:tool add_cylindrical_joint
// mcp:summary Add a cylindrical joint — translation along and rotation about the joint axis (2 DOF) — between two component joint origins. Solves and returns the joint.
func (a AssemblyJoints) AddCylindrical(args wire.AddJointArgs) (wire.AssemblyJointResult, error) {
	var r wire.AssemblyJointResult
	return r, a.c.call(wire.MethodAssemblyJointsAddCylindrical, args, &r)
}

// AddPlanar allows two in-plane translations and a rotation about the normal (3 DOF).
//
// mcp:tool add_planar_joint
// mcp:summary Add a planar joint — two in-plane translations + rotation about the plane normal (3 DOF) — between two component joint origins. Solves and returns the joint.
func (a AssemblyJoints) AddPlanar(args wire.AddJointArgs) (wire.AssemblyJointResult, error) {
	var r wire.AssemblyJointResult
	return r, a.c.call(wire.MethodAssemblyJointsAddPlanar, args, &r)
}

// AddBall allows three rotations about a common point (3 DOF).
//
// mcp:tool add_ball_joint
// mcp:summary Add a ball joint — three rotations about a common point (3 DOF) — between two component joint origins. Solves and returns the joint.
func (a AssemblyJoints) AddBall(args wire.AddJointArgs) (wire.AssemblyJointResult, error) {
	var r wire.AssemblyJointResult
	return r, a.c.call(wire.MethodAssemblyJointsAddBall, args, &r)
}

// Delete removes the joint and returns the refreshed set, e.g. Delete(id).
//
// mcp:tool delete_assembly_joint
// mcp:summary Delete an assembly joint (id) and re-solve. Returns the remaining joint set.
func (a AssemblyJoints) Delete(id uint64) (wire.AssemblyJointsResult, error) {
	var r wire.AssemblyJointsResult
	return r, a.c.call(wire.MethodAssemblyJointsDelete, wire.DeleteJointArgs{ID: id}, &r)
}

// SetLimits sets a joint's linear/angular bounds, e.g.
// SetLimits(wire.SetJointLimitsArgs{ID: id, Limits: wire.JointLimits{HasAngularMax: true, AngularMax: 1.57}}).
//
// mcp:tool set_joint_limits
// mcp:summary Set a joint's driven-value limits — a linear range (cm) and/or an angular range (radians); each bound optional via its has* flag. Returns the updated joint.
func (a AssemblyJoints) SetLimits(args wire.SetJointLimitsArgs) (wire.AssemblyJointResult, error) {
	var r wire.AssemblyJointResult
	return r, a.c.call(wire.MethodAssemblyJointsSetLimits, args, &r)
}

// SetFlip sets a joint's flip sense, e.g. SetFlip(id, true).
//
// mcp:tool set_joint_flip
// mcp:summary Flip (flip:true) or unflip a joint's primary-axis sense (which way the components face). Returns the updated joint.
func (a AssemblyJoints) SetFlip(id uint64, flip bool) (wire.AssemblyJointResult, error) {
	var r wire.AssemblyJointResult
	return r, a.c.call(wire.MethodAssemblyJointsSetFlip, wire.SetJointFlipArgs{ID: id, Flip: flip}, &r)
}

// DSJoints is the DS-joint (degrees-of-freedom / imposed-motion) operation group — the
// kinematic view of joints motion and simulation consumers read.
type DSJoints struct{ c *Client }

// DSJoints returns the DS-joint operation group.
func (c *Client) DSJoints() DSJoints { return DSJoints{c} }

// List returns the active assembly's DS-joint set.
//
// mcp:tool list_ds_joints
// mcp:summary List the active assembly's DS joints: each with id, kind, name, and per-DOF imposed-motion (free/driven/locked) breakdown.
func (d DSJoints) List() (wire.DSJointsResult, error) {
	var r wire.DSJointsResult
	return r, d.c.call(wire.MethodDSJointsList, struct{}{}, &r)
}

// Add adds a DS joint of the given kind between two origins, e.g.
// Add(wire.AddDSJointArgs{A: a, B: b, Type: "rotational"}).
//
// mcp:tool add_ds_joint
// mcp:summary Add a DS joint (type: rigid|rotational|prismatic|cylindrical|planar|spherical) between two component joint origins. Returns the joint with its degrees of freedom.
func (d DSJoints) Add(args wire.AddDSJointArgs) (wire.DSJointResult, error) {
	var r wire.DSJointResult
	return r, d.c.call(wire.MethodDSJointsAdd, args, &r)
}

// SetImposedMotion sets the imposed motion of a DS joint's DOF, e.g.
// SetImposedMotion(wire.SetImposedMotionArgs{ID: id, DOFIndex: 0, ImposedMotion: "locked"}).
//
// mcp:tool set_ds_joint_imposed_motion
// mcp:summary Set a DS joint DOF's imposed motion (free|driven|locked) and value (cm or radians) by dofIndex. Locked removes that DOF from the solve. Returns the updated joint.
func (d DSJoints) SetImposedMotion(args wire.SetImposedMotionArgs) (wire.DSJointResult, error) {
	var r wire.DSJointResult
	return r, d.c.call(wire.MethodDSJointsSetImposedMotion, args, &r)
}

// Delete removes the DS joint and returns the refreshed set.
//
// mcp:tool delete_ds_joint
// mcp:summary Delete a DS joint (id). Returns the remaining DS-joint set.
func (d DSJoints) Delete(id uint64) (wire.DSJointsResult, error) {
	var r wire.DSJointsResult
	return r, d.c.call(wire.MethodDSJointsDelete, wire.DeleteDSJointArgs{ID: id}, &r)
}
