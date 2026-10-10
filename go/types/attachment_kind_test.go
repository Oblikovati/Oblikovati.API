// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestAttachmentKindFrozenBlock pins the reference ids and wire spellings.
func TestAttachmentKindFrozenBlock(t *testing.T) {
	want := map[AttachmentKind]string{
		3329: "generic", 3330: "embedded", 3331: "linked",
	}
	if len(want) != len(attachmentKindNames) {
		t.Fatalf("attachment kind count = %d, want %d", len(attachmentKindNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("AttachmentKind(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseAttachmentKind(name); !ok || parsed != v {
			t.Errorf("ParseAttachmentKind(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
	if _, ok := ParseAttachmentKind("noSuchKind"); ok {
		t.Error("ParseAttachmentKind must reject unknown spellings")
	}
}
