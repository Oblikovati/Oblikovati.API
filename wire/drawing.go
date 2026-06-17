// SPDX-License-Identifier: Apache-2.0

package wire

// Drawing document wire DTOs (M14-F01, Oblikovati/Oblikovati#384). These cross the
// router for the active drawing document: list/add/remove sheets, pick the active
// sheet, set the primary referenced model (the document whose iProperties feed the
// title block), and read a title block's resolved field values. Sheet sizes and
// orientations cross as their types.SheetSize / types.SheetOrientation wire spellings
// ("a3", "ansiC", "landscape").

// SheetInfo is the JSON shape of one drawing sheet.
type SheetInfo struct {
	Name          string  `json:"name"`
	Size          string  `json:"size"`        // types.SheetSize spelling ("a3", "custom")
	Orientation   string  `json:"orientation"` // types.SheetOrientation spelling
	WidthMM       float64 `json:"widthMm"`     // laid-out width (orientation applied)
	HeightMM      float64 `json:"heightMm"`    // laid-out height
	Active        bool    `json:"active"`
	HasBorder     bool    `json:"hasBorder"`
	HasTitleBlock bool    `json:"hasTitleBlock"`
}

// ListSheetsResult is the response of [MethodDrawingListSheets]: every sheet (the
// active one flagged) and the drawing's primary referenced model, if set.
type ListSheetsResult struct {
	Sheets         []SheetInfo `json:"sheets"`
	ModelReference string      `json:"modelReference,omitempty"`
}

// AddSheetArgs is the request of [MethodDrawingAddSheet]. Size is a types.SheetSize
// spelling; for "custom" (or an empty Size) WidthMM/HeightMM give the dimensions. An
// empty Name lets the host assign the next "Sheet:N" name.
type AddSheetArgs struct {
	Name        string  `json:"name,omitempty"`
	Size        string  `json:"size,omitempty"`
	Orientation string  `json:"orientation,omitempty"`
	WidthMM     float64 `json:"widthMm,omitempty"`
	HeightMM    float64 `json:"heightMm,omitempty"`
}

// SheetResult is the response of [MethodDrawingAddSheet] / [MethodDrawingSetActiveSheet]:
// the affected sheet.
type SheetResult struct {
	Sheet SheetInfo `json:"sheet"`
}

// RemoveSheetArgs is the request of [MethodDrawingRemoveSheet]: the sheet to remove
// (a drawing must keep at least one sheet).
type RemoveSheetArgs struct {
	Name string `json:"name"`
}

// SetActiveSheetArgs is the request of [MethodDrawingSetActiveSheet]: the sheet to
// make active.
type SetActiveSheetArgs struct {
	Name string `json:"name"`
}

// SetModelReferenceArgs is the request of [MethodDrawingSetModelReference]: the full
// document name of the model the drawing documents (its title-block fields resolve
// against this model's iProperties). An empty name clears the reference.
type SetModelReferenceArgs struct {
	FullDocumentName string `json:"fullDocumentName"`
}

// SetModelReferenceResult is the response of [MethodDrawingSetModelReference]: the
// reference as stored (empty when cleared).
type SetModelReferenceResult struct {
	ModelReference string `json:"modelReference"`
}

// TitleBlockFieldsArgs is the request of [MethodDrawingTitleBlockFields]: the sheet
// whose title-block fields to resolve. An empty Sheet uses the active sheet.
type TitleBlockFieldsArgs struct {
	Sheet string `json:"sheet,omitempty"`
}

// TitleBlockField is one resolved title-block field: its name, the resolved text, and
// the source it resolved from (a "Set:Property" iProperty token, or "" for static text).
type TitleBlockField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Source string `json:"source,omitempty"`
}

// TitleBlockFieldsResult is the response of [MethodDrawingTitleBlockFields]: the title
// block's definition name and its resolved fields. Fields is empty when the sheet has
// no title block.
type TitleBlockFieldsResult struct {
	DefinitionName string            `json:"definitionName"`
	Fields         []TitleBlockField `json:"fields"`
}

// ExportDrawingDXFArgs is the request of [MethodDrawingExportDXF]: write the active sheet to
// the DXF file at Path. Version is a types.DXFVersion spelling ("r2000"/"r2018"; "" ⇒ r2000).
type ExportDrawingDXFArgs struct {
	Path    string `json:"path"`
	Version string `json:"version,omitempty"`
}

// ExportDrawingDXFResult is the response of [MethodDrawingExportDXF]: the file written and the
// number of DXF entities (view edges, border and title-block lines/text) in it.
type ExportDrawingDXFResult struct {
	Path     string `json:"path"`
	Entities int    `json:"entities"`
}
