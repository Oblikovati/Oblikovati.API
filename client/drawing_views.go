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
	var r wire.ListDrawingViewsResult
	return r, d.c.call(wire.MethodDrawingViewsList, struct{}{}, &r)
}

// AddBase adds a base view of the referenced model at the given orientation/scale/style,
// centred on the active sheet, and computes its hidden-line curves.
//
// mcp:tool drawing_add_base_view
// mcp:summary Add a base view of the drawing's referenced model (orientation = front|top|right|back|left|bottom|iso, style = hiddenLine|wireframe|shaded, scale e.g. 0.5 for 1:2) centred at centerXmm/centerYmm; computes visible/hidden edges.
func (d DrawingViews) AddBase(args wire.AddBaseViewArgs) (wire.ViewResult, error) {
	var r wire.ViewResult
	return r, d.c.call(wire.MethodDrawingViewsAddBase, args, &r)
}

// AddProjected adds a view projected from a base view in the given direction.
//
// mcp:tool drawing_add_projected_view
// mcp:summary Add a projected view off a base view (direction = right|left|up|down), inheriting the base's scale/style; placed at centerXmm/centerYmm.
func (d DrawingViews) AddProjected(args wire.AddProjectedViewArgs) (wire.ViewResult, error) {
	var r wire.ViewResult
	return r, d.c.call(wire.MethodDrawingViewsAddProjected, args, &r)
}

// AddAuxiliary adds a view projected perpendicular to a fold line on a parent view.
//
// mcp:tool drawing_add_auxiliary_view
// mcp:summary Add an auxiliary view off a parent view: projected perpendicular to a fold line at foldAngleDeg (0 folds down like a top view, 90 folds to the side), inheriting the parent's scale/style; placed at centerXmm/centerYmm. Shows an inclined face true-size.
func (d DrawingViews) AddAuxiliary(args wire.AddAuxiliaryViewArgs) (wire.ViewResult, error) {
	var r wire.ViewResult
	return r, d.c.call(wire.MethodDrawingViewsAddAuxiliary, args, &r)
}

// Delete removes the named view (and any views projected from it).
//
// mcp:tool drawing_delete_view
// mcp:summary Delete a drawing view by name (deleting a base view also removes views projected from it).
func (d DrawingViews) Delete(args wire.DeleteViewArgs) (wire.ListDrawingViewsResult, error) {
	var r wire.ListDrawingViewsResult
	return r, d.c.call(wire.MethodDrawingViewsDelete, args, &r)
}

// Curves returns a view's drawing curves — the projected edge segments classified visible
// (solid) or hidden (dashed), in sheet millimetres.
//
// mcp:tool drawing_view_curves
// mcp:summary Read a view's hidden-line drawing curves: 2D segments (sheet mm) flagged visible (solid) or hidden (dashed), each carrying the source model edge's reference key.
func (d DrawingViews) Curves(args wire.ViewCurvesArgs) (wire.ViewCurvesResult, error) {
	var r wire.ViewCurvesResult
	return r, d.c.call(wire.MethodDrawingViewsCurves, args, &r)
}
