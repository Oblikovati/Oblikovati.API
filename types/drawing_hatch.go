// SPDX-License-Identifier: Apache-2.0

package types

// Drawing hatch value types (M14-F08, #638). A hatch region fills a closed boundary with a line
// family — a hatch pattern. These are the canonical Apache-2.0 definitions; the GPL model generates
// the fill lines from the pattern's angle/spacing.

// HatchPattern is a built-in hatch line family. The zero value is HatchGeneral (single 45° lines).
type HatchPattern int32

const (
	// HatchGeneral is the general-purpose single hatch: parallel lines at 45°.
	HatchGeneral HatchPattern = iota
	// HatchCross is cross-hatch: two perpendicular line families at ±45°.
	HatchCross
	// HatchANSI31 is the ANSI iron/general pattern: parallel lines at 45° (a closer spacing).
	HatchANSI31
)

var hatchPatternNames = map[HatchPattern]string{
	HatchGeneral: "general",
	HatchCross:   "cross",
	HatchANSI31:  "ansi31",
}

// String returns the pattern's wire spelling ("general", "cross", "ansi31").
func (p HatchPattern) String() string { return enumName(hatchPatternNames, p) }

// ParseHatchPattern resolves a wire spelling back to its pattern.
//
//	p, ok := types.ParseHatchPattern("cross") // HatchCross, true
func ParseHatchPattern(s string) (HatchPattern, bool) {
	return enumFromName(hatchPatternNames, s)
}
