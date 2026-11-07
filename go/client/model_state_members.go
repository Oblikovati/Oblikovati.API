// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Rename changes a non-primary model-state member's unique name without changing its id.
//
// Example: client.ModelStates().Rename(wire.ModelStateNameArgs{ID: 7, Name: "Prototype"}).
//
// mcp:tool rename_model_state
// mcp:summary Rename a model-state member while retaining its identity and active status. Primary cannot be renamed.
func (m ModelStates) Rename(args wire.ModelStateNameArgs) (wire.ModelStateResult, error) {
	return call[wire.ModelStateResult](m.c, wire.MethodModelStatesRename, args)
}

// Copy makes an inactive member from the source; an empty name requests a generated name.
//
// Example: client.ModelStates().Copy(wire.ModelStateNameArgs{ID: 7, Name: "Prototype Copy"}).
//
// mcp:tool copy_model_state
// mcp:summary Copy a model-state member without changing the active member. An empty name requests a unique generated name.
func (m ModelStates) Copy(args wire.ModelStateNameArgs) (wire.ModelStateResult, error) {
	return call[wire.ModelStateResult](m.c, wire.MethodModelStatesCopy, args)
}
