// SPDX-License-Identifier: Apache-2.0

package contract

// EditableModelState is a model-state member whose owning collection supports
// naming and copying. The built-in Primary member cannot be renamed or deleted.
// A copy receives a new identity and does not replace the active source member.
type EditableModelState interface {
	ModelState
	// Primary identifies the built-in member without relying on its display name.
	Primary() bool
	// SetName changes a non-primary member's unique name without changing its ID.
	SetName(name string) error
	// Copy makes an inactive copy; an empty name asks the owner for a unique name.
	Copy(newName string) (ModelState, error)
}
