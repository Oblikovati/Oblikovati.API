// SPDX-License-Identifier: Apache-2.0

package types

// Occurrence-pattern vocabulary (M11-F04 parity, Oblikovati/Oblikovati#1976). A persistent
// assembly occurrence pattern can be suppressed as a whole or element by element; the
// pattern-level state reports whether none, some, or all of its elements are suppressed. The
// numeric ids are FROZEN to the reference enum and must never be renumbered.
//
// This is the canonical, Apache-2.0 definition; the GPL implementation aliases it (ADR-0018).

// OccurrencePatternSuppression reports how much of an occurrence pattern is suppressed
// (parity: OccurrencePatternSuppressionEnum). It is a derived read state, not a request:
// suppressing the pattern or its elements moves it between these.
type OccurrencePatternSuppression int32

const (
	// AllElementsSuppressed means every element of the pattern is suppressed.
	AllElementsSuppressed OccurrencePatternSuppression = 118529
	// NoneSuppressed means no element of the pattern is suppressed.
	NoneSuppressed OccurrencePatternSuppression = 118530
	// SomeElementsSuppressed means the pattern is partly suppressed (mixed elements).
	SomeElementsSuppressed OccurrencePatternSuppression = 118531
)

var occurrencePatternSuppressionNames = map[OccurrencePatternSuppression]string{
	AllElementsSuppressed:  "all",
	NoneSuppressed:         "none",
	SomeElementsSuppressed: "some",
}

// String returns the suppression state's wire spelling.
func (s OccurrencePatternSuppression) String() string {
	return enumName(occurrencePatternSuppressionNames, s)
}

// ParseOccurrencePatternSuppression resolves a wire spelling back to its suppression state.
func ParseOccurrencePatternSuppression(s string) (OccurrencePatternSuppression, bool) {
	return enumFromName(occurrencePatternSuppressionNames, s)
}
