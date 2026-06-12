// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestDimensionDisplayTypeFrozenBlock pins the reference ids and wire
// spellings — the block is frozen and a renumber or respell is a breaking
// wire change.
func TestDimensionDisplayTypeFrozenBlock(t *testing.T) {
	want := map[DimensionDisplayType]string{
		34817: "value", 34818: "name", 34819: "expression",
		34820: "tolerance", 34821: "preciseValue",
	}
	if len(want) != len(dimensionDisplayNames) {
		t.Fatalf("dimension display type count = %d, want %d", len(dimensionDisplayNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("DimensionDisplayType(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseDimensionDisplayType(name); !ok || parsed != v {
			t.Errorf("ParseDimensionDisplayType(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
	if _, ok := ParseDimensionDisplayType("noSuchDisplay"); ok {
		t.Error("ParseDimensionDisplayType must reject unknown spellings")
	}
}
