// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The drawing views operation group (M14-F02 PBI-139, Oblikovati/Oblikovati#386): place base
// and projected views of the drawing's referenced model on the active sheet, and read the
// hidden-line drawing curves (visible/hidden segments) each view produces.

// DrawingViews is the drawing-views operation group.
type DrawingViews struct{ c *Client }

// DrawingViews returns the drawing-views operation group.
func (c *Client) DrawingViews() DrawingViews { return DrawingViews{c} }

// List returns the active sheet's views.
//
// mcp:tool drawing_list_views
// mcp:summary List the active sheet's drawing views (name, base/projected, orientation, scale, style, sheet centre, and visible/hidden curve counts).
func (d DrawingViews) List() (wire.ListDrawingViewsResult, error) {
	return call[wire.ListDrawingViewsResult](d.c, wire.MethodDrawingViewsList, struct{}{})
}

// AddBase adds a base view of the referenced model at the given orientation/scale/style,
// centred on the active sheet, and computes its hidden-line curves.
//
// mcp:tool drawing_add_base_view
// mcp:summary Add a base view of the drawing's referenced model (orientation = front|top|right|back|left|bottom|iso, style = hiddenLine|wireframe|shaded, scale e.g. 0.5 for 1:2) centred at centerXmm/centerYmm; computes visible/hidden edges.
func (d DrawingViews) AddBase(args wire.AddBaseViewArgs) (wire.ViewResult, error) {
	return call[wire.ViewResult](d.c, wire.MethodDrawingViewsAddBase, args)
}

// AddProjected adds a view projected from a base view in the given direction.
//
// mcp:tool drawing_add_projected_view
// mcp:summary Add a projected view off a base view (direction = right|left|up|down), inheriting the base's scale/style; placed at centerXmm/centerYmm.
func (d DrawingViews) AddProjected(args wire.AddProjectedViewArgs) (wire.ViewResult, error) {
	return call[wire.ViewResult](d.c, wire.MethodDrawingViewsAddProjected, args)
}

// AddAuxiliary adds a view projected perpendicular to a fold line on a parent view.
//
// mcp:tool drawing_add_auxiliary_view
// mcp:summary Add an auxiliary view off a parent view: projected perpendicular to a fold line at foldAngleDeg (0 folds down like a top view, 90 folds to the side), inheriting the parent's scale/style; placed at centerXmm/centerYmm. Shows an inclined face true-size.
func (d DrawingViews) AddAuxiliary(args wire.AddAuxiliaryViewArgs) (wire.ViewResult, error) {
	return call[wire.ViewResult](d.c, wire.MethodDrawingViewsAddAuxiliary, args)
}

// AddSection adds a section view cutting the parent's model along a section line.
//
// mcp:tool drawing_add_section_view
// mcp:summary Add a section view off a parent view: the model is cut by the plane through the section line (x1,y1)-(x2,y2) drawn on the parent (sheet mm), the near half removed, the cut outline drawn bold and the exposed faces hatched; placed at centerXmm/centerYmm.
func (d DrawingViews) AddSection(args wire.AddSectionViewArgs) (wire.ViewResult, error) {
	return call[wire.ViewResult](d.c, wire.MethodDrawingViewsAddSection, args)
}

// AddDetail adds a magnified detail view of a circular region of a parent view.
//
// mcp:tool drawing_add_detail_view
// mcp:summary Add a detail view: a magnified circular region (boundaryXmm/boundaryYmm/radiusMm on the parent, sheet mm) of a parent view, at the larger scale, placed at centerXmm/centerYmm.
func (d DrawingViews) AddDetail(args wire.AddDetailViewArgs) (wire.ViewResult, error) {
	return call[wire.ViewResult](d.c, wire.MethodDrawingViewsAddDetail, args)
}

// AddBreak adds a break view: the parent compressed by removing a band along an axis.
//
// mcp:tool drawing_add_break_view
// mcp:summary Add a break view: the parent view compressed by removing a band (orientation = horizontal removes a vertical band, vertical removes a horizontal one) between gapStartMm and gapEndMm on the parent (sheet mm), with break lines at the cut; placed at centerXmm/centerYmm.
func (d DrawingViews) AddBreak(args wire.AddBreakViewArgs) (wire.ViewResult, error) {
	return call[wire.ViewResult](d.c, wire.MethodDrawingViewsAddBreak, args)
}

// AddSlice adds a slice view: only the zero-thickness cut outline at a section line on the parent.
//
// mcp:tool drawing_add_slice_view
// mcp:summary Add a slice view off a parent: only the zero-thickness slice outline at the section line (x1,y1)-(x2,y2) on the parent (sheet mm), with nothing projected behind it; placed at centerXmm/centerYmm.
func (d DrawingViews) AddSlice(args wire.AddSliceViewArgs) (wire.ViewResult, error) {
	return call[wire.ViewResult](d.c, wire.MethodDrawingViewsAddSlice, args)
}

// AddBreakout adds a breakout view: a parent copy with the interior revealed in a bounded region.
//
// mcp:tool drawing_add_breakout_view
// mcp:summary Add a breakout view off a parent: a local cut-away revealing the interior inside the circular region (boundaryXmm/boundaryYmm/radiusMm on the parent, sheet mm); placed at centerXmm/centerYmm.
func (d DrawingViews) AddBreakout(args wire.AddBreakoutViewArgs) (wire.ViewResult, error) {
	return call[wire.ViewResult](d.c, wire.MethodDrawingViewsAddBreakout, args)
}

// AddDraft adds a model-less framed draft view for manual 2D geometry.
//
// mcp:tool drawing_add_draft_view
// mcp:summary Add a draft view: a model-less framed container (widthMm × heightMm, sheet mm) at centerXmm/centerYmm for manually-drawn 2D geometry.
func (d DrawingViews) AddDraft(args wire.AddDraftViewArgs) (wire.ViewResult, error) {
	return call[wire.ViewResult](d.c, wire.MethodDrawingViewsAddDraft, args)
}

// Delete removes the named view (and any views projected from it).
//
// mcp:tool drawing_delete_view
// mcp:summary Delete a drawing view by name (deleting a base view also removes views projected from it).
func (d DrawingViews) Delete(args wire.DeleteViewArgs) (wire.ListDrawingViewsResult, error) {
	return call[wire.ListDrawingViewsResult](d.c, wire.MethodDrawingViewsDelete, args)
}

// SetLabel changes any subset of a view's label — free text, the show-label/name/scale flags, and
// the caption position — leaving the unset ones alone (#1983).
//
// mcp:tool set_view_label
// mcp:summary Change a drawing view's label (text / showLabel / showName / showScale / position); unset fields unchanged.
func (d DrawingViews) SetLabel(args wire.SetViewLabelArgs) (wire.ListDrawingViewsResult, error) {
	return call[wire.ListDrawingViewsResult](d.c, wire.MethodDrawingViewsSetLabel, args)
}

// SetDisplay changes a view's edge-display toggles — currently the tangent-edge (fillet/blend
// transition) display — leaving the unset ones alone (#1984).
//
// mcp:tool set_view_display
// mcp:summary Change a drawing view's edge display: displayTangentEdges=false drops smooth tangent (fillet/blend) edges. Unset fields unchanged.
func (d DrawingViews) SetDisplay(args wire.SetViewDisplayArgs) (wire.ListDrawingViewsResult, error) {
	return call[wire.ListDrawingViewsResult](d.c, wire.MethodDrawingViewsSetDisplay, args)
}

// Rotate sets a view's rotation about its centre (degrees, CCW positive), rotating its curves (#1988).
//
// mcp:tool rotate_view
// mcp:summary Rotate a drawing view about its centre to angleDeg degrees (CCW positive).
func (d DrawingViews) Rotate(args wire.RotateViewArgs) (wire.ListDrawingViewsResult, error) {
	return call[wire.ListDrawingViewsResult](d.c, wire.MethodDrawingViewsRotate, args)
}

// Align locks a view to an anchor view on a shared axis (horizontal shares Y, vertical shares X) so
// moving the anchor drags it, or frees it with inPosition (#1988).
//
// mcp:tool align_view
// mcp:summary Align a drawing view to an anchor: "horizontal" (shared Y), "vertical" (shared X), or "inPosition" (free). Optional justification (centered/fixed).
func (d DrawingViews) Align(args wire.AlignViewArgs) (wire.ListDrawingViewsResult, error) {
	return call[wire.ListDrawingViewsResult](d.c, wire.MethodDrawingViewsAlign, args)
}

// AddCrop clips a view to a rectangular or circular fence (sheet mm), keeping the view's scale,
// with an optional continuous/zigzag break-mark boundary (#1987).
//
// mcp:tool drawing_add_view_crop
// mcp:summary Crop a drawing view to a fence (shape=rectangle x0,y0,x1,y1 | circle circleXmm,circleYmm,radiusMm; sheet mm), dropping curves outside it. breakMark=none|continuous|zigzag draws the boundary. The view keeps its scale (unlike a detail view).
func (d DrawingViews) AddCrop(args wire.AddViewCropArgs) (wire.ListDrawingViewsResult, error) {
	return call[wire.ListDrawingViewsResult](d.c, wire.MethodDrawingViewsAddCrop, args)
}

// RemoveCrop drops every crop on a view, restoring its full curve set (#1987).
//
// mcp:tool drawing_remove_view_crop
// mcp:summary Remove all crops from a drawing view, restoring its full (uncropped) curve set.
func (d DrawingViews) RemoveCrop(args wire.RemoveViewCropArgs) (wire.ListDrawingViewsResult, error) {
	return call[wire.ListDrawingViewsResult](d.c, wire.MethodDrawingViewsRemoveCrop, args)
}

// Curves returns a view's drawing curves — the projected edge segments classified visible
// (solid) or hidden (dashed), in sheet millimetres.
//
// mcp:tool drawing_view_curves
// mcp:summary Read a view's hidden-line drawing curves: 2D segments (sheet mm) flagged visible (solid) or hidden (dashed), each carrying the source model edge's reference key.
func (d DrawingViews) Curves(args wire.ViewCurvesArgs) (wire.ViewCurvesResult, error) {
	return call[wire.ViewCurvesResult](d.c, wire.MethodDrawingViewsCurves, args)
}
