// SPDX-License-Identifier: Apache-2.0

package types

// MoveOperationType discriminates one entry in a Move feature's ordered operation list
// (M20·F20 parity). A move definition composes a sequence of independently parametric
// operations rather than a single baked transform, so "rotate 30° about this edge then
// slide 5 mm along it" round-trips as two separately driven operations. The reference
// API models these as typed operation objects (free-drag / move-along-ray /
// rotate-about-line) with no numeric enum; the string values are the stable wire
// vocabulary — treat them as frozen.
//
// This is the canonical, Apache-2.0 definition; the GPL implementation maps each to its
// model operation builder.
type MoveOperationType string

const (
	// MoveFreeDrag offsets the body by explicit X/Y/Z distances.
	MoveFreeDrag MoveOperationType = "freeDrag"
	// MoveAlongRay slides the body a distance along a direction entity.
	MoveAlongRay MoveOperationType = "alongRay"
	// MoveRotateAboutLine rotates the body by an angle about an axis entity.
	MoveRotateAboutLine MoveOperationType = "rotateAboutLine"
)

// IsValid reports whether t is a defined move-operation type.
func (t MoveOperationType) IsValid() bool {
	switch t {
	case MoveFreeDrag, MoveAlongRay, MoveRotateAboutLine:
		return true
	default:
		return false
	}
}

// AllMoveOperationTypes returns every move-operation type in definition order — the
// source list for a builder or a schema enum.
func AllMoveOperationTypes() []MoveOperationType {
	return []MoveOperationType{MoveFreeDrag, MoveAlongRay, MoveRotateAboutLine}
}
