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
	// CustomInferenceTypes selects InferenceTypes; false retains the legacy inference policy.
	CustomInferenceTypes bool                 `json:"customInferenceTypes,omitempty"`
	InferenceTypes       SketchInferenceTypes `json:"inferenceTypes,omitempty"`

	// Snap/grid defaults (the reference CAD API's SketchSettings, #1877). XSnapSpacing/YSnapSpacing are the grid
	// snap increments in cm; SnapsPerMinorGrid subdivides a minor grid cell and
	// MinorLinesPerMajorGridLine sets how many minor lines fall between major grid lines.
	XSnapSpacing               float64 `json:"xSnapSpacing"`
	YSnapSpacing               float64 `json:"ySnapSpacing"`
	SnapsPerMinorGrid          int     `json:"snapsPerMinorGrid"`
	MinorLinesPerMajorGridLine int     `json:"minorLinesPerMajorGridLine"`

	// Constraint display / creation behaviour (the reference CAD API's SketchConstraintSettings, #1877).
	// PersistInferredConstraints keeps inferred constraints as persistent relations — the reference CAD API's
	// EnablePersistConstraints, distinct from AutoApplyConstraints (which governs whether inference
	// applies them at all). DisplayConstraintsOnCreation shows constraint glyphs as geometry is
	// placed; EditDimensionsWhenCreated pops the value editor for a new dimension; OverConstrained
	// Behavior decides what a redundant dimension does.
	PersistInferredConstraints   bool                             `json:"persistInferredConstraints"`
	DisplayConstraintsOnCreation bool                             `json:"displayConstraintsOnCreation"`
	EditDimensionsWhenCreated    bool                             `json:"editDimensionsWhenCreated"`
	OverConstrainedBehavior      OverConstrainedDimensionBehavior `json:"overConstrainedBehavior"`

	// CustomRelaxTypes selects RelaxTypes; false retains the legacy removal policy.
	// Selective removal was previously unavailable because the solver did not expose it (#1877).
	CustomRelaxTypes bool             `json:"customRelaxTypes,omitempty"`
	RelaxTypes       SketchRelaxTypes `json:"relaxTypes,omitempty"`
	// EnableRelaxMode and KeepDimensionsWithEquationInRelaxMode control relaxation.
	EnableRelaxMode                       bool `json:"enableRelaxMode"`
	KeepDimensionsWithEquationInRelaxMode bool `json:"keepDimensionsWithEquationInRelaxMode"`
}

// DefaultSketchSettings is the out-of-the-box configuration: inference and auto-apply on with
// horizontal/vertical preferred (the pre-#1877 behaviour), plus the reference-API-aligned grid/snap and
// constraint-display defaults — a 1 mm snap grid, persisted inferred constraints, constraint
// glyphs shown while geometry is placed, the dimension editor on create, and redundant
// dimensions added as driven.
func DefaultSketchSettings() SketchSettings {
	return SketchSettings{
		InferConstraints:     true,
		AutoApplyConstraints: true,
		ConstraintPriority:   PriorityHorizontalVertical,

		XSnapSpacing:               0.1,
		YSnapSpacing:               0.1,
		SnapsPerMinorGrid:          1,
		MinorLinesPerMajorGridLine: 10,

		PersistInferredConstraints:   true,
		DisplayConstraintsOnCreation: true,
		EditDimensionsWhenCreated:    true,
		OverConstrainedBehavior:      OverConstrainedApplyDriven,

		EnableRelaxMode:                       false,
		KeepDimensionsWithEquationInRelaxMode: true,
	}
}
