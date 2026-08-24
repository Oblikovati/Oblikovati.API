// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// Appearance is the in-process contract for one appearance described by the full
// OpenPBR Surface v1.1.1 specification (github.com/AcademySoftwareFoundation/OpenPBR) —
// every lobe (Base, Specular, Transmission, Subsurface, Coat, Fuzz, Thin-film, Emission,
// Geometry). The GPL model/material.Appearance satisfies it (compile-time asserted
// there).
//
// Colors are [types.Color3] in the ACEScg working space (OpenPBR's default).
//
// Example — reading the Base group's albedo for a renderer surface:
//
//	base := appearance.Base()
//	albedo := base.Color // types.Color3, ACEScg
type Appearance interface {
	// ID is the stable identity used by assignments and library lookups.
	ID() string
	// DisplayName is the label shown in the appearance browser.
	DisplayName() string
	// Source says whether this is a built-in, project, or document-embedded asset.
	Source() types.AssetSource

	// Base is the diffuse/metal foundation lobe (spec §Base).
	Base() types.OpenPBRBase
	// Specular is the dielectric microfacet lobe over Base (spec §Specular).
	Specular() types.OpenPBRSpecular
	// Transmission is refraction through Base, with scatter and dispersion (spec
	// §Transmission).
	Transmission() types.OpenPBRTransmission
	// Subsurface is volumetric subsurface scattering coupled to Base (spec §Subsurface).
	Subsurface() types.OpenPBRSubsurface
	// Coat is the clear dielectric layer above every other lobe (spec §Coat).
	Coat() types.OpenPBRCoat
	// Fuzz is the sheen/velvet retroreflective lobe above Coat (spec §Fuzz).
	Fuzz() types.OpenPBRFuzz
	// ThinFilm is wavelength-dependent iridescence over Specular/Base (spec §Thin-film).
	ThinFilm() types.OpenPBRThinFilm
	// Emission is self-emitted radiance, independent of every other lobe (spec
	// §Emission).
	Emission() types.OpenPBREmission
	// Geometry is the surface-level override group (opacity, thin-walled, normal/tangent
	// perturbation; spec §Geometry).
	Geometry() types.OpenPBRGeometry
}
