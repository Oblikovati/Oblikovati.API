// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestReferenceStatusFrozenBlock pins the reference ids and wire spellings —
// the block is frozen and a renumber or respell is a breaking wire change.
func TestReferenceStatusFrozenBlock(t *testing.T) {
	want := map[ReferenceStatus]string{
		49665: "unknown", 49666: "upToDate", 49667: "outOfDate",
		49668: "missing", 49669: "replaced",
	}
	if len(want) != len(referenceStatusNames) {
		t.Fatalf("reference status count = %d, want %d", len(referenceStatusNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("ReferenceStatus(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseReferenceStatus(name); !ok || parsed != v {
			t.Errorf("ParseReferenceStatus(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
	if _, ok := ParseReferenceStatus("noSuchStatus"); ok {
		t.Error("ParseReferenceStatus must reject unknown spellings")
	}
}

// TestFileLocationTypeFrozenBlock pins the reference ids and wire spellings.
func TestFileLocationTypeFrozenBlock(t *testing.T) {
	want := map[FileLocationType]string{
		45057: "workspace", 45058: "local", 45059: "workgroup", 45060: "library",
		45061: "unknown", 45062: "ownerDirectory", 45063: "cloud",
	}
	if len(want) != len(fileLocationTypeNames) {
		t.Fatalf("location type count = %d, want %d", len(fileLocationTypeNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("FileLocationType(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseFileLocationType(name); !ok || parsed != v {
			t.Errorf("ParseFileLocationType(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
	if _, ok := ParseFileLocationType("noSuchLocation"); ok {
		t.Error("ParseFileLocationType must reject unknown spellings")
	}
}
