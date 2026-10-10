// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// MiniToolbars is the in-canvas mini-toolbar operation group (M05-F07): declare a
// floating toolbar near the work, stream the user's edits back as
// miniToolbar.changed events and the OK/Apply/Cancel as miniToolbar.committed.
type MiniToolbars struct{ c *Client }

// MiniToolbars returns the mini-toolbar operation group.
func (c *Client) MiniToolbars() MiniToolbars { return MiniToolbars{c} }

// Set creates the toolbar or replaces it entirely.
//
//	client.MiniToolbars().Set(wire.MiniToolbarSpec{
//	    ID: "sim.probe", Visible: true, HeadsUpText: "Probe the result",
//	    ShowOK: true, ShowCancel: true,
//	    Controls: []wire.MiniToolbarControlSpec{
//	        {Kind: types.MiniToolbarValueEditor, ID: "depth", Label: "Depth", Value: "10 mm"},
//	    },
//	})
//
// mcp:tool mini_toolbar_set
// mcp:summary Creates the toolbar or replaces it entirely.
func (m MiniToolbars) Set(tb wire.MiniToolbarSpec) (wire.OKResult, error) {
	return call[wire.OKResult](m.c, wire.MethodMiniToolbarSet, wire.SetMiniToolbarArgs{Toolbar: tb})
}

// Update merges the given controls' values into the toolbar by control id.
//
// mcp:tool mini_toolbar_update
// mcp:summary Merges the given controls' values into the toolbar by control id.
func (m MiniToolbars) Update(id string, controls []wire.MiniToolbarControlSpec) (wire.OKResult, error) {
	args := wire.UpdateMiniToolbarArgs{ID: id, Controls: controls}
	return call[wire.OKResult](m.c, wire.MethodMiniToolbarUpdate, args)
}

// Remove dismisses the toolbar.
//
// mcp:tool mini_toolbar_remove
// mcp:summary Dismisses the toolbar.
func (m MiniToolbars) Remove(id string) (wire.OKResult, error) {
	return call[wire.OKResult](m.c, wire.MethodMiniToolbarRemove, wire.RemoveMiniToolbarArgs{ID: id})
}

// List returns the declared toolbars in creation order.
//
// mcp:tool mini_toolbar_list
// mcp:summary Returns the declared toolbars in creation order.
func (m MiniToolbars) List() (wire.ListMiniToolbarsResult, error) {
	return call[wire.ListMiniToolbarsResult](m.c, wire.MethodMiniToolbarList, nil)
}
