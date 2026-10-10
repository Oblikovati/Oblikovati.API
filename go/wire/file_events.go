// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// File-access and file-UI push events (M04-F05, Oblikovati/Oblikovati#613):
// the host announces reference resolution, dirty transitions, and the file
// new/open/save-as flows. In-process subscribers may answer the Before phase of
// these hooks (supply a resolved path, pre-seed a dialog); the push events an
// add-in receives over Notify are the After-phase observations of the outcome
// (out-of-proc answering is the event-transport work, Oblikovati#148).

// FileResolutionEventPayload is the JSON shape of [EventFileResolution]: a
// referenced document name failed to resolve. ResolvedName is the substitute a
// subscriber supplied ("" when nothing resolved it and the open failed).
type FileResolutionEventPayload struct {
	Type          string `json:"type"` // always EventFileResolution
	RequestedName string `json:"requestedName"`
	ResolvedName  string `json:"resolvedName,omitempty"`
}

// FileDirtyEventPayload is the JSON shape of [EventFileDirty]: a document
// gained unsaved changes (the clean→dirty transition; further edits while
// already dirty do not re-fire).
type FileDirtyEventPayload struct {
	Type             string `json:"type"` // always EventFileDirty
	Document         uint64 `json:"document"`
	FullDocumentName string `json:"fullDocumentName"`
}

// FileMetadataEntry is one name/value pair collected by the
// [EventFilePopulateMetadata] hook (file properties gathered around a save).
type FileMetadataEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// FileDialogHookPayload is the JSON shape of the file-UI hook events
// ([EventFileNew], [EventFileNewDialog], [EventFileOpenDialog],
// [EventFileSaveAsDialog], [EventFileOpenFromMRU],
// [EventFilePopulateMetadata]). Only the fields relevant to the hook are set:
// DocumentType for a new document, TemplateFile/FileName for the dialog flows
// (the path a subscriber pre-seeded, or the one the user chose), SaveCopyAs for
// the save-as variant, Metadata for the populate-metadata collection.
type FileDialogHookPayload struct {
	Type         string              `json:"type"`
	DocumentType types.DocumentType  `json:"documentType,omitempty"`
	TemplateFile string              `json:"templateFile,omitempty"`
	FileName     string              `json:"fileName,omitempty"`
	SaveCopyAs   bool                `json:"saveCopyAs,omitempty"`
	Metadata     []FileMetadataEntry `json:"metadata,omitempty"`
}
