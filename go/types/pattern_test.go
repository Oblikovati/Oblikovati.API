// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// The M20·F18 pattern option enums are FROZEN to the reference ids and wire spellings.

func TestPatternSpacingTypeFrozenBlock(t *testing.T) {
	want := map[PatternSpacingType]string{
		33537: "spacing", 33538: "fitted", 33539: "fitToPathLength",
	}
	assertFrozenBlock(t, "PatternSpacingType", want, patternSpacingTypeNames, ParsePatternSpacingType)
}

func TestPatternComputeTypeFrozenBlock(t *testing.T) {
	want := map[PatternComputeType]string{
		47361: "identical", 47362: "adjustToModel", 47363: "optimized",
	}
	assertFrozenBlock(t, "PatternComputeType", want, patternComputeTypeNames, ParsePatternComputeType)
}

func TestPatternOrientationFrozenBlock(t *testing.T) {
	want := map[PatternOrientation]string{
		33793: "identical", 33794: "direction1", 33795: "direction2",
	}
	assertFrozenBlock(t, "PatternOrientation", want, patternOrientationNames, ParsePatternOrientation)
}

func TestPatternPositioningMethodFrozenBlock(t *testing.T) {
	want := map[PatternPositioningMethod]string{
		113409: "fitted", 113410: "incremental",
	}
	assertFrozenBlock(t, "PatternPositioningMethod", want, patternPositioningMethodNames,
		ParsePatternPositioningMethod)
}

func TestPatternBoundaryInclusionFrozenBlock(t *testing.T) {
	want := map[PatternBoundaryInclusion]string{
		128769: "enclosed", 128770: "centroid", 128771: "basePoint",
	}
	assertFrozenBlock(t, "PatternBoundaryInclusion", want, patternBoundaryInclusionNames,
		ParsePatternBoundaryInclusion)
}
