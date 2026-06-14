// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Browser is the add-in browser-pane operation group: declare a named tree shown
// alongside the host's Model pane and observe node interaction through
// [wire.BrowserNodeEvent] push events (the ClientBrowserNodeDefinition
// equivalent, M05-F03 #256).
type Browser struct{ c *Client }

// Browser returns the add-in browser-pane operation group.
func (c *Client) Browser() Browser { return Browser{c} }

// SetPane creates the pane or replaces its whole tree — declared bulk state, like
// the client-graphics groups.
//
//	client.Browser().SetPane(wire.BrowserPaneSpec{
//	    ID: "sim", Title: "Simulation",
//	    Nodes: []wire.BrowserNodeSpec{{ID: "loads", Label: "Loads", Expanded: true}},
//	})
//
// mcp:tool browser_set_pane
// mcp:summary Creates the pane or replaces its whole tree — declared bulk state, like the client-graphics groups.
func (b Browser) SetPane(pane wire.BrowserPaneSpec) (wire.OKResult, error) {
	var r wire.OKResult
	return r, b.c.call(wire.MethodBrowserSetPane, wire.SetBrowserPaneArgs{Pane: pane}, &r)
}

// DeletePane removes an add-in pane.
//
// mcp:tool browser_delete_pane
// mcp:summary Removes an add-in pane.
func (b Browser) DeletePane(id string) (wire.OKResult, error) {
	var r wire.OKResult
	return r, b.c.call(wire.MethodBrowserDeletePane, wire.DeleteBrowserPaneArgs{ID: id}, &r)
}

// ListPanes returns every add-in pane in creation order.
//
// mcp:tool browser_list_panes
// mcp:summary Returns every add-in pane in creation order.
func (b Browser) ListPanes() (wire.ListBrowserPanesResult, error) {
	var r wire.ListBrowserPanesResult
	return r, b.c.call(wire.MethodBrowserListPanes, nil, &r)
}
