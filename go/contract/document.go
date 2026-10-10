// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// Document is the in-process contract for an open document's identity and state —
// the scalar surface every document kind shares. The GPL implementation's
// model/doc.Document satisfies it (compile-time asserted there).
//
// Navigation into a document's content (component definition, parameters,
// features, sketches) is intentionally NOT here yet: those return collections and
// sub-objects, and Go interfaces are invariant in return position, so exposing
// them as a contract requires generics or adapters — a deliberate later step.
// Out-of-process add-ins reach that structure through
// [oblikovati.org/api/wire] today.
//
// The surface is segregated into three embedded capability families
// ([DocumentIdentity], [DocumentDirtyState], [DocumentLifecycle]) so a consumer that
// only reads identity does not depend on the dirty/lifecycle verbs (audit I9). Document
// stays their union — every existing implementer and caller is unaffected.
type Document interface {
	DocumentIdentity
	DocumentDirtyState
	DocumentLifecycle
}

// DocumentIdentity is a document's naming and kind — the read-mostly identity surface.
type DocumentIdentity interface {
	// DocumentType is the kind discriminator (part/assembly/drawing/presentation).
	DocumentType() types.DocumentType

	// SubType refines the base type with a flavored sub-type id (M03-F11) —
	// types.SubTypePlain for an unflavored document.
	SubType() types.DocumentSubTypeID

	// DisplayName is the user-facing name (derived from the file name unless an
	// explicit override is set).
	DisplayName() string
	SetDisplayName(name string)

	// FullDocumentName is the canonical full identity; FullFileName is its on-disk
	// path form.
	FullDocumentName() string
	FullFileName() string
}

// DocumentDirtyState is a document's unsaved-changes flag.
type DocumentDirtyState interface {
	// Dirty reports unsaved changes; MarkDirty/ClearDirty toggle it.
	Dirty() bool
	MarkDirty()
	ClearDirty()
}

// DocumentLifecycle is a document's load, visibility and packaging state.
type DocumentLifecycle interface {
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
