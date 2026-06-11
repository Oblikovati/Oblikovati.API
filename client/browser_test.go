// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestBrowserSetPaneMarshalsTree(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	pane := wire.BrowserPaneSpec{
		ID: "sim", Title: "Simulation",
		Nodes: []wire.BrowserNodeSpec{{
			ID: "loads", Label: "Loads", Expanded: true,
			Children: []wire.BrowserNodeSpec{{ID: "f1", Label: "Force 10N"}},
		}},
	}
	if _, err := c.Browser().SetPane(pane); err != nil {
		t.Fatalf("SetPane: %v", err)
	}
	if ft.gotMethod != wire.MethodBrowserSetPane {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodBrowserSetPane)
	}
	var sent wire.SetBrowserPaneArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Pane.ID != "sim" || len(sent.Pane.Nodes) != 1 ||
		len(sent.Pane.Nodes[0].Children) != 1 ||
		sent.Pane.Nodes[0].Children[0].Label != "Force 10N" {
		t.Errorf("sent = %+v, want the nested tree intact", sent.Pane)
	}
}

func TestBrowserDeleteAndListPanes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	if _, err := c.Browser().DeletePane("sim"); err != nil {
		t.Fatalf("DeletePane: %v", err)
	}
	if ft.gotMethod != wire.MethodBrowserDeletePane {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodBrowserDeletePane)
	}

	ft.reply = []byte(`{"panes":[{"id":"sim","title":"Simulation"}]}`)
	lst, err := c.Browser().ListPanes()
	if err != nil || len(lst.Panes) != 1 || lst.Panes[0].Title != "Simulation" {
		t.Fatalf("ListPanes = (%+v, %v), want one Simulation pane", lst, err)
	}
}

func TestDockableWindowsSetMarshalsControls(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	w := wire.DockableWindowSpec{
		ID: "sim.panel", Title: "Simulation", Dock: types.DockRight, Visible: true,
		Controls: []wire.PanelControlSpec{
			{Kind: types.PanelLabel, Text: "Mesh: 12k"},
			{Kind: types.PanelButton, Text: "Run", CommandID: "Sim.Run"},
		},
	}
	if _, err := c.DockableWindows().Set(w); err != nil {
		t.Fatalf("Set: %v", err)
	}
	if ft.gotMethod != wire.MethodDockableWindowsSet {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDockableWindowsSet)
	}
	var sent wire.SetDockableWindowArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Window.Dock != types.DockRight || len(sent.Window.Controls) != 2 ||
		sent.Window.Controls[1].CommandID != "Sim.Run" {
		t.Errorf("sent = %+v, want dock-right window with a Sim.Run button", sent.Window)
	}
}

func TestDockableWindowsVisibilityDeleteList(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	if _, err := c.DockableWindows().SetVisible("sim.panel", false); err != nil {
		t.Fatalf("SetVisible: %v", err)
	}
	var vis wire.SetDockableWindowVisibleArgs
	if err := json.Unmarshal(ft.gotReq, &vis); err != nil || vis.ID != "sim.panel" || vis.Visible {
		t.Errorf("SetVisible sent %s, want id=sim.panel visible=false", ft.gotReq)
	}

	if _, err := c.DockableWindows().Delete("sim.panel"); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	if ft.gotMethod != wire.MethodDockableWindowsDelete {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDockableWindowsDelete)
	}

	ft.reply = []byte(`{"windows":[{"id":"sim.panel","title":"Simulation","visible":true}]}`)
	lst, err := c.DockableWindows().List()
	if err != nil || len(lst.Windows) != 1 || !lst.Windows[0].Visible {
		t.Fatalf("List = (%+v, %v), want one visible window", lst, err)
	}
}

func TestRibbonEnvironmentsDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(
		`{"environments":[{"environment":0,"name":"Base","active":true},{"environment":1,"name":"Sketch","active":false}]}`)}
	c := New(ft)
	r, err := c.Ribbon().Environments()
	if err != nil {
		t.Fatalf("Environments: %v", err)
	}
	if ft.gotMethod != wire.MethodUIListEnvironments {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodUIListEnvironments)
	}
	if len(r.Environments) != 2 || !r.Environments[0].Active ||
		r.Environments[1].Environment != types.SketchEnvironment {
		t.Errorf("decoded = %+v, want active Base + Sketch", r.Environments)
	}
}
