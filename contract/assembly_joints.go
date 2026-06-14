// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The scalar read surface of the assembly joint set (M12-F02, Oblikovati/Oblikovati#359/
// #364). An in-proc consumer reads a joint's kind, free DOF, flip, limits, and health
// directly; every mutation travels over api/wire (assemblyJoints.*). The host
// implementations live in /source (model/assembly).

// JointLimits bounds a joint's driven values: a linear range and an angular range, each
// bound independently optional.
type JointLimits interface {
	// LinearMinimum returns the lower linear bound (cm) and whether it is set.
	LinearMinimum() (float64, bool)
	// LinearMaximum returns the upper linear bound (cm) and whether it is set.
	LinearMaximum() (float64, bool)
	// AngularMinimum returns the lower angular bound (radians) and whether it is set.
	AngularMinimum() (float64, bool)
	// AngularMaximum returns the upper angular bound (radians) and whether it is set.
	AngularMaximum() (float64, bool)
}

// AssemblyJoint is the read surface of one joint: its session id, kind, name, flip sense,
// the free DOF it leaves, limits, and health.
type AssemblyJoint interface {
	// ID is the joint's session id, unique within its assembly's set.
	ID() uint64
	// Type is the joint kind.
	Type() types.AssemblyJointType
	// Name is the joint's display name (e.g. "Rotational:1").
	Name() string
	// Flip reports the reversed primary-axis sense.
	Flip() bool
	// DegreesOfFreedom is the number of free DOF the joint leaves between the two occurrences.
	DegreesOfFreedom() int
	// Health reports whether the joint is fully evaluated, or sick (lost geometry).
	Health() types.HealthStatus
	// Suppressed reports whether the joint is excluded from the solve.
	Suppressed() bool
	// Limits returns the joint's driven-value bounds, or nil when unbounded.
	Limits() JointLimits
}

// AssemblyJointDefinition is the read surface of a joint's definition — the kind and the
// origin-definition kinds of its two joint origins (how each frame is built).
type AssemblyJointDefinition interface {
	// JointType is the kind of joint this definition builds.
	JointType() types.AssemblyJointType
	// OriginTypes returns how each of the two joint origins is defined (geometry kind).
	OriginTypes() (a, b types.AssemblyJointOriginDefinitionType)
}

// AssemblyJoints is the assembly's joint collection in creation order
// (host: assembly.JointSet).
type AssemblyJoints interface {
	// Count returns the number of joints in the set.
	Count() int
	// Item returns the joint at index i (0-based), or nil when out of range.
	Item(i int) AssemblyJoint
}

// AssemblyJointProxy is a joint viewed in the context of a specific occurrence path — the
// reference API's per-instance joint proxy, so a joint in a placed sub-assembly is
// addressable per placement.
type AssemblyJointProxy interface {
	AssemblyJoint
	// NativeJoint returns the underlying joint this proxy views.
	NativeJoint() AssemblyJoint
}

// AssemblyJointsEnumerator is the per-occurrence view of the joints that reference one
// occurrence (the reference API's per-component Joints collection).
type AssemblyJointsEnumerator interface {
	// Count returns the number of joints referencing the occurrence.
	Count() int
	// Item returns the i-th joint referencing the occurrence, or nil when out of range.
	Item(i int) AssemblyJoint
}

// DSDegreesOfFreedom is one degree of freedom of a DS joint: translational or rotational,
// its imposed-motion mode, and its current value.
type DSDegreesOfFreedom interface {
	// Rotational reports whether this DOF is rotational (false ⇒ translational).
	Rotational() bool
	// ImposedMotion reports how the DOF's motion is imposed (free/driven/locked).
	ImposedMotion() types.DSDOFImposedMotionType
	// Value is the DOF's current value (cm or radians).
	Value() float64
}

// DSJoint is the degrees-of-freedom / imposed-motion view of a joint — the kinematic
// surface motion and simulation consumers read.
type DSJoint interface {
	// ID is the DS joint's session id.
	ID() uint64
	// Type is the DS joint kind (mechanism vocabulary: prismatic/spherical/…).
	Type() types.DSJointType
	// Name is the DS joint's display name.
	Name() string
	// DOFCount returns the number of degrees of freedom.
	DOFCount() int
	// DOF returns the i-th degree of freedom, or nil when out of range.
	DOF(i int) DSDegreesOfFreedom
}

// DSJoints is the DS-joint collection (host: assembly.DSJointSet).
type DSJoints interface {
	// Count returns the number of DS joints.
	Count() int
	// Item returns the DS joint at index i, or nil when out of range.
	Item(i int) DSJoint
}

// DSJointDefinition is the read surface of a DS joint's definition (its kind).
type DSJointDefinition interface {
	// DSJointType is the kind of DS joint this definition builds.
	DSJointType() types.DSJointType
}
