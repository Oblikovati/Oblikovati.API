// SPDX-License-Identifier: Apache-2.0

package wire

// Sketch inference (M06-F10, Oblikovati/Oblikovati#625): while geometry is
// placed, the inference engine snaps points onto nearby geometry and
// auto-applies geometric constraints; these DTOs reify what was inferred so
// add-ins can drive and audit it.

// AppliedConstraintInference reports one geometric constraint the engine
// auto-applied while an entity was created. Kind is the
// [oblikovati.org/api/types.ConstraintInferenceKind] wire spelling;
// ConstraintIndex addresses the created constraint in the sketch's constraint
// enumeration; Entities are the session ids the constraint ties together.
type AppliedConstraintInference struct {
	Kind            string   `json:"kind"`
	ConstraintIndex int      `json:"constraintIndex"`
	Entities        []uint64 `json:"entities"`
}

// AppliedPointInference reports how one defining point of a created entity was
// inferred. Kind is the [oblikovati.org/api/types.SketchPointInferenceKind]
// wire spelling; PointID is the created point's session id; Entities are the
// session ids of the geometry the inference referenced.
type AppliedPointInference struct {
	Kind     string   `json:"kind"`
	PointID  uint64   `json:"pointId"`
	Entities []uint64 `json:"entities,omitempty"`
}

// InferenceOptionsView is the sketch inference configuration: whether point
// and constraint inference run at all, and which constraint family wins when
// two could apply (the [oblikovati.org/api/types.ConstraintInferencePriority]
// wire spelling). It is the result of [MethodSketchGetInferenceOptions] and
// the request of [MethodSketchSetInferenceOptions]; on set, empty Priority
// keeps the current value.
type InferenceOptionsView struct {
	InferEnabled     bool   `json:"inferEnabled"`
	ConstrainEnabled bool   `json:"constrainEnabled"`
	Priority         string `json:"priority,omitempty"`
}
