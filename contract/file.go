// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The file surface (M03-F07, Oblikovati/Oblikovati#608): the file as an object
// distinct from the documents it contains. One file can host several documents
// (model states and library members later; today exactly one), carries its own
// identity, and owns the persisted records of its file-to-file references.

// File is one document file: identity, load state, and its reference records.
type File interface {
	// FullFileName returns the file's absolute on-disk name.
	FullFileName() string
	// InternalName returns the stable GUID minted when the file was created;
	// it survives renames and save-as copies are re-minted.
	InternalName() string
	// RevisionID returns the GUID stamping the file's content as of its last
	// save. Any save mints a new one.
	RevisionID() string
	// DatabaseRevisionID returns the GUID stamping only the model content
	// (geometry/reference changes); a property-only save keeps it.
	DatabaseRevisionID() string
	// FileSaveCounter returns how many times the file has been saved.
	FileSaveCounter() int
	// VersionCreated returns the software version that created the file.
	VersionCreated() string
	// VersionSaved returns the software version that last saved the file.
	VersionSaved() string
	// Loaded reports whether any document within this file is loaded.
	Loaded() bool
	// Referenced reports whether any other in-memory file references this one.
	Referenced() bool
}

// FileDescriptor is the persisted "as-saved" record of one file-to-file
// reference held by a file: the logical (relative + library) name, where it
// resolved this session, and the flags explaining a broken reference.
//
// It is segregated into three embedded capability families
// ([FileReferenceIdentity], [FileReferenceStatus], [FileReferenceRepair]) so a
// read-only consumer of a reference's status does not depend on the repair
// mutation (audit I9). FileDescriptor stays their union — every existing
// implementer and caller is unaffected.
type FileDescriptor interface {
	FileReferenceIdentity
	FileReferenceStatus
	FileReferenceRepair
}

// FileReferenceIdentity is a file reference's as-saved logical and resolved names.
type FileReferenceIdentity interface {
	// FullFileName returns the reference's as-saved full file name.
	FullFileName() string
	// RelativeFileName returns the workspace-relative spelling, "" when the
	// target lies outside the project workspace.
	RelativeFileName() string
	// LibraryName returns the owning library's name, "" for non-library refs.
	LibraryName() string
	// LocationType classifies where the reference resolved.
	LocationType() types.FileLocationType
	// ResolvedFileName returns where the reference resolved this session,
	// "" while missing.
	ResolvedFileName() string
	// ReferencedFileInternalName returns the target's identity GUID as saved.
	ReferencedFileInternalName() string
	// FileSaveCounter returns the target's save counter as saved, for
	// out-of-date detection.
	FileSaveCounter() int
}

// FileReferenceStatus is a file reference's resolution state — the derived status
// and the flags that explain a broken reference.
type FileReferenceStatus interface {
	// Status derives the single status vocabulary from the flags below.
	Status() types.ReferenceStatus
	// ReferenceMissing reports that the target cannot be found anywhere.
	ReferenceMissing() bool
	// ReferenceReplaced reports that the reference was re-pointed (repaired)
	// and the owner has not been saved since.
	ReferenceReplaced() bool
	// ReferenceLocationDifferent reports the target was found somewhere other
	// than the as-saved location.
	ReferenceLocationDifferent() bool
	// ReferenceInternalNameDifferent reports the found file carries a
	// different identity GUID than the one saved against.
	ReferenceInternalNameDifferent() bool
}

// FileReferenceRepair re-points a broken reference — the mutation a read-only
// consumer of a reference's status never depends on.
type FileReferenceRepair interface {
	// ReplaceReference re-points this record at fullFileName (a repair),
	// erroring when the replacement cannot be loaded.
	ReplaceReference(fullFileName string) error
}

// ReferencedFileDescriptor is the document-side view of one file reference:
// status plus the bridge to the resolved document.
type ReferencedFileDescriptor interface {
	// DisplayName returns the reference's human-readable name.
	DisplayName() string
	// FullFileName returns the reference's as-saved full file name.
	FullFileName() string
	// Status derives the reference's resolution state.
	Status() types.ReferenceStatus
	// DocumentFound reports whether a document resolved for this reference.
	DocumentFound() bool
	// DifferentDocument reports the resolved document is not the one saved
	// against (a substitute supplied by resolution or repair).
	DifferentDocument() bool
}
