// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The document open/save lifecycle (#138) and the policy layer around it:
// SaveCopyAs and batch save (M03-F09, Oblikovati/Oblikovati#610). Save-time
// behavior (thumbnail capture, dependents, old-version retention) is the
// "save" application-option group ([OptionGroupSave]).

// OpenDocumentArgs is the request of [MethodDocumentsOpen]: the typed open
// options (replacing the reference API's untyped name-value bag). DeferContent
// opens a reference stub — identity registered, content not paged in.
type OpenDocumentArgs struct {
	FullDocumentName string `json:"fullDocumentName"`
	Visible          bool   `json:"visible"`
	DeferContent     bool   `json:"deferContent,omitempty"`
	// Express requests a validated saved presentation without loading assembly components.
	// The host rejects a missing or stale cache and the combination with DeferContent.
	Express bool `json:"express,omitempty"`
}

// LoadFullDocumentArgs identifies an express assembly to load with its components.
type LoadFullDocumentArgs struct {
	ID uint64 `json:"id"`
}

// SaveDocumentArgs is the request of [MethodDocumentsSave]: save the document
// at its current file binding.
type SaveDocumentArgs struct {
	Document uint64 `json:"document"`
}

// SaveDocumentAsArgs is the request of [MethodDocumentsSaveAs]: save under a
// new full document name, which becomes the document's identity.
type SaveDocumentAsArgs struct {
	Document            uint64 `json:"document"`
	NewFullDocumentName string `json:"newFullDocumentName"`
}

// SaveWarningAssemblyCacheUnavailable means the model was saved but its
// optional Express assembly presentation could not be qualified.
const SaveWarningAssemblyCacheUnavailable = "assemblyCacheUnavailable"

// SaveWarning reports a nonfatal outcome for a successfully saved document.
// Code identifies the condition and Detail explains the specific cause.
// Document identifies the session that was saved, including during close-all.
type SaveWarning struct {
	Document uint64 `json:"document"`
	Code     string `json:"code"`
	Detail   string `json:"detail"`
}

// SaveDocumentResult is the response of [MethodDocumentsSave],
// [MethodDocumentsSaveAs] and [MethodDocumentsSaveCopyAs]: where the bytes
// landed. Warnings report nonfatal outcomes after a successful save.
type SaveDocumentResult struct {
	FullDocumentName string        `json:"fullDocumentName"`
	Warnings         []SaveWarning `json:"warnings,omitempty"`
}

// SaveCopyAsArgs is the request of [MethodDocumentsSaveCopyAs]: write a copy
// to TargetFileName without retargeting the in-memory document (the export/
// archival workhorse — the document keeps its current file binding and dirty
// state). Metadata, when set, customizes the minted copy.
type SaveCopyAsArgs struct {
	Document       uint64                 `json:"document"`
	TargetFileName string                 `json:"targetFileName"`
	Metadata       *types.NewFileMetadata `json:"metadata,omitempty"`
}

// BatchSaveItem is one (document → target) pair of [MethodDocumentsBatchSave].
// TargetFileName is required for the saveAs and saveCopyAs operations and
// ignored for save.
type BatchSaveItem struct {
	Document       uint64 `json:"document"`
	TargetFileName string `json:"targetFileName,omitempty"`
}

// BatchSaveArgs is the request of [MethodDocumentsBatchSave]: execute one
// save operation — "save", "saveAs" or "saveCopyAs" — over every item,
// continuing past per-item failures.
type BatchSaveArgs struct {
	Operation string          `json:"operation"`
	Items     []BatchSaveItem `json:"items"`
}

// BatchSaveItemResult is one per-file outcome of [MethodDocumentsBatchSave].
// Warnings report nonfatal outcomes only when OK is true.
type BatchSaveItemResult struct {
	Document         uint64        `json:"document"`
	FullDocumentName string        `json:"fullDocumentName,omitempty"`
	OK               bool          `json:"ok"`
	Error            string        `json:"error,omitempty"`
	Warnings         []SaveWarning `json:"warnings,omitempty"`
}

// BatchSaveResult is the response of [MethodDocumentsBatchSave].
type BatchSaveResult struct {
	Saved   int                   `json:"saved"`
	Results []BatchSaveItemResult `json:"results"`
}

// SaveOptionsView is the "save" option group (M03-F09): application-level
// save policy. Thumbnail support is host-dependent — writing an unsupported
// capture mode errors rather than persisting a dead setting. OldVersionsToKeep
// moves the prior file into an OldVersions sibling directory on save, pruned
// to the configured count; 0 disables retention.
type SaveOptionsView struct {
	Thumbnail         types.ThumbnailSaveOption `json:"thumbnail"`
	SaveDependents    bool                      `json:"saveDependents"`
	OldVersionsToKeep int                       `json:"oldVersionsToKeep"`
}
