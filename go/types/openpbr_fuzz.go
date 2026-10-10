// SPDX-License-Identifier: Apache-2.0

package types

// OpenPBRFuzz groups OpenPBR Surface's Fuzz parameter group (spec §Fuzz) — a
// sheen/velvet-like retroreflective lobe (Zeltner sheen BRDF) layered above Coat. Field
// defaults mirror the spec's table exactly; see [DefaultOpenPBRFuzz].
type OpenPBRFuzz struct {
	Weight    float32 `json:"weight" yaml:"weight"`       // fuzz_weight, [0,1]
	Color     Color3  `json:"color" yaml:"color"`         // fuzz_color, [0,1]^3
	Roughness float32 `json:"roughness" yaml:"roughness"` // fuzz_roughness, [0,1]
}

// DefaultOpenPBRFuzz returns the spec's default Fuzz group: zero weight (disabled by
// default), white fuzz, mid roughness.
func DefaultOpenPBRFuzz() OpenPBRFuzz {
	return OpenPBRFuzz{
		Weight:    0,
		Color:     Color3{R: 1, G: 1, B: 1},
		Roughness: 0.5,
	}
}
