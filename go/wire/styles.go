// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// ColorStyleView is the JSON shape of one color style — the element of [ColorStylesResult] and
// the response of [MethodStylesGet] / [MethodStylesSet] and the request of the latter. The
// color components are full [types.Color] value objects; Shininess/Opacity are in [0,1].
type ColorStyleView struct {
	Name      string                  `json:"name"`
	Diffuse   types.Color             `json:"diffuse"`
	Ambient   types.Color             `json:"ambient"`
	Specular  types.Color             `json:"specular"`
	Emissive  types.Color             `json:"emissive"`
	Shininess float64                 `json:"shininess"`
	Opacity   float64                 `json:"opacity"`
	Location  types.StyleLocationEnum `json:"location"`
}

// ColorStylesResult is the response of [MethodStylesList]: every color style in the document.
type ColorStylesResult struct {
	Styles []ColorStyleView `json:"styles"`
}

// GetStyleArgs is the request of [MethodStylesGet] / [MethodStylesDelete]: the style by name.
type GetStyleArgs struct {
	Name string `json:"name"`
}

// StyleLibraryInfo is one entry of [StyleLibrariesResult]: a loaded style library and its
// cascade position (lower Order shadows higher Order for a same-named style).
type StyleLibraryInfo struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Order int    `json:"order"`
}

// StyleLibrariesResult is the response of [MethodStylesListLibraries]: the loaded libraries in
// cascade order.
type StyleLibrariesResult struct {
	Libraries []StyleLibraryInfo `json:"libraries"`
}

// ImportStyleLibraryArgs is the request of [MethodStylesImportLibrary]: the path to a style
// library file to load into the cascade.
type ImportStyleLibraryArgs struct {
	Path string `json:"path"`
}

// StyleChangedEvent is the payload of the [EventStyleAdded] / [EventStyleChanged] /
// [EventStyleDeleted] push events: the style that changed and its kind ("color" / "lighting").
type StyleChangedEvent struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}
