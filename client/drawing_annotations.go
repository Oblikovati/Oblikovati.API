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
	return call[wire.ListDrawingAnnotationsResult](d.c, wire.MethodDrawingAnnotationsList, struct{}{})
}

// AddCoGMarker adds a centre-of-gravity marker on a view, positioned at the model's centre of mass.
//
// mcp:tool drawing_add_cog_marker
// mcp:summary Add a centre-of-gravity marker on a drawing view (viewName); it is placed at the referenced model's centre of mass projected into that view and updates with the model.
func (d DrawingAnnotations) AddCoGMarker(args wire.AddCoGMarkerArgs) (wire.AnnotationResult, error) {
	return call[wire.AnnotationResult](d.c, wire.MethodDrawingAnnotationsAddCoG, args)
}

// AddRevisionCloud adds a scalloped revision cloud over a sheet region.
//
// mcp:tool drawing_add_revision_cloud
// mcp:summary Add a revision cloud (scalloped markup) over the sheet rectangle xmm/ymm/widthMm/heightMm, with an optional revision tag.
func (d DrawingAnnotations) AddRevisionCloud(args wire.AddRevisionCloudArgs) (wire.AnnotationResult, error) {
	return call[wire.AnnotationResult](d.c, wire.MethodDrawingAnnotationsAddRevisionCloud, args)
}

// AddCenterMarks adds a centre mark (crosshair) at every circular model edge's centre in a view.
//
// mcp:tool drawing_add_center_marks
// mcp:summary Add a centre mark (crosshair) at the centre of every circular model edge in a drawing view (the auto centre-mark-all-holes action). Each mark attaches to its edge and re-projects when the model changes.
func (d DrawingAnnotations) AddCenterMarks(args wire.AddCenterMarksArgs) (wire.CenterMarksResult, error) {
	return call[wire.CenterMarksResult](d.c, wire.MethodDrawingAnnotationsAddCenterMarks, args)
}

// AddCenterlines adds the horizontal+vertical symmetry centerlines through a view's centre.
//
// mcp:tool drawing_add_centerlines
// mcp:summary Add the horizontal and vertical dash-dot symmetry centerlines through a drawing view's centre, spanning its extent. The lines re-derive from the view's bounds, so they track the model.
func (d DrawingAnnotations) AddCenterlines(args wire.AddCenterlinesArgs) (wire.AnnotationResult, error) {
	return call[wire.AnnotationResult](d.c, wire.MethodDrawingAnnotationsAddCenterlines, args)
}

// AddFeatureControlFrame adds a GD&T feature control frame at a sheet point.
//
// mcp:tool drawing_add_feature_control_frame
// mcp:summary Add a GD&T feature control frame at a sheet point (xmm/ymm): a boxed geometric-tolerance callout with a characteristic (types.GeometricCharacteristic wire spelling, e.g. position|flatness|perpendicularity|parallelism|straightness|circularity), a tolerance value, and ordered datum reference letters.
func (d DrawingAnnotations) AddFeatureControlFrame(args wire.AddFeatureControlFrameArgs) (wire.AnnotationResult, error) {
	return call[wire.AnnotationResult](d.c, wire.MethodDrawingAnnotationsAddFCF, args)
}

// AddDatumFeature adds a GD&T datum feature symbol (a lettered box + datum triangle) at a sheet point.
//
// mcp:tool drawing_add_datum_feature
// mcp:summary Add a GD&T datum feature symbol at a sheet point (xmm/ymm): the datum letter (e.g. "A") in a box with a filled datum triangle, marking a datum that feature control frames reference.
func (d DrawingAnnotations) AddDatumFeature(args wire.AddDatumFeatureArgs) (wire.AnnotationResult, error) {
	return call[wire.AnnotationResult](d.c, wire.MethodDrawingAnnotationsAddDatum, args)
}

// AddSurfaceTexture adds an ISO 1302 surface texture symbol (a roughness checkmark) at a sheet point.
//
// mcp:tool drawing_add_surface_texture
// mcp:summary Add an ISO 1302 surface texture symbol at a sheet point (xmm/ymm): the roughness checkmark glyph with a finish value (roughness, e.g. "1.6"). materialRemoval = any (basic, default) | required (machined, with bar) | prohibited (as-cast, with vertex circle).
func (d DrawingAnnotations) AddSurfaceTexture(args wire.AddSurfaceTextureArgs) (wire.AnnotationResult, error) {
	return call[wire.AnnotationResult](d.c, wire.MethodDrawingAnnotationsAddSurfaceText, args)
}

// AddPartsList adds a parts list table sourced from the referenced assembly's BOM at a sheet point.
//
// mcp:tool drawing_add_parts_list
// mcp:summary Add a parts list table at a sheet point (xmm/ymm = top-left): a grid sourced from the referenced assembly's parts-only BOM (item number, part number, description, quantity). The rowCount in the result is the number of BOM items; the table updates with the assembly.
func (d DrawingAnnotations) AddPartsList(args wire.AddPartsListArgs) (wire.AnnotationResult, error) {
	return call[wire.AnnotationResult](d.c, wire.MethodDrawingAnnotationsAddPartsList, args)
}

// AddBalloon adds a balloon (a circled parts-list item number with an optional leader) at a sheet point.
//
// mcp:tool drawing_add_balloon
// mcp:summary Add a balloon at a sheet point (xmm/ymm = circle centre): a circle holding the parts-list item number, with an optional leader to the component it tags (leaderXmm/leaderYmm). Balloons reference parts-list items.
func (d DrawingAnnotations) AddBalloon(args wire.AddBalloonArgs) (wire.AnnotationResult, error) {
	return call[wire.AnnotationResult](d.c, wire.MethodDrawingAnnotationsAddBalloon, args)
}

// AddHoleTable adds a hole table for a base view's circular edges at a sheet point.
//
// mcp:tool drawing_add_hole_table
// mcp:summary Add a hole table at a sheet point (xmm/ymm = top-left) listing every circular edge in a base view: HOLE / X / Y (from the view's datum origin) / ⌀ (diameter). The rowCount in the result is the hole count; the table updates with the model.
func (d DrawingAnnotations) AddHoleTable(args wire.AddHoleTableArgs) (wire.AnnotationResult, error) {
	return call[wire.AnnotationResult](d.c, wire.MethodDrawingAnnotationsAddHoleTable, args)
}

// AddRevisionTable adds a revision table (revision/date/description rows) at a sheet point.
//
// mcp:tool drawing_add_revision_table
// mcp:summary Add a revision table at a sheet point (xmm/ymm = top-left) with rows of {revision, date, description}, recording the drawing's change history. The rowCount in the result is the revision count.
func (d DrawingAnnotations) AddRevisionTable(args wire.AddRevisionTableArgs) (wire.AnnotationResult, error) {
	return call[wire.AnnotationResult](d.c, wire.MethodDrawingAnnotationsAddRevTable, args)
}

// AddRevisionTag adds a revision tag (a triangle holding a revision letter) at a sheet point.
//
// mcp:tool drawing_add_revision_tag
// mcp:summary Add a revision tag (a triangle holding the revision letter) centred at a sheet point (xmm/ymm), flagging where that revision changed the drawing.
func (d DrawingAnnotations) AddRevisionTag(args wire.AddRevisionTagArgs) (wire.AnnotationResult, error) {
	return call[wire.AnnotationResult](d.c, wire.MethodDrawingAnnotationsAddRevTag, args)
}

// AddNote adds a free text note (with an optional leader) at a sheet point.
//
// mcp:tool drawing_add_note
// mcp:summary Add a free text note anchored at a sheet point (xmm/ymm); if leaderXmm/leaderYmm are given, a leader is drawn from the note to that point.
func (d DrawingAnnotations) AddNote(args wire.AddDrawingNoteArgs) (wire.AnnotationResult, error) {
	return call[wire.AnnotationResult](d.c, wire.MethodDrawingAnnotationsAddNote, args)
}

// AddCustomTable adds a general-purpose table (arbitrary headers + rows) at a sheet point.
//
// mcp:tool drawing_add_custom_table
// mcp:summary Add a general-purpose table at a sheet point (xmm/ymm = top-left) with the given column headers and rows (each row's cells align to the headers). The rowCount in the result is the data-row count.
func (d DrawingAnnotations) AddCustomTable(args wire.AddCustomTableArgs) (wire.AnnotationResult, error) {
	return call[wire.AnnotationResult](d.c, wire.MethodDrawingAnnotationsAddCustomTable, args)
}

// AddHoleNotes adds a feature note on each hole in a base view: a leadered diameter callout.
//
// mcp:tool drawing_add_hole_notes
// mcp:summary Add hole notes to a base view (viewName): a leadered Ø-diameter callout computed from each hole's circular edge and re-resolved when the model changes. quantity "combined" groups holes by diameter into one "<n>x Ø<d>" callout per size (default "perHole" = one per hole). Optional format template with {d} (diameter) and {n} (count) placeholders, e.g. "Ø{d} THRU". The rowCount in the result is the callout count.
func (d DrawingAnnotations) AddHoleNotes(args wire.AddHoleNotesArgs) (wire.AnnotationResult, error) {
	return call[wire.AnnotationResult](d.c, wire.MethodDrawingAnnotationsAddHoleNotes, args)
}

// Delete removes the named annotation.
//
// mcp:tool drawing_delete_annotation
// mcp:summary Delete a drawing annotation by name.
func (d DrawingAnnotations) Delete(args wire.DeleteAnnotationArgs) (wire.ListDrawingAnnotationsResult, error) {
	return call[wire.ListDrawingAnnotationsResult](d.c, wire.MethodDrawingAnnotationsDelete, args)
}
