// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestDocumentSubTypeBuiltIns pins the reserved sub-type ids: persisted .obk
// files carry these strings, so a respell is a breaking format change.
func TestDocumentSubTypeBuiltIns(t *testing.T) {
	if SubTypePlain != "" {
		t.Errorf("SubTypePlain = %q, want the empty id", SubTypePlain)
	}
	if SubTypeSheetMetalPart != "org.oblikovati.part.sheetMetal" {
		t.Errorf("SubTypeSheetMetalPart = %q respelled — breaking format change", SubTypeSheetMetalPart)
	}
	for id, want := range map[DocumentSubTypeID]bool{
		SubTypePlain:                    true,
		SubTypeSheetMetalPart:           true,
		"org.oblikovati.future.flavor":  true,
		"com.x.sim.study":               false,
		"org.oblikovatiX.part.flavored": false,
	} {
		if got := id.BuiltIn(); got != want {
			t.Errorf("DocumentSubTypeID(%q).BuiltIn() = %v, want %v", id, got, want)
		}
	}
}
