// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestDefaultOpenPBRSurfaceParams asserts every OpenPBR Surface v1.1.1 parameter
// (parametrization.md.html) has a typed field carrying the spec's default value —
// one row per Identifier in the spec's table, in the spec's own order.
func TestDefaultOpenPBRSurfaceParams(t *testing.T) {
	p := struct {
		Base         OpenPBRBase
		Specular     OpenPBRSpecular
		Transmission OpenPBRTransmission
		Subsurface   OpenPBRSubsurface
		Coat         OpenPBRCoat
		Fuzz         OpenPBRFuzz
		ThinFilm     OpenPBRThinFilm
		Emission     OpenPBREmission
		Geometry     OpenPBRGeometry
	}{
		Base: DefaultOpenPBRBase(), Specular: DefaultOpenPBRSpecular(),
		Transmission: DefaultOpenPBRTransmission(), Subsurface: DefaultOpenPBRSubsurface(),
		Coat: DefaultOpenPBRCoat(), Fuzz: DefaultOpenPBRFuzz(), ThinFilm: DefaultOpenPBRThinFilm(),
		Emission: DefaultOpenPBREmission(), Geometry: DefaultOpenPBRGeometry(),
	}

	cases := []struct {
		identifier string
		got, want  any
	}{
		{"base_weight", p.Base.Weight, float32(1)},
		{"base_color", p.Base.Color, Color3{0.8, 0.8, 0.8}},
		{"base_metalness", p.Base.Metalness, float32(0)},
		{"base_diffuse_roughness", p.Base.DiffuseRoughness, float32(0)},

		{"specular_weight", p.Specular.Weight, float32(1)},
		{"specular_color", p.Specular.Color, Color3{1, 1, 1}},
		{"specular_roughness", p.Specular.Roughness, float32(0.3)},
		{"specular_roughness_anisotropy", p.Specular.RoughnessAnisotropy, float32(0)},
		{"specular_ior", p.Specular.IOR, float32(1.5)},

		{"transmission_weight", p.Transmission.Weight, float32(0)},
		{"transmission_color", p.Transmission.Color, Color3{1, 1, 1}},
		{"transmission_depth", p.Transmission.Depth, float32(0)},
		{"transmission_scatter", p.Transmission.Scatter, Color3{0, 0, 0}},
		{"transmission_scatter_anisotropy", p.Transmission.ScatterAnisotropy, float32(0)},
		{"transmission_dispersion_scale", p.Transmission.DispersionScale, float32(0)},
		{"transmission_dispersion_abbe_number", p.Transmission.DispersionAbbeNumber, float32(20)},

		{"subsurface_weight", p.Subsurface.Weight, float32(0)},
		{"subsurface_color", p.Subsurface.Color, Color3{0.8, 0.8, 0.8}},
		{"subsurface_radius", p.Subsurface.Radius, float32(1)},
		{"subsurface_radius_scale", p.Subsurface.RadiusScale, Color3{1.0, 0.5, 0.25}},
		{"subsurface_scatter_anisotropy", p.Subsurface.ScatterAnisotropy, float32(0)},

		{"coat_weight", p.Coat.Weight, float32(0)},
		{"coat_color", p.Coat.Color, Color3{1, 1, 1}},
		{"coat_roughness", p.Coat.Roughness, float32(0)},
		{"coat_roughness_anisotropy", p.Coat.RoughnessAnisotropy, float32(0)},
		{"coat_ior", p.Coat.IOR, float32(1.6)},
		{"coat_darkening", p.Coat.Darkening, float32(1)},

		{"fuzz_weight", p.Fuzz.Weight, float32(0)},
		{"fuzz_color", p.Fuzz.Color, Color3{1, 1, 1}},
		{"fuzz_roughness", p.Fuzz.Roughness, float32(0.5)},

		{"emission_luminance", p.Emission.Luminance, float32(0)},
		{"emission_color", p.Emission.Color, Color3{1, 1, 1}},

		{"thin_film_weight", p.ThinFilm.Weight, float32(0)},
		{"thin_film_thickness", p.ThinFilm.Thickness, float32(0.5)},
		{"thin_film_ior", p.ThinFilm.IOR, float32(1.4)},

		{"geometry_opacity", p.Geometry.Opacity, float32(1)},
		{"geometry_thin_walled", p.Geometry.ThinWalled, false},
	}

	for _, c := range cases {
		t.Run(c.identifier, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("%s default = %v, want %v", c.identifier, c.got, c.want)
			}
		})
	}

	if p.Geometry.Normal != nil || p.Geometry.Tangent != nil ||
		p.Geometry.CoatNormal != nil || p.Geometry.CoatTangent != nil {
		t.Errorf("geometry_normal/tangent/coat_normal/coat_tangent default = non-nil, want nil (unperturbed)")
	}
}
