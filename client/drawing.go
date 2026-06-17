// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The drawing operation group (M14-F01, Oblikovati/Oblikovati#384): manage the active
// drawing document's sheets, choose the active sheet, point the drawing at the model it
// documents (whose iProperties drive the title block), and read a title block's resolved
// field values.

// Drawing is the drawing-document operation group.
type Drawing struct{ c *Client }

// Drawing returns the drawing-document operation group.
func (c *Client) Drawing() Drawing { return Drawing{c} }

// ListSheets returns the active drawing's sheets (the active one flagged) and its primary
// referenced model, if set.
//
// mcp:tool drawing_list_sheets
// mcp:summary List the active drawing document's sheets (name, size, orientation, width/height in mm, border/title-block presence; the active sheet is flagged) and the primary referenced model.
func (d Drawing) ListSheets() (wire.ListSheetsResult, error) {
	var r wire.ListSheetsResult
	return r, d.c.call(wire.MethodDrawingListSheets, struct{}{}, &r)
}

// AddSheet adds a sheet of the given standard size (or a custom width×height) and
// orientation, and makes it active.
//
// mcp:tool drawing_add_sheet
// mcp:summary Add a sheet to the active drawing (size = a0..a4|ansiA..ansiE|custom, orientation = portrait|landscape; for custom give widthMm/heightMm); an empty name auto-assigns. Becomes the active sheet.
func (d Drawing) AddSheet(args wire.AddSheetArgs) (wire.SheetResult, error) {
	var r wire.SheetResult
	return r, d.c.call(wire.MethodDrawingAddSheet, args, &r)
}

// RemoveSheet deletes the named sheet (a drawing must keep at least one sheet).
//
// mcp:tool drawing_remove_sheet
// mcp:summary Remove the named sheet from the active drawing (a drawing must keep at least one sheet).
func (d Drawing) RemoveSheet(args wire.RemoveSheetArgs) (wire.ListSheetsResult, error) {
	var r wire.ListSheetsResult
	return r, d.c.call(wire.MethodDrawingRemoveSheet, args, &r)
}

// SetActiveSheet makes the named sheet the active sheet.
//
// mcp:tool drawing_set_active_sheet
// mcp:summary Make the named sheet the active sheet of the active drawing.
func (d Drawing) SetActiveSheet(args wire.SetActiveSheetArgs) (wire.SheetResult, error) {
	var r wire.SheetResult
	return r, d.c.call(wire.MethodDrawingSetActiveSheet, args, &r)
}

// SetModelReference points the drawing at the model it documents; its title-block fields
// resolve against that model's iProperties. An empty name clears the reference.
//
// mcp:tool drawing_set_model_reference
// mcp:summary Set the drawing's primary referenced model by full document name (its iProperties feed the title block); an empty name clears the reference.
func (d Drawing) SetModelReference(args wire.SetModelReferenceArgs) (wire.SetModelReferenceResult, error) {
	var r wire.SetModelReferenceResult
	return r, d.c.call(wire.MethodDrawingSetModelReference, args, &r)
}

// TitleBlockFields returns a sheet's title-block definition name and resolved field
// values (an empty Sheet uses the active sheet).
//
// mcp:tool drawing_title_block_fields
// mcp:summary Read a sheet's title-block resolved fields (name, value, and source iProperty token); an empty sheet uses the active sheet. Fields are empty when the sheet has no title block.
func (d Drawing) TitleBlockFields(args wire.TitleBlockFieldsArgs) (wire.TitleBlockFieldsResult, error) {
	var r wire.TitleBlockFieldsResult
	return r, d.c.call(wire.MethodDrawingTitleBlockFields, args, &r)
}
