// SPDX-License-Identifier: Apache-2.0

package types

// SplineFitMethod selects the parameterization an interpolation spline uses
// to fit its points (M06-F11, Oblikovati/Oblikovati#626). The default is
// SplineFitSmooth, which matches the behavior shipped before the field existed.
//
// The values are a frozen block matching the reference API's spline-fit enum;
// never renumber them.
type SplineFitMethod int32

const (
	// SplineFitSmooth uses centripetal parameterization (the default).
	SplineFitSmooth SplineFitMethod = 26369
	// SplineFitSweet uses minimum-energy fitting.
	SplineFitSweet SplineFitMethod = 26370
	// SplineFitChord uses chord-length parameterization — the reference
	// API's third, legacy-drafting-compatible fit method.
	SplineFitChord SplineFitMethod = 26371
)

// splineFitMethodNames are the frozen wire spellings.
var splineFitMethodNames = map[SplineFitMethod]string{
	SplineFitSmooth: "smooth",
	SplineFitSweet:  "sweet",
	SplineFitChord:  "chord",
}

// String returns the fit method's wire spelling.
func (m SplineFitMethod) String() string { return enumName(splineFitMethodNames, m, "enum(?)") }

// ParseSplineFitMethod resolves a wire spelling back to its method.
func ParseSplineFitMethod(s string) (SplineFitMethod, bool) {
	return enumFromName(splineFitMethodNames, s)
}
