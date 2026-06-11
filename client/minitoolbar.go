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
func (m MiniToolbars) Set(tb wire.MiniToolbarSpec) (wire.OKResult, error) {
	var r wire.OKResult
	return r, m.c.call(wire.MethodMiniToolbarSet, wire.SetMiniToolbarArgs{Toolbar: tb}, &r)
}

// Update merges the given controls' values into the toolbar by control id.
func (m MiniToolbars) Update(id string, controls []wire.MiniToolbarControlSpec) (wire.OKResult, error) {
	var r wire.OKResult
	args := wire.UpdateMiniToolbarArgs{ID: id, Controls: controls}
	return r, m.c.call(wire.MethodMiniToolbarUpdate, args, &r)
}

// Remove dismisses the toolbar.
func (m MiniToolbars) Remove(id string) (wire.OKResult, error) {
	var r wire.OKResult
	return r, m.c.call(wire.MethodMiniToolbarRemove, wire.RemoveMiniToolbarArgs{ID: id}, &r)
}

// List returns the declared toolbars in creation order.
func (m MiniToolbars) List() (wire.ListMiniToolbarsResult, error) {
	var r wire.ListMiniToolbarsResult
	return r, m.c.call(wire.MethodMiniToolbarList, nil, &r)
}
