// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati/api/types"

// Light is the in-process contract for one scene light, scoped to the
// properties our renderer consumes. It is read-mostly here; mutation goes through the wire
// methods / the app. The GPL app satisfies it (compile-time asserted there).
//
// Direction and Position are world-space [3]float64 (cm); Direction points from a lit surface
// toward the light. Color is a [types.Rgba]; Intensity scales it. Spot and attenuation getters
// are meaningful only for the matching [types.LightDefinitionTypeEnum].
type Light interface {
	// LightType is the coordinate space the light lives in (model/view/ground-plane).
	LightType() types.LightTypeEnum
	// LightDefinitionType is the emission shape (directional/point/spot).
	LightDefinitionType() types.LightDefinitionTypeEnum
	// On reports whether the light contributes to the scene.
	On() bool
	// Color is the light's color.
	Color() types.Rgba
	// Intensity scales the color (≥ 0).
	Intensity() float64
	// Direction is the unit vector from a lit surface toward the light (directional/spot).
	Direction() [3]float64
	// Position is the light's world-space position (point/spot).
	Position() [3]float64
	// SpotInnerAngle is the spot cone's inner half-angle in radians (full intensity within).
	SpotInnerAngle() float64
	// SpotOuterAngle is the spot cone's outer half-angle in radians (zero intensity beyond).
	SpotOuterAngle() float64
	// Attenuation is the (constant, linear, quadratic) distance falloff (point/spot).
	Attenuation() [3]float64
}
