// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati/api/types"

// LightingStyle is the in-process contract for a lighting rig,
// scoped to the global controls our renderer consumes plus the discrete [Light]s. The GPL app
// satisfies it (compile-time asserted there).
//
// Ambience, Brightness and the IBL/shadow scalars are unit-free multipliers/intensities in
// [0,1] unless noted; Exposure is in stops. StyleType distinguishes a standard light rig from
// an image-based (HDR) one.
type LightingStyle interface {
	// Name is the style's user-facing label.
	Name() string
	// StyleType is whether this is a standard or image-based style.
	StyleType() types.LightingStyleTypeEnum
	// Ambience is the global ambient fill in [0,1].
	Ambience() float64
	// Brightness is the global light multiplier.
	Brightness() float64
	// Exposure is the tone-map exposure in stops.
	Exposure() float64
	// Lights are the discrete lights in this style (may be empty for a pure IBL style).
	Lights() []Light
	// ImageBasedLightingBrightness scales the IBL contribution in [0,1].
	ImageBasedLightingBrightness() float64
	// ImageBasedLightingRotation spins the environment about vertical, in radians.
	ImageBasedLightingRotation() float64
	// ShadowDensity is the cast-shadow darkness in [0,1].
	ShadowDensity() float64
	// ShadowSoftness is the cast-shadow edge blur in [0,1].
	ShadowSoftness() float64
	// ShadowDirection is the source direction shadows are cast from.
	ShadowDirection() types.ShadowDirectionEnum
}
