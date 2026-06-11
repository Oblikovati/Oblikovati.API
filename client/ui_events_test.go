// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestUISearchDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"commands":[{"id":"Model.Extrude","displayName":"Extrude","enabled":true}]}`)}
	c := New(ft)
	r, err := c.UI().Search("extr")
	if err != nil || len(r.Commands) != 1 || r.Commands[0].ID != "Model.Extrude" {
		t.Fatalf("Search = (%+v, %v), want the Extrude hit", r, err)
	}
	if ft.gotMethod != wire.MethodUISearch {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodUISearch)
	}
}

func TestUIMarkingMenuRoundTrip(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	menu := wire.MarkingMenuView{
		Environment: types.SketchEnvironment,
		Quadrants: []wire.MarkingMenuItem{
			{Quadrant: types.QuadrantNorth, CommandID: "Sketch.Line"},
			{Quadrant: types.QuadrantEast, CommandID: "Sketch.Circle"},
		},
		Overflow: []string{"Sketch.Finish"},
	}
	if _, err := c.UI().SetMarkingMenu(menu); err != nil {
		t.Fatalf("SetMarkingMenu: %v", err)
	}
	var sent wire.SetMarkingMenuArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil ||
		len(sent.Menu.Quadrants) != 2 || sent.Menu.Quadrants[1].Quadrant != types.QuadrantEast {
		t.Errorf("sent = %s, want the two-slot sketch menu", ft.gotReq)
	}

	ft.reply = []byte(`{"environment":1,"quadrants":[{"quadrant":0,"commandId":"Sketch.Line"}]}`)
	got, err := c.UI().MarkingMenu(types.SketchEnvironment)
	if err != nil || got.Quadrants[0].CommandID != "Sketch.Line" {
		t.Fatalf("MarkingMenu = (%+v, %v), want the line slot", got, err)
	}
}

func TestUIContextMenuAndVisibility(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	if _, err := c.UI().SetContextMenu("com.x.sim", "feature", []wire.ContextMenuItemSpec{
		{Label: "Analyze stress", CommandID: "Sim.Analyze"},
	}); err != nil {
		t.Fatalf("SetContextMenu: %v", err)
	}
	var sent wire.SetContextMenuArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.Kind != "feature" || sent.AddIn != "com.x.sim" {
		t.Errorf("sent = %s, want the feature-kind injection", ft.gotReq)
	}

	if _, err := c.UI().SetObjectVisibility(wire.ObjectVisibilityView{WorkPlanes: false, Sketches: true}); err != nil {
		t.Fatalf("SetObjectVisibility: %v", err)
	}
	ft.reply = []byte(`{"workPlanes":false,"workAxes":true,"workPoints":true,"sketches":true}`)
	vis, err := c.UI().ObjectVisibility()
	if err != nil || vis.WorkPlanes || !vis.WorkAxes {
		t.Fatalf("ObjectVisibility = (%+v, %v), want planes hidden, axes shown", vis, err)
	}
}
