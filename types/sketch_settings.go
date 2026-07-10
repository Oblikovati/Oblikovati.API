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

	// Snap/grid defaults (Inventor SketchSettings, #1877). XSnapSpacing/YSnapSpacing are the grid
	// snap increments in cm; SnapsPerMinorGrid subdivides a minor grid cell and
	// MinorLinesPerMajorGridLine sets how many minor lines fall between major grid lines.
	XSnapSpacing               float64 `json:"xSnapSpacing"`
	YSnapSpacing               float64 `json:"ySnapSpacing"`
	SnapsPerMinorGrid          int     `json:"snapsPerMinorGrid"`
	MinorLinesPerMajorGridLine int     `json:"minorLinesPerMajorGridLine"`

	// Constraint display / creation behaviour (Inventor SketchConstraintSettings, #1877).
	// PersistInferredConstraints keeps inferred constraints as persistent relations — Inventor's
	// EnablePersistConstraints, distinct from AutoApplyConstraints (which governs whether inference
	// applies them at all). DisplayConstraintsOnCreation shows constraint glyphs as geometry is
	// placed; EditDimensionsWhenCreated pops the value editor for a new dimension; OverConstrained
	// Behavior decides what a redundant dimension does.
	PersistInferredConstraints   bool                             `json:"persistInferredConstraints"`
	DisplayConstraintsOnCreation bool                             `json:"displayConstraintsOnCreation"`
	EditDimensionsWhenCreated    bool                             `json:"editDimensionsWhenCreated"`
	OverConstrainedBehavior      OverConstrainedDimensionBehavior `json:"overConstrainedBehavior"`

	// Relax-mode settings (Inventor SketchConstraintSettings, #1877): EnableRelaxMode lets dragging
	// drop conflicting constraints, and KeepDimensionsWithEquationInRelaxMode preserves dimensions
	// whose value is an equation while relaxing. Inventor's GeometricConstraintsToRemoveInRelaxMode
	// (a constraint-family bitmask selecting which relations relax may drop) is intentionally out of
	// scope: Oblikovati's solver has no selective per-family relax pass to configure yet, and adding
	// a list field would also make this value struct non-comparable. #1877.
	EnableRelaxMode                       bool `json:"enableRelaxMode"`
	KeepDimensionsWithEquationInRelaxMode bool `json:"keepDimensionsWithEquationInRelaxMode"`
}

// DefaultSketchSettings is the out-of-the-box configuration: inference and auto-apply on with
// horizontal/vertical preferred (the pre-#1877 behaviour), plus the Inventor-aligned grid/snap and
// constraint-display defaults — a 1 mm snap grid, persisted inferred constraints, the dimension
// editor on create, and redundant dimensions added as driven.
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
		DisplayConstraintsOnCreation: false,
		EditDimensionsWhenCreated:    true,
		OverConstrainedBehavior:      OverConstrainedApplyDriven,

		EnableRelaxMode:                       false,
		KeepDimensionsWithEquationInRelaxMode: true,
	}
}
