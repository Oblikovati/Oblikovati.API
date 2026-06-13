// SPDX-License-Identifier: Apache-2.0

package types

import "strings"

// DocumentType discriminates the four document kinds plus an unknown sentinel.
//
// Values are STABLE ACROSS SESSIONS and must never be renumbered: they are
// persisted in package manifests and the document reference graph (architecture
// core/05). This is the canonical, Apache-2.0 definition; the GPL implementation
// aliases it (model/doc.DocumentType) so existing call sites are unaffected.
type DocumentType uint32

const (
	// DocumentUnknown is the zero value: a document whose kind has not been
	// resolved (e.g. an unresolved reference stub before its manifest is read).
	DocumentUnknown DocumentType = 0
	// DocumentPart holds a single modeled part.
	DocumentPart DocumentType = 1
	// DocumentAssembly holds component occurrences and constraints.
	DocumentAssembly DocumentType = 2
	// DocumentDrawing holds annotated sheets/views of other documents.
	DocumentDrawing DocumentType = 3
	// DocumentPresentation holds exploded/animated views of an assembly.
	DocumentPresentation DocumentType = 4
)

// IsValid reports whether t is one of the four real document kinds (not unknown).
func (t DocumentType) IsValid() bool {
	return t >= DocumentPart && t <= DocumentPresentation
}

// String returns a stable lowercase name for the kind, used in diagnostics and the
// package manifest. The values, not these names, are the persisted identity.
func (t DocumentType) String() string {
	switch t {
	case DocumentPart:
		return "part"
	case DocumentAssembly:
		return "assembly"
	case DocumentDrawing:
		return "drawing"
	case DocumentPresentation:
		return "presentation"
	default:
		return "unknown"
	}
}

// Per-document-type on-disk extensions (ADR-0034). Each document kind carries its
// own user-facing extension so the OS, file dialogs, and the reference graph can
// identify a file's kind *before* its manifest is read, and so assemblies and
// drawings name their referenced files unambiguously. This supersedes the single
// ".obk" package extension (ADR-0020 amended): the manifest's documentType stays
// the canonical identity — the extension mirrors it and must agree.
const (
	// PartFileExtension is the extension for a part document (DocumentPart).
	PartFileExtension = ".opd"
	// AssemblyFileExtension is the extension for an assembly document (DocumentAssembly).
	AssemblyFileExtension = ".oad"
	// DrawingFileExtension is the extension for a drawing document (DocumentDrawing).
	DrawingFileExtension = ".odd"
	// PresentationFileExtension is the extension for a presentation document (DocumentPresentation).
	PresentationFileExtension = ".ord"
	// ProjectFileExtension is the extension for a design-project file. A project is
	// NOT a document (it has no DocumentType) — it is the portable search-path
	// config that resolves a document's referenced files (architecture core/05).
	ProjectFileExtension = ".opj"
)

// Extension returns the on-disk file extension (leading dot) for the document
// kind, or "" for DocumentUnknown.
//
//	types.DocumentPart.Extension() // ".opd"
func (t DocumentType) Extension() string {
	switch t {
	case DocumentPart:
		return PartFileExtension
	case DocumentAssembly:
		return AssemblyFileExtension
	case DocumentDrawing:
		return DrawingFileExtension
	case DocumentPresentation:
		return PresentationFileExtension
	default:
		return ""
	}
}

// DocumentTypeFromExtension maps a file extension (leading dot, any case) to its
// document kind, returning DocumentUnknown for the project extension or anything
// unrecognized. It is the inverse of [DocumentType.Extension].
//
//	types.DocumentTypeFromExtension(".OAD") // DocumentAssembly
func DocumentTypeFromExtension(ext string) DocumentType {
	switch strings.ToLower(ext) {
	case PartFileExtension:
		return DocumentPart
	case AssemblyFileExtension:
		return DocumentAssembly
	case DrawingFileExtension:
		return DocumentDrawing
	case PresentationFileExtension:
		return DocumentPresentation
	default:
		return DocumentUnknown
	}
}
