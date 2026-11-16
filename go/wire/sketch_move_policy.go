// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// PlanSketchMoveArgs requests a read-only move plan for entities in the active document.
// SketchIndex is the sketch index; Entities is a nonempty set of unique session
// IDs in that sketch; Vector is exactly two finite values in cm.
// Both policies are required. Empty, unknown or unsupported policies are errors.
// The host must not substitute the legacy sketch.transform operation.
type PlanSketchMoveArgs struct {
	SketchIndex     int                             `json:"sketchIndex"`
	Entities        []uint64                        `json:"entities"`
	Vector          []float64                       `json:"vector"`
	DimensionPolicy types.SketchMoveDimensionPolicy `json:"dimensionPolicy"`
	GeometricPolicy types.SketchMoveGeometricPolicy `json:"geometricPolicy"`
}

// SketchMoveDimensionChange describes a dimension in the exact planned revision.
// Index is the index returned by sketch.dimensions for that same revision.
type SketchMoveDimensionChange struct {
	Index            int    `json:"index"`
	BeforeExpression string `json:"beforeExpression"`
	AfterExpression  string `json:"afterExpression"`
	ApprovalRequired bool   `json:"approvalRequired"`
}

// SketchMoveGeometricChange describes a relation proposed for removal.
// Index is the index returned by sketch.constraints for that same revision.
type SketchMoveGeometricChange struct {
	Index            int      `json:"index"`
	Kind             string   `json:"kind"`
	Entities         []uint64 `json:"entities"`
	ApprovalRequired bool     `json:"approvalRequired"`
}

// PlanSketchMoveResult reports every proposed constraint change without editing.
// Token is nonempty, opaque, one-use, scoped to this session, document, sketch,
// request and exact document state, and expires within a bounded host lifetime.
// Dimensions and Geometric list all proposed changes, including indirect effects;
// each index occurs at most once in its list. Hosts send empty arrays when there
// are none. Under a Prompt policy, each change of that kind requires approval.
// A Never policy cannot produce a change in its category. The result must satisfy all
// retained constraints. A host must refuse a plan it cannot describe, solve,
// or preserve; it must not move geometry or change history in planning.
type PlanSketchMoveResult struct {
	Token      string                      `json:"token"`
	Dimensions []SketchMoveDimensionChange `json:"dimensions"`
	Geometric  []SketchMoveGeometricChange `json:"geometric"`
}

// CommitSketchMoveArgs accepts a previously inspected plan. The caller sets each
// approval flag only after accepting all approval-required changes of that kind.
// To cancel or decline, discard the token; do not call commit. A noninteractive
// caller may choose a non-prompt policy or explicitly accept a described plan.
type CommitSketchMoveArgs struct {
	Token             string `json:"token"`
	ApproveDimensions bool   `json:"approveDimensions"`
	ApproveGeometric  bool   `json:"approveGeometric"`
}

// CommitSketchMoveResult confirms one atomic history edit was committed.
// Missing approvals, an expired, used, unknown or stale token, changed document
// state, invalid geometry, or any unsupported policy must return an error without
// changing geometry, constraints, dimensions, or history. The committed result
// must match the described plan; it must not add effects after approval. Undo and
// Redo restore all effects of the accepted edit together. Committed must be true
// on success. A caller must check Committed even when transport returned no error.
type CommitSketchMoveResult struct {
	Committed bool `json:"committed"`
}
