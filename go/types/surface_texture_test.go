// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestMaterialRemovalRoundTrip: every variant has a stable wire spelling that parses back, the zero
// value is MaterialRemovalAny, and an unknown spelling is rejected.
func TestMaterialRemovalRoundTrip(t *testing.T) {
	if MaterialRemoval(0) != MaterialRemovalAny {
		t.Errorf("zero MaterialRemoval = %v, want MaterialRemovalAny", MaterialRemoval(0))
	}
	want := map[MaterialRemoval]string{
		MaterialRemovalAny: "any", MaterialRemovalRequired: "required", MaterialRemovalProhibited: "prohibited",
	}
	for m, spelling := range want {
		if got := m.String(); got != spelling {
			t.Errorf("%v.String() = %q, want %q", m, got, spelling)
		}
		parsed, ok := ParseMaterialRemoval(spelling)
		if !ok || parsed != m {
			t.Errorf("ParseMaterialRemoval(%q) = (%v,%v), want (%v,true)", spelling, parsed, ok, m)
		}
	}
	if _, ok := ParseMaterialRemoval("maybe"); ok {
		t.Error("ParseMaterialRemoval should reject an unknown spelling")
	}
}
