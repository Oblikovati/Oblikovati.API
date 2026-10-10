// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// The occurrence-pattern suppression enum is FROZEN to the reference ids and wire spellings
// (Oblikovati/Oblikovati#1976).

func TestOccurrencePatternSuppressionFrozenBlock(t *testing.T) {
	want := map[OccurrencePatternSuppression]string{
		118529: "all", 118530: "none", 118531: "some",
	}
	assertFrozenBlock(t, "OccurrencePatternSuppression", want,
		occurrencePatternSuppressionNames, ParseOccurrencePatternSuppression)
}
