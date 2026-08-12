// SPDX-License-Identifier: Apache-2.0

package wire

// Drawing view wire DTOs (M14-F02 PBI-139, Oblikovati/Oblikovati#386). Views project the
// drawing's referenced model onto the active sheet: a base view from a standard orientation
// (front/top/right/…/iso), and projected views placed in an orthographic direction off a base.
// A view's geometry is its drawing curves — 2D segments classified visible (solid) or hidden
// (dashed), each carrying the hex reference key of the model edge it came from for associativity.

// DrawingViewInfo is the JSON shape of one drawing view.
type DrawingViewInfo struct {
	Name         string  `json:"name"`
	Type         string  `json:"type"`                   // types.DrawingViewType spelling (base/projected/auxiliary/…)
	Projected    bool    `json:"projected"`              // true only for an orthographic projected view
	BaseView     string  `json:"baseView,omitempty"`     // the parent view (projected/auxiliary/…)
	Orientation  string  `json:"orientation"`            // types.BaseViewOrientation spelling
	Direction    string  `json:"direction,omitempty"`    // types.ProjectionDirection (projected views)
	FoldAngleDeg float64 `json:"foldAngleDeg,omitempty"` // fold-line angle on the parent (auxiliary views)
	Scale        float64 `json:"scale"`
	Style        string  `json:"style"` // types.DrawingViewStyle spelling
	CenterXMM    float64 `json:"centerXmm"`
	CenterYMM    float64 `json:"centerYmm"`
	VisibleCount int     `json:"visibleCount"`
	HiddenCount  int     `json:"hiddenCount"`
	// Label (#1983). Label is the composed caption drawn under the view (empty ⇒ none); the Show*
	// flags report whether the label, its name, and its scale note are shown; LabelXMM/YMM place it.
	Label     string  `json:"label,omitempty"`
	ShowLabel bool    `json:"showLabel"`
	ShowName  bool    `json:"showName"`
	ShowScale bool    `json:"showScale"`
	LabelXMM  float64 `json:"labelXmm,omitempty"`
	LabelYMM  float64 `json:"labelYmm,omitempty"`
	// Section options (#1982), reported for section views. SectionDepthMM is 0 for a full
	// through-cut; SectionReverse keeps the far half; SectionType is the partial-cut kind.
	SectionDepthMM float64 `json:"sectionDepthMm,omitempty"`
	SectionReverse bool    `json:"sectionReverse,omitempty"`
	SectionType    string  `json:"sectionType,omitempty"`
}

// SetViewLabelArgs is the request of [MethodDrawingViewsSetLabel]: change any subset of the named
// view's label (#1983). Each pointer field is applied only when present; Text="" restores the
// default caption. LabelXMM and LabelYMM must both be set to reposition the caption.
type SetViewLabelArgs struct {
	Name      string   `json:"name"`
	Text      *string  `json:"text,omitempty"`
	ShowLabel *bool    `json:"showLabel,omitempty"`
	ShowName  *bool    `json:"showName,omitempty"`
	ShowScale *bool    `json:"showScale,omitempty"`
	LabelXMM  *float64 `json:"labelXmm,omitempty"`
	LabelYMM  *float64 `json:"labelYmm,omitempty"`
}

// ListDrawingViewsResult is the response of [MethodDrawingViewsList]: the active sheet's views.
type ListDrawingViewsResult struct {
	Views []DrawingViewInfo `json:"views"`
}

// AddBaseViewArgs is the request of [MethodDrawingViewsAddBase]: a base view of the drawing's
// referenced model at the given orientation/scale/style, centred at (CenterXMM, CenterYMM) on
// the active sheet. An empty Name auto-assigns; empty Orientation/Style default to front /
// hiddenLine; a non-positive Scale defaults to 1.
type AddBaseViewArgs struct {
	Name        string  `json:"name,omitempty"`
	Orientation string  `json:"orientation,omitempty"`
	Style       string  `json:"style,omitempty"`
	Scale       float64 `json:"scale,omitempty"`
	CenterXMM   float64 `json:"centerXmm,omitempty"`
	CenterYMM   float64 `json:"centerYmm,omitempty"`
}

// AddProjectedViewArgs is the request of [MethodDrawingViewsAddProjected]: a view projected
// from BaseView in Direction (right/left/up/down), inheriting the base's scale and style.
type AddProjectedViewArgs struct {
	Name      string  `json:"name,omitempty"`
	BaseView  string  `json:"baseView"`
	Direction string  `json:"direction"`
	CenterXMM float64 `json:"centerXmm,omitempty"`
	CenterYMM float64 `json:"centerYmm,omitempty"`
}

// AddAuxiliaryViewArgs is the request of [MethodDrawingViewsAddAuxiliary]: a view projected
// perpendicular to a fold line drawn on the parent view at FoldAngleDeg (degrees, measured
// from the parent's horizontal axis), inheriting the parent's scale and style. A fold angle of
// 0 folds down (like a top projection); 90 folds to the side.
type AddAuxiliaryViewArgs struct {
	Name         string  `json:"name,omitempty"`
	ParentView   string  `json:"parentView"`
	FoldAngleDeg float64 `json:"foldAngleDeg"`
	CenterXMM    float64 `json:"centerXmm,omitempty"`
	CenterYMM    float64 `json:"centerYmm,omitempty"`
}

// AddSectionViewArgs is the request of [MethodDrawingViewsAddSection]: a section view of the
// parent's model, cut by the plane through the section line (X1,Y1)-(X2,Y2) on the parent (sheet
// millimetres), perpendicular to the parent. The near half is removed, the cut outline drawn
// bold and the exposed faces hatched; the view is placed at (CenterXMM, CenterYMM).
//
// The retained material is tunable (#1982): FullDepth (the default) keeps everything behind the
// plane, or FullDepth=false limits it to a slab SectionDepthMM deep so only geometry within that
// distance of the plane participates. Reverse keeps the opposite half. SectionType selects a
// partial cut (none/quarter/half/threeQuarter; "" ⇒ none, a plain full cut).
type AddSectionViewArgs struct {
	Name           string  `json:"name,omitempty"`
	ParentView     string  `json:"parentView"`
	X1             float64 `json:"x1"`
	Y1             float64 `json:"y1"`
	X2             float64 `json:"x2"`
	Y2             float64 `json:"y2"`
	CenterXMM      float64 `json:"centerXmm,omitempty"`
	CenterYMM      float64 `json:"centerYmm,omitempty"`
	FullDepth      bool    `json:"fullDepth,omitempty"`
	SectionDepthMM float64 `json:"sectionDepthMm,omitempty"`
	Reverse        bool    `json:"reverse,omitempty"`
	SectionType    string  `json:"sectionType,omitempty"`
}

// AddDetailViewArgs is the request of [MethodDrawingViewsAddDetail]: a magnified view of the
// circular region (BoundaryXMM, BoundaryYMM, RadiusMM — on the parent, sheet millimetres) of
// ParentView, at the larger Scale, placed at (CenterXMM, CenterYMM).
type AddDetailViewArgs struct {
	Name        string  `json:"name,omitempty"`
	ParentView  string  `json:"parentView"`
	BoundaryXMM float64 `json:"boundaryXmm"`
	BoundaryYMM float64 `json:"boundaryYmm"`
	RadiusMM    float64 `json:"radiusMm"`
	Scale       float64 `json:"scale"`
	CenterXMM   float64 `json:"centerXmm,omitempty"`
	CenterYMM   float64 `json:"centerYmm,omitempty"`
}

// AddBreakViewArgs is the request of [MethodDrawingViewsAddBreak]: a compressed view of the
// parent with a band removed. Orientation is "horizontal" (remove a vertical band) or
// "vertical"; GapStartMM/GapEndMM bound the removed band along that axis on the parent (sheet
// millimetres). The view is placed at (CenterXMM, CenterYMM).
type AddBreakViewArgs struct {
	Name        string  `json:"name,omitempty"`
	ParentView  string  `json:"parentView"`
	Orientation string  `json:"orientation,omitempty"` // types.BreakOrientation ("" ⇒ horizontal)
	GapStartMM  float64 `json:"gapStartMm"`
	GapEndMM    float64 `json:"gapEndMm"`
	CenterXMM   float64 `json:"centerXmm,omitempty"`
	CenterYMM   float64 `json:"centerYmm,omitempty"`
}

// AddSliceViewArgs is the request of [MethodDrawingViewsAddSlice]: a zero-thickness slice at the
// section line (X1,Y1)-(X2,Y2) on the parent (sheet mm) — only the cut outline, nothing behind.
type AddSliceViewArgs struct {
	Name       string  `json:"name,omitempty"`
	ParentView string  `json:"parentView"`
	X1         float64 `json:"x1"`
	Y1         float64 `json:"y1"`
	X2         float64 `json:"x2"`
	Y2         float64 `json:"y2"`
	CenterXMM  float64 `json:"centerXmm,omitempty"`
	CenterYMM  float64 `json:"centerYmm,omitempty"`
}

// AddBreakoutViewArgs is the request of [MethodDrawingViewsAddBreakout]: a copy of ParentView
// with the interior revealed inside the circular region (BoundaryXMM, BoundaryYMM, RadiusMM on
// the parent, sheet mm) — a local cut-away. Placed at (CenterXMM, CenterYMM).
type AddBreakoutViewArgs struct {
	Name        string  `json:"name,omitempty"`
	ParentView  string  `json:"parentView"`
	BoundaryXMM float64 `json:"boundaryXmm"`
	BoundaryYMM float64 `json:"boundaryYmm"`
	RadiusMM    float64 `json:"radiusMm"`
	CenterXMM   float64 `json:"centerXmm,omitempty"`
	CenterYMM   float64 `json:"centerYmm,omitempty"`
}

// AddDraftViewArgs is the request of [MethodDrawingViewsAddDraft]: a model-less framed view of
// WidthMM × HeightMM (sheet mm) at (CenterXMM, CenterYMM) — a container for manual 2D geometry.
type AddDraftViewArgs struct {
	Name      string  `json:"name,omitempty"`
	WidthMM   float64 `json:"widthMm"`
	HeightMM  float64 `json:"heightMm"`
	CenterXMM float64 `json:"centerXmm,omitempty"`
	CenterYMM float64 `json:"centerYmm,omitempty"`
}

// ViewResult is the response of [MethodDrawingViewsAddBase] / [MethodDrawingViewsAddProjected]:
// the created view.
type ViewResult struct {
	View DrawingViewInfo `json:"view"`
}

// DeleteViewArgs is the request of [MethodDrawingViewsDelete]: the view to remove (deleting a
// base view also removes the views projected from it).
type DeleteViewArgs struct {
	Name string `json:"name"`
}

// ViewCurvesArgs is the request of [MethodDrawingViewsCurves]: the view whose drawing curves
// to return.
type ViewCurvesArgs struct {
	View string `json:"view"`
}

// DrawingCurveSegment is one projected edge segment in sheet millimetres: its endpoints, its
// visibility (solid vs dashed), and the hex reference key of the source model edge.
type DrawingCurveSegment struct {
	AX      float64 `json:"ax"`
	AY      float64 `json:"ay"`
	BX      float64 `json:"bx"`
	BY      float64 `json:"by"`
	Visible bool    `json:"visible"`
	Kind    string  `json:"kind,omitempty"` // types.DrawingCurveKind ("edge" default; section/hatch/break)
	EdgeKey string  `json:"edgeKey,omitempty"`
}

// ViewCurvesResult is the response of [MethodDrawingViewsCurves]: the view's drawing curves.
type ViewCurvesResult struct {
	Segments []DrawingCurveSegment `json:"segments"`
}
