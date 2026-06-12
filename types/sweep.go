// SPDX-License-Identifier: Apache-2.0

package types

// Sweep vocabulary (M08-F03 PBI-094, Oblikovati/Oblikovati#314). All numeric
// blocks are FROZEN to the reference enums (SweepDefinitionTypeEnum /
// SweepProfileOrientationEnum / SweepProfileScalingEnum / SweepTypeEnum);
// never renumber. SolidSweepDef is an Oblikovati extension continuing the
// definition-type block — the reference carries solid sweeps as a separate
// definition class with no enum member of its own.

// SweepDefinitionType discriminates the sweep definition union.
type SweepDefinitionType int32

const (
	PathSweepDef                 SweepDefinitionType = 59137
	PathAndGuideRailSweepDef     SweepDefinitionType = 59138
	PathAndGuideSurfaceSweepDef  SweepDefinitionType = 59139
	PathAndSectionTwistsSweepDef SweepDefinitionType = 59140
	// SolidSweepDef sweeps a tool BODY along the path (the reference
	// SolidSweepDefinition) — Oblikovati extension id.
	SolidSweepDef SweepDefinitionType = 59141
)

var sweepDefinitionTypeNames = map[SweepDefinitionType]string{
	PathSweepDef:                 "path",
	PathAndGuideRailSweepDef:     "pathAndGuideRail",
	PathAndGuideSurfaceSweepDef:  "pathAndGuideSurface",
	PathAndSectionTwistsSweepDef: "pathAndSectionTwists",
	SolidSweepDef:                "solid",
}

// String returns the definition type's wire spelling.
func (t SweepDefinitionType) String() string { return enumName(sweepDefinitionTypeNames, t) }

// ParseSweepDefinitionType resolves a wire spelling back to its type.
func ParseSweepDefinitionType(s string) (SweepDefinitionType, bool) {
	return enumFromName(sweepDefinitionTypeNames, s)
}

// SweepProfileOrientation controls how the profile rides the path.
type SweepProfileOrientation int32

const (
	// NormalToPath keeps the profile perpendicular to the path tangent.
	NormalToPath SweepProfileOrientation = 59649
	// ParallelToOriginalProfile translates the profile without rotating it.
	ParallelToOriginalProfile SweepProfileOrientation = 59650
	// AlignToVector keeps the profile normal locked to a fixed vector.
	AlignToVector SweepProfileOrientation = 59651
)

var sweepProfileOrientationNames = map[SweepProfileOrientation]string{
	NormalToPath:              "normalToPath",
	ParallelToOriginalProfile: "parallelToOriginalProfile",
	AlignToVector:             "alignToVector",
}

// String returns the orientation's wire spelling.
func (o SweepProfileOrientation) String() string { return enumName(sweepProfileOrientationNames, o) }

// ParseSweepProfileOrientation resolves a wire spelling back to its orientation.
func ParseSweepProfileOrientation(s string) (SweepProfileOrientation, bool) {
	return enumFromName(sweepProfileOrientationNames, s)
}

// SweepProfileScaling controls how a guide rail scales the profile.
type SweepProfileScaling int32

const (
	// XYProfileScaling scales the whole section with the rail distance.
	XYProfileScaling SweepProfileScaling = 59393
	// XProfileScaling scales only along the rail direction.
	XProfileScaling SweepProfileScaling = 59394
	// NoProfileScaling lets the rail control orientation only.
	NoProfileScaling SweepProfileScaling = 59395
)

var sweepProfileScalingNames = map[SweepProfileScaling]string{
	XYProfileScaling: "xy",
	XProfileScaling:  "x",
	NoProfileScaling: "none",
}

// String returns the scaling's wire spelling.
func (s SweepProfileScaling) String() string { return enumName(sweepProfileScalingNames, s) }

// ParseSweepProfileScaling resolves a wire spelling back to its scaling.
func ParseSweepProfileScaling(s string) (SweepProfileScaling, bool) {
	return enumFromName(sweepProfileScalingNames, s)
}

// SweepType is the placed feature's sweep kind (the SweepFeature property
// mirroring its definition type).
type SweepType int32

const (
	PathSweepType                SweepType = 104449
	PathAndGuideRailSweepType    SweepType = 104450
	PathAndGuideSurfaceSweepType SweepType = 104451
	PathAndSectionTwistSweepType SweepType = 104452
)

var sweepTypeNames = map[SweepType]string{
	PathSweepType:                "path",
	PathAndGuideRailSweepType:    "pathAndGuideRail",
	PathAndGuideSurfaceSweepType: "pathAndGuideSurface",
	PathAndSectionTwistSweepType: "pathAndSectionTwists",
}

// String returns the sweep type's wire spelling.
func (t SweepType) String() string { return enumName(sweepTypeNames, t) }

// ParseSweepType resolves a wire spelling back to its type.
func ParseSweepType(s string) (SweepType, bool) { return enumFromName(sweepTypeNames, s) }
