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
//
// mcp:tool dockable_windows_set
// mcp:summary Creates the window or replaces its title/content if it exists.
func (d DockableWindows) Set(w wire.DockableWindowSpec) (wire.OKResult, error) {
	var r wire.OKResult
	return r, d.c.call(wire.MethodDockableWindowsSet, wire.SetDockableWindowArgs{Window: w}, &r)
}

// SetVisible shows or hides the window without touching its content.
//
// mcp:tool dockable_windows_set_visible
// mcp:summary Shows or hides the window without touching its content.
func (d DockableWindows) SetVisible(id string, visible bool) (wire.OKResult, error) {
	var r wire.OKResult
	args := wire.SetDockableWindowVisibleArgs{ID: id, Visible: visible}
	return r, d.c.call(wire.MethodDockableWindowsSetVisible, args, &r)
}

// SetValue drives one editable control of the window to a value, exactly as a user edit would: the
// host updates the stored control and notifies the owning add-in, which may react (e.g. switch a
// view and re-render). Value is the control's string form — the option text for a dropdown/combo,
// "true"/"false" for a checkbox, the number for a value editor/slider, the text for a text box.
//
// mcp:tool set_panel_value
// mcp:summary Set one editable control of an add-in dockable window to a value (as a user edit would), notifying the add-in.
func (d DockableWindows) SetValue(windowID, controlID, value string) (wire.OKResult, error) {
	var r wire.OKResult
	args := wire.SetDockableWindowValueArgs{WindowId: windowID, ControlId: controlID, Value: value}
	return r, d.c.call(wire.MethodDockableWindowsSetValue, args, &r)
}

// SetReferences replaces a referenceList control's rows exactly as an Add-from-selection would:
// the host updates the stored rows and notifies the owning add-in with a
// wire.PanelReferencesChangedEvent. Refs is the full new set.
//
// mcp:tool set_panel_references
// mcp:summary Replace a reference-list control's rows (as Add-from-selection would), notifying the add-in.
func (d DockableWindows) SetReferences(windowID, controlID string, refs []string) (wire.OKResult, error) {
	var r wire.OKResult
	args := wire.SetDockableWindowReferencesArgs{WindowId: windowID, ControlId: controlID, Refs: refs}
	return r, d.c.call(wire.MethodDockableWindowsSetReferences, args, &r)
}

// Delete removes the window entirely.
//
// mcp:tool dockable_windows_delete
// mcp:summary Removes the window entirely.
func (d DockableWindows) Delete(id string) (wire.OKResult, error) {
	var r wire.OKResult
	return r, d.c.call(wire.MethodDockableWindowsDelete, wire.DeleteDockableWindowArgs{ID: id}, &r)
}

// List returns every add-in dockable window in creation order.
//
// mcp:tool dockable_windows_list
// mcp:summary Returns every add-in dockable window in creation order.
func (d DockableWindows) List() (wire.ListDockableWindowsResult, error) {
	var r wire.ListDockableWindowsResult
	return r, d.c.call(wire.MethodDockableWindowsList, nil, &r)
}
