// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati/api/types"

// DisplayModeView is the JSON shape of the viewport's current display mode: the enum value and
// its label. The response of [MethodViewGetDisplayMode] and [MethodViewSetDisplayMode].
type DisplayModeView struct {
	Mode types.DisplayModeEnum `json:"mode"`
	Name string                `json:"name"`
}

// SetDisplayModeArgs is the request of [MethodViewSetDisplayMode]: the mode to switch the
// viewport to.
type SetDisplayModeArgs struct {
	Mode types.DisplayModeEnum `json:"mode"`
}

// DisplayModeInfo is one entry of [ListDisplayModesResult]: a selectable mode, its label, and
// whether it is the active one — enough to populate a display-mode picker.
type DisplayModeInfo struct {
	Mode   types.DisplayModeEnum `json:"mode"`
	Name   string                `json:"name"`
	Active bool                  `json:"active"`
}

// ListDisplayModesResult is the response of [MethodViewListDisplayModes].
type ListDisplayModesResult struct {
	Modes []DisplayModeInfo `json:"modes"`
}
