// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Linked external-file attachments (M03-F08, Oblikovati/Oblikovati#609): a
// document's named references to foreign files — linked spreadsheets, sketch
// images, design-data side files — tracked in the reference graph with status
// and timestamp-based out-of-date detection.

// AttachmentInfo is one attachment record: the unique name other objects bind
// to, the kind, the linked path (resolved through the project search paths)
// or the embedded resource id, and the freshness state. LastKnownFileTime is
// RFC 3339; for linked attachments a newer on-disk time reports outOfDate.
type AttachmentInfo struct {
	Name              string                `json:"name"`
	Kind              types.AttachmentKind  `json:"kind"`
	FullFileName      string                `json:"fullFileName,omitempty"`
	ResolvedFileName  string                `json:"resolvedFileName,omitempty"`
	Resource          string                `json:"resource,omitempty"`
	Status            types.ReferenceStatus `json:"status"`
	LastKnownFileTime string                `json:"lastKnownFileTime,omitempty"`
	BrowserVisible    bool                  `json:"browserVisible"`
}

// ListAttachmentsArgs is the request of [MethodDocumentsListAttachments].
type ListAttachmentsArgs struct {
	Document uint64 `json:"document"`
}

// ListAttachmentsResult is the response of [MethodDocumentsListAttachments].
type ListAttachmentsResult struct {
	Attachments []AttachmentInfo `json:"attachments"`
}

// AddAttachmentArgs is the request of [MethodDocumentsAddAttachment]: attach
// the file at FullFileName under the unique Name. An embedded attachment reads
// the file once and carries its bytes; linked/generic record the path.
type AddAttachmentArgs struct {
	Document     uint64               `json:"document"`
	Name         string               `json:"name"`
	Kind         types.AttachmentKind `json:"kind"`
	FullFileName string               `json:"fullFileName"`
}

// RemoveAttachmentArgs is the request of [MethodDocumentsRemoveAttachment]:
// delete the named attachment record (and an embedded payload with it).
type RemoveAttachmentArgs struct {
	Document uint64 `json:"document"`
	Name     string `json:"name"`
}
