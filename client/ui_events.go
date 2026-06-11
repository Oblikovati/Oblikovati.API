// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// UI is the shell-customization operation group (M05-F12): command search, the
// radial marking menu, context-menu injection, and the object-visibility toggles.
// The matching push events are command.started, selection.changed and
// ui.environmentChanged.
type UI struct{ c *Client }

// UI returns the shell-customization operation group.
func (c *Client) UI() UI { return UI{c} }

// Search finds registered commands matching the query (the command search box's
// backing method).
func (u UI) Search(query string) (wire.SearchCommandsResult, error) {
	var r wire.SearchCommandsResult
	return r, u.c.call(wire.MethodUISearch, wire.SearchCommandsArgs{Query: query}, &r)
}

// MarkingMenu returns one environment's radial marking menu.
func (u UI) MarkingMenu(env types.Environment) (wire.MarkingMenuView, error) {
	var r wire.MarkingMenuView
	return r, u.c.call(wire.MethodUIGetMarkingMenu, wire.GetMarkingMenuArgs{Environment: env}, &r)
}

// SetMarkingMenu replaces one environment's radial marking menu — the
// OnRadialMarkingMenu customization hook, declarative.
func (u UI) SetMarkingMenu(menu wire.MarkingMenuView) (wire.OKResult, error) {
	var r wire.OKResult
	return r, u.c.call(wire.MethodUISetMarkingMenu, wire.SetMarkingMenuArgs{Menu: menu}, &r)
}

// SetContextMenu replaces this add-in's injected context-menu entries for one
// browser node kind ("" injects everywhere) — the OnContextMenu equivalent.
func (u UI) SetContextMenu(addin, kind string, items []wire.ContextMenuItemSpec) (wire.OKResult, error) {
	var r wire.OKResult
	args := wire.SetContextMenuArgs{AddIn: addin, Kind: kind, Items: items}
	return r, u.c.call(wire.MethodUISetContextMenu, args, &r)
}

// ObjectVisibility returns the View ▸ Object-visibility toggles.
func (u UI) ObjectVisibility() (wire.ObjectVisibilityView, error) {
	var r wire.ObjectVisibilityView
	return r, u.c.call(wire.MethodUIGetObjectVisibility, nil, &r)
}

// SetObjectVisibility writes the View ▸ Object-visibility toggles (hidden
// geometry also stops being pickable).
func (u UI) SetObjectVisibility(v wire.ObjectVisibilityView) (wire.OKResult, error) {
	var r wire.OKResult
	return r, u.c.call(wire.MethodUISetObjectVisibility, wire.SetObjectVisibilityArgs{Visibility: v}, &r)
}
