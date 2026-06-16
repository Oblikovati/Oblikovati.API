// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// ColorSchemeView is the JSON shape of one color scheme — the element of
// [ColorSchemesResult] and the response of [MethodColorSchemesGetActive] /
// [MethodColorSchemesSetActive]. The background colors honor BackgroundType: ScreenColor for
// a solid background, Top/BottomScreenColor for a gradient. Highlight/select colors feed the
// selection pipeline.
type ColorSchemeView struct {
	Name           string                   `json:"name"`
	Active         bool                     `json:"active"`
	BackgroundType types.BackgroundTypeEnum `json:"backgroundType"`
	ScreenColor    types.Color              `json:"screenColor"`
	TopScreen      types.Color              `json:"topScreenColor"`
	BottomScreen   types.Color              `json:"bottomScreenColor"`
	Highlight      types.Color              `json:"highlightColor"`
	PrimarySelect  types.Color              `json:"primarySelectColor"`
	SecondSelect   types.Color              `json:"secondarySelectColor"`
}

// ColorSchemesResult is the response of [MethodColorSchemesList]: every scheme, with the
// active one flagged.
type ColorSchemesResult struct {
	Schemes []ColorSchemeView `json:"schemes"`
}

// SetColorSchemeArgs is the request of [MethodColorSchemesSetActive]: activate the scheme by
// name (schemes are identified by name).
type SetColorSchemeArgs struct {
	Name string `json:"name"`
}
