// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// SketchInferenceScopeArgs identifies a 2D sketch by its zero-based index in
// the current part. It is the request of
// [MethodSketchGetInferenceScope].
// Example: `args := wire.SketchInferenceScopeArgs{SketchIndex: 0}`.
type SketchInferenceScopeArgs struct {
	SketchIndex int `json:"sketchIndex"`
}

// SketchInferenceScopeView is the applied inference scope for a 2D sketch in
// the current document session. Entities are nonzero, distinct entity IDs owned by that sketch and
// are valid for that sketch session. SelectedGeometry requires at least one
// eligible entity; the other modes require an empty list. Selection is a set,
// and hosts return it in ascending ID order. A set request replaces the entire
// scope atomically; invalid references leave the prior scope unchanged.
// The scope is retained during the active sketch session. No default, disk
// persistence, or cross-session retention is implied. Preview and commit use
// the same applied scope. Canceling a UI draft does not set this value.
// Example: `wire.SketchInferenceScopeView{SketchIndex: 0, Scope: types.InferenceScopeAllGeometry}`.
type SketchInferenceScopeView struct {
	SketchIndex int                        `json:"sketchIndex"`
	Scope       types.SketchInferenceScope `json:"scope"`
	Entities    []uint64                   `json:"entities"`
}
