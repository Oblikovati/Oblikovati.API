// SPDX-License-Identifier: Apache-2.0

package types

// OpenPBRBase groups OpenPBR Surface's Base parameter group (spec §Base,
// parametrization.md.html) — the diffuse/metal foundation lobe every other layer sits on
// top of. Field defaults mirror the spec's table exactly; see [DefaultOpenPBRBase].
type OpenPBRBase struct {
	Weight           float32 `json:"weight" yaml:"weight"`                     // base_weight, [0,1]
	Color            Color3  `json:"color" yaml:"color"`                       // base_color, [0,1]^3
	Metalness        float32 `json:"metalness" yaml:"metalness"`               // base_metalness, [0,1]
	DiffuseRoughness float32 `json:"diffuseRoughness" yaml:"diffuseRoughness"` // base_diffuse_roughness, [0,1]
}

// DefaultOpenPBRBase returns the spec's default Base group: fully-weighted 80%-grey
// diffuse, zero metalness, zero (Lambertian) diffuse roughness.
func DefaultOpenPBRBase() OpenPBRBase {
	return OpenPBRBase{
		Weight:           1,
		Color:            Color3{R: 0.8, G: 0.8, B: 0.8},
		Metalness:        0,
		DiffuseRoughness: 0,
	}
}
