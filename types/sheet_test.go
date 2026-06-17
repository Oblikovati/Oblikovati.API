// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestSheetSizeRoundTrip pins each sheet size's wire spelling, that the zero value is
// SheetSizeCustom, and that Parse inverts String.
func TestSheetSizeRoundTrip(t *testing.T) {
	if SheetSize(0) != SheetSizeCustom {
		t.Errorf("zero SheetSize = %v, want SheetSizeCustom", SheetSize(0))
	}
	cases := map[SheetSize]string{
		SheetSizeCustom: "custom",
		SheetSizeA0:     "a0",
		SheetSizeA1:     "a1",
		SheetSizeA2:     "a2",
		SheetSizeA3:     "a3",
		SheetSizeA4:     "a4",
		SheetSizeAnsiA:  "ansiA",
		SheetSizeAnsiB:  "ansiB",
		SheetSizeAnsiC:  "ansiC",
		SheetSizeAnsiD:  "ansiD",
		SheetSizeAnsiE:  "ansiE",
	}
	for s, want := range cases {
		if got := s.String(); got != want {
			t.Errorf("SheetSize(%d).String() = %q, want %q", s, got, want)
		}
		parsed, ok := ParseSheetSize(want)
		if !ok || parsed != s {
			t.Errorf("ParseSheetSize(%q) = (%d, %v), want (%d, true)", want, parsed, ok, s)
		}
	}
	if _, ok := ParseSheetSize("a99"); ok {
		t.Error("ParseSheetSize(\"a99\") = ok, want not ok")
	}
}

// TestSheetDimensionsMM pins the standard dimensions (portrait, width ≤ height) and that
// a custom size is not table-driven. ANSI sizes are their inch dimensions at 25.4 mm/in.
func TestSheetDimensionsMM(t *testing.T) {
	cases := map[SheetSize][2]float64{
		SheetSizeA0:    {841, 1189},
		SheetSizeA4:    {210, 297},
		SheetSizeAnsiA: {215.9, 279.4},
		SheetSizeAnsiE: {863.6, 1117.6},
	}
	for s, want := range cases {
		w, h, ok := SheetDimensionsMM(s)
		if !ok || w != want[0] || h != want[1] {
			t.Errorf("SheetDimensionsMM(%v) = (%g, %g, %v), want (%g, %g, true)", s, w, h, ok, want[0], want[1])
		}
		if w > h {
			t.Errorf("SheetDimensionsMM(%v) = %g×%g, want portrait (width ≤ height)", s, w, h)
		}
	}
	if w, h, ok := SheetDimensionsMM(SheetSizeCustom); ok || w != 0 || h != 0 {
		t.Errorf("SheetDimensionsMM(custom) = (%g, %g, %v), want (0, 0, false)", w, h, ok)
	}
}

// TestSheetOrientationRoundTrip pins the orientation spellings and the portrait default.
func TestSheetOrientationRoundTrip(t *testing.T) {
	if SheetOrientation(0) != SheetPortrait {
		t.Errorf("zero SheetOrientation = %v, want SheetPortrait", SheetOrientation(0))
	}
	cases := map[SheetOrientation]string{
		SheetPortrait:  "portrait",
		SheetLandscape: "landscape",
	}
	for o, want := range cases {
		if got := o.String(); got != want {
			t.Errorf("SheetOrientation(%d).String() = %q, want %q", o, got, want)
		}
		parsed, ok := ParseSheetOrientation(want)
		if !ok || parsed != o {
			t.Errorf("ParseSheetOrientation(%q) = (%d, %v), want (%d, true)", want, parsed, ok, o)
		}
	}
}
