// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// TransformSketchArgs is the request of [MethodSketchTransform] — the discriminated edit
// operation on a selection of entities. Op is "move" | "rotate" | "copy" | "mirror":
//
//   - move:   translate the selection in place by Vector ([dx,dy] cm).
//   - copy:   duplicate the selection, offset by Vector; the copies are returned.
//   - rotate: rotate the selection in place about Center ([x,y] cm) by Angle (a
//     unit-bearing expression like "90 deg").
//   - mirror: reflect the selection across the line MirrorLine (an entity id); the
//     mirrored copies are returned.
//
// DimensionPolicy and GeometricPolicy apply to move only. An omitted policy means
// "never". A host that needs a decision for "prompt" returns DecisionRequired
// without changing the sketch; the caller can ask the user and retry explicitly.
type TransformSketchArgs struct {
	SketchIndex     int                             `json:"sketchIndex"`
	Op              string                          `json:"op"`
	Entities        []uint64                        `json:"entities"`
	Vector          []float64                       `json:"vector,omitempty"`
	Center          []float64                       `json:"center,omitempty"`
	Angle           string                          `json:"angle,omitempty"`
	MirrorLine      uint64                          `json:"mirrorLine,omitempty"`
	DimensionPolicy types.SketchMoveDimensionPolicy `json:"dimensionPolicy,omitempty"`
	GeometricPolicy types.SketchMoveGeometricPolicy `json:"geometricPolicy,omitempty"`
}

// SketchMoveDecision identifies which constraint policies require a user decision.
// At least one field is true. No sketch edit has been committed when this is returned.
type SketchMoveDecision struct {
	Dimension bool `json:"dimension,omitempty"`
	Geometric bool `json:"geometric,omitempty"`
}

// TransformSketchResult is the response of [MethodSketchTransform]: the ids of any newly
// created entities, or a decision needed before a move can be committed.
type TransformSketchResult struct {
	Created          []uint64            `json:"created,omitempty"`
	DecisionRequired *SketchMoveDecision `json:"decisionRequired,omitempty"`
}
