// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// MaterialInfo is the JSON shape of a material: identity, density, the property groups,
// and the id of the appearance it renders with.
type MaterialInfo struct {
	ID          string           `json:"id"`
	DisplayName string           `json:"displayName"`
	Source      string           `json:"source"`
	Density     float64          `json:"density"`
	Mechanical  types.Mechanical `json:"mechanical"`
	Thermal     types.Thermal    `json:"thermal"`
	Electrical  types.Electrical `json:"electrical"`
	// Magnetic carries the magnetostatics constitutive data (μr, Br, Hc, Bsat) for
	// soft-magnetic cores and permanent magnets; the zero value is a non-magnetic material.
	Magnetic types.Magnetic `json:"magnetic"`
	// IsotropyClass is "isotropic" (or empty), "orthotropic", or "transversely-isotropic".
	IsotropyClass string `json:"isotropyClass,omitempty"`
	// Anisotropic carries the direction-dependent elastic constants when IsotropyClass is
	// not isotropic; zero-valued otherwise.
	Anisotropic  types.AnisotropicElastic `json:"anisotropic"`
	AppearanceID string                   `json:"appearanceId"`
}

// ListMaterialsResult is the list response.
type ListMaterialsResult struct {
	Materials []MaterialInfo `json:"materials"`
}

// AssetRefArgs identifies an asset by id ([MethodAppearancesGet]/[MethodMaterialsGet]).
type AssetRefArgs struct {
	ID string `json:"id"`
}

// DuplicateAssetArgs creates a custom asset by copying an existing one under a new name
// ([MethodAppearancesCreate]/[MethodMaterialsCreate]).
type DuplicateAssetArgs struct {
	BaseID string `json:"baseId"`
	Name   string `json:"name"`
}

// AssignMaterialArgs assigns a material to a body, or to the part default when BodyKey is
// empty. BodyKey is the body's hex reference key (stable across recompute).
type AssignMaterialArgs struct {
	BodyKey    string `json:"bodyKey,omitempty"`
	MaterialID string `json:"materialId"`
}
