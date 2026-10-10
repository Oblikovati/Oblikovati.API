// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// Documents is the document-management operation group.
type Documents struct{ c *Client }

// Documents returns the document operation group.
func (c *Client) Documents() Documents { return Documents{c} }

// List returns every open document and which one is active.
//
// mcp:tool list_documents
// mcp:summary List open documents and which one is active.
func (d Documents) List() (wire.ListDocumentsResult, error) {
	return call[wire.ListDocumentsResult](d.c, wire.MethodDocumentsList, nil)
}

// Update recomputes the active document's out-of-date features (#139). With
// acceptErrorsAndContinue it succeeds and reports sick features; otherwise a sick feature fails.
//
// mcp:tool update_document
// mcp:summary Recompute the active document's out-of-date features; reports sick features.
func (d Documents) Update(acceptErrorsAndContinue bool) (wire.UpdateDocumentResult, error) {
	return call[wire.UpdateDocumentResult](d.c, wire.MethodDocumentsUpdate, wire.UpdateDocumentArgs{AcceptErrorsAndContinue: acceptErrorsAndContinue})
}

// Rebuild recomputes the active document's entire feature program as if all entities were
// dirtied (#139) — the full parametric rebuild.
//
// mcp:tool rebuild_document
// mcp:summary Recompute the active document's entire feature program (full rebuild).
func (d Documents) Rebuild(acceptErrorsAndContinue bool) (wire.UpdateDocumentResult, error) {
	return call[wire.UpdateDocumentResult](d.c, wire.MethodDocumentsRebuild, wire.UpdateDocumentArgs{AcceptErrorsAndContinue: acceptErrorsAndContinue})
}

// RequiresUpdate reports whether the active document has out-of-date features a recompute would
// change (#139) — the read-only "needs update" flag.
//
// mcp:tool document_requires_update
// mcp:summary Report whether the active document has out-of-date features needing a recompute.
func (d Documents) RequiresUpdate() (wire.RequiresUpdateResult, error) {
	return call[wire.RequiresUpdateResult](d.c, wire.MethodDocumentsRequiresUpdate, nil)
}

// Create makes a new document of the given kind active and returns it.
//
// mcp:tool create_document
// mcp:summary Create a new document (type: part|assembly|drawing|presentation) and make it active.
func (d Documents) Create(args wire.CreateDocumentArgs) (wire.DocumentInfo, error) {
	return call[wire.DocumentInfo](d.c, wire.MethodDocumentsCreate, args)
}

// Activate makes the document with the given session id active.
//
// mcp:tool activate_document
// mcp:summary Make the document with the given id active.
func (d Documents) Activate(id uint64) (wire.OKResult, error) {
	return call[wire.OKResult](d.c, wire.MethodDocumentsActivate, wire.ActivateDocumentArgs{ID: id})
}

// Close closes the document with the given session id; force discards unsaved
// changes instead of saving them first.
//
//	closed, err := c.Documents().Close(doc.ID, false)
//
// mcp:tool close_document
// mcp:summary Close the document with the given id. Set force:true to discard unsaved changes.
func (d Documents) Close(id uint64, force bool) (wire.CloseDocumentsResult, error) {
	return call[wire.CloseDocumentsResult](d.c, wire.MethodDocumentsClose, wire.CloseDocumentArgs{ID: id, Force: force})
}

// CloseAll closes every open document; force discards unsaved changes (the usual
// choice to reset to a clean session).
//
//	closed, err := c.Documents().CloseAll(true)
//
// mcp:tool close_all_documents
// mcp:summary Close every open document to start a clean session. Set force:true to discard unsaved changes.
func (d Documents) CloseAll(force bool) (wire.CloseDocumentsResult, error) {
	return call[wire.CloseDocumentsResult](d.c, wire.MethodDocumentsCloseAll, wire.CloseAllDocumentsArgs{Force: force})
}

// RegisterSubType declares a flavored document subtype over a base type; the
// flavor's lifecycle reaches the owner as client.operation push events (M05-F15).
//
//	client.Documents().RegisterSubType(wire.RegisterDocumentSubTypeArgs{
//	    ID: "com.x.sim.study", BaseType: "part", DisplayName: "Simulation Study",
//	})
//
// mcp:tool documents_register_sub_type
// mcp:summary Declares a flavored document subtype over a base type; the flavor's lifecycle reaches the owner as client.operation push events (M05-F15).
func (d Documents) RegisterSubType(args wire.RegisterDocumentSubTypeArgs) (wire.OKResult, error) {
	return call[wire.OKResult](d.c, wire.MethodDocumentsRegisterSubType, args)
}

// SubTypes returns the registered flavored subtypes.
//
// mcp:tool documents_list_sub_types
// mcp:summary Returns the registered flavored subtypes.
func (d Documents) SubTypes() (wire.ListDocumentSubTypesResult, error) {
	return call[wire.ListDocumentSubTypesResult](d.c, wire.MethodDocumentsListSubTypes, nil)
}

// FileReferences returns the document-side view of a document's file
// references: status plus the bridge to the resolved document (M03-F07).
//
// mcp:tool documents_list_file_references
// mcp:summary Returns the document-side view of a document's file references: status plus the bridge to the resolved document (M03-F07).
func (d Documents) FileReferences(id uint64) (wire.ListDocumentFileReferencesResult, error) {
	args := wire.ListDocumentFileReferencesArgs{Document: id}
	return call[wire.ListDocumentFileReferencesResult](d.c, wire.MethodDocumentsListFileReferences, args)
}

// Open loads the document at the given full document name (or returns the
// already-open one) and makes it active (#138).
//
//	info, err := c.Documents().Open(wire.OpenDocumentArgs{FullDocumentName: "/w/bracket.obk", Visible: true})
//
// mcp:tool documents_open
// mcp:summary Loads the document at the given full document name (or returns the already-open one) and makes it active (#138).
func (d Documents) Open(args wire.OpenDocumentArgs) (wire.DocumentInfo, error) {
	return call[wire.DocumentInfo](d.c, wire.MethodDocumentsOpen, args)
}

// Save writes the document at its current file binding (#138).
//
// mcp:tool documents_save
// mcp:summary Writes the document at its current file binding (#138).
func (d Documents) Save(id uint64) (wire.SaveDocumentResult, error) {
	return call[wire.SaveDocumentResult](d.c, wire.MethodDocumentsSave, wire.SaveDocumentArgs{Document: id})
}

// SaveAs writes the document under a new full document name, which becomes its
// identity (#138).
//
// mcp:tool documents_save_as
// mcp:summary Writes the document under a new full document name, which becomes its identity (#138).
func (d Documents) SaveAs(id uint64, newFullDocumentName string) (wire.SaveDocumentResult, error) {
	args := wire.SaveDocumentAsArgs{Document: id, NewFullDocumentName: newFullDocumentName}
	return call[wire.SaveDocumentResult](d.c, wire.MethodDocumentsSaveAs, args)
}

// SaveCopyAs writes a copy of the document to a target file without
// retargeting the in-memory document (M03-F09).
//
// mcp:tool documents_save_copy_as
// mcp:summary Writes a copy of the document to a target file without retargeting the in-memory document (M03-F09).
func (d Documents) SaveCopyAs(args wire.SaveCopyAsArgs) (wire.SaveDocumentResult, error) {
	return call[wire.SaveDocumentResult](d.c, wire.MethodDocumentsSaveCopyAs, args)
}

// BatchSave executes one save operation over several documents, continuing
// past per-item failures and returning per-file outcomes (M03-F09).
//
// mcp:tool documents_batch_save
// mcp:summary Executes one save operation over several documents, continuing past per-item failures and returning per-file outcomes (M03-F09).
func (d Documents) BatchSave(args wire.BatchSaveArgs) (wire.BatchSaveResult, error) {
	return call[wire.BatchSaveResult](d.c, wire.MethodDocumentsBatchSave, args)
}

// ListProperties returns every iProperty of the document, across all its sets (#156).
//
//	props, err := c.Documents().ListProperties(doc.ID)
//
// mcp:tool documents_list_properties
// mcp:summary Returns every iProperty of the document, across all its sets (#156).
func (d Documents) ListProperties(id uint64) (wire.ListPropertiesResult, error) {
	return call[wire.ListPropertiesResult](d.c, wire.MethodDocumentsListProperties, wire.ListPropertiesArgs{Document: id})
}

// GetProperty returns one document property addressed by its set and name (#156).
//
// mcp:tool documents_get_property
// mcp:summary Returns one document property addressed by its set and name (#156).
func (d Documents) GetProperty(id uint64, set, name string) (wire.PropertyResult, error) {
	args := wire.GetPropertyArgs{Document: id, Set: set, Name: name}
	return call[wire.PropertyResult](d.c, wire.MethodDocumentsGetProperty, args)
}

// SetProperty creates or replaces a document property's typed value, returning its new
// state (#156).
//
//	c.Documents().SetProperty(doc.ID, "Design Tracking Properties", "Part Number",
//	    types.StringVariant("BRK-001"))
//
// mcp:tool documents_set_property
// mcp:summary Creates or replaces a document property's typed value, returning its new state (#156).
func (d Documents) SetProperty(id uint64, set, name string, value types.Variant) (wire.PropertyResult, error) {
	args := wire.SetPropertyArgs{Document: id, Set: set, Name: name, Value: value}
	return call[wire.PropertyResult](d.c, wire.MethodDocumentsSetProperty, args)
}

// GetSketchSettings returns the document's persisted sketch-authoring defaults — the constraint-
// inference toggles and family priority the sketch tools read (#147).
//
//	s, err := c.Documents().GetSketchSettings(doc.ID)
//
// mcp:tool documents_get_sketch_settings
// mcp:summary Returns the document's persisted sketch settings (constraint inference toggles and priority) (#147).
func (d Documents) GetSketchSettings(id uint64) (wire.SketchSettingsResult, error) {
	return call[wire.SketchSettingsResult](d.c, wire.MethodDocumentGetSketchSettings, wire.GetSketchSettingsArgs{Document: id})
}

// SetSketchSettings replaces the document's sketch settings, returning their new state (#147).
//
//	c.Documents().SetSketchSettings(doc.ID, types.SketchSettings{
//	    InferConstraints: true, AutoApplyConstraints: false,
//	    ConstraintPriority: types.PriorityHorizontalVertical})
//
// mcp:tool documents_set_sketch_settings
// mcp:summary Replaces the document's sketch settings (constraint inference toggles and priority), returning their new state (#147).
func (d Documents) SetSketchSettings(id uint64, settings types.SketchSettings) (wire.SketchSettingsResult, error) {
	return call[wire.SketchSettingsResult](d.c, wire.MethodDocumentSetSketchSettings, wire.SetSketchSettingsArgs{Document: id, Settings: settings})
}

// GetEndOfPart returns the active part's end-of-part rollback marker: its feature-index position
// (-1 at the end) and whether the part is currently rolled back (#141).
//
//	eop, err := c.Documents().GetEndOfPart()
//
// mcp:tool documents_get_end_of_part
// mcp:summary Returns the active part's end-of-part marker position (-1 at the end) and whether it is rolled back (#141).
func (d Documents) GetEndOfPart() (wire.EndOfPartResult, error) {
	return call[wire.EndOfPartResult](d.c, wire.MethodDocumentGetEndOfPart, struct{}{})
}

// SetEndOfPart moves the active part's end-of-part marker to the feature index position (a negative
// index restores it to the end, re-including every feature), returning the marker's new state (#141).
//
//	c.Documents().SetEndOfPart(2) // roll back so features after index 2 are suppressed
//
// mcp:tool documents_set_end_of_part
// mcp:summary Moves the active part's end-of-part marker to a feature index (negative restores it to the end), returning its new state (#141).
func (d Documents) SetEndOfPart(position int) (wire.EndOfPartResult, error) {
	return call[wire.EndOfPartResult](d.c, wire.MethodDocumentSetEndOfPart, wire.SetEndOfPartArgs{Position: position})
}
