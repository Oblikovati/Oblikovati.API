// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Windows is the window-management operation group (M05-F10): the view frame and
// the document tab strip. View tiling inside the frame is [Client.Views]'s layout
// surface.
type Windows struct{ c *Client }

// Windows returns the window-management operation group.
func (c *Client) Windows() Windows { return Windows{c} }

// Frames returns the top-level view frames (one, on the single-frame host).
//
// mcp:tool windows_list_frames
// mcp:summary Returns the top-level view frames (one, on the single-frame host).
func (w Windows) Frames() (wire.ListViewFramesResult, error) {
	var r wire.ListViewFramesResult
	return r, w.c.call(wire.MethodWindowsListFrames, nil, &r)
}

// Tabs returns the document tab strip in order, flagging the active tab.
//
// mcp:tool windows_list_tabs
// mcp:summary Returns the document tab strip in order, flagging the active tab.
func (w Windows) Tabs() (wire.ListViewTabsResult, error) {
	var r wire.ListViewTabsResult
	return r, w.c.call(wire.MethodWindowsListTabs, nil, &r)
}

// ActivateTab brings a document tab to the front.
//
// mcp:tool windows_activate_tab
// mcp:summary Brings a document tab to the front.
func (w Windows) ActivateTab(document uint64) (wire.OKResult, error) {
	var r wire.OKResult
	return r, w.c.call(wire.MethodWindowsActivateTab, wire.ActivateViewTabArgs{Document: document}, &r)
}

// CloseTab closes a document tab; force discards unsaved changes.
//
// mcp:tool windows_close_tab
// mcp:summary Closes a document tab; force discards unsaved changes.
func (w Windows) CloseTab(document uint64, force bool) (wire.OKResult, error) {
	var r wire.OKResult
	return r, w.c.call(wire.MethodWindowsCloseTab, wire.CloseViewTabArgs{Document: document, Force: force}, &r)
}
