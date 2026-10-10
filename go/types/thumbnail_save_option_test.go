// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestThumbnailSaveOptionFrozenBlock pins the reference ids and wire spellings.
func TestThumbnailSaveOptionFrozenBlock(t *testing.T) {
	want := map[ThumbnailSaveOption]string{
		79873: "none", 79874: "isoViewOnSave", 79875: "activeWindowOnSave",
		79876: "activeWindow", 79877: "importFromFile",
	}
	if len(want) != len(thumbnailSaveOptionNames) {
		t.Fatalf("thumbnail option count = %d, want %d", len(thumbnailSaveOptionNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("ThumbnailSaveOption(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseThumbnailSaveOption(name); !ok || parsed != v {
			t.Errorf("ParseThumbnailSaveOption(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
	if _, ok := ParseThumbnailSaveOption("noSuchOption"); ok {
		t.Error("ParseThumbnailSaveOption must reject unknown spellings")
	}
}
