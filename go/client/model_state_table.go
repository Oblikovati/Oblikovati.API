// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// GetTable reads the active part's complete model-state table.
// Example: `view, err := client.ModelStates().GetTable()`.
//
// mcp:tool get_model_state_table
// mcp:summary Read the active part's model-state table and its generation.
func (m ModelStates) GetTable() (wire.ModelStateTableView, error) {
	return call[wire.ModelStateTableView](m.c, wire.MethodModelStatesGetTable, struct{}{})
}

// SetTable applies a mixed table draft as one undoable transaction.
// Example: `result, err := client.ModelStates().SetTable(args)`.
//
// mcp:tool set_model_state_table
// mcp:summary Apply a model-state table draft as one undoable edit. Use the current table generation.
func (m ModelStates) SetTable(args wire.SetModelStateTableArgs) (wire.ModelStateTableResult, error) {
	return call[wire.ModelStateTableResult](m.c, wire.MethodModelStatesSetTable, args)
}

// VerifyTable checks the same mixed draft without creating an edit.
// Example: `result, err := client.ModelStates().VerifyTable(args)`.
//
// mcp:tool verify_model_state_table
// mcp:summary Validate a model-state table draft without applying the edit.
func (m ModelStates) VerifyTable(args wire.SetModelStateTableArgs) (wire.ModelStateTableVerification, error) {
	return call[wire.ModelStateTableVerification](m.c, wire.MethodModelStatesVerifyTable, args)
}
