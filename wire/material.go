// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// AppearanceInfo is the JSON shape of a PBR appearance. Albedo and Emissive are
// "#RRGGBBAA" hex (compact, readable, matching the on-disk and theme conventions); the
// scalar PBR terms are in [0,1].
type AppearanceInfo struct {
	ID          string  `json:"id"`
	DisplayName string  `json:"displayName"`
	Source      string  `json:"source"`
	Albedo      string  `json:"albedo"`
	Metallic    float32 `json:"metallic"`
	Roughness   float32 `json:"roughness"`
	Emissive    string  `json:"emissive"`
	Opacity     float32 `json:"opacity"`
}

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
	// IsotropyClass is "isotropic" (or empty), "orthotropic", or "transversely-isotropic".
	IsotropyClass string `json:"isotropyClass,omitempty"`
	// Anisotropic carries the direction-dependent elastic constants when IsotropyClass is
	// not isotropic; zero-valued otherwise.
	Anisotropic  types.AnisotropicElastic `json:"anisotropic"`
	AppearanceID string                   `json:"appearanceId"`
}

// ListAppearancesResult / ListMaterialsResult are the list responses.
type ListAppearancesResult struct {
	Appearances []AppearanceInfo `json:"appearances"`
}

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

// AssignAppearanceArgs overrides the appearance at a scope ("part", "body", or "face").
// Key is the hex reference key of the body/face (empty for the part default).
type AssignAppearanceArgs struct {
	Scope        string `json:"scope"`
	Key          string `json:"key,omitempty"`
	AppearanceID string `json:"appearanceId"`
}
