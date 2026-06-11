// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestAddInLoadBehaviorNamesAreStable pins the wire-visible names and numeric values
// (persisted in the per-user add-in preferences file and in saved automations).
func TestAddInLoadBehaviorNamesAreStable(t *testing.T) {
	want := map[AddInLoadBehavior]string{
		LoadOnStartup: "startup", LoadOnDemand: "demand", LoadDisabled: "disabled",
	}
	if LoadOnStartup != 0 || LoadOnDemand != 1 || LoadDisabled != 2 {
		t.Errorf("load-behavior values = %d,%d,%d, want 0,1,2",
			LoadOnStartup, LoadOnDemand, LoadDisabled)
	}
	for b, name := range want {
		if b.String() != name {
			t.Errorf("%d.String() = %q, want %q", b, b.String(), name)
		}
		got, ok := ParseAddInLoadBehavior(name)
		if !ok || got != b {
			t.Errorf("ParseAddInLoadBehavior(%q) = (%v, %v), want (%v, true)", name, got, ok, b)
		}
	}
}

// TestParseAddInLoadBehaviorRejectsUnknown checks an unknown name reports !ok so
// callers keep their current value instead of silently defaulting.
func TestParseAddInLoadBehaviorRejectsUnknown(t *testing.T) {
	if _, ok := ParseAddInLoadBehavior("sometimes"); ok {
		t.Error(`ParseAddInLoadBehavior("sometimes") ok = true, want false`)
	}
}

// TestAddInKindNamesAreStable pins the kind names and values.
func TestAddInKindNamesAreStable(t *testing.T) {
	if StandardAddIn != 0 || TranslatorAddIn != 1 {
		t.Errorf("kind values = %d,%d, want 0,1", StandardAddIn, TranslatorAddIn)
	}
	if StandardAddIn.String() != "standard" || TranslatorAddIn.String() != "translator" {
		t.Errorf("kind names = %q,%q, want standard,translator",
			StandardAddIn.String(), TranslatorAddIn.String())
	}
}
