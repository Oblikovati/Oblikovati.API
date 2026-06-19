// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestPanelControlConstructors(t *testing.T) {
	cases := []struct {
		got  wire.PanelControlSpec
		kind types.PanelControlKind
	}{
		{PanelLabel("l", "text"), types.PanelLabel},
		{PanelButton("b", "Go", "Cmd.Go"), types.PanelButton},
		{PanelSeparator(), types.PanelSeparator},
		{PanelTextBox("name", "Name", "x"), types.PanelTextBox},
		{PanelValueEditor("len", "Length", "46.7 mm"), types.PanelValueEditor},
		{PanelCheckBox("on", "Enable", true), types.PanelCheckBox},
		{PanelDropdown("t", "Type", []string{"a", "b"}, "b"), types.PanelDropdown},
		{PanelComboBox("g", "Grade", []string{"N42"}, "N42"), types.PanelComboBox},
		{PanelSlider("arc", "Arc", 0.83, 0.5, 1.0, 0.01), types.PanelSlider},
	}
	for _, c := range cases {
		if c.got.Kind != c.kind {
			t.Errorf("kind = %v, want %v", c.got.Kind, c.kind)
		}
	}
	// Spot-check the editable payloads survive the field encoding.
	if PanelCheckBox("on", "Enable", true).Value != "true" {
		t.Error("checkbox value should be \"true\"")
	}
	ve := PanelValueEditor("len", "Length", "46.7 mm")
	if ve.Value != "46.7 mm" {
		t.Errorf("value-editor value = %q", ve.Value)
	}
	dd := PanelDropdown("t", "Type", []string{"a", "b"}, "b")
	if len(dd.Options) != 2 || dd.Value != "b" {
		t.Errorf("dropdown = %+v", dd)
	}
}

func TestPanelValueChangedEventRoundTrips(t *testing.T) {
	in := wire.PanelValueChangedEvent{
		Type: wire.EventPanelValueChanged, WindowId: "motor.panel", ControlId: "poles", Value: "12",
	}
	b, _ := json.Marshal(in)
	var out wire.PanelValueChangedEvent
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out != in {
		t.Errorf("round-trip = %+v, want %+v", out, in)
	}
	if out.Type != "panel.valueChanged" {
		t.Errorf("event type = %q", out.Type)
	}
}
