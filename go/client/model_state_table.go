// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// GetTable reads the active part's complete model-state table.
// Example: `view, err := client.ModelStates().GetTable()`.
func (m ModelStates) GetTable() (wire.ModelStateTableView, error) {
	return call[wire.ModelStateTableView](m.c, wire.MethodModelStatesGetTable, struct{}{})
}

// SetTable applies a mixed table draft as one undoable transaction.
// Example: `result, err := client.ModelStates().SetTable(args)`.
func (m ModelStates) SetTable(args wire.SetModelStateTableArgs) (wire.ModelStateTableResult, error) {
	return call[wire.ModelStateTableResult](m.c, wire.MethodModelStatesSetTable, args)
}

// VerifyTable checks the same mixed draft without creating an edit.
// Example: `result, err := client.ModelStates().VerifyTable(args)`.
func (m ModelStates) VerifyTable(args wire.SetModelStateTableArgs) (wire.ModelStateTableVerification, error) {
	return call[wire.ModelStateTableVerification](m.c, wire.MethodModelStatesVerifyTable, args)
}
