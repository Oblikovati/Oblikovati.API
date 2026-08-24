// SPDX-License-Identifier: Apache-2.0

package types

// OpenPBRSubsurface groups OpenPBR Surface's Subsurface parameter group (spec
// §Subsurface) — volumetric subsurface scattering coupled to the Base diffuse lobe.
// Radius is in scene length units; norm range [0,1]. Field defaults mirror the spec's
// table exactly; see [DefaultOpenPBRSubsurface].
type OpenPBRSubsurface struct {
	Weight            float32 `json:"weight" yaml:"weight"`                       // subsurface_weight, [0,1]
	Color             Color3  `json:"color" yaml:"color"`                         // subsurface_color, [0,1]^3
	Radius            float32 `json:"radius" yaml:"radius"`                       // subsurface_radius, [0,∞), length, norm [0,1]
	RadiusScale       Color3  `json:"radiusScale" yaml:"radiusScale"`             // subsurface_radius_scale, [0,1]^3
	ScatterAnisotropy float32 `json:"scatterAnisotropy" yaml:"scatterAnisotropy"` // subsurface_scatter_anisotropy, [-1,1]
}

// DefaultOpenPBRSubsurface returns the spec's default Subsurface group: zero weight
// (disabled by default), 80%-grey color, unit radius, RGB radius falloff (1, 0.5, 0.25),
// no anisotropy.
func DefaultOpenPBRSubsurface() OpenPBRSubsurface {
	return OpenPBRSubsurface{
		Weight:            0,
		Color:             Color3{R: 0.8, G: 0.8, B: 0.8},
		Radius:            1,
		RadiusScale:       Color3{R: 1.0, G: 0.5, B: 0.25},
		ScatterAnisotropy: 0,
	}
}
