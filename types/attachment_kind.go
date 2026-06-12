// SPDX-License-Identifier: Apache-2.0

package types

// AttachmentKind classifies a document's reference to a foreign (non-native)
// file: linked in place, embedded as a payload carried inside the document, or
// a generic tracked path (M03-F08, Oblikovati/Oblikovati#609).
//
// The values are a frozen block matching the reference API's foreign-document
// type enum; never renumber them.
type AttachmentKind int32

const (
	// AttachmentGeneric is an opaque tracked path: the document records the
	// file but assigns it no further semantics.
	AttachmentGeneric AttachmentKind = 3329
	// AttachmentEmbedded carries the foreign file's bytes inside the document
	// (as an embedded resource), so it travels with the .obk.
	AttachmentEmbedded AttachmentKind = 3330
	// AttachmentLinked references the foreign file in place; the document
	// tracks its path and last-known modification time.
	AttachmentLinked AttachmentKind = 3331
)

// attachmentKindNames are the frozen wire spellings.
var attachmentKindNames = map[AttachmentKind]string{
	AttachmentGeneric:  "generic",
	AttachmentEmbedded: "embedded",
	AttachmentLinked:   "linked",
}

// String returns the attachment kind's wire spelling.
func (k AttachmentKind) String() string { return enumName(attachmentKindNames, k) }

// ParseAttachmentKind resolves a wire spelling back to its AttachmentKind.
func ParseAttachmentKind(s string) (AttachmentKind, bool) {
	return enumFromName(attachmentKindNames, s)
}
