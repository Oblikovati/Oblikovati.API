// SPDX-License-Identifier: Apache-2.0

package types

// OpenPBRGeometry groups OpenPBR Surface's Geometry parameter group (spec §Geometry) —
// surface-level overrides that apply regardless of which lobes are active. Normal/Tangent/
// CoatNormal/CoatTangent are nil by default ("unperturbed": use the surface's own
// geometric normal/tangent); a non-nil override is a per-point perturbation (e.g. a normal
// map), not a per-appearance constant, but the constant form is represented the same way.
// Field defaults mirror the spec's table exactly; see [DefaultOpenPBRGeometry].
type OpenPBRGeometry struct {
	Opacity     float32 `json:"opacity" yaml:"opacity"`                             // geometry_opacity, [0,1]
	ThinWalled  bool    `json:"thinWalled" yaml:"thinWalled"`                       // geometry_thin_walled
	Normal      *Vector `json:"normal,omitempty" yaml:"normal,omitempty"`           // geometry_normal
	Tangent     *Vector `json:"tangent,omitempty" yaml:"tangent,omitempty"`         // geometry_tangent
	CoatNormal  *Vector `json:"coatNormal,omitempty" yaml:"coatNormal,omitempty"`   // geometry_coat_normal
	CoatTangent *Vector `json:"coatTangent,omitempty" yaml:"coatTangent,omitempty"` // geometry_coat_tangent
}

// DefaultOpenPBRGeometry returns the spec's default Geometry group: fully opaque, not
// thin-walled, every normal/tangent unperturbed (nil).
func DefaultOpenPBRGeometry() OpenPBRGeometry {
	return OpenPBRGeometry{
		Opacity:    1,
		ThinWalled: false,
	}
}
