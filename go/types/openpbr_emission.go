// SPDX-License-Identifier: Apache-2.0

package types

// OpenPBREmission groups OpenPBR Surface's Emission parameter group (spec §Emission) —
// self-emitted radiance, independent of every reflective/transmissive lobe. Luminance is
// in nits (cd/m²) with norm range [0,1000] though its full range is [0,∞); Color is
// unbounded above (unlike every other Color3 field in this package's OpenPBR groups).
// Field defaults mirror the spec's table exactly; see [DefaultOpenPBREmission].
type OpenPBREmission struct {
	Luminance float32 `json:"luminance" yaml:"luminance"` // emission_luminance, [0,∞), nits, norm [0,1000]
	Color     Color3  `json:"color" yaml:"color"`         // emission_color, [0,∞)^3
}

// DefaultOpenPBREmission returns the spec's default Emission group: zero luminance (no
// emission by default), white color.
func DefaultOpenPBREmission() OpenPBREmission {
	return OpenPBREmission{
		Luminance: 0,
		Color:     Color3{R: 1, G: 1, B: 1},
	}
}
