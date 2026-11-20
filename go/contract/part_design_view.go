// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// PartDesignView is a document-owned design representation. Its ID remains stable
// across save and reopen; it does not identify an assembly representation.
// Example: `if view.Active() { _ = view.AnnotationScale() }`.
type PartDesignView interface {
	ID() uint64
	Name() string
	BuiltIn() bool
	Primary() bool
	Active() bool
	Locked() bool
	Published() bool
	AutosaveCamera() bool
	AnnotationScale() types.PartAnnotationScale
	StandardOrientation() types.ViewOrientationTypeEnum
}
