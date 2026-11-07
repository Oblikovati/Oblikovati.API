// SPDX-License-Identifier: Apache-2.0

package wire

const (
	// MethodModelStatesRename changes a non-primary member's name, keeping its id.
	MethodModelStatesRename = "modelStates.rename"
	// MethodModelStatesCopy copies a source member without changing which member is active.
	MethodModelStatesCopy = "modelStates.copy"
)

// ModelStateNameArgs selects an existing member and requests a name for rename or copy.
// Rename requires a nonempty unique name. Copy may omit Name to request a generated name.
type ModelStateNameArgs struct {
	ID   uint64 `json:"id"`
	Name string `json:"name,omitempty"`
}
