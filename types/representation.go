// SPDX-License-Identifier: Apache-2.0

package types

// The representation vocabulary (M12-F04, Oblikovati/Oblikovati#361/#367): the three assembly
// representation families — design-view (visibility/appearance/section/camera), positional
// (constraint/joint value overrides) and level-of-detail (component suppression) — plus model
// states that select one of each. Representations are named override layers over an immutable
// base assembly. These are the canonical Apache-2.0 definitions; the GPL host (model/assembly)
// aliases them.
//
// Values are STABLE ACROSS SESSIONS and must never be renumbered: a model recipe persists a
// representation's kind.

// RepresentationKind discriminates the three representation families.
type RepresentationKind uint32

const (
	// RepresentationUnknown is the zero value: an unresolved or not-yet-typed representation.
	RepresentationUnknown RepresentationKind = 0
	// RepresentationDesignView overrides visibility, appearance, section planes, and the camera.
	RepresentationDesignView RepresentationKind = 1
	// RepresentationPositional overrides constraint/joint values (and per-occurrence flexibility).
	RepresentationPositional RepresentationKind = 2
	// RepresentationLevelOfDetail suppresses occurrences for performance on large assemblies.
	RepresentationLevelOfDetail RepresentationKind = 3
)

// IsValid reports whether k names a real representation family (not unknown).
func (k RepresentationKind) IsValid() bool {
	return k >= RepresentationDesignView && k <= RepresentationLevelOfDetail
}

// String returns a stable lowercase name. The value, not this name, is the persisted identity.
func (k RepresentationKind) String() string {
	switch k {
	case RepresentationDesignView:
		return "designView"
	case RepresentationPositional:
		return "positional"
	case RepresentationLevelOfDetail:
		return "levelOfDetail"
	default:
		return "unknown"
	}
}

// SectionPlane is a clipping plane carried by a design-view representation: a point on the
// plane and a normal. When the representation is active the half-space on the normal side is
// clipped away (Flipped swaps which side is removed), revealing the model's interior.
type SectionPlane struct {
	Origin  Point  `json:"origin"`
	Normal  Vector `json:"normal"`
	Flipped bool   `json:"flipped,omitempty"`
}
