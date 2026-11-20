// SPDX-License-Identifier: Apache-2.0

package wire

// ConstraintGlyphVisibilityArgs identifies one geometric constraint by its
// stable relation ID and changes only its glyph visibility. The host
// rejects stale or invalid IDs without modifying the sketch. Example:
// ConstraintGlyphVisibilityArgs{SketchIndex: 0, ConstraintID: 1, Visible: false}.
type ConstraintGlyphVisibilityArgs struct {
	SketchIndex  int    `json:"sketchIndex"`
	ConstraintID uint64 `json:"constraintId"`
	Visible      bool   `json:"visible"`
}

// AllConstraintGlyphsVisibilityArgs changes every geometric-constraint glyph
// in the sketch. Showing all clears individual hidden overrides; hiding all
// does not change geometry, dimensions, or constraint enforcement. Example:
// AllConstraintGlyphsVisibilityArgs{SketchIndex: 0, Visible: true}.
type AllConstraintGlyphsVisibilityArgs struct {
	SketchIndex int  `json:"sketchIndex"`
	Visible     bool `json:"visible"`
}

// ConstraintGlyphsVisibilityArgs selects a sketch for a visibility query.
// Example: ConstraintGlyphsVisibilityArgs{SketchIndex: 0}.
type ConstraintGlyphsVisibilityArgs struct {
	SketchIndex int `json:"sketchIndex"`
}

// ConstraintGlyphsVisibilityResult reports the global display switch and the
// stable geometric relation IDs and those hidden individually, in ascending
// order. Shown false masks every relation; HiddenIDs still reports overrides.
// The IDs belong to the specified sketch. Example:
// ConstraintGlyphsVisibilityResult{Shown: true, RelationIDs: []uint64{1}, HiddenIDs: []uint64{1}}.
type ConstraintGlyphsVisibilityResult struct {
	Shown       bool     `json:"shown"`
	RelationIDs []uint64 `json:"relationIds"`
	HiddenIDs   []uint64 `json:"hiddenIds"`
}
