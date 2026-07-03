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
	return call[wire.ListSheetsResult](d.c, wire.MethodDrawingListSheets, struct{}{})
}

// AddSheet adds a sheet of the given standard size (or a custom width×height) and
// orientation, and makes it active.
//
// mcp:tool drawing_add_sheet
// mcp:summary Add a sheet to the active drawing (size = a0..a4|ansiA..ansiE|custom, orientation = portrait|landscape; for custom give widthMm/heightMm); an empty name auto-assigns. Becomes the active sheet.
func (d Drawing) AddSheet(args wire.AddSheetArgs) (wire.SheetResult, error) {
	return call[wire.SheetResult](d.c, wire.MethodDrawingAddSheet, args)
}

// RemoveSheet deletes the named sheet (a drawing must keep at least one sheet).
//
// mcp:tool drawing_remove_sheet
// mcp:summary Remove the named sheet from the active drawing (a drawing must keep at least one sheet).
func (d Drawing) RemoveSheet(args wire.RemoveSheetArgs) (wire.ListSheetsResult, error) {
	return call[wire.ListSheetsResult](d.c, wire.MethodDrawingRemoveSheet, args)
}

// SetActiveSheet makes the named sheet the active sheet.
//
// mcp:tool drawing_set_active_sheet
// mcp:summary Make the named sheet the active sheet of the active drawing.
func (d Drawing) SetActiveSheet(args wire.SetActiveSheetArgs) (wire.SheetResult, error) {
	return call[wire.SheetResult](d.c, wire.MethodDrawingSetActiveSheet, args)
}

// SetModelReference points the drawing at the model it documents; its title-block fields
// resolve against that model's iProperties. An empty name clears the reference.
//
// mcp:tool drawing_set_model_reference
// mcp:summary Set the drawing's primary referenced model by full document name (its iProperties feed the title block); an empty name clears the reference.
func (d Drawing) SetModelReference(args wire.SetModelReferenceArgs) (wire.SetModelReferenceResult, error) {
	return call[wire.SetModelReferenceResult](d.c, wire.MethodDrawingSetModelReference, args)
}

// TitleBlockFields returns a sheet's title-block definition name and resolved field
// values (an empty Sheet uses the active sheet).
//
// mcp:tool drawing_title_block_fields
// mcp:summary Read a sheet's title-block resolved fields (name, value, and source iProperty token); an empty sheet uses the active sheet. Fields are empty when the sheet has no title block.
func (d Drawing) TitleBlockFields(args wire.TitleBlockFieldsArgs) (wire.TitleBlockFieldsResult, error) {
	return call[wire.TitleBlockFieldsResult](d.c, wire.MethodDrawingTitleBlockFields, args)
}

// ExportDXF writes the active sheet to a DXF file — its views' visible/hidden edges, border and
// title block on named layers.
//
// mcp:tool drawing_export_dxf
// mcp:summary Export the active drawing sheet to a DXF file (path, version r2000|r2018): view edges on Visible/Hidden layers, the border, and the title-block grid + field text.
func (d Drawing) ExportDXF(args wire.ExportDrawingDXFArgs) (wire.ExportDrawingDXFResult, error) {
	return call[wire.ExportDrawingDXFResult](d.c, wire.MethodDrawingExportDXF, args)
}
