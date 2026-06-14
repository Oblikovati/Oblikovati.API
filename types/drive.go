// SPDX-License-Identifier: Apache-2.0

package types

// The drive vocabulary (M12-F03, Oblikovati/Oblikovati#366): driving a constraint or joint
// sweeps one of its driven variables through a value range, re-solving at each step so the
// assembly animates (a kinematic motion study). These are the canonical Apache-2.0
// definitions; the GPL solver (model/assembly) aliases them.
//
// Values are STABLE ACROSS SESSIONS and must never be renumbered: a drive recipe persists
// which variable it sweeps.

// DriveVariable selects which of a relationship's driven variables a drive sweeps. A
// rotational joint has one angular variable; a slider one linear; a cylindrical joint both,
// so the caller must pick. The zero value asks for the relationship's natural variable.
type DriveVariable uint32

const (
	// DriveNatural sweeps the relationship's single natural driven variable (the rotation of
	// a rotational joint, the translation of a slider). Ambiguous for multi-DOF kinds.
	DriveNatural DriveVariable = 0
	// DriveAngular sweeps a rotation about the joint axis (radians).
	DriveAngular DriveVariable = 1
	// DriveLinear sweeps a translation along the joint axis (centimetres).
	DriveLinear DriveVariable = 2
)

// IsValid reports whether v names a real selector value.
func (v DriveVariable) IsValid() bool { return v <= DriveLinear }

// String returns a stable lowercase name. The value, not this name, is the persisted identity.
func (v DriveVariable) String() string {
	switch v {
	case DriveNatural:
		return "natural"
	case DriveAngular:
		return "angular"
	case DriveLinear:
		return "linear"
	default:
		return "unknown"
	}
}
