// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// The M08 sweep enum blocks are FROZEN to the reference ids and wire
// spellings (Oblikovati/Oblikovati#314); SolidSweepDef 59141 is the
// documented Oblikovati extension continuing the definition-type block.

func TestSweepDefinitionTypeFrozenBlock(t *testing.T) {
	want := map[SweepDefinitionType]string{
		59137: "path", 59138: "pathAndGuideRail", 59139: "pathAndGuideSurface",
		59140: "pathAndSectionTwists", 59141: "solid",
	}
	assertFrozenBlock(t, "SweepDefinitionType", want, sweepDefinitionTypeNames, ParseSweepDefinitionType)
}

func TestSweepProfileOrientationFrozenBlock(t *testing.T) {
	want := map[SweepProfileOrientation]string{
		59649: "normalToPath", 59650: "parallelToOriginalProfile", 59651: "alignToVector",
	}
	assertFrozenBlock(t, "SweepProfileOrientation", want, sweepProfileOrientationNames, ParseSweepProfileOrientation)
}

func TestSweepProfileScalingFrozenBlock(t *testing.T) {
	want := map[SweepProfileScaling]string{59393: "xy", 59394: "x", 59395: "none"}
	assertFrozenBlock(t, "SweepProfileScaling", want, sweepProfileScalingNames, ParseSweepProfileScaling)
}

func TestSweepTypeFrozenBlock(t *testing.T) {
	want := map[SweepType]string{
		104449: "path", 104450: "pathAndGuideRail", 104451: "pathAndGuideSurface", 104452: "pathAndSectionTwists",
	}
	assertFrozenBlock(t, "SweepType", want, sweepTypeNames, ParseSweepType)
}
