// SPDX-License-Identifier: Apache-2.0

package wire

// Command alias & keyboard-shortcut customization (M05-F17, Oblikovati#831). The keymap
// is the catalog of bindable actions — every registered command plus the built-in
// session actions (undo/redo/cancel/commit/visibility) — each with its effective and
// default keyboard shortcut and any user-defined alias. Chords travel as their canonical
// text form ([oblikovati.org/api/types.KeyChord.String], e.g. "Ctrl+Shift+E"); an empty
// string means "unbound". Only the user's deltas from the defaults are persisted.

// BindingInfo is one bindable action in the keymap catalog ([MethodKeymapList]).
type BindingInfo struct {
	ActionID     string `json:"actionId"`
	DisplayName  string `json:"displayName"`
	Kind         string `json:"kind"`                   // "command" or "builtin"
	Chord        string `json:"chord,omitempty"`        // effective shortcut; "" if unbound
	DefaultChord string `json:"defaultChord,omitempty"` // out-of-the-box shortcut
	Alias        string `json:"alias,omitempty"`        // effective user alias; "" if none
	Customized   bool   `json:"customized,omitempty"`   // chord or alias differs from default
}

// ListBindingsResult is the response of [MethodKeymapList]: the full catalog, in a
// stable display order (commands first in registration order, then built-ins).
type ListBindingsResult struct {
	Bindings []BindingInfo `json:"bindings"`
}

// SetChordArgs is the request of [MethodKeymapSetChord]: rebind one action's keyboard
// shortcut. Chord is the canonical chord string; an empty Chord clears the binding. The
// host rejects a chord already bound to another action.
type SetChordArgs struct {
	ActionID string `json:"actionId"`
	Chord    string `json:"chord"`
}

// SetAliasArgs is the request of [MethodKeymapSetAlias]: set one action's typed command
// alias. An empty Alias clears it. The host rejects an alias already bound to another
// action.
type SetAliasArgs struct {
	ActionID string `json:"actionId"`
	Alias    string `json:"alias"`
}

// ResetBindingArgs is the request of [MethodKeymapReset]: restore one action's shortcut
// and alias to their defaults.
type ResetBindingArgs struct {
	ActionID string `json:"actionId"`
}

// KeymapExport is the response of [MethodKeymapExport] and the request of
// [MethodKeymapImport]: the user's full customization delta, portable across installs.
// Keys are action ids; chord values are canonical chord strings. Importing replaces the
// current customization.
type KeymapExport struct {
	Chords  map[string]string `json:"chords,omitempty"`
	Aliases map[string]string `json:"aliases,omitempty"`
}
