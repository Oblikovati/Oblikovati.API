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
	// Sheet authoring (#1989). Revision is the sheet's revision string; BorderHZones/VZones report a
	// zoned border's grid (0 ⇒ plain); TitleBlockLocation is the title block's corner.
	Revision           string `json:"revision,omitempty"`
	BorderHZones       int    `json:"borderHZones,omitempty"`
	BorderVZones       int    `json:"borderVZones,omitempty"`
	TitleBlockLocation string `json:"titleBlockLocation,omitempty"`
}

// AddDefaultBorderArgs is the request of [MethodDrawingAddDefaultBorder]: replace the active sheet's
// border with a zoned one — HZones columns × VZones rows, labelled per HLabelMode / VLabelMode
// (types.BorderLabelMode: "alphabetical"/"numeric"/"none"; "" ⇒ alphabetical). Sheet names the target
// sheet ("" ⇒ active) (#1989).
type AddDefaultBorderArgs struct {
	Sheet      string `json:"sheet,omitempty"`
	HZones     int    `json:"hZones"`
	VZones     int    `json:"vZones"`
	HLabelMode string `json:"hLabelMode,omitempty"`
	VLabelMode string `json:"vLabelMode,omitempty"`
}

// SetTitleBlockArgs is the request of [MethodDrawingSetTitleBlock]: move a sheet's title block to a
// corner (types.TitleBlockLocation: "bottomRight"/"bottomLeft"/"topLeft"/"topRight"; "" ⇒ bottomRight),
// seeding the default block when the sheet has none. Sheet names the target ("" ⇒ active) (#1989).
type SetTitleBlockArgs struct {
	Sheet    string `json:"sheet,omitempty"`
	Location string `json:"location,omitempty"`
}

// SetSheetRevisionArgs is the request of [MethodDrawingSetSheetRevision]: set a sheet's revision
// string. Sheet names the target ("" ⇒ active) (#1989).
type SetSheetRevisionArgs struct {
	Sheet    string `json:"sheet,omitempty"`
	Revision string `json:"revision"`
}

// DefineSheetFormatArgs is the request of [MethodDrawingDefineSheetFormat]: register a reusable sheet
// format under Name — a size/orientation, an optional zoned border (HZones×VZones, 0 ⇒ plain) and a
// title-block corner — that [MethodDrawingAddSheetUsingFormat] stamps new sheets from (#1989).
type DefineSheetFormatArgs struct {
	Name               string  `json:"name"`
	Size               string  `json:"size,omitempty"`
	Orientation        string  `json:"orientation,omitempty"`
	WidthMM            float64 `json:"widthMm,omitempty"`
	HeightMM           float64 `json:"heightMm,omitempty"`
	HZones             int     `json:"hZones,omitempty"`
	VZones             int     `json:"vZones,omitempty"`
	HLabelMode         string  `json:"hLabelMode,omitempty"`
	VLabelMode         string  `json:"vLabelMode,omitempty"`
	TitleBlockLocation string  `json:"titleBlockLocation,omitempty"`
}

// AddSheetUsingFormatArgs is the request of [MethodDrawingAddSheetUsingFormat]: add a sheet named Name
// stamped from the registered format Format (#1989).
type AddSheetUsingFormatArgs struct {
	Name   string `json:"name,omitempty"`
	Format string `json:"format"`
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
