// SPDX-License-Identifier: Apache-2.0

package types

// GroundShadowEnum is how the scene casts shadows onto the ground plane — Inventor's
// GroundShadowEnum: none, a standard ground shadow, or an X-ray (see-through) ground shadow.
// The numeric ids are Inventor's exact frozen values (69121–69123).
//
// This is the canonical Apache-2.0 definition; the GPL implementation aliases it
// (app.GroundShadowEnum).
type GroundShadowEnum int32

const (
	// NoGroundShadow casts no ground shadow (69121).
	NoGroundShadow GroundShadowEnum = 69121
	// GroundShadow casts a standard ground shadow (69122).
	GroundShadow GroundShadowEnum = 69122
	// XRayGroundShadow casts a see-through (X-ray) ground shadow (69123).
	XRayGroundShadow GroundShadowEnum = 69123
)

var groundShadowNames = map[GroundShadowEnum]string{
	NoGroundShadow:   "None",
	GroundShadow:     "Ground Shadow",
	XRayGroundShadow: "X-Ray Ground Shadow",
}

// String returns the ground-shadow mode's user-facing name.
func (g GroundShadowEnum) String() string {
	if name, ok := groundShadowNames[g]; ok {
		return name
	}
	return "groundShadow(?)"
}

// IsValid reports whether g is a defined ground-shadow mode.
func (g GroundShadowEnum) IsValid() bool {
	_, ok := groundShadowNames[g]
	return ok
}

// AllGroundShadows returns every defined ground-shadow mode, in picker order.
func AllGroundShadows() []GroundShadowEnum {
	return []GroundShadowEnum{NoGroundShadow, GroundShadow, XRayGroundShadow}
}
