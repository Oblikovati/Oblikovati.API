// SPDX-License-Identifier: Apache-2.0

package contract

// EndOfFeatures is the assembly's rollback marker — how far down the assembly feature
// program the model evaluates. Moving it up rolls the assembly back, suppressing the
// trailing features from evaluation (M11-F08, Oblikovati/Oblikovati#633/#725). The host
// implementation lives in /source (compdef.AssemblyFeatures).
type EndOfFeatures interface {
	// EndOfFeaturesPosition returns the feature index evaluated up to, or -1 when the
	// whole program is evaluated (the marker is at the end).
	EndOfFeaturesPosition() int
	// IsRolledBack reports whether the marker sits before the end of the program.
	IsRolledBack() bool
	// SetEndOfFeatures moves the marker to position (negative restores it to the end).
	SetEndOfFeatures(position int)
	// RollToEnd moves the marker back to the end, re-including every feature.
	RollToEnd()
}

// AssemblyFeatures is the assembly's machining-feature program — the features authored
// in the assembly that cut/modify placed component geometry in place (M11-F08). It is
// the assembly analogue of the part feature program, adding the end-of-features
// rollback marker and batch suppression over the scalar surface a consumer reads.
type AssemblyFeatures interface {
	EndOfFeatures
	// Count returns the number of features in the program.
	Count() int
	// SuppressFeatures suppresses every named feature in one batch.
	SuppressFeatures(ids ...uint64)
	// UnsuppressFeatures clears suppression on every named feature in one batch.
	UnsuppressFeatures(ids ...uint64)
}
