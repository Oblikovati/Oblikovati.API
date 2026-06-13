// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The assembly occurrence surface (M11-F01/F02, Oblikovati/Oblikovati#728): read the
// active assembly's occurrence tree and place/transform/ground/suppress/replace/remove
// components. Occurrences are addressed by session id (uint64) — the same ids the
// occurrence push events carry. Placements are an axis-aligned 4×4 transform
// ([types.Matrix], row-major) in the assembly's space.

// OccurrenceInfo is one node of the assembly occurrence tree: a placed component, its
// per-instance state, and any nested sub-assembly occurrences. Default-false state flags
// and an empty child list are omitted.
type OccurrenceInfo struct {
	ID         uint64           `json:"id"`
	Name       string           `json:"name"`
	Transform  types.Matrix     `json:"transform"`
	Suppressed bool             `json:"suppressed,omitempty"`
	Grounded   bool             `json:"grounded,omitempty"`
	Adaptive   bool             `json:"adaptive,omitempty"`
	Substitute bool             `json:"substitute,omitempty"`
	Children   []OccurrenceInfo `json:"children,omitempty"`
}

// OccurrencesResult is the reply of [MethodAssemblyOccurrences] and of [MethodAssemblyRemove]:
// the active assembly's occurrence tree (top-level occurrences, each with nested children).
type OccurrencesResult struct {
	Occurrences []OccurrenceInfo `json:"occurrences"`
}

// OccurrenceResult is the reply of the single-occurrence operations (place, transform,
// ground, suppress, replace): the affected occurrence's refreshed info.
type OccurrenceResult struct {
	Occurrence OccurrenceInfo `json:"occurrence"`
}

// PlaceOccurrenceArgs is the request of [MethodAssemblyPlace]: place the component held by
// the open Document (by document id) into the active assembly under Name at Transform. The
// document's content is shared (the flyweight), so every placement tracks its edits. The
// document must be an open part or assembly, not a drawing or reference stub.
type PlaceOccurrenceArgs struct {
	Document  uint64       `json:"document"`
	Name      string       `json:"name"`
	Transform types.Matrix `json:"transform"`
}

// PlaceByDefinitionArgs is the request of [MethodAssemblyPlaceByDefinition]: place another
// instance of the component that occurrence Source (by session id) already instances, under
// Name at Transform — reusing the shared definition without re-resolving a document.
type PlaceByDefinitionArgs struct {
	Source    uint64       `json:"source"`
	Name      string       `json:"name"`
	Transform types.Matrix `json:"transform"`
}

// TransformOccurrenceArgs is the request of [MethodAssemblyTransform]: reposition the
// occurrence with id ID to Transform (its placement in the assembly's space).
type TransformOccurrenceArgs struct {
	ID        uint64       `json:"id"`
	Transform types.Matrix `json:"transform"`
}

// GroundOccurrenceArgs is the request of [MethodAssemblyGround]: fix (Grounded=true) or
// release the occurrence with id ID in the assembly's space.
type GroundOccurrenceArgs struct {
	ID       uint64 `json:"id"`
	Grounded bool   `json:"grounded"`
}

// SuppressOccurrenceArgs is the request of [MethodAssemblySuppress]: exclude
// (Suppressed=true) or restore the occurrence with id ID from/to the model.
type SuppressOccurrenceArgs struct {
	ID         uint64 `json:"id"`
	Suppressed bool   `json:"suppressed"`
}

// ReplaceOccurrenceArgs is the request of [MethodAssemblyReplace]: swap the component of the
// occurrence with id ID for the one held by the open Document (by document id), keeping the
// occurrence's id, name, transform, and state — the "replace component" operation.
type ReplaceOccurrenceArgs struct {
	ID       uint64 `json:"id"`
	Document uint64 `json:"document"`
}

// RemoveOccurrenceArgs is the request of [MethodAssemblyRemove]: delete the occurrence with
// id ID from the active assembly.
type RemoveOccurrenceArgs struct {
	ID uint64 `json:"id"`
}
