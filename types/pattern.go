// SPDX-License-Identifier: Apache-2.0

package types

// Pattern definition vocabulary (M20·F18 parity). Rectangular and circular feature
// patterns carry option surfaces beyond a count and spacing: how spacing is interpreted,
// how each occurrence is computed and oriented, how the run is positioned, and how a
// bounding profile clips the run. Every numeric block below is FROZEN to the reference
// enum ids and must never be renumbered.
//
// These are the canonical, Apache-2.0 definitions; the GPL implementation aliases them
// (ADR-0018) and maps each onto the pattern feature definition.

// PatternSpacingType interprets a pattern direction's spacing value (parity:
// PatternSpacingTypeEnum).
type PatternSpacingType int32

const (
	// SpacingBetween places occurrences a fixed distance apart (count + spacing).
	SpacingBetween PatternSpacingType = 33537
	// SpacingFitted fits all occurrences within a span (count + span).
	SpacingFitted PatternSpacingType = 33538
	// SpacingFitToPathLength fits occurrences along a curve's length (spacing + span).
	SpacingFitToPathLength PatternSpacingType = 33539
)

var patternSpacingTypeNames = map[PatternSpacingType]string{
	SpacingBetween:         "spacing",
	SpacingFitted:          "fitted",
	SpacingFitToPathLength: "fitToPathLength",
}

// String returns the spacing type's wire spelling.
func (t PatternSpacingType) String() string { return enumName(patternSpacingTypeNames, t, "enum(?)") }

// ParsePatternSpacingType resolves a wire spelling back to its spacing type.
func ParsePatternSpacingType(s string) (PatternSpacingType, bool) {
	return enumFromName(patternSpacingTypeNames, s)
}

// PatternComputeType selects how each occurrence's geometry is computed (parity:
// PatternComputeTypeEnum).
type PatternComputeType int32

const (
	// ComputeIdentical creates every occurrence identically to the seed (fastest).
	ComputeIdentical PatternComputeType = 47361
	// ComputeAdjustToModel recomputes each occurrence against the model it lands on.
	ComputeAdjustToModel PatternComputeType = 47362
	// ComputeOptimized reuses one body transformed into place (lightest representation).
	ComputeOptimized PatternComputeType = 47363
)

var patternComputeTypeNames = map[PatternComputeType]string{
	ComputeIdentical:     "identical",
	ComputeAdjustToModel: "adjustToModel",
	ComputeOptimized:     "optimized",
}

// String returns the compute type's wire spelling.
func (t PatternComputeType) String() string { return enumName(patternComputeTypeNames, t, "enum(?)") }

// ParsePatternComputeType resolves a wire spelling back to its compute type.
func ParsePatternComputeType(s string) (PatternComputeType, bool) {
	return enumFromName(patternComputeTypeNames, s)
}

// PatternOrientation selects how each occurrence is rotated relative to the seed
// (parity: PatternOrientationEnum).
type PatternOrientation int32

const (
	// OrientIdentical keeps every occurrence in the seed's orientation.
	OrientIdentical PatternOrientation = 33793
	// OrientToDirection1 rotates occurrences to follow the first pattern direction.
	OrientToDirection1 PatternOrientation = 33794
	// OrientToDirection2 rotates occurrences to follow the second pattern direction.
	OrientToDirection2 PatternOrientation = 33795
)

var patternOrientationNames = map[PatternOrientation]string{
	OrientIdentical:    "identical",
	OrientToDirection1: "direction1",
	OrientToDirection2: "direction2",
}

// String returns the orientation's wire spelling.
func (t PatternOrientation) String() string { return enumName(patternOrientationNames, t, "enum(?)") }

// ParsePatternOrientation resolves a wire spelling back to its orientation.
func ParsePatternOrientation(s string) (PatternOrientation, bool) {
	return enumFromName(patternOrientationNames, s)
}

// PatternPositioningMethod selects how occurrences are positioned along a direction
// (parity: PatternPositioningMethodEnum).
type PatternPositioningMethod int32

const (
	// PositionFitted distributes occurrences evenly across the span.
	PositionFitted PatternPositioningMethod = 113409
	// PositionIncremental steps occurrences by the spacing increment.
	PositionIncremental PatternPositioningMethod = 113410
)

var patternPositioningMethodNames = map[PatternPositioningMethod]string{
	PositionFitted:      "fitted",
	PositionIncremental: "incremental",
}

// String returns the positioning method's wire spelling.
func (t PatternPositioningMethod) String() string {
	return enumName(patternPositioningMethodNames, t, "enum(?)")
}

// ParsePatternPositioningMethod resolves a wire spelling back to its positioning method.
func ParsePatternPositioningMethod(s string) (PatternPositioningMethod, bool) {
	return enumFromName(patternPositioningMethodNames, s)
}

// PatternBoundaryInclusion decides whether an occurrence inside/outside a bounding
// profile is kept, by which point of the occurrence is tested (parity:
// PatternBoundaryInclusionEnum).
type PatternBoundaryInclusion int32

const (
	// IncludeEnclosedGeometry keeps an occurrence whose whole geometry is enclosed.
	IncludeEnclosedGeometry PatternBoundaryInclusion = 128769
	// IncludeByCentroid keeps an occurrence whose centroid is enclosed.
	IncludeByCentroid PatternBoundaryInclusion = 128770
	// IncludeByBasePoint keeps an occurrence whose base point is enclosed.
	IncludeByBasePoint PatternBoundaryInclusion = 128771
)

var patternBoundaryInclusionNames = map[PatternBoundaryInclusion]string{
	IncludeEnclosedGeometry: "enclosed",
	IncludeByCentroid:       "centroid",
	IncludeByBasePoint:      "basePoint",
}

// String returns the boundary-inclusion's wire spelling.
func (t PatternBoundaryInclusion) String() string {
	return enumName(patternBoundaryInclusionNames, t, "enum(?)")
}

// ParsePatternBoundaryInclusion resolves a wire spelling back to its inclusion rule.
func ParsePatternBoundaryInclusion(s string) (PatternBoundaryInclusion, bool) {
	return enumFromName(patternBoundaryInclusionNames, s)
}
