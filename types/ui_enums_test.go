// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestControlKindValuesAreStable pins the control-kind values: 0–3 mirror the
// pre-contract app-internal enum (aliased there, ADR-0018) so existing registrations
// keep their meaning; 4 is the popup control added by M05-F03.
func TestControlKindValuesAreStable(t *testing.T) {
	want := map[ControlKind]string{
		ButtonControl: "button", ToggleControl: "toggle", ComboControl: "combo",
		SpinnerControl: "spinner", PopupControl: "popup",
	}
	if ButtonControl != 0 || ToggleControl != 1 || ComboControl != 2 ||
		SpinnerControl != 3 || PopupControl != 4 {
		t.Errorf("control-kind values changed: %d %d %d %d %d, want 0..4",
			ButtonControl, ToggleControl, ComboControl, SpinnerControl, PopupControl)
	}
	for k, name := range want {
		if k.String() != name {
			t.Errorf("%d.String() = %q, want %q", k, k.String(), name)
		}
	}
}

// TestDockingStateNamesAreStable pins the docking-state names and values.
func TestDockingStateNamesAreStable(t *testing.T) {
	want := map[DockingState]string{
		DockFloating: "floating", DockLeft: "left", DockRight: "right", DockBottom: "bottom",
	}
	if DockFloating != 0 || DockLeft != 1 || DockRight != 2 || DockBottom != 3 {
		t.Errorf("docking values = %d,%d,%d,%d, want 0..3",
			DockFloating, DockLeft, DockRight, DockBottom)
	}
	for d, name := range want {
		if d.String() != name {
			t.Errorf("%d.String() = %q, want %q", d, d.String(), name)
		}
	}
}

// TestPanelControlKindNamesAreStable pins the panel-control names and values.
func TestPanelControlKindNamesAreStable(t *testing.T) {
	want := map[PanelControlKind]string{
		PanelLabel: "label", PanelButton: "button", PanelSeparator: "separator",
	}
	if PanelLabel != 0 || PanelButton != 1 || PanelSeparator != 2 {
		t.Errorf("panel-control values = %d,%d,%d, want 0..2",
			PanelLabel, PanelButton, PanelSeparator)
	}
	for k, name := range want {
		if k.String() != name {
			t.Errorf("%d.String() = %q, want %q", k, k.String(), name)
		}
	}
}

func TestPanelTreeTableKinds(t *testing.T) {
	if PanelTree != 13 || PanelTable != 14 {
		t.Fatalf("ordinals: PanelTree=%d PanelTable=%d, want 13,14 (appended, stable)", PanelTree, PanelTable)
	}
	if PanelTree.String() != "tree" || PanelTable.String() != "table" {
		t.Fatalf("names: %q, %q, want \"tree\", \"table\"", PanelTree.String(), PanelTable.String())
	}
}
