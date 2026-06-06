// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati/api/types"

// Document is the in-process contract for an open document's identity and state —
// the scalar surface every document kind shares. The GPL implementation's
// model/doc.Document satisfies it (compile-time asserted there).
//
// Navigation into a document's content (component definition, parameters,
// features, sketches) is intentionally NOT here yet: those return collections and
// sub-objects, and Go interfaces are invariant in return position, so exposing
// them as a contract requires generics or adapters — a deliberate later step.
// Out-of-process add-ins reach that structure through
// [oblikovati/api/wire] today.
type Document interface {
	// DocumentType is the kind discriminator (part/assembly/drawing/presentation).
	DocumentType() types.DocumentType

	// DisplayName is the user-facing name (derived from the file name unless an
	// explicit override is set).
	DisplayName() string
	SetDisplayName(name string)

	// FullDocumentName is the canonical full identity; FullFileName is its on-disk
	// path form.
	FullDocumentName() string
	FullFileName() string

	// Dirty reports unsaved changes; MarkDirty/ClearDirty toggle it.
	Dirty() bool
	MarkDirty()
	ClearDirty()

	// Open reports whether the document's content is paged in (false for an
	// unopened reference stub); IsReferenceStub is the inverse predicate.
	Open() bool
	IsReferenceStub() bool

	// Visible reports whether the document is shown (an open document may be
	// hidden); Referenced reports whether other open documents depend on it.
	Visible() bool
	SetVisible(visible bool)
	Referenced() bool

	// Compacted reports whether the saved package is in compacted form.
	Compacted() bool
}
