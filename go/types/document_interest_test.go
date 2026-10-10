// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestDocumentInterestTypeFrozenBlock pins the reference ids and wire spellings.
func TestDocumentInterestTypeFrozenBlock(t *testing.T) {
	want := map[DocumentInterestType]string{
		68865: "notInterested", 68866: "interested",
	}
	if len(want) != len(documentInterestTypeNames) {
		t.Fatalf("interest type count = %d, want %d", len(documentInterestTypeNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("DocumentInterestType(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseDocumentInterestType(name); !ok || parsed != v {
			t.Errorf("ParseDocumentInterestType(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
	if _, ok := ParseDocumentInterestType("noSuchType"); ok {
		t.Error("ParseDocumentInterestType must reject unknown spellings")
	}
}
