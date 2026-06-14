// SPDX-License-Identifier: Apache-2.0

package types

// The assembly-joint vocabulary (M12-F02, Oblikovati/Oblikovati#359/#364): the simplified
// joint model — one relationship that establishes a degree-of-freedom set between two
// occurrences — plus the DS-joint (degrees-of-freedom / imposed-motion) surface. These are
// the canonical, Apache-2.0 definitions; the GPL solver (model/assembly) aliases them.
//
// Values are STABLE ACROSS SESSIONS and must never be renumbered: an assembly recipe
// persists a joint's kind, its origin-definition kind, and a DS DOF's imposed-motion kind.

// AssemblyJointType discriminates the standard joint kinds, ordered by remaining freedom:
// rigid (0 DOF) through ball (3 rotational DOF). A joint reduces the relative placement of
// two occurrences to exactly its degrees of freedom.
type AssemblyJointType uint32

const (
	// JointUnknown is the zero value: an unresolved or not-yet-typed joint.
	JointUnknown AssemblyJointType = 0
	// JointRigid fully fixes the two occurrences together (0 DOF).
	JointRigid AssemblyJointType = 1
	// JointRotational allows one rotation about the joint axis (1 DOF).
	JointRotational AssemblyJointType = 2
	// JointSlider allows one translation along the joint axis (1 DOF).
	JointSlider AssemblyJointType = 3
	// JointCylindrical allows translation along and rotation about the axis (2 DOF).
	JointCylindrical AssemblyJointType = 4
	// JointPlanar allows two translations in a plane and one rotation about its normal (3 DOF).
	JointPlanar AssemblyJointType = 5
	// JointBall allows three rotations about a common point (3 DOF).
	JointBall AssemblyJointType = 6
)

// IsValid reports whether t names a real joint kind (not unknown).
func (t AssemblyJointType) IsValid() bool { return t >= JointRigid && t <= JointBall }

// String returns a stable lowercase name. The value, not this name, is the persisted identity.
func (t AssemblyJointType) String() string {
	switch t {
	case JointRigid:
		return "rigid"
	case JointRotational:
		return "rotational"
	case JointSlider:
		return "slider"
	case JointCylindrical:
		return "cylindrical"
	case JointPlanar:
		return "planar"
	case JointBall:
		return "ball"
	default:
		return "unknown"
	}
}

// DegreesOfFreedom returns the number of free DOF a joint of this kind leaves between the
// two occurrences (rigid 0 … ball 3) — the value the solver's rank analysis must produce.
func (t AssemblyJointType) DegreesOfFreedom() int {
	switch t {
	case JointRotational, JointSlider:
		return 1
	case JointCylindrical:
		return 2
	case JointPlanar, JointBall:
		return 3
	default:
		return 0
	}
}

// AssemblyJointOriginDefinitionType discriminates how a joint origin (the frame the joint is
// built on, on each component) is defined: from a point, an edge/axis, or a planar face.
type AssemblyJointOriginDefinitionType uint32

const (
	// JointOriginUnknown is the zero value.
	JointOriginUnknown AssemblyJointOriginDefinitionType = 0
	// JointOriginPoint derives the origin frame from a vertex / work point.
	JointOriginPoint AssemblyJointOriginDefinitionType = 1
	// JointOriginEdge derives the origin frame from an edge / axis (its line).
	JointOriginEdge AssemblyJointOriginDefinitionType = 2
	// JointOriginPlane derives the origin frame from a planar face / work plane (its normal).
	JointOriginPlane AssemblyJointOriginDefinitionType = 3
)

// String returns a stable lowercase name.
func (t AssemblyJointOriginDefinitionType) String() string {
	switch t {
	case JointOriginPoint:
		return "point"
	case JointOriginEdge:
		return "edge"
	case JointOriginPlane:
		return "plane"
	default:
		return "unknown"
	}
}

// DSJointType discriminates the DS-joint kinds — the degrees-of-freedom view of a joint,
// named in the mechanism vocabulary (prismatic = slider, spherical = ball).
type DSJointType uint32

const (
	// DSJointUnknown is the zero value.
	DSJointUnknown DSJointType = 0
	// DSJointRigid fixes the two bodies (0 DOF).
	DSJointRigid DSJointType = 1
	// DSJointRotational allows one rotation (1 DOF).
	DSJointRotational DSJointType = 2
	// DSJointPrismatic allows one translation (1 DOF) — the slider.
	DSJointPrismatic DSJointType = 3
	// DSJointCylindrical allows one translation and one rotation (2 DOF).
	DSJointCylindrical DSJointType = 4
	// DSJointPlanar allows two translations and one rotation (3 DOF).
	DSJointPlanar DSJointType = 5
	// DSJointSpherical allows three rotations (3 DOF) — the ball.
	DSJointSpherical DSJointType = 6
)

// String returns a stable lowercase name.
func (t DSJointType) String() string {
	switch t {
	case DSJointRigid:
		return "rigid"
	case DSJointRotational:
		return "rotational"
	case DSJointPrismatic:
		return "prismatic"
	case DSJointCylindrical:
		return "cylindrical"
	case DSJointPlanar:
		return "planar"
	case DSJointSpherical:
		return "spherical"
	default:
		return "unknown"
	}
}

// DSDOFImposedMotionType discriminates how a single DS degree of freedom is driven: free to
// move, driven to an imposed value (a motor / drive, M12-F03), or locked at its value.
type DSDOFImposedMotionType uint32

const (
	// DSDOFFree: the DOF moves freely (no imposed motion).
	DSDOFFree DSDOFImposedMotionType = 0
	// DSDOFDriven: the DOF is driven to an imposed value (a motor).
	DSDOFDriven DSDOFImposedMotionType = 1
	// DSDOFLocked: the DOF is held fixed at its current value.
	DSDOFLocked DSDOFImposedMotionType = 2
)

// String returns a stable lowercase name.
func (t DSDOFImposedMotionType) String() string {
	switch t {
	case DSDOFDriven:
		return "driven"
	case DSDOFLocked:
		return "locked"
	default:
		return "free"
	}
}
