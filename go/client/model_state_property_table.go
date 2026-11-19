// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// GetPropertyTable reads the active part's model-state property table.
// Example: `view, err := client.ModelStates().GetPropertyTable()`.
func (m ModelStates) GetPropertyTable() (wire.ModelStatePropertyTableView, error) {
	return call[wire.ModelStatePropertyTableView](m.c, wire.MethodModelStatesGetPropertyTable, struct{}{})
}

// SetPropertyTable atomically applies displayed columns and member values.
// Example: `view, err := client.ModelStates().SetPropertyTable(args)`.
func (m ModelStates) SetPropertyTable(args wire.SetModelStatePropertyTableArgs) (wire.ModelStatePropertyTableView, error) {
	return call[wire.ModelStatePropertyTableView](m.c, wire.MethodModelStatesSetPropertyTable, args)
}
