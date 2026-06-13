// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestDocumentExtensionsAreStable pins the per-type extensions: persisted files
// carry these on disk and the OS associates them, so a respell is a breaking
// format change (ADR-0034).
func TestDocumentExtensionsAreStable(t *testing.T) {
	for _, c := range []struct {
		kind DocumentType
		want string
	}{
		{DocumentPart, ".opd"},
		{DocumentAssembly, ".oad"},
		{DocumentDrawing, ".odd"},
		{DocumentPresentation, ".ord"},
		{DocumentUnknown, ""},
	} {
		if got := c.kind.Extension(); got != c.want {
			t.Errorf("%v.Extension() = %q, want %q", c.kind, got, c.want)
		}
	}
	if ProjectFileExtension != ".opj" {
		t.Errorf("ProjectFileExtension = %q respelled — breaking format change", ProjectFileExtension)
	}
}

// TestDocumentTypeFromExtensionRoundTrips checks the inverse mapping, including
// case-insensitivity and that the project extension is not a document kind.
func TestDocumentTypeFromExtensionRoundTrips(t *testing.T) {
	for _, kind := range []DocumentType{DocumentPart, DocumentAssembly, DocumentDrawing, DocumentPresentation} {
		if got := DocumentTypeFromExtension(kind.Extension()); got != kind {
			t.Errorf("DocumentTypeFromExtension(%q) = %v, want %v", kind.Extension(), got, kind)
		}
	}
	if got := DocumentTypeFromExtension(".OAD"); got != DocumentAssembly {
		t.Errorf("DocumentTypeFromExtension(\".OAD\") = %v, want assembly (case-insensitive)", got)
	}
	for _, ext := range []string{ProjectFileExtension, ".obk", ".stl", ".txt", ""} {
		if got := DocumentTypeFromExtension(ext); got != DocumentUnknown {
			t.Errorf("DocumentTypeFromExtension(%q) = %v, want unknown", ext, got)
		}
	}
}
