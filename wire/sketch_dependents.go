// SPDX-License-Identifier: Apache-2.0

package wire

// Sketch dependents (Oblikovati/Oblikovati#154). A sketch may be consumed by one or more
// features (an extrude's profile, a revolve's profile/centerline, a loft section). Enumerating
// them predicts the impact of deleting or editing the sketch — [MethodSketchDelete] rejects a
// sketch that a feature still uses and reports the offending dependents.

// SketchDependent is one thing that depends on a sketch — today a feature that consumes it,
// addressed by its stable id (from model.tree), with its display name and kind.
type SketchDependent struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Kind string `json:"kind"`
}

// SketchDependentsResult is the response of [MethodSketchDependents]: every dependent of the
// sketch addressed by [SketchArgs.SketchIndex], in history order (empty when nothing uses it).
type SketchDependentsResult struct {
	Dependents []SketchDependent `json:"dependents"`
}
