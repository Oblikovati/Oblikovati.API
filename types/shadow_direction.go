// SPDX-License-Identifier: Apache-2.0

package types

// ShadowDirectionEnum is the light source that casts the scene's shadows — Inventor's
// ShadowDirectionEnum: a fixed 45°/overhead direction, one of the four standard lights, or
// the environment image. The numeric ids are Inventor's exact frozen values (92161–92168).
//
// This is the canonical Apache-2.0 definition; the GPL implementation aliases it
// (app.ShadowDirectionEnum).
type ShadowDirectionEnum int32

const (
	// FortyFiveDegreesLeftShadow casts from 45° to the left (92161).
	FortyFiveDegreesLeftShadow ShadowDirectionEnum = 92161
	// FortyFiveDegreesRightShadow casts from 45° to the right (92162).
	FortyFiveDegreesRightShadow ShadowDirectionEnum = 92162
	// AboveShadow casts straight down from above (92163).
	AboveShadow ShadowDirectionEnum = 92163
	// LightOneShadow follows light 1 (92164).
	LightOneShadow ShadowDirectionEnum = 92164
	// LightTwoShadow follows light 2 (92165).
	LightTwoShadow ShadowDirectionEnum = 92165
	// LightThreeShadow follows light 3 (92166).
	LightThreeShadow ShadowDirectionEnum = 92166
	// LightFourShadow follows light 4 (92167).
	LightFourShadow ShadowDirectionEnum = 92167
	// EnvironmentShadow follows the environment image's dominant light (92168).
	EnvironmentShadow ShadowDirectionEnum = 92168
)

var shadowDirectionNames = map[ShadowDirectionEnum]string{
	FortyFiveDegreesLeftShadow:  "45° Left",
	FortyFiveDegreesRightShadow: "45° Right",
	AboveShadow:                 "Above",
	LightOneShadow:              "Light 1",
	LightTwoShadow:              "Light 2",
	LightThreeShadow:            "Light 3",
	LightFourShadow:             "Light 4",
	EnvironmentShadow:           "Environment",
}

// String returns the shadow direction's user-facing name.
func (d ShadowDirectionEnum) String() string {
	if name, ok := shadowDirectionNames[d]; ok {
		return name
	}
	return "shadowDirection(?)"
}

// IsValid reports whether d is a defined shadow direction.
func (d ShadowDirectionEnum) IsValid() bool {
	_, ok := shadowDirectionNames[d]
	return ok
}

// AllShadowDirections returns every defined shadow direction, in picker order.
func AllShadowDirections() []ShadowDirectionEnum {
	return []ShadowDirectionEnum{
		FortyFiveDegreesLeftShadow, FortyFiveDegreesRightShadow, AboveShadow,
		LightOneShadow, LightTwoShadow, LightThreeShadow, LightFourShadow, EnvironmentShadow,
	}
}
