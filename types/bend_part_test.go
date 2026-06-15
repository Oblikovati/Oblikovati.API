// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestBendPartTypeFrozenBlock pins the reference ids and wire spellings.
func TestBendPartTypeFrozenBlock(t *testing.T) {
	want := map[BendPartType]string{
		83457: "arcLengthAndAngle", 83458: "radiusAndAngle", 83459: "radiusAndArcLength",
	}
	assertFrozenBlock(t, "BendPartType", want, bendPartTypeNames, ParseBendPartType)
}
