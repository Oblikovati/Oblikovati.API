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
	var r wire.ModelTreeResult
	return r, m.c.call(wire.MethodModelTree, nil, &r)
}

// Selection returns the current selection summary.
//
// mcp:tool get_selection
// mcp:summary Read the current selection.
func (m Model) Selection() (wire.SelectionResult, error) {
	var r wire.SelectionResult
	return r, m.c.call(wire.MethodModelSelection, nil, &r)
}

// Select selects the entities named by their reference strings (from a SelectionResult). Mode
// "add" extends the current selection; "replace" (or empty) replaces it. Returns the new selection.
//
// mcp:tool select_entities
// mcp:summary Select model entities by their reference strings (mode add|replace).
func (m Model) Select(refs []string, mode string) (wire.SelectionResult, error) {
	var r wire.SelectionResult
	return r, m.c.call(wire.MethodModelSelect, wire.SelectArgs{Refs: refs, Mode: mode}, &r)
}

// Deselect removes the named entities from the current selection. Returns the new selection.
//
// mcp:tool deselect_entities
// mcp:summary Remove model entities (by reference string) from the current selection.
func (m Model) Deselect(refs []string) (wire.SelectionResult, error) {
	var r wire.SelectionResult
	return r, m.c.call(wire.MethodModelDeselect, wire.DeselectArgs{Refs: refs}, &r)
}

// ClearSelection clears the whole selection. Returns the now-empty selection.
//
// mcp:tool clear_selection
// mcp:summary Clear the current selection.
func (m Model) ClearSelection() (wire.SelectionResult, error) {
	var r wire.SelectionResult
	return r, m.c.call(wire.MethodModelClearSelection, nil, &r)
}

// ReferenceKeys returns the active part's topology (faces/edges/vertices) with their
// persistent reference keys — the keys consumed by Include / AddSurfaceCurve / Project /
// attributes. It is how an add-in obtains a key without a viewport pick.
//
// mcp:tool get_reference_keys
// mcp:summary List the active part's persistent topology reference keys (stable ids for faces/edges/vertices) — the refs project_geometry, include, and work planes consume.
func (m Model) ReferenceKeys() (wire.ReferenceKeysResult, error) {
	var r wire.ReferenceKeysResult
	return r, m.c.call(wire.MethodModelReferenceKeys, nil, &r)
}
