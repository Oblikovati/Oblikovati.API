// SPDX-License-Identifier: Apache-2.0

package types

// SketchSettings is a part document's persisted sketch-authoring defaults (#147) — the per-document
// surface of the Document Settings dialog's Sketch tab. It lifts the constraint-inference preferences
// that were previously session-only into the document (persisted in the .obk), so each part keeps its
// own: whether inference snaps points and auto-applies constraints while sketching, and which
// constraint family wins when two could apply. (3D-sketch and modeling settings are separate future
// objects; this starts with the inference toggles the sketch tools already read.)
type SketchSettings struct {
	// InferConstraints runs point snapping (endpoint → existing point, intersection, midpoint, …)
	// while sketching, surfacing inferred relations.
	InferConstraints bool `json:"inferConstraints"`
	// AutoApplyConstraints commits the inferred constraints automatically as geometry is placed
	// (off leaves the snap as a hint without persisting the relation).
	AutoApplyConstraints bool `json:"autoApplyConstraints"`
	// ConstraintPriority picks the constraint family when the inference engine could apply either.
	ConstraintPriority ConstraintInferencePriority `json:"constraintPriority"`
}

// DefaultSketchSettings is the out-of-the-box configuration: inference and auto-apply on, with
// horizontal/vertical preferred — matching a new document's behaviour before this surface existed.
func DefaultSketchSettings() SketchSettings {
	return SketchSettings{
		InferConstraints:     true,
		AutoApplyConstraints: true,
		ConstraintPriority:   PriorityHorizontalVertical,
	}
}
