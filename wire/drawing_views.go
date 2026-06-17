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
	Projected    bool    `json:"projected"`
	BaseView     string  `json:"baseView,omitempty"`  // set for a projected view
	Orientation  string  `json:"orientation"`         // types.BaseViewOrientation spelling
	Direction    string  `json:"direction,omitempty"` // types.ProjectionDirection (projected views)
	Scale        float64 `json:"scale"`
	Style        string  `json:"style"` // types.DrawingViewStyle spelling
	CenterXMM    float64 `json:"centerXmm"`
	CenterYMM    float64 `json:"centerYmm"`
	VisibleCount int     `json:"visibleCount"`
	HiddenCount  int     `json:"hiddenCount"`
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
	EdgeKey string  `json:"edgeKey,omitempty"`
}

// ViewCurvesResult is the response of [MethodDrawingViewsCurves]: the view's drawing curves.
type ViewCurvesResult struct {
	Segments []DrawingCurveSegment `json:"segments"`
}
