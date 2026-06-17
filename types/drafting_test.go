// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestDraftingStandardRoundTrip pins each standard's wire spelling, the ISO default, and
// that Parse inverts String.
func TestDraftingStandardRoundTrip(t *testing.T) {
	if DraftingStandard(0) != DraftingISO {
		t.Errorf("zero DraftingStandard = %v, want DraftingISO", DraftingStandard(0))
	}
	cases := map[DraftingStandard]string{
		DraftingISO:  "iso",
		DraftingANSI: "ansi",
	}
	for d, want := range cases {
		if got := d.String(); got != want {
			t.Errorf("DraftingStandard(%d).String() = %q, want %q", d, got, want)
		}
		parsed, ok := ParseDraftingStandard(want)
		if !ok || parsed != d {
			t.Errorf("ParseDraftingStandard(%q) = (%d, %v), want (%d, true)", want, parsed, ok, d)
		}
	}
	if _, ok := ParseDraftingStandard("din"); ok {
		t.Error("ParseDraftingStandard(\"din\") = ok, want not ok")
	}
}

// TestDimensionUnitRoundTrip pins the unit spellings and the millimetre default.
func TestDimensionUnitRoundTrip(t *testing.T) {
	if DimensionUnit(0) != DimensionMillimeter {
		t.Errorf("zero DimensionUnit = %v, want DimensionMillimeter", DimensionUnit(0))
	}
	cases := map[DimensionUnit]string{
		DimensionMillimeter: "mm",
		DimensionInch:       "in",
	}
	for u, want := range cases {
		if got := u.String(); got != want {
			t.Errorf("DimensionUnit(%d).String() = %q, want %q", u, got, want)
		}
		parsed, ok := ParseDimensionUnit(want)
		if !ok || parsed != u {
			t.Errorf("ParseDimensionUnit(%q) = (%d, %v), want (%d, true)", want, parsed, ok, u)
		}
	}
}
