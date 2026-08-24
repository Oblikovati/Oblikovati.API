// SPDX-License-Identifier: Apache-2.0

package types

// OpenPBRThinFilm groups OpenPBR Surface's Thin-film parameter group (spec §Thin-film) —
// wavelength-dependent iridescence from a microscopically thin dielectric film over
// Specular/Base (Airy-summation interference). Thickness is in micrometres, norm range
// [0,1]. IOR has norm range [1,3] though its full range is (0,∞). Field defaults mirror
// the spec's table exactly; see [DefaultOpenPBRThinFilm].
type OpenPBRThinFilm struct {
	Weight    float32 `json:"weight" yaml:"weight"`       // thin_film_weight, [0,1]
	Thickness float32 `json:"thickness" yaml:"thickness"` // thin_film_thickness, [0,∞), µm, norm [0,1]
	IOR       float32 `json:"ior" yaml:"ior"`             // thin_film_ior, (0,∞), norm [1,3]
}

// DefaultOpenPBRThinFilm returns the spec's default Thin-film group: zero weight
// (disabled by default), 0.5µm thickness, IOR 1.4.
func DefaultOpenPBRThinFilm() OpenPBRThinFilm {
	return OpenPBRThinFilm{
		Weight:    0,
		Thickness: 0.5,
		IOR:       1.4,
	}
}
