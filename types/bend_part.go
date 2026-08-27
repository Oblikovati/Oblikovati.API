// SPDX-License-Identifier: Apache-2.0

package types

// BendPartType discriminates how a Bend Part feature's geometry is driven (parity:
// BendPartTypeEnum). A bend wraps a solid body around a sketch bend line; exactly two
// of {radius, angle, arc length} are supplied and the third is derived, so the type
// names which pair the definition carries. The numeric values are frozen at the
// reference API's ids and must never be renumbered.
//
// This is the canonical, Apache-2.0 definition; the GPL implementation aliases it
// (ADR-0018) and maps it onto the bend kernel op.
type BendPartType int32

const (
	// ArcLengthAndAngleBend drives the bend by its arc length and bend angle (radius derived).
	ArcLengthAndAngleBend BendPartType = 83457
	// RadiusAndAngleBend drives the bend by its bend radius and angle (arc length derived).
	RadiusAndAngleBend BendPartType = 83458
	// RadiusAndArcLengthBend drives the bend by its radius and arc length (angle derived).
	RadiusAndArcLengthBend BendPartType = 83459
)

var bendPartTypeNames = map[BendPartType]string{
	ArcLengthAndAngleBend:  "arcLengthAndAngle",
	RadiusAndAngleBend:     "radiusAndAngle",
	RadiusAndArcLengthBend: "radiusAndArcLength",
}

// String returns the bend type's wire spelling.
func (t BendPartType) String() string { return enumName(bendPartTypeNames, t, "enum(?)") }

// ParseBendPartType resolves a wire spelling back to its bend type.
func ParseBendPartType(s string) (BendPartType, bool) { return enumFromName(bendPartTypeNames, s) }
