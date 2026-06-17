// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestUnitsTypeFrozenBlock pins the category ids and wire spellings — a
// renumber or respell is a breaking wire change.
func TestUnitsTypeFrozenBlock(t *testing.T) {
	want := map[UnitsType]string{
		0: "unitless", 1: "length", 2: "angle", 3: "area",
		4: "volume", 5: "mass", 6: "time",
	}
	if len(want) != len(unitsTypeNames) {
		t.Fatalf("units type count = %d, want %d", len(unitsTypeNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("UnitsType(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseUnitsType(name); !ok || parsed != v {
			t.Errorf("ParseUnitsType(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
	if _, ok := ParseUnitsType("noSuchCategory"); ok {
		t.Error("ParseUnitsType must reject unknown spellings")
	}
}
