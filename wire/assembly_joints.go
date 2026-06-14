// SPDX-License-Identifier: Apache-2.0

package wire

// The assembly joint surface (M12-F02, Oblikovati/Oblikovati#359/#364): author the
// simplified joints that establish a degree-of-freedom set between two occurrences, set
// their limits and flip, solve, and read DOF — plus the DS-joint (degrees-of-freedom /
// imposed-motion) view. Joint origins are geometry inputs addressed exactly like assembly
// constraints (an occurrence id + a reference key, [ConstraintGeomRef]); the solve and
// health report reuse [AssemblyHealthResult] (constraints and joints solve together).

// JointLimits bounds a joint's driven values: a linear range (slider/cylindrical
// translation, cm) and an angular range (rotational/cylindrical rotation, radians). Each
// bound is independently optional.
type JointLimits struct {
	HasLinearMin bool    `json:"hasLinearMin,omitempty"`
	LinearMin    float64 `json:"linearMin,omitempty"`
	HasLinearMax bool    `json:"hasLinearMax,omitempty"`
	LinearMax    float64 `json:"linearMax,omitempty"`

	HasAngularMin bool    `json:"hasAngularMin,omitempty"`
	AngularMin    float64 `json:"angularMin,omitempty"`
	HasAngularMax bool    `json:"hasAngularMax,omitempty"`
	AngularMax    float64 `json:"angularMax,omitempty"`
}

// JointInfo is one row of the assembly's joint set: its session id, kind, name, the two
// joint-origin geometry inputs, the flip flag, optional limits, the free DOF the joint
// leaves, health, and suppression.
type JointInfo struct {
	ID               uint64            `json:"id"`
	Type             string            `json:"type"`
	Name             string            `json:"name"`
	A                ConstraintGeomRef `json:"a"`
	B                ConstraintGeomRef `json:"b"`
	Flip             bool              `json:"flip,omitempty"`
	DegreesOfFreedom int               `json:"degreesOfFreedom"`
	Limits           *JointLimits      `json:"limits,omitempty"`
	Health           string            `json:"health,omitempty"`
	Suppressed       bool              `json:"suppressed,omitempty"`
}

// AssemblyJointsResult is the reply of [MethodAssemblyJointsList]: the active assembly's
// joint set in creation order.
type AssemblyJointsResult struct {
	Joints []JointInfo `json:"joints"`
}

// AssemblyJointResult is the reply of the single-joint add/mutate operations: the affected
// joint's info (after the assembly re-solves).
type AssemblyJointResult struct {
	Joint JointInfo `json:"joint"`
}

// AddJointArgs is the request of the add-joint methods ([MethodAssemblyJointsAddRigid] …
// [MethodAssemblyJointsAddBall]): build the joint between joint origins A and B. Flip
// reverses the primary-axis sense (which way the components face).
type AddJointArgs struct {
	A    ConstraintGeomRef `json:"a"`
	B    ConstraintGeomRef `json:"b"`
	Flip bool              `json:"flip,omitempty"`
}

// DeleteJointArgs is the request of [MethodAssemblyJointsDelete]: remove the joint with id ID.
type DeleteJointArgs struct {
	ID uint64 `json:"id"`
}

// SetJointLimitsArgs is the request of [MethodAssemblyJointsSetLimits]: set the linear/angular
// bounds of the joint with id ID.
type SetJointLimitsArgs struct {
	ID     uint64      `json:"id"`
	Limits JointLimits `json:"limits"`
}

// SetJointFlipArgs is the request of [MethodAssemblyJointsSetFlip]: set the flip sense of the
// joint with id ID.
type SetJointFlipArgs struct {
	ID   uint64 `json:"id"`
	Flip bool   `json:"flip"`
}

// DSDOFInfo is one degree of freedom of a DS joint: whether it is translational or
// rotational, how its motion is imposed (free/driven/locked), and its current value
// (cm or radians).
type DSDOFInfo struct {
	Rotational    bool    `json:"rotational"`
	ImposedMotion string  `json:"imposedMotion"`
	Value         float64 `json:"value,omitempty"`
}

// DSJointInfo is one row of the DS-joint set: its id, kind, name, and per-DOF imposed-motion
// breakdown — the kinematic (degrees-of-freedom) view of a joint that motion/simulation
// consumers read.
type DSJointInfo struct {
	ID               uint64            `json:"id"`
	Type             string            `json:"type"`
	Name             string            `json:"name"`
	A                ConstraintGeomRef `json:"a"`
	B                ConstraintGeomRef `json:"b"`
	DegreesOfFreedom []DSDOFInfo       `json:"degreesOfFreedom"`
}

// DSJointsResult is the reply of [MethodDSJointsList]: the DS-joint set.
type DSJointsResult struct {
	Joints []DSJointInfo `json:"joints"`
}

// DSJointResult is the reply of the single DS-joint operations.
type DSJointResult struct {
	Joint DSJointInfo `json:"joint"`
}

// AddDSJointArgs is the request of [MethodDSJointsAdd]: add a DS joint of the given Type
// (a DSJointType string) between origins A and B.
type AddDSJointArgs struct {
	A    ConstraintGeomRef `json:"a"`
	B    ConstraintGeomRef `json:"b"`
	Type string            `json:"type"`
}

// SetImposedMotionArgs is the request of [MethodDSJointsSetImposedMotion]: set the imposed
// motion (free/driven/locked) and value of the DS joint's DOF at index DOFIndex.
type SetImposedMotionArgs struct {
	ID            uint64  `json:"id"`
	DOFIndex      int     `json:"dofIndex"`
	ImposedMotion string  `json:"imposedMotion"`
	Value         float64 `json:"value,omitempty"`
}

// DeleteDSJointArgs is the request of [MethodDSJointsDelete]: remove the DS joint with id ID.
type DeleteDSJointArgs struct {
	ID uint64 `json:"id"`
}
