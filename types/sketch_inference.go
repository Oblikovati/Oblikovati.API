// SPDX-License-Identifier: Apache-2.0

package types

// SketchPointInferenceKind describes how a placed sketch point was inferred
// from nearby geometry (M06-F10, Oblikovati/Oblikovati#625). This is the
// sketch-domain inference vocabulary; the host-level drag-snap kind stays in
// [PointInferenceKind].
//
// The values are a frozen block matching the reference API's point-inference
// enum; never renumber them.
type SketchPointInferenceKind int32

const (
	// SketchInferenceAtIntersection placed the point at two curves' crossing.
	SketchInferenceAtIntersection SketchPointInferenceKind = 22273
	// SketchInferenceOnCurve placed the point on a curve's body.
	SketchInferenceOnCurve SketchPointInferenceKind = 22274
	// SketchInferenceOnPoint snapped the point onto an existing sketch point.
	SketchInferenceOnPoint SketchPointInferenceKind = 22275
	// SketchInferenceAtMidpoint placed the point at a curve's midpoint.
	SketchInferenceAtMidpoint SketchPointInferenceKind = 22276
)

// sketchPointInferenceNames are the frozen wire spellings.
var sketchPointInferenceNames = map[SketchPointInferenceKind]string{
	SketchInferenceAtIntersection: "atIntersection",
	SketchInferenceOnCurve:        "onCurve",
	SketchInferenceOnPoint:        "onPoint",
	SketchInferenceAtMidpoint:     "atMidpoint",
}

// String returns the inference kind's wire spelling.
func (k SketchPointInferenceKind) String() string {
	return enumName(sketchPointInferenceNames, k, "enum(?)")
}

// ParseSketchPointInferenceKind resolves a wire spelling back to its kind.
func ParseSketchPointInferenceKind(s string) (SketchPointInferenceKind, bool) {
	return enumFromName(sketchPointInferenceNames, s)
}

// ConstraintInferenceKind types one geometric constraint the inference engine
// proposed or auto-applied while sketching (M06-F10).
//
// The values are a frozen block matching the reference API's
// constraint-inference enum. They are individual bit flags (1 << n) so a set
// of inference kinds can travel as a mask; never renumber them.
type ConstraintInferenceKind int32

const (
	InferCoincident    ConstraintInferenceKind = 1
	InferHorizontal    ConstraintInferenceKind = 2
	InferIntersection  ConstraintInferenceKind = 4
	InferMidpoint      ConstraintInferenceKind = 8
	InferOnCurve       ConstraintInferenceKind = 16
	InferParallel      ConstraintInferenceKind = 32
	InferPerpendicular ConstraintInferenceKind = 64
	InferTangent       ConstraintInferenceKind = 128
	InferVertical      ConstraintInferenceKind = 256
)

// constraintInferenceNames are the frozen wire spellings.
var constraintInferenceNames = map[ConstraintInferenceKind]string{
	InferCoincident:    "coincident",
	InferHorizontal:    "horizontal",
	InferIntersection:  "intersection",
	InferMidpoint:      "midpoint",
	InferOnCurve:       "onCurve",
	InferParallel:      "parallel",
	InferPerpendicular: "perpendicular",
	InferTangent:       "tangent",
	InferVertical:      "vertical",
}

// String returns the inference kind's wire spelling.
func (k ConstraintInferenceKind) String() string {
	return enumName(constraintInferenceNames, k, "enum(?)")
}

// ParseConstraintInferenceKind resolves a wire spelling back to its kind.
func ParseConstraintInferenceKind(s string) (ConstraintInferenceKind, bool) {
	return enumFromName(constraintInferenceNames, s)
}

// ConstraintInferencePriority is the user preference for which constraint
// family wins when the inference engine could apply either (M06-F10).
//
// The values are a frozen block matching the reference API's
// constraint-priority enum; never renumber them.
type ConstraintInferencePriority int32

const (
	// PriorityParallelPerpendicular prefers parallel/perpendicular over
	// horizontal/vertical when both fit.
	PriorityParallelPerpendicular ConstraintInferencePriority = 50433
	// PriorityHorizontalVertical prefers horizontal/vertical (the default).
	PriorityHorizontalVertical ConstraintInferencePriority = 50434
	// PriorityNone applies no family preference.
	PriorityNone ConstraintInferencePriority = 50435
)

// constraintInferencePriorityNames are the frozen wire spellings.
var constraintInferencePriorityNames = map[ConstraintInferencePriority]string{
	PriorityParallelPerpendicular: "parallelPerpendicular",
	PriorityHorizontalVertical:    "horizontalVertical",
	PriorityNone:                  "none",
}

// String returns the priority's wire spelling.
func (p ConstraintInferencePriority) String() string {
	return enumName(constraintInferencePriorityNames, p, "enum(?)")
}

// ParseConstraintInferencePriority resolves a wire spelling back to its priority.
func ParseConstraintInferencePriority(s string) (ConstraintInferencePriority, bool) {
	return enumFromName(constraintInferencePriorityNames, s)
}
