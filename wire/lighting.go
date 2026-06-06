// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati/api/types"

// LightInfo is the JSON shape of one scene light — the element of [LightListResult] and the
// payload of [MethodLightingAddLight] / [MethodLightingSetLight]. Direction/Position are
// world-space (cm); Color is the light color and Intensity scales it.
type LightInfo struct {
	Index               int                           `json:"index"`
	LightType           types.LightTypeEnum           `json:"lightType"`
	LightDefinitionType types.LightDefinitionTypeEnum `json:"definitionType"`
	On                  bool                          `json:"on"`
	Color               types.Rgba                    `json:"color"`
	Intensity           float64                       `json:"intensity"`
	Direction           [3]float64                    `json:"direction"`
	Position            [3]float64                    `json:"position"`
	SpotInnerAngle      float64                       `json:"spotInnerAngle"`
	SpotOuterAngle      float64                       `json:"spotOuterAngle"`
	Attenuation         [3]float64                    `json:"attenuation"`
}

// LightListResult is the response of [MethodLightingListLights].
type LightListResult struct {
	Lights []LightInfo `json:"lights"`
}

// AddLightArgs is the request of [MethodLightingAddLight]: the emission shape of the light to
// add (it is created with neutral defaults the caller then tunes via SetLight).
type AddLightArgs struct {
	DefinitionType types.LightDefinitionTypeEnum `json:"definitionType"`
}

// SetLightArgs is the request of [MethodLightingSetLight]: the index of the light to update
// and its new full state.
type SetLightArgs struct {
	Index int       `json:"index"`
	Light LightInfo `json:"light"`
}

// LightingStyleView is the JSON shape of the active lighting style — the response of
// [MethodLightingGetStyle] and [MethodLightingSetStyle].
type LightingStyleView struct {
	Name           string                      `json:"name"`
	StyleType      types.LightingStyleTypeEnum `json:"styleType"`
	Ambience       float64                     `json:"ambience"`
	Brightness     float64                     `json:"brightness"`
	Exposure       float64                     `json:"exposure"`
	IBLBrightness  float64                     `json:"iblBrightness"`
	IBLRotation    float64                     `json:"iblRotation"`
	ShadowDensity  float64                     `json:"shadowDensity"`
	ShadowSoftness float64                     `json:"shadowSoftness"`
	ShadowDir      types.ShadowDirectionEnum   `json:"shadowDirection"`
}

// SetLightingStyleArgs is the request of [MethodLightingSetStyle]: the style to activate, by
// name (Inventor identifies lighting styles by name, not by an enum).
type SetLightingStyleArgs struct {
	Name string `json:"name"`
}

// LightingStyleInfo is one entry of [LightingStyleListResult]: a selectable style and whether
// it is the active one.
type LightingStyleInfo struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// LightingStyleListResult is the response of [MethodLightingListStyles].
type LightingStyleListResult struct {
	Styles []LightingStyleInfo `json:"styles"`
}

// ShadowSettings is the JSON shape of the View's shadow toggles — the response of
// [MethodViewGetShadows] / [MethodViewSetShadows] and the request of the latter.
// Density and Softness are in [0,1]. GroundShadow selects the ground-shadow style.
type ShadowSettings struct {
	GroundShadow   types.GroundShadowEnum `json:"groundShadow"`
	ObjectShadows  bool                   `json:"objectShadows"`
	AmbientShadows bool                   `json:"ambientShadows"`
	Density        float64                `json:"density"`
	Softness       float64                `json:"softness"`
}

// EnvironmentView is the JSON shape of the active IBL environment — the response of
// [MethodEnvironmentGet] / [MethodEnvironmentSet] / [MethodEnvironmentLoadImage]. Preset is
// the built-in name (empty when a file is loaded); FilePath is the loaded HDR (empty for a
// preset). Rotation is radians about vertical; Intensity scales the IBL; ShowImage draws the
// sky as the background.
type EnvironmentView struct {
	Preset    string  `json:"preset"`
	FilePath  string  `json:"filePath"`
	Rotation  float64 `json:"rotation"`
	Intensity float64 `json:"intensity"`
	ShowImage bool    `json:"showImage"`
}

// SetEnvironmentArgs is the request of [MethodEnvironmentSet]: select a built-in preset by
// name and set its display parameters.
type SetEnvironmentArgs struct {
	Preset    string  `json:"preset"`
	Rotation  float64 `json:"rotation"`
	Intensity float64 `json:"intensity"`
	ShowImage bool    `json:"showImage"`
}

// LoadEnvironmentImageArgs is the request of [MethodEnvironmentLoadImage]: the path to an
// equirectangular HDR file (.hdr) to use as the environment (Inventor's
// LightingStyle.UploadImage).
type LoadEnvironmentImageArgs struct {
	FilePath string `json:"filePath"`
}

// EnvironmentPresetInfo is one entry of [EnvironmentPresetListResult]: a selectable built-in
// environment and whether it is the active one.
type EnvironmentPresetInfo struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// EnvironmentPresetListResult is the response of [MethodEnvironmentListPresets].
type EnvironmentPresetListResult struct {
	Presets []EnvironmentPresetInfo `json:"presets"`
}
