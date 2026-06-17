// SPDX-License-Identifier: Apache-2.0

package types

// Drawing view value types (M14-F02 PBI-139). A drawing view projects a 3D model onto the
// sheet: a base view from a standard orientation, and projected views derived from it by an
// orthographic direction. A rendering style fixes whether hidden edges are computed. These
// are the canonical, Apache-2.0 definitions; the GPL model aliases them and runs the
// hidden-line projection.

// BaseViewOrientation names the standard orientation a base view projects from. The zero
// value is BaseViewFront.
type BaseViewOrientation int32

const (
	// BaseViewFront looks along +Y (the front elevation).
	BaseViewFront BaseViewOrientation = iota
	// BaseViewTop looks straight down (-Z), the plan.
	BaseViewTop
	// BaseViewRight looks along -X (the right side).
	BaseViewRight
	// BaseViewBack looks along -Y.
	BaseViewBack
	// BaseViewLeft looks along +X.
	BaseViewLeft
	// BaseViewBottom looks straight up (+Z).
	BaseViewBottom
	// BaseViewIso is the top-front-right isometric.
	BaseViewIso
)

var baseViewOrientationNames = map[BaseViewOrientation]string{
	BaseViewFront:  "front",
	BaseViewTop:    "top",
	BaseViewRight:  "right",
	BaseViewBack:   "back",
	BaseViewLeft:   "left",
	BaseViewBottom: "bottom",
	BaseViewIso:    "iso",
}

// String returns the orientation's wire spelling ("front", "iso").
func (o BaseViewOrientation) String() string { return enumName(baseViewOrientationNames, o) }

// ParseBaseViewOrientation resolves a wire spelling back to its orientation.
func ParseBaseViewOrientation(s string) (BaseViewOrientation, bool) {
	return enumFromName(baseViewOrientationNames, s)
}

// DrawingViewStyle fixes how a view's edges are rendered. The zero value is
// HiddenLineViewStyle (visible solid + hidden dashed).
type DrawingViewStyle int32

const (
	// HiddenLineViewStyle shows visible edges solid and hidden edges dashed (the default).
	HiddenLineViewStyle DrawingViewStyle = iota
	// WireframeViewStyle shows every edge as visible (no hidden-line removal).
	WireframeViewStyle
	// ShadedViewStyle shades the view (reserved; renders as hidden-line until shading lands).
	ShadedViewStyle
)

var drawingViewStyleNames = map[DrawingViewStyle]string{
	HiddenLineViewStyle: "hiddenLine",
	WireframeViewStyle:  "wireframe",
	ShadedViewStyle:     "shaded",
}

// String returns the style's wire spelling.
func (s DrawingViewStyle) String() string { return enumName(drawingViewStyleNames, s) }

// ParseDrawingViewStyle resolves a wire spelling back to its style.
func ParseDrawingViewStyle(s string) (DrawingViewStyle, bool) {
	return enumFromName(drawingViewStyleNames, s)
}

// ProjectionDirection names where a projected view sits relative to its base view; it also
// fixes the orthographic direction the projection looks from. The zero value is ProjectRight.
type ProjectionDirection int32

const (
	// ProjectRight places the projected view to the right of the base (looking from the right).
	ProjectRight ProjectionDirection = iota
	// ProjectLeft places it to the left.
	ProjectLeft
	// ProjectUp places it above the base.
	ProjectUp
	// ProjectDown places it below the base.
	ProjectDown
)

var projectionDirectionNames = map[ProjectionDirection]string{
	ProjectRight: "right",
	ProjectLeft:  "left",
	ProjectUp:    "up",
	ProjectDown:  "down",
}

// String returns the direction's wire spelling.
func (d ProjectionDirection) String() string { return enumName(projectionDirectionNames, d) }

// ParseProjectionDirection resolves a wire spelling back to its direction.
func ParseProjectionDirection(s string) (ProjectionDirection, bool) {
	return enumFromName(projectionDirectionNames, s)
}
