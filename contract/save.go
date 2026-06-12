// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The save policy layer (M03-F09, Oblikovati/Oblikovati#610): application-wide
// save options and the batch-save service that executes one operation over
// several documents with per-file outcomes.

// SaveOptions is the application-level save policy. Setters error when the
// host does not implement the requested behavior (no dead settings).
type SaveOptions interface {
	// Thumbnail returns how a preview thumbnail is captured on save.
	Thumbnail() types.ThumbnailSaveOption
	// SetThumbnail selects the capture mode, erroring on modes this host
	// cannot perform.
	SetThumbnail(option types.ThumbnailSaveOption) error
	// SaveDependents reports whether saving a document also saves its dirty
	// referenced documents first.
	SaveDependents() bool
	// SetSaveDependents toggles dependent saving.
	SetSaveDependents(save bool)
	// OldVersionsToKeep returns how many prior versions a save retains in the
	// OldVersions sibling directory; 0 disables retention.
	OldVersionsToKeep() int
	// SetOldVersionsToKeep sets the retention count, erroring on negatives.
	SetOldVersionsToKeep(count int) error
}

// BatchSaveOutcome is one per-file result of a batch-save execution.
type BatchSaveOutcome struct {
	// FullDocumentName is where the file landed (or would have).
	FullDocumentName string
	// Err is the per-file failure, nil on success; the batch continues past
	// individual failures.
	Err error
}

// BatchSave queues (document → target) pairs and executes one save operation
// over all of them. A queue is single-use: executing drains it.
type BatchSave interface {
	// AddFileToSave queues a document with its target file name (ignored for
	// ExecuteSave); it errors on a nil document or a duplicate target.
	AddFileToSave(document Document, targetFileName string) error
	// Count returns the number of queued pairs.
	Count() int
	// ExecuteSave saves every queued document at its current binding.
	ExecuteSave() []BatchSaveOutcome
	// ExecuteSaveAs saves every queued document under its target name,
	// retargeting each document's identity.
	ExecuteSaveAs() []BatchSaveOutcome
	// ExecuteSaveCopyAs writes a copy of every queued document to its target
	// without retargeting the in-memory documents.
	ExecuteSaveCopyAs() []BatchSaveOutcome
}
