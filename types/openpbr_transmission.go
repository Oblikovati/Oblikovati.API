// SPDX-License-Identifier: Apache-2.0

package types

// OpenPBRTransmission groups OpenPBR Surface's Transmission parameter group (spec
// §Transmission) — dielectric refraction through the Base layer, with volumetric
// scattering/absorption and physical dispersion. Depth is in scene length units; norm
// range [0,1]. Abbe number has norm range [9,91] (its full range is (0,∞)). Field defaults
// mirror the spec's table exactly; see [DefaultOpenPBRTransmission].
type OpenPBRTransmission struct {
	Weight               float32 `json:"weight" yaml:"weight"`                             // transmission_weight, [0,1]
	Color                Color3  `json:"color" yaml:"color"`                               // transmission_color, [0,1]^3
	Depth                float32 `json:"depth" yaml:"depth"`                               // transmission_depth, [0,∞), length, norm [0,1]
	Scatter              Color3  `json:"scatter" yaml:"scatter"`                           // transmission_scatter, [0,1]^3
	ScatterAnisotropy    float32 `json:"scatterAnisotropy" yaml:"scatterAnisotropy"`       // transmission_scatter_anisotropy, [-1,1]
	DispersionScale      float32 `json:"dispersionScale" yaml:"dispersionScale"`           // transmission_dispersion_scale, [0,1]
	DispersionAbbeNumber float32 `json:"dispersionAbbeNumber" yaml:"dispersionAbbeNumber"` // transmission_dispersion_abbe_number, (0,∞), norm [9,91]
}

// DefaultOpenPBRTransmission returns the spec's default Transmission group: zero weight
// (opaque by default), white transmission, no scatter/dispersion, Abbe number 20.
func DefaultOpenPBRTransmission() OpenPBRTransmission {
	return OpenPBRTransmission{
		Weight:               0,
		Color:                Color3{R: 1, G: 1, B: 1},
		Depth:                0,
		Scatter:              Color3{R: 0, G: 0, B: 0},
		ScatterAnisotropy:    0,
		DispersionScale:      0,
		DispersionAbbeNumber: 20,
	}
}
