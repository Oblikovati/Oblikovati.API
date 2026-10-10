// SPDX-License-Identifier: Apache-2.0

package contract

// EndOfPart is a part's rollback marker (the "end-of-part" marker, #141) — how far down the part
// feature program the model evaluates. Moving it up rolls the part back, suppressing the trailing
// features from evaluation so a user can inspect an earlier state or author a feature mid-history.
// It is the part analogue of [EndOfFeatures] (which is the assembly's marker); the host
// implementation lives in /source (compdef.PartComponentDefinition).
type EndOfPart interface {
	// EndOfPartPosition returns the feature index evaluated up to, or -1 when the whole program
	// is evaluated (the marker is at the end).
	EndOfPartPosition() int
	// IsRolledBack reports whether the marker sits before the end of the program, so some trailing
	// features are currently suppressed.
	IsRolledBack() bool
	// SetEndOfPart moves the marker to position (a negative index restores it to the end).
	SetEndOfPart(position int)
	// RollToEnd moves the marker back to the end, re-including every feature.
	RollToEnd()
}
