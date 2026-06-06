// SPDX-License-Identifier: Apache-2.0

package types

// LightingStyleTypeEnum says whether a lighting style is a traditional multi-light rig or an
// image-based (HDR environment) style. The numeric ids are stable, frozen values
// (50750977–50750978).
//
// This is the canonical Apache-2.0 definition; the GPL implementation aliases it
// (app.LightingStyleTypeEnum).
type LightingStyleTypeEnum int32

const (
	// StandardLightingStyle is a traditional rig of discrete lights (50750977).
	StandardLightingStyle LightingStyleTypeEnum = 50750977
	// ImageBasedLightingStyle derives illumination from an HDR environment image (50750978).
	ImageBasedLightingStyle LightingStyleTypeEnum = 50750978
)

var lightingStyleTypeNames = map[LightingStyleTypeEnum]string{
	StandardLightingStyle:   "Standard",
	ImageBasedLightingStyle: "Image Based",
}

// String returns the lighting-style type's user-facing name.
func (t LightingStyleTypeEnum) String() string {
	if name, ok := lightingStyleTypeNames[t]; ok {
		return name
	}
	return "lightingStyleType(?)"
}

// IsValid reports whether t is a defined lighting-style type.
func (t LightingStyleTypeEnum) IsValid() bool {
	_, ok := lightingStyleTypeNames[t]
	return ok
}

// AllLightingStyleTypes returns every defined lighting-style type.
func AllLightingStyleTypes() []LightingStyleTypeEnum {
	return []LightingStyleTypeEnum{StandardLightingStyle, ImageBasedLightingStyle}
}
