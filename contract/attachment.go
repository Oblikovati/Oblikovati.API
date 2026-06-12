// SPDX-License-Identifier: Apache-2.0

package contract

import (
	"time"

	"oblikovati.org/api/types"
)

// Linked external-file attachments (M03-F08, Oblikovati/Oblikovati#609): a
// document's named references to foreign files, persisted with the document
// and tracked with status and timestamp-based out-of-date detection.

// FileAttachment is one attachment record on a document.
type FileAttachment interface {
	// Name returns the unique name other objects bind to.
	Name() string
	// Kind returns whether the foreign file is linked, embedded or generic.
	Kind() types.AttachmentKind
	// FullFileName returns the as-recorded path, "" for embedded payloads.
	FullFileName() string
	// ResolvedFileName returns where the path resolved this session, "" while
	// missing (linked/generic only).
	ResolvedFileName() string
	// Status derives the attachment's freshness from resolution and the
	// last-known file time.
	Status() types.ReferenceStatus
	// LastKnownFileTime returns the target's modification time as recorded at
	// attach/save; the zero time when unknown or embedded.
	LastKnownFileTime() time.Time
	// BrowserVisible reports whether the attachment shows in the model browser.
	BrowserVisible() bool
	// SetBrowserVisible toggles the attachment's browser visibility.
	SetBrowserVisible(visible bool)
}

// FileAttachments is a document's attachment collection.
type FileAttachments interface {
	// Count returns the number of attachment records.
	Count() int
	// Names returns the attachment names in insertion order.
	Names() []string
	// ByName returns the named attachment, ok=false when absent.
	ByName(name string) (FileAttachment, bool)
	// Add attaches the file at fullFileName under the unique name, erroring on
	// a duplicate name or (for embedded kinds) an unreadable file.
	Add(name string, kind types.AttachmentKind, fullFileName string) (FileAttachment, error)
	// Remove deletes the named record, reporting whether it existed.
	Remove(name string) bool
}
