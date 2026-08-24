// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// OpenPBRAppearanceInfo is the JSON shape of a full OpenPBR Surface v1.1.1 appearance —
// additive alongside [AppearanceInfo]'s metallic-roughness subset. Every group is the
// [types] value type from PBI-335 (colors are ACEScg [types.Color3], not hex, since
// emission_color is unbounded above and hex cannot represent that).
type OpenPBRAppearanceInfo struct {
	ID           string                    `json:"id"`
	DisplayName  string                    `json:"displayName"`
	Source       string                    `json:"source"`
	Base         types.OpenPBRBase         `json:"base"`
	Specular     types.OpenPBRSpecular     `json:"specular"`
	Transmission types.OpenPBRTransmission `json:"transmission"`
	Subsurface   types.OpenPBRSubsurface   `json:"subsurface"`
	Coat         types.OpenPBRCoat         `json:"coat"`
	Fuzz         types.OpenPBRFuzz         `json:"fuzz"`
	ThinFilm     types.OpenPBRThinFilm     `json:"thinFilm"`
	Emission     types.OpenPBREmission     `json:"emission"`
	Geometry     types.OpenPBRGeometry     `json:"geometry"`
}

// ListOpenPBRAppearancesResult is the [MethodOpenPBRAppearancesList] response.
type ListOpenPBRAppearancesResult struct {
	Appearances []OpenPBRAppearanceInfo `json:"appearances"`
}

// CreateOpenPBRAppearanceArgs creates a custom OpenPBR appearance by copying an existing
// one under a new name ([MethodOpenPBRAppearancesCreate]), mirroring
// [DuplicateAssetArgs]'s shape for the existing appearances.create method.
type CreateOpenPBRAppearanceArgs struct {
	BaseID string `json:"baseId"`
	Name   string `json:"name"`
}

// UpdateOpenPBRAppearanceArgs replaces every group of an existing, editable OpenPBR
// appearance ([MethodOpenPBRAppearancesUpdate]).
type UpdateOpenPBRAppearanceArgs struct {
	ID           string                    `json:"id"`
	Base         types.OpenPBRBase         `json:"base"`
	Specular     types.OpenPBRSpecular     `json:"specular"`
	Transmission types.OpenPBRTransmission `json:"transmission"`
	Subsurface   types.OpenPBRSubsurface   `json:"subsurface"`
	Coat         types.OpenPBRCoat         `json:"coat"`
	Fuzz         types.OpenPBRFuzz         `json:"fuzz"`
	ThinFilm     types.OpenPBRThinFilm     `json:"thinFilm"`
	Emission     types.OpenPBREmission     `json:"emission"`
	Geometry     types.OpenPBRGeometry     `json:"geometry"`
}

// AssignOpenPBRAppearanceArgs overrides the OpenPBR appearance at a scope ("part",
// "body", or "face"), mirroring [AssignAppearanceArgs]. Key is the hex reference key of
// the body/face (empty for the part default).
type AssignOpenPBRAppearanceArgs struct {
	Scope        string `json:"scope"`
	Key          string `json:"key,omitempty"`
	AppearanceID string `json:"appearanceId"`
}
