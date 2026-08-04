// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Sketch Format panel over the wire (Oblikovati/Oblikovati#2015): an entity's formatting
// overrides, and the panel's armed creation modes.

// SketchEntityFormatArgs addresses one sketch entity in the active sketch. It is the request of
// [MethodSketchGetEntityFormat].
type SketchEntityFormatArgs struct {
	EntityID uint64 `json:"entityId"`
}

// SketchEntityFormatView is one entity's formatting. HasFormat distinguishes an entity that
// overrides nothing from one whose overrides happen to be empty — without it a caller could not
// tell "inherits everything" from "explicitly default", which are the same picture but different
// states to write back.
type SketchEntityFormatView struct {
	EntityID  uint64                   `json:"entityId"`
	HasFormat bool                     `json:"hasFormat"`
	Format    types.SketchEntityFormat `json:"format"`
}

// SetSketchEntityFormatArgs sets one entity's formatting. A format that overrides nothing clears
// the entity's overrides, so writing back a default is how a caller removes them. It is the
// request of [MethodSketchSetEntityFormat].
type SetSketchEntityFormatArgs struct {
	EntityID uint64                   `json:"entityId"`
	Format   types.SketchEntityFormat `json:"format"`
}

// SketchFormatModesView is the Format panel's armed creation state. It is the result of
// [MethodSketchGetFormatModes] and the request of [MethodSketchSetFormatModes]; a set replaces
// every field, so read-modify-write to change one.
type SketchFormatModesView struct {
	Modes types.SketchFormatModes `json:"modes"`
}
