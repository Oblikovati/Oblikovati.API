// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// Material is the in-process contract for one physical-world material: its density and
// mechanical/thermal/electrical properties, plus the appearance it renders with. The GPL
// model/material.Material satisfies it (compile-time asserted there).
//
// A material references its appearance by id ([Material.AppearanceID]); the appearance
// itself is a separate asset, so several materials can share one look and an appearance
// can be edited independently.
type Material interface {
	// ID is the stable identity used by assignments and library lookups.
	ID() string
	// DisplayName is the label shown in the material browser.
	DisplayName() string
	// Source says whether this is a built-in, project, or document-embedded asset.
	Source() types.AssetSource
	// Density is the material density in g/cm³ (mass = density × volume).
	Density() float64
	// Mechanical returns the structural properties.
	Mechanical() types.Mechanical
	// Thermal returns the heat-related properties.
	Thermal() types.Thermal
	// Electrical returns the electrical properties.
	Electrical() types.Electrical
	// IsotropyClass declares the material's elastic symmetry. An isotropic material is
	// fully described by Mechanical; orthotropic / transversely-isotropic materials also
	// carry Anisotropic. Never returns the empty string — an unset class reports Isotropic.
	IsotropyClass() types.IsotropyClass
	// Anisotropic returns the direction-dependent elastic constants, or the zero value for
	// an isotropic material. Meaningful only when IsotropyClass is not Isotropic.
	Anisotropic() types.AnisotropicElastic
	// AppearanceID is the id of the appearance this material renders with.
	AppearanceID() string
}
