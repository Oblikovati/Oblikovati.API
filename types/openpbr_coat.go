// SPDX-License-Identifier: Apache-2.0

package types

// OpenPBRCoat groups OpenPBR Surface's Coat parameter group (spec §Coat) — a clear
// dielectric layer above everything else (Base, Specular, Transmission, Subsurface). IOR
// has norm range [1,3] though its full range is (0,∞). Field defaults mirror the spec's
// table exactly; see [DefaultOpenPBRCoat].
type OpenPBRCoat struct {
	Weight              float32 `json:"weight" yaml:"weight"`                           // coat_weight, [0,1]
	Color               Color3  `json:"color" yaml:"color"`                             // coat_color, [0,1]^3
	Roughness           float32 `json:"roughness" yaml:"roughness"`                     // coat_roughness, [0,1]
	RoughnessAnisotropy float32 `json:"roughnessAnisotropy" yaml:"roughnessAnisotropy"` // coat_roughness_anisotropy, [0,1]
	IOR                 float32 `json:"ior" yaml:"ior"`                                 // coat_ior, (0,∞), norm [1,3]
	Darkening           float32 `json:"darkening" yaml:"darkening"`                     // coat_darkening, [0,1]
}

// DefaultOpenPBRCoat returns the spec's default Coat group: zero weight (disabled by
// default), white coat, no roughness/anisotropy, IOR 1.6, full darkening.
func DefaultOpenPBRCoat() OpenPBRCoat {
	return OpenPBRCoat{
		Weight:              0,
		Color:               Color3{R: 1, G: 1, B: 1},
		Roughness:           0,
		RoughnessAnisotropy: 0,
		IOR:                 1.6,
		Darkening:           1,
	}
}
