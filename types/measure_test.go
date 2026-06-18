// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestMeasureTypeRoundTrip checks every measure type spells to its wire name and parses back.
func TestMeasureTypeRoundTrip(t *testing.T) {
	cases := map[MeasureType]string{
		MeasureLength:      "length",
		MeasureArea:        "area",
		MeasureDistance:    "distance",
		MeasureMinDistance: "minDistance",
	}
	for mt, name := range cases {
		if got := mt.String(); got != name {
			t.Errorf("%d.String() = %q, want %q", mt, got, name)
		}
		got, ok := ParseMeasureType(name)
		if !ok || got != mt {
			t.Errorf("ParseMeasureType(%q) = (%d, %v), want (%d, true)", name, got, ok, mt)
		}
	}
	if _, ok := ParseMeasureType("volume"); ok {
		t.Error(`ParseMeasureType("volume") = ok, want unknown`)
	}
}
