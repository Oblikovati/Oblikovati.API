// SPDX-License-Identifier: Apache-2.0

package types

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
