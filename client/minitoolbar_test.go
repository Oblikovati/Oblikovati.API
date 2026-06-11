// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestMiniToolbarSetMarshalsSpec(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	anchor := [3]float64{1, 2, 3}
	tb := wire.MiniToolbarSpec{
		ID: "sim.probe", Visible: true, HeadsUpText: "Probe", Anchor: &anchor,
		ShowOK: true, ShowCancel: true,
		Controls: []wire.MiniToolbarControlSpec{
			{Kind: types.MiniToolbarValueEditor, ID: "depth", Label: "Depth", Value: "10 mm"},
			{Kind: types.MiniToolbarSlider, ID: "blend", Number: 0.5, Min: 0, Max: 1},
		},
	}
	if _, err := c.MiniToolbars().Set(tb); err != nil {
		t.Fatalf("Set: %v", err)
	}
	if ft.gotMethod != wire.MethodMiniToolbarSet {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodMiniToolbarSet)
	}
	var sent wire.SetMiniToolbarArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Toolbar.Anchor == nil || sent.Toolbar.Anchor[2] != 3 ||
		len(sent.Toolbar.Controls) != 2 || sent.Toolbar.Controls[0].Value != "10 mm" {
		t.Errorf("sent = %+v, want the anchored two-control spec intact", sent.Toolbar)
	}
}

func TestMiniToolbarUpdateRemoveList(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	if _, err := c.MiniToolbars().Update("sim.probe", []wire.MiniToolbarControlSpec{
		{ID: "depth", Value: "12 mm"},
	}); err != nil {
		t.Fatalf("Update: %v", err)
	}
	var upd wire.UpdateMiniToolbarArgs
	if err := json.Unmarshal(ft.gotReq, &upd); err != nil || upd.ID != "sim.probe" ||
		upd.Controls[0].Value != "12 mm" {
		t.Errorf("Update sent %s, want depth=12 mm", ft.gotReq)
	}

	if _, err := c.MiniToolbars().Remove("sim.probe"); err != nil {
		t.Fatalf("Remove: %v", err)
	}
	if ft.gotMethod != wire.MethodMiniToolbarRemove {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodMiniToolbarRemove)
	}

	ft.reply = []byte(`{"toolbars":[{"id":"sim.probe","visible":true}]}`)
	lst, err := c.MiniToolbars().List()
	if err != nil || len(lst.Toolbars) != 1 || !lst.Toolbars[0].Visible {
		t.Fatalf("List = (%+v, %v), want one visible toolbar", lst, err)
	}
}
