// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// GroundPlaneView is the JSON shape of a document's ground-plane settings — a member of
// [DisplaySettingsView]. Lengths are cm; Opacity/Reflectivity in [0,1].
type GroundPlaneView struct {
	Visible                    bool        `json:"visible"`
	Color                      types.Color `json:"color"`
	HeightOffset               float64     `json:"heightOffset"`
	DisplayGridLines           bool        `json:"displayGridLines"`
	MinorGridLineSpacing       float64     `json:"minorGridLineSpacing"`
	MinorLinesPerMajorGridLine int         `json:"minorLinesPerMajorGridLine"`
	Opacity                    float64     `json:"opacity"`
	Reflectivity               float64     `json:"reflectivity"`
}

// ShadedDisplayModeOptionsView is the JSON shape of the shaded-mode sub-options — a member of
// [DisplayModeOptionsView].
type ShadedDisplayModeOptionsView struct {
	EdgeDisplay      bool                       `json:"edgeDisplay"`
	EdgeColor        types.Color                `json:"edgeColor"`
	Silhouettes      bool                       `json:"silhouettes"`
	TransparencyType types.TransparencyTypeEnum `json:"transparencyType"`
}

// WireframeDisplayModeOptionsView is the JSON shape of the wireframe-mode sub-options — a member
// of [DisplayModeOptionsView].
type WireframeDisplayModeOptionsView struct {
	DepthDimming      bool `json:"depthDimming"`
	Silhouettes       bool `json:"silhouettes"`
	DimmedHiddenEdges bool `json:"dimmedHiddenEdges"`
}

// DisplayModeOptionsView is the JSON shape of the application-level display options — the response
// of [MethodDisplayGetOptions] and the request of [MethodDisplaySetOptions].
type DisplayModeOptionsView struct {
	DisplayQuality           types.DisplayQualityEnum        `json:"displayQuality"`
	ViewTransitionTime       float64                         `json:"viewTransitionTime"`
	MinimumFrameRate         float64                         `json:"minimumFrameRate"`
	HiddenLineDimmingPercent int                             `json:"hiddenLineDimmingPercent"`
	EdgeColor                types.Color                     `json:"edgeColor"`
	NewWindowDisplayMode     types.DisplayModeEnum           `json:"newWindowDisplayMode"`
	NewWindowProjection      types.ProjectionTypeEnum        `json:"newWindowProjection"`
	BackFaceCulling          types.BackFaceCullingEnum       `json:"backFaceCulling"`
	UseRayTracing            bool                            `json:"useRayTracing"`
	RayTracingQuality        types.RayTracingQualityEnum     `json:"rayTracingQuality"`
	Shaded                   ShadedDisplayModeOptionsView    `json:"shaded"`
	Wireframe                WireframeDisplayModeOptionsView `json:"wireframe"`
}

// DisplaySettingsView is the JSON shape of a document's per-document display settings — the
// response of [MethodDocumentGetDisplaySettings] and the request of
// [MethodDocumentSetDisplaySettings].
type DisplaySettingsView struct {
	BackgroundType           types.BackgroundTypeEnum        `json:"backgroundType"`
	EdgeColor                types.Color                     `json:"edgeColor"`
	DepthDimming             bool                            `json:"depthDimming"`
	DisplaySilhouettes       bool                            `json:"displaySilhouettes"`
	HiddenLineDimmingPercent int                             `json:"hiddenLineDimmingPercent"`
	NewWindowDisplayMode     types.DisplayModeEnum           `json:"newWindowDisplayMode"`
	DisplayModeSource        types.DisplayModeSourceTypeEnum `json:"displayModeSource"`
	NewWindowProjection      types.ProjectionTypeEnum        `json:"newWindowProjection"`
	GroundPlane              GroundPlaneView                 `json:"groundPlane"`
	GroundShadow             types.GroundShadowEnum          `json:"groundShadow"`
	ShadowDirection          types.ShadowDirectionEnum       `json:"shadowDirection"`
	ShowGroundReflections    bool                            `json:"showGroundReflections"`
	ShowObjectShadows        bool                            `json:"showObjectShadows"`
	ShowAmbientShadows       bool                            `json:"showAmbientShadows"`
	TexturesOn               bool                            `json:"texturesOn"`
}

// GetDisplaySettingsArgs is the request of [MethodDocumentGetDisplaySettings]: which document's
// settings to read (empty selects the active document).
type GetDisplaySettingsArgs struct {
	Document string `json:"document,omitempty"`
}

// SetDisplaySettingsArgs is the request of [MethodDocumentSetDisplaySettings]: the document to
// update and the settings to apply.
type SetDisplaySettingsArgs struct {
	Document string              `json:"document,omitempty"`
	Settings DisplaySettingsView `json:"settings"`
}
