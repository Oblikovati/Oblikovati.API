// SPDX-License-Identifier: Apache-2.0

package types

// ThemeToken names one semantic color slot of the UI theme — the curated, stable
// vocabulary the host styles itself from and add-ins read to match the host's look.
// One token may drive several concrete Dear ImGui colors (e.g. ChromeAccent feeds the
// active tab, the checkmark, and the slider grab), keeping the user-facing palette
// small. The string values are the wire/file keys; treat them as frozen.
//
// Tokens cover the application shell only — window, menus, viewport 2D overlays, and
// 3D gizmos. They deliberately do NOT cover 3D body appearance, which the material/
// appearance subsystem owns.
type ThemeToken string

// Chrome — the windowed shell drawn by Dear ImGui (menus, panels, ribbon, controls).
const (
	TokenChromeWindowBg      ThemeToken = "chrome.window_bg"
	TokenChromePanelBg       ThemeToken = "chrome.panel_bg"
	TokenChromePopupBg       ThemeToken = "chrome.popup_bg"
	TokenChromeMenuBarBg     ThemeToken = "chrome.menu_bar_bg"
	TokenChromeHeaderBg      ThemeToken = "chrome.header_bg"
	TokenChromeText          ThemeToken = "chrome.text"
	TokenChromeTextDisabled  ThemeToken = "chrome.text_disabled"
	TokenChromeBorder        ThemeToken = "chrome.border"
	TokenChromeControlBg     ThemeToken = "chrome.control_bg"
	TokenChromeControlHover  ThemeToken = "chrome.control_hover"
	TokenChromeControlActive ThemeToken = "chrome.control_active"
	TokenChromeButton        ThemeToken = "chrome.button"
	TokenChromeButtonHover   ThemeToken = "chrome.button_hover"
	TokenChromeButtonActive  ThemeToken = "chrome.button_active"
	TokenChromeAccent        ThemeToken = "chrome.accent"
	TokenChromeScrollbar     ThemeToken = "chrome.scrollbar"
)

// Viewport 2D — the sketch/grid/dimension overlays drawn into the 3D view.
const (
	TokenViewportBg       ThemeToken = "viewport.bg"
	TokenGridMinor        ThemeToken = "viewport.grid_minor"
	TokenGridMajor        ThemeToken = "viewport.grid_major"
	TokenGridAxis         ThemeToken = "viewport.grid_axis"
	TokenSketchGeometry   ThemeToken = "viewport.sketch_geometry"
	TokenSketchSelected   ThemeToken = "viewport.sketch_selected"
	TokenSketchCandidate  ThemeToken = "viewport.sketch_candidate"
	TokenSketchPreview    ThemeToken = "viewport.sketch_preview"
	TokenDimensionDriving ThemeToken = "viewport.dimension_driving"
	TokenDimensionDriven  ThemeToken = "viewport.dimension_driven"
	TokenSnapGlyph        ThemeToken = "viewport.snap_glyph"
	// TokenViewportActiveBorder is the outline drawn around the focused view tile in a
	// split (multi-view) layout, so the user can tell which view is active.
	TokenViewportActiveBorder ThemeToken = "viewport.active_border"
)

// Gizmos 3D — manipulator/affordance geometry (work planes, selection highlight).
const (
	TokenPlaneFaint         ThemeToken = "gizmo.plane_faint"
	TokenPlaneHover         ThemeToken = "gizmo.plane_hover"
	TokenPlaneSelected      ThemeToken = "gizmo.plane_selected"
	TokenSelectionHighlight ThemeToken = "gizmo.selection_highlight"
	// TokenPlaneFill is the translucent fill of a work plane's display square — its
	// alpha sets how see-through the plane is, so the user configures the look here.
	TokenPlaneFill ThemeToken = "gizmo.plane_fill"
)

// Icons — ribbon glyphs, rasterized as white alpha masks and tinted at draw time.
const (
	TokenIconTint     ThemeToken = "icon.tint"
	TokenIconDisabled ThemeToken = "icon.disabled"
)

// AllThemeTokens lists every token once, in display order (Chrome, Viewport, Gizmos,
// Icons). A complete theme palette defines a color for each; the editor iterates this
// to render its grouped color rows. Keep new tokens appended to their group.
func AllThemeTokens() []ThemeToken {
	return []ThemeToken{
		TokenChromeWindowBg, TokenChromePanelBg, TokenChromePopupBg, TokenChromeMenuBarBg,
		TokenChromeHeaderBg, TokenChromeText, TokenChromeTextDisabled, TokenChromeBorder,
		TokenChromeControlBg, TokenChromeControlHover, TokenChromeControlActive,
		TokenChromeButton, TokenChromeButtonHover, TokenChromeButtonActive,
		TokenChromeAccent, TokenChromeScrollbar,
		TokenViewportBg, TokenGridMinor, TokenGridMajor, TokenGridAxis,
		TokenSketchGeometry, TokenSketchSelected, TokenSketchCandidate, TokenSketchPreview,
		TokenDimensionDriving, TokenDimensionDriven, TokenSnapGlyph,
		TokenViewportActiveBorder,
		TokenPlaneFaint, TokenPlaneHover, TokenPlaneSelected, TokenPlaneFill,
		TokenSelectionHighlight,
		TokenIconTint, TokenIconDisabled,
	}
}

// ThemeKind classifies a theme as one of the two shipped built-ins or a user copy.
// It picks the right default seed and tells the editor which themes are read-only
// (built-ins) versus editable (custom).
type ThemeKind string

const (
	ThemeLight  ThemeKind = "light"
	ThemeDark   ThemeKind = "dark"
	ThemeCustom ThemeKind = "custom"
)

// Editable reports whether the user may recolor a theme of this kind (only customs).
func (k ThemeKind) Editable() bool { return k == ThemeCustom }
