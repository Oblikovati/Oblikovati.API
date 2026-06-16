// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The flat-pattern operation group (M13-F05, Oblikovati/Oblikovati#635): manage the developed
// flat's named orientations — saved alignment states that frame it for drawing views and
// export and drive its reported length/width.

// FlatPattern is the flat-pattern operation group.
type FlatPattern struct{ c *Client }

// FlatPattern returns the flat-pattern operation group.
func (c *Client) FlatPattern() FlatPattern { return FlatPattern{c} }

// ListOrientations returns the active part's flat-pattern orientations, the active one flagged.
//
// mcp:tool flat_pattern_list_orientations
// mcp:summary List the active sheet-metal part's flat-pattern orientations (name, alignment, flips, and the flat's length/width under each); the active orientation is flagged.
func (f FlatPattern) ListOrientations() (wire.OrientationsResult, error) {
	var r wire.OrientationsResult
	return r, f.c.call(wire.MethodFlatPatternListOrientations, struct{}{}, &r)
}

// AddOrientation creates a named orientation (optionally activating it).
//
// mcp:tool flat_pattern_add_orientation
// mcp:summary Add a named flat-pattern orientation (alignment type horizontal|vertical, alignment rotation in degrees, optional alignment-axis reference key, flip flags); set activate to make it current.
func (f FlatPattern) AddOrientation(args wire.AddOrientationArgs) (wire.OrientationResult, error) {
	var r wire.OrientationResult
	return r, f.c.call(wire.MethodFlatPatternAddOrientation, args, &r)
}

// ActivateOrientation makes the named orientation current (it frames the flat and drives
// length/width).
//
// mcp:tool flat_pattern_activate_orientation
// mcp:summary Activate a flat-pattern orientation by name; it frames the flat for export/drawing and drives the reported length/width.
func (f FlatPattern) ActivateOrientation(args wire.ActivateOrientationArgs) (wire.OrientationResult, error) {
	var r wire.OrientationResult
	return r, f.c.call(wire.MethodFlatPatternActivateOrientation, args, &r)
}

// DeleteOrientation removes the named orientation (the default orientation cannot be deleted).
//
// mcp:tool flat_pattern_delete_orientation
// mcp:summary Delete a flat-pattern orientation by name (the default orientation cannot be deleted).
func (f FlatPattern) DeleteOrientation(args wire.DeleteOrientationArgs) (wire.OrientationsResult, error) {
	var r wire.OrientationsResult
	return r, f.c.call(wire.MethodFlatPatternDeleteOrientation, args, &r)
}

// EdgesOfType returns the developed flat's classified fold/tangent edges, optionally filtered
// to one type (bendUp/bendDown/tangent).
//
// mcp:tool flat_pattern_edges_of_type
// mcp:summary List the developed flat's classified edges (bend-up/bend-down fold lines, tangent lines), optionally filtered to one type — the bend layer of a flat-pattern drawing/export.
func (f FlatPattern) EdgesOfType(args wire.EdgesOfTypeArgs) (wire.EdgesResult, error) {
	var r wire.EdgesResult
	return r, f.c.call(wire.MethodFlatPatternEdgesOfType, args, &r)
}

// Faces returns the developed flat's classified faces (front/back) with their areas.
//
// mcp:tool flat_pattern_faces
// mcp:summary Report the developed flat's classified faces — the front (top) and back (bottom) faces and their developed areas.
func (f FlatPattern) Faces() (wire.FacesResult, error) {
	var r wire.FacesResult
	return r, f.c.call(wire.MethodFlatPatternFaces, struct{}{}, &r)
}

// MapEntity maps a topology entity between the folded model and the developed flat by
// reference key (folded→flat when ToFlat, else flat→folded), so a drawing dimension or
// selection survives recompute.
//
// mcp:tool flat_pattern_map_entity
// mcp:summary Map a topology entity (by reference key) between the folded sheet-metal model and its developed flat pattern (set toFlat for folded→flat, else flat→folded). Face-level: top/bottom faces map to the flat front/back face.
func (f FlatPattern) MapEntity(args wire.MapEntityArgs) (wire.MapEntityResult, error) {
	var r wire.MapEntityResult
	return r, f.c.call(wire.MethodFlatPatternMapEntity, args, &r)
}

// ListPlates returns the developed flat's plates — one per connected flat region — with each
// plate's extents/area under the active orientation.
//
// mcp:tool flat_pattern_list_plates
// mcp:summary List the developed flat's plates (one per connected flat region of the sheet-metal part) with each plate's length/width/area under the active orientation.
func (f FlatPattern) ListPlates() (wire.PlatesResult, error) {
	var r wire.PlatesResult
	return r, f.c.call(wire.MethodFlatPatternListPlates, struct{}{}, &r)
}

// GetSettings returns the part's flat-pattern settings.
//
// mcp:tool flat_pattern_get_settings
// mcp:summary Report the active sheet-metal part's flat-pattern settings (deferUpdate: whether the flat only recomputes on demand).
func (f FlatPattern) GetSettings() (wire.SettingsResult, error) {
	var r wire.SettingsResult
	return r, f.c.call(wire.MethodFlatPatternGetSettings, struct{}{}, &r)
}

// SetSettings edits the part's flat-pattern settings.
//
// mcp:tool flat_pattern_set_settings
// mcp:summary Edit the flat-pattern settings (deferUpdate: suppress the automatic flat recompute so a heavy flat develops only on demand). Returns the updated settings.
func (f FlatPattern) SetSettings(args wire.SetSettingsArgs) (wire.SettingsResult, error) {
	var r wire.SettingsResult
	return r, f.c.call(wire.MethodFlatPatternSetSettings, args, &r)
}

// ListBendOrder returns the part's bends in their press-brake sequence (each with its 1-based
// order, angle and radius).
//
// mcp:tool flat_pattern_list_bend_order
// mcp:summary List the sheet-metal part's bends in press-brake sequence order (feature, 1-based order, angle, radius) — the bend-order annotation shown on the flat pattern.
func (f FlatPattern) ListBendOrder() (wire.BendOrderResult, error) {
	var r wire.BendOrderResult
	return r, f.c.call(wire.MethodFlatPatternListBendOrder, struct{}{}, &r)
}

// SetBendOrder sets the bend sequence: Order lists the bend features by name; omitted bends
// keep their natural order after the listed ones (an empty Order resets to natural).
//
// mcp:tool flat_pattern_set_bend_order
// mcp:summary Set the press-brake bend sequence by listing the bend features in order (omitted bends keep natural order after them; empty resets to creation order). Returns the new order.
func (f FlatPattern) SetBendOrder(args wire.SetBendOrderArgs) (wire.BendOrderResult, error) {
	var r wire.BendOrderResult
	return r, f.c.call(wire.MethodFlatPatternSetBendOrder, args, &r)
}
