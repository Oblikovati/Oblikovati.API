// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Model is the read-only model-inspection operation group.
type Model struct{ c *Client }

// Model returns the model-inspection operation group.
func (c *Client) Model() Model { return Model{c} }

// Tree returns a read-only snapshot of the active part's structure.
//
// mcp:tool get_model_tree
// mcp:summary Read the active part's structure: parameters, sketches, features, body count.
func (m Model) Tree() (wire.ModelTreeResult, error) {
	return call[wire.ModelTreeResult](m.c, wire.MethodModelTree, nil)
}

// Selection returns the current selection summary.
//
// mcp:tool get_selection
// mcp:summary Read the current selection.
func (m Model) Selection() (wire.SelectionResult, error) {
	return call[wire.SelectionResult](m.c, wire.MethodModelSelection, nil)
}

// Select selects the entities named by their reference strings (from a SelectionResult). Mode
// "add" extends the current selection; "replace" (or empty) replaces it. Returns the new selection.
//
// mcp:tool select_entities
// mcp:summary Select model entities by their reference strings (mode add|replace).
func (m Model) Select(refs []string, mode string) (wire.SelectionResult, error) {
	return call[wire.SelectionResult](m.c, wire.MethodModelSelect, wire.SelectArgs{Refs: refs, Mode: mode})
}

// Deselect removes the named entities from the current selection. Returns the new selection.
//
// mcp:tool deselect_entities
// mcp:summary Remove model entities (by reference string) from the current selection.
func (m Model) Deselect(refs []string) (wire.SelectionResult, error) {
	return call[wire.SelectionResult](m.c, wire.MethodModelDeselect, wire.DeselectArgs{Refs: refs})
}

// ClearSelection clears the whole selection. Returns the now-empty selection.
//
// mcp:tool clear_selection
// mcp:summary Clear the current selection.
func (m Model) ClearSelection() (wire.SelectionResult, error) {
	return call[wire.SelectionResult](m.c, wire.MethodModelClearSelection, nil)
}

// CreateHighlightSet adds a named, colored emphasis group (#157) the viewport outlines without
// selecting; add references with AddHighlightItems.
//
// mcp:tool create_highlight_set
// mcp:summary Create a named, colored highlight set (emphasis group) the viewport outlines without selecting.
func (m Model) CreateHighlightSet(name, color string) (wire.HighlightSetInfo, error) {
	return call[wire.HighlightSetInfo](m.c, wire.MethodModelHighlightSetCreate, wire.CreateHighlightSetArgs{Name: name, Color: color})
}

// AddHighlightItems adds model references (from ReferenceKeys) to a highlight set (#157).
//
// mcp:tool add_highlight_items
// mcp:summary Add model references to a highlight set.
func (m Model) AddHighlightItems(name string, refs []string) (wire.HighlightSetInfo, error) {
	return call[wire.HighlightSetInfo](m.c, wire.MethodModelHighlightSetAddItems, wire.HighlightSetItemsArgs{Name: name, Refs: refs})
}

// SetHighlightSetColor re-colours a highlight set to a "#rrggbb" colour (#157).
//
// mcp:tool set_highlight_set_color
// mcp:summary Re-colour a highlight set.
func (m Model) SetHighlightSetColor(name, color string) (wire.HighlightSetInfo, error) {
	return call[wire.HighlightSetInfo](m.c, wire.MethodModelHighlightSetSetColor, wire.SetHighlightSetColorArgs{Name: name, Color: color})
}

// DeleteHighlightSet removes a highlight set (#157).
//
// mcp:tool delete_highlight_set
// mcp:summary Delete a highlight set.
func (m Model) DeleteHighlightSet(name string) (wire.OKResult, error) {
	return call[wire.OKResult](m.c, wire.MethodModelHighlightSetDelete, wire.HighlightSetRefArgs{Name: name})
}

// HighlightSets lists the active session's highlight sets (#157).
//
// mcp:tool list_highlight_sets
// mcp:summary List the session's highlight sets (name, colour, item count).
func (m Model) HighlightSets() (wire.ListHighlightSetsResult, error) {
	return call[wire.ListHighlightSetsResult](m.c, wire.MethodModelHighlightSetList, nil)
}

// ReferenceKeys returns the active part's topology (faces/edges/vertices) with their
// persistent reference keys — the keys consumed by Include / AddSurfaceCurve / Project /
// attributes. It is how an add-in obtains a key without a viewport pick.
//
// mcp:tool get_reference_keys
// mcp:summary List the active part's persistent topology reference keys (stable ids for faces/edges/vertices) — the refs project_geometry, include, and work planes consume.
func (m Model) ReferenceKeys() (wire.ReferenceKeysResult, error) {
	return call[wire.ReferenceKeysResult](m.c, wire.MethodModelReferenceKeys, nil)
}
