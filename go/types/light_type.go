// SPDX-License-Identifier: Apache-2.0

package types

// LightTypeEnum is the coordinate space a light lives in. A model-space light is fixed to
// the model; a view-space light follows the camera (a headlight); a ground-plane-space
// light is fixed to the ground/orientation-cube frame. The numeric ids are stable, frozen
// values (52993–52995).
//
// This is the canonical Apache-2.0 definition; the GPL implementation aliases it
// (app.LightTypeEnum) and maps it onto the renderer's lighting rig.
type LightTypeEnum int32

const (
	// ModelSpaceLight is fixed in model space (52993).
	ModelSpaceLight LightTypeEnum = 52993
	// ViewSpaceLight follows the camera — a headlight (52994).
	ViewSpaceLight LightTypeEnum = 52994
	// GroundPlaneSpaceLight is fixed to the ground-plane/ViewCube frame (52995).
	GroundPlaneSpaceLight LightTypeEnum = 52995
)

var lightTypeNames = map[LightTypeEnum]string{
	ModelSpaceLight:       "Model Space",
	ViewSpaceLight:        "View Space",
	GroundPlaneSpaceLight: "Ground Plane Space",
}

// String returns the light type's user-facing name.
func (t LightTypeEnum) String() string {
	return enumName(lightTypeNames, t, "lightType(?)")
}

// IsValid reports whether t is a defined light type.
func (t LightTypeEnum) IsValid() bool {
	return enumValid(lightTypeNames, t)
}

// AllLightTypes returns every defined light type, for building a picker.
func AllLightTypes() []LightTypeEnum {
	return []LightTypeEnum{ModelSpaceLight, ViewSpaceLight, GroundPlaneSpaceLight}
}
