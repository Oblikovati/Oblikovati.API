// SPDX-License-Identifier: Apache-2.0

package types

// OpenPBRSurfaceParams bundles every OpenPBR Surface v1.1.1 parameter group
// (parametrization.md.html), grouped exactly as the spec groups them. It is the value
// type behind [contract.OpenPBRAppearance]'s grouped accessors and the
// OpenPBRAppearanceInfo wire DTO — the full lobe set, not the cut-down metallic-roughness
// subset [Appearance] carries.
type OpenPBRSurfaceParams struct {
	Base         OpenPBRBase
	Specular     OpenPBRSpecular
	Transmission OpenPBRTransmission
	Subsurface   OpenPBRSubsurface
	Coat         OpenPBRCoat
	Fuzz         OpenPBRFuzz
	ThinFilm     OpenPBRThinFilm
	Emission     OpenPBREmission
	Geometry     OpenPBRGeometry
}

// DefaultOpenPBRSurfaceParams returns the spec's default value for every parameter group
// — the value an OpenPBRAppearance has when no group has been customized.
func DefaultOpenPBRSurfaceParams() OpenPBRSurfaceParams {
	return OpenPBRSurfaceParams{
		Base:         DefaultOpenPBRBase(),
		Specular:     DefaultOpenPBRSpecular(),
		Transmission: DefaultOpenPBRTransmission(),
		Subsurface:   DefaultOpenPBRSubsurface(),
		Coat:         DefaultOpenPBRCoat(),
		Fuzz:         DefaultOpenPBRFuzz(),
		ThinFilm:     DefaultOpenPBRThinFilm(),
		Emission:     DefaultOpenPBREmission(),
		Geometry:     DefaultOpenPBRGeometry(),
	}
}
