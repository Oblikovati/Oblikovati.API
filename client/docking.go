// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// DockableWindows is the add-in dockable-window operation group: declare a titled
// panel the host docks into its layout, with declarative content the host renders
// (M05-F03, #247). Visibility changes come back as
// [wire.DockableWindowChangedEvent] push events; button clicks arrive as ordinary
// command-ended events for the command each button names.
type DockableWindows struct{ c *Client }

// DockableWindows returns the add-in dockable-window operation group.
func (c *Client) DockableWindows() DockableWindows { return DockableWindows{c} }

// Set creates the window or replaces its title/content if it exists.
//
//	client.DockableWindows().Set(wire.DockableWindowSpec{
//	    ID: "sim.panel", Title: "Simulation", Dock: types.DockRight, Visible: true,
//	    Controls: []wire.PanelControlSpec{
//	        {Kind: types.PanelLabel, Text: "Mesh: 12k elements"},
//	        {Kind: types.PanelButton, Text: "Run", CommandID: "Sim.Run"},
//	    },
//	})
func (d DockableWindows) Set(w wire.DockableWindowSpec) (wire.OKResult, error) {
	var r wire.OKResult
	return r, d.c.call(wire.MethodDockableWindowsSet, wire.SetDockableWindowArgs{Window: w}, &r)
}

// SetVisible shows or hides the window without touching its content.
func (d DockableWindows) SetVisible(id string, visible bool) (wire.OKResult, error) {
	var r wire.OKResult
	args := wire.SetDockableWindowVisibleArgs{ID: id, Visible: visible}
	return r, d.c.call(wire.MethodDockableWindowsSetVisible, args, &r)
}

// Delete removes the window entirely.
func (d DockableWindows) Delete(id string) (wire.OKResult, error) {
	var r wire.OKResult
	return r, d.c.call(wire.MethodDockableWindowsDelete, wire.DeleteDockableWindowArgs{ID: id}, &r)
}

// List returns every add-in dockable window in creation order.
func (d DockableWindows) List() (wire.ListDockableWindowsResult, error) {
	var r wire.ListDockableWindowsResult
	return r, d.c.call(wire.MethodDockableWindowsList, nil, &r)
}
