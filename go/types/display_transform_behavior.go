// SPDX-License-Identifier: Apache-2.0

package types

// DisplayTransformBehaviorEnum is how a view-relative graphic (an image billboard, a screen
// label) reacts to the camera: it may always face the viewer (front-facing), keep a constant
// pixel size regardless of zoom (pixel-scaling), both, or neither. The numeric ids are stable,
// frozen values (26113–26116).
//
// Used by the display-options surface and by client-graphics image/text primitives. This is
// the canonical Apache-2.0 definition; the GPL implementation aliases it
// (app.DisplayTransformBehaviorEnum).
type DisplayTransformBehaviorEnum int32

const (
	// FrontFacingBehavior keeps the graphic facing the viewer (billboarding) (26113).
	FrontFacingBehavior DisplayTransformBehaviorEnum = 26113
	// PixelScalingBehavior keeps a constant pixel size regardless of zoom (26114).
	PixelScalingBehavior DisplayTransformBehaviorEnum = 26114
	// FrontFacingAndPixelScalingBehavior applies both front-facing and pixel-scaling (26115).
	FrontFacingAndPixelScalingBehavior DisplayTransformBehaviorEnum = 26115
	// NoTransformBehaviors applies neither — the graphic transforms with the model (26116).
	NoTransformBehaviors DisplayTransformBehaviorEnum = 26116
)

var displayTransformBehaviorNames = map[DisplayTransformBehaviorEnum]string{
	FrontFacingBehavior:                "Front Facing",
	PixelScalingBehavior:               "Pixel Scaling",
	FrontFacingAndPixelScalingBehavior: "Front Facing and Pixel Scaling",
	NoTransformBehaviors:               "None",
}

// String returns the transform-behavior's user-facing name.
func (d DisplayTransformBehaviorEnum) String() string {
	return enumName(displayTransformBehaviorNames, d, "displayTransformBehavior(?)")
}

// IsValid reports whether d is a defined transform behavior.
func (d DisplayTransformBehaviorEnum) IsValid() bool {
	return enumValid(displayTransformBehaviorNames, d)
}

// AllDisplayTransformBehaviors returns every defined transform behavior.
func AllDisplayTransformBehaviors() []DisplayTransformBehaviorEnum {
	return []DisplayTransformBehaviorEnum{
		FrontFacingBehavior, PixelScalingBehavior,
		FrontFacingAndPixelScalingBehavior, NoTransformBehaviors,
	}
}
