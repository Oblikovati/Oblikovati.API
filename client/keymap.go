// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// Keymap is the command alias & keyboard-shortcut customization group (M05-F17): the
// catalog of bindable actions and the operations to rebind shortcuts, set aliases,
// reset to defaults, and import/export the whole customization. Chords are passed as
// typed [types.KeyChord] values and travel as their canonical string form.
type Keymap struct{ c *Client }

// Keymap returns the keyboard-customization operation group.
func (c *Client) Keymap() Keymap { return Keymap{c} }

// List returns the full keymap catalog: every command and built-in action with its
// effective and default shortcut and any alias.
//
// mcp:tool keymap_list
// mcp:summary Returns the full catalog of bindable commands with their shortcuts and aliases.
func (k Keymap) List() (wire.ListBindingsResult, error) {
	return call[wire.ListBindingsResult](k.c, wire.MethodKeymapList, nil)
}

// SetChord rebinds one action's keyboard shortcut. A zero chord clears the binding; the
// host rejects a chord already bound to another action.
//
//	client.Keymap().SetChord("Feature.Extrude", types.KeyChord{Key: "E", Ctrl: true})
//
// mcp:tool keymap_set_chord
// mcp:summary Rebinds a command's keyboard shortcut.
func (k Keymap) SetChord(actionID types.ActionID, chord types.KeyChord) (wire.OKResult, error) {
	args := wire.SetChordArgs{ActionID: actionID, Chord: chord.String()}
	return call[wire.OKResult](k.c, wire.MethodKeymapSetChord, args)
}

// SetAlias sets one action's typed command alias. An empty alias clears it; the host
// rejects an alias already bound to another action.
//
// mcp:tool keymap_set_alias
// mcp:summary Sets a command's typed alias.
func (k Keymap) SetAlias(actionID types.ActionID, alias string) (wire.OKResult, error) {
	args := wire.SetAliasArgs{ActionID: actionID, Alias: alias}
	return call[wire.OKResult](k.c, wire.MethodKeymapSetAlias, args)
}

// Reset restores one action's shortcut and alias to their defaults.
//
// mcp:tool keymap_reset
// mcp:summary Restores one command's shortcut and alias to their defaults.
func (k Keymap) Reset(actionID types.ActionID) (wire.OKResult, error) {
	return call[wire.OKResult](k.c, wire.MethodKeymapReset, wire.ResetBindingArgs{ActionID: actionID})
}

// ResetAll restores every binding to its default, discarding all customization.
//
// mcp:tool keymap_reset_all
// mcp:summary Restores every keyboard shortcut and alias to its default.
func (k Keymap) ResetAll() (wire.OKResult, error) {
	return call[wire.OKResult](k.c, wire.MethodKeymapResetAll, nil)
}

// Export returns the user's full customization delta, portable across installs.
//
// mcp:tool keymap_export
// mcp:summary Exports the user's keyboard customization as a portable delta.
func (k Keymap) Export() (wire.KeymapExport, error) {
	return call[wire.KeymapExport](k.c, wire.MethodKeymapExport, nil)
}

// Import replaces the current customization with the given delta.
//
// mcp:tool keymap_import
// mcp:summary Replaces the keyboard customization with an imported delta.
func (k Keymap) Import(exp wire.KeymapExport) (wire.OKResult, error) {
	return call[wire.OKResult](k.c, wire.MethodKeymapImport, exp)
}
