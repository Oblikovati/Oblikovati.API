// SPDX-License-Identifier: Apache-2.0

package types

// LightDefinitionTypeEnum is the emission shape of a light: a directional (parallel-ray)
// sun, an omnidirectional point, or a cone spotlight. The numeric ids are stable, frozen
// values (53249–53251).
//
// This is the canonical Apache-2.0 definition; the GPL implementation aliases it
// (app.LightDefinitionTypeEnum) and maps it onto the renderer's LightKind.
type LightDefinitionTypeEnum int32

const (
	// DirectionalLight emits parallel rays (a sun); only its direction matters (53249).
	DirectionalLight LightDefinitionTypeEnum = 53249
	// PointLight emits from a position in all directions, with distance attenuation (53250).
	PointLight LightDefinitionTypeEnum = 53250
	// SpotLight emits from a position within a cone about its direction (53251).
	SpotLight LightDefinitionTypeEnum = 53251
)

var lightDefinitionTypeNames = map[LightDefinitionTypeEnum]string{
	DirectionalLight: "Directional",
	PointLight:       "Point",
	SpotLight:        "Spot",
}

// String returns the definition type's user-facing name.
func (t LightDefinitionTypeEnum) String() string {
	if name, ok := lightDefinitionTypeNames[t]; ok {
		return name
	}
	return "lightDefinitionType(?)"
}

// IsValid reports whether t is a defined light-definition type.
func (t LightDefinitionTypeEnum) IsValid() bool {
	_, ok := lightDefinitionTypeNames[t]
	return ok
}

// AllLightDefinitionTypes returns every defined light-definition type, for building a picker.
func AllLightDefinitionTypes() []LightDefinitionTypeEnum {
	return []LightDefinitionTypeEnum{DirectionalLight, PointLight, SpotLight}
}
