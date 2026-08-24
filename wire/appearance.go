// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// AppearanceInfo is the JSON shape of a full OpenPBR Surface v1.1.1 appearance. Every
// group is the [types] value type from PBI-335 (colors are ACEScg [types.Color3], not
// hex, since emission_color is unbounded above and hex cannot represent that).
type AppearanceInfo struct {
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

// ListAppearancesResult is the [MethodAppearancesList] response.
type ListAppearancesResult struct {
	Appearances []AppearanceInfo `json:"appearances"`
}

// CreateAppearanceArgs creates a custom appearance by copying an existing one under a
// new name ([MethodAppearancesCreate]), mirroring [DuplicateAssetArgs]'s shape for the
// materials.create method.
type CreateAppearanceArgs struct {
	BaseID string `json:"baseId"`
	Name   string `json:"name"`
}

// UpdateAppearanceArgs replaces the display name and every group of an existing,
// editable appearance ([MethodAppearancesUpdate]).
type UpdateAppearanceArgs struct {
	ID           string                    `json:"id"`
	DisplayName  string                    `json:"displayName"`
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

// AssignAppearanceArgs overrides the appearance at a scope ("part", "body", or "face").
// Key is the hex reference key of the body/face (empty for the part default).
type AssignAppearanceArgs struct {
	Scope        string `json:"scope"`
	Key          string `json:"key,omitempty"`
	AppearanceID string `json:"appearanceId"`
}
