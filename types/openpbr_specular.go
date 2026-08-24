// SPDX-License-Identifier: Apache-2.0

package types

// OpenPBRSpecular groups OpenPBR Surface's Specular parameter group (spec §Specular) — the
// dielectric microfacet lobe layered over Base. Weight has norm range [0,1] though its full
// range is unbounded; IOR has norm range [1,3] though its full range is (0,∞). Field
// defaults mirror the spec's table exactly; see [DefaultOpenPBRSpecular].
type OpenPBRSpecular struct {
	Weight              float32 `json:"weight" yaml:"weight"`                           // specular_weight, [0,∞), norm [0,1]
	Color               Color3  `json:"color" yaml:"color"`                             // specular_color, [0,1]^3
	Roughness           float32 `json:"roughness" yaml:"roughness"`                     // specular_roughness, [0,1]
	RoughnessAnisotropy float32 `json:"roughnessAnisotropy" yaml:"roughnessAnisotropy"` // specular_roughness_anisotropy, [0,1]
	IOR                 float32 `json:"ior" yaml:"ior"`                                 // specular_ior, (0,∞), norm [1,3]
}

// DefaultOpenPBRSpecular returns the spec's default Specular group: full-weight white
// specular, 0.3 roughness, no anisotropy, IOR 1.5 (typical dielectric).
func DefaultOpenPBRSpecular() OpenPBRSpecular {
	return OpenPBRSpecular{
		Weight:              1,
		Color:               Color3{R: 1, G: 1, B: 1},
		Roughness:           0.3,
		RoughnessAnisotropy: 0,
		IOR:                 1.5,
	}
}
