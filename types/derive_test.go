// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestDeriveStyleSpellings pins the derive style wire spellings and round-trip.
func TestDeriveStyleSpellings(t *testing.T) {
	want := map[DeriveStyle]string{
		DeriveInclude: "include", DeriveExclude: "exclude", DeriveSubtract: "subtract",
	}
	if len(want) != len(deriveStyleNames) {
		t.Fatalf("derive style count = %d, want %d", len(deriveStyleNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("DeriveStyle(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseDeriveStyle(name); !ok || parsed != v {
			t.Errorf("ParseDeriveStyle(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
	if _, ok := ParseDeriveStyle("nope"); ok {
		t.Error("ParseDeriveStyle must reject unknown spellings")
	}
}

// TestShrinkwrapStyleSpellings pins the shrinkwrap removal/envelope wire spellings.
func TestShrinkwrapStyleSpellings(t *testing.T) {
	removes := map[ShrinkwrapRemoveStyle]string{
		RemoveNone: "none", RemoveSmallParts: "smallParts", RemoveInternalParts: "internalParts",
	}
	for v, name := range removes {
		if v.String() != name {
			t.Errorf("ShrinkwrapRemoveStyle(%d).String() = %q, want %q", int32(v), v.String(), name)
		}
		if parsed, ok := ParseShrinkwrapRemoveStyle(name); !ok || parsed != v {
			t.Errorf("ParseShrinkwrapRemoveStyle(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
	envelopes := map[ShrinkwrapEnvelopeStyle]string{
		EnvelopeNone: "none", EnvelopePerPart: "perPart", EnvelopeWhole: "whole",
	}
	for v, name := range envelopes {
		if v.String() != name {
			t.Errorf("ShrinkwrapEnvelopeStyle(%d).String() = %q, want %q", int32(v), v.String(), name)
		}
		if parsed, ok := ParseShrinkwrapEnvelopeStyle(name); !ok || parsed != v {
			t.Errorf("ParseShrinkwrapEnvelopeStyle(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
}
