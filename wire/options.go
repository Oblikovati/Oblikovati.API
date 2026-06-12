// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Application-option groups (M05-F11, #618). Deliberately narrow: every option here
// is read by something real in the host — startup behavior, the ViewCube and color
// scheme, the sketch grid, the default chamfer corner treatment. Reference options
// with no live reader are declined into the won't-implement ADR rather than shipped
// as dead settings.

// The option group names of [MethodOptionsGetGroup] / [MethodOptionsSetGroup].
const (
	OptionGroupGeneral = "general"
	OptionGroupDisplay = "display"
	OptionGroupSketch  = "sketch"
	OptionGroupPart    = "part"
	OptionGroupSave    = "save"
)

// GeneralOptionsView is the "general" group: application-level behavior.
type GeneralOptionsView struct {
	StartupAction types.StartupActionType `json:"startupAction"`
}

// DisplayOptionsView is the "display" group: the color scheme (the active theme's
// name — themes themselves are managed by the theme.* methods) and the ViewCube.
// CubeCorner is the viewport corner the cube anchors to (0=top-right, 1=top-left,
// 2=bottom-right, 3=bottom-left).
type DisplayOptionsView struct {
	ColorScheme         string  `json:"colorScheme"`
	ViewCubeHidden      bool    `json:"viewCubeHidden,omitempty"`
	CompassHidden       bool    `json:"compassHidden,omitempty"`
	LockToSelection     bool    `json:"lockToSelection,omitempty"`
	CubeInactiveOpacity float64 `json:"cubeInactiveOpacity,omitempty"`
	CubeSizePx          int     `json:"cubeSizePx,omitempty"`
	CubeCorner          int     `json:"cubeCorner,omitempty"`
}

// SketchOptionsView is the "sketch" group: the grid and click snapping. Spacing is
// in model/database units (cm) — unit-independent, like the stored preference.
type SketchOptionsView struct {
	GridSpacingCm  float64 `json:"gridSpacingCm"`
	GridVisible    bool    `json:"gridVisible"`
	GridMajorEvery int     `json:"gridMajorEvery"`
	SnapToPoints   bool    `json:"snapToPoints"`
	SnapToGrid     bool    `json:"snapToGrid"`
}

// PartOptionsView is the "part" group: part-modeling defaults.
type PartOptionsView struct {
	ChamferFlatCorners bool `json:"chamferFlatCorners"`
}

// GetOptionGroupArgs is the request of [MethodOptionsGetGroup].
type GetOptionGroupArgs struct {
	Group string `json:"group"`
}

// OptionGroupView is the response of [MethodOptionsGetGroup] and the request of
// [MethodOptionsSetGroup]: exactly one of the group fields is set, matching Group.
type OptionGroupView struct {
	Group   string              `json:"group"`
	General *GeneralOptionsView `json:"general,omitempty"`
	Display *DisplayOptionsView `json:"display,omitempty"`
	Sketch  *SketchOptionsView  `json:"sketch,omitempty"`
	Part    *PartOptionsView    `json:"part,omitempty"`
	Save    *SaveOptionsView    `json:"save,omitempty"`
}

// ListOptionGroupsResult is the response of [MethodOptionsListGroups].
type ListOptionGroupsResult struct {
	Groups []string `json:"groups"`
}
