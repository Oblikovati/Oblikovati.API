// SPDX-License-Identifier: Apache-2.0

package types

// OverConstrainedDimensionBehavior is a document's preference for what happens when a dimension the
// user adds would over-constrain the sketch (the reference CAD API's SketchConstraintSettings.
// OverConstrainedDimensionBehavior, #1877). The reference enum's numeric ids are not in-repo, so
// this is an Oblikovati-owned frozen block — the wire contract is the string spelling; never
// renumber the ids.
type OverConstrainedDimensionBehavior int32

const (
	// OverConstrainedApplyDriven adds the redundant dimension as a driven (reference) dimension,
	// which measures but does not constrain — the safe default.
	OverConstrainedApplyDriven OverConstrainedDimensionBehavior = iota
	// OverConstrainedApplyDriving adds it as a driving dimension, accepting the redundancy (the
	// sketch reports over-constrained).
	OverConstrainedApplyDriving
	// OverConstrainedPrompt defers the choice to the user each time.
	OverConstrainedPrompt
)

// overConstrainedBehaviorNames are the frozen wire spellings.
var overConstrainedBehaviorNames = map[OverConstrainedDimensionBehavior]string{
	OverConstrainedApplyDriven:  "applyDriven",
	OverConstrainedApplyDriving: "applyDriving",
	OverConstrainedPrompt:       "prompt",
}

// String returns the behaviour's wire spelling.
func (b OverConstrainedDimensionBehavior) String() string {
	return enumName(overConstrainedBehaviorNames, b, "enum(?)")
}

// ParseOverConstrainedDimensionBehavior resolves a wire spelling back to its behaviour.
func ParseOverConstrainedDimensionBehavior(s string) (OverConstrainedDimensionBehavior, bool) {
	return enumFromName(overConstrainedBehaviorNames, s)
}
