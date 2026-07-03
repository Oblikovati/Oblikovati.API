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
//
// mcp:tool ui_search
// mcp:summary Finds registered commands matching the query (the command search box's backing method).
func (u UI) Search(query string) (wire.SearchCommandsResult, error) {
	return call[wire.SearchCommandsResult](u.c, wire.MethodUISearch, wire.SearchCommandsArgs{Query: query})
}

// MarkingMenu returns one environment's radial marking menu.
//
// mcp:tool ui_get_marking_menu
// mcp:summary Returns one environment's radial marking menu.
func (u UI) MarkingMenu(env types.Environment) (wire.MarkingMenuView, error) {
	return call[wire.MarkingMenuView](u.c, wire.MethodUIGetMarkingMenu, wire.GetMarkingMenuArgs{Environment: env})
}

// SetMarkingMenu replaces one environment's radial marking menu — the
// OnRadialMarkingMenu customization hook, declarative.
//
// mcp:tool ui_set_marking_menu
// mcp:summary Replaces one environment's radial marking menu — the OnRadialMarkingMenu customization hook, declarative.
func (u UI) SetMarkingMenu(menu wire.MarkingMenuView) (wire.OKResult, error) {
	return call[wire.OKResult](u.c, wire.MethodUISetMarkingMenu, wire.SetMarkingMenuArgs{Menu: menu})
}

// SetContextMenu replaces this add-in's injected context-menu entries for one
// browser node kind ("" injects everywhere) — the OnContextMenu equivalent.
//
// mcp:tool ui_set_context_menu
// mcp:summary Replaces this add-in's injected context-menu entries for one browser node kind ("" injects everywhere) — the OnContextMenu equivalent.
func (u UI) SetContextMenu(addin, kind string, items []wire.ContextMenuItemSpec) (wire.OKResult, error) {
	args := wire.SetContextMenuArgs{AddIn: addin, Kind: kind, Items: items}
	return call[wire.OKResult](u.c, wire.MethodUISetContextMenu, args)
}

// ObjectVisibility returns the View ▸ Object-visibility toggles.
//
// mcp:tool ui_get_object_visibility
// mcp:summary Returns the View ▸ Object-visibility toggles.
func (u UI) ObjectVisibility() (wire.ObjectVisibilityView, error) {
	return call[wire.ObjectVisibilityView](u.c, wire.MethodUIGetObjectVisibility, nil)
}

// SetObjectVisibility writes the View ▸ Object-visibility toggles (hidden
// geometry also stops being pickable).
//
// mcp:tool ui_set_object_visibility
// mcp:summary Writes the View ▸ Object-visibility toggles (hidden geometry also stops being pickable).
func (u UI) SetObjectVisibility(v wire.ObjectVisibilityView) (wire.OKResult, error) {
	return call[wire.OKResult](u.c, wire.MethodUISetObjectVisibility, wire.SetObjectVisibilityArgs{Visibility: v})
}

// RegisterEnvironment declares this add-in's contextual UI environment (value ≥ 2;
// commands created with it form the environment's tabs) — M05-F16.
//
//	client.UI().RegisterEnvironment(7, "Weldment")
//
// mcp:tool ui_register_environment
// mcp:summary Declares this add-in's contextual UI environment (value ≥ 2; commands created with it form the environment's tabs) — M05-F16.
func (u UI) RegisterEnvironment(env types.Environment, name string) (wire.OKResult, error) {
	args := wire.RegisterEnvironmentArgs{Environment: env, Name: name}
	return call[wire.OKResult](u.c, wire.MethodUIRegisterEnvironment, args)
}

// ActivateEnvironment enters a registered environment (base, 0, leaves it); the
// switch reaches every add-in as a ui.environmentChanged event.
//
// mcp:tool ui_activate_environment
// mcp:summary Enters a registered environment (base, 0, leaves it); the switch reaches every add-in as a ui.environmentChanged event.
func (u UI) ActivateEnvironment(env types.Environment) (wire.OKResult, error) {
	return call[wire.OKResult](u.c, wire.MethodUIActivateEnvironment, wire.ActivateEnvironmentArgs{Environment: env})
}
