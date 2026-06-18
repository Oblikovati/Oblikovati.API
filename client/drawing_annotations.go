// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The drawing annotations operation group (M14-F02 #813): place a centre-of-gravity marker on a
// view (driven by the referenced model's mass properties) and revision-cloud markup over a sheet
// region.

// DrawingAnnotations is the drawing-annotations operation group.
type DrawingAnnotations struct{ c *Client }

// DrawingAnnotations returns the drawing-annotations operation group.
func (c *Client) DrawingAnnotations() DrawingAnnotations { return DrawingAnnotations{c} }

// List returns the active sheet's annotations.
//
// mcp:tool drawing_list_annotations
// mcp:summary List the active sheet's drawing annotations (name, kind = cog|revisionCloud, the view a CoG marker is on, a revision cloud's tag, and curve count).
func (d DrawingAnnotations) List() (wire.ListDrawingAnnotationsResult, error) {
	var r wire.ListDrawingAnnotationsResult
	return r, d.c.call(wire.MethodDrawingAnnotationsList, struct{}{}, &r)
}

// AddCoGMarker adds a centre-of-gravity marker on a view, positioned at the model's centre of mass.
//
// mcp:tool drawing_add_cog_marker
// mcp:summary Add a centre-of-gravity marker on a drawing view (viewName); it is placed at the referenced model's centre of mass projected into that view and updates with the model.
func (d DrawingAnnotations) AddCoGMarker(args wire.AddCoGMarkerArgs) (wire.AnnotationResult, error) {
	var r wire.AnnotationResult
	return r, d.c.call(wire.MethodDrawingAnnotationsAddCoG, args, &r)
}

// AddRevisionCloud adds a scalloped revision cloud over a sheet region.
//
// mcp:tool drawing_add_revision_cloud
// mcp:summary Add a revision cloud (scalloped markup) over the sheet rectangle xmm/ymm/widthMm/heightMm, with an optional revision tag.
func (d DrawingAnnotations) AddRevisionCloud(args wire.AddRevisionCloudArgs) (wire.AnnotationResult, error) {
	var r wire.AnnotationResult
	return r, d.c.call(wire.MethodDrawingAnnotationsAddRevisionCloud, args, &r)
}

// AddCenterMarks adds a centre mark (crosshair) at every circular model edge's centre in a view.
//
// mcp:tool drawing_add_center_marks
// mcp:summary Add a centre mark (crosshair) at the centre of every circular model edge in a drawing view (the auto centre-mark-all-holes action). Each mark attaches to its edge and re-projects when the model changes.
func (d DrawingAnnotations) AddCenterMarks(args wire.AddCenterMarksArgs) (wire.CenterMarksResult, error) {
	var r wire.CenterMarksResult
	return r, d.c.call(wire.MethodDrawingAnnotationsAddCenterMarks, args, &r)
}

// Delete removes the named annotation.
//
// mcp:tool drawing_delete_annotation
// mcp:summary Delete a drawing annotation by name.
func (d DrawingAnnotations) Delete(args wire.DeleteAnnotationArgs) (wire.ListDrawingAnnotationsResult, error) {
	var r wire.ListDrawingAnnotationsResult
	return r, d.c.call(wire.MethodDrawingAnnotationsDelete, args, &r)
}
