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
	// WireframeViewStyle shows every edge as visible (no hidden-line removal). An Oblikovati-only
	// style with no Inventor equivalent (#1985), kept for the wireframe preview.
	WireframeViewStyle
	// ShadedViewStyle shades the view (reserved; renders as hidden-line until shading lands).
	ShadedViewStyle
	// HiddenLineRemovedViewStyle shows visible edges only — the canonical drafting style with no
	// dashed hidden lines (Inventor's kHiddenLineRemovedDrawingViewStyle, #1985).
	HiddenLineRemovedViewStyle
	// FromBaseViewStyle inherits the parent view's style associatively (Inventor's kFromBaseDrawingViewStyle).
	FromBaseViewStyle
	// ShadedHiddenLineViewStyle overlays shading with hidden edges (reserved until shading lands).
	ShadedHiddenLineViewStyle
)

var drawingViewStyleNames = map[DrawingViewStyle]string{
	HiddenLineViewStyle:        "hiddenLine",
	WireframeViewStyle:         "wireframe",
	ShadedViewStyle:            "shaded",
	HiddenLineRemovedViewStyle: "hiddenLineRemoved",
	FromBaseViewStyle:          "fromBase",
	ShadedHiddenLineViewStyle:  "shadedHiddenLine",
}

// String returns the style's wire spelling.
func (s DrawingViewStyle) String() string { return enumName(drawingViewStyleNames, s) }

// ParseDrawingViewStyle resolves a wire spelling back to its style.
func ParseDrawingViewStyle(s string) (DrawingViewStyle, bool) {
	return enumFromName(drawingViewStyleNames, s)
}

// RemovesHiddenEdges reports whether the style drops hidden edges entirely (visible edges only) —
// true only for the hidden-line-removed style, so a view recompute can filter them out (#1985).
// Wireframe keeps hidden edges but draws them as visible, so it does not remove them.
func (s DrawingViewStyle) RemovesHiddenEdges() bool {
	return s == HiddenLineRemovedViewStyle
}

// DrawingViewType discriminates the kind of a drawing view. The reference contracts model
// auxiliary/overlay/slice as a base DrawingView carrying this discriminator (not distinct
// interfaces); only section and detail are also distinct contracts. The zero value is
// DrawingViewBase.
type DrawingViewType int32

const (
	// DrawingViewBase projects a standard orientation of the model.
	DrawingViewBase DrawingViewType = iota
	// DrawingViewProjected is an orthographic view derived from a base view by a direction.
	DrawingViewProjected
	// DrawingViewAuxiliary is projected perpendicular to a fold line drawn on a parent view —
	// it shows an inclined face true-size.
	DrawingViewAuxiliary
	// DrawingViewSection shows the model cut by a plane (with cut-face hatching).
	DrawingViewSection
	// DrawingViewDetail magnifies a circular region of a parent view at a larger scale.
	DrawingViewDetail
	// DrawingViewBreak removes a band of a view to compress a long part (with break lines).
	DrawingViewBreak
	// DrawingViewSlice shows only the zero-thickness slice at a section line (no projection behind).
	DrawingViewSlice
	// DrawingViewBreakout reveals the interior within a bounded region of a parent view (a local
	// cut-away).
	DrawingViewBreakout
	// DrawingViewDraft is a model-less view: a framed container for manually-drawn 2D geometry.
	DrawingViewDraft
)

var drawingViewTypeNames = map[DrawingViewType]string{
	DrawingViewBase:      "base",
	DrawingViewProjected: "projected",
	DrawingViewAuxiliary: "auxiliary",
	DrawingViewSection:   "section",
	DrawingViewDetail:    "detail",
	DrawingViewBreak:     "break",
	DrawingViewSlice:     "slice",
	DrawingViewBreakout:  "breakout",
	DrawingViewDraft:     "draft",
}

// String returns the view type's wire spelling ("base", "auxiliary").
func (t DrawingViewType) String() string { return enumName(drawingViewTypeNames, t) }

// ParseDrawingViewType resolves a wire spelling back to its view type.
func ParseDrawingViewType(s string) (DrawingViewType, bool) {
	return enumFromName(drawingViewTypeNames, s)
}

// BreakOrientation is the axis along which a break view compresses: a horizontal break removes
// a vertical band (shortening a wide part), a vertical break removes a horizontal band. The zero
// value is BreakHorizontal.
type BreakOrientation int32

const (
	// BreakHorizontal removes a vertical band, compressing the view horizontally.
	BreakHorizontal BreakOrientation = iota
	// BreakVertical removes a horizontal band, compressing the view vertically.
	BreakVertical
)

var breakOrientationNames = map[BreakOrientation]string{
	BreakHorizontal: "horizontal",
	BreakVertical:   "vertical",
}

// String returns the break orientation's wire spelling.
func (o BreakOrientation) String() string { return enumName(breakOrientationNames, o) }

// ParseBreakOrientation resolves a wire spelling back to its break orientation.
func ParseBreakOrientation(s string) (BreakOrientation, bool) {
	return enumFromName(breakOrientationNames, s)
}

// DrawingCurveKind classifies a drawing curve so the head can style it: an edge of the model
// (visible/hidden), a section-cut outline, a hatch line, or a break-line glyph. The zero value
// is DrawingEdgeCurve, so the existing visible/hidden edge curves keep their meaning.
type DrawingCurveKind int32

const (
	// DrawingEdgeCurve is a projected model edge (the visible/hidden flag styles it solid/dashed).
	DrawingEdgeCurve DrawingCurveKind = iota
	// DrawingSectionCurve is a section-cut outline (drawn bold).
	DrawingSectionCurve
	// DrawingHatchCurve is one hatch line filling a section face.
	DrawingHatchCurve
	// DrawingBreakCurve is a break-line glyph segment.
	DrawingBreakCurve
)

var drawingCurveKindNames = map[DrawingCurveKind]string{
	DrawingEdgeCurve:    "edge",
	DrawingSectionCurve: "section",
	DrawingHatchCurve:   "hatch",
	DrawingBreakCurve:   "break",
}

// String returns the curve kind's wire spelling.
func (k DrawingCurveKind) String() string { return enumName(drawingCurveKindNames, k) }

// ParseDrawingCurveKind resolves a wire spelling back to its curve kind.
func ParseDrawingCurveKind(s string) (DrawingCurveKind, bool) {
	return enumFromName(drawingCurveKindNames, s)
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

// DrawingAnnotationKind classifies a drawing annotation. The zero value is CoGMarkerAnnotation.
type DrawingAnnotationKind int32

const (
	// CoGMarkerAnnotation is a centre-of-gravity indicator on a view, driven by the model's mass.
	CoGMarkerAnnotation DrawingAnnotationKind = iota
	// RevisionCloudAnnotation is a scalloped cloud highlighting a changed sheet region.
	RevisionCloudAnnotation
	// CenterMarkAnnotation is a crosshair at a circular edge's centre, associative to that edge.
	CenterMarkAnnotation
	// CenterlineAnnotation is the horizontal+vertical dash-dot symmetry axes through a view's
	// centre, spanning its extent; associative to the view (re-derives from its bounds).
	CenterlineAnnotation
	// FeatureControlFrameAnnotation is a GD&T feature control frame: a boxed geometric-tolerance
	// callout (characteristic symbol · tolerance · datum references) placed on the sheet.
	FeatureControlFrameAnnotation
	// DatumFeatureAnnotation is a GD&T datum feature symbol: a letter in a box with a filled
	// datum triangle, marking the datum a feature control frame references.
	DatumFeatureAnnotation
	// SurfaceTextureAnnotation is an ISO 1302 surface texture symbol: the checkmark glyph with a
	// roughness value, stating a surface's finish requirement.
	SurfaceTextureAnnotation
	// PartsListAnnotation is a parts list table sourced from the referenced assembly's BOM (item
	// number, part number, description, quantity), updating with the assembly.
	PartsListAnnotation
	// BalloonAnnotation is a balloon: a circle holding a parts-list item number, with an optional
	// leader to the component it tags.
	BalloonAnnotation
	// HoleTableAnnotation is a hole table: a row per circular edge in a base view, with its X/Y
	// position from a datum origin and its diameter, updating with the model.
	HoleTableAnnotation
	// RevisionTableAnnotation is a revision table: a row per revision (revision, date, description),
	// recording the drawing's change history.
	RevisionTableAnnotation
	// RevisionTagAnnotation is a revision tag: a triangle holding a revision letter, placed on the
	// sheet to flag where that revision changed the drawing.
	RevisionTagAnnotation
	// DrawingNoteAnnotation is a free text note on the sheet, with an optional leader line to the
	// point it annotates.
	DrawingNoteAnnotation
	// CustomTableAnnotation is a general-purpose table: arbitrary column headers and rows in a grid.
	CustomTableAnnotation
	// HoleNoteAnnotation is a feature note on a base view's holes: a leadered diameter callout per
	// hole, computed from the hole's circular edge and re-resolved when the model changes.
	HoleNoteAnnotation
)

var drawingAnnotationKindNames = map[DrawingAnnotationKind]string{
	CoGMarkerAnnotation:           "cog",
	RevisionCloudAnnotation:       "revisionCloud",
	CenterMarkAnnotation:          "centerMark",
	CenterlineAnnotation:          "centerline",
	FeatureControlFrameAnnotation: "featureControlFrame",
	DatumFeatureAnnotation:        "datumFeature",
	SurfaceTextureAnnotation:      "surfaceTexture",
	PartsListAnnotation:           "partsList",
	BalloonAnnotation:             "balloon",
	HoleTableAnnotation:           "holeTable",
	RevisionTableAnnotation:       "revisionTable",
	RevisionTagAnnotation:         "revisionTag",
	DrawingNoteAnnotation:         "drawingNote",
	CustomTableAnnotation:         "customTable",
	HoleNoteAnnotation:            "holeNote",
}

// String returns the annotation kind's wire spelling.
func (k DrawingAnnotationKind) String() string { return enumName(drawingAnnotationKindNames, k) }

// ParseDrawingAnnotationKind resolves a wire spelling back to its annotation kind.
func ParseDrawingAnnotationKind(s string) (DrawingAnnotationKind, bool) {
	return enumFromName(drawingAnnotationKindNames, s)
}
