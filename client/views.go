// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Views is the per-document view-collection operation group: a document owns a set of
// views (the document's view collection), each with its own camera, tiled per a layout. Use it
// to enumerate, add, activate, close, and rename views, and to choose the tiling layout.
// A Document field of 0 in any request targets the active document.
type Views struct{ c *Client }

// Views returns the view-collection operation group.
func (c *Client) Views() Views { return Views{c} }

// List enumerates a document's views (document 0 = active), with the active index and layout.
//
//	vs, _ := client.Views().List(0)
//
// mcp:tool list_views
// mcp:summary List a document's views (each with its camera), the active index, and the tiling layout; document 0 = active.
func (v Views) List(document uint64) (wire.ListViewsResult, error) {
	var r wire.ListViewsResult
	return r, v.c.call(wire.MethodViewsList, wire.ListViewsArgs{Document: document}, &r)
}

// Add creates a new view of a document and makes it active, returning the new view.
//
// mcp:tool add_view
// mcp:summary Add a new view to a document and make it active (copyActiveCamera starts it at the current view's camera); document 0 = active.
func (v Views) Add(args wire.AddViewArgs) (wire.ViewInfo, error) {
	var r wire.ViewInfo
	return r, v.c.call(wire.MethodViewsAdd, args, &r)
}

// Activate makes the indexed view of a document the active view.
//
// mcp:tool activate_view
// mcp:summary Make the view at the given index the active view; document 0 = active.
func (v Views) Activate(args wire.ActivateViewArgs) (wire.ListViewsResult, error) {
	var r wire.ListViewsResult
	return r, v.c.call(wire.MethodViewsActivate, args, &r)
}

// Close removes the indexed view; closing the last view of a document is refused.
//
// mcp:tool close_view
// mcp:summary Close the view at the given index (the last view cannot be closed); document 0 = active.
func (v Views) Close(args wire.CloseViewArgs) (wire.ListViewsResult, error) {
	var r wire.ListViewsResult
	return r, v.c.call(wire.MethodViewsClose, args, &r)
}

// Rename sets the indexed view's name.
//
// mcp:tool views_rename
// mcp:summary Sets the indexed view's name.
func (v Views) Rename(args wire.RenameViewArgs) (wire.ViewInfo, error) {
	var r wire.ViewInfo
	return r, v.c.call(wire.MethodViewsRename, args, &r)
}

// Layout returns a document's current tiling layout.
//
// mcp:tool views_get_layout
// mcp:summary Returns a document's current tiling layout.
func (v Views) Layout(document uint64) (wire.LayoutResult, error) {
	var r wire.LayoutResult
	return r, v.c.call(wire.MethodViewsGetLayout, wire.ListViewsArgs{Document: document}, &r)
}

// SetLayout chooses how a document's views are tiled.
//
// mcp:tool set_view_layout
// mcp:summary Set how a document's views tile the viewport (0 single, 1 two-H, 2 two-V, 3 three, 4 quad); document 0 = active.
func (v Views) SetLayout(args wire.SetLayoutArgs) (wire.LayoutResult, error) {
	var r wire.LayoutResult
	return r, v.c.call(wire.MethodViewsSetLayout, args, &r)
}
